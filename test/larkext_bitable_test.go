package test

import (
	"testing"

	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_LarkExt_Bitable(t *testing.T) {
	as := assert.New(t)
	as.Nil(nil)

	larkClient := AppAllPermission.Ins()
	// folderClient, err := larkext.NewRootFolder(ctx, larkClient)
	// as.Nil(err)

	larkext.NewBitable(larkClient, "x")
}
