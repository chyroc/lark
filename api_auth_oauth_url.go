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
package lark

import (
	"context"
	"fmt"
	"net/url"
)

// GenOAuthURL 请求用户身份验证
//
// 应用请求用户身份验证时，需按如下方式构造登录链接，并引导用户跳转至此链接。飞书客户端内用户免登，系统浏览器内用户需完成扫码登录。登录成功后会生成登录预授权码 code，并作为参数重定向到重定向URL。 详细开发指南：第三方网站免登
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN
func (r *AuthService) GenOAuthURL(ctx context.Context, request *GenOAuthURLReq) string {
	return fmt.Sprintf(
		r.cli.openBaseURL+`/open-apis/authen/v1/index?redirect_uri=%s&app_id=%s&state=%s`,
		url.QueryEscape(request.RedirectURI),
		r.cli.appID,
		url.QueryEscape(request.State),
	)
}

// GenOAuthURLReq ...
type GenOAuthURLReq struct {
	RedirectURI string `json:"redirect_uri,omitempty"` // 在本流程中，此值为 authorization_code
	State       string `json:"state,omitempty"`        // 来自[请求身份验证(新)](/ssl:ttdoc/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN)流程，用户扫码登录后会自动302到redirect_uri并带上此参数
}
