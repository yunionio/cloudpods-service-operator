// Copyright 2020 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-yaml/yaml"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +kubebuilder:object:generate=false
type IValue interface {
	IsZero() bool
	Interface() interface{}
}

// +kubebuilder:object:generate=false
type IStore interface {
	GetValue(ctx context.Context) (IValue, error)
}

type StringStore struct {
	// +optional
	Value string `json:"value,omitempty"`
	// +optional
	Reference *ObjectFieldReference `json:"reference,omitempty"`
}

type String string

func (sv String) IsZero() bool {
	return sv == ""
}

func (sv String) String() string {
	return string(sv)
}

func (sv String) Interface() interface{} {
	return sv.String()
}

func (st StringStore) GetValue(ctx context.Context) (IValue, error) {
	if len(st.Value) > 0 {
		return String(st.Value), nil
	}
	in, err := st.Reference.Value(ctx)
	if err != nil || in == nil{
		return nil, err
	}
	s, ok := in.(string)
	if !ok {
		ts := reflect.TypeOf(in).String()
		return nil, fmt.Errorf("Type of ObjectFieldReference' Value in not 'string' but '%s'", ts)
	}
	return String(s), nil
}

type IntOrStringStore struct {
	// +optional
	Value *IntOrString `json:"value,omitempty"`
	// +optional
	Reference *ObjectFieldReference `json:"reference,omitempty"`
}

type IntOrString struct {
	intstr.IntOrString `json:",inline"`
}

func (isv IntOrString) String() (string, bool) {
	if isv.Type == intstr.String {
		return isv.StrVal, true
	}
	return "", false
}

func (isv IntOrString) Int() (int32, bool) {
	if isv.Type == intstr.Int {
		return isv.IntVal, true
	}
	return 0, false
}

func (isv IntOrString) IsZero() bool {
	if s, ok := isv.String(); ok {
		return s == ""
	}
	if i, ok := isv.Int(); ok {
		return i == 0
	}
	return true
}

func (isv IntOrString) Interface() interface{} {
	if s, ok := isv.String(); ok {
		return s
	}
	if i, ok := isv.Int(); ok {
		return i
	}
	return isv
}

func (ist IntOrStringStore) GetValue(ctx context.Context) (IValue, error) {
	if ist.Value != nil {
		return *ist.Value, nil
	}
	in, err := ist.Reference.Value(ctx)
	if err != nil || in == nil {
		return nil, err
	}

	var is intstr.IntOrString
	switch v := in.(type) {
	case string:
		is = intstr.FromString(v)
	case int:
		is = intstr.FromInt(v)
	case int32:
		is = intstr.FromInt(int(v))
	case int64:
		is = intstr.FromInt(int(v))
	default:
		ts := reflect.TypeOf(in).String()
		return nil, fmt.Errorf("Type of ObjectFieldReference' Value in not 'string' but '%s'", ts)
	}
	return IntOrString{is}, nil
}

type IntOrStringOrYamlStore struct {
	// IsYaml determines whether the string in IntOrStringStore is a yaml string
	// +optional
	IsYaml *bool `json:"isYaml,omitempty"`

	IntOrStringStore `json:",inline"`
}

type Yaml []byte

func (y Yaml) MarshalYAML() (interface{}, error) {
	s := y
	if s[0] == '-' {
		var in []map[string]interface{}
		err := yaml.Unmarshal(s, &in)
		if err != nil {
			return nil, err
		}
		return in, nil
	}
	var in map[string]interface{}
	err := yaml.Unmarshal(s, &in)
	if err != nil {
		return string(s), nil
	}
	return in, nil
}

func (y Yaml) IsZero() bool {
	return len(y) == 0
}

func (y Yaml) Interface() interface{} {
	return y
}

func (isys IntOrStringOrYamlStore) GetValue(ctx context.Context) (IValue, error) {
	value, err := isys.IntOrStringStore.GetValue(ctx)
	if err != nil {
		return nil, err
	}
	if isys.IsYaml == nil || !*isys.IsYaml {
		return value, nil
	}
	is := value.(IntOrString)
	s, ok := is.String()
	if !ok {
		return nil, fmt.Errorf("Not support yaml string for integer")
	}
	return Yaml(s), nil
}
