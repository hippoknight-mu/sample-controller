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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "k8s.io/api/scheduling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	schedulingv1 "k8s.io/client-go/applyconfigurations/scheduling/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// PriorityClassesGetter has a method to return a PriorityClassInterface.
// A group's client should implement this interface.
type PriorityClassesGetter interface {
	PriorityClasses() PriorityClassInterface
}

// PriorityClassInterface has methods to work with PriorityClass resources.
type PriorityClassInterface interface {
	Create(ctx context.Context, priorityClass *v1.PriorityClass, opts metav1.CreateOptions) (*v1.PriorityClass, error)
	Update(ctx context.Context, priorityClass *v1.PriorityClass, opts metav1.UpdateOptions) (*v1.PriorityClass, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PriorityClass, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PriorityClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PriorityClass, err error)
	Apply(ctx context.Context, priorityClass *schedulingv1.PriorityClassApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityClass, err error)
	PriorityClassExpansion
}

// priorityClasses implements PriorityClassInterface
type priorityClasses struct {
	client rest.Interface
}

// newPriorityClasses returns a PriorityClasses
func newPriorityClasses(c *SchedulingV1Client) *priorityClasses {
	return &priorityClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the priorityClass, and returns the corresponding priorityClass object, and an error if there is any.
func (c *priorityClasses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PriorityClass, err error) {
	result = &v1.PriorityClass{}
	err = c.client.Get().
		Resource("priorityclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PriorityClasses that match those selectors.
func (c *priorityClasses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PriorityClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PriorityClassList{}
	err = c.client.Get().
		Resource("priorityclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested priorityClasses.
func (c *priorityClasses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("priorityclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a priorityClass and creates it.  Returns the server's representation of the priorityClass, and an error, if there is any.
func (c *priorityClasses) Create(ctx context.Context, priorityClass *v1.PriorityClass, opts metav1.CreateOptions) (result *v1.PriorityClass, err error) {
	result = &v1.PriorityClass{}
	err = c.client.Post().
		Resource("priorityclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(priorityClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a priorityClass and updates it. Returns the server's representation of the priorityClass, and an error, if there is any.
func (c *priorityClasses) Update(ctx context.Context, priorityClass *v1.PriorityClass, opts metav1.UpdateOptions) (result *v1.PriorityClass, err error) {
	result = &v1.PriorityClass{}
	err = c.client.Put().
		Resource("priorityclasses").
		Name(priorityClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(priorityClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the priorityClass and deletes it. Returns an error if one occurs.
func (c *priorityClasses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("priorityclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *priorityClasses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("priorityclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched priorityClass.
func (c *priorityClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PriorityClass, err error) {
	result = &v1.PriorityClass{}
	err = c.client.Patch(pt).
		Resource("priorityclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied priorityClass.
func (c *priorityClasses) Apply(ctx context.Context, priorityClass *schedulingv1.PriorityClassApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PriorityClass, err error) {
	if priorityClass == nil {
		return nil, fmt.Errorf("priorityClass provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(priorityClass)
	if err != nil {
		return nil, err
	}
	name := priorityClass.Name
	if name == nil {
		return nil, fmt.Errorf("priorityClass.Name must be provided to Apply")
	}
	result = &v1.PriorityClass{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("priorityclasses").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
