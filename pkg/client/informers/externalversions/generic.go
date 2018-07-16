/*
Copyright 2018 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/knative/eventing/pkg/apis/channels/v1alpha1"
	feeds_v1alpha1 "github.com/knative/eventing/pkg/apis/feeds/v1alpha1"
	flows_v1alpha1 "github.com/knative/eventing/pkg/apis/flows/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=channels.knative.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("buses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Channels().V1alpha1().Buses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("channels"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Channels().V1alpha1().Channels().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterbuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Channels().V1alpha1().ClusterBuses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("subscriptions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Channels().V1alpha1().Subscriptions().Informer()}, nil

		// Group=feeds.knative.dev, Version=v1alpha1
	case feeds_v1alpha1.SchemeGroupVersion.WithResource("clustereventsources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Feeds().V1alpha1().ClusterEventSources().Informer()}, nil
	case feeds_v1alpha1.SchemeGroupVersion.WithResource("clustereventtypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Feeds().V1alpha1().ClusterEventTypes().Informer()}, nil
	case feeds_v1alpha1.SchemeGroupVersion.WithResource("eventsources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Feeds().V1alpha1().EventSources().Informer()}, nil
	case feeds_v1alpha1.SchemeGroupVersion.WithResource("eventtypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Feeds().V1alpha1().EventTypes().Informer()}, nil
	case feeds_v1alpha1.SchemeGroupVersion.WithResource("feeds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Feeds().V1alpha1().Feeds().Informer()}, nil

		// Group=flows.knative.dev, Version=v1alpha1
	case flows_v1alpha1.SchemeGroupVersion.WithResource("flows"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Flows().V1alpha1().Flows().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
