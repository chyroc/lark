package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_FolderExt(t *testing.T) {
	as := assert.New(t)
	ctx := context.Background()
	larkCli := AppAllPermission.Ins()
	f := larkext.NewFolder(larkCli, "")

	t.Run("new-sheet", func(t *testing.T) {
		sheet, err := f.NewSheet(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		err = sheet.Delete(ctx)
		as.Nil(err)
	})

	// t.Run("meta", func(t *testing.T) {
	// 	meta, err := f.Meta(ctx)
	// 	as.Nil(err)
	// 	as.NotNil(meta)
	// })
}
