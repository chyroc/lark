package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Attendance_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		chatCli := cli.Attendance()

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.UpdateUserSettings(ctx, &lark.UpdateUserSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		chatCli := cli.Attendance()

		t.Run("", func(t *testing.T) {
			_, _, err := chatCli.UpdateUserSettings(ctx, &lark.UpdateUserSettingsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})
	})
}
