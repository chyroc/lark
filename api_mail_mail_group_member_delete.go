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

// DeleteMailGroupMember 删除邮件组单个成员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-member/delete
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/mail-group/mailgroup-member/delete
func (r *MailService) DeleteMailGroupMember(ctx context.Context, request *DeleteMailGroupMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupMemberResp, *Response, error) {
	if r.cli.mock.mockMailDeleteMailGroupMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailGroupMember mock enable")
		return r.cli.mock.mockMailDeleteMailGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "DeleteMailGroupMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteMailGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailDeleteMailGroupMember mock MailDeleteMailGroupMember method
func (r *Mock) MockMailDeleteMailGroupMember(f func(ctx context.Context, request *DeleteMailGroupMemberReq, options ...MethodOptionFunc) (*DeleteMailGroupMemberResp, *Response, error)) {
	r.mockMailDeleteMailGroupMember = f
}

// UnMockMailDeleteMailGroupMember un-mock MailDeleteMailGroupMember method
func (r *Mock) UnMockMailDeleteMailGroupMember() {
	r.mockMailDeleteMailGroupMember = nil
}

// DeleteMailGroupMemberReq ...
type DeleteMailGroupMemberReq struct {
	MailGroupID string `path:"mailgroup_id" json:"-"` // The unique ID or email address of a mail group, 示例值: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx"
	MemberID    string `path:"member_id" json:"-"`    // The unique ID of a member in this mail group, 示例值: "xxxxxxxxxxxxxxx"
}

// DeleteMailGroupMemberResp ...
type DeleteMailGroupMemberResp struct {
}

// deleteMailGroupMemberResp ...
type deleteMailGroupMemberResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailGroupMemberResp `json:"data,omitempty"`
}
