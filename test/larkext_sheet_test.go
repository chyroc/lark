package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_SheetExt(t *testing.T) {
	as := assert.New(t)

	ctx := context.Background()
	larkCli := AppAllPermission.Ins()

	folderIns := larkext.NewFolder(larkCli, "")

	sheetIns, err := folderIns.NewSheet(ctx, "")
	as.Nil(err)

	err = sheetIns.SetTitle(ctx, fmt.Sprintf("sheet title %d", randInt64()))
	as.Nil(err)

	_, _, err = larkCli.Drive.UpdateDriveMemberPermission(context.Background(), &lark.UpdateDriveMemberPermissionReq{
		NeedNotification: ptr.Bool(true),
		Type:             "sheet",
		Token:            sheetIns.SheetToken(),
		MemberID:         UserAdmin.OpenID,
		MemberType:       "openid",
		Perm:             "full_access",
	})
	as.Nil(err)

	meta, err := sheetIns.Meta(ctx)
	as.Nil(err)

	defaultSheetID := meta.Sheets[0].SheetID

	sheetID, err := sheetIns.NewSheet(ctx, fmt.Sprintf("sheet %d", randInt64()))
	as.Nil(err)
	as.NotEmpty(sheetID)

	err = sheetIns.DeleteSheet(ctx, sheetID)
	as.Nil(err)

	err = sheetIns.InsertRows(ctx, defaultSheetID, 0, 1)
	as.Nil(err)
	err = sheetIns.InsertCols(ctx, defaultSheetID, 0, 1)
	as.Nil(err)

	// spew.Dump(ins.SearchSheet(ctx, defaultSheetID, "22", nil))
	// spew.Dump(ins.CopySheet(ctx, "3xtK6c", ptr.String("xxxxx")))
	// spew.Dump(ins.CopySheet(ctx, "72e2fa", nil))
}
