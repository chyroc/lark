package card

import (
	"github.com/chyroc/lark"
)

// option 有三种使用方式，分别是：selectMenu选项，selectMenu选人，overflow选项

// 基础使用
func SelectOption(text, value string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Text:  Text(text),
		Value: value,
	}
}

// 选人模式
// 指定待选人员，只需指定待选人的openId
func PersonOption(id string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Value: id,
	}
}

// 跳转链接 option
// overflow 中可用
func LinkOption(text, url string) *lark.MessageContentCardObjectOption {
	return &lark.MessageContentCardObjectOption{
		Text: Text(text),
		URL:  url,
	}
}
