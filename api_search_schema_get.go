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

// GetSearchSchema 获取单个数据范式
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/schema/get
func (r *SearchService) GetSearchSchema(ctx context.Context, request *GetSearchSchemaReq, options ...MethodOptionFunc) (*GetSearchSchemaResp, *Response, error) {
	if r.cli.mock.mockSearchGetSearchSchema != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#GetSearchSchema mock enable")
		return r.cli.mock.mockSearchGetSearchSchema(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "GetSearchSchema",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/schemas/:schema_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getSearchSchemaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchGetSearchSchema mock SearchGetSearchSchema method
func (r *Mock) MockSearchGetSearchSchema(f func(ctx context.Context, request *GetSearchSchemaReq, options ...MethodOptionFunc) (*GetSearchSchemaResp, *Response, error)) {
	r.mockSearchGetSearchSchema = f
}

// UnMockSearchGetSearchSchema un-mock SearchGetSearchSchema method
func (r *Mock) UnMockSearchGetSearchSchema() {
	r.mockSearchGetSearchSchema = nil
}

// GetSearchSchemaReq ...
type GetSearchSchemaReq struct {
	SchemaID string `path:"schema_id" json:"-"` // 用户自定义数据范式的唯一标识, 示例值: "custom_schema_id", 最大长度: `40` 字符, 正则校验: `^[a-zA-Z][a-zA-Z0-9-_].*$`
}

// GetSearchSchemaResp ...
type GetSearchSchemaResp struct {
	Schema *GetSearchSchemaRespSchema `json:"schema,omitempty"` // 数据范式
}

// GetSearchSchemaRespSchema ...
type GetSearchSchemaRespSchema struct {
	Properties []*GetSearchSchemaRespSchemaPropertie `json:"properties,omitempty"` // 数据范式的属性定义
	Display    *GetSearchSchemaRespSchemaDisplay     `json:"display,omitempty"`    // 数据展示相关配置
	SchemaID   string                                `json:"schema_id,omitempty"`  // 用户自定义数据范式的唯一标识
}

// GetSearchSchemaRespSchemaDisplay ...
type GetSearchSchemaRespSchemaDisplay struct {
	CardKey       string                                           `json:"card_key,omitempty"`       // 搜索数据的展示卡片, 可选值有: search_common_card: 普通 common 卡片
	FieldsMapping []*GetSearchSchemaRespSchemaDisplayFieldsMapping `json:"fields_mapping,omitempty"` // 数据字段名称和展示字段名称的映射关系。如果没有设置, 则只会展示 与展示字段名称同名的 数据字段
}

// GetSearchSchemaRespSchemaDisplayFieldsMapping ...
type GetSearchSchemaRespSchemaDisplayFieldsMapping struct {
	DisplayField string `json:"display_field,omitempty"` // 展示字段名称, 与 card_key 有关, 每个模版能展示的字段不同。该字段不能重复
	DataField    string `json:"data_field,omitempty"`    // 数据字段的名称。需要确保该字段对应在 schema 属性定义中的 is_returnable 为 true, 否则无法展示。需要使用 ${xxx} 的规则来描述
}

// GetSearchSchemaRespSchemaPropertie ...
type GetSearchSchemaRespSchemaPropertie struct {
	Name            string                                             `json:"name,omitempty"`             // 属性名
	Type            string                                             `json:"type,omitempty"`             // 属性类型, 可选值有: text: -长文本类型, 长度大于20的文本, int: 64位整数类型, tag: 标签类型, timestamp: Unix 时间戳类型（单位为秒）, double: 浮点数类型（小数）, tinytext: 短文本类型, 长度小于等于20的文本
	IsSearchable    bool                                               `json:"is_searchable,omitempty"`    // 该属性是否可用作搜索, 默认为 false
	IsSortable      bool                                               `json:"is_sortable,omitempty"`      // 该属性是否可用作搜索结果排序, 默认为 false。如果为 true, 需要再配置 sortOptions
	IsReturnable    bool                                               `json:"is_returnable,omitempty"`    // 该属性是否可用作返回字段, 为 false 时, 该字段不会被召回和展示。默认为 false
	SortOptions     *GetSearchSchemaRespSchemaPropertieSortOptions     `json:"sort_options,omitempty"`     // 属性排序的可选配置, 当 is_sortable 为 true 时, 该字段为必填字段
	TypeDefinitions *GetSearchSchemaRespSchemaPropertieTypeDefinitions `json:"type_definitions,omitempty"` // 相关类型数据的定义和约束
	SearchOptions   *GetSearchSchemaRespSchemaPropertieSearchOptions   `json:"search_options,omitempty"`   // 属性搜索的可选配置, 当 is_searchable 为 true 时, 该字段为必填参数
}

// GetSearchSchemaRespSchemaPropertieSearchOptions ...
type GetSearchSchemaRespSchemaPropertieSearchOptions struct {
	EnableSemanticMatch     bool `json:"enable_semantic_match,omitempty"`      // 是否支持语义切词召回。默认不支持（推荐使用在长文本的场景）
	EnableExactMatch        bool `json:"enable_exact_match,omitempty"`         // 是否支持精确匹配。默认不支持（推荐使用在短文本、需要精确查找的场景）
	EnablePrefixMatch       bool `json:"enable_prefix_match,omitempty"`        // 是否支持前缀匹配（短文本的默认的分词/召回策略。前缀长度为 1-12）
	EnableNumberSuffixMatch bool `json:"enable_number_suffix_match,omitempty"` // 是否支持数据后缀匹配。默认不支持（推荐使用在短文本、有数字后缀查找的场景。后缀长度为3-12）
	EnableCamelMatch        bool `json:"enable_camel_match,omitempty"`         // 是否支持驼峰英文匹配。默认不支持（推荐使用在短文本, 且包含驼峰形式英文的查找场景）
}

// GetSearchSchemaRespSchemaPropertieSortOptions ...
type GetSearchSchemaRespSchemaPropertieSortOptions struct {
	Priority int64  `json:"priority,omitempty"` // 排序的优先级, 可选范围为 0~4, 0为最高优先级。如果优先级相同, 则随机进行排序。默认为0, 可选值有: 0: 最高优先级, 1: 次高优先级, 2: 次次高优先级, 3: 次低优先级, 4: 最低优先级
	Order    string `json:"order,omitempty"`    // 排序的顺序。默认为 desc, 可选值有: asc: 升序, desc: 降序
}

// GetSearchSchemaRespSchemaPropertieTypeDefinitions ...
type GetSearchSchemaRespSchemaPropertieTypeDefinitions struct {
	Tag []*GetSearchSchemaRespSchemaPropertieTypeDefinitionsTag `json:"tag,omitempty"` // 标签类型的定义
}

// GetSearchSchemaRespSchemaPropertieTypeDefinitionsTag ...
type GetSearchSchemaRespSchemaPropertieTypeDefinitionsTag struct {
	Name  string `json:"name,omitempty"`  // tag 对应的枚举值名称
	Color string `json:"color,omitempty"` // 标签对应的颜色, 可选值有: red: 含警示性、敏感性的提示信息, green: 表示成功、完成、完毕的提示信息, blue: 组件架构、职能等中性信息, grey: 中立系统提示信息（慎重使用）, yellow: 焦点信息、推广性信息
	Text  string `json:"text,omitempty"`  // 标签中展示的文本
}

// getSearchSchemaResp ...
type getSearchSchemaResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetSearchSchemaResp `json:"data,omitempty"`
}