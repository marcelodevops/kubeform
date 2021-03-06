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

// FakePinpointApnsSandboxChannels implements PinpointApnsSandboxChannelInterface
type FakePinpointApnsSandboxChannels struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var pinpointapnssandboxchannelsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "pinpointapnssandboxchannels"}

var pinpointapnssandboxchannelsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "PinpointApnsSandboxChannel"}

// Get takes name of the pinpointApnsSandboxChannel, and returns the corresponding pinpointApnsSandboxChannel object, and an error if there is any.
func (c *FakePinpointApnsSandboxChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointApnsSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinpointapnssandboxchannelsResource, c.ns, name), &v1alpha1.PinpointApnsSandboxChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsSandboxChannel), err
}

// List takes label and field selectors, and returns the list of PinpointApnsSandboxChannels that match those selectors.
func (c *FakePinpointApnsSandboxChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointApnsSandboxChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinpointapnssandboxchannelsResource, pinpointapnssandboxchannelsKind, c.ns, opts), &v1alpha1.PinpointApnsSandboxChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinpointApnsSandboxChannelList{ListMeta: obj.(*v1alpha1.PinpointApnsSandboxChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinpointApnsSandboxChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinpointApnsSandboxChannels.
func (c *FakePinpointApnsSandboxChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinpointapnssandboxchannelsResource, c.ns, opts))

}

// Create takes the representation of a pinpointApnsSandboxChannel and creates it.  Returns the server's representation of the pinpointApnsSandboxChannel, and an error, if there is any.
func (c *FakePinpointApnsSandboxChannels) Create(pinpointApnsSandboxChannel *v1alpha1.PinpointApnsSandboxChannel) (result *v1alpha1.PinpointApnsSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinpointapnssandboxchannelsResource, c.ns, pinpointApnsSandboxChannel), &v1alpha1.PinpointApnsSandboxChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsSandboxChannel), err
}

// Update takes the representation of a pinpointApnsSandboxChannel and updates it. Returns the server's representation of the pinpointApnsSandboxChannel, and an error, if there is any.
func (c *FakePinpointApnsSandboxChannels) Update(pinpointApnsSandboxChannel *v1alpha1.PinpointApnsSandboxChannel) (result *v1alpha1.PinpointApnsSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinpointapnssandboxchannelsResource, c.ns, pinpointApnsSandboxChannel), &v1alpha1.PinpointApnsSandboxChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsSandboxChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePinpointApnsSandboxChannels) UpdateStatus(pinpointApnsSandboxChannel *v1alpha1.PinpointApnsSandboxChannel) (*v1alpha1.PinpointApnsSandboxChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pinpointapnssandboxchannelsResource, "status", c.ns, pinpointApnsSandboxChannel), &v1alpha1.PinpointApnsSandboxChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsSandboxChannel), err
}

// Delete takes name of the pinpointApnsSandboxChannel and deletes it. Returns an error if one occurs.
func (c *FakePinpointApnsSandboxChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinpointapnssandboxchannelsResource, c.ns, name), &v1alpha1.PinpointApnsSandboxChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinpointApnsSandboxChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinpointapnssandboxchannelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinpointApnsSandboxChannelList{})
	return err
}

// Patch applies the patch and returns the patched pinpointApnsSandboxChannel.
func (c *FakePinpointApnsSandboxChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointApnsSandboxChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinpointapnssandboxchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinpointApnsSandboxChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinpointApnsSandboxChannel), err
}
