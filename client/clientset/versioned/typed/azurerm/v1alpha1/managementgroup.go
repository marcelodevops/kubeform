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

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagementGroupsGetter has a method to return a ManagementGroupInterface.
// A group's client should implement this interface.
type ManagementGroupsGetter interface {
	ManagementGroups(namespace string) ManagementGroupInterface
}

// ManagementGroupInterface has methods to work with ManagementGroup resources.
type ManagementGroupInterface interface {
	Create(*v1alpha1.ManagementGroup) (*v1alpha1.ManagementGroup, error)
	Update(*v1alpha1.ManagementGroup) (*v1alpha1.ManagementGroup, error)
	UpdateStatus(*v1alpha1.ManagementGroup) (*v1alpha1.ManagementGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ManagementGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.ManagementGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ManagementGroup, err error)
	ManagementGroupExpansion
}

// managementGroups implements ManagementGroupInterface
type managementGroups struct {
	client rest.Interface
	ns     string
}

// newManagementGroups returns a ManagementGroups
func newManagementGroups(c *AzurermV1alpha1Client, namespace string) *managementGroups {
	return &managementGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managementGroup, and returns the corresponding managementGroup object, and an error if there is any.
func (c *managementGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ManagementGroup, err error) {
	result = &v1alpha1.ManagementGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managementgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagementGroups that match those selectors.
func (c *managementGroups) List(opts v1.ListOptions) (result *v1alpha1.ManagementGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ManagementGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managementgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managementGroups.
func (c *managementGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managementgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a managementGroup and creates it.  Returns the server's representation of the managementGroup, and an error, if there is any.
func (c *managementGroups) Create(managementGroup *v1alpha1.ManagementGroup) (result *v1alpha1.ManagementGroup, err error) {
	result = &v1alpha1.ManagementGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managementgroups").
		Body(managementGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a managementGroup and updates it. Returns the server's representation of the managementGroup, and an error, if there is any.
func (c *managementGroups) Update(managementGroup *v1alpha1.ManagementGroup) (result *v1alpha1.ManagementGroup, err error) {
	result = &v1alpha1.ManagementGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managementgroups").
		Name(managementGroup.Name).
		Body(managementGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *managementGroups) UpdateStatus(managementGroup *v1alpha1.ManagementGroup) (result *v1alpha1.ManagementGroup, err error) {
	result = &v1alpha1.ManagementGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managementgroups").
		Name(managementGroup.Name).
		SubResource("status").
		Body(managementGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the managementGroup and deletes it. Returns an error if one occurs.
func (c *managementGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managementgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managementGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managementgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched managementGroup.
func (c *managementGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ManagementGroup, err error) {
	result = &v1alpha1.ManagementGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managementgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
