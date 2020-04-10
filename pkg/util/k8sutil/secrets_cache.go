//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
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
// Author Ewout Prangsma
//

package k8sutil

import (
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// secretsCache implements a cached version of a SecretInterface.
// It is NOT go-routine safe.
type secretsCache struct {
	cli   corev1.SecretInterface
	cache []v1.Secret
}

// NewSecretCache creates a cached version of the given SecretInterface.
func NewSecretCache(cli corev1.SecretInterface) SecretInterface {
	return &secretsCache{cli: cli}
}

var (
	secretGroupResource = schema.GroupResource{
		Group:    v1.GroupName,
		Resource: "Secret",
	}
)

func (sc *secretsCache) Create(s *v1.Secret) (*v1.Secret, error) {
	sc.cache = nil
	result, err := sc.cli.Create(s)
	if err != nil {
		return nil, maskAny(err)
	}
	return result, nil
}

func (sc *secretsCache) Delete(name string, options *metav1.DeleteOptions) error {
	sc.cache = nil
	err := sc.cli.Delete(name, options)
	if err != nil {
		return maskAny(err)
	}
	return nil
}

func (sc *secretsCache) Get(name string, options metav1.GetOptions) (*v1.Secret, error) {
	if sc.cache == nil {
		list, err := sc.cli.List(metav1.ListOptions{})
		if err != nil {
			return nil, maskAny(err)
		}
		sc.cache = list.Items
	}
	for _, s := range sc.cache {
		if s.GetName() == name {
			return &s, nil
		}
	}
	return nil, maskAny(apierrors.NewNotFound(secretGroupResource, name))
}
