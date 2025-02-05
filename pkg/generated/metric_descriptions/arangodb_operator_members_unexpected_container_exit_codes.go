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

package metric_descriptions

import "github.com/arangodb/kube-arangodb/pkg/util/metrics"

var (
	arangodbOperatorMembersUnexpectedContainerExitCodes = metrics.NewDescription("arangodb_operator_members_unexpected_container_exit_codes", "Counter of unexpected restarts in pod (Containers/InitContainers/EphemeralContainers)", []string{`namespace`, `name`, `member`, `container`, `container_type`, `code`}, nil)
)

func init() {
	registerDescription(arangodbOperatorMembersUnexpectedContainerExitCodes)
}

func ArangodbOperatorMembersUnexpectedContainerExitCodes() metrics.Description {
	return arangodbOperatorMembersUnexpectedContainerExitCodes
}

func ArangodbOperatorMembersUnexpectedContainerExitCodesCounter(value float64, namespace string, name string, member string, container string, containerType string, code string) metrics.Metric {
	return ArangodbOperatorMembersUnexpectedContainerExitCodes().Gauge(value, namespace, name, member, container, containerType, code)
}
