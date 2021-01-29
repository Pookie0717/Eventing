/*
Copyright 2021 The Knative Authors

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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "knative.dev/eventing/pkg/apis/eventing/v1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// TriggersGetter has a method to return a TriggerInterface.
// A group's client should implement this interface.
type TriggersGetter interface {
	Triggers(namespace string) TriggerInterface
}

// TriggerInterface has methods to work with Trigger resources.
type TriggerInterface interface {
	Create(ctx context.Context, trigger *v1.Trigger, opts metav1.CreateOptions) (*v1.Trigger, error)
	Update(ctx context.Context, trigger *v1.Trigger, opts metav1.UpdateOptions) (*v1.Trigger, error)
	UpdateStatus(ctx context.Context, trigger *v1.Trigger, opts metav1.UpdateOptions) (*v1.Trigger, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Trigger, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TriggerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Trigger, err error)
	TriggerExpansion
}

// triggers implements TriggerInterface
type triggers struct {
	client rest.Interface
	ns     string
}

// newTriggers returns a Triggers
func newTriggers(c *EventingV1Client, namespace string) *triggers {
	return &triggers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trigger, and returns the corresponding trigger object, and an error if there is any.
func (c *triggers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Trigger, err error) {
	result = &v1.Trigger{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("triggers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Triggers that match those selectors.
func (c *triggers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TriggerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TriggerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("triggers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested triggers.
func (c *triggers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("triggers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a trigger and creates it.  Returns the server's representation of the trigger, and an error, if there is any.
func (c *triggers) Create(ctx context.Context, trigger *v1.Trigger, opts metav1.CreateOptions) (result *v1.Trigger, err error) {
	result = &v1.Trigger{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("triggers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trigger).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a trigger and updates it. Returns the server's representation of the trigger, and an error, if there is any.
func (c *triggers) Update(ctx context.Context, trigger *v1.Trigger, opts metav1.UpdateOptions) (result *v1.Trigger, err error) {
	result = &v1.Trigger{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("triggers").
		Name(trigger.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trigger).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *triggers) UpdateStatus(ctx context.Context, trigger *v1.Trigger, opts metav1.UpdateOptions) (result *v1.Trigger, err error) {
	result = &v1.Trigger{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("triggers").
		Name(trigger.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(trigger).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the trigger and deletes it. Returns an error if one occurs.
func (c *triggers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("triggers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *triggers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("triggers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched trigger.
func (c *triggers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Trigger, err error) {
	result = &v1.Trigger{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("triggers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
