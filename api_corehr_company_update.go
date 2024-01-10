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

// UpdateCoreHRCompany 更新公司信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/patch
func (r *CoreHRService) UpdateCoreHRCompany(ctx context.Context, request *UpdateCoreHRCompanyReq, options ...MethodOptionFunc) (*UpdateCoreHRCompanyResp, *Response, error) {
	if r.cli.mock.mockCoreHRUpdateCoreHRCompany != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#UpdateCoreHRCompany mock enable")
		return r.cli.mock.mockCoreHRUpdateCoreHRCompany(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "UpdateCoreHRCompany",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/companies/:company_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateCoreHRCompanyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRUpdateCoreHRCompany mock CoreHRUpdateCoreHRCompany method
func (r *Mock) MockCoreHRUpdateCoreHRCompany(f func(ctx context.Context, request *UpdateCoreHRCompanyReq, options ...MethodOptionFunc) (*UpdateCoreHRCompanyResp, *Response, error)) {
	r.mockCoreHRUpdateCoreHRCompany = f
}

// UnMockCoreHRUpdateCoreHRCompany un-mock CoreHRUpdateCoreHRCompany method
func (r *Mock) UnMockCoreHRUpdateCoreHRCompany() {
	r.mockCoreHRUpdateCoreHRCompany = nil
}

// UpdateCoreHRCompanyReq ...
type UpdateCoreHRCompanyReq struct {
	CompanyID           string                                       `path:"company_id" json:"-"`            // 需要更新的公司 ID, 示例值: "1616161616"
	ClientToken         *string                                      `query:"client_token" json:"-"`         // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	HiberarchyCommon    *UpdateCoreHRCompanyReqHiberarchyCommon      `json:"hiberarchy_common,omitempty"`    // 层级关系, 内层字段见实体
	Type                *UpdateCoreHRCompanyReqType                  `json:"type,omitempty"`                 // 性质, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)公司类型（company_type）枚举定义部分获得
	IndustryList        []*UpdateCoreHRCompanyReqIndustry            `json:"industry_list,omitempty"`        // 行业, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)行业（industry）枚举定义部分获得
	LegalRepresentative []*UpdateCoreHRCompanyReqLegalRepresentative `json:"legal_representative,omitempty"` // 法定代表人
	PostCode            *string                                      `json:"post_code,omitempty"`            // 邮编, 示例值: "邮编"
	TaxPayerID          *string                                      `json:"tax_payer_id,omitempty"`         // 纳税人识别号, 示例值: "123456840"
	Confidential        *bool                                        `json:"confidential,omitempty"`         // confidential, 示例值: true
	SubTypeList         []*UpdateCoreHRCompanyReqSubType             `json:"sub_type_list,omitempty"`        // 主体类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)主体类型（company_sub_type）枚举定义部分获得
	BranchCompany       *bool                                        `json:"branch_company,omitempty"`       // 是否为分公司, 示例值: true
	PrimaryManager      []*UpdateCoreHRCompanyReqPrimaryManager      `json:"primary_manager,omitempty"`      // 主要负责人
	CustomFields        []*UpdateCoreHRCompanyReqCustomField         `json:"custom_fields,omitempty"`        // 自定义字段
	Currency            *UpdateCoreHRCompanyReqCurrency              `json:"currency,omitempty"`             // 默认币种
	Phone               *UpdateCoreHRCompanyReqPhone                 `json:"phone,omitempty"`                // 电话
	Fax                 *UpdateCoreHRCompanyReqFax                   `json:"fax,omitempty"`                  // 传真
}

// UpdateCoreHRCompanyReqCurrency ...
type UpdateCoreHRCompanyReqCurrency struct {
	CurrencyName       []*UpdateCoreHRCompanyReqCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        *int64                                        `json:"numeric_code,omitempty"`          // 数字代码, 示例值: 12
	CurrencyAlpha3Code *string                                       `json:"currency_alpha_3_code,omitempty"` // 三位字母代码, 示例值: "12"
}

// UpdateCoreHRCompanyReqCurrencyCurrencyName ...
type UpdateCoreHRCompanyReqCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "刘梓新"
}

// UpdateCoreHRCompanyReqCustomField ...
type UpdateCoreHRCompanyReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05]), 示例值: "Sandy"
}

// UpdateCoreHRCompanyReqFax ...
type UpdateCoreHRCompanyReqFax struct {
	AreaCode    *UpdateCoreHRCompanyReqFaxAreaCode `json:"area_code,omitempty"`    // 区号, 示例值: 123123
	PhoneNumber string                             `json:"phone_number,omitempty"` // 号码, 示例值: "213213"
}

// UpdateCoreHRCompanyReqFaxAreaCode ...
type UpdateCoreHRCompanyReqFaxAreaCode struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "phone_type"
}

// UpdateCoreHRCompanyReqHiberarchyCommon ...
type UpdateCoreHRCompanyReqHiberarchyCommon struct {
	ParentID       *string                                              `json:"parent_id,omitempty"`       // 上级组织, 示例值: "4719168654814483759"
	Name           []*UpdateCoreHRCompanyReqHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *UpdateCoreHRCompanyReqHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                 `json:"active,omitempty"`          // 是否启用, 示例值: true
	EffectiveTime  *string                                              `json:"effective_time,omitempty"`  // 生效时间, 示例值: "2020-05-01 00:00:00"
	ExpirationTime *string                                              `json:"expiration_time,omitempty"` // 失效时间, 示例值: "2020-05-02 00:00:00"
	Code           *string                                              `json:"code,omitempty"`            // 编码, 示例值: "12456"
	Description    []*UpdateCoreHRCompanyReqHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	CustomFields   []*UpdateCoreHRCompanyReqHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// UpdateCoreHRCompanyReqHiberarchyCommonCustomField ...
type UpdateCoreHRCompanyReqHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05]), 示例值: "Sandy"
}

// UpdateCoreHRCompanyReqHiberarchyCommonDescription ...
type UpdateCoreHRCompanyReqHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "刘梓新"
}

// UpdateCoreHRCompanyReqHiberarchyCommonName ...
type UpdateCoreHRCompanyReqHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "刘梓新"
}

// UpdateCoreHRCompanyReqHiberarchyCommonType ...
type UpdateCoreHRCompanyReqHiberarchyCommonType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "phone_type"
}

// UpdateCoreHRCompanyReqIndustry ...
type UpdateCoreHRCompanyReqIndustry struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "phone_type"
}

// UpdateCoreHRCompanyReqLegalRepresentative ...
type UpdateCoreHRCompanyReqLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "刘梓新"
}

// UpdateCoreHRCompanyReqPhone ...
type UpdateCoreHRCompanyReqPhone struct {
	AreaCode    *UpdateCoreHRCompanyReqPhoneAreaCode `json:"area_code,omitempty"`    // 区号, 示例值: 123123
	PhoneNumber string                               `json:"phone_number,omitempty"` // 号码, 示例值: "213213"
}

// UpdateCoreHRCompanyReqPhoneAreaCode ...
type UpdateCoreHRCompanyReqPhoneAreaCode struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "phone_type"
}

// UpdateCoreHRCompanyReqPrimaryManager ...
type UpdateCoreHRCompanyReqPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "刘梓新"
}

// UpdateCoreHRCompanyReqSubType ...
type UpdateCoreHRCompanyReqSubType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "phone_type"
}

// UpdateCoreHRCompanyReqType ...
type UpdateCoreHRCompanyReqType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "phone_type"
}

// UpdateCoreHRCompanyResp ...
type UpdateCoreHRCompanyResp struct {
	Company *UpdateCoreHRCompanyRespCompany `json:"company,omitempty"` // 公司
}

// UpdateCoreHRCompanyRespCompany ...
type UpdateCoreHRCompanyRespCompany struct {
	ID                      string                                                  `json:"id,omitempty"`                        // 实体在CoreHR内部的唯一键
	HiberarchyCommon        *UpdateCoreHRCompanyRespCompanyHiberarchyCommon         `json:"hiberarchy_common,omitempty"`         // 层级关系, 内层字段见实体
	Type                    *UpdateCoreHRCompanyRespCompanyType                     `json:"type,omitempty"`                      // 性质, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)公司类型（company_type）枚举定义部分获得
	IndustryList            []*UpdateCoreHRCompanyRespCompanyIndustry               `json:"industry_list,omitempty"`             // 行业, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)行业（industry）枚举定义部分获得
	LegalRepresentative     []*UpdateCoreHRCompanyRespCompanyLegalRepresentative    `json:"legal_representative,omitempty"`      // 法定代表人
	PostCode                string                                                  `json:"post_code,omitempty"`                 // 邮编
	TaxPayerID              string                                                  `json:"tax_payer_id,omitempty"`              // 纳税人识别号
	Confidential            bool                                                    `json:"confidential,omitempty"`              // confidential
	SubTypeList             []*UpdateCoreHRCompanyRespCompanySubType                `json:"sub_type_list,omitempty"`             // 主体类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)主体类型（company_sub_type）枚举定义部分获得
	BranchCompany           bool                                                    `json:"branch_company,omitempty"`            // 是否为分公司
	PrimaryManager          []*UpdateCoreHRCompanyRespCompanyPrimaryManager         `json:"primary_manager,omitempty"`           // 主要负责人
	CustomFields            []*UpdateCoreHRCompanyRespCompanyCustomField            `json:"custom_fields,omitempty"`             // 自定义字段
	Currency                *UpdateCoreHRCompanyRespCompanyCurrency                 `json:"currency,omitempty"`                  // 默认币种
	Phone                   *UpdateCoreHRCompanyRespCompanyPhone                    `json:"phone,omitempty"`                     // 电话
	Fax                     *UpdateCoreHRCompanyRespCompanyFax                      `json:"fax,omitempty"`                       // 传真
	RegisteredOfficeAddress []*UpdateCoreHRCompanyRespCompanyRegisteredOfficeAddres `json:"registered_office_address,omitempty"` // 注册地址
	OfficeAddress           []*UpdateCoreHRCompanyRespCompanyOfficeAddres           `json:"office_address,omitempty"`            // 办公地址
}

// UpdateCoreHRCompanyRespCompanyCurrency ...
type UpdateCoreHRCompanyRespCompanyCurrency struct {
	ID                 string                                                `json:"id,omitempty"`                    // 货币id
	CountryRegionID    string                                                `json:"country_region_id,omitempty"`     // 货币所属国家/地区id, 详细信息可通过[【查询国家/地区信息】](https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/location_data/list)接口查询获得
	CurrencyName       []*UpdateCoreHRCompanyRespCompanyCurrencyCurrencyName `json:"currency_name,omitempty"`         // 货币名称
	NumericCode        int64                                                 `json:"numeric_code,omitempty"`          // 数字代码
	CurrencyAlpha3Code string                                                `json:"currency_alpha_3_code,omitempty"` // 三位字母代码
}

// UpdateCoreHRCompanyRespCompanyCurrencyCurrencyName ...
type UpdateCoreHRCompanyRespCompanyCurrencyCurrencyName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyCustomField ...
type UpdateCoreHRCompanyRespCompanyCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// UpdateCoreHRCompanyRespCompanyFax ...
type UpdateCoreHRCompanyRespCompanyFax struct {
	AreaCode    *UpdateCoreHRCompanyRespCompanyFaxAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                     `json:"phone_number,omitempty"` // 号码
}

// UpdateCoreHRCompanyRespCompanyFaxAreaCode ...
type UpdateCoreHRCompanyRespCompanyFaxAreaCode struct {
	EnumName string                                              `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRCompanyRespCompanyFaxAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRCompanyRespCompanyFaxAreaCodeDisplay ...
type UpdateCoreHRCompanyRespCompanyFaxAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyHiberarchyCommon ...
type UpdateCoreHRCompanyRespCompanyHiberarchyCommon struct {
	ParentID       string                                                       `json:"parent_id,omitempty"`       // 上级组织
	Name           []*UpdateCoreHRCompanyRespCompanyHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *UpdateCoreHRCompanyRespCompanyHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                         `json:"active,omitempty"`          // 启用
	EffectiveTime  string                                                       `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                       `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                       `json:"code,omitempty"`            // 编码
	Description    []*UpdateCoreHRCompanyRespCompanyHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                       `json:"tree_order,omitempty"`      // 树形排序, 代表同层级的部门排序序号
	ListOrder      string                                                       `json:"list_order,omitempty"`      // 列表排序, 代表所有部门的混排序号
	CustomFields   []*UpdateCoreHRCompanyRespCompanyHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// UpdateCoreHRCompanyRespCompanyHiberarchyCommonCustomField ...
type UpdateCoreHRCompanyRespCompanyHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(123, 123.23, true, [\"id1\", \"id2\], 2006-01-02 15:04:05])
}

// UpdateCoreHRCompanyRespCompanyHiberarchyCommonDescription ...
type UpdateCoreHRCompanyRespCompanyHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyHiberarchyCommonName ...
type UpdateCoreHRCompanyRespCompanyHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyHiberarchyCommonType ...
type UpdateCoreHRCompanyRespCompanyHiberarchyCommonType struct {
	EnumName string                                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay ...
type UpdateCoreHRCompanyRespCompanyHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyIndustry ...
type UpdateCoreHRCompanyRespCompanyIndustry struct {
	EnumName string                                           `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRCompanyRespCompanyIndustryDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRCompanyRespCompanyIndustryDisplay ...
type UpdateCoreHRCompanyRespCompanyIndustryDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyLegalRepresentative ...
type UpdateCoreHRCompanyRespCompanyLegalRepresentative struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyOfficeAddres ...
type UpdateCoreHRCompanyRespCompanyOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyPhone ...
type UpdateCoreHRCompanyRespCompanyPhone struct {
	AreaCode    *UpdateCoreHRCompanyRespCompanyPhoneAreaCode `json:"area_code,omitempty"`    // 区号
	PhoneNumber string                                       `json:"phone_number,omitempty"` // 号码
}

// UpdateCoreHRCompanyRespCompanyPhoneAreaCode ...
type UpdateCoreHRCompanyRespCompanyPhoneAreaCode struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRCompanyRespCompanyPhoneAreaCodeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRCompanyRespCompanyPhoneAreaCodeDisplay ...
type UpdateCoreHRCompanyRespCompanyPhoneAreaCodeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyPrimaryManager ...
type UpdateCoreHRCompanyRespCompanyPrimaryManager struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyRegisteredOfficeAddres ...
type UpdateCoreHRCompanyRespCompanyRegisteredOfficeAddres struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanySubType ...
type UpdateCoreHRCompanyRespCompanySubType struct {
	EnumName string                                          `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRCompanyRespCompanySubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRCompanyRespCompanySubTypeDisplay ...
type UpdateCoreHRCompanyRespCompanySubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// UpdateCoreHRCompanyRespCompanyType ...
type UpdateCoreHRCompanyRespCompanyType struct {
	EnumName string                                       `json:"enum_name,omitempty"` // 枚举值
	Display  []*UpdateCoreHRCompanyRespCompanyTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// UpdateCoreHRCompanyRespCompanyTypeDisplay ...
type UpdateCoreHRCompanyRespCompanyTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// updateCoreHRCompanyResp ...
type updateCoreHRCompanyResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCoreHRCompanyResp `json:"data,omitempty"`
}
