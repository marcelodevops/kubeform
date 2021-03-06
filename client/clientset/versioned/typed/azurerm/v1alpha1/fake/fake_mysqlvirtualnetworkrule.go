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

// FakeMysqlVirtualNetworkRules implements MysqlVirtualNetworkRuleInterface
type FakeMysqlVirtualNetworkRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mysqlvirtualnetworkrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mysqlvirtualnetworkrules"}

var mysqlvirtualnetworkrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MysqlVirtualNetworkRule"}

// Get takes name of the mysqlVirtualNetworkRule, and returns the corresponding mysqlVirtualNetworkRule object, and an error if there is any.
func (c *FakeMysqlVirtualNetworkRules) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlvirtualnetworkrulesResource, c.ns, name), &v1alpha1.MysqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlVirtualNetworkRule), err
}

// List takes label and field selectors, and returns the list of MysqlVirtualNetworkRules that match those selectors.
func (c *FakeMysqlVirtualNetworkRules) List(opts v1.ListOptions) (result *v1alpha1.MysqlVirtualNetworkRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlvirtualnetworkrulesResource, mysqlvirtualnetworkrulesKind, c.ns, opts), &v1alpha1.MysqlVirtualNetworkRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlVirtualNetworkRuleList{ListMeta: obj.(*v1alpha1.MysqlVirtualNetworkRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlVirtualNetworkRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqlVirtualNetworkRules.
func (c *FakeMysqlVirtualNetworkRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlvirtualnetworkrulesResource, c.ns, opts))

}

// Create takes the representation of a mysqlVirtualNetworkRule and creates it.  Returns the server's representation of the mysqlVirtualNetworkRule, and an error, if there is any.
func (c *FakeMysqlVirtualNetworkRules) Create(mysqlVirtualNetworkRule *v1alpha1.MysqlVirtualNetworkRule) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlvirtualnetworkrulesResource, c.ns, mysqlVirtualNetworkRule), &v1alpha1.MysqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlVirtualNetworkRule), err
}

// Update takes the representation of a mysqlVirtualNetworkRule and updates it. Returns the server's representation of the mysqlVirtualNetworkRule, and an error, if there is any.
func (c *FakeMysqlVirtualNetworkRules) Update(mysqlVirtualNetworkRule *v1alpha1.MysqlVirtualNetworkRule) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlvirtualnetworkrulesResource, c.ns, mysqlVirtualNetworkRule), &v1alpha1.MysqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlVirtualNetworkRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMysqlVirtualNetworkRules) UpdateStatus(mysqlVirtualNetworkRule *v1alpha1.MysqlVirtualNetworkRule) (*v1alpha1.MysqlVirtualNetworkRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqlvirtualnetworkrulesResource, "status", c.ns, mysqlVirtualNetworkRule), &v1alpha1.MysqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlVirtualNetworkRule), err
}

// Delete takes name of the mysqlVirtualNetworkRule and deletes it. Returns an error if one occurs.
func (c *FakeMysqlVirtualNetworkRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlvirtualnetworkrulesResource, c.ns, name), &v1alpha1.MysqlVirtualNetworkRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqlVirtualNetworkRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlvirtualnetworkrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlVirtualNetworkRuleList{})
	return err
}

// Patch applies the patch and returns the patched mysqlVirtualNetworkRule.
func (c *FakeMysqlVirtualNetworkRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlVirtualNetworkRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlvirtualnetworkrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlVirtualNetworkRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlVirtualNetworkRule), err
}
