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

// SimpledbDomainsGetter has a method to return a SimpledbDomainInterface.
// A group's client should implement this interface.
type SimpledbDomainsGetter interface {
	SimpledbDomains(namespace string) SimpledbDomainInterface
}

// SimpledbDomainInterface has methods to work with SimpledbDomain resources.
type SimpledbDomainInterface interface {
	Create(*v1alpha1.SimpledbDomain) (*v1alpha1.SimpledbDomain, error)
	Update(*v1alpha1.SimpledbDomain) (*v1alpha1.SimpledbDomain, error)
	UpdateStatus(*v1alpha1.SimpledbDomain) (*v1alpha1.SimpledbDomain, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SimpledbDomain, error)
	List(opts v1.ListOptions) (*v1alpha1.SimpledbDomainList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SimpledbDomain, err error)
	SimpledbDomainExpansion
}

// simpledbDomains implements SimpledbDomainInterface
type simpledbDomains struct {
	client rest.Interface
	ns     string
}

// newSimpledbDomains returns a SimpledbDomains
func newSimpledbDomains(c *AwsV1alpha1Client, namespace string) *simpledbDomains {
	return &simpledbDomains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the simpledbDomain, and returns the corresponding simpledbDomain object, and an error if there is any.
func (c *simpledbDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.SimpledbDomain, err error) {
	result = &v1alpha1.SimpledbDomain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("simpledbdomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SimpledbDomains that match those selectors.
func (c *simpledbDomains) List(opts v1.ListOptions) (result *v1alpha1.SimpledbDomainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SimpledbDomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("simpledbdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested simpledbDomains.
func (c *simpledbDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("simpledbdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a simpledbDomain and creates it.  Returns the server's representation of the simpledbDomain, and an error, if there is any.
func (c *simpledbDomains) Create(simpledbDomain *v1alpha1.SimpledbDomain) (result *v1alpha1.SimpledbDomain, err error) {
	result = &v1alpha1.SimpledbDomain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("simpledbdomains").
		Body(simpledbDomain).
		Do().
		Into(result)
	return
}

// Update takes the representation of a simpledbDomain and updates it. Returns the server's representation of the simpledbDomain, and an error, if there is any.
func (c *simpledbDomains) Update(simpledbDomain *v1alpha1.SimpledbDomain) (result *v1alpha1.SimpledbDomain, err error) {
	result = &v1alpha1.SimpledbDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("simpledbdomains").
		Name(simpledbDomain.Name).
		Body(simpledbDomain).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *simpledbDomains) UpdateStatus(simpledbDomain *v1alpha1.SimpledbDomain) (result *v1alpha1.SimpledbDomain, err error) {
	result = &v1alpha1.SimpledbDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("simpledbdomains").
		Name(simpledbDomain.Name).
		SubResource("status").
		Body(simpledbDomain).
		Do().
		Into(result)
	return
}

// Delete takes name of the simpledbDomain and deletes it. Returns an error if one occurs.
func (c *simpledbDomains) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("simpledbdomains").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *simpledbDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("simpledbdomains").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched simpledbDomain.
func (c *simpledbDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SimpledbDomain, err error) {
	result = &v1alpha1.SimpledbDomain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("simpledbdomains").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
