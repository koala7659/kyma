// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/kyma-project/kyma/components/application-registry/pkg/client/clientset/versioned/typed/istio/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIstioV1alpha2 struct {
	*testing.Fake
}

func (c *FakeIstioV1alpha2) AuthorizationPolicies(namespace string) v1alpha2.AuthorizationPolicyInterface {
	return &FakeAuthorizationPolicies{c, namespace}
}

func (c *FakeIstioV1alpha2) Handlers(namespace string) v1alpha2.HandlerInterface {
	return &FakeHandlers{c, namespace}
}

func (c *FakeIstioV1alpha2) Instances(namespace string) v1alpha2.InstanceInterface {
	return &FakeInstances{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIstioV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
