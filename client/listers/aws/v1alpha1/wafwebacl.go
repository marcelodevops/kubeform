/*
Copyright The Kubeform Authors.

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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WafWebACLLister helps list WafWebACLs.
type WafWebACLLister interface {
	// List lists all WafWebACLs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafWebACL, err error)
	// WafWebACLs returns an object that can list and get WafWebACLs.
	WafWebACLs(namespace string) WafWebACLNamespaceLister
	WafWebACLListerExpansion
}

// wafWebACLLister implements the WafWebACLLister interface.
type wafWebACLLister struct {
	indexer cache.Indexer
}

// NewWafWebACLLister returns a new WafWebACLLister.
func NewWafWebACLLister(indexer cache.Indexer) WafWebACLLister {
	return &wafWebACLLister{indexer: indexer}
}

// List lists all WafWebACLs in the indexer.
func (s *wafWebACLLister) List(selector labels.Selector) (ret []*v1alpha1.WafWebACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafWebACL))
	})
	return ret, err
}

// WafWebACLs returns an object that can list and get WafWebACLs.
func (s *wafWebACLLister) WafWebACLs(namespace string) WafWebACLNamespaceLister {
	return wafWebACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafWebACLNamespaceLister helps list and get WafWebACLs.
type WafWebACLNamespaceLister interface {
	// List lists all WafWebACLs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafWebACL, err error)
	// Get retrieves the WafWebACL from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafWebACL, error)
	WafWebACLNamespaceListerExpansion
}

// wafWebACLNamespaceLister implements the WafWebACLNamespaceLister
// interface.
type wafWebACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafWebACLs in the indexer for a given namespace.
func (s wafWebACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafWebACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafWebACL))
	})
	return ret, err
}

// Get retrieves the WafWebACL from the indexer for a given namespace and name.
func (s wafWebACLNamespaceLister) Get(name string) (*v1alpha1.WafWebACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafwebacl"), name)
	}
	return obj.(*v1alpha1.WafWebACL), nil
}
