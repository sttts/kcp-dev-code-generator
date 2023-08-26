/*
Copyright The KCP Authors.

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

// Code generated by client-gen-v0.28.1. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	testing "k8s.io/client-go/testing"

	v1beta1 "acme.corp/pkg/apis/example/v1beta1"
)

// FakeTestTypes implements TestTypeInterface
type FakeTestTypes struct {
	Fake *FakeExampleV1beta1
	ns   string
}

var testtypesResource = v1beta1.SchemeGroupVersion.WithResource("testtypes")

var testtypesKind = v1beta1.SchemeGroupVersion.WithKind("TestType")

// Get takes name of the testType, and returns the corresponding testType object, and an error if there is any.
func (c *FakeTestTypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testtypesResource, c.ns, name), &v1beta1.TestType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TestType), err
}

// Create takes the representation of a testType and creates it.  Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Create(ctx context.Context, testType *v1beta1.TestType, opts v1.CreateOptions) (result *v1beta1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testtypesResource, c.ns, testType), &v1beta1.TestType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TestType), err
}

// Update takes the representation of a testType and updates it. Returns the server's representation of the testType, and an error, if there is any.
func (c *FakeTestTypes) Update(ctx context.Context, testType *v1beta1.TestType, opts v1.UpdateOptions) (result *v1beta1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testtypesResource, c.ns, testType), &v1beta1.TestType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TestType), err
}

// Delete takes name of the testType and deletes it. Returns an error if one occurs.
func (c *FakeTestTypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(testtypesResource, c.ns, name, opts), &v1beta1.TestType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestTypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testtypesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.TestTypeList{})
	return err
}

// Patch applies the patch and returns the patched testType.
func (c *FakeTestTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.TestType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testtypesResource, c.ns, name, pt, data, subresources...), &v1beta1.TestType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TestType), err
}
