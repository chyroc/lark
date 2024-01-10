// Code generated by lark_sdk_gen. DO NOT EDIT.
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
)

// GetMailPublicMailboxAliasList 获取所有公共邮箱别名。
//
// 该接口一次性返回所有数据, 分页参数无效
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/list
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/public-mailbox/public_mailbox-alias/list
func (r *MailService) GetMailPublicMailboxAliasList(ctx context.Context, request *GetMailPublicMailboxAliasListReq, options ...MethodOptionFunc) (*GetMailPublicMailboxAliasListResp, *Response, error) {
	if r.cli.mock.mockMailGetMailPublicMailboxAliasList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#GetMailPublicMailboxAliasList mock enable")
		return r.cli.mock.mockMailGetMailPublicMailboxAliasList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "GetMailPublicMailboxAliasList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getMailPublicMailboxAliasListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailGetMailPublicMailboxAliasList mock MailGetMailPublicMailboxAliasList method
func (r *Mock) MockMailGetMailPublicMailboxAliasList(f func(ctx context.Context, request *GetMailPublicMailboxAliasListReq, options ...MethodOptionFunc) (*GetMailPublicMailboxAliasListResp, *Response, error)) {
	r.mockMailGetMailPublicMailboxAliasList = f
}

// UnMockMailGetMailPublicMailboxAliasList un-mock MailGetMailPublicMailboxAliasList method
func (r *Mock) UnMockMailGetMailPublicMailboxAliasList() {
	r.mockMailGetMailPublicMailboxAliasList = nil
}

// GetMailPublicMailboxAliasListReq ...
type GetMailPublicMailboxAliasListReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱id或公共邮箱邮件地址, 示例值: "xxxxxx 或 xxx@xx.xxx"
}

// GetMailPublicMailboxAliasListResp ...
type GetMailPublicMailboxAliasListResp struct {
	Items []*GetMailPublicMailboxAliasListRespItem `json:"items,omitempty"` // 公共邮箱别名
}

// GetMailPublicMailboxAliasListRespItem ...
type GetMailPublicMailboxAliasListRespItem struct {
	PrimaryEmail string `json:"primary_email,omitempty"` // 主邮箱地址
	EmailAlias   string `json:"email_alias,omitempty"`   // 邮箱别名
}

// getMailPublicMailboxAliasListResp ...
type getMailPublicMailboxAliasListResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *GetMailPublicMailboxAliasListResp `json:"data,omitempty"`
}
