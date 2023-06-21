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

// CreateHireExternalApplication 导入来自其他系统的投递信息, 创建为外部投递。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/external_application/create
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/get-candidates/import-external-system-information/create
func (r *HireService) CreateHireExternalApplication(ctx context.Context, request *CreateHireExternalApplicationReq, options ...MethodOptionFunc) (*CreateHireExternalApplicationResp, *Response, error) {
	if r.cli.mock.mockHireCreateHireExternalApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#CreateHireExternalApplication mock enable")
		return r.cli.mock.mockHireCreateHireExternalApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CreateHireExternalApplication",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/external_applications",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createHireExternalApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCreateHireExternalApplication mock HireCreateHireExternalApplication method
func (r *Mock) MockHireCreateHireExternalApplication(f func(ctx context.Context, request *CreateHireExternalApplicationReq, options ...MethodOptionFunc) (*CreateHireExternalApplicationResp, *Response, error)) {
	r.mockHireCreateHireExternalApplication = f
}

// UnMockHireCreateHireExternalApplication un-mock HireCreateHireExternalApplication method
func (r *Mock) UnMockHireCreateHireExternalApplication() {
	r.mockHireCreateHireExternalApplication = nil
}

// CreateHireExternalApplicationReq ...
type CreateHireExternalApplicationReq struct {
	ExternalID         *string `json:"external_id,omitempty"`          // 外部系统投递主键 （仅用于幂等）, 示例值: "123"
	JobRecruitmentType *int64  `json:"job_recruitment_type,omitempty"` // 职位招聘类型, 示例值: 1, 可选值有: 1: 社招, 2: 校招
	JobTitle           *string `json:"job_title,omitempty"`            // 职位名称, 示例值: "高级Java"
	ResumeSource       *string `json:"resume_source,omitempty"`        // 简历来源, 示例值: "lagou"
	Stage              *string `json:"stage,omitempty"`                // 阶段, 示例值: "1"
	TalentID           string  `json:"talent_id,omitempty"`            // 人才 ID, 示例值: "6960663240925956459"
	TerminationReason  *string `json:"termination_reason,omitempty"`   // 终止原因, 示例值: "不合适"
	DeliveryType       *int64  `json:"delivery_type,omitempty"`        // 投递类型, 示例值: 1, 可选值有: 1: HR 寻访, 2: 候选人主动投递, 3: 人才推荐, 4: 其他
	ModifyTime         *int64  `json:"modify_time,omitempty"`          // 更新时间, 示例值: 1618500278645
	TerminationType    *string `json:"termination_type,omitempty"`     // 终止类型, 示例值: "health"
}

// CreateHireExternalApplicationResp ...
type CreateHireExternalApplicationResp struct {
	ExternalApplication *CreateHireExternalApplicationRespExternalApplication `json:"external_application,omitempty"` // 外部投递信息
}

// CreateHireExternalApplicationRespExternalApplication ...
type CreateHireExternalApplicationRespExternalApplication struct {
	ID                 string `json:"id,omitempty"`                   // 外部投递 ID
	ExternalID         string `json:"external_id,omitempty"`          // 外部系统投递主键 （仅用于幂等）
	JobRecruitmentType int64  `json:"job_recruitment_type,omitempty"` // 职位招聘类型, 可选值有: 1: 社招, 2: 校招
	JobTitle           string `json:"job_title,omitempty"`            // 职位名称
	ResumeSource       string `json:"resume_source,omitempty"`        // 简历来源
	Stage              string `json:"stage,omitempty"`                // 阶段
	TalentID           string `json:"talent_id,omitempty"`            // 人才 ID
	TerminationReason  string `json:"termination_reason,omitempty"`   // 终止原因
	DeliveryType       int64  `json:"delivery_type,omitempty"`        // 投递类型, 可选值有: 1: HR 寻访, 2: 候选人主动投递, 3: 人才推荐, 4: 其他
	ModifyTime         int64  `json:"modify_time,omitempty"`          // 更新时间
	TerminationType    string `json:"termination_type,omitempty"`     // 终止类型
}

// createHireExternalApplicationResp ...
type createHireExternalApplicationResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *CreateHireExternalApplicationResp `json:"data,omitempty"`
}
