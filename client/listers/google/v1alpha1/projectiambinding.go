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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectIamBindingLister helps list ProjectIamBindings.
type ProjectIamBindingLister interface {
	// List lists all ProjectIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectIamBinding, err error)
	// ProjectIamBindings returns an object that can list and get ProjectIamBindings.
	ProjectIamBindings(namespace string) ProjectIamBindingNamespaceLister
	ProjectIamBindingListerExpansion
}

// projectIamBindingLister implements the ProjectIamBindingLister interface.
type projectIamBindingLister struct {
	indexer cache.Indexer
}

// NewProjectIamBindingLister returns a new ProjectIamBindingLister.
func NewProjectIamBindingLister(indexer cache.Indexer) ProjectIamBindingLister {
	return &projectIamBindingLister{indexer: indexer}
}

// List lists all ProjectIamBindings in the indexer.
func (s *projectIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectIamBinding))
	})
	return ret, err
}

// ProjectIamBindings returns an object that can list and get ProjectIamBindings.
func (s *projectIamBindingLister) ProjectIamBindings(namespace string) ProjectIamBindingNamespaceLister {
	return projectIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectIamBindingNamespaceLister helps list and get ProjectIamBindings.
type ProjectIamBindingNamespaceLister interface {
	// List lists all ProjectIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectIamBinding, err error)
	// Get retrieves the ProjectIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ProjectIamBinding, error)
	ProjectIamBindingNamespaceListerExpansion
}

// projectIamBindingNamespaceLister implements the ProjectIamBindingNamespaceLister
// interface.
type projectIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectIamBindings in the indexer for a given namespace.
func (s projectIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectIamBinding))
	})
	return ret, err
}

// Get retrieves the ProjectIamBinding from the indexer for a given namespace and name.
func (s projectIamBindingNamespaceLister) Get(name string) (*v1alpha1.ProjectIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("projectiambinding"), name)
	}
	return obj.(*v1alpha1.ProjectIamBinding), nil
}
