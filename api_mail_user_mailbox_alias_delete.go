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

// DeleteMailUserMailboxAlias 删除用户邮箱别名。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/user_mailbox-alias/delete
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/user_mailbox-alias/delete-2
func (r *MailService) DeleteMailUserMailboxAlias(ctx context.Context, request *DeleteMailUserMailboxAliasReq, options ...MethodOptionFunc) (*DeleteMailUserMailboxAliasResp, *Response, error) {
	if r.cli.mock.mockMailDeleteMailUserMailboxAlias != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailUserMailboxAlias mock enable")
		return r.cli.mock.mockMailDeleteMailUserMailboxAlias(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "DeleteMailUserMailboxAlias",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/user_mailboxes/:user_mailbox_id/aliases/:alias_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteMailUserMailboxAliasResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailDeleteMailUserMailboxAlias mock MailDeleteMailUserMailboxAlias method
func (r *Mock) MockMailDeleteMailUserMailboxAlias(f func(ctx context.Context, request *DeleteMailUserMailboxAliasReq, options ...MethodOptionFunc) (*DeleteMailUserMailboxAliasResp, *Response, error)) {
	r.mockMailDeleteMailUserMailboxAlias = f
}

// UnMockMailDeleteMailUserMailboxAlias un-mock MailDeleteMailUserMailboxAlias method
func (r *Mock) UnMockMailDeleteMailUserMailboxAlias() {
	r.mockMailDeleteMailUserMailboxAlias = nil
}

// DeleteMailUserMailboxAliasReq ...
type DeleteMailUserMailboxAliasReq struct {
	UserMailboxID string `path:"user_mailbox_id" json:"-"` // 用户邮箱地址, 示例值: "user@xxx.xx"
	AliasID       string `path:"alias_id" json:"-"`        // 别名邮箱地址, 示例值: "user_alias@xxx.xx"
}

// DeleteMailUserMailboxAliasResp ...
type DeleteMailUserMailboxAliasResp struct {
}

// deleteMailUserMailboxAliasResp ...
type deleteMailUserMailboxAliasResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailUserMailboxAliasResp `json:"data,omitempty"`
}
