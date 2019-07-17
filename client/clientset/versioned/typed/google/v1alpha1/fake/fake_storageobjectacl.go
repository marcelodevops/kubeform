/*
Copyright 2019 The Kubeform Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// FakeStorageObjectAcls implements StorageObjectAclInterface
type FakeStorageObjectAcls struct {
	Fake *FakeGoogleV1alpha1
}

var storageobjectaclsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storageobjectacls"}

var storageobjectaclsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageObjectAcl"}

// Get takes name of the storageObjectAcl, and returns the corresponding storageObjectAcl object, and an error if there is any.
func (c *FakeStorageObjectAcls) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storageobjectaclsResource, name), &v1alpha1.StorageObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectAcl), err
}

// List takes label and field selectors, and returns the list of StorageObjectAcls that match those selectors.
func (c *FakeStorageObjectAcls) List(opts v1.ListOptions) (result *v1alpha1.StorageObjectAclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storageobjectaclsResource, storageobjectaclsKind, opts), &v1alpha1.StorageObjectAclList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageObjectAclList{ListMeta: obj.(*v1alpha1.StorageObjectAclList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageObjectAclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageObjectAcls.
func (c *FakeStorageObjectAcls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storageobjectaclsResource, opts))
}

// Create takes the representation of a storageObjectAcl and creates it.  Returns the server's representation of the storageObjectAcl, and an error, if there is any.
func (c *FakeStorageObjectAcls) Create(storageObjectAcl *v1alpha1.StorageObjectAcl) (result *v1alpha1.StorageObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storageobjectaclsResource, storageObjectAcl), &v1alpha1.StorageObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectAcl), err
}

// Update takes the representation of a storageObjectAcl and updates it. Returns the server's representation of the storageObjectAcl, and an error, if there is any.
func (c *FakeStorageObjectAcls) Update(storageObjectAcl *v1alpha1.StorageObjectAcl) (result *v1alpha1.StorageObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storageobjectaclsResource, storageObjectAcl), &v1alpha1.StorageObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectAcl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageObjectAcls) UpdateStatus(storageObjectAcl *v1alpha1.StorageObjectAcl) (*v1alpha1.StorageObjectAcl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(storageobjectaclsResource, "status", storageObjectAcl), &v1alpha1.StorageObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectAcl), err
}

// Delete takes name of the storageObjectAcl and deletes it. Returns an error if one occurs.
func (c *FakeStorageObjectAcls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storageobjectaclsResource, name), &v1alpha1.StorageObjectAcl{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageObjectAcls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storageobjectaclsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageObjectAclList{})
	return err
}

// Patch applies the patch and returns the patched storageObjectAcl.
func (c *FakeStorageObjectAcls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageObjectAcl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storageobjectaclsResource, name, pt, data, subresources...), &v1alpha1.StorageObjectAcl{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageObjectAcl), err
}