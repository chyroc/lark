package card

import (
	"github.com/chyroc/lark"
)

func ColumnSet(columns ...*lark.MessageContentCardModuleColumn) *lark.MessageContentCardModuleColumnSet {
	flexMode := "none"
	bgStyle := "default"
	hSpacing := "default"
	return &lark.MessageContentCardModuleColumnSet{
		Columns:           removeNilMessageContentCardModuleColumn(columns),
		FlexMode:          &flexMode,
		BackgroundStyle:   &bgStyle,
		HorizontalSpacing: &hSpacing,
	}
}
