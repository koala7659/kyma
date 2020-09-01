// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kyma-project/kyma/components/application-registry/pkg/apis/istio/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthorizationPolicyLister helps list AuthorizationPolicies.
type AuthorizationPolicyLister interface {
	// List lists all AuthorizationPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.AuthorizationPolicy, err error)
	// AuthorizationPolicies returns an object that can list and get AuthorizationPolicies.
	AuthorizationPolicies(namespace string) AuthorizationPolicyNamespaceLister
	AuthorizationPolicyListerExpansion
}

// authorizationPolicyLister implements the AuthorizationPolicyLister interface.
type authorizationPolicyLister struct {
	indexer cache.Indexer
}

// NewAuthorizationPolicyLister returns a new AuthorizationPolicyLister.
func NewAuthorizationPolicyLister(indexer cache.Indexer) AuthorizationPolicyLister {
	return &authorizationPolicyLister{indexer: indexer}
}

// List lists all AuthorizationPolicies in the indexer.
func (s *authorizationPolicyLister) List(selector labels.Selector) (ret []*v1alpha2.AuthorizationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.AuthorizationPolicy))
	})
	return ret, err
}

// AuthorizationPolicies returns an object that can list and get AuthorizationPolicies.
func (s *authorizationPolicyLister) AuthorizationPolicies(namespace string) AuthorizationPolicyNamespaceLister {
	return authorizationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuthorizationPolicyNamespaceLister helps list and get AuthorizationPolicies.
type AuthorizationPolicyNamespaceLister interface {
	// List lists all AuthorizationPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.AuthorizationPolicy, err error)
	// Get retrieves the AuthorizationPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.AuthorizationPolicy, error)
	AuthorizationPolicyNamespaceListerExpansion
}

// authorizationPolicyNamespaceLister implements the AuthorizationPolicyNamespaceLister
// interface.
type authorizationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuthorizationPolicies in the indexer for a given namespace.
func (s authorizationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.AuthorizationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.AuthorizationPolicy))
	})
	return ret, err
}

// Get retrieves the AuthorizationPolicy from the indexer for a given namespace and name.
func (s authorizationPolicyNamespaceLister) Get(name string) (*v1alpha2.AuthorizationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("authorizationpolicy"), name)
	}
	return obj.(*v1alpha2.AuthorizationPolicy), nil
}
