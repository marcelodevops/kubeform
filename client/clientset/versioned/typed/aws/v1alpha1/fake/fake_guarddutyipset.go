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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGuarddutyIpsets implements GuarddutyIpsetInterface
type FakeGuarddutyIpsets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var guarddutyipsetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "guarddutyipsets"}

var guarddutyipsetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GuarddutyIpset"}

// Get takes name of the guarddutyIpset, and returns the corresponding guarddutyIpset object, and an error if there is any.
func (c *FakeGuarddutyIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.GuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(guarddutyipsetsResource, c.ns, name), &v1alpha1.GuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyIpset), err
}

// List takes label and field selectors, and returns the list of GuarddutyIpsets that match those selectors.
func (c *FakeGuarddutyIpsets) List(opts v1.ListOptions) (result *v1alpha1.GuarddutyIpsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(guarddutyipsetsResource, guarddutyipsetsKind, c.ns, opts), &v1alpha1.GuarddutyIpsetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GuarddutyIpsetList{ListMeta: obj.(*v1alpha1.GuarddutyIpsetList).ListMeta}
	for _, item := range obj.(*v1alpha1.GuarddutyIpsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested guarddutyIpsets.
func (c *FakeGuarddutyIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(guarddutyipsetsResource, c.ns, opts))

}

// Create takes the representation of a guarddutyIpset and creates it.  Returns the server's representation of the guarddutyIpset, and an error, if there is any.
func (c *FakeGuarddutyIpsets) Create(guarddutyIpset *v1alpha1.GuarddutyIpset) (result *v1alpha1.GuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(guarddutyipsetsResource, c.ns, guarddutyIpset), &v1alpha1.GuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyIpset), err
}

// Update takes the representation of a guarddutyIpset and updates it. Returns the server's representation of the guarddutyIpset, and an error, if there is any.
func (c *FakeGuarddutyIpsets) Update(guarddutyIpset *v1alpha1.GuarddutyIpset) (result *v1alpha1.GuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(guarddutyipsetsResource, c.ns, guarddutyIpset), &v1alpha1.GuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyIpset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGuarddutyIpsets) UpdateStatus(guarddutyIpset *v1alpha1.GuarddutyIpset) (*v1alpha1.GuarddutyIpset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(guarddutyipsetsResource, "status", c.ns, guarddutyIpset), &v1alpha1.GuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyIpset), err
}

// Delete takes name of the guarddutyIpset and deletes it. Returns an error if one occurs.
func (c *FakeGuarddutyIpsets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(guarddutyipsetsResource, c.ns, name), &v1alpha1.GuarddutyIpset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGuarddutyIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(guarddutyipsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GuarddutyIpsetList{})
	return err
}

// Patch applies the patch and returns the patched guarddutyIpset.
func (c *FakeGuarddutyIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyIpset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(guarddutyipsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GuarddutyIpset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GuarddutyIpset), err
}
