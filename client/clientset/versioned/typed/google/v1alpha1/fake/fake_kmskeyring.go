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

// FakeKmsKeyRings implements KmsKeyRingInterface
type FakeKmsKeyRings struct {
	Fake *FakeGoogleV1alpha1
}

var kmskeyringsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "kmskeyrings"}

var kmskeyringsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "KmsKeyRing"}

// Get takes name of the kmsKeyRing, and returns the corresponding kmsKeyRing object, and an error if there is any.
func (c *FakeKmsKeyRings) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsKeyRing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kmskeyringsResource, name), &v1alpha1.KmsKeyRing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRing), err
}

// List takes label and field selectors, and returns the list of KmsKeyRings that match those selectors.
func (c *FakeKmsKeyRings) List(opts v1.ListOptions) (result *v1alpha1.KmsKeyRingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kmskeyringsResource, kmskeyringsKind, opts), &v1alpha1.KmsKeyRingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KmsKeyRingList{ListMeta: obj.(*v1alpha1.KmsKeyRingList).ListMeta}
	for _, item := range obj.(*v1alpha1.KmsKeyRingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kmsKeyRings.
func (c *FakeKmsKeyRings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kmskeyringsResource, opts))
}

// Create takes the representation of a kmsKeyRing and creates it.  Returns the server's representation of the kmsKeyRing, and an error, if there is any.
func (c *FakeKmsKeyRings) Create(kmsKeyRing *v1alpha1.KmsKeyRing) (result *v1alpha1.KmsKeyRing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kmskeyringsResource, kmsKeyRing), &v1alpha1.KmsKeyRing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRing), err
}

// Update takes the representation of a kmsKeyRing and updates it. Returns the server's representation of the kmsKeyRing, and an error, if there is any.
func (c *FakeKmsKeyRings) Update(kmsKeyRing *v1alpha1.KmsKeyRing) (result *v1alpha1.KmsKeyRing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kmskeyringsResource, kmsKeyRing), &v1alpha1.KmsKeyRing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRing), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKmsKeyRings) UpdateStatus(kmsKeyRing *v1alpha1.KmsKeyRing) (*v1alpha1.KmsKeyRing, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kmskeyringsResource, "status", kmsKeyRing), &v1alpha1.KmsKeyRing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRing), err
}

// Delete takes name of the kmsKeyRing and deletes it. Returns an error if one occurs.
func (c *FakeKmsKeyRings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kmskeyringsResource, name), &v1alpha1.KmsKeyRing{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKmsKeyRings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kmskeyringsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KmsKeyRingList{})
	return err
}

// Patch applies the patch and returns the patched kmsKeyRing.
func (c *FakeKmsKeyRings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsKeyRing, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kmskeyringsResource, name, pt, data, subresources...), &v1alpha1.KmsKeyRing{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKeyRing), err
}