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

// SearchCoreHRNationality 根据国家 ID、国籍 ID 查询国籍信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-nationality/search
func (r *CoreHRService) SearchCoreHRNationality(ctx context.Context, request *SearchCoreHRNationalityReq, options ...MethodOptionFunc) (*SearchCoreHRNationalityResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRNationality != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRNationality mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRNationality(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRNationality",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/basic_info/nationalities/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRNationalityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRNationality mock CoreHRSearchCoreHRNationality method
func (r *Mock) MockCoreHRSearchCoreHRNationality(f func(ctx context.Context, request *SearchCoreHRNationalityReq, options ...MethodOptionFunc) (*SearchCoreHRNationalityResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRNationality = f
}

// UnMockCoreHRSearchCoreHRNationality un-mock CoreHRSearchCoreHRNationality method
func (r *Mock) UnMockCoreHRSearchCoreHRNationality() {
	r.mockCoreHRSearchCoreHRNationality = nil
}

// SearchCoreHRNationalityReq ...
type SearchCoreHRNationalityReq struct {
	PageSize            int64    `query:"page_size" json:"-"`              // 分页大小, 最大 100, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken           *string  `query:"page_token" json:"-"`             // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	NationalityIDList   []string `json:"nationality_id_list,omitempty"`    // 国籍 ID 列表, 示例值: ["6891251722631891445"], 最大长度: `100`
	CountryRegionIDList []string `json:"country_region_id_list,omitempty"` // 国家 / 地区 ID 列表, 可通过[查询国家 / 地区信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)接口查询, 示例值: ["6891251722631891441"], 最大长度: `100`
	StatusList          []int64  `json:"status_list,omitempty"`            // 状态列表, 示例值: [1], 可选值有: 1: 生效, 0: 失效, 默认值: `[1]`, 最大长度: `2`
}

// SearchCoreHRNationalityResp ...
type SearchCoreHRNationalityResp struct {
	Items     []*SearchCoreHRNationalityRespItem `json:"items,omitempty"`      // 查询的国籍信息
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchCoreHRNationalityRespItem ...
type SearchCoreHRNationalityRespItem struct {
	NationalityID   string                                 `json:"nationality_id,omitempty"`    // 国籍 ID（对应其他查询结果的 nationality_id_v2 字段）
	Name            []*SearchCoreHRNationalityRespItemName `json:"name,omitempty"`              // 名称
	Alpha2Code      string                                 `json:"alpha_2_code,omitempty"`      // 二字码
	Alpha3Code      string                                 `json:"alpha_3_code,omitempty"`      // 三字码
	NumericCode     int64                                  `json:"numeric_code,omitempty"`      // 数字代码
	CountryRegionID string                                 `json:"country_region_id,omitempty"` // 国家 / 地区 ID, 可通过[查询国家 / 地区信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-country_region/search)接口查询
	Status          int64                                  `json:"status,omitempty"`            // 状态, 可选值有: 1: 生效, 0: 失效
}

// SearchCoreHRNationalityRespItemName ...
type SearchCoreHRNationalityRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// searchCoreHRNationalityResp ...
type searchCoreHRNationalityResp struct {
	Code  int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                       `json:"msg,omitempty"`  // 错误描述
	Data  *SearchCoreHRNationalityResp `json:"data,omitempty"`
	Error *ErrorDetail                 `json:"error,omitempty"`
}
