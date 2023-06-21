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

// BatchDeleteMailGroupPermissionMember 一次请求可以删除一个邮件组中的多个权限成员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup-permission_member/batch_delete
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/mail-group/mailgroup-permission_member/batch_delete
func (r *MailService) BatchDeleteMailGroupPermissionMember(ctx context.Context, request *BatchDeleteMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*BatchDeleteMailGroupPermissionMemberResp, *Response, error) {
	if r.cli.mock.mockMailBatchDeleteMailGroupPermissionMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#BatchDeleteMailGroupPermissionMember mock enable")
		return r.cli.mock.mockMailBatchDeleteMailGroupPermissionMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "BatchDeleteMailGroupPermissionMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id/permission_members/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchDeleteMailGroupPermissionMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailBatchDeleteMailGroupPermissionMember mock MailBatchDeleteMailGroupPermissionMember method
func (r *Mock) MockMailBatchDeleteMailGroupPermissionMember(f func(ctx context.Context, request *BatchDeleteMailGroupPermissionMemberReq, options ...MethodOptionFunc) (*BatchDeleteMailGroupPermissionMemberResp, *Response, error)) {
	r.mockMailBatchDeleteMailGroupPermissionMember = f
}

// UnMockMailBatchDeleteMailGroupPermissionMember un-mock MailBatchDeleteMailGroupPermissionMember method
func (r *Mock) UnMockMailBatchDeleteMailGroupPermissionMember() {
	r.mockMailBatchDeleteMailGroupPermissionMember = nil
}

// BatchDeleteMailGroupPermissionMemberReq ...
type BatchDeleteMailGroupPermissionMemberReq struct {
	MailGroupID            string   `path:"mailgroup_id" json:"-"`               // The unique ID or email address of a mail group, 示例值: "xxxxxxxxxxxxxxx or test_mail_group@xxx.xx"
	PermissionMemberIDList []string `json:"permission_member_id_list,omitempty"` // 本次调用删除的权限成员ID列表, 示例值: ["xxxxxxx", "yyyyyyy"], 长度范围: `1` ～ `200`
}

// BatchDeleteMailGroupPermissionMemberResp ...
type BatchDeleteMailGroupPermissionMemberResp struct {
}

// batchDeleteMailGroupPermissionMemberResp ...
type batchDeleteMailGroupPermissionMemberResp struct {
	Code int64                                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                    `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeleteMailGroupPermissionMemberResp `json:"data,omitempty"`
}
