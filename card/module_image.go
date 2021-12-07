package card

import (
	"github.com/chyroc/lark"
)

func ModuleImage(image string) *lark.MessageContentCardModuleImage {
	return &lark.MessageContentCardModuleImage{
		ImgKey: image,
		Alt: Text(""),
	}
}
