package test

import "context"

var ctx = context.Background()

func ptrStringNoNonePtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ptrValueString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func ptrString(s string) *string {
	return &s
}

func ptrInt(s int) *int {
	return &s
}

func ptrInt8(s int8) *int8 {
	return &s
}

func ptrInt16(s int16) *int16 {
	return &s
}

func ptrInt32(s int32) *int32 {
	return &s
}

func ptrInt64(s int64) *int64 {
	return &s
}

func ptrUInt(s uint) *uint {
	return &s
}

func ptrUInt8(s uint8) *uint8 {
	return &s
}

func ptrUInt16(s uint16) *uint16 {
	return &s
}

func ptrUInt32(s uint32) *uint32 {
	return &s
}

func ptrUInt64(s uint64) *uint64 {
	return &s
}

func ptrBool(s bool) *bool {
	return &s
}
