package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Auth(t *testing.T) {
	as := assert.New(t)

	if IsInCI() {
		cli := AppALLPermission.Ins()

		resp, _, err := cli.Auth.GetUserInfo(ctx, &lark.GetUserInfoReq{}, lark.WithUserAccessToken(UserAdmin.AccessToken[AppALLPermission.AppID]))
		as.Nil(err)
		as.NotNil(resp)
		as.Equal(UserAdmin.OpenID, resp.OpenID)
		printData(resp)
	}
}

func Test_GenOauth(t *testing.T) {
	fmt.Println(AppALLPermission.Ins().Auth.GenOAuthURL(ctx, &lark.GenOAuthURLReq{
		RedirectURI: "http://127.0.0.1:5000/",
		State:       "",
	}))

	return
	// AppALLPermission.Ins().Auth.GetAccessToken(ctx,&lark.)
}
