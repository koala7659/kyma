package gateway

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/kyma-project/kyma/tests/application-gateway-tests/test/gateway/mock"

	"github.com/kyma-project/kyma/tests/application-gateway-tests/test/gateway/legacy"

	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/stretchr/testify/assert"

	"github.com/kyma-project/kyma/tests/application-gateway-tests/test/tools"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

const (
	legacyDefaultCheckInterval           = 2 * time.Second
	legacyAppGatewayHealthCheckTimeout   = 15 * time.Second
	legacyAccessServiceConnectionTimeout = 60 * time.Second
	legacyApiServerAccessTimeout         = 60 * time.Second
	legacyDnsWaitTime                    = 30 * time.Second

	mockServiceNameFormat     = "%s-gateway-test-mock-service"
	testExecutorPodNameFormat = "%s-tests-test-executor"
)

type updatePodFunc func(pod *v1.Pod)

type LegacyTestSuite struct {
	httpClient      *http.Client
	k8sClient       *kubernetes.Clientset
	podClient       corev1.PodInterface
	serviceClient   corev1.ServiceInterface
	config          legacy.LegacyModeTestConfig
	appMockServer   *mock.AppMockServer
	mockServiceName string
}

func NewLegacyTestSuite(t *testing.T) *LegacyTestSuite {
	config, err := legacy.ReadConfig()
	require.NoError(t, err)

	if config.GatewayLegacyMode == false {
		return nil
	}

	k8sConfig, err := restclient.InClusterConfig()
	require.NoError(t, err)

	coreClientset, err := kubernetes.NewForConfig(k8sConfig)
	require.NoError(t, err)

	appMockServer := mock.NewAppMockServerForLegacy(config.MockServicePort)

	return &LegacyTestSuite{
		httpClient:      &http.Client{},
		k8sClient:       coreClientset,
		podClient:       coreClientset.CoreV1().Pods(config.Namespace),
		serviceClient:   coreClientset.CoreV1().Services(config.Namespace),
		config:          config,
		appMockServer:   appMockServer,
		mockServiceName: fmt.Sprintf(mockServiceNameFormat, config.Application),
	}
}

func (ts *LegacyTestSuite) Setup(t *testing.T) {
	ts.WaitForAccessToAPIServer(t)

	ts.appMockServer.Start()
	ts.createMockService(t)

	ts.CheckApplicationGatewayHealth(t)
}

func (ts *LegacyTestSuite) Cleanup(t *testing.T) {
	t.Log("Calling cleanup")

	err := ts.appMockServer.Kill()
	assert.NoError(t, err)

	ts.deleteMockService(t)
}

func (ts *LegacyTestSuite) ApplicationName() string {
	return ts.config.Application
}

// WaitForAccessToAPIServer waits for access to API Server which might be delayed by initialization of Istio sidecar
func (ts *LegacyTestSuite) WaitForAccessToAPIServer(t *testing.T) {
	err := tools.WaitForFunction(legacyDefaultCheckInterval, legacyApiServerAccessTimeout, func() bool {
		t.Log("Trying to access API Server...")
		_, err := ts.k8sClient.ServerVersion()
		if err != nil {
			t.Log(err.Error())
			return false
		}

		return true
	})

	require.NoError(t, err)
}

func (ts *LegacyTestSuite) CheckApplicationGatewayHealth(t *testing.T) {
	t.Log("Checking application gateway health...")

	err := tools.WaitForFunction(legacyDefaultCheckInterval, legacyApiServerAccessTimeout, func() bool {
		req, err := http.NewRequest(http.MethodGet, ts.proxyURL()+"/v1/health", nil)
		if err != nil {
			return false
		}

		res, err := ts.httpClient.Do(req)
		if err != nil {
			return false
		}

		if res.StatusCode != http.StatusOK {
			return false
		}

		return true
	})

	require.NoError(t, err, "Failed to check health of Application Gateway.")
}

func (ts *LegacyTestSuite) CallAccessService(t *testing.T, apiId, path string) *http.Response {
	url := fmt.Sprintf("http://%s-%s/%s", ts.config.Application, apiId, path)

	t.Log("Waiting for DNS in Istio Proxy...")
	// Wait for Istio Pilot to propagate DNS
	time.Sleep(legacyDnsWaitTime)

	var resp *http.Response

	err := tools.WaitForFunction(legacyDefaultCheckInterval, legacyAccessServiceConnectionTimeout, func() bool {
		t.Logf("Accessing proxy at: %s", url)
		var err error

		resp, err = http.Get(url)
		if err != nil {
			t.Logf("Access service not ready: %s", err.Error())
			return false
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusServiceUnavailable {
			t.Logf("Invalid response from access service, status: %d.", resp.StatusCode)
			bytes, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			t.Log(string(bytes))
			t.Logf("Access service is not ready. Retrying.")
			return false
		}

		return true
	})
	require.NoError(t, err)

	return resp
}

func (ts *LegacyTestSuite) proxyURL() string {
	return fmt.Sprintf("http://%s-application-gateway:8081", ts.config.Application)
}

func (ts *LegacyTestSuite) GetMockServiceURL() string {
	return fmt.Sprintf("http://%s:%d", ts.mockServiceName, ts.config.MockServicePort)
}

func (ts *LegacyTestSuite) createMockService(t *testing.T) {
	selectors := map[string]string{
		ts.config.MockSelectorKey: ts.config.MockSelectorValue,
	}

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ts.mockServiceName,
			Namespace: ts.config.Namespace,
			Labels:    selectors,
		},
		Spec: v1.ServiceSpec{
			Selector: selectors,
			Ports: []v1.ServicePort{
				{Port: ts.config.MockServicePort, Name: "http-port"},
			},
		},
	}

	_, err := ts.serviceClient.Create(context.Background(), service, metav1.CreateOptions{})
	require.NoError(t, err)
}

func (ts *LegacyTestSuite) deleteMockService(t *testing.T) {
	err := ts.serviceClient.Delete(context.Background(), ts.mockServiceName, metav1.DeleteOptions{})
	assert.NoError(t, err)
}

// not needed anymore - istio denier functionality is removed
//func (ts *LegacyTestSuite) AddDenierLabel(t *testing.T, apiId string) {
//	podName := fmt.Sprintf(testExecutorPodNameFormat, ts.config.Application)
//	serviceName := fmt.Sprintf("%s-%s", ts.config.Application, apiId)
//
//	updateFunc := func(pod *v1.Pod) {
//		if pod.Labels == nil {
//			pod.Labels = map[string]string{}
//		}
//
//		pod.Labels[serviceName] = "true"
//	}
//
//	err := ts.updatePod(podName, updateFunc)
//	require.NoError(t, err)
//}
//
//func (ts *LegacyTestSuite) updatePod(podName string, updateFunc updatePodFunc) error {
//	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
//		newPod, err := ts.podClient.Get(podName, metav1.GetOptions{})
//		if err != nil {
//			return err
//		}
//
//		updateFunc(newPod)
//		_, err = ts.podClient.Update(newPod)
//		return err
//	})
//}
