package card

import (
	"unsafe"

	"github.com/chyroc/lark"
)

type face struct {
	x    uintptr
	data unsafe.Pointer
}

func isNil(v any) bool {
	return (*face)(unsafe.Pointer(&v)).data == nil
}

func removeNilMessageContentCardModule(data []lark.MessageContentCardModule) []lark.MessageContentCardModule {
	res := make([]lark.MessageContentCardModule, 0, len(data))
	for _, v := range data {
		if isNil(v) {
			continue
		}
		res = append(res, v)
	}
	return res
}

func removeNilMessageContentCardElement(data []lark.MessageContentCardElement) []lark.MessageContentCardElement {
	res := make([]lark.MessageContentCardElement, 0, len(data))
	for _, v := range data {
		if isNil(v) {
			continue
		}
		res = append(res, v)
	}
	return res
}

func removeNilMessageContentCardObjectOption(data []*lark.MessageContentCardObjectOption) []*lark.MessageContentCardObjectOption {
	res := make([]*lark.MessageContentCardObjectOption, 0, len(data))
	for _, v := range data {
		if isNil(v) {
			continue
		}
		res = append(res, v)
	}
	return res
}

func removeNilMessageContentCardModuleColumn(data []*lark.MessageContentCardModuleColumn) []*lark.MessageContentCardModuleColumn {
	res := make([]*lark.MessageContentCardModuleColumn, 0, len(data))
	for _, v := range data {
		if isNil(v) {
			continue
		}
		res = append(res, v)
	}
	return res
}

func removeNilString(data []string) []string {
	res := make([]string, 0, len(data))
	for _, v := range data {
		if isNil(v) {
			continue
		}
		res = append(res, v)
	}
	return res
}
