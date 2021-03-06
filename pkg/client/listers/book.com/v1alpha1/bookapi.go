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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sakiib/crd/pkg/apis/book.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BookAPILister helps list BookAPIs.
// All objects returned here must be treated as read-only.
type BookAPILister interface {
	// List lists all BookAPIs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BookAPI, err error)
	// BookAPIs returns an object that can list and get BookAPIs.
	BookAPIs(namespace string) BookAPINamespaceLister
	BookAPIListerExpansion
}

// bookAPILister implements the BookAPILister interface.
type bookAPILister struct {
	indexer cache.Indexer
}

// NewBookAPILister returns a new BookAPILister.
func NewBookAPILister(indexer cache.Indexer) BookAPILister {
	return &bookAPILister{indexer: indexer}
}

// List lists all BookAPIs in the indexer.
func (s *bookAPILister) List(selector labels.Selector) (ret []*v1alpha1.BookAPI, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BookAPI))
	})
	return ret, err
}

// BookAPIs returns an object that can list and get BookAPIs.
func (s *bookAPILister) BookAPIs(namespace string) BookAPINamespaceLister {
	return bookAPINamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BookAPINamespaceLister helps list and get BookAPIs.
// All objects returned here must be treated as read-only.
type BookAPINamespaceLister interface {
	// List lists all BookAPIs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BookAPI, err error)
	// Get retrieves the BookAPI from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BookAPI, error)
	BookAPINamespaceListerExpansion
}

// bookAPINamespaceLister implements the BookAPINamespaceLister
// interface.
type bookAPINamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BookAPIs in the indexer for a given namespace.
func (s bookAPINamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BookAPI, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BookAPI))
	})
	return ret, err
}

// Get retrieves the BookAPI from the indexer for a given namespace and name.
func (s bookAPINamespaceLister) Get(name string) (*v1alpha1.BookAPI, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bookapi"), name)
	}
	return obj.(*v1alpha1.BookAPI), nil
}
