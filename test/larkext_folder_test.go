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

	sheet, err := f.NewSheet(ctx, fmt.Sprintf("rand %d", randInt64()))
	as.Nil(err)

	err = sheet.Delete(ctx)
	as.Nil(err)
}
