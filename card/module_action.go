package card

import (
	"github.com/chyroc/lark"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
func Action(actions ...lark.MessageContentCardAction) *lark.MessageContentCardModuleAction {
	return &lark.MessageContentCardModuleAction{
		Actions: actions,
	}
}
