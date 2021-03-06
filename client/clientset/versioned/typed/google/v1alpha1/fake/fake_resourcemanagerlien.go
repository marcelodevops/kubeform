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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceManagerLiens implements ResourceManagerLienInterface
type FakeResourceManagerLiens struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var resourcemanagerliensResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "resourcemanagerliens"}

var resourcemanagerliensKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "ResourceManagerLien"}

// Get takes name of the resourceManagerLien, and returns the corresponding resourceManagerLien object, and an error if there is any.
func (c *FakeResourceManagerLiens) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcemanagerliensResource, c.ns, name), &v1alpha1.ResourceManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceManagerLien), err
}

// List takes label and field selectors, and returns the list of ResourceManagerLiens that match those selectors.
func (c *FakeResourceManagerLiens) List(opts v1.ListOptions) (result *v1alpha1.ResourceManagerLienList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcemanagerliensResource, resourcemanagerliensKind, c.ns, opts), &v1alpha1.ResourceManagerLienList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceManagerLienList{ListMeta: obj.(*v1alpha1.ResourceManagerLienList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceManagerLienList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceManagerLiens.
func (c *FakeResourceManagerLiens) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcemanagerliensResource, c.ns, opts))

}

// Create takes the representation of a resourceManagerLien and creates it.  Returns the server's representation of the resourceManagerLien, and an error, if there is any.
func (c *FakeResourceManagerLiens) Create(resourceManagerLien *v1alpha1.ResourceManagerLien) (result *v1alpha1.ResourceManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcemanagerliensResource, c.ns, resourceManagerLien), &v1alpha1.ResourceManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceManagerLien), err
}

// Update takes the representation of a resourceManagerLien and updates it. Returns the server's representation of the resourceManagerLien, and an error, if there is any.
func (c *FakeResourceManagerLiens) Update(resourceManagerLien *v1alpha1.ResourceManagerLien) (result *v1alpha1.ResourceManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcemanagerliensResource, c.ns, resourceManagerLien), &v1alpha1.ResourceManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceManagerLien), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceManagerLiens) UpdateStatus(resourceManagerLien *v1alpha1.ResourceManagerLien) (*v1alpha1.ResourceManagerLien, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcemanagerliensResource, "status", c.ns, resourceManagerLien), &v1alpha1.ResourceManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceManagerLien), err
}

// Delete takes name of the resourceManagerLien and deletes it. Returns an error if one occurs.
func (c *FakeResourceManagerLiens) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcemanagerliensResource, c.ns, name), &v1alpha1.ResourceManagerLien{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceManagerLiens) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcemanagerliensResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceManagerLienList{})
	return err
}

// Patch applies the patch and returns the patched resourceManagerLien.
func (c *FakeResourceManagerLiens) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceManagerLien, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcemanagerliensResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceManagerLien{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceManagerLien), err
}
