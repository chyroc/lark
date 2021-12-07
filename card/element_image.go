package card

import (
	"github.com/chyroc/lark"
)

func ElementImage(image string) *lark.MessageContentCardElementImage {
	return &lark.MessageContentCardElementImage{
		ImgKey: image,
		Alt:    Text(""),
	}
}
