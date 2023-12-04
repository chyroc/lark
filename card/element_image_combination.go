package card

import "github.com/chyroc/lark"

func ElementImageCombination(mode lark.MessageContentCardElementCombinationMode, images ...string) *lark.MessageContentCardElementImageCombination {
	res := &lark.MessageContentCardElementImageCombination{
		CombinationMode: mode,
	}
	res.SetImageList(removeNilString(images)...)
	return res
}
