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

// FakeCloudwatchLogDestinationPolicies implements CloudwatchLogDestinationPolicyInterface
type FakeCloudwatchLogDestinationPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudwatchlogdestinationpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudwatchlogdestinationpolicies"}

var cloudwatchlogdestinationpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudwatchLogDestinationPolicy"}

// Get takes name of the cloudwatchLogDestinationPolicy, and returns the corresponding cloudwatchLogDestinationPolicy object, and an error if there is any.
func (c *FakeCloudwatchLogDestinationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudwatchLogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudwatchlogdestinationpoliciesResource, c.ns, name), &v1alpha1.CloudwatchLogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestinationPolicy), err
}

// List takes label and field selectors, and returns the list of CloudwatchLogDestinationPolicies that match those selectors.
func (c *FakeCloudwatchLogDestinationPolicies) List(opts v1.ListOptions) (result *v1alpha1.CloudwatchLogDestinationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudwatchlogdestinationpoliciesResource, cloudwatchlogdestinationpoliciesKind, c.ns, opts), &v1alpha1.CloudwatchLogDestinationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudwatchLogDestinationPolicyList{ListMeta: obj.(*v1alpha1.CloudwatchLogDestinationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudwatchLogDestinationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudwatchLogDestinationPolicies.
func (c *FakeCloudwatchLogDestinationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudwatchlogdestinationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a cloudwatchLogDestinationPolicy and creates it.  Returns the server's representation of the cloudwatchLogDestinationPolicy, and an error, if there is any.
func (c *FakeCloudwatchLogDestinationPolicies) Create(cloudwatchLogDestinationPolicy *v1alpha1.CloudwatchLogDestinationPolicy) (result *v1alpha1.CloudwatchLogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudwatchlogdestinationpoliciesResource, c.ns, cloudwatchLogDestinationPolicy), &v1alpha1.CloudwatchLogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestinationPolicy), err
}

// Update takes the representation of a cloudwatchLogDestinationPolicy and updates it. Returns the server's representation of the cloudwatchLogDestinationPolicy, and an error, if there is any.
func (c *FakeCloudwatchLogDestinationPolicies) Update(cloudwatchLogDestinationPolicy *v1alpha1.CloudwatchLogDestinationPolicy) (result *v1alpha1.CloudwatchLogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudwatchlogdestinationpoliciesResource, c.ns, cloudwatchLogDestinationPolicy), &v1alpha1.CloudwatchLogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestinationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudwatchLogDestinationPolicies) UpdateStatus(cloudwatchLogDestinationPolicy *v1alpha1.CloudwatchLogDestinationPolicy) (*v1alpha1.CloudwatchLogDestinationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudwatchlogdestinationpoliciesResource, "status", c.ns, cloudwatchLogDestinationPolicy), &v1alpha1.CloudwatchLogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestinationPolicy), err
}

// Delete takes name of the cloudwatchLogDestinationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCloudwatchLogDestinationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudwatchlogdestinationpoliciesResource, c.ns, name), &v1alpha1.CloudwatchLogDestinationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudwatchLogDestinationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudwatchlogdestinationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudwatchLogDestinationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched cloudwatchLogDestinationPolicy.
func (c *FakeCloudwatchLogDestinationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudwatchLogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudwatchlogdestinationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudwatchLogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudwatchLogDestinationPolicy), err
}
