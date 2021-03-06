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

// DirectoryServiceConditionalForwarderLister helps list DirectoryServiceConditionalForwarders.
type DirectoryServiceConditionalForwarderLister interface {
	// List lists all DirectoryServiceConditionalForwarders in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceConditionalForwarder, err error)
	// DirectoryServiceConditionalForwarders returns an object that can list and get DirectoryServiceConditionalForwarders.
	DirectoryServiceConditionalForwarders(namespace string) DirectoryServiceConditionalForwarderNamespaceLister
	DirectoryServiceConditionalForwarderListerExpansion
}

// directoryServiceConditionalForwarderLister implements the DirectoryServiceConditionalForwarderLister interface.
type directoryServiceConditionalForwarderLister struct {
	indexer cache.Indexer
}

// NewDirectoryServiceConditionalForwarderLister returns a new DirectoryServiceConditionalForwarderLister.
func NewDirectoryServiceConditionalForwarderLister(indexer cache.Indexer) DirectoryServiceConditionalForwarderLister {
	return &directoryServiceConditionalForwarderLister{indexer: indexer}
}

// List lists all DirectoryServiceConditionalForwarders in the indexer.
func (s *directoryServiceConditionalForwarderLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceConditionalForwarder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryServiceConditionalForwarder))
	})
	return ret, err
}

// DirectoryServiceConditionalForwarders returns an object that can list and get DirectoryServiceConditionalForwarders.
func (s *directoryServiceConditionalForwarderLister) DirectoryServiceConditionalForwarders(namespace string) DirectoryServiceConditionalForwarderNamespaceLister {
	return directoryServiceConditionalForwarderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectoryServiceConditionalForwarderNamespaceLister helps list and get DirectoryServiceConditionalForwarders.
type DirectoryServiceConditionalForwarderNamespaceLister interface {
	// List lists all DirectoryServiceConditionalForwarders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceConditionalForwarder, err error)
	// Get retrieves the DirectoryServiceConditionalForwarder from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DirectoryServiceConditionalForwarder, error)
	DirectoryServiceConditionalForwarderNamespaceListerExpansion
}

// directoryServiceConditionalForwarderNamespaceLister implements the DirectoryServiceConditionalForwarderNamespaceLister
// interface.
type directoryServiceConditionalForwarderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DirectoryServiceConditionalForwarders in the indexer for a given namespace.
func (s directoryServiceConditionalForwarderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryServiceConditionalForwarder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryServiceConditionalForwarder))
	})
	return ret, err
}

// Get retrieves the DirectoryServiceConditionalForwarder from the indexer for a given namespace and name.
func (s directoryServiceConditionalForwarderNamespaceLister) Get(name string) (*v1alpha1.DirectoryServiceConditionalForwarder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("directoryserviceconditionalforwarder"), name)
	}
	return obj.(*v1alpha1.DirectoryServiceConditionalForwarder), nil
}
