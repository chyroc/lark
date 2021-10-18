package test

import (
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Contact_Failed(t *testing.T) {
	as := assert.New(t)

	cli := AppAllPermission.Ins()
	moduleCli := cli.Contact

	t.Run("get-user", func(t *testing.T) {
		resp, _, err := moduleCli.GetUser(ctx, &lark.GetUserReq{
			UserIDType:       lark.IDTypePtr(lark.IDTypeUserID),
			DepartmentIDType: nil,
			UserID:           UserAdmin.UserID,
		})
		as.Nil(err)
		as.NotNil(resp)
		printData(resp)
		as.Equal(UserAdmin.OpenID, resp.User.OpenID)
		as.True(resp.User.IsTenantManager)
		as.True(resp.User.Status.IsActivated)
	})

	t.Run("get-dep", func(t *testing.T) {
		resp, _, err := moduleCli.GetDepartment(ctx, &lark.GetDepartmentReq{
			UserIDType:       lark.IDTypePtr(lark.IDTypeUserID),
			DepartmentIDType: lark.DepartmentIDTypePtr(lark.DepartmentIDTypeDepartmentID),
			DepartmentID:     "0",
		})
		as.Nil(err)
		as.NotNil(resp)
		printData(resp)
		as.True(resp.Department.MemberCount > 0)
		as.Equal("0", resp.Department.OpenDepartmentID)
	})

	t.Run("get-dep-list", func(t *testing.T) {
		resp, _, err := moduleCli.GetDepartmentList(ctx, &lark.GetDepartmentListReq{
			// UserIDType:         nil,
			DepartmentIDType:   nil,
			ParentDepartmentID: nil,
			FetchChild:         nil,
			PageToken:          nil,
			PageSize:           nil,
		})
		printData(resp)
		as.Nil(err)
		as.NotNil(resp)
		as.True(len(resp.Items) > 0)
	})

	if IsInCI() {
		t.Run("search-dep", func(t *testing.T) {
			resp, _, err := moduleCli.SearchDepartment(ctx, &lark.SearchDepartmentReq{
				UserIDType:       nil,
				DepartmentIDType: nil,
				PageToken:        nil,
				PageSize:         nil,
				Query:            "lark-sdk",
			}, lark.WithUserAccessToken(UserAdmin.AccessToken[AppAllPermission.AppID]))
			as.Nil(err)
			as.NotNil(resp)
			printData(resp)
		})
	}

	t.Run("", func(t *testing.T) {
		childOpenDepartmentID := ""
		{
			resp, _, err := moduleCli.CreateDepartment(ctx, &lark.CreateDepartmentReq{
				// ClientToken:        nil,
				// I18nName:           nil,
				// DepartmentID:       nil,
				// Order:              nil,
				// UnitIDs:            nil,
				// CreateGroupChat:    nil,
				UserIDType:         lark.IDTypePtr(lark.IDTypeUserID),
				DepartmentIDType:   lark.DepartmentIDTypePtr(lark.DepartmentIDTypeDepartmentID),
				Name:               "dep-" + randInt64String(),
				ParentDepartmentID: "0",
				LeaderUserID:       ptr.String(UserAdmin.UserID),
			})
			as.Nil(err)
			as.NotNil(resp)
			as.NotNil(resp.Department)
			childOpenDepartmentID = resp.Department.OpenDepartmentID
		}

		defer func() {
			var err error
			for i := 0; i < 10; i++ {
				_, _, err = moduleCli.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{
					// UserIDType:       lark,
					DepartmentIDType: lark.DepartmentIDTypePtr(lark.DepartmentIDTypeOpenDepartmentID),
					DepartmentID:     childOpenDepartmentID,
				})
				if err == nil {
					break
				}
			}
			as.Nil(err)
		}()

		{
			resp, _, err := moduleCli.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{
				// UserIDType:       nil,
				DepartmentIDType: lark.DepartmentIDTypePtr(lark.DepartmentIDTypeOpenDepartmentID),
				DepartmentID:     childOpenDepartmentID,
				// PageToken:        ptr.String(""),
				// PageSize:         ptr.Int64(10),
			})
			printData(resp)
			as.Nil(err)
			as.NotNil(resp)
			// as.Len(resp.Items, 1)
			// as.Equal("0", resp.Items[0].OpenDepartmentID)
		}

		{
			resp, _, err := moduleCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{
				// UserIDType:         nil,
				DepartmentIDType: lark.DepartmentIDTypePtr(lark.DepartmentIDTypeOpenDepartmentID),
				DepartmentID:     childOpenDepartmentID,
				Name:             ptr.String("dep-update-" + randInt64String()),
				// I18nName:           nil,
				// ParentDepartmentID: nil,
				// LeaderUserID:       nil,
				// Order:              nil,
				// UnitIDs:            nil,
				// CreateGroupChat:    nil,
			})
			printData(resp)
			as.Nil(err)
			as.NotNil(resp)
		}

		{
			resp, _, err := moduleCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{
				// UserIDType:         nil,
				DepartmentIDType: lark.DepartmentIDTypePtr(lark.DepartmentIDTypeOpenDepartmentID),
				DepartmentID:     childOpenDepartmentID,
				Name:             ptr.String("dep-patch-" + randInt64String()),
				// I18nName:           nil,
				// ParentDepartmentID: nil,
				// LeaderUserID:       nil,
				// Order:              nil,
				// UnitIDs:            nil,
				// CreateGroupChat:    nil,
			})
			printData(resp)
			as.Nil(err)
			as.NotNil(resp)
		}
	})
}
