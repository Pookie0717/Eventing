/*
Copyright 2020 The Knative Authors

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
	context "context"

	informers "k8s.io/client-go/informers"
	fake "knative.dev/pkg/client/injection/kube/client/fake"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	factory "knative.dev/pkg/injection/clients/namespacedkube/informers/factory"
	"knative.dev/pkg/system"
)

var Get = factory.Get

func init() {
	injection.Fake.RegisterInformerFactory(withInformerFactory)
}

func withInformerFactory(ctx context.Context) context.Context {
	c := fake.Get(ctx)
	return context.WithValue(ctx, factory.Key{},
		informers.NewSharedInformerFactoryWithOptions(c, controller.GetResyncPeriod(ctx),
			// This factory scopes things to the system namespace.
			informers.WithNamespace(system.Namespace())))
}
