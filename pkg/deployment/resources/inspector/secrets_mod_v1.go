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

package inspector

import (
	"context"

	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	typedCore "k8s.io/client-go/kubernetes/typed/core/v1"

	secretv1 "github.com/arangodb/kube-arangodb/pkg/util/k8sutil/inspector/secret/v1"
)

func (p secretsMod) V1() secretv1.ModInterface {
	return secretsModV1(p)
}

type secretsModV1 struct {
	i *inspectorState
}

func (p secretsModV1) client() typedCore.SecretInterface {
	return p.i.Client().Kubernetes().CoreV1().Secrets(p.i.Namespace())
}

func (p secretsModV1) Create(ctx context.Context, secret *core.Secret, opts meta.CreateOptions) (*core.Secret, error) {
	if secret, err := p.client().Create(ctx, secret, opts); err != nil {
		return secret, err
	} else {
		p.i.GetThrottles().Secret().Invalidate()
		return secret, err
	}
}

func (p secretsModV1) Update(ctx context.Context, secret *core.Secret, opts meta.UpdateOptions) (*core.Secret, error) {
	if secret, err := p.client().Update(ctx, secret, opts); err != nil {
		return secret, err
	} else {
		p.i.GetThrottles().Secret().Invalidate()
		return secret, err
	}
}

func (p secretsModV1) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts meta.PatchOptions, subresources ...string) (result *core.Secret, err error) {
	if secret, err := p.client().Patch(ctx, name, pt, data, opts, subresources...); err != nil {
		return secret, err
	} else {
		p.i.GetThrottles().Secret().Invalidate()
		return secret, err
	}
}

func (p secretsModV1) Delete(ctx context.Context, name string, opts meta.DeleteOptions) error {
	if err := p.client().Delete(ctx, name, opts); err != nil {
		return err
	} else {
		p.i.GetThrottles().Secret().Invalidate()
		return err
	}
}
