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

// CreateContactJobLevel 该接口可以创建职级。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/job_level/create
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/job_level/create
func (r *ContactService) CreateContactJobLevel(ctx context.Context, request *CreateContactJobLevelReq, options ...MethodOptionFunc) (*CreateContactJobLevelResp, *Response, error) {
	if r.cli.mock.mockContactCreateContactJobLevel != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#CreateContactJobLevel mock enable")
		return r.cli.mock.mockContactCreateContactJobLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "CreateContactJobLevel",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/job_levels",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createContactJobLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactCreateContactJobLevel mock ContactCreateContactJobLevel method
func (r *Mock) MockContactCreateContactJobLevel(f func(ctx context.Context, request *CreateContactJobLevelReq, options ...MethodOptionFunc) (*CreateContactJobLevelResp, *Response, error)) {
	r.mockContactCreateContactJobLevel = f
}

// UnMockContactCreateContactJobLevel un-mock ContactCreateContactJobLevel method
func (r *Mock) UnMockContactCreateContactJobLevel() {
	r.mockContactCreateContactJobLevel = nil
}

// CreateContactJobLevelReq ...
type CreateContactJobLevelReq struct {
	Name            string                                     `json:"name,omitempty"`             // 职级名称, 示例值: "高级专家", 长度范围: `1` ～ `255` 字符
	Description     *string                                    `json:"description,omitempty"`      // 职级描述, 示例值: "公司内部中高级职称, 有一定专业技术能力的人员"
	Order           *int64                                     `json:"order,omitempty"`            // 职级的排序, 可填入自然数100-100000的数值, 系统按照数值大小从小到大排序。不填写该字段时, 默认新增排序在当前职级列表中最后位（最大值）, 示例值: 200, 取值范围: `100` ～ `100000`
	Status          bool                                       `json:"status,omitempty"`           // 是否启用, 示例值: true
	I18nName        []*CreateContactJobLevelReqI18nName        `json:"i18n_name,omitempty"`        // 多语言名称
	I18nDescription []*CreateContactJobLevelReqI18nDescription `json:"i18n_description,omitempty"` // 多语言描述
}

// CreateContactJobLevelReqI18nDescription ...
type CreateContactJobLevelReqI18nDescription struct {
	Locale *string `json:"locale,omitempty"` // 语言版本, 示例值: "zh_cn"
	Value  *string `json:"value,omitempty"`  // 字段名, 示例值: "多语言内容"
}

// CreateContactJobLevelReqI18nName ...
type CreateContactJobLevelReqI18nName struct {
	Locale *string `json:"locale,omitempty"` // 语言版本, 示例值: "zh_cn"
	Value  *string `json:"value,omitempty"`  // 字段名, 示例值: "多语言内容"
}

// CreateContactJobLevelResp ...
type CreateContactJobLevelResp struct {
	JobLevel *CreateContactJobLevelRespJobLevel `json:"job_level,omitempty"` // 职级信息
}

// CreateContactJobLevelRespJobLevel ...
type CreateContactJobLevelRespJobLevel struct {
	Name            string                                              `json:"name,omitempty"`             // 职级名称
	Description     string                                              `json:"description,omitempty"`      // 职级描述
	Order           int64                                               `json:"order,omitempty"`            // 职级的排序, 可填入自然数100-100000的数值, 系统按照数值大小从小到大排序。不填写该字段时, 默认新增排序在当前职级列表中最后位（最大值）
	Status          bool                                                `json:"status,omitempty"`           // 是否启用
	JobLevelID      string                                              `json:"job_level_id,omitempty"`     // 职级ID
	I18nName        []*CreateContactJobLevelRespJobLevelI18nName        `json:"i18n_name,omitempty"`        // 多语言名称
	I18nDescription []*CreateContactJobLevelRespJobLevelI18nDescription `json:"i18n_description,omitempty"` // 多语言描述
}

// CreateContactJobLevelRespJobLevelI18nDescription ...
type CreateContactJobLevelRespJobLevelI18nDescription struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// CreateContactJobLevelRespJobLevelI18nName ...
type CreateContactJobLevelRespJobLevelI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// createContactJobLevelResp ...
type createContactJobLevelResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateContactJobLevelResp `json:"data,omitempty"`
}
