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

// GetBitableDashboardList 根据 app_token, 获取多维表格下的所有仪表盘
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-dashboard/list
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-dashboard/list
func (r *BitableService) GetBitableDashboardList(ctx context.Context, request *GetBitableDashboardListReq, options ...MethodOptionFunc) (*GetBitableDashboardListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableDashboardList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableDashboardList mock enable")
		return r.cli.mock.mockBitableGetBitableDashboardList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableDashboardList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/dashboards",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableDashboardListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableDashboardList mock BitableGetBitableDashboardList method
func (r *Mock) MockBitableGetBitableDashboardList(f func(ctx context.Context, request *GetBitableDashboardListReq, options ...MethodOptionFunc) (*GetBitableDashboardListResp, *Response, error)) {
	r.mockBitableGetBitableDashboardList = f
}

// UnMockBitableGetBitableDashboardList un-mock BitableGetBitableDashboardList method
func (r *Mock) UnMockBitableGetBitableDashboardList() {
	r.mockBitableGetBitableDashboardList = nil
}

// GetBitableDashboardListReq ...
type GetBitableDashboardListReq struct {
	AppToken  string  `path:"app_token" json:"-"`   // 多维表格文档 Token, 示例值: "bascng7vrxcxpig7geggXiCtadY"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `500`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "blknkqrP3RqUkcAW"
}

// GetBitableDashboardListResp ...
type GetBitableDashboardListResp struct {
	Dashboards []*GetBitableDashboardListRespDashboard `json:"dashboards,omitempty"` // 仪表盘信息
	PageToken  string                                  `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore    bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetBitableDashboardListRespDashboard ...
type GetBitableDashboardListRespDashboard struct {
	BlockID string `json:"block_id,omitempty"` // 仪表盘 ID
	Name    string `json:"name,omitempty"`     // 仪表盘名字
}

// getBitableDashboardListResp ...
type getBitableDashboardListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableDashboardListResp `json:"data,omitempty"`
}
