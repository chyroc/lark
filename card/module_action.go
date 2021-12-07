package card

import (
	"github.com/chyroc/lark"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
func Action(actions ...lark.MessageContentCardElement) *lark.MessageContentCardModuleAction {
	return &lark.MessageContentCardModuleAction{
		Actions: actions,
	}
}
