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

// UpdatePublicMailbox 更新公共邮箱所有信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/update
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/public-mailbox/public_mailbox/update
func (r *MailService) UpdatePublicMailbox(ctx context.Context, request *UpdatePublicMailboxReq, options ...MethodOptionFunc) (*UpdatePublicMailboxResp, *Response, error) {
	if r.cli.mock.mockMailUpdatePublicMailbox != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#UpdatePublicMailbox mock enable")
		return r.cli.mock.mockMailUpdatePublicMailbox(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "UpdatePublicMailbox",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updatePublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailUpdatePublicMailbox mock MailUpdatePublicMailbox method
func (r *Mock) MockMailUpdatePublicMailbox(f func(ctx context.Context, request *UpdatePublicMailboxReq, options ...MethodOptionFunc) (*UpdatePublicMailboxResp, *Response, error)) {
	r.mockMailUpdatePublicMailbox = f
}

// UnMockMailUpdatePublicMailbox un-mock MailUpdatePublicMailbox method
func (r *Mock) UnMockMailUpdatePublicMailbox() {
	r.mockMailUpdatePublicMailbox = nil
}

// UpdatePublicMailboxReq ...
type UpdatePublicMailboxReq struct {
	PublicMailboxID string  `path:"public_mailbox_id" json:"-"` // 公共邮箱唯一标识或公共邮箱地址, 示例值: "xxxxxxxxxxxxxxx 或 test_public_mailbox@xxx.xx"
	Email           *string `json:"email,omitempty"`            // 公共邮箱地址, 示例值: "test_public_mailbox@xxx.xx"
	Name            *string `json:"name,omitempty"`             // 公共邮箱名称, 示例值: "test public mailbox"
}

// UpdatePublicMailboxResp ...
type UpdatePublicMailboxResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}

// updatePublicMailboxResp ...
type updatePublicMailboxResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdatePublicMailboxResp `json:"data,omitempty"`
}
