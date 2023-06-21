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

// GetCoreHrCurrencyList 批量查询货币信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/currency/list
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/currency/list
func (r *CoreHrService) GetCoreHrCurrencyList(ctx context.Context, request *GetCoreHrCurrencyListReq, options ...MethodOptionFunc) (*GetCoreHrCurrencyListResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrCurrencyList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrCurrencyList mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrCurrencyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrCurrencyList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/currencies",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrCurrencyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrCurrencyList mock CoreHrGetCoreHrCurrencyList method
func (r *Mock) MockCoreHrGetCoreHrCurrencyList(f func(ctx context.Context, request *GetCoreHrCurrencyListReq, options ...MethodOptionFunc) (*GetCoreHrCurrencyListResp, *Response, error)) {
	r.mockCoreHrGetCoreHrCurrencyList = f
}

// UnMockCoreHrGetCoreHrCurrencyList un-mock CoreHrGetCoreHrCurrencyList method
func (r *Mock) UnMockCoreHrGetCoreHrCurrencyList() {
	r.mockCoreHrGetCoreHrCurrencyList = nil
}

// GetCoreHrCurrencyListReq ...
type GetCoreHrCurrencyListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100
}

// GetCoreHrCurrencyListResp ...
type GetCoreHrCurrencyListResp struct {
	Items     []*GetCoreHrCurrencyListRespItem `json:"items,omitempty"`      // 货币信息
	HasMore   bool                             `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHrCurrencyListRespItem ...
type GetCoreHrCurrencyListRespItem struct {
	ID                 string                                       `json:"id,omitempty"`                    // 货币id
	CountryRegionID    string                                       `json:"country_region_id,omitempty"`     // 货币所属国家/地区id, 详细信息可通过【查询国家/地区信息】接口查询获得
	CurrencyName       []*GetCoreHrCurrencyListRespItemCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                        `json:"numeric_code,omitempty"`          // 数字代码
	CurrencyAlpha3Code string                                       `json:"currency_alpha_3_code,omitempty"` // 三位字母代码
}

// GetCoreHrCurrencyListRespItemCurrencyName ...
type GetCoreHrCurrencyListRespItemCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHrCurrencyListResp ...
type getCoreHrCurrencyListResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrCurrencyListResp `json:"data,omitempty"`
}
