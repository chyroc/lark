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

// GetCoreHRProcessFormVariableData 根据流程实例 id（process_id）获取流程表单字段数据, 包括表单里的业务字段和自定义字段。仅支持飞书人事、假勤相关业务流程。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/process-form_variable_data/get
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/process-form_variable_data/get
func (r *CoreHRService) GetCoreHRProcessFormVariableData(ctx context.Context, request *GetCoreHRProcessFormVariableDataReq, options ...MethodOptionFunc) (*GetCoreHRProcessFormVariableDataResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRProcessFormVariableData != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRProcessFormVariableData mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRProcessFormVariableData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRProcessFormVariableData",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/processes/:process_id/form_variable_data",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRProcessFormVariableDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRProcessFormVariableData mock CoreHRGetCoreHRProcessFormVariableData method
func (r *Mock) MockCoreHRGetCoreHRProcessFormVariableData(f func(ctx context.Context, request *GetCoreHRProcessFormVariableDataReq, options ...MethodOptionFunc) (*GetCoreHRProcessFormVariableDataResp, *Response, error)) {
	r.mockCoreHRGetCoreHRProcessFormVariableData = f
}

// UnMockCoreHRGetCoreHRProcessFormVariableData un-mock CoreHRGetCoreHRProcessFormVariableData method
func (r *Mock) UnMockCoreHRGetCoreHRProcessFormVariableData() {
	r.mockCoreHRGetCoreHRProcessFormVariableData = nil
}

// GetCoreHRProcessFormVariableDataReq ...
type GetCoreHRProcessFormVariableDataReq struct {
	ProcessID string `path:"process_id" json:"-"` // 流程实例 ID, 示例值: "123456987"
}

// GetCoreHRProcessFormVariableDataResp ...
type GetCoreHRProcessFormVariableDataResp struct {
	FieldVariableValues []*GetCoreHRProcessFormVariableDataRespFieldVariableValue `json:"field_variable_values,omitempty"` // 流程变量
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValue struct {
	VariableApiName string                                                               `json:"variable_api_name,omitempty"` // 变量api名称
	VariableName    *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableName  `json:"variable_name,omitempty"`     // 变量名称的i18n描述
	VariableValue   *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValue `json:"variable_value,omitempty"`    // 变量值的对象
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableName ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableName struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValue struct {
	TextValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueTextValue       `json:"text_value,omitempty"`       // 文本变量对象
	NumberValue     *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueNumberValue     `json:"number_value,omitempty"`     // 数值变量对象
	DateValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDateValue       `json:"date_value,omitempty"`       // 日期变量对象
	EmploymentValue *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEmploymentValue `json:"employment_value,omitempty"` // 员工变量对象
	DateTimeValue   *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDateTimeValue   `json:"date_time_value,omitempty"`  // 日期时间变量对象
	EnumValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValue       `json:"enum_value,omitempty"`       // 枚举变量对象
	BoolValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueBoolValue       `json:"bool_value,omitempty"`       // 布尔变量对象
	DepartmentValue *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDepartmentValue `json:"department_value,omitempty"` // 部门变量对象
	FileValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueFileValue       `json:"file_value,omitempty"`       // 文件变量对象
	I18nValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue       `json:"i18n_value,omitempty"`       // i18n变量对象
	ObjectValue     *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue     `json:"object_value,omitempty"`     // 对象变量
	ListValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValue       `json:"list_value,omitempty"`       // 列表对象
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueBoolValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueBoolValue struct {
	Value bool `json:"value,omitempty"` // 布尔变量的值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDateTimeValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDateTimeValue struct {
	Value int64  `json:"value,omitempty"` // 毫秒的时间戳
	Zone  string `json:"zone,omitempty"`  // 时区
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDateValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDateValue struct {
	Value int64 `json:"value,omitempty"` // 日期变量的值, 从1970起的天数
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDepartmentValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueDepartmentValue struct {
	Value string `json:"value,omitempty"` // 部门ID
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEmploymentValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEmploymentValue struct {
	Value  string `json:"value,omitempty"`   // employmentID
	UserID string `json:"user_id,omitempty"` // 员工ID 如3158117
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValue struct {
	Value string                                                                            `json:"value,omitempty"` // 枚举值
	Name  *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueName `json:"name,omitempty"`  // 枚举的名称
	Desc  *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueDesc `json:"desc,omitempty"`  // 枚举的描述
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueDesc ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueDesc struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueName ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueEnumValueName struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueFileValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueFileValue struct {
	SourceType int64  `json:"source_type,omitempty"` // 文件源类型（1BPM; 2主数据）
	FileID     string `json:"file_id,omitempty"`     // 文件id
	FileName   string `json:"file_name,omitempty"`   // 文件名称
	Length     int64  `json:"length,omitempty"`      // 文件长度
	MimeType   string `json:"mime_type,omitempty"`   // mime type
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValue struct {
	Value *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValueValue `json:"value,omitempty"` // i18n值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValueValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueI18nValueValue struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValue struct {
	Values []*GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValue `json:"values,omitempty"` // 列表值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValue struct {
	TextValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueTextValue       `json:"text_value,omitempty"`       // 文本变量对象
	NumberValue     *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNumberValue     `json:"number_value,omitempty"`     // 数值变量对象
	DateValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateValue       `json:"date_value,omitempty"`       // 日期变量对象
	EmploymentValue *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEmploymentValue `json:"employment_value,omitempty"` // 员工变量对象
	DateTimeValue   *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateTimeValue   `json:"date_time_value,omitempty"`  // 日期时间变量对象
	EnumValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValue       `json:"enum_value,omitempty"`       // 枚举变量对象
	BoolValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueBoolValue       `json:"bool_value,omitempty"`       // 布尔变量对象
	DepartmentValue *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDepartmentValue `json:"department_value,omitempty"` // 部门变量对象
	FileValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueFileValue       `json:"file_value,omitempty"`       // 文件变量对象
	I18nValue       *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValue       `json:"i18n_value,omitempty"`       // i18n变量对象
	ObjectValue     *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueObjectValue     `json:"object_value,omitempty"`     // 对象变量
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueBoolValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueBoolValue struct {
	Value bool `json:"value,omitempty"` // 布尔变量的值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateTimeValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateTimeValue struct {
	Value int64  `json:"value,omitempty"` // 毫秒的时间戳
	Zone  string `json:"zone,omitempty"`  // 时区
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDateValue struct {
	Value int64 `json:"value,omitempty"` // 日期变量的值, 从1970起的天数
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDepartmentValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueDepartmentValue struct {
	Value string `json:"value,omitempty"` // 部门ID
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEmploymentValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEmploymentValue struct {
	Value  string `json:"value,omitempty"`   // employmentID
	UserID string `json:"user_id,omitempty"` // 员工ID 如3158117
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValue struct {
	Value string                                                                                          `json:"value,omitempty"` // 枚举值
	Name  *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueName `json:"name,omitempty"`  // 枚举的名称
	Desc  *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueDesc `json:"desc,omitempty"`  // 枚举的描述
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueDesc ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueDesc struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueName ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueEnumValueName struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueFileValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueFileValue struct {
	SourceType int64  `json:"source_type,omitempty"` // 文件源类型（1BPM; 2主数据）
	FileID     string `json:"file_id,omitempty"`     // 文件id
	FileName   string `json:"file_name,omitempty"`   // 文件名称
	Length     int64  `json:"length,omitempty"`      // 文件长度
	MimeType   string `json:"mime_type,omitempty"`   // mime type
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValue struct {
	Value *GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValueValue `json:"value,omitempty"` // i18n值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValueValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueI18nValueValue struct {
	ZhCn string `json:"zh_cn,omitempty"` // i18n类型字段, 中文值
	EnUs string `json:"en_us,omitempty"` // i18n类型字段, 英文值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNumberValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueNumberValue struct {
	Value string `json:"value,omitempty"` // 数值类型变量的值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueObjectValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueObjectValue struct {
	Value     string `json:"value,omitempty"`       // 对象ID
	WkApiName string `json:"wk_api_name,omitempty"` // 主数据apiName
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueTextValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueListValueValueTextValue struct {
	Value string `json:"value,omitempty"` // 文本类型变量的值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueNumberValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueNumberValue struct {
	Value string `json:"value,omitempty"` // 数值类型变量的值
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueObjectValue struct {
	Value     string `json:"value,omitempty"`       // 对象ID
	WkApiName string `json:"wk_api_name,omitempty"` // 主数据apiName
}

// GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueTextValue ...
type GetCoreHRProcessFormVariableDataRespFieldVariableValueVariableValueTextValue struct {
	Value string `json:"value,omitempty"` // 文本类型变量的值
}

// getCoreHRProcessFormVariableDataResp ...
type getCoreHRProcessFormVariableDataResp struct {
	Code  int64                                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRProcessFormVariableDataResp `json:"data,omitempty"`
	Error *ErrorDetail                          `json:"error,omitempty"`
}
