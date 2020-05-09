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
	"fmt"
	"reflect"
)

type StringValue struct {
	// +optional
	Direct string `json:"direct,omitempty"`
	// +optional
	Indirect *ObjectFieldReference `json:"indirect,omitempty"`
}

func (sv *StringValue) Value(ctx context.Context) (string, error) {
	if len(sv.Direct) > 0 {
		return sv.Direct, nil
	}
	in, err := sv.Indirect.Value(ctx)
	if err != nil {
		return "", err
	}
	if in == nil {
		return "", nil
	}
	s, ok := in.(string)
	if !ok {
		ts := reflect.TypeOf(in).String()
		return "", fmt.Errorf("Type of ObjectFieldReference' Value in not 'string' but '%s'", ts)
	}
	return s, nil
}
