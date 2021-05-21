package test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_Builder(t *testing.T) {
	as := assert.New(t)

	as.Equal(`<at user_id="1"></at>`, lark.AtBuilder.AtOpenID("1"))
	as.Equal(`<at user_id="all"></at>`, lark.AtBuilder.AtALL())
}
