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
	f, err := larkext.NewRootFolder(ctx, larkCli)
	as.Nil(err)

	t.Run("meta", func(t *testing.T) {
		meta1, err := larkext.NewFolder(larkCli, "").Meta(ctx)
		as.Nil(err)

		meta2, err := f.Meta(ctx)
		as.Nil(err)

		as.Equal(meta1.Token, meta2.Token)
	})

	t.Run("list", func(t *testing.T) {
		files, err := f.ListFiles(ctx)
		as.Nil(err)
		as.True(len(files) > 0)
	})

	t.Run("new-sheet sheet-self-delete", func(t *testing.T) {
		sheet, err := f.NewSheet(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(sheet.Delete(ctx))
	})

	t.Run("new-sheet folder-delete", func(t *testing.T) {
		sheet, err := f.NewSheet(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(f.DeleteSheet(ctx, sheet.SheetToken()))
	})

	t.Run("new-doc doc-self-delete", func(t *testing.T) {
		doc, err := f.NewDoc(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(doc.Delete(ctx))
	})

	t.Run("new-doc folder-delete", func(t *testing.T) {
		doc, err := f.NewDoc(ctx, fmt.Sprintf("rand %d", randInt64()))
		as.Nil(err)

		as.Nil(f.DeleteDoc(ctx, doc.DocToken()))
	})
}
