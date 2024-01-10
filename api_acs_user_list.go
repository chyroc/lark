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

// GetACSUserList 使用该接口获取智能门禁中所有用户信息。
//
// 只能获取已加入智能门禁权限组的用户。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/user/list
// new doc: https://open.feishu.cn/document/server-docs/acs-v1/user/list
func (r *ACSService) GetACSUserList(ctx context.Context, request *GetACSUserListReq, options ...MethodOptionFunc) (*GetACSUserListResp, *Response, error) {
	if r.cli.mock.mockACSGetACSUserList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] ACS#GetACSUserList mock enable")
		return r.cli.mock.mockACSGetACSUserList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "ACS",
		API:                   "GetACSUserList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/acs/v1/users",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getACSUserListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSGetACSUserList mock ACSGetACSUserList method
func (r *Mock) MockACSGetACSUserList(f func(ctx context.Context, request *GetACSUserListReq, options ...MethodOptionFunc) (*GetACSUserListResp, *Response, error)) {
	r.mockACSGetACSUserList = f
}

// UnMockACSGetACSUserList un-mock ACSGetACSUserList method
func (r *Mock) UnMockACSGetACSUserList() {
	r.mockACSGetACSUserList = nil
}

// GetACSUserListReq ...
type GetACSUserListReq struct {
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `50`
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "10"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetACSUserListResp ...
type GetACSUserListResp struct {
	Items     []*GetACSUserListRespItem `json:"items,omitempty"`
	PageToken string                    `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                      `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetACSUserListRespItem ...
type GetACSUserListRespItem struct {
	Feature *GetACSUserListRespItemFeature `json:"feature,omitempty"` // 用户特征
	UserID  string                         `json:"user_id,omitempty"` // 用户 ID
}

// GetACSUserListRespItemFeature ...
type GetACSUserListRespItemFeature struct {
	Card         int64 `json:"card,omitempty"`          // 卡号
	FaceUploaded bool  `json:"face_uploaded,omitempty"` // 是否已上传人脸图片
}

// getACSUserListResp ...
type getACSUserListResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetACSUserListResp `json:"data,omitempty"`
}
