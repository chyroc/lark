/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Auth(t *testing.T) {
	SkipRealAPITest(t)
	t.Skip()
	as := assert.New(t)

	if IsInCI() {
		cli := AppAllPermission.Ins()

		resp, _, err := cli.Auth.GetUserInfo(ctx, &lark.GetUserInfoReq{}, lark.WithUserAccessToken(UserAdmin.AccessToken[AppAllPermission.AppID]))
		as.Nil(err)
		as.NotNil(resp)
		as.Equal(UserAdmin.OpenID, resp.OpenID)
		printData(resp)
	}
}

func Test_GenOauth(t *testing.T) {
	SkipRealAPITest(t)
	fmt.Println(AppAllPermission.Ins().Auth.GenOAuthURL(ctx, &lark.GenOAuthURLReq{
		RedirectURI: "http://127.0.0.1:5000/",
		State:       "",
	}))

	return
	// AppAllPermission.Ins().Auth.GetAccessToken(ctx,&lark.)
}
