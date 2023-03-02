/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package internal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReflectToString(t *testing.T) {
	as := assert.New(t)

	as.Equal("1", ReflectToString(reflect.ValueOf(string("1"))))
	as.Equal("1", ReflectToString(reflect.ValueOf(int(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(int8(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(int16(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(int32(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(int64(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(uint(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(uint8(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(uint16(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(uint32(1))))
	as.Equal("1", ReflectToString(reflect.ValueOf(uint64(1))))
	as.Equal("true", ReflectToString(reflect.ValueOf(true)))
	as.Equal("false", ReflectToString(reflect.ValueOf(false)))
	as.Equal("<map[string]string Value>", ReflectToString(reflect.ValueOf(map[string]string{})))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrString(string("1")))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrInt(int(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrInt8(int8(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrInt16(int16(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrInt32(int32(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrInt64(int64(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrUInt(uint(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrUInt8(uint8(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrUInt16(uint16(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrUInt32(uint32(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptrUInt64(uint64(1)))))
	as.Equal("true", ReflectToString(reflect.ValueOf(ptrBool(true))))
	as.Equal("false", ReflectToString(reflect.ValueOf(ptrBool(false))))
}

func TestReflectToQueryString(t *testing.T) {
	tests := []struct {
		name  string
		args  reflect.Value
		wantS []string
	}{
		{"1", reflect.ValueOf(int(1)), []string{"1"}},
		{"1", reflect.ValueOf(int8(1)), []string{"1"}},
		{"1", reflect.ValueOf(int16(1)), []string{"1"}},
		{"1", reflect.ValueOf(int32(1)), []string{"1"}},
		{"1", reflect.ValueOf(int64(1)), []string{"1"}},
		{"1", reflect.ValueOf(uint(1)), []string{"1"}},
		{"1", reflect.ValueOf(uint8(1)), []string{"1"}},
		{"1", reflect.ValueOf(uint16(1)), []string{"1"}},
		{"1", reflect.ValueOf(uint32(1)), []string{"1"}},
		{"1", reflect.ValueOf(uint64(1)), []string{"1"}},
		{"1", reflect.ValueOf(true), []string{"true"}},
		{"1", reflect.ValueOf(false), []string{"false"}},
		{"1", reflect.ValueOf([]int{1, 2}), []string{"1", "2"}},
		{"1", reflect.ValueOf([...]int{1, 2}), []string{"1", "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := ReflectToQueryString(tt.args); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("ReflectToQueryString() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestIsInReflectKind(t *testing.T) {
	type args struct {
		v    reflect.Kind
		list []reflect.Kind
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{reflect.String, []reflect.Kind{reflect.String}}, true},
		{"1", args{reflect.String, []reflect.Kind{reflect.Int}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInReflectKind(tt.args.v, tt.args.list); got != tt.want {
				t.Errorf("IsInReflectKind() = %v, want %v", got, tt.want)
			}
		})
	}
}
