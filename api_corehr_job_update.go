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

// UpdateCoreHrJob 更新职务。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job/patch
func (r *CoreHrService) UpdateCoreHrJob(ctx context.Context, request *UpdateCoreHrJobReq, options ...MethodOptionFunc) (*UpdateCoreHrJobResp, *Response, error) {
	if r.cli.mock.mockCoreHrUpdateCoreHrJob != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#UpdateCoreHrJob mock enable")
		return r.cli.mock.mockCoreHrUpdateCoreHrJob(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "UpdateCoreHrJob",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/jobs/:job_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateCoreHrJobResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrUpdateCoreHrJob mock CoreHrUpdateCoreHrJob method
func (r *Mock) MockCoreHrUpdateCoreHrJob(f func(ctx context.Context, request *UpdateCoreHrJobReq, options ...MethodOptionFunc) (*UpdateCoreHrJobResp, *Response, error)) {
	r.mockCoreHrUpdateCoreHrJob = f
}

// UnMockCoreHrUpdateCoreHrJob un-mock CoreHrUpdateCoreHrJob method
func (r *Mock) UnMockCoreHrUpdateCoreHrJob() {
	r.mockCoreHrUpdateCoreHrJob = nil
}

// UpdateCoreHrJobReq ...
type UpdateCoreHrJobReq struct {
	JobID              string                           `path:"job_id" json:"-"`                 // 职务ID, 示例值: "1616161616"
	ClientToken        *string                          `query:"client_token" json:"-"`          // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	Code               *string                          `json:"code,omitempty"`                  // 编码, 示例值: "JP422119"
	Name               []*UpdateCoreHrJobReqName        `json:"name,omitempty"`                  // 名称
	Description        []*UpdateCoreHrJobReqDescription `json:"description,omitempty"`           // 描述
	Active             *bool                            `json:"active,omitempty"`                // 是否启用, 示例值: true
	JobTitle           []*UpdateCoreHrJobReqJobTitle    `json:"job_title,omitempty"`             // 职务头衔
	JobFamilyIDList    []string                         `json:"job_family_id_list,omitempty"`    // 职务序列 ID 列表, 枚举值及详细信息可通过【批量查询职务序列】接口查询获得, 示例值: ["7873827832"]
	JobLevelIDList     []string                         `json:"job_level_id_list,omitempty"`     // 职务级别 ID 列表, 枚举值及详细信息可通过【批量查询职务级别】接口查询获得, 示例值: ["4374837438"]
	WorkingHoursTypeID *string                          `json:"working_hours_type_id,omitempty"` // 工时制度 ID, 枚举值及详细信息可通过【批量查询工时制度】接口查询获得, 示例值: "6890452208593372679"
	EffectiveTime      *string                          `json:"effective_time,omitempty"`        // 生效时间, 示例值: "2020-01-01 00:00:00"
	ExpirationTime     *string                          `json:"expiration_time,omitempty"`       // 失效时间, 示例值: "2021-01-01 00:00:00"
	CustomFields       []*UpdateCoreHrJobReqCustomField `json:"custom_fields,omitempty"`         // 自定义字段
}

// UpdateCoreHrJobReqCustomField ...
type UpdateCoreHrJobReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// UpdateCoreHrJobReqDescription ...
type UpdateCoreHrJobReqDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHrJobReqJobTitle ...
type UpdateCoreHrJobReqJobTitle struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHrJobReqName ...
type UpdateCoreHrJobReqName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// UpdateCoreHrJobResp ...
type UpdateCoreHrJobResp struct {
	Job *UpdateCoreHrJobRespJob `json:"job,omitempty"` // 职务
}

// UpdateCoreHrJobRespJob ...
type UpdateCoreHrJobRespJob struct {
	ID                 string                               `json:"id,omitempty"`                    // 职务 ID
	Code               string                               `json:"code,omitempty"`                  // 编码
	Name               []*UpdateCoreHrJobRespJobName        `json:"name,omitempty"`                  // 名称
	Description        []*UpdateCoreHrJobRespJobDescription `json:"description,omitempty"`           // 描述
	Active             bool                                 `json:"active,omitempty"`                // 是否启用
	JobTitle           []*UpdateCoreHrJobRespJobJobTitle    `json:"job_title,omitempty"`             // 职务头衔
	JobFamilyIDList    []string                             `json:"job_family_id_list,omitempty"`    // 职务序列 ID 列表, 枚举值及详细信息可通过【批量查询职务序列】接口查询获得
	JobLevelIDList     []string                             `json:"job_level_id_list,omitempty"`     // 职务级别 ID 列表, 枚举值及详细信息可通过【批量查询职务级别】接口查询获得
	WorkingHoursTypeID string                               `json:"working_hours_type_id,omitempty"` // 工时制度 ID, 枚举值及详细信息可通过【批量查询工时制度】接口查询获得
	EffectiveTime      string                               `json:"effective_time,omitempty"`        // 生效时间
	ExpirationTime     string                               `json:"expiration_time,omitempty"`       // 失效时间
	CustomFields       []*UpdateCoreHrJobRespJobCustomField `json:"custom_fields,omitempty"`         // 自定义字段
}

// UpdateCoreHrJobRespJobCustomField ...
type UpdateCoreHrJobRespJobCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// UpdateCoreHrJobRespJobDescription ...
type UpdateCoreHrJobRespJobDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHrJobRespJobJobTitle ...
type UpdateCoreHrJobRespJobJobTitle struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// UpdateCoreHrJobRespJobName ...
type UpdateCoreHrJobRespJobName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// updateCoreHrJobResp ...
type updateCoreHrJobResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCoreHrJobResp `json:"data,omitempty"`
}
