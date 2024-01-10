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

// TransferApprovalInstance 对于单个审批任务进行转交操作。转交后审批流程流转给被转交人。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/task/transfer
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/task/transfer
func (r *ApprovalService) TransferApprovalInstance(ctx context.Context, request *TransferApprovalInstanceReq, options ...MethodOptionFunc) (*TransferApprovalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalTransferApprovalInstance != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Approval#TransferApprovalInstance mock enable")
		return r.cli.mock.mockApprovalTransferApprovalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "TransferApprovalInstance",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/tasks/transfer",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(transferApprovalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalTransferApprovalInstance mock ApprovalTransferApprovalInstance method
func (r *Mock) MockApprovalTransferApprovalInstance(f func(ctx context.Context, request *TransferApprovalInstanceReq, options ...MethodOptionFunc) (*TransferApprovalInstanceResp, *Response, error)) {
	r.mockApprovalTransferApprovalInstance = f
}

// UnMockApprovalTransferApprovalInstance un-mock ApprovalTransferApprovalInstance method
func (r *Mock) UnMockApprovalTransferApprovalInstance() {
	r.mockApprovalTransferApprovalInstance = nil
}

// TransferApprovalInstanceReq ...
type TransferApprovalInstanceReq struct {
	UserIDType     *IDType `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ApprovalCode   string  `json:"approval_code,omitempty"`    // 审批定义 Code, 示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	InstanceCode   string  `json:"instance_code,omitempty"`    // 审批实例 Code, 示例值: "81D31358-93AF-92D6-7425-01A5D67C4E71"
	UserID         string  `json:"user_id,omitempty"`          // 根据user_id_type填写操作用户id, 示例值: "f7cb567e"
	Comment        *string `json:"comment,omitempty"`          // 意见, 示例值: "OK"
	TransferUserID string  `json:"transfer_user_id,omitempty"` // 根据user_id_type填写被转交人唯一 ID, 示例值: "f4ip317q"
	TaskID         string  `json:"task_id,omitempty"`          // 任务 ID, 审批实例详情task_list中id, 示例值: "12345"
}

// TransferApprovalInstanceResp ...
type TransferApprovalInstanceResp struct {
}

// transferApprovalInstanceResp ...
type transferApprovalInstanceResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *TransferApprovalInstanceResp `json:"data,omitempty"`
}
