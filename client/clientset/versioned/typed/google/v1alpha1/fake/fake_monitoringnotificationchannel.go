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

// FakeMonitoringNotificationChannels implements MonitoringNotificationChannelInterface
type FakeMonitoringNotificationChannels struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var monitoringnotificationchannelsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "monitoringnotificationchannels"}

var monitoringnotificationchannelsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "MonitoringNotificationChannel"}

// Get takes name of the monitoringNotificationChannel, and returns the corresponding monitoringNotificationChannel object, and an error if there is any.
func (c *FakeMonitoringNotificationChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.MonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(monitoringnotificationchannelsResource, c.ns, name), &v1alpha1.MonitoringNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringNotificationChannel), err
}

// List takes label and field selectors, and returns the list of MonitoringNotificationChannels that match those selectors.
func (c *FakeMonitoringNotificationChannels) List(opts v1.ListOptions) (result *v1alpha1.MonitoringNotificationChannelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(monitoringnotificationchannelsResource, monitoringnotificationchannelsKind, c.ns, opts), &v1alpha1.MonitoringNotificationChannelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MonitoringNotificationChannelList{ListMeta: obj.(*v1alpha1.MonitoringNotificationChannelList).ListMeta}
	for _, item := range obj.(*v1alpha1.MonitoringNotificationChannelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested monitoringNotificationChannels.
func (c *FakeMonitoringNotificationChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(monitoringnotificationchannelsResource, c.ns, opts))

}

// Create takes the representation of a monitoringNotificationChannel and creates it.  Returns the server's representation of the monitoringNotificationChannel, and an error, if there is any.
func (c *FakeMonitoringNotificationChannels) Create(monitoringNotificationChannel *v1alpha1.MonitoringNotificationChannel) (result *v1alpha1.MonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(monitoringnotificationchannelsResource, c.ns, monitoringNotificationChannel), &v1alpha1.MonitoringNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringNotificationChannel), err
}

// Update takes the representation of a monitoringNotificationChannel and updates it. Returns the server's representation of the monitoringNotificationChannel, and an error, if there is any.
func (c *FakeMonitoringNotificationChannels) Update(monitoringNotificationChannel *v1alpha1.MonitoringNotificationChannel) (result *v1alpha1.MonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(monitoringnotificationchannelsResource, c.ns, monitoringNotificationChannel), &v1alpha1.MonitoringNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringNotificationChannel), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMonitoringNotificationChannels) UpdateStatus(monitoringNotificationChannel *v1alpha1.MonitoringNotificationChannel) (*v1alpha1.MonitoringNotificationChannel, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(monitoringnotificationchannelsResource, "status", c.ns, monitoringNotificationChannel), &v1alpha1.MonitoringNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringNotificationChannel), err
}

// Delete takes name of the monitoringNotificationChannel and deletes it. Returns an error if one occurs.
func (c *FakeMonitoringNotificationChannels) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(monitoringnotificationchannelsResource, c.ns, name), &v1alpha1.MonitoringNotificationChannel{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMonitoringNotificationChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(monitoringnotificationchannelsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MonitoringNotificationChannelList{})
	return err
}

// Patch applies the patch and returns the patched monitoringNotificationChannel.
func (c *FakeMonitoringNotificationChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MonitoringNotificationChannel, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(monitoringnotificationchannelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MonitoringNotificationChannel{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MonitoringNotificationChannel), err
}
