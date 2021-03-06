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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSchedulerJobCollections implements SchedulerJobCollectionInterface
type FakeSchedulerJobCollections struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var schedulerjobcollectionsResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "schedulerjobcollections"}

var schedulerjobcollectionsKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "SchedulerJobCollection"}

// Get takes name of the schedulerJobCollection, and returns the corresponding schedulerJobCollection object, and an error if there is any.
func (c *FakeSchedulerJobCollections) Get(name string, options v1.GetOptions) (result *v1alpha1.SchedulerJobCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(schedulerjobcollectionsResource, c.ns, name), &v1alpha1.SchedulerJobCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerJobCollection), err
}

// List takes label and field selectors, and returns the list of SchedulerJobCollections that match those selectors.
func (c *FakeSchedulerJobCollections) List(opts v1.ListOptions) (result *v1alpha1.SchedulerJobCollectionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(schedulerjobcollectionsResource, schedulerjobcollectionsKind, c.ns, opts), &v1alpha1.SchedulerJobCollectionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SchedulerJobCollectionList{ListMeta: obj.(*v1alpha1.SchedulerJobCollectionList).ListMeta}
	for _, item := range obj.(*v1alpha1.SchedulerJobCollectionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested schedulerJobCollections.
func (c *FakeSchedulerJobCollections) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(schedulerjobcollectionsResource, c.ns, opts))

}

// Create takes the representation of a schedulerJobCollection and creates it.  Returns the server's representation of the schedulerJobCollection, and an error, if there is any.
func (c *FakeSchedulerJobCollections) Create(schedulerJobCollection *v1alpha1.SchedulerJobCollection) (result *v1alpha1.SchedulerJobCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(schedulerjobcollectionsResource, c.ns, schedulerJobCollection), &v1alpha1.SchedulerJobCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerJobCollection), err
}

// Update takes the representation of a schedulerJobCollection and updates it. Returns the server's representation of the schedulerJobCollection, and an error, if there is any.
func (c *FakeSchedulerJobCollections) Update(schedulerJobCollection *v1alpha1.SchedulerJobCollection) (result *v1alpha1.SchedulerJobCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(schedulerjobcollectionsResource, c.ns, schedulerJobCollection), &v1alpha1.SchedulerJobCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerJobCollection), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSchedulerJobCollections) UpdateStatus(schedulerJobCollection *v1alpha1.SchedulerJobCollection) (*v1alpha1.SchedulerJobCollection, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(schedulerjobcollectionsResource, "status", c.ns, schedulerJobCollection), &v1alpha1.SchedulerJobCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerJobCollection), err
}

// Delete takes name of the schedulerJobCollection and deletes it. Returns an error if one occurs.
func (c *FakeSchedulerJobCollections) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(schedulerjobcollectionsResource, c.ns, name), &v1alpha1.SchedulerJobCollection{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSchedulerJobCollections) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(schedulerjobcollectionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SchedulerJobCollectionList{})
	return err
}

// Patch applies the patch and returns the patched schedulerJobCollection.
func (c *FakeSchedulerJobCollections) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SchedulerJobCollection, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(schedulerjobcollectionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SchedulerJobCollection{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SchedulerJobCollection), err
}
