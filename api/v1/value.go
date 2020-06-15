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
	"strings"

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
	if err != nil {
		return nil, err
	}
	if in == nil {
		return nil, nil
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
	if err != nil {
		return nil, err
	}
	if in == nil {
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

//type YamlStore struct {
//	// +optional
//	Value *Yaml `json:"value,omitempty"`
//	// +optional
//	Reference *ObjectFieldReference `json:"reference,omitempty"`
//}
//
//func (ys YamlStore) GetValue(ctx context.Context) (IValue, error) {
//	if ys.Value != nil {
//		return ys.Value, nil
//	}
//	in, err := ys.Reference.Value(ctx)
//	if err != nil {
//		return nil, err
//	}
//	if in == nil {
//		return nil, err
//	}
//
//	raw, err := yaml.Marshal(in)
//	if err != nil {
//		return nil, err
//	}
//	return Yaml{runtime.RawExtension{Raw: raw}}, nil
//}
//
//type Yaml struct {
//	runtime.RawExtension `json:",inline"`
//}
//
//func (y Yaml) IsZero() bool {
//	return len(y.Raw) == 0
//}
//
//func (y Yaml) Interface() interface{} {
//	return y
//}
//
//func (y Yaml) MarshalYAML() (interface{}, error) {
//	switch  {
//	case y.Raw[0] == '"':
//		var s string
//		err := yaml.Unmarshal(y.Raw, &s)
//		if err != nil {
//			return nil, err
//		}
//		return s, nil
//	case y.Raw[0] == '-':
//		var s []map[string]interface{}
//		err := yaml.Unmarshal(y.Raw, &s)
//		if err != nil {
//			return nil, err
//		}
//		return s, nil
//	case !bytes.Contains(y.Raw, []byte{':'}):
//		var s int
//		err := yaml.Unmarshal(y.Raw, &s)
//		if err != nil {
//			return nil, err
//		}
//		return s, nil
//	default:
//		var s map[string]interface{}
//		err := yaml.Unmarshal(y.Raw, &s)
//		if err != nil {
//			return nil, err
//		}
//		return s, nil
//	}
//}

type IntOrStringOrYamlStore struct {
	// +optional
	Value *IntOrStringOrYaml `json:"value,omitempty"`
	// +optional
	Reference *ObjectFieldReference `json:"reference,omitempty"`
}

type IntOrStringOrYaml struct {
	IntOrString `json:",inline"`
}

type Yaml []byte

func (y *Yaml) MarshalYAML() (interface{}, error) {
	s := *y
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

func (isy IntOrStringOrYaml) Interface() interface{} {
	if s, ok := isy.String(); ok {
		if !strings.Contains(s, ":") && !strings.Contains(s, "\n") {
			return s
		}
		y := Yaml(s)
		return &y
	}
	if i, ok := isy.Int(); ok {
		return i
	}
	return isy
}

func (isys IntOrStringOrYamlStore) GetValue(ctx context.Context) (IValue, error) {
	var iss IntOrStringStore
	if isys.Value == nil {
		iss = IntOrStringStore{
			Reference: isys.Reference,
		}
	} else {
		iss = IntOrStringStore{
			Value:     &isys.Value.IntOrString,
			Reference: isys.Reference,
		}
	}
	value, err := iss.GetValue(ctx)
	if err != nil {
		return nil, err
	}
	is := value.(IntOrString)
	return IntOrStringOrYaml{IntOrString: is}, nil
}
