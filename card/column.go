package card

import (
	"github.com/chyroc/lark"
)

func Column(modules ...lark.MessageContentCardModule) *lark.MessageContentCardModuleColumn {
	return &lark.MessageContentCardModuleColumn{
		Modules:       removeNilMessageContentCardModule(modules),
		Width:         nil,
		Weight:        nil,
		VerticalAlign: nil,
	}
}
