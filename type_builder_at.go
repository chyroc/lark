package lark

import (
	"fmt"
)

var AtBuilder atBuilder

type atBuilder string

func (r atBuilder) AtAll() string {
	return `<at user_id="all"></at>`
}

func (r atBuilder) AtOpenID(openID string) string {
	return fmt.Sprintf(`<at user_id=%q></at>`, openID)
}
