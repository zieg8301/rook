/*
Copyright 2018 The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ObjectStoreUsersGetter has a method to return a ObjectStoreUserInterface.
// A group's client should implement this interface.
type ObjectStoreUsersGetter interface {
	ObjectStoreUsers(namespace string) ObjectStoreUserInterface
}

// ObjectStoreUserInterface has methods to work with ObjectStoreUser resources.
type ObjectStoreUserInterface interface {
	Create(*v1beta1.ObjectStoreUser) (*v1beta1.ObjectStoreUser, error)
	Update(*v1beta1.ObjectStoreUser) (*v1beta1.ObjectStoreUser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ObjectStoreUser, error)
	List(opts v1.ListOptions) (*v1beta1.ObjectStoreUserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ObjectStoreUser, err error)
	ObjectStoreUserExpansion
}

// objectStoreUsers implements ObjectStoreUserInterface
type objectStoreUsers struct {
	client rest.Interface
	ns     string
}

// newObjectStoreUsers returns a ObjectStoreUsers
func newObjectStoreUsers(c *CephV1beta1Client, namespace string) *objectStoreUsers {
	return &objectStoreUsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the objectStoreUser, and returns the corresponding objectStoreUser object, and an error if there is any.
func (c *objectStoreUsers) Get(name string, options v1.GetOptions) (result *v1beta1.ObjectStoreUser, err error) {
	result = &v1beta1.ObjectStoreUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("objectstoreusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ObjectStoreUsers that match those selectors.
func (c *objectStoreUsers) List(opts v1.ListOptions) (result *v1beta1.ObjectStoreUserList, err error) {
	result = &v1beta1.ObjectStoreUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("objectstoreusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested objectStoreUsers.
func (c *objectStoreUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("objectstoreusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a objectStoreUser and creates it.  Returns the server's representation of the objectStoreUser, and an error, if there is any.
func (c *objectStoreUsers) Create(objectStoreUser *v1beta1.ObjectStoreUser) (result *v1beta1.ObjectStoreUser, err error) {
	result = &v1beta1.ObjectStoreUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("objectstoreusers").
		Body(objectStoreUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a objectStoreUser and updates it. Returns the server's representation of the objectStoreUser, and an error, if there is any.
func (c *objectStoreUsers) Update(objectStoreUser *v1beta1.ObjectStoreUser) (result *v1beta1.ObjectStoreUser, err error) {
	result = &v1beta1.ObjectStoreUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("objectstoreusers").
		Name(objectStoreUser.Name).
		Body(objectStoreUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the objectStoreUser and deletes it. Returns an error if one occurs.
func (c *objectStoreUsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("objectstoreusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *objectStoreUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("objectstoreusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched objectStoreUser.
func (c *objectStoreUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ObjectStoreUser, err error) {
	result = &v1beta1.ObjectStoreUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("objectstoreusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
