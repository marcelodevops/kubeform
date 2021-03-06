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

// DocdbClustersGetter has a method to return a DocdbClusterInterface.
// A group's client should implement this interface.
type DocdbClustersGetter interface {
	DocdbClusters(namespace string) DocdbClusterInterface
}

// DocdbClusterInterface has methods to work with DocdbCluster resources.
type DocdbClusterInterface interface {
	Create(*v1alpha1.DocdbCluster) (*v1alpha1.DocdbCluster, error)
	Update(*v1alpha1.DocdbCluster) (*v1alpha1.DocdbCluster, error)
	UpdateStatus(*v1alpha1.DocdbCluster) (*v1alpha1.DocdbCluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DocdbCluster, error)
	List(opts v1.ListOptions) (*v1alpha1.DocdbClusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DocdbCluster, err error)
	DocdbClusterExpansion
}

// docdbClusters implements DocdbClusterInterface
type docdbClusters struct {
	client rest.Interface
	ns     string
}

// newDocdbClusters returns a DocdbClusters
func newDocdbClusters(c *AwsV1alpha1Client, namespace string) *docdbClusters {
	return &docdbClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the docdbCluster, and returns the corresponding docdbCluster object, and an error if there is any.
func (c *docdbClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.DocdbCluster, err error) {
	result = &v1alpha1.DocdbCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("docdbclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DocdbClusters that match those selectors.
func (c *docdbClusters) List(opts v1.ListOptions) (result *v1alpha1.DocdbClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DocdbClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("docdbclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested docdbClusters.
func (c *docdbClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("docdbclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a docdbCluster and creates it.  Returns the server's representation of the docdbCluster, and an error, if there is any.
func (c *docdbClusters) Create(docdbCluster *v1alpha1.DocdbCluster) (result *v1alpha1.DocdbCluster, err error) {
	result = &v1alpha1.DocdbCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("docdbclusters").
		Body(docdbCluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a docdbCluster and updates it. Returns the server's representation of the docdbCluster, and an error, if there is any.
func (c *docdbClusters) Update(docdbCluster *v1alpha1.DocdbCluster) (result *v1alpha1.DocdbCluster, err error) {
	result = &v1alpha1.DocdbCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("docdbclusters").
		Name(docdbCluster.Name).
		Body(docdbCluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *docdbClusters) UpdateStatus(docdbCluster *v1alpha1.DocdbCluster) (result *v1alpha1.DocdbCluster, err error) {
	result = &v1alpha1.DocdbCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("docdbclusters").
		Name(docdbCluster.Name).
		SubResource("status").
		Body(docdbCluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the docdbCluster and deletes it. Returns an error if one occurs.
func (c *docdbClusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("docdbclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *docdbClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("docdbclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched docdbCluster.
func (c *docdbClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DocdbCluster, err error) {
	result = &v1alpha1.DocdbCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("docdbclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
