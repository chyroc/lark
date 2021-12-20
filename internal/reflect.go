package internal

import (
	"reflect"
	"strconv"
)

// ReflectToString ...
func ReflectToString(v reflect.Value) (s string) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	default:
		return v.String()
	}
}

// ReflectToQueryString ...
func ReflectToQueryString(v reflect.Value) (s []string) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.String:
		return []string{v.String()}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []string{strconv.FormatInt(v.Int(), 10)}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []string{strconv.FormatUint(v.Uint(), 10)}
	case reflect.Bool:
		return []string{strconv.FormatBool(v.Bool())}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			s = append(s, ReflectToString(v.Index(i)))
		}
		return s
	default:
		return []string{v.String()}
	}
}

// IsInReflectKind ...
func IsInReflectKind(v reflect.Kind, list []reflect.Kind) bool {
	for _, vv := range list {
		if vv == v {
			return true
		}
	}
	return false
}
