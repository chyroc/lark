package card

import (
	"github.com/chyroc/lark"
)

func StaticSelectMenu(options ...*lark.MessageContentCardObjectOption) *lark.MessageContentCardElementSelectMenu {
	return &lark.MessageContentCardElementSelectMenu{
		Tag:     lark.MessageContentCardElementTagSelectStatic,
		Options: options,
		Value:   nil,
	}
}

func PersonSelectMenuForIDs(ids ...string) *lark.MessageContentCardElementSelectMenu {
	var options []*lark.MessageContentCardObjectOption
	for _, v := range ids {
		options = append(options, PersonOption(v))
	}
	return &lark.MessageContentCardElementSelectMenu{
		Tag:     lark.MessageContentCardElementTagSelectPerson,
		Options: options,
		Value:   nil,
	}
}
