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

// CreateCoreHRPreHire 创建待入职人员。
//
// Offer ID 仅支持传入飞书招聘 ID, 如果未使用飞书招聘请置空。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/pre_hire/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/pre_hire/create
func (r *CoreHRService) CreateCoreHRPreHire(ctx context.Context, request *CreateCoreHRPreHireReq, options ...MethodOptionFunc) (*CreateCoreHRPreHireResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRPreHire != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRPreHire mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRPreHire(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRPreHire",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/pre_hires",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRPreHireResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRPreHire mock CoreHRCreateCoreHRPreHire method
func (r *Mock) MockCoreHRCreateCoreHRPreHire(f func(ctx context.Context, request *CreateCoreHRPreHireReq, options ...MethodOptionFunc) (*CreateCoreHRPreHireResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRPreHire = f
}

// UnMockCoreHRCreateCoreHRPreHire un-mock CoreHRCreateCoreHRPreHire method
func (r *Mock) UnMockCoreHRCreateCoreHRPreHire() {
	r.mockCoreHRCreateCoreHRPreHire = nil
}

// CreateCoreHRPreHireReq ...
type CreateCoreHRPreHireReq struct {
	BasicInfo        *CreateCoreHRPreHireReqBasicInfo        `json:"basic_info,omitempty"`         // 个人信息
	OfferInfo        *CreateCoreHRPreHireReqOfferInfo        `json:"offer_info,omitempty"`         // 职位信息
	EducationInfo    []*CreateCoreHRPreHireReqEducationInfo  `json:"education_info,omitempty"`     // 教育经历
	WorkExperience   []*CreateCoreHRPreHireReqWorkExperience `json:"work_experience,omitempty"`    // 工作经历
	AtsApplicationID *string                                 `json:"ats_application_id,omitempty"` // 招聘应用 ID, 仅支持飞书招聘 ID, 详细信息可以通过招聘【获取投递信息】接口查询获得, 示例值: "7140946969586010376"
}

// CreateCoreHRPreHireReqBasicInfo ...
type CreateCoreHRPreHireReqBasicInfo struct {
	Name                  *CreateCoreHRPreHireReqBasicInfoName `json:"name,omitempty"`                    // 描述
	PhoneNumber           *string                              `json:"phone_number,omitempty"`            // 手机号, 示例值: "31123127"
	InternationalAreaCode *string                              `json:"international_area_code,omitempty"` // 区号, 示例值: "86_china"
	Email                 *string                              `json:"email,omitempty"`                   // 个人邮箱, 示例值: "xx@xx.com"
	DateOfBirth           *string                              `json:"date_of_birth,omitempty"`           // 生日, 示例值: "2011-99-99"
	PersonalIDNumber      *string                              `json:"personal_id_number,omitempty"`      // 证件号, 示例值: "31123127"
	DateEnteredWorkforce  *string                              `json:"date_entered_workforce,omitempty"`  // 参加工作日期, 示例值: "2100-09-09"
	GenderID              *string                              `json:"gender_id,omitempty"`               // 性别, 示例值: "male"
	NationalityID         *string                              `json:"nationality_id,omitempty"`          // 国籍, 示例值: "6862995757234914824"
	HomeAddress           *string                              `json:"home_address,omitempty"`            // 家庭地址, 示例值: "home addr"
	WorkerID              *string                              `json:"worker_id,omitempty"`               // 人员编号, 示例值: "6862995757234914824"
}

// CreateCoreHRPreHireReqBasicInfoName ...
type CreateCoreHRPreHireReqBasicInfoName struct {
	FullName        *string `json:"full_name,omitempty"`         // 全名, 示例值: "李一一"
	FirstName       *string `json:"first_name,omitempty"`        // 名, 示例值: "一"
	MiddleName      *string `json:"middle_name,omitempty"`       // 中间名, 示例值: "一"
	NamePrimary     *string `json:"name_primary,omitempty"`      // 姓, 示例值: "李"
	LocalFirstName  *string `json:"local_first_name,omitempty"`  // 名 - 本地文字, 示例值: "一"
	LocalMiddleName *string `json:"local_middle_name,omitempty"` // 本地中间名, 示例值: "一"
	LocalPrimary    *string `json:"local_primary,omitempty"`     // 姓 - 本地文字, 示例值: "李"
}

// CreateCoreHRPreHireReqEducationInfo ...
type CreateCoreHRPreHireReqEducationInfo struct {
	SchoolName   *string `json:"school_name,omitempty"`    // 学校名称, 示例值: "长安大学"
	Education    *string `json:"education,omitempty"`      // 学历, 示例值: "phd"
	StartTime    *string `json:"start_time,omitempty"`     // 开始时间, 示例值: "2017-04-01"
	EndTime      *string `json:"end_time,omitempty"`       // 结束时间, 示例值: "2018-04-01"
	FieldOfStudy *string `json:"field_of_study,omitempty"` // 专业, 示例值: "医学影像技术"
}

// CreateCoreHRPreHireReqOfferInfo ...
type CreateCoreHRPreHireReqOfferInfo struct {
	OfferID              *string                                          `json:"offer_id,omitempty"`                // Offer ID, 仅支持飞书招聘 ID, 可以通过飞书招聘【获取 Offer 申请表列表】接口获取, 示例值: "7032210902531327521"
	OfferHRID            *string                                          `json:"offer_hr_id,omitempty"`             // Offer HR ID, 示例值: "7032210902531327521"
	DepartmentID         *string                                          `json:"department_id,omitempty"`           // 部门 ID, 可以通过【批量获取部门】接口获取, 示例值: "7147562782945478177"
	DirectLeaderID       *string                                          `json:"direct_leader_id,omitempty"`        // 直属领导的雇佣 ID, 可以通过【查询员工信息】接口获取, 示例值: "7032210902531327521"
	JobID                *string                                          `json:"job_id,omitempty"`                  // 职务 ID, 可以通过招聘【批量查询职务】接口获取, 示例值: "6977976735715378724"
	JobFamilyID          *string                                          `json:"job_family_id,omitempty"`           // 职务序列 ID, 可以通过招聘【批量查询职务序列】接口获取, 示例值: "6977972856625939999"
	JobLevelID           *string                                          `json:"job_level_id,omitempty"`            // 职务级别 ID, 可以通过招聘【批量查询职务级别】接口获取, 示例值: "6977971894960145950"
	JobTitle             *string                                          `json:"job_title,omitempty"`               // 职务头衔, 示例值: "java"
	ProbationStartDate   *string                                          `json:"probation_start_date,omitempty"`    // 试用期开始日期, 示例值: "2022-07-29"
	ProbationEndDate     *string                                          `json:"probation_end_date,omitempty"`      // 试用期结束日期, 示例值: "2023-04-07"
	ContractStartDate    *string                                          `json:"contract_start_date,omitempty"`     // 合同开始日期, 示例值: "2022-10-08"
	ContractEndDate      *string                                          `json:"contract_end_date,omitempty"`       // 合同结束日期, 示例值: "2025-10-07"
	OnboardingDate       *string                                          `json:"onboarding_date,omitempty"`         // 入职日期, 示例值: "2022-10-08"
	OnboardingLocationID *string                                          `json:"onboarding_location_id,omitempty"`  // 入职地点 ID, 详细信息可通过【批量查询地点】接口获得, 示例值: "6977976687350924832"
	OfficeLocationID     *string                                          `json:"office_location_id,omitempty"`      // 办公地点 ID, 详细信息可通过【批量查询地点】接口获得, 示例值: "6977976687350924832"
	RecruitmentTypeID    *string                                          `json:"recruitment_type_id,omitempty"`     // 招聘来源, 枚举值 api_name 可通过【获取自定义字段详情】接口查询, 查询参数如下: object_api_name = "pre_hire", custom_api_name = "recruitment_type", 示例值: "experienced_professionals"
	ProbationPeriod      *string                                          `json:"probation_period,omitempty"`        // 试用期时长, 示例值: "6"
	EmployeeTypeID       *string                                          `json:"employee_type_id,omitempty"`        // 人员类型 ID, 可以通过招聘【批量查询人员类型】接口获取, 示例值: "6977973225846343171"
	EmploymentTypeID     *string                                          `json:"employment_type_id,omitempty"`      // 雇佣类型, 枚举值 api_name 可通过【获取自定义字段详情】接口查询, 查询参数如下: object_api_name = "pre_hire", custom_api_name = "employment_type", 示例值: "6977973225846343171"
	WorkEmail            *string                                          `json:"work_email,omitempty"`              // 工作邮箱, 示例值: "joshua@bytedance.com"
	DurationTypeID       *string                                          `json:"duration_type_id,omitempty"`        // 期限类型, 枚举值 api_name 可通过【获取自定义字段详情】接口查询, 查询参数如下: object_api_name = "pre_hire", custom_api_name = "duration_type", 示例值: "6977973225846343171"
	SigningTypeID        *string                                          `json:"signing_type_id,omitempty"`         // 签订类型, 枚举值 api_name 可通过【获取自定义字段详情】接口查询, 查询参数如下: object_api_name = "pre_hire", custom_api_name = "signing_type", 示例值: "6738317738688661772"
	EntryMode            *string                                          `json:"entry_mode,omitempty"`              // 入职方式, 枚举值 api_name 可通过【获取自定义字段详情】接口查询, 查询参数如下: object_api_name = "pre_hire", custom_api_name = "onboarding_method", 示例值: "1"
	SocialSecurityCityID *string                                          `json:"social_security_city_id,omitempty"` // 社保城市 ID, 详细信息可通过【批量查询地点】接口获得, 示例值: "xxx"
	ContractType         *string                                          `json:"contract_type,omitempty"`           // 合同类型, 枚举值 api_name 可通过【获取自定义字段详情】接口查询, 查询参数如下: object_api_name = "pre_hire", custom_api_name = "contract_type", 示例值: "6738317738688661772"
	Company              *string                                          `json:"company,omitempty"`                 // 公司 ID, 详细信息可通过【批量查询公司】接口获得, 示例值: "6738317738688661772"
	CostCenterRate       []*CreateCoreHRPreHireReqOfferInfoCostCenterRate `json:"cost_center_rate,omitempty"`        // 成本中心分摊信息
	CustomFields         []*CreateCoreHRPreHireReqOfferInfoCustomField    `json:"custom_fields,omitempty"`           // 自定义字段
}

// CreateCoreHRPreHireReqOfferInfoCostCenterRate ...
type CreateCoreHRPreHireReqOfferInfoCostCenterRate struct {
	CostCenterID *string `json:"cost_center_id,omitempty"` // 成本中心 ID, 可以通过【查询单个成本中心信息】接口获取对应的成本中心信息, 示例值: "6950635856373745165"
	Rate         *int64  `json:"rate,omitempty"`           // 分摊比例, 示例值: 100
}

// CreateCoreHRPreHireReqOfferInfoCustomField ...
type CreateCoreHRPreHireReqOfferInfoCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 请转换为字符串数组的方式写入, 示例值: "[\"Sandy\"]"
}

// CreateCoreHRPreHireReqWorkExperience ...
type CreateCoreHRPreHireReqWorkExperience struct {
	CompanyName *string `json:"company_name,omitempty"` // 公司名称, 示例值: "猎豹"
	StartTime   *string `json:"start_time,omitempty"`   // 开始时间, 示例值: "2015-02-01"
	EndTime     *string `json:"end_time,omitempty"`     // 结束时间, 示例值: "2017-02-01"
	JobTitle    *string `json:"job_title,omitempty"`    // 岗位, 示例值: "产品经理"
	Description *string `json:"description,omitempty"`  // 工作描述, 示例值: "app"
}

// CreateCoreHRPreHireResp ...
type CreateCoreHRPreHireResp struct {
	PreHireID string `json:"pre_hire_id,omitempty"` // 待入职 ID
}

// createCoreHRPreHireResp ...
type createCoreHRPreHireResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHRPreHireResp `json:"data,omitempty"`
}
