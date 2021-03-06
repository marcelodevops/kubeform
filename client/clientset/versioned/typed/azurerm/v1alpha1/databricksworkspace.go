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

// DatabricksWorkspacesGetter has a method to return a DatabricksWorkspaceInterface.
// A group's client should implement this interface.
type DatabricksWorkspacesGetter interface {
	DatabricksWorkspaces(namespace string) DatabricksWorkspaceInterface
}

// DatabricksWorkspaceInterface has methods to work with DatabricksWorkspace resources.
type DatabricksWorkspaceInterface interface {
	Create(*v1alpha1.DatabricksWorkspace) (*v1alpha1.DatabricksWorkspace, error)
	Update(*v1alpha1.DatabricksWorkspace) (*v1alpha1.DatabricksWorkspace, error)
	UpdateStatus(*v1alpha1.DatabricksWorkspace) (*v1alpha1.DatabricksWorkspace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DatabricksWorkspace, error)
	List(opts v1.ListOptions) (*v1alpha1.DatabricksWorkspaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatabricksWorkspace, err error)
	DatabricksWorkspaceExpansion
}

// databricksWorkspaces implements DatabricksWorkspaceInterface
type databricksWorkspaces struct {
	client rest.Interface
	ns     string
}

// newDatabricksWorkspaces returns a DatabricksWorkspaces
func newDatabricksWorkspaces(c *AzurermV1alpha1Client, namespace string) *databricksWorkspaces {
	return &databricksWorkspaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the databricksWorkspace, and returns the corresponding databricksWorkspace object, and an error if there is any.
func (c *databricksWorkspaces) Get(name string, options v1.GetOptions) (result *v1alpha1.DatabricksWorkspace, err error) {
	result = &v1alpha1.DatabricksWorkspace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatabricksWorkspaces that match those selectors.
func (c *databricksWorkspaces) List(opts v1.ListOptions) (result *v1alpha1.DatabricksWorkspaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatabricksWorkspaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested databricksWorkspaces.
func (c *databricksWorkspaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a databricksWorkspace and creates it.  Returns the server's representation of the databricksWorkspace, and an error, if there is any.
func (c *databricksWorkspaces) Create(databricksWorkspace *v1alpha1.DatabricksWorkspace) (result *v1alpha1.DatabricksWorkspace, err error) {
	result = &v1alpha1.DatabricksWorkspace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		Body(databricksWorkspace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a databricksWorkspace and updates it. Returns the server's representation of the databricksWorkspace, and an error, if there is any.
func (c *databricksWorkspaces) Update(databricksWorkspace *v1alpha1.DatabricksWorkspace) (result *v1alpha1.DatabricksWorkspace, err error) {
	result = &v1alpha1.DatabricksWorkspace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		Name(databricksWorkspace.Name).
		Body(databricksWorkspace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *databricksWorkspaces) UpdateStatus(databricksWorkspace *v1alpha1.DatabricksWorkspace) (result *v1alpha1.DatabricksWorkspace, err error) {
	result = &v1alpha1.DatabricksWorkspace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		Name(databricksWorkspace.Name).
		SubResource("status").
		Body(databricksWorkspace).
		Do().
		Into(result)
	return
}

// Delete takes name of the databricksWorkspace and deletes it. Returns an error if one occurs.
func (c *databricksWorkspaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *databricksWorkspaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("databricksworkspaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched databricksWorkspace.
func (c *databricksWorkspaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatabricksWorkspace, err error) {
	result = &v1alpha1.DatabricksWorkspace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("databricksworkspaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
