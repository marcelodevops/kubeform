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

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PinpointEmailChannelsGetter has a method to return a PinpointEmailChannelInterface.
// A group's client should implement this interface.
type PinpointEmailChannelsGetter interface {
	PinpointEmailChannels(namespace string) PinpointEmailChannelInterface
}

// PinpointEmailChannelInterface has methods to work with PinpointEmailChannel resources.
type PinpointEmailChannelInterface interface {
	Create(*v1alpha1.PinpointEmailChannel) (*v1alpha1.PinpointEmailChannel, error)
	Update(*v1alpha1.PinpointEmailChannel) (*v1alpha1.PinpointEmailChannel, error)
	UpdateStatus(*v1alpha1.PinpointEmailChannel) (*v1alpha1.PinpointEmailChannel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PinpointEmailChannel, error)
	List(opts v1.ListOptions) (*v1alpha1.PinpointEmailChannelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointEmailChannel, err error)
	PinpointEmailChannelExpansion
}

// pinpointEmailChannels implements PinpointEmailChannelInterface
type pinpointEmailChannels struct {
	client rest.Interface
	ns     string
}

// newPinpointEmailChannels returns a PinpointEmailChannels
func newPinpointEmailChannels(c *AwsV1alpha1Client, namespace string) *pinpointEmailChannels {
	return &pinpointEmailChannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pinpointEmailChannel, and returns the corresponding pinpointEmailChannel object, and an error if there is any.
func (c *pinpointEmailChannels) Get(name string, options v1.GetOptions) (result *v1alpha1.PinpointEmailChannel, err error) {
	result = &v1alpha1.PinpointEmailChannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PinpointEmailChannels that match those selectors.
func (c *pinpointEmailChannels) List(opts v1.ListOptions) (result *v1alpha1.PinpointEmailChannelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PinpointEmailChannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pinpointEmailChannels.
func (c *pinpointEmailChannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pinpointEmailChannel and creates it.  Returns the server's representation of the pinpointEmailChannel, and an error, if there is any.
func (c *pinpointEmailChannels) Create(pinpointEmailChannel *v1alpha1.PinpointEmailChannel) (result *v1alpha1.PinpointEmailChannel, err error) {
	result = &v1alpha1.PinpointEmailChannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		Body(pinpointEmailChannel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pinpointEmailChannel and updates it. Returns the server's representation of the pinpointEmailChannel, and an error, if there is any.
func (c *pinpointEmailChannels) Update(pinpointEmailChannel *v1alpha1.PinpointEmailChannel) (result *v1alpha1.PinpointEmailChannel, err error) {
	result = &v1alpha1.PinpointEmailChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		Name(pinpointEmailChannel.Name).
		Body(pinpointEmailChannel).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pinpointEmailChannels) UpdateStatus(pinpointEmailChannel *v1alpha1.PinpointEmailChannel) (result *v1alpha1.PinpointEmailChannel, err error) {
	result = &v1alpha1.PinpointEmailChannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		Name(pinpointEmailChannel.Name).
		SubResource("status").
		Body(pinpointEmailChannel).
		Do().
		Into(result)
	return
}

// Delete takes name of the pinpointEmailChannel and deletes it. Returns an error if one occurs.
func (c *pinpointEmailChannels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pinpointEmailChannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pinpointEmailChannel.
func (c *pinpointEmailChannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PinpointEmailChannel, err error) {
	result = &v1alpha1.PinpointEmailChannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pinpointemailchannels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
