/*
Copyright 2019 The Kubeform Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ComputeVpnGatewaysGetter has a method to return a ComputeVpnGatewayInterface.
// A group's client should implement this interface.
type ComputeVpnGatewaysGetter interface {
	ComputeVpnGateways() ComputeVpnGatewayInterface
}

// ComputeVpnGatewayInterface has methods to work with ComputeVpnGateway resources.
type ComputeVpnGatewayInterface interface {
	Create(*v1alpha1.ComputeVpnGateway) (*v1alpha1.ComputeVpnGateway, error)
	Update(*v1alpha1.ComputeVpnGateway) (*v1alpha1.ComputeVpnGateway, error)
	UpdateStatus(*v1alpha1.ComputeVpnGateway) (*v1alpha1.ComputeVpnGateway, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeVpnGateway, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeVpnGatewayList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeVpnGateway, err error)
	ComputeVpnGatewayExpansion
}

// computeVpnGateways implements ComputeVpnGatewayInterface
type computeVpnGateways struct {
	client rest.Interface
}

// newComputeVpnGateways returns a ComputeVpnGateways
func newComputeVpnGateways(c *GoogleV1alpha1Client) *computeVpnGateways {
	return &computeVpnGateways{
		client: c.RESTClient(),
	}
}

// Get takes name of the computeVpnGateway, and returns the corresponding computeVpnGateway object, and an error if there is any.
func (c *computeVpnGateways) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeVpnGateway, err error) {
	result = &v1alpha1.ComputeVpnGateway{}
	err = c.client.Get().
		Resource("computevpngateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeVpnGateways that match those selectors.
func (c *computeVpnGateways) List(opts v1.ListOptions) (result *v1alpha1.ComputeVpnGatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeVpnGatewayList{}
	err = c.client.Get().
		Resource("computevpngateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeVpnGateways.
func (c *computeVpnGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("computevpngateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeVpnGateway and creates it.  Returns the server's representation of the computeVpnGateway, and an error, if there is any.
func (c *computeVpnGateways) Create(computeVpnGateway *v1alpha1.ComputeVpnGateway) (result *v1alpha1.ComputeVpnGateway, err error) {
	result = &v1alpha1.ComputeVpnGateway{}
	err = c.client.Post().
		Resource("computevpngateways").
		Body(computeVpnGateway).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeVpnGateway and updates it. Returns the server's representation of the computeVpnGateway, and an error, if there is any.
func (c *computeVpnGateways) Update(computeVpnGateway *v1alpha1.ComputeVpnGateway) (result *v1alpha1.ComputeVpnGateway, err error) {
	result = &v1alpha1.ComputeVpnGateway{}
	err = c.client.Put().
		Resource("computevpngateways").
		Name(computeVpnGateway.Name).
		Body(computeVpnGateway).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeVpnGateways) UpdateStatus(computeVpnGateway *v1alpha1.ComputeVpnGateway) (result *v1alpha1.ComputeVpnGateway, err error) {
	result = &v1alpha1.ComputeVpnGateway{}
	err = c.client.Put().
		Resource("computevpngateways").
		Name(computeVpnGateway.Name).
		SubResource("status").
		Body(computeVpnGateway).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeVpnGateway and deletes it. Returns an error if one occurs.
func (c *computeVpnGateways) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("computevpngateways").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeVpnGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("computevpngateways").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeVpnGateway.
func (c *computeVpnGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeVpnGateway, err error) {
	result = &v1alpha1.ComputeVpnGateway{}
	err = c.client.Patch(pt).
		Resource("computevpngateways").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}