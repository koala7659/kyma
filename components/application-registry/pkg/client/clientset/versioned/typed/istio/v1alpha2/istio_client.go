// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kyma-project/kyma/components/application-registry/pkg/apis/istio/v1alpha2"
	"github.com/kyma-project/kyma/components/application-registry/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type IstioV1alpha2Interface interface {
	RESTClient() rest.Interface
	AuthorizationPoliciesGetter
	HandlersGetter
	InstancesGetter
}

// IstioV1alpha2Client is used to interact with features provided by the istio group.
type IstioV1alpha2Client struct {
	restClient rest.Interface
}

func (c *IstioV1alpha2Client) AuthorizationPolicies(namespace string) AuthorizationPolicyInterface {
	return newAuthorizationPolicies(c, namespace)
}

func (c *IstioV1alpha2Client) Handlers(namespace string) HandlerInterface {
	return newHandlers(c, namespace)
}

func (c *IstioV1alpha2Client) Instances(namespace string) InstanceInterface {
	return newInstances(c, namespace)
}

// NewForConfig creates a new IstioV1alpha2Client for the given config.
func NewForConfig(c *rest.Config) (*IstioV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IstioV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new IstioV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IstioV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IstioV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *IstioV1alpha2Client {
	return &IstioV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *IstioV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
