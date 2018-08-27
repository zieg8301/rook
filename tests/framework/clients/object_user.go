/*
Copyright 2016 The Rook Authors. All rights reserved.

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

package clients

import (
	"fmt"
	"github.com/rook/rook/tests/framework/installer"
	"github.com/rook/rook/tests/framework/utils"
)

// ObjectUserOperation is wrapper for k8s rook object user operations
type ObjectUserOperation struct {
	k8sh      *utils.K8sHelper
	manifests installer.CephManifests
}

// CreateObjectUserOperation creates new rook object user client
func CreateObjectUserOperation(k8sh *utils.K8sHelper, manifests installer.CephManifests) *ObjectUserOperation {
	return &ObjectUserOperation{k8sh, manifests}
}

// // List Function to list a filesystem in rook
// // Output - output returned by the call
// func (o *ObjectUserOperation) List(namespace string) ([]rgw.ObjectUser, error) {
// 	context := o.k8sh.MakeContext()
// 	users, err := rgw.ListUsers(context)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to list users: %+v", err)
// 	}
// 	return users, nil
// }

// ObjectUserList Function to list object store users in rook
func (o *ObjectUserOperation) List(namespace string, store string) (string, error) {
	result, err := o.k8sh.GetResource("-n", namespace, "secrets", "-l", "rook_object_store="+store)
	if err != nil {
		return "null", fmt.Errorf("failed to list users: %+v", err)
	}
	return result, nil
}

// ObjectUserCreate Function to create a object store user in rook
func (o *ObjectUserOperation) Create(namespace string, name string, displayName string, store string) error {

	logger.Infof("creating the object store user via CRD")
	if _, err := o.k8sh.ResourceOperation("create", o.manifests.GetObjectStoreUser(namespace, name, displayName, store)); err != nil {
		return err
	}
	return nil
}

func (o *ObjectUserOperation) Delete(namespace string, name string, displayName string, store string) error {

	logger.Infof("Deleting the object store user via CRD")
	if _, err := o.k8sh.DeleteResource("-n", namespace, "ObjectStoreUser", store); err != nil {
		return err
	}
	return nil
}
