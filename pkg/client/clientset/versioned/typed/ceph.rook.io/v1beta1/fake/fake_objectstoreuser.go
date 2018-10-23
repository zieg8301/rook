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
 package fake
 import (
	v1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)
 // FakeObjectStoreUsers implements ObjectStoreUserInterface
type FakeObjectStoreUsers struct {
	Fake *FakeCephV1beta1
	ns   string
}
 var objectstoreusersResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1beta1", Resource: "objectstoreusers"}
 var objectstoreusersKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1beta1", Kind: "ObjectStoreUser"}
 // Get takes name of the objectStoreUser, and returns the corresponding objectStoreUser object, and an error if there is any.
func (c *FakeObjectStoreUsers) Get(name string, options v1.GetOptions) (result *v1beta1.ObjectStoreUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objectstoreusersResource, c.ns, name), &v1beta1.ObjectStoreUser{})
 	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectStoreUser), err
}
 // List takes label and field selectors, and returns the list of ObjectStoreUsers that match those selectors.
func (c *FakeObjectStoreUsers) List(opts v1.ListOptions) (result *v1beta1.ObjectStoreUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objectstoreusersResource, objectstoreusersKind, c.ns, opts), &v1beta1.ObjectStoreUserList{})
 	if obj == nil {
		return nil, err
	}
 	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ObjectStoreUserList{}
	for _, item := range obj.(*v1beta1.ObjectStoreUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
 // Watch returns a watch.Interface that watches the requested objectStoreUsers.
func (c *FakeObjectStoreUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objectstoreusersResource, c.ns, opts))
 }
 // Create takes the representation of a objectStoreUser and creates it.  Returns the server's representation of the objectStoreUser, and an error, if there is any.
func (c *FakeObjectStoreUsers) Create(objectStoreUser *v1beta1.ObjectStoreUser) (result *v1beta1.ObjectStoreUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objectstoreusersResource, c.ns, objectStoreUser), &v1beta1.ObjectStoreUser{})
 	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectStoreUser), err
}
 // Update takes the representation of a objectStoreUser and updates it. Returns the server's representation of the objectStoreUser, and an error, if there is any.
func (c *FakeObjectStoreUsers) Update(objectStoreUser *v1beta1.ObjectStoreUser) (result *v1beta1.ObjectStoreUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objectstoreusersResource, c.ns, objectStoreUser), &v1beta1.ObjectStoreUser{})
 	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectStoreUser), err
}
 // Delete takes name of the objectStoreUser and deletes it. Returns an error if one occurs.
func (c *FakeObjectStoreUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objectstoreusersResource, c.ns, name), &v1beta1.ObjectStoreUser{})
 	return err
}
 // DeleteCollection deletes a collection of objects.
func (c *FakeObjectStoreUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objectstoreusersResource, c.ns, listOptions)
 	_, err := c.Fake.Invokes(action, &v1beta1.ObjectStoreUserList{})
	return err
}
 // Patch applies the patch and returns the patched objectStoreUser.
func (c *FakeObjectStoreUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ObjectStoreUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objectstoreusersResource, c.ns, name, data, subresources...), &v1beta1.ObjectStoreUser{})
 	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ObjectStoreUser), err
}
