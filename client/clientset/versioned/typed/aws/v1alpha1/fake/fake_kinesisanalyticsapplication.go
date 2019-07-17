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

// FakeKinesisAnalyticsApplications implements KinesisAnalyticsApplicationInterface
type FakeKinesisAnalyticsApplications struct {
	Fake *FakeAwsV1alpha1
}

var kinesisanalyticsapplicationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "kinesisanalyticsapplications"}

var kinesisanalyticsapplicationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "KinesisAnalyticsApplication"}

// Get takes name of the kinesisAnalyticsApplication, and returns the corresponding kinesisAnalyticsApplication object, and an error if there is any.
func (c *FakeKinesisAnalyticsApplications) Get(name string, options v1.GetOptions) (result *v1alpha1.KinesisAnalyticsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kinesisanalyticsapplicationsResource, name), &v1alpha1.KinesisAnalyticsApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisAnalyticsApplication), err
}

// List takes label and field selectors, and returns the list of KinesisAnalyticsApplications that match those selectors.
func (c *FakeKinesisAnalyticsApplications) List(opts v1.ListOptions) (result *v1alpha1.KinesisAnalyticsApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kinesisanalyticsapplicationsResource, kinesisanalyticsapplicationsKind, opts), &v1alpha1.KinesisAnalyticsApplicationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KinesisAnalyticsApplicationList{ListMeta: obj.(*v1alpha1.KinesisAnalyticsApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.KinesisAnalyticsApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kinesisAnalyticsApplications.
func (c *FakeKinesisAnalyticsApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kinesisanalyticsapplicationsResource, opts))
}

// Create takes the representation of a kinesisAnalyticsApplication and creates it.  Returns the server's representation of the kinesisAnalyticsApplication, and an error, if there is any.
func (c *FakeKinesisAnalyticsApplications) Create(kinesisAnalyticsApplication *v1alpha1.KinesisAnalyticsApplication) (result *v1alpha1.KinesisAnalyticsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kinesisanalyticsapplicationsResource, kinesisAnalyticsApplication), &v1alpha1.KinesisAnalyticsApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisAnalyticsApplication), err
}

// Update takes the representation of a kinesisAnalyticsApplication and updates it. Returns the server's representation of the kinesisAnalyticsApplication, and an error, if there is any.
func (c *FakeKinesisAnalyticsApplications) Update(kinesisAnalyticsApplication *v1alpha1.KinesisAnalyticsApplication) (result *v1alpha1.KinesisAnalyticsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kinesisanalyticsapplicationsResource, kinesisAnalyticsApplication), &v1alpha1.KinesisAnalyticsApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisAnalyticsApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKinesisAnalyticsApplications) UpdateStatus(kinesisAnalyticsApplication *v1alpha1.KinesisAnalyticsApplication) (*v1alpha1.KinesisAnalyticsApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kinesisanalyticsapplicationsResource, "status", kinesisAnalyticsApplication), &v1alpha1.KinesisAnalyticsApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisAnalyticsApplication), err
}

// Delete takes name of the kinesisAnalyticsApplication and deletes it. Returns an error if one occurs.
func (c *FakeKinesisAnalyticsApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(kinesisanalyticsapplicationsResource, name), &v1alpha1.KinesisAnalyticsApplication{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKinesisAnalyticsApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kinesisanalyticsapplicationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KinesisAnalyticsApplicationList{})
	return err
}

// Patch applies the patch and returns the patched kinesisAnalyticsApplication.
func (c *FakeKinesisAnalyticsApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KinesisAnalyticsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kinesisanalyticsapplicationsResource, name, pt, data, subresources...), &v1alpha1.KinesisAnalyticsApplication{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KinesisAnalyticsApplication), err
}