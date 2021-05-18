// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateUserApproval
//
// 由于部分企业使用的是自己的审批系统，而不是飞书审批系统，因此员工的请假、加班等数据无法流入到飞书考勤系统中，导致员工在请假时间段内依然收到打卡提醒，并且被记为缺卡。
// 对于这些只使用飞书考勤系统，而未使用飞书审批系统的企业，可以通过考勤开放接口的形式，将三方审批结果数据回写到飞书的考勤系统中。
// 目前支持加班、请假、出差和外出这四种审批结果的写入。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//AddApprovalsInLarkAttendance
func (r *AttendanceService) CreateUserApproval(ctx context.Context, request *CreateUserApprovalReq, options ...MethodOptionFunc) (*CreateUserApprovalResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateUserApproval != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateUserApproval mock enable")
		return r.cli.mock.mockAttendanceCreateUserApproval(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Attendance#CreateUserApproval call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateUserApproval request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_approvals",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createUserApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Attendance#CreateUserApproval POST https://open.feishu.cn/open-apis/attendance/v1/user_approvals failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Attendance#CreateUserApproval POST https://open.feishu.cn/open-apis/attendance/v1/user_approvals failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "CreateUserApproval", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateUserApproval success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceCreateUserApproval(f func(ctx context.Context, request *CreateUserApprovalReq, options ...MethodOptionFunc) (*CreateUserApprovalResp, *Response, error)) {
	r.mockAttendanceCreateUserApproval = f
}

func (r *Mock) UnMockAttendanceCreateUserApproval() {
	r.mockAttendanceCreateUserApproval = nil
}

type CreateUserApprovalReq struct {
	EmployeeType EmployeeType                       `query:"employee_type" json:"-"` // 请求体中的 user_id 的员工工号类型，必选字段，可用值：【employee_id（员工employeeId），employee_no（员工工号）】，示例值："employee_id"
	UserApproval *CreateUserApprovalReqUserApproval `json:"user_approval,omitempty"` // 审批信息
}

type CreateUserApprovalReqUserApproval struct {
	UserID        string                                           `json:"user_id,omitempty"`        // 审批用户
	Date          string                                           `json:"date,omitempty"`           // 审批作用时间
	Outs          []*CreateUserApprovalReqUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*CreateUserApprovalReqUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*CreateUserApprovalReqUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*CreateUserApprovalReqUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

type CreateUserApprovalReqUserApprovalOut struct {
	UniqID        string     `json:"uniq_id,omitempty"`        // 外出类型唯一 ID，代表一种外出类型，长度小于 14
	Unit          int        `json:"unit,omitempty"`           // 外出时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval      int        `json:"interval,omitempty"`       // 假期时长（单位秒）
	StartTime     string     `json:"start_time,omitempty"`     // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime       string     `json:"end_time,omitempty"`       // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames     *I18nNames `json:"i18n_names,omitempty"`     // 外出多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale string     `json:"default_locale,omitempty"` // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果外出名称没有所对应语言的名称，则会使用默认语言的名称
	Reason        string     `json:"reason,omitempty"`         // 外出理由
}

type CreateUserApprovalReqUserApprovalLeave struct {
	UniqID        string     `json:"uniq_id,omitempty"`        // 假期类型唯一 ID，代表一种假期类型，长度小于 14
	Unit          int        `json:"unit,omitempty"`           // 假期时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval      int        `json:"interval,omitempty"`       // 假期时长（单位秒）
	StartTime     string     `json:"start_time,omitempty"`     // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime       string     `json:"end_time,omitempty"`       // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames     *I18nNames `json:"i18n_names,omitempty"`     // 假期多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale string     `json:"default_locale,omitempty"` // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果假期名称没有所对应语言的名称，则会使用默认语言的名称，可用值：【ch（中文），en（英文），ja（日文）】
	Reason        string     `json:"reason,omitempty"`         // 请假理由，必选字段
}

type CreateUserApprovalReqUserApprovalOvertimeWork struct {
	Duration  float64 `json:"duration,omitempty"`   // 加班时长
	Unit      int     `json:"unit,omitempty"`       // 加班时长单位，可用值：【1（天），2（小时）】
	Category  int     `json:"category,omitempty"`   // 加班日期类型，可用值：【1（工作日），2（休息日），3（节假日）】
	Type      int     `json:"type,omitempty"`       // 加班规则类型，可用值：【0（不关联加班规则），1（调休），2（加班费）】
	StartTime string  `json:"start_time,omitempty"` // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime   string  `json:"end_time,omitempty"`   // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type CreateUserApprovalReqUserApprovalTrip struct {
	StartTime        string `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	Reason           string `json:"reason,omitempty"`             // 出差理由
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type createUserApprovalResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *CreateUserApprovalResp `json:"data,omitempty"` // -
}

type CreateUserApprovalResp struct {
	UserApprovals []*CreateUserApprovalRespUserApproval `json:"user_approvals,omitempty"` // 审批结果列表
}

type CreateUserApprovalRespUserApproval struct {
	UserID        string                                            `json:"user_id,omitempty"`        // 审批用户 ID
	Date          string                                            `json:"date,omitempty"`           // 审批作用时间
	Outs          []*CreateUserApprovalRespUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*CreateUserApprovalRespUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*CreateUserApprovalRespUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*CreateUserApprovalRespUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

type CreateUserApprovalRespUserApprovalOut struct {
	UniqID           string     `json:"uniq_id,omitempty"`            // 外出类型唯一 ID，代表一种外出类型，长度小于 14
	Unit             int        `json:"unit,omitempty"`               // 外出时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval         int        `json:"interval,omitempty"`           // 假期时长（单位秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 外出多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果外出名称没有所对应语言的名称，则会使用默认语言的名称
	Reason           string     `json:"reason,omitempty"`             // 外出理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间
}

type CreateUserApprovalRespUserApprovalLeave struct {
	UniqID           string     `json:"uniq_id,omitempty"`            // 假期类型唯一 ID，代表一种假期类型，长度小于 14
	Unit             int        `json:"unit,omitempty"`               // 假期时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval         int        `json:"interval,omitempty"`           // 假期时长（单位秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 假期多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果假期名称没有所对应语言的名称，则会使用默认语言的名称，可用值：【ch（中文），en（英文），ja（日文）】
	Reason           string     `json:"reason,omitempty"`             // 请假理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type CreateUserApprovalRespUserApprovalOvertimeWork struct {
	Duration  float64 `json:"duration,omitempty"`   // 加班时长
	Unit      int     `json:"unit,omitempty"`       // 加班时长单位，可用值：【1（天），2（小时）】
	Category  int     `json:"category,omitempty"`   // 加班日期类型，可用值：【1（工作日），2（休息日），3（节假日）】
	Type      int     `json:"type,omitempty"`       // 加班规则类型，可用值：【0（不关联加班规则），1（调休），2（加班费），3（关联加班规则，没有调休或加班费）】
	StartTime string  `json:"start_time,omitempty"` // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime   string  `json:"end_time,omitempty"`   // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type CreateUserApprovalRespUserApprovalTrip struct {
	StartTime        string `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	Reason           string `json:"reason,omitempty"`             // 出差理由
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}
