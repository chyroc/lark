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

// BatchDeleteTaskV1Follower 该接口用于批量删除关注人。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/batch_delete_follower
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task-follower/batch_delete_follower
//
// Deprecated
func (r *TaskV1Service) BatchDeleteTaskV1Follower(ctx context.Context, request *BatchDeleteTaskV1FollowerReq, options ...MethodOptionFunc) (*BatchDeleteTaskV1FollowerResp, *Response, error) {
	if r.cli.mock.mockTaskV1BatchDeleteTaskV1Follower != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] TaskV1#BatchDeleteTaskV1Follower mock enable")
		return r.cli.mock.mockTaskV1BatchDeleteTaskV1Follower(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "TaskV1",
		API:                   "BatchDeleteTaskV1Follower",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/batch_delete_follower",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchDeleteTaskV1FollowerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskV1BatchDeleteTaskV1Follower mock TaskV1BatchDeleteTaskV1Follower method
func (r *Mock) MockTaskV1BatchDeleteTaskV1Follower(f func(ctx context.Context, request *BatchDeleteTaskV1FollowerReq, options ...MethodOptionFunc) (*BatchDeleteTaskV1FollowerResp, *Response, error)) {
	r.mockTaskV1BatchDeleteTaskV1Follower = f
}

// UnMockTaskV1BatchDeleteTaskV1Follower un-mock TaskV1BatchDeleteTaskV1Follower method
func (r *Mock) UnMockTaskV1BatchDeleteTaskV1Follower() {
	r.mockTaskV1BatchDeleteTaskV1Follower = nil
}

// BatchDeleteTaskV1FollowerReq ...
type BatchDeleteTaskV1FollowerReq struct {
	TaskID     string   `path:"task_id" json:"-"`       // 任务ID, 示例值: "83912691-2e43-47fc-94a4-d512e03984fa"
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	IDList     []string `json:"id_list,omitempty"`      // 要删除的关注人ID列表, 示例值: ["ou_99e1a581b36ecc4862cbfbce473f3123"]
}

// BatchDeleteTaskV1FollowerResp ...
type BatchDeleteTaskV1FollowerResp struct {
	Followers []string `json:"followers,omitempty"` // 实际删除的关注人用户ID列表
}

// batchDeleteTaskV1FollowerResp ...
type batchDeleteTaskV1FollowerResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *BatchDeleteTaskV1FollowerResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}