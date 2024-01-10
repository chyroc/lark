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

// GetCoreHRLocationList 批量查询地点。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/location/list
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/location/list
func (r *CoreHRService) GetCoreHRLocationList(ctx context.Context, request *GetCoreHRLocationListReq, options ...MethodOptionFunc) (*GetCoreHRLocationListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRLocationList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRLocationList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRLocationList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRLocationList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/locations",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRLocationListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRLocationList mock CoreHRGetCoreHRLocationList method
func (r *Mock) MockCoreHRGetCoreHRLocationList(f func(ctx context.Context, request *GetCoreHRLocationListReq, options ...MethodOptionFunc) (*GetCoreHRLocationListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRLocationList = f
}

// UnMockCoreHRGetCoreHRLocationList un-mock CoreHRGetCoreHRLocationList method
func (r *Mock) UnMockCoreHRGetCoreHRLocationList() {
	r.mockCoreHRGetCoreHRLocationList = nil
}

// GetCoreHRLocationListReq ...
type GetCoreHRLocationListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1231231987
	PageSize  int64   `query:"page_size" json:"-"`  // 分页大小, 示例值: 100
}

// GetCoreHRLocationListResp ...
type GetCoreHRLocationListResp struct {
	Items     []*GetCoreHRLocationListRespItem `json:"items,omitempty"`      // 查询的地点信息
	HasMore   bool                             `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHRLocationListRespItem ...
type GetCoreHRLocationListRespItem struct {
	ID                 string                                         `json:"id,omitempty"`                    // 实体在CoreHR内部的唯一键
	HiberarchyCommon   *GetCoreHRLocationListRespItemHiberarchyCommon `json:"hiberarchy_common,omitempty"`     // 层级关系, 内层字段见实体
	LocationUsageList  []*GetCoreHRLocationListRespItemLocationUsage  `json:"location_usage_list,omitempty"`   // 地点用途
	Address            []*GetCoreHRLocationListRespItemAddres         `json:"address,omitempty"`               // 地址
	WorkingHoursTypeID string                                         `json:"working_hours_type_id,omitempty"` // 工时制度
	EffectiveTime      string                                         `json:"effective_time,omitempty"`        // 生效时间
	ExpirationTime     string                                         `json:"expiration_time,omitempty"`       // 失效时间
	CustomFields       []*GetCoreHRLocationListRespItemCustomField    `json:"custom_fields,omitempty"`         // 自定义字段
	Locale             *GetCoreHRLocationListRespItemLocale           `json:"locale,omitempty"`                // 区域设置
	TimeZoneID         string                                         `json:"time_zone_id,omitempty"`          // 时区
	DisplayLanguageID  string                                         `json:"display_language_id,omitempty"`   // 默认显示语言
}

// GetCoreHRLocationListRespItemAddres ...
type GetCoreHRLocationListRespItemAddres struct {
	FullAddressLocalScript   string                                            `json:"full_address_local_script,omitempty"`   // 完整地址（本地文字）
	FullAddressWesternScript string                                            `json:"full_address_western_script,omitempty"` // 完整地址（西方文字）
	ID                       string                                            `json:"id,omitempty"`                          // 地址ID
	CountryRegionID          string                                            `json:"country_region_id,omitempty"`           // 国家 / 地区
	RegionID                 string                                            `json:"region_id,omitempty"`                   // 主要行政区
	CityID                   string                                            `json:"city_id,omitempty"`                     // 城市, 该字段已作废, 请使用 city_id_v2 字段
	DistinctID               string                                            `json:"distinct_id,omitempty"`                 // 区/县, 该字段已作废, 请使用 district_id_v2 字段
	LocalAddressLine1        string                                            `json:"local_address_line1,omitempty"`         // 地址行 1（非拉丁语系的本地文字）
	LocalAddressLine2        string                                            `json:"local_address_line2,omitempty"`         // 地址行 2（非拉丁语系的本地文字）
	LocalAddressLine3        string                                            `json:"local_address_line3,omitempty"`         // 地址行 3（非拉丁语系的本地文字）
	LocalAddressLine4        string                                            `json:"local_address_line4,omitempty"`         // 地址行 4（非拉丁语系的本地文字）
	LocalAddressLine5        string                                            `json:"local_address_line5,omitempty"`         // 地址行 5（非拉丁语系的本地文字）
	LocalAddressLine6        string                                            `json:"local_address_line6,omitempty"`         // 地址行 6（非拉丁语系的本地文字）
	LocalAddressLine7        string                                            `json:"local_address_line7,omitempty"`         // 地址行 7（非拉丁语系的本地文字）
	LocalAddressLine8        string                                            `json:"local_address_line8,omitempty"`         // 地址行 8（非拉丁语系的本地文字）
	LocalAddressLine9        string                                            `json:"local_address_line9,omitempty"`         // 地址行 9（非拉丁语系的本地文字）
	PostalCode               string                                            `json:"postal_code,omitempty"`                 // 邮政编码
	AddressTypeList          []*GetCoreHRLocationListRespItemAddresAddressType `json:"address_type_list,omitempty"`           // 地址类型
	IsPrimary                bool                                              `json:"is_primary,omitempty"`                  // 主要地址
	IsPublic                 bool                                              `json:"is_public,omitempty"`                   // 公开地址
	CustomFields             []*GetCoreHRLocationListRespItemAddresCustomField `json:"custom_fields,omitempty"`               // 自定义字段
}

// GetCoreHRLocationListRespItemAddresAddressType ...
type GetCoreHRLocationListRespItemAddresAddressType struct {
	EnumName string                                                   `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRLocationListRespItemAddresAddressTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRLocationListRespItemAddresAddressTypeDisplay ...
type GetCoreHRLocationListRespItemAddresAddressTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRLocationListRespItemAddresCustomField ...
type GetCoreHRLocationListRespItemAddresCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// GetCoreHRLocationListRespItemCustomField ...
type GetCoreHRLocationListRespItemCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// GetCoreHRLocationListRespItemHiberarchyCommon ...
type GetCoreHRLocationListRespItemHiberarchyCommon struct {
	ParentID       string                                                      `json:"parent_id,omitempty"`       // 上级组织
	Name           []*GetCoreHRLocationListRespItemHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *GetCoreHRLocationListRespItemHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型
	Active         bool                                                        `json:"active,omitempty"`          // 启用
	EffectiveTime  string                                                      `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                      `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                      `json:"code,omitempty"`            // 编码
	Description    []*GetCoreHRLocationListRespItemHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                      `json:"tree_order,omitempty"`      // 树形排序, 代表同层级的部门排序序号
	ListOrder      string                                                      `json:"list_order,omitempty"`      // 列表排序, 代表所有部门的混排序号
	CustomFields   []*GetCoreHRLocationListRespItemHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// GetCoreHRLocationListRespItemHiberarchyCommonCustomField ...
type GetCoreHRLocationListRespItemHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// GetCoreHRLocationListRespItemHiberarchyCommonDescription ...
type GetCoreHRLocationListRespItemHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRLocationListRespItemHiberarchyCommonName ...
type GetCoreHRLocationListRespItemHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRLocationListRespItemHiberarchyCommonType ...
type GetCoreHRLocationListRespItemHiberarchyCommonType struct {
	EnumName string                                                      `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRLocationListRespItemHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRLocationListRespItemHiberarchyCommonTypeDisplay ...
type GetCoreHRLocationListRespItemHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRLocationListRespItemLocale ...
type GetCoreHRLocationListRespItemLocale struct {
	EnumName string                                        `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRLocationListRespItemLocaleDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRLocationListRespItemLocaleDisplay ...
type GetCoreHRLocationListRespItemLocaleDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// GetCoreHRLocationListRespItemLocationUsage ...
type GetCoreHRLocationListRespItemLocationUsage struct {
	EnumName string                                               `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRLocationListRespItemLocationUsageDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRLocationListRespItemLocationUsageDisplay ...
type GetCoreHRLocationListRespItemLocationUsageDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// getCoreHRLocationListResp ...
type getCoreHRLocationListResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRLocationListResp `json:"data,omitempty"`
}
