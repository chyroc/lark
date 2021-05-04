package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Contact(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		contactCli := cli.Contact()

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.CreateDepartment(ctx, &lark.CreateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetDepartment(ctx, &lark.GetDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetDepartmentList(ctx, &lark.GetDepartmentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := contactCli.SearchDepartment(ctx, &lark.SearchDepartmentReq{})
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateDepartment(ctx, &lark.UpdateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.CreateUser(ctx, &lark.CreateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.DeleteUser(ctx, &lark.DeleteUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetUser(ctx, &lark.GetUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetUserList(ctx, &lark.GetUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateUser(ctx, &lark.UpdateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateUserPatch(ctx, &lark.UpdateUserPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})
}
