package card

import (
	"github.com/chyroc/lark"
)

// 注意：note 只能填充 text 和 image 对象
func Note(elements ...lark.MessageContentCardElement) *lark.MessageContentCardModuleNote {
	return &lark.MessageContentCardModuleNote{
		Elements: elements,
	}
}
