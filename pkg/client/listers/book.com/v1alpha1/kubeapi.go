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

// KubeApiLister helps list KubeApis.
// All objects returned here must be treated as read-only.
type KubeApiLister interface {
	// List lists all KubeApis in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KubeApi, err error)
	// KubeApis returns an object that can list and get KubeApis.
	KubeApis(namespace string) KubeApiNamespaceLister
	KubeApiListerExpansion
}

// kubeApiLister implements the KubeApiLister interface.
type kubeApiLister struct {
	indexer cache.Indexer
}

// NewKubeApiLister returns a new KubeApiLister.
func NewKubeApiLister(indexer cache.Indexer) KubeApiLister {
	return &kubeApiLister{indexer: indexer}
}

// List lists all KubeApis in the indexer.
func (s *kubeApiLister) List(selector labels.Selector) (ret []*v1alpha1.KubeApi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubeApi))
	})
	return ret, err
}

// KubeApis returns an object that can list and get KubeApis.
func (s *kubeApiLister) KubeApis(namespace string) KubeApiNamespaceLister {
	return kubeApiNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KubeApiNamespaceLister helps list and get KubeApis.
// All objects returned here must be treated as read-only.
type KubeApiNamespaceLister interface {
	// List lists all KubeApis in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KubeApi, err error)
	// Get retrieves the KubeApi from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KubeApi, error)
	KubeApiNamespaceListerExpansion
}

// kubeApiNamespaceLister implements the KubeApiNamespaceLister
// interface.
type kubeApiNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KubeApis in the indexer for a given namespace.
func (s kubeApiNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KubeApi, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KubeApi))
	})
	return ret, err
}

// Get retrieves the KubeApi from the indexer for a given namespace and name.
func (s kubeApiNamespaceLister) Get(name string) (*v1alpha1.KubeApi, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubeapi"), name)
	}
	return obj.(*v1alpha1.KubeApi), nil
}
