// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/kyma-project/kyma/components/application-registry/pkg/apis/istio/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthorizationPolicies implements AuthorizationPolicyInterface
type FakeAuthorizationPolicies struct {
	Fake *FakeIstioV1alpha2
	ns   string
}

var authorizationpoliciesResource = schema.GroupVersionResource{Group: "istio", Version: "v1alpha2", Resource: "authorizationpolicies"}

var authorizationpoliciesKind = schema.GroupVersionKind{Group: "istio", Version: "v1alpha2", Kind: "AuthorizationPolicy"}

// Get takes name of the authorizationPolicy, and returns the corresponding authorizationPolicy object, and an error if there is any.
func (c *FakeAuthorizationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha2.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(authorizationpoliciesResource, c.ns, name), &v1alpha2.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AuthorizationPolicy), err
}

// List takes label and field selectors, and returns the list of AuthorizationPolicies that match those selectors.
func (c *FakeAuthorizationPolicies) List(opts v1.ListOptions) (result *v1alpha2.AuthorizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(authorizationpoliciesResource, authorizationpoliciesKind, c.ns, opts), &v1alpha2.AuthorizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AuthorizationPolicyList{ListMeta: obj.(*v1alpha2.AuthorizationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha2.AuthorizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authorizationPolicies.
func (c *FakeAuthorizationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(authorizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a authorizationPolicy and creates it.  Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *FakeAuthorizationPolicies) Create(authorizationPolicy *v1alpha2.AuthorizationPolicy) (result *v1alpha2.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(authorizationpoliciesResource, c.ns, authorizationPolicy), &v1alpha2.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AuthorizationPolicy), err
}

// Update takes the representation of a authorizationPolicy and updates it. Returns the server's representation of the authorizationPolicy, and an error, if there is any.
func (c *FakeAuthorizationPolicies) Update(authorizationPolicy *v1alpha2.AuthorizationPolicy) (result *v1alpha2.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(authorizationpoliciesResource, c.ns, authorizationPolicy), &v1alpha2.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AuthorizationPolicy), err
}

// Delete takes name of the authorizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAuthorizationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(authorizationpoliciesResource, c.ns, name), &v1alpha2.AuthorizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthorizationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(authorizationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.AuthorizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched authorizationPolicy.
func (c *FakeAuthorizationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(authorizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha2.AuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AuthorizationPolicy), err
}
