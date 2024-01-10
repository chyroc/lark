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

// CreateSearchDataSource 创建一个数据源。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/create
// new doc: https://open.feishu.cn/document/server-docs/search-v2/open-search/data_source/create
func (r *SearchService) CreateSearchDataSource(ctx context.Context, request *CreateSearchDataSourceReq, options ...MethodOptionFunc) (*CreateSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchCreateSearchDataSource != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Search#CreateSearchDataSource mock enable")
		return r.cli.mock.mockSearchCreateSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "CreateSearchDataSource",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchCreateSearchDataSource mock SearchCreateSearchDataSource method
func (r *Mock) MockSearchCreateSearchDataSource(f func(ctx context.Context, request *CreateSearchDataSourceReq, options ...MethodOptionFunc) (*CreateSearchDataSourceResp, *Response, error)) {
	r.mockSearchCreateSearchDataSource = f
}

// UnMockSearchCreateSearchDataSource un-mock SearchCreateSearchDataSource method
func (r *Mock) UnMockSearchCreateSearchDataSource() {
	r.mockSearchCreateSearchDataSource = nil
}

// CreateSearchDataSourceReq ...
type CreateSearchDataSourceReq struct {
	Name             string                                    `json:"name,omitempty"`              // data_source的展示名称, 示例值: "客服工单"
	State            *int64                                    `json:"state,omitempty"`             // 数据源状态, 0-已上线, 1-未上线。如果未填, 默认是未上线状态, 示例值: 0, 可选值有: 0: 已上线, 1: 未上线
	Description      *string                                   `json:"description,omitempty"`       // 对于数据源的描述, 示例值: "搜索客服工单数据"
	IconURL          *string                                   `json:"icon_url,omitempty"`          // 数据源在 search tab 上的展示图标路径, 建议使用png或jpeg格式, 否则可能无法在客户端正常展示, 示例值: "https://www.xxx.com/open.jpg"
	Template         *string                                   `json:"template,omitempty"`          // 数据源采用的展示模版名称, 示例值: "search_common_card", 默认值: `search_common_card`
	SearchableFields []string                                  `json:"searchable_fields,omitempty"` // 【已废弃, 如有定制需要请使用“数据范式”接口】描述哪些字段可以被搜索, 示例值: 【已废弃, 如有定制需要请使用“数据范式”接口】["field1", "field2"]
	I18nName         *CreateSearchDataSourceReqI18nName        `json:"i18n_name,omitempty"`         // 数据源的国际化展示名称
	I18nDescription  *CreateSearchDataSourceReqI18nDescription `json:"i18n_description,omitempty"`  // 数据源的国际化描述
	SchemaID         *string                                   `json:"schema_id,omitempty"`         // 数据源关联的 schema 标识, 示例值: "custom_schema"
}

// CreateSearchDataSourceReqI18nDescription ...
type CreateSearchDataSourceReqI18nDescription struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 国际化字段: 中文, 示例值: "任务"
	EnUs *string `json:"en_us,omitempty"` // 国际化字段: 英文, 示例值: "TODO"
	JaJp *string `json:"ja_jp,omitempty"` // 国际化字段: 日文, 示例值: "タスク"
}

// CreateSearchDataSourceReqI18nName ...
type CreateSearchDataSourceReqI18nName struct {
	ZhCn *string `json:"zh_cn,omitempty"` // 国际化字段: 中文, 示例值: "任务"
	EnUs *string `json:"en_us,omitempty"` // 国际化字段: 英文, 示例值: "TODO"
	JaJp *string `json:"ja_jp,omitempty"` // 国际化字段: 日文, 示例值: "タスク"
}

// CreateSearchDataSourceResp ...
type CreateSearchDataSourceResp struct {
	DataSource *CreateSearchDataSourceRespDataSource `json:"data_source,omitempty"` // 数据源实例
}

// CreateSearchDataSourceRespDataSource ...
type CreateSearchDataSourceRespDataSource struct {
	ID               string                                               `json:"id,omitempty"`                // 数据源的唯一标识
	Name             string                                               `json:"name,omitempty"`              // data_source的展示名称
	State            int64                                                `json:"state,omitempty"`             // 数据源状态, 0-已上线, 1-未上线。如果未填, 默认是未上线状态, 可选值有: 0: 已上线, 1: 未上线
	Description      string                                               `json:"description,omitempty"`       // 对于数据源的描述
	CreateTime       string                                               `json:"create_time,omitempty"`       // 创建时间, 使用Unix时间戳, 单位为“秒”
	UpdateTime       string                                               `json:"update_time,omitempty"`       // 更新时间, 使用Unix时间戳, 单位为“秒”
	IsExceedQuota    bool                                                 `json:"is_exceed_quota,omitempty"`   // 是否超限
	IconURL          string                                               `json:"icon_url,omitempty"`          // 数据源在 search tab 上的展示图标路径, 建议使用png或jpeg格式, 否则可能无法在客户端正常展示
	Template         string                                               `json:"template,omitempty"`          // 数据源采用的展示模版名称
	SearchableFields []string                                             `json:"searchable_fields,omitempty"` // 【已废弃, 如有定制需要请使用“数据范式”接口】描述哪些字段可以被搜索
	I18nName         *CreateSearchDataSourceRespDataSourceI18nName        `json:"i18n_name,omitempty"`         // 数据源的国际化展示名称
	I18nDescription  *CreateSearchDataSourceRespDataSourceI18nDescription `json:"i18n_description,omitempty"`  // 数据源的国际化描述
	SchemaID         string                                               `json:"schema_id,omitempty"`         // 数据源关联的 schema 标识
}

// CreateSearchDataSourceRespDataSourceI18nDescription ...
type CreateSearchDataSourceRespDataSourceI18nDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 国际化字段: 中文
	EnUs string `json:"en_us,omitempty"` // 国际化字段: 英文
	JaJp string `json:"ja_jp,omitempty"` // 国际化字段: 日文
}

// CreateSearchDataSourceRespDataSourceI18nName ...
type CreateSearchDataSourceRespDataSourceI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 国际化字段: 中文
	EnUs string `json:"en_us,omitempty"` // 国际化字段: 英文
	JaJp string `json:"ja_jp,omitempty"` // 国际化字段: 日文
}

// createSearchDataSourceResp ...
type createSearchDataSourceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateSearchDataSourceResp `json:"data,omitempty"`
}
