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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeCloudfrontPublicKeys implements CloudfrontPublicKeyInterface
type FakeCloudfrontPublicKeys struct {
	Fake *FakeAwsV1alpha1
}

var cloudfrontpublickeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudfrontpublickeys"}

var cloudfrontpublickeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudfrontPublicKey"}

// Get takes name of the cloudfrontPublicKey, and returns the corresponding cloudfrontPublicKey object, and an error if there is any.
func (c *FakeCloudfrontPublicKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudfrontPublicKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudfrontpublickeysResource, name), &v1alpha1.CloudfrontPublicKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudfrontPublicKey), err
}

// List takes label and field selectors, and returns the list of CloudfrontPublicKeys that match those selectors.
func (c *FakeCloudfrontPublicKeys) List(opts v1.ListOptions) (result *v1alpha1.CloudfrontPublicKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudfrontpublickeysResource, cloudfrontpublickeysKind, opts), &v1alpha1.CloudfrontPublicKeyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudfrontPublicKeyList{ListMeta: obj.(*v1alpha1.CloudfrontPublicKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudfrontPublicKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudfrontPublicKeys.
func (c *FakeCloudfrontPublicKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudfrontpublickeysResource, opts))
}

// Create takes the representation of a cloudfrontPublicKey and creates it.  Returns the server's representation of the cloudfrontPublicKey, and an error, if there is any.
func (c *FakeCloudfrontPublicKeys) Create(cloudfrontPublicKey *v1alpha1.CloudfrontPublicKey) (result *v1alpha1.CloudfrontPublicKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudfrontpublickeysResource, cloudfrontPublicKey), &v1alpha1.CloudfrontPublicKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudfrontPublicKey), err
}

// Update takes the representation of a cloudfrontPublicKey and updates it. Returns the server's representation of the cloudfrontPublicKey, and an error, if there is any.
func (c *FakeCloudfrontPublicKeys) Update(cloudfrontPublicKey *v1alpha1.CloudfrontPublicKey) (result *v1alpha1.CloudfrontPublicKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudfrontpublickeysResource, cloudfrontPublicKey), &v1alpha1.CloudfrontPublicKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudfrontPublicKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudfrontPublicKeys) UpdateStatus(cloudfrontPublicKey *v1alpha1.CloudfrontPublicKey) (*v1alpha1.CloudfrontPublicKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cloudfrontpublickeysResource, "status", cloudfrontPublicKey), &v1alpha1.CloudfrontPublicKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudfrontPublicKey), err
}

// Delete takes name of the cloudfrontPublicKey and deletes it. Returns an error if one occurs.
func (c *FakeCloudfrontPublicKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cloudfrontpublickeysResource, name), &v1alpha1.CloudfrontPublicKey{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudfrontPublicKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudfrontpublickeysResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudfrontPublicKeyList{})
	return err
}

// Patch applies the patch and returns the patched cloudfrontPublicKey.
func (c *FakeCloudfrontPublicKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfrontPublicKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudfrontpublickeysResource, name, pt, data, subresources...), &v1alpha1.CloudfrontPublicKey{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudfrontPublicKey), err
}