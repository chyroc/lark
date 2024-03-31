package helper

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseAST(t *testing.T) {
	t.Run("", func(t *testing.T) {
		dir := "/Users/chyroc/code/chyroc/lark/.github/script/helper/fortest"
		res, err := ParseAST(dir)
		if err != nil {
			panic(err)
		}
		spew.Dump(res.Consts)
	})
}
