/*


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

package v1

import (
	"context"

	"github.com/mcuadros/go-lookup"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"yunion.io/x/pkg/errors"
)

var (
	schemei *runtime.Scheme
	clienti client.Client
)

func InitReferenceManager(client client.Client, scheme *runtime.Scheme) {
	schemei = scheme
	clienti = client
}

type ObjectFieldReference struct {
	// +optional
	Group string `json:"group,omitempty"`
	// +optional
	Version   string `json:"version,omitempty"`
	Kind      string `json:"kind"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	FieldPath string `json:"fieldPath"`
}

func (fr *ObjectFieldReference) GroupVersionKind() schema.GroupVersionKind {
	gvk := schema.GroupVersionKind{
		Group:   fr.Group,
		Version: fr.Version,
		Kind:    fr.Kind,
	}
	if len(gvk.Group) == 0 {
		gvk.Group = "onecloud.yunion.io"
	}
	if len(gvk.Version) == 0 {
		gvk.Version = "v1"
	}
	return gvk
}

func (fr *ObjectFieldReference) NamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: fr.Namespace,
		Name:      fr.Name,
	}
}

func (fr *ObjectFieldReference) String() string {
	return "(" + fr.GroupVersionKind().String() + ")." + fr.NamespacedName().String() + "." + fr.FieldPath
}

func (fr *ObjectFieldReference) Value(ctx context.Context) (interface{}, error) {
	obj, err := schemei.New(fr.GroupVersionKind())
	if err != nil {
		return nil, errors.Wrap(err, "scheme.New")
	}
	err = clienti.Get(ctx, fr.NamespacedName(), obj)
	if err != nil {
		return nil, errors.Wrap(err, "client.Get")
	}
	value, err := lookup.LookupString(obj, fr.FieldPath)
	if err != nil {
		if err == lookup.ErrIndexOutOfRange {
			return nil, nil
		}
		return nil, errors.Wrap(err, "lookup.Lookup")
	}
	return value.Interface(), nil
}
