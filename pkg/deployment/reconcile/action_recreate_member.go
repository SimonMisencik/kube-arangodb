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

package reconcile

import (
	"context"

	api "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v1"
	"github.com/arangodb/kube-arangodb/pkg/util/errors"
)

// newRecreateMemberAction creates a new Action that implements the given
// planned RecreateMember action.
func newRecreateMemberAction(action api.Action, actionCtx ActionContext) Action {
	a := &actionRecreateMember{}

	a.actionImpl = newActionImplDefRef(action, actionCtx)

	return a
}

// actionRecreateMember implements an RecreateMemberAction.
type actionRecreateMember struct {
	// actionImpl implement timeout and member id functions
	actionImpl

	// actionEmptyCheckProgress implement check progress with empty implementation
	actionEmptyCheckProgress
}

// Start performs the start of the action.
// Returns true if the action is completely finished, false in case
// the start time needs to be recorded and a ready condition needs to be checked.
func (a *actionRecreateMember) Start(ctx context.Context) (bool, error) {
	m, g, ok := a.actionCtx.GetMemberStatusAndGroupByID(a.action.MemberID)
	if !ok {
		return false, errors.Newf("expecting member to be present in list, but it is not")
	}

	cache, ok := a.actionCtx.ACS().ClusterCache(m.ClusterID)
	if !ok {
		return true, errors.Newf("Cluster is not ready")
	}

	switch g {
	case api.ServerGroupDBServers, api.ServerGroupAgents: // Only DBServers and Agents use persistent data
		_, ok := cache.PersistentVolumeClaim().V1().GetSimple(m.PersistentVolumeClaimName)
		if !ok {
			return false, errors.Newf("PVC is missing %s. Members won't be recreated without old PV", m.PersistentVolumeClaimName)
		}
	}

	if m.Phase == api.MemberPhaseFailed {
		// Change cluster phase to ensure it wont be removed
		m.Phase = api.MemberPhaseNone
	}

	if err := a.actionCtx.UpdateMember(ctx, m); err != nil {
		return false, errors.WithStack(err)
	}

	return true, nil
}
