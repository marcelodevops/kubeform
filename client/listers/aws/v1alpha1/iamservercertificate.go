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

// IamServerCertificateLister helps list IamServerCertificates.
type IamServerCertificateLister interface {
	// List lists all IamServerCertificates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamServerCertificate, err error)
	// IamServerCertificates returns an object that can list and get IamServerCertificates.
	IamServerCertificates(namespace string) IamServerCertificateNamespaceLister
	IamServerCertificateListerExpansion
}

// iamServerCertificateLister implements the IamServerCertificateLister interface.
type iamServerCertificateLister struct {
	indexer cache.Indexer
}

// NewIamServerCertificateLister returns a new IamServerCertificateLister.
func NewIamServerCertificateLister(indexer cache.Indexer) IamServerCertificateLister {
	return &iamServerCertificateLister{indexer: indexer}
}

// List lists all IamServerCertificates in the indexer.
func (s *iamServerCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.IamServerCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamServerCertificate))
	})
	return ret, err
}

// IamServerCertificates returns an object that can list and get IamServerCertificates.
func (s *iamServerCertificateLister) IamServerCertificates(namespace string) IamServerCertificateNamespaceLister {
	return iamServerCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamServerCertificateNamespaceLister helps list and get IamServerCertificates.
type IamServerCertificateNamespaceLister interface {
	// List lists all IamServerCertificates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamServerCertificate, err error)
	// Get retrieves the IamServerCertificate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamServerCertificate, error)
	IamServerCertificateNamespaceListerExpansion
}

// iamServerCertificateNamespaceLister implements the IamServerCertificateNamespaceLister
// interface.
type iamServerCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamServerCertificates in the indexer for a given namespace.
func (s iamServerCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamServerCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamServerCertificate))
	})
	return ret, err
}

// Get retrieves the IamServerCertificate from the indexer for a given namespace and name.
func (s iamServerCertificateNamespaceLister) Get(name string) (*v1alpha1.IamServerCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamservercertificate"), name)
	}
	return obj.(*v1alpha1.IamServerCertificate), nil
}
