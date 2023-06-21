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

// UpdateHireJobConfig 更新职位设置, 包括面试评价表、Offer 申请表等。接口将按照所选择的「更新选项」进行设置参数校验和更新。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job/update_config
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/recruitment-related-configuration/job/update_config
func (r *HireService) UpdateHireJobConfig(ctx context.Context, request *UpdateHireJobConfigReq, options ...MethodOptionFunc) (*UpdateHireJobConfigResp, *Response, error) {
	if r.cli.mock.mockHireUpdateHireJobConfig != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#UpdateHireJobConfig mock enable")
		return r.cli.mock.mockHireUpdateHireJobConfig(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "UpdateHireJobConfig",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/jobs/:job_id/update_config",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateHireJobConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireUpdateHireJobConfig mock HireUpdateHireJobConfig method
func (r *Mock) MockHireUpdateHireJobConfig(f func(ctx context.Context, request *UpdateHireJobConfigReq, options ...MethodOptionFunc) (*UpdateHireJobConfigResp, *Response, error)) {
	r.mockHireUpdateHireJobConfig = f
}

// UnMockHireUpdateHireJobConfig un-mock HireUpdateHireJobConfig method
func (r *Mock) UnMockHireUpdateHireJobConfig() {
	r.mockHireUpdateHireJobConfig = nil
}

// UpdateHireJobConfigReq ...
type UpdateHireJobConfigReq struct {
	JobID                         string                                            `path:"job_id" json:"-"`                            // 职位 ID, 示例值: "6960663240925956660"
	UserIDType                    *IDType                                           `query:"user_id_type" json:"-"`                     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	OfferApplySchemaID            *string                                           `json:"offer_apply_schema_id,omitempty"`            // Offer 申请表, 枚举通过接口「获取 Offer 申请表列表」获取, 示例值: "6960663240925956573"
	OfferProcessConf              *string                                           `json:"offer_process_conf,omitempty"`               // Offer 审批流, 枚举通过接口「获取 Offer 审批流列表」获取, 示例值: "6960663240925956572"
	RecommendedEvaluatorIDList    []string                                          `json:"recommended_evaluator_id_list,omitempty"`    // 建议评估人 ID 列表, 示例值: ["6960663240925956571"]
	UpdateOptionList              []int64                                           `json:"update_option_list,omitempty"`               // 更新选项, 传入要更新的配置项, 1=更新面试评价表, 2=更新 Offer 申请表, 3=更新 Offer 审批流程, 4=更新招聘需求, 5=更新建议面试官, 6=更新推荐评估人, 示例值: [1], 可选值有: 1: 更新面试评价表, 2: 更新 Offer 申请表, 3: 更新 Offer 审批流程, 4: 更新招聘需求, 5: 更新建议面试官, 6: 更新推荐评估人, 8: 更新关联职位, 9: 更新自助约面配置, 10: 更新面试登记表
	AssessmentTemplateBizID       *string                                           `json:"assessment_template_biz_id,omitempty"`       // 面试评价表, 枚举通过接口「获取面试评价表列表」获取, 示例值: "6960663240925956571"
	InterviewRoundConfList        []*UpdateHireJobConfigReqInterviewRoundConf       `json:"interview_round_conf_list,omitempty"`        // 建议面试官列表
	JrIDList                      []string                                          `json:"jr_id_list,omitempty"`                       // 关联招聘需求, 支持关联多个, 枚举通过接口「获取招聘需求」获取, 示例值: ["6960663240925956571"]
	InterviewRegistrationSchemaID *string                                           `json:"interview_registration_schema_id,omitempty"` // 面试登记表 ID, 当在飞书招聘「设置 - 面试登记表使用设置 - 面试登记表使用方式」中选择「全部职位应用同一登记表」时, 「职位设置」下的面试登记表不生效；选择「HR 按职位选择登记表」时, 该字段为必填, 示例值: "6930815272790114324"
	InterviewRoundTypeConfList    []*UpdateHireJobConfigReqInterviewRoundTypeConf   `json:"interview_round_type_conf_list,omitempty"`   // 面试轮次类型 ID 列表
	RelatedJobIDList              []string                                          `json:"related_job_id_list,omitempty"`              // 关联职位列表, 如职位为实体职位则关联虚拟职位id, 如职位为虚拟职位则关联实体职位id, 示例值: ["6960663240925956573"]
	InterviewAppointmentConfig    *UpdateHireJobConfigReqInterviewAppointmentConfig `json:"interview_appointment_config,omitempty"`     // 自助约面配置
}

// UpdateHireJobConfigReqInterviewAppointmentConfig ...
type UpdateHireJobConfigReqInterviewAppointmentConfig struct {
	EnableInterviewAppointmentByInterviewer *bool                                                   `json:"enable_interview_appointment_by_interviewer,omitempty"` // 是否开启面试官自助约面, 示例值: true
	Config                                  *UpdateHireJobConfigReqInterviewAppointmentConfigConfig `json:"config,omitempty"`                                      // 配置详情
}

// UpdateHireJobConfigReqInterviewAppointmentConfigConfig ...
type UpdateHireJobConfigReqInterviewAppointmentConfigConfig struct {
	InterviewType                         *int64   `json:"interview_type,omitempty"`                            // 面试类型, 示例值: 1, 可选值有: 1: 现场面试, 2: 电话面试, 3: 视频面试
	TalentTimezoneCode                    *string  `json:"talent_timezone_code,omitempty"`                      // 时区, 示例值: "Asia/Shanghai"
	ContactUserID                         *string  `json:"contact_user_id,omitempty"`                           // 联系人id, 示例值: "ou_c99c5f35d542efc7ee492afe11af19ef"
	ContactMobile                         *string  `json:"contact_mobile,omitempty"`                            // 联系人电话, 示例值: "151"
	ContactEmail                          *string  `json:"contact_email,omitempty"`                             // 联系人邮箱, 示例值: "test@email"
	AddressID                             *string  `json:"address_id,omitempty"`                                // 地址id, 示例值: "6960663240925956576"
	VideoType                             *int64   `json:"video_type,omitempty"`                                // 地址id, 示例值: 1, 可选值有: 1: zoom, 2: 牛客技术类型, 3: 牛客非技术类型, 4: 赛码, 5: 飞书, 8: Hackerrank, 9: 飞书(含代码考核), 100: 不使用系统工具
	Cc                                    []string `json:"cc,omitempty"`                                        // 抄送人id lsit, 示例值: ["user_id"]
	Remark                                *string  `json:"remark,omitempty"`                                    // 备注, 示例值: "备注"
	InterviewNotificationTemplateID       *string  `json:"interview_notification_template_id,omitempty"`        // 面试通知模板, 示例值: "6960663240925956573"
	AppointmentNotificationTemplateID     *string  `json:"appointment_notification_template_id,omitempty"`      // 预约通知模板, 示例值: "6960663240925956573"
	CancelInterviewNotificationTemplateID *string  `json:"cancel_interview_notification_template_id,omitempty"` // 取消面试通知, 示例值: "6960663240925956573"
}

// UpdateHireJobConfigReqInterviewRoundConf ...
type UpdateHireJobConfigReqInterviewRoundConf struct {
	InterviewerIDList []string `json:"interviewer_id_list,omitempty"` // 建议面试官 ID 列表, 示例值: ["6960663240925956571"]
	Round             *int64   `json:"round,omitempty"`               // 面试轮次, 示例值: 1
}

// UpdateHireJobConfigReqInterviewRoundTypeConf ...
type UpdateHireJobConfigReqInterviewRoundTypeConf struct {
	RoundBizID              *string `json:"round_biz_id,omitempty"`               // 面试轮次类型业务 ID, 示例值: "7012129842917837100"
	AssessmentTemplateBizID *string `json:"assessment_template_biz_id,omitempty"` // 面试评价表业务 ID, 示例值: "6960663240925956632"
}

// UpdateHireJobConfigResp ...
type UpdateHireJobConfigResp struct {
	JobConfig *UpdateHireJobConfigRespJobConfig `json:"job_config,omitempty"` // 职位信息
}

// UpdateHireJobConfigRespJobConfig ...
type UpdateHireJobConfigRespJobConfig struct {
	OfferApplySchema            *UpdateHireJobConfigRespJobConfigOfferApplySchema            `json:"offer_apply_schema,omitempty"`            // Offer 申请表
	OfferProcessConf            *UpdateHireJobConfigRespJobConfigOfferProcessConf            `json:"offer_process_conf,omitempty"`            // Offer 审批流
	RecommendedEvaluatorList    []*UpdateHireJobConfigRespJobConfigRecommendedEvaluator      `json:"recommended_evaluator_list,omitempty"`    // 建议评估人列表
	AssessmentTemplate          *UpdateHireJobConfigRespJobConfigAssessmentTemplate          `json:"assessment_template,omitempty"`           // 面试评价表
	ID                          string                                                       `json:"id,omitempty"`                            // 职位 ID
	InterviewRoundList          []*UpdateHireJobConfigRespJobConfigInterviewRound            `json:"interview_round_list,omitempty"`          // 建议面试官列表
	JobRequirementList          []*UpdateHireJobConfigRespJobConfigJobRequirement            `json:"job_requirement_list,omitempty"`          // 招聘需求
	InterviewRegistrationSchema *UpdateHireJobConfigRespJobConfigInterviewRegistrationSchema `json:"interview_registration_schema,omitempty"` // 面试登记表
	InterviewRoundTypeList      []*UpdateHireJobConfigRespJobConfigInterviewRoundType        `json:"interview_round_type_list,omitempty"`     // 面试轮次类型列表
	RelatedJobList              []*UpdateHireJobConfigRespJobConfigRelatedJob                `json:"related_job_list,omitempty"`              // 关联职位列表
	JobAttribute                int64                                                        `json:"job_attribute,omitempty"`                 // 职位属性, 1是实体职位, 2是虚拟职位
}

// UpdateHireJobConfigRespJobConfigAssessmentTemplate ...
type UpdateHireJobConfigRespJobConfigAssessmentTemplate struct {
	ID   string                                                  `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigAssessmentTemplateName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigAssessmentTemplateName ...
type UpdateHireJobConfigRespJobConfigAssessmentTemplateName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigInterviewRegistrationSchema ...
type UpdateHireJobConfigRespJobConfigInterviewRegistrationSchema struct {
	ID   string `json:"id,omitempty"`   // 面试登记表ID
	Name string `json:"name,omitempty"` // 面试登记表名称
}

// UpdateHireJobConfigRespJobConfigInterviewRound ...
type UpdateHireJobConfigRespJobConfigInterviewRound struct {
	InterviewerList []*UpdateHireJobConfigRespJobConfigInterviewRoundInterviewer `json:"interviewer_list,omitempty"` // 面试官列表
	Round           int64                                                        `json:"round,omitempty"`            // 面试轮次
}

// UpdateHireJobConfigRespJobConfigInterviewRoundInterviewer ...
type UpdateHireJobConfigRespJobConfigInterviewRoundInterviewer struct {
	ID   string                                                         `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigInterviewRoundInterviewerName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigInterviewRoundInterviewerName ...
type UpdateHireJobConfigRespJobConfigInterviewRoundInterviewerName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigInterviewRoundType ...
type UpdateHireJobConfigRespJobConfigInterviewRoundType struct {
	AssessmentRound    *UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentRound    `json:"assessment_round,omitempty"`    // 面试轮次类型
	AssessmentTemplate *UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentTemplate `json:"assessment_template,omitempty"` // 面试评价表
}

// UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentRound ...
type UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentRound struct {
	ID   string                                                                 `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentRoundName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentRoundName ...
type UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentRoundName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentTemplate ...
type UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentTemplate struct {
	ID   string                                                                    `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentTemplateName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentTemplateName ...
type UpdateHireJobConfigRespJobConfigInterviewRoundTypeAssessmentTemplateName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigJobRequirement ...
type UpdateHireJobConfigRespJobConfigJobRequirement struct {
	ID   string                                              `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigJobRequirementName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigJobRequirementName ...
type UpdateHireJobConfigRespJobConfigJobRequirementName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigOfferApplySchema ...
type UpdateHireJobConfigRespJobConfigOfferApplySchema struct {
	ID   string                                                `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigOfferApplySchemaName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigOfferApplySchemaName ...
type UpdateHireJobConfigRespJobConfigOfferApplySchemaName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigOfferProcessConf ...
type UpdateHireJobConfigRespJobConfigOfferProcessConf struct {
	ID   string                                                `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigOfferProcessConfName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigOfferProcessConfName ...
type UpdateHireJobConfigRespJobConfigOfferProcessConfName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigRecommendedEvaluator ...
type UpdateHireJobConfigRespJobConfigRecommendedEvaluator struct {
	ID   string                                                    `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigRecommendedEvaluatorName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigRecommendedEvaluatorName ...
type UpdateHireJobConfigRespJobConfigRecommendedEvaluatorName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// UpdateHireJobConfigRespJobConfigRelatedJob ...
type UpdateHireJobConfigRespJobConfigRelatedJob struct {
	ID   string                                          `json:"id,omitempty"`   // ID
	Name *UpdateHireJobConfigRespJobConfigRelatedJobName `json:"name,omitempty"` // 名称
}

// UpdateHireJobConfigRespJobConfigRelatedJobName ...
type UpdateHireJobConfigRespJobConfigRelatedJobName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// updateHireJobConfigResp ...
type updateHireJobConfigResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHireJobConfigResp `json:"data,omitempty"`
}
