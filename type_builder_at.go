package lark

import (
	"fmt"
)

// AtBuilder ...
var AtBuilder atBuilder

type atBuilder string

// AtAll ...
func (r atBuilder) AtAll() string {
	return `<at user_id="all"></at>`
}

// AtOpenID ...
func (r atBuilder) AtOpenID(openID string) string {
	return fmt.Sprintf(`<at user_id=%q></at>`, openID)
}
