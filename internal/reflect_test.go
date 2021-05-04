package internal

import (
	"reflect"
	"testing"

	"github.com/chyroc/go-ptr"
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
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.String(string("1")))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.Int(int(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.Int8(int8(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.Int16(int16(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.Int32(int32(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.Int64(int64(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.UInt(uint(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.UInt8(uint8(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.UInt16(uint16(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.UInt32(uint32(1)))))
	as.Equal("1", ReflectToString(reflect.ValueOf(ptr.UInt64(uint64(1)))))
	as.Equal("true", ReflectToString(reflect.ValueOf(ptr.Bool(true))))
	as.Equal("false", ReflectToString(reflect.ValueOf(ptr.Bool(false))))
}
