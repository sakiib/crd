/*
Copyright The Kubernetes Authors.

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
	"context"
	"time"

	v1alpha1 "github.com/sakiib/crd/pkg/apis/book.com/v1alpha1"
	scheme "github.com/sakiib/crd/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BookAPIsGetter has a method to return a BookAPIInterface.
// A group's client should implement this interface.
type BookAPIsGetter interface {
	BookAPIs(namespace string) BookAPIInterface
}

// BookAPIInterface has methods to work with BookAPI resources.
type BookAPIInterface interface {
	Create(ctx context.Context, bookAPI *v1alpha1.BookAPI, opts v1.CreateOptions) (*v1alpha1.BookAPI, error)
	Update(ctx context.Context, bookAPI *v1alpha1.BookAPI, opts v1.UpdateOptions) (*v1alpha1.BookAPI, error)
	UpdateStatus(ctx context.Context, bookAPI *v1alpha1.BookAPI, opts v1.UpdateOptions) (*v1alpha1.BookAPI, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BookAPI, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BookAPIList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BookAPI, err error)
	BookAPIExpansion
}

// bookAPIs implements BookAPIInterface
type bookAPIs struct {
	client rest.Interface
	ns     string
}

// newBookAPIs returns a BookAPIs
func newBookAPIs(c *BookV1alpha1Client, namespace string) *bookAPIs {
	return &bookAPIs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bookAPI, and returns the corresponding bookAPI object, and an error if there is any.
func (c *bookAPIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BookAPI, err error) {
	result = &v1alpha1.BookAPI{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bookapis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BookAPIs that match those selectors.
func (c *bookAPIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BookAPIList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BookAPIList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bookapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bookAPIs.
func (c *bookAPIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bookapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bookAPI and creates it.  Returns the server's representation of the bookAPI, and an error, if there is any.
func (c *bookAPIs) Create(ctx context.Context, bookAPI *v1alpha1.BookAPI, opts v1.CreateOptions) (result *v1alpha1.BookAPI, err error) {
	result = &v1alpha1.BookAPI{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bookapis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bookAPI).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bookAPI and updates it. Returns the server's representation of the bookAPI, and an error, if there is any.
func (c *bookAPIs) Update(ctx context.Context, bookAPI *v1alpha1.BookAPI, opts v1.UpdateOptions) (result *v1alpha1.BookAPI, err error) {
	result = &v1alpha1.BookAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bookapis").
		Name(bookAPI.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bookAPI).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bookAPIs) UpdateStatus(ctx context.Context, bookAPI *v1alpha1.BookAPI, opts v1.UpdateOptions) (result *v1alpha1.BookAPI, err error) {
	result = &v1alpha1.BookAPI{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bookapis").
		Name(bookAPI.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bookAPI).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bookAPI and deletes it. Returns an error if one occurs.
func (c *bookAPIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bookapis").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bookAPIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bookapis").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bookAPI.
func (c *bookAPIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BookAPI, err error) {
	result = &v1alpha1.BookAPI{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bookapis").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}