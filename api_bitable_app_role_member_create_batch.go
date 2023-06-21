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

// BatchCreateBitableAppRoleMember 批量新增自定义角色的协作者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/batch_create
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/advanced-permission/app-role-member/batch_create
func (r *BitableService) BatchCreateBitableAppRoleMember(ctx context.Context, request *BatchCreateBitableAppRoleMemberReq, options ...MethodOptionFunc) (*BatchCreateBitableAppRoleMemberResp, *Response, error) {
	if r.cli.mock.mockBitableBatchCreateBitableAppRoleMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#BatchCreateBitableAppRoleMember mock enable")
		return r.cli.mock.mockBitableBatchCreateBitableAppRoleMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "BatchCreateBitableAppRoleMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/batch_create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchCreateBitableAppRoleMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableBatchCreateBitableAppRoleMember mock BitableBatchCreateBitableAppRoleMember method
func (r *Mock) MockBitableBatchCreateBitableAppRoleMember(f func(ctx context.Context, request *BatchCreateBitableAppRoleMemberReq, options ...MethodOptionFunc) (*BatchCreateBitableAppRoleMemberResp, *Response, error)) {
	r.mockBitableBatchCreateBitableAppRoleMember = f
}

// UnMockBitableBatchCreateBitableAppRoleMember un-mock BitableBatchCreateBitableAppRoleMember method
func (r *Mock) UnMockBitableBatchCreateBitableAppRoleMember() {
	r.mockBitableBatchCreateBitableAppRoleMember = nil
}

// BatchCreateBitableAppRoleMemberReq ...
type BatchCreateBitableAppRoleMemberReq struct {
	AppToken   string                                      `path:"app_token" json:"-"`    // 多维表格的唯一标识符 [app_token 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe), 示例值: "bascnnKKvcoUblgmmhZkYqabcef"
	RoleID     string                                      `path:"role_id" json:"-"`      // 自定义角色 ID, 示例值: "rolNGhPqks"
	MemberList []*BatchCreateBitableAppRoleMemberReqMember `json:"member_list,omitempty"` // 协作者列表, 最大长度: `100`
}

// BatchCreateBitableAppRoleMemberReqMember ...
type BatchCreateBitableAppRoleMemberReqMember struct {
	Type *string `json:"type,omitempty"` // 协作者 ID 类型, 示例值: "open_id", 可选值有: open_id: 协作者 ID 类型为 open_id, union_id: 协作者 ID 类型为 union_id, user_id: 协作者 ID 类型为 user_id, chat_id: 协作者 ID 类型为 chat_id, department_id: 协作者 ID 类型为 department_id, open_department_id: 协作者 ID 类型为 open_department_id, 默认值: `open_id`
	ID   string  `json:"id,omitempty"`   // 协作者 ID, 示例值: "ou_35990a9d9052051a2fae9b2f1afabcef"
}

// BatchCreateBitableAppRoleMemberResp ...
type BatchCreateBitableAppRoleMemberResp struct {
}

// batchCreateBitableAppRoleMemberResp ...
type batchCreateBitableAppRoleMemberResp struct {
	Code int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *BatchCreateBitableAppRoleMemberResp `json:"data,omitempty"`
}
