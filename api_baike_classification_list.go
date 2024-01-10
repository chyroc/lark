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

// GetBaikeClassificationList 获取飞书词典当前分类。
//
// 飞书词典目前为二级分类体系, 每个词条可添加多个二级分类, 但选择的二级分类必须从属于不同的一级分类。
// 为了更好地提升接口文档的的易理解性, 我们对文档进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/classification/list)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/classification/list
// new doc: https://open.feishu.cn/document/server-docs/baike-v1/classification/list
//
// Deprecated
func (r *BaikeService) GetBaikeClassificationList(ctx context.Context, request *GetBaikeClassificationListReq, options ...MethodOptionFunc) (*GetBaikeClassificationListResp, *Response, error) {
	if r.cli.mock.mockBaikeGetBaikeClassificationList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Baike#GetBaikeClassificationList mock enable")
		return r.cli.mock.mockBaikeGetBaikeClassificationList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "GetBaikeClassificationList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/classifications",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBaikeClassificationListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeGetBaikeClassificationList mock BaikeGetBaikeClassificationList method
func (r *Mock) MockBaikeGetBaikeClassificationList(f func(ctx context.Context, request *GetBaikeClassificationListReq, options ...MethodOptionFunc) (*GetBaikeClassificationListResp, *Response, error)) {
	r.mockBaikeGetBaikeClassificationList = f
}

// UnMockBaikeGetBaikeClassificationList un-mock BaikeGetBaikeClassificationList method
func (r *Mock) UnMockBaikeGetBaikeClassificationList() {
	r.mockBaikeGetBaikeClassificationList = nil
}

// GetBaikeClassificationListReq ...
type GetBaikeClassificationListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 20, 默认值: `20`, 取值范围: `1` ～ `500`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 408ecac018b2e3518db37275e812bb8ad3e755fc886f322ac6c430ba
}

// GetBaikeClassificationListResp ...
type GetBaikeClassificationListResp struct {
	Items     []*GetBaikeClassificationListRespItem `json:"items,omitempty"`      // 分类
	PageToken string                                `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetBaikeClassificationListRespItem ...
type GetBaikeClassificationListRespItem struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	Name     string `json:"name,omitempty"`      // 二级分类名称
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// getBaikeClassificationListResp ...
type getBaikeClassificationListResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetBaikeClassificationListResp `json:"data,omitempty"`
}
