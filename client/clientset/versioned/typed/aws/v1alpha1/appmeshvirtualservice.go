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

// AppmeshVirtualServicesGetter has a method to return a AppmeshVirtualServiceInterface.
// A group's client should implement this interface.
type AppmeshVirtualServicesGetter interface {
	AppmeshVirtualServices(namespace string) AppmeshVirtualServiceInterface
}

// AppmeshVirtualServiceInterface has methods to work with AppmeshVirtualService resources.
type AppmeshVirtualServiceInterface interface {
	Create(*v1alpha1.AppmeshVirtualService) (*v1alpha1.AppmeshVirtualService, error)
	Update(*v1alpha1.AppmeshVirtualService) (*v1alpha1.AppmeshVirtualService, error)
	UpdateStatus(*v1alpha1.AppmeshVirtualService) (*v1alpha1.AppmeshVirtualService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppmeshVirtualService, error)
	List(opts v1.ListOptions) (*v1alpha1.AppmeshVirtualServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppmeshVirtualService, err error)
	AppmeshVirtualServiceExpansion
}

// appmeshVirtualServices implements AppmeshVirtualServiceInterface
type appmeshVirtualServices struct {
	client rest.Interface
	ns     string
}

// newAppmeshVirtualServices returns a AppmeshVirtualServices
func newAppmeshVirtualServices(c *AwsV1alpha1Client, namespace string) *appmeshVirtualServices {
	return &appmeshVirtualServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appmeshVirtualService, and returns the corresponding appmeshVirtualService object, and an error if there is any.
func (c *appmeshVirtualServices) Get(name string, options v1.GetOptions) (result *v1alpha1.AppmeshVirtualService, err error) {
	result = &v1alpha1.AppmeshVirtualService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppmeshVirtualServices that match those selectors.
func (c *appmeshVirtualServices) List(opts v1.ListOptions) (result *v1alpha1.AppmeshVirtualServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppmeshVirtualServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appmeshVirtualServices.
func (c *appmeshVirtualServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a appmeshVirtualService and creates it.  Returns the server's representation of the appmeshVirtualService, and an error, if there is any.
func (c *appmeshVirtualServices) Create(appmeshVirtualService *v1alpha1.AppmeshVirtualService) (result *v1alpha1.AppmeshVirtualService, err error) {
	result = &v1alpha1.AppmeshVirtualService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		Body(appmeshVirtualService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appmeshVirtualService and updates it. Returns the server's representation of the appmeshVirtualService, and an error, if there is any.
func (c *appmeshVirtualServices) Update(appmeshVirtualService *v1alpha1.AppmeshVirtualService) (result *v1alpha1.AppmeshVirtualService, err error) {
	result = &v1alpha1.AppmeshVirtualService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		Name(appmeshVirtualService.Name).
		Body(appmeshVirtualService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *appmeshVirtualServices) UpdateStatus(appmeshVirtualService *v1alpha1.AppmeshVirtualService) (result *v1alpha1.AppmeshVirtualService, err error) {
	result = &v1alpha1.AppmeshVirtualService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		Name(appmeshVirtualService.Name).
		SubResource("status").
		Body(appmeshVirtualService).
		Do().
		Into(result)
	return
}

// Delete takes name of the appmeshVirtualService and deletes it. Returns an error if one occurs.
func (c *appmeshVirtualServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appmeshVirtualServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appmeshVirtualService.
func (c *appmeshVirtualServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppmeshVirtualService, err error) {
	result = &v1alpha1.AppmeshVirtualService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appmeshvirtualservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
