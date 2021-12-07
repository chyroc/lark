package card

import (
	"github.com/chyroc/lark"
)

// plain_text
func Text(text string) *lark.MessageContentCardObjectText {
	return &lark.MessageContentCardObjectText{
		Tag:     lark.MessageContentCardTextTypePlainText,
		Content: text,
	}
}

// lark_md
func MarkdownText(md string) *lark.MessageContentCardObjectText {
	return &lark.MessageContentCardObjectText{
		Tag:     lark.MessageContentCardTextTypeLarkMd,
		Content: md,
	}
}
