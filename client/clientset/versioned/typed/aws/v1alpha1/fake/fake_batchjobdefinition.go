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

// FakeBatchJobDefinitions implements BatchJobDefinitionInterface
type FakeBatchJobDefinitions struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var batchjobdefinitionsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "batchjobdefinitions"}

var batchjobdefinitionsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "BatchJobDefinition"}

// Get takes name of the batchJobDefinition, and returns the corresponding batchJobDefinition object, and an error if there is any.
func (c *FakeBatchJobDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.BatchJobDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(batchjobdefinitionsResource, c.ns, name), &v1alpha1.BatchJobDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchJobDefinition), err
}

// List takes label and field selectors, and returns the list of BatchJobDefinitions that match those selectors.
func (c *FakeBatchJobDefinitions) List(opts v1.ListOptions) (result *v1alpha1.BatchJobDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(batchjobdefinitionsResource, batchjobdefinitionsKind, c.ns, opts), &v1alpha1.BatchJobDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BatchJobDefinitionList{ListMeta: obj.(*v1alpha1.BatchJobDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.BatchJobDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested batchJobDefinitions.
func (c *FakeBatchJobDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(batchjobdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a batchJobDefinition and creates it.  Returns the server's representation of the batchJobDefinition, and an error, if there is any.
func (c *FakeBatchJobDefinitions) Create(batchJobDefinition *v1alpha1.BatchJobDefinition) (result *v1alpha1.BatchJobDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(batchjobdefinitionsResource, c.ns, batchJobDefinition), &v1alpha1.BatchJobDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchJobDefinition), err
}

// Update takes the representation of a batchJobDefinition and updates it. Returns the server's representation of the batchJobDefinition, and an error, if there is any.
func (c *FakeBatchJobDefinitions) Update(batchJobDefinition *v1alpha1.BatchJobDefinition) (result *v1alpha1.BatchJobDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(batchjobdefinitionsResource, c.ns, batchJobDefinition), &v1alpha1.BatchJobDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchJobDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBatchJobDefinitions) UpdateStatus(batchJobDefinition *v1alpha1.BatchJobDefinition) (*v1alpha1.BatchJobDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(batchjobdefinitionsResource, "status", c.ns, batchJobDefinition), &v1alpha1.BatchJobDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchJobDefinition), err
}

// Delete takes name of the batchJobDefinition and deletes it. Returns an error if one occurs.
func (c *FakeBatchJobDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(batchjobdefinitionsResource, c.ns, name), &v1alpha1.BatchJobDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBatchJobDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(batchjobdefinitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BatchJobDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched batchJobDefinition.
func (c *FakeBatchJobDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BatchJobDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(batchjobdefinitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BatchJobDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BatchJobDefinition), err
}
