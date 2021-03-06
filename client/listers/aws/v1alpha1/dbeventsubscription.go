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

// DbEventSubscriptionLister helps list DbEventSubscriptions.
type DbEventSubscriptionLister interface {
	// List lists all DbEventSubscriptions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DbEventSubscription, err error)
	// DbEventSubscriptions returns an object that can list and get DbEventSubscriptions.
	DbEventSubscriptions(namespace string) DbEventSubscriptionNamespaceLister
	DbEventSubscriptionListerExpansion
}

// dbEventSubscriptionLister implements the DbEventSubscriptionLister interface.
type dbEventSubscriptionLister struct {
	indexer cache.Indexer
}

// NewDbEventSubscriptionLister returns a new DbEventSubscriptionLister.
func NewDbEventSubscriptionLister(indexer cache.Indexer) DbEventSubscriptionLister {
	return &dbEventSubscriptionLister{indexer: indexer}
}

// List lists all DbEventSubscriptions in the indexer.
func (s *dbEventSubscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.DbEventSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbEventSubscription))
	})
	return ret, err
}

// DbEventSubscriptions returns an object that can list and get DbEventSubscriptions.
func (s *dbEventSubscriptionLister) DbEventSubscriptions(namespace string) DbEventSubscriptionNamespaceLister {
	return dbEventSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbEventSubscriptionNamespaceLister helps list and get DbEventSubscriptions.
type DbEventSubscriptionNamespaceLister interface {
	// List lists all DbEventSubscriptions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DbEventSubscription, err error)
	// Get retrieves the DbEventSubscription from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DbEventSubscription, error)
	DbEventSubscriptionNamespaceListerExpansion
}

// dbEventSubscriptionNamespaceLister implements the DbEventSubscriptionNamespaceLister
// interface.
type dbEventSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbEventSubscriptions in the indexer for a given namespace.
func (s dbEventSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DbEventSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbEventSubscription))
	})
	return ret, err
}

// Get retrieves the DbEventSubscription from the indexer for a given namespace and name.
func (s dbEventSubscriptionNamespaceLister) Get(name string) (*v1alpha1.DbEventSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbeventsubscription"), name)
	}
	return obj.(*v1alpha1.DbEventSubscription), nil
}
