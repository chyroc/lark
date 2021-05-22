package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Auth(t *testing.T) {
	as := assert.New(t)
	if IsNotInCI() {
		return
	}

	ctx := context.Background()
	cli := AppALLPermission.Ins()

	resp, _, err := cli.Auth.GetUserInfo(ctx, &lark.GetUserInfoReq{}, lark.WithUserAccessToken(UserAdmin.AccessToken[AppALLPermission.AppID]))
	as.Nil(err)
	as.NotNil(resp)
	as.Equal(UserAdmin.OpenID, resp.OpenID)
	printData(resp)
}
