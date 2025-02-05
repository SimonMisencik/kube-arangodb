//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package resources

import "github.com/arangodb/kube-arangodb/pkg/logging"

// Resources is a service that creates low level resources for members
// and inspects low level resources, put the inspection result in members.
type Resources struct {
	log             logging.Logger
	namespace, name string
	context         Context

	metrics Metrics
}

// NewResources creates a new Resources service, used to
// create and inspect low level resources such as pods and services.
func NewResources(namespace, name string, context Context) *Resources {
	r := &Resources{
		namespace: namespace,
		name:      name,
		context:   context,
	}

	r.log = logger.WrapObj(r)

	return r
}
