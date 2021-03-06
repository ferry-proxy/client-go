/*
Copyright 2022 FerryProxy Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/ferryproxy/api/apis/traffic/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoutePolicyLister helps list RoutePolicies.
// All objects returned here must be treated as read-only.
type RoutePolicyLister interface {
	// List lists all RoutePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.RoutePolicy, err error)
	// RoutePolicies returns an object that can list and get RoutePolicies.
	RoutePolicies(namespace string) RoutePolicyNamespaceLister
	RoutePolicyListerExpansion
}

// routePolicyLister implements the RoutePolicyLister interface.
type routePolicyLister struct {
	indexer cache.Indexer
}

// NewRoutePolicyLister returns a new RoutePolicyLister.
func NewRoutePolicyLister(indexer cache.Indexer) RoutePolicyLister {
	return &routePolicyLister{indexer: indexer}
}

// List lists all RoutePolicies in the indexer.
func (s *routePolicyLister) List(selector labels.Selector) (ret []*v1alpha2.RoutePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.RoutePolicy))
	})
	return ret, err
}

// RoutePolicies returns an object that can list and get RoutePolicies.
func (s *routePolicyLister) RoutePolicies(namespace string) RoutePolicyNamespaceLister {
	return routePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RoutePolicyNamespaceLister helps list and get RoutePolicies.
// All objects returned here must be treated as read-only.
type RoutePolicyNamespaceLister interface {
	// List lists all RoutePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.RoutePolicy, err error)
	// Get retrieves the RoutePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.RoutePolicy, error)
	RoutePolicyNamespaceListerExpansion
}

// routePolicyNamespaceLister implements the RoutePolicyNamespaceLister
// interface.
type routePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RoutePolicies in the indexer for a given namespace.
func (s routePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.RoutePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.RoutePolicy))
	})
	return ret, err
}

// Get retrieves the RoutePolicy from the indexer for a given namespace and name.
func (s routePolicyNamespaceLister) Get(name string) (*v1alpha2.RoutePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("routepolicy"), name)
	}
	return obj.(*v1alpha2.RoutePolicy), nil
}
