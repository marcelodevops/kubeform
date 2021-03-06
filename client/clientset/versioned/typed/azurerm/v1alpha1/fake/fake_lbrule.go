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

// FakeLbRules implements LbRuleInterface
type FakeLbRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var lbrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "lbrules"}

var lbrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "LbRule"}

// Get takes name of the lbRule, and returns the corresponding lbRule object, and an error if there is any.
func (c *FakeLbRules) Get(name string, options v1.GetOptions) (result *v1alpha1.LbRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lbrulesResource, c.ns, name), &v1alpha1.LbRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbRule), err
}

// List takes label and field selectors, and returns the list of LbRules that match those selectors.
func (c *FakeLbRules) List(opts v1.ListOptions) (result *v1alpha1.LbRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lbrulesResource, lbrulesKind, c.ns, opts), &v1alpha1.LbRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LbRuleList{ListMeta: obj.(*v1alpha1.LbRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.LbRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lbRules.
func (c *FakeLbRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lbrulesResource, c.ns, opts))

}

// Create takes the representation of a lbRule and creates it.  Returns the server's representation of the lbRule, and an error, if there is any.
func (c *FakeLbRules) Create(lbRule *v1alpha1.LbRule) (result *v1alpha1.LbRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lbrulesResource, c.ns, lbRule), &v1alpha1.LbRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbRule), err
}

// Update takes the representation of a lbRule and updates it. Returns the server's representation of the lbRule, and an error, if there is any.
func (c *FakeLbRules) Update(lbRule *v1alpha1.LbRule) (result *v1alpha1.LbRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lbrulesResource, c.ns, lbRule), &v1alpha1.LbRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLbRules) UpdateStatus(lbRule *v1alpha1.LbRule) (*v1alpha1.LbRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lbrulesResource, "status", c.ns, lbRule), &v1alpha1.LbRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbRule), err
}

// Delete takes name of the lbRule and deletes it. Returns an error if one occurs.
func (c *FakeLbRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lbrulesResource, c.ns, name), &v1alpha1.LbRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLbRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lbrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LbRuleList{})
	return err
}

// Patch applies the patch and returns the patched lbRule.
func (c *FakeLbRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LbRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lbrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LbRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LbRule), err
}
