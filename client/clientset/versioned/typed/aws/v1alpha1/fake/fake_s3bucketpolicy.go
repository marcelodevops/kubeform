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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeS3BucketPolicies implements S3BucketPolicyInterface
type FakeS3BucketPolicies struct {
	Fake *FakeAwsV1alpha1
}

var s3bucketpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "s3bucketpolicies"}

var s3bucketpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "S3BucketPolicy"}

// Get takes name of the s3BucketPolicy, and returns the corresponding s3BucketPolicy object, and an error if there is any.
func (c *FakeS3BucketPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.S3BucketPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(s3bucketpoliciesResource, name), &v1alpha1.S3BucketPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketPolicy), err
}

// List takes label and field selectors, and returns the list of S3BucketPolicies that match those selectors.
func (c *FakeS3BucketPolicies) List(opts v1.ListOptions) (result *v1alpha1.S3BucketPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(s3bucketpoliciesResource, s3bucketpoliciesKind, opts), &v1alpha1.S3BucketPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.S3BucketPolicyList{ListMeta: obj.(*v1alpha1.S3BucketPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.S3BucketPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested s3BucketPolicies.
func (c *FakeS3BucketPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(s3bucketpoliciesResource, opts))
}

// Create takes the representation of a s3BucketPolicy and creates it.  Returns the server's representation of the s3BucketPolicy, and an error, if there is any.
func (c *FakeS3BucketPolicies) Create(s3BucketPolicy *v1alpha1.S3BucketPolicy) (result *v1alpha1.S3BucketPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(s3bucketpoliciesResource, s3BucketPolicy), &v1alpha1.S3BucketPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketPolicy), err
}

// Update takes the representation of a s3BucketPolicy and updates it. Returns the server's representation of the s3BucketPolicy, and an error, if there is any.
func (c *FakeS3BucketPolicies) Update(s3BucketPolicy *v1alpha1.S3BucketPolicy) (result *v1alpha1.S3BucketPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(s3bucketpoliciesResource, s3BucketPolicy), &v1alpha1.S3BucketPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeS3BucketPolicies) UpdateStatus(s3BucketPolicy *v1alpha1.S3BucketPolicy) (*v1alpha1.S3BucketPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(s3bucketpoliciesResource, "status", s3BucketPolicy), &v1alpha1.S3BucketPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketPolicy), err
}

// Delete takes name of the s3BucketPolicy and deletes it. Returns an error if one occurs.
func (c *FakeS3BucketPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(s3bucketpoliciesResource, name), &v1alpha1.S3BucketPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeS3BucketPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(s3bucketpoliciesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.S3BucketPolicyList{})
	return err
}

// Patch applies the patch and returns the patched s3BucketPolicy.
func (c *FakeS3BucketPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.S3BucketPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(s3bucketpoliciesResource, name, pt, data, subresources...), &v1alpha1.S3BucketPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.S3BucketPolicy), err
}