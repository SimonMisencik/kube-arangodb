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

package storage

import (
	"context"
	"time"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/arangodb/kube-arangodb/pkg/util/errors"
)

// inspectPVs queries all PersistentVolume's and triggers a cleanup for
// released volumes.
// Returns the number of available PV's.
func (ls *LocalStorage) inspectPVs() (int, error) {
	list, err := ls.deps.Client.Kubernetes().CoreV1().PersistentVolumes().List(context.Background(), meta.ListOptions{})
	if err != nil {
		return 0, errors.WithStack(err)
	}
	spec := ls.apiObject.Spec
	availableVolumes := 0
	cleanupBeforeTimestamp := time.Now().Add(time.Hour * -24)
	for _, pv := range list.Items {
		if pv.Spec.StorageClassName != spec.StorageClass.Name {
			// Not our storage class
			continue
		}
		switch pv.Status.Phase {
		case core.VolumeAvailable:
			// Is this an old volume?
			if pv.GetObjectMeta().GetCreationTimestamp().Time.Before(cleanupBeforeTimestamp) {
				// Let's clean it up
				if ls.isOwnerOf(&pv) {
					// Cleanup this volume
					ls.log.Str("name", pv.GetName()).Debug("Added PersistentVolume to cleaner")
					ls.pvCleaner.Add(pv)
				} else {
					ls.log.Str("name", pv.GetName()).Debug("PersistentVolume is not owned by us")
					availableVolumes++
				}
			} else {
				availableVolumes++
			}
		case core.VolumeReleased:
			if ls.isOwnerOf(&pv) {
				// Cleanup this volume
				ls.log.Str("name", pv.GetName()).Debug("Added PersistentVolume to cleaner")
				ls.pvCleaner.Add(pv)
			} else {
				ls.log.Str("name", pv.GetName()).Debug("PersistentVolume is not owned by us")
			}
		}
	}
	return availableVolumes, nil
}
