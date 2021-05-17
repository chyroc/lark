// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetEmployeeList 根据员工飞书用户 ID / 员工状态 / 雇员类型等搜索条件 ，批量获取员工花名册字段信息。字段包括「系统标准字段 / system_fields」和「自定义字段 / custom_fields」
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/employee/list
func (r *EHRService) GetEmployeeList(ctx context.Context, request *GetEmployeeListReq, options ...MethodOptionFunc) (*GetEmployeeListResp, *Response, error) {
	if r.cli.mock.mockEHRGetEmployeeList != nil {
		r.cli.logDebug(ctx, "[lark] EHR#GetEmployeeList mock enable")
		return r.cli.mock.mockEHRGetEmployeeList(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] EHR#GetEmployeeList call api")
	r.cli.logDebug(ctx, "[lark] EHR#GetEmployeeList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/ehr/v1/employees",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getEmployeeListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] EHR#GetEmployeeList GET https://open.feishu.cn/open-apis/ehr/v1/employees failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] EHR#GetEmployeeList GET https://open.feishu.cn/open-apis/ehr/v1/employees failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("EHR", "GetEmployeeList", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] EHR#GetEmployeeList request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockEHRGetEmployeeList(f func(ctx context.Context, request *GetEmployeeListReq, options ...MethodOptionFunc) (*GetEmployeeListResp, *Response, error)) {
	r.mockEHRGetEmployeeList = f
}

func (r *Mock) UnMockEHRGetEmployeeList() {
	r.mockEHRGetEmployeeList = nil
}

type GetEmployeeListReq struct {
	View       *string `query:"view" json:"-"`         // 返回数据类型, 示例值："basic", 可选值有: `basic`：概览，只返回 id、name 等基本信息, `full`：明细，返回系统标准字段和自定义字段集合
	Status     []int   `query:"status" json:"-"`       // 员工状态，不传代表查询所有员工状态,实际在职 = 2&4
	Type       []int   `query:"type" json:"-"`         // 雇员类型，不传代表查询所有雇员类型
	StartTime  *int    `query:"start_time" json:"-"`   // 查询开始时间（创建时间 &gt;= 此时间）, 示例值：1608690517811
	EndTime    *int    `query:"end_time" json:"-"`     // 查询结束时间（创建时间 &lt;= 此时间）, 示例值：1608690517811
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	UserIDs    *string `query:"user_ids" json:"-"`     // user_id、open_id 或 union_id，默认为 open_id。,如果传入的值不是 open_id，需要一并传入 user_id_type 参数。, 示例值："ou_8ebd4f35d7101ffdeb4771d7c8ec517e"
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："1cGaS5rM1F9u9g/xjVMipKrmQHo5CHYa+uE3D3OnAmc="
	PageSize   *int    `query:"page_size" json:"-"`    // 分页大小, 示例值：10, 最大值：`100`
}

type getEmployeeListResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetEmployeeListResp `json:"data,omitempty"` //
}

type GetEmployeeListResp struct {
	Items     []*GetEmployeeListRespItem `json:"items,omitempty"`      // 员工列表
	PageToken string                     `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                       `json:"has_more,omitempty"`   // 是否还有更多项
}

type GetEmployeeListRespItem struct {
	UserID       string                                `json:"user_id,omitempty"`       // 员工的用户 ID,user_id_type 为 user_id 时返回 user_id；,user_id_type 为 open_id 时返回 open_id；,user_id_type 为 union_id 时返回 uion_id；,「待入职」和「已取消入职」的员工，此字段值为 null
	SystemFields *GetEmployeeListRespItemSystemFields  `json:"system_fields,omitempty"` // 系统字段
	CustomFields []*GetEmployeeListRespItemCustomField `json:"custom_fields,omitempty"` // 自定义字段
}

type GetEmployeeListRespItemSystemFields struct {
	Name                    string                                              `json:"name,omitempty"`                      // 中文姓名
	EnName                  string                                              `json:"en_name,omitempty"`                   // 英文姓名
	Email                   string                                              `json:"email,omitempty"`                     // 邮箱
	Mobile                  string                                              `json:"mobile,omitempty"`                    // 手机号码
	DepartmentID            string                                              `json:"department_id,omitempty"`             // 部门的飞书 open_department_id
	Manager                 *GetEmployeeListRespItemSystemFieldsManager         `json:"manager,omitempty"`                   // 上级
	Job                     *GetEmployeeListRespItemSystemFieldsJob             `json:"job,omitempty"`                       // 职位
	JobLevel                *GetEmployeeListRespItemSystemFieldsJobLevel        `json:"job_level,omitempty"`                 // 职级
	WorkLocation            *GetEmployeeListRespItemSystemFieldsWorkLocation    `json:"work_location,omitempty"`             // 工作地点
	Gender                  int                                                 `json:"gender,omitempty"`                    // 性别, 可选值有: `1`：男, `2`：女
	Birthday                string                                              `json:"birthday,omitempty"`                  // 出生日期
	NativeRegion            *GetEmployeeListRespItemSystemFieldsNativeRegion    `json:"native_region,omitempty"`             // 籍贯
	Ethnicity               int                                                 `json:"ethnicity,omitempty"`                 // 民族, 可选值有: `1`：汉族, `2`：蒙古族, `3`：回族, `4`：藏族, `5`：维吾尔族, `6`：苗族, `7`：彝族, `8`：壮族, `9`：布依族, `10`：朝鲜族, `11`：满族, `12`：侗族, `13`：瑶族, `14`：白族, `15`：土家族, `16`：哈尼族, `17`：哈萨克族, `18`：傣族, `19`：黎族, `20`：傈僳族, `21`：佤族, `22`：畲族, `23`：高山族, `24`：拉祜族, `25`：水族, `26`：东乡族, `27`：纳西族, `28`：景颇族, `29`：阿昌族, `30`：柯尔克孜族, `31`：土族, `32`：达斡尔族, `33`：仫佬族, `34`：羌族, `35`：布朗族, `36`：撒拉族, `37`：毛南族, `38`：仡佬族, `39`：锡伯族, `40`：普米族, `41`：塔吉克族, `42`：怒族, `43`：乌孜别克族, `44`：俄罗斯族, `45`：鄂温克族, `46`：德昂族, `47`：保安族, `48`：裕固族, `49`：京族, `50`：塔塔尔族, `51`：独龙族, `52`：鄂伦春族, `53`：赫哲族, `54`：门巴族, `55`：珞巴族, `56`：基诺族, `57`：其他
	MaritalStatus           int                                                 `json:"marital_status,omitempty"`            // 婚姻状况, 可选值有: `1`：未婚, `2`：已婚, `3`：离异, `4`：其他
	PoliticalStatus         int                                                 `json:"political_status,omitempty"`          // 政治面貌, 可选值有: `1`：中共党员, `2`：中国农工民主党, `3`：中国国民党革命委员会, `4`：中国民主促进会会员, `5`：中国民主同盟成员, `6`：中国民主建国会, `7`：中国致公党党员, `8`：九三学社社员, `9`：共青团员, `10`：其它党派成员, `11`：民主人士, `12`：群众
	EnteredWorkforceDate    string                                              `json:"entered_workforce_date,omitempty"`    // 参加工作日期
	IDType                  IDType                                              `json:"id_type,omitempty"`                   // 证件类型, 可选值有: `1`：居民身份证, `2`：港澳居民来往内地通行证, `3`：台湾居民来往大陆通行证, `4`：护照, `5`：其他
	IDNumber                string                                              `json:"id_number,omitempty"`                 // 证件号
	HukouType               int                                                 `json:"hukou_type,omitempty"`                // 户口类型, 可选值有: `1`：本市城镇, `2`：外埠城镇, `3`：本市农村, `4`：外埠农村
	HukouLocation           string                                              `json:"hukou_location,omitempty"`            // 户口所在地
	BankAccountNumber       string                                              `json:"bank_account_number,omitempty"`       // 银行卡号
	BankName                string                                              `json:"bank_name,omitempty"`                 // 开户行
	SocialSecurityAccount   string                                              `json:"social_security_account,omitempty"`   // 社保账号
	ProvidentFundAccount    string                                              `json:"provident_fund_account,omitempty"`    // 公积金账号
	EmployeeNo              string                                              `json:"employee_no,omitempty"`               // 工号
	EmployeeType            int                                                 `json:"employee_type,omitempty"`             // 雇员类型, 可选值有: `1`：全职, `2`：实习, `3`：顾问, `4`：外包, `5`：劳务
	Status                  int                                                 `json:"status,omitempty"`                    // 员工状态, 可选值有: `1`：待入职, `2`：在职, `3`：已取消入职, `4`：待离职, `5`：已离职
	HireDate                string                                              `json:"hire_date,omitempty"`                 // 入职日期
	ProbationMonths         float64                                             `json:"probation_months,omitempty"`          // 试用期（月）
	ConversionDate          string                                              `json:"conversion_date,omitempty"`           // 转正日期
	Application             int                                                 `json:"application,omitempty"`               // 转正申请, 可选值有: `1`：未申请, `2`：审批中, `3`：被驳回, `4`：已通过
	ApplicationStatus       int                                                 `json:"application_status,omitempty"`        // 转正状态, 可选值有: `1`：无需转正, `2`：待转正, `3`：已转正
	LastDay                 string                                              `json:"last_day,omitempty"`                  // 离职日期
	DepartureType           int                                                 `json:"departure_type,omitempty"`            // 离职类型, 可选值有: `1`：主动, `2`：被动
	DepartureReason         int                                                 `json:"departure_reason,omitempty"`          // 离职原因, 可选值有: `1`：身体、家庭原因, `2`：职业发展, `3`：薪资福利不满意, `4`：工作压力大, `5`：合同到期不续签, `6`：其他, `7`：无法胜任工作, `8`：组织业务调整和岗位优化, `9`：违反公司条例, `10`：试用期未通过, `11`：其他
	DepartureNotes          string                                              `json:"departure_notes,omitempty"`           // 离职备注
	ContractCompany         *GetEmployeeListRespItemSystemFieldsContractCompany `json:"contract_company,omitempty"`          // 合同公司
	ContractType            int                                                 `json:"contract_type,omitempty"`             // 合同类型, 可选值有: `1`：固定期限劳动合同, `2`：无固定期限劳动合同, `3`：实习协议, `4`：外包协议, `5`：劳务派遣合同, `6`：返聘协议, `7`：其他
	ContractStartDate       string                                              `json:"contract_start_date,omitempty"`       // 合同开始日期
	ContractExpirationDate  string                                              `json:"contract_expiration_date,omitempty"`  // 合同到期日期
	ContractSignTimes       int                                                 `json:"contract_sign_times,omitempty"`       // 劳动合同签订次数
	PersonalEmail           string                                              `json:"personal_email,omitempty"`            // 个人邮箱
	FamilyAddress           string                                              `json:"family_address,omitempty"`            // 家庭地址
	PrimaryEmergencyContact *EHREmergencyContact                                `json:"primary_emergency_contact,omitempty"` // 主要紧急联系人
	EmergencyContact        []*EHREmergencyContact                              `json:"emergency_contact,omitempty"`         // 紧急联系人
	HighestLevelOfEdu       *EHREducation                                       `json:"highest_level_of_edu,omitempty"`      // 最高学历
	Education               []*EHREducation                                     `json:"education,omitempty"`                 // 教育经历
	FormerWorkExp           *EHRWorkExperience                                  `json:"former_work_exp,omitempty"`           // 前工作经历
	WorkExp                 []*EHRWorkExperience                                `json:"work_exp,omitempty"`                  // 工作经历
	IDPhotoPoSide           []*EHRAttachment                                    `json:"id_photo_po_side,omitempty"`          // 身份证照片（人像面）
	IDPhotoEmSide           []*EHRAttachment                                    `json:"id_photo_em_side,omitempty"`          // 身份证照片（国徽面）
	IDPhoto                 []*EHRAttachment                                    `json:"id_photo,omitempty"`                  // 证件照
	DiplomaPhoto            []*EHRAttachment                                    `json:"diploma_photo,omitempty"`             // 学位证书
	GraduationCert          []*EHRAttachment                                    `json:"graduation_cert,omitempty"`           // 毕业证书
	CertOfMerit             []*EHRAttachment                                    `json:"cert_of_merit,omitempty"`             // 奖励证明
	OffboardingFile         []*EHRAttachment                                    `json:"offboarding_file,omitempty"`          // 离职证明
	CancelOnboardingReason  int                                                 `json:"cancel_onboarding_reason,omitempty"`  // 取消入职原因, 可选值有: `1`：个人原因, `2`：原单位留任, `3`：接受其他 Offer, `4`：其他
	CancelOnboardingNotes   string                                              `json:"cancel_onboarding_notes,omitempty"`   // 取消入职备注
	EmployeeFormStatus      int                                                 `json:"employee_form_status,omitempty"`      // 入职登记表状态, 可选值有: `1`：未发送, `2`：待提交, `3`：已提交
	CreateTime              int                                                 `json:"create_time,omitempty"`               // 创建时间
	UpdateTime              int                                                 `json:"update_time,omitempty"`               // 更新时间
}

type GetEmployeeListRespItemSystemFieldsManager struct {
	UserID string `json:"user_id,omitempty"` // 上级的用户 ID（user_id）
	Name   string `json:"name,omitempty"`    // 中文名
	EnName string `json:"en_name,omitempty"` // 英文名
}

type GetEmployeeListRespItemSystemFieldsJob struct {
	ID   int    `json:"id,omitempty"`   // 职位 ID
	Name string `json:"name,omitempty"` // 职位名称
}

type GetEmployeeListRespItemSystemFieldsJobLevel struct {
	ID   int    `json:"id,omitempty"`   // 职级 ID
	Name string `json:"name,omitempty"` // 职级名称
}

type GetEmployeeListRespItemSystemFieldsWorkLocation struct {
	ID   int    `json:"id,omitempty"`   // 工作地点 ID
	Name string `json:"name,omitempty"` // 工作地点名称
}

type GetEmployeeListRespItemSystemFieldsNativeRegion struct {
	IsoCode string `json:"iso_code,omitempty"` // ISO 编码
	Name    string `json:"name,omitempty"`     // 名称
}

type GetEmployeeListRespItemSystemFieldsContractCompany struct {
	ID   int    `json:"id,omitempty"`   // 公司 ID
	Name string `json:"name,omitempty"` // 公司名称
}

type GetEmployeeListRespItemCustomField struct {
	Key   string `json:"key,omitempty"`   // 自定义字段key
	Label string `json:"label,omitempty"` // 自定义字段名称
	Type  string `json:"type,omitempty"`  // 自定义字段类型, 可选值有: `text`：文本类型, `date`：日期类型，如 2020-01-01, `option`：枚举类型, `file`：附件类型
	Value string `json:"value,omitempty"` // 根据 type 不同，结构不同，不同 type 对应的数据结构在 type 的枚举值中有描述
}
