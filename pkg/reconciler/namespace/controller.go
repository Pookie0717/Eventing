/*
Copyright 2019 The Knative Authors

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

package namespace

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	client "knative.dev/pkg/client/injection/kube/client"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	eventingclient "knative.dev/eventing/pkg/client/injection/client"
	"knative.dev/eventing/pkg/client/injection/informers/configs/v1alpha1/configmappropagation"
	"knative.dev/eventing/pkg/client/injection/informers/eventing/v1beta1/broker"
	"knative.dev/pkg/client/injection/kube/informers/core/v1/namespace"
	"knative.dev/pkg/client/injection/kube/informers/core/v1/serviceaccount"
	"knative.dev/pkg/client/injection/kube/informers/rbac/v1/rolebinding"
)

type envConfig struct {
	BrokerPullSecretName string `envconfig:"BROKER_IMAGE_PULL_SECRET_NAME" required:"false"`
}

const (
	// ReconcilerName is the name of the reconciler
	ReconcilerName = "Namespace" // TODO: Namespace is not a very good name for this controller.

	// controllerAgentName is the string used by this controller to identify
	// itself when creating events.
	controllerAgentName = "knative-eventing-namespace-controller"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)
	namespaceInformer := namespace.Get(ctx)
	serviceAccountInformer := serviceaccount.Get(ctx)
	roleBindingInformer := rolebinding.Get(ctx)
	brokerInformer := broker.Get(ctx)
	configMapPropagationInformer := configmappropagation.Get(ctx)

	recorder := controller.GetEventRecorder(ctx)
	if recorder == nil {
		// Create event broadcaster
		logger.Debug("Creating event broadcaster")
		eventBroadcaster := record.NewBroadcaster()
		watches := []watch.Interface{
			eventBroadcaster.StartLogging(logger.Named("event-broadcaster").Infof),
			eventBroadcaster.StartRecordingToSink(
				&v1.EventSinkImpl{Interface: client.Get(ctx).CoreV1().Events("")}),
		}
		recorder = eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})
		go func() {
			<-ctx.Done()
			for _, w := range watches {
				w.Stop()
			}
		}()
	}

	r := &Reconciler{
		eventingClientSet:          eventingclient.Get(ctx),
		kubeClientSet:              kubeclient.Get(ctx),
		namespaceLister:            namespaceInformer.Lister(),
		serviceAccountLister:       serviceAccountInformer.Lister(),
		roleBindingLister:          roleBindingInformer.Lister(),
		brokerLister:               brokerInformer.Lister(),
		configMapPropagationLister: configMapPropagationInformer.Lister(),
		recorder:                   recorder,
	}

	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		logger.Info("no broker image pull secret name defined")
	}
	r.brokerPullSecretName = env.BrokerPullSecretName

	impl := controller.NewImpl(r, logger, ReconcilerName)
	// TODO: filter label selector: on InjectionEnabledLabels()

	logger.Info("Setting up event handlers")
	namespaceInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	// Watch all the resources that this reconciler reconciles.
	serviceAccountInformer.Informer().AddEventHandler(
		cache.FilteringResourceEventHandler{
			FilterFunc: controller.FilterGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Namespace")),
			Handler:    controller.HandleAll(impl.EnqueueControllerOf),
		})
	roleBindingInformer.Informer().AddEventHandler(
		cache.FilteringResourceEventHandler{
			FilterFunc: controller.FilterGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Namespace")),
			Handler:    controller.HandleAll(impl.EnqueueControllerOf),
		})
	brokerInformer.Informer().AddEventHandler(
		cache.FilteringResourceEventHandler{
			FilterFunc: controller.FilterGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Namespace")),
			Handler:    controller.HandleAll(impl.EnqueueControllerOf),
		})
	configMapPropagationInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterGroupVersionKind(corev1.SchemeGroupVersion.WithKind("Namespace")),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
