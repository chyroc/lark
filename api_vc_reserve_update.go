// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateReserve 更新一个预约
//
// 只能更新归属于自己的预约，不需要更新的字段不传（如果传空则会被更新为空）；可用于续期操作，到期时间距离当前时间不超过30天
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/update
func (r *VCService) UpdateReserve(ctx context.Context, request *UpdateReserveReq, options ...MethodOptionFunc) (*UpdateReserveResp, *Response, error) {
	if r.cli.mock.mockVCUpdateReserve != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateReserve mock enable")
		return r.cli.mock.mockVCUpdateReserve(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] VC#UpdateReserve call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateReserve request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "PUT",
		URL:          "https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(updateReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] VC#UpdateReserve PUT https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] VC#UpdateReserve PUT https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("VC", "UpdateReserve", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateReserve success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCUpdateReserve(f func(ctx context.Context, request *UpdateReserveReq, options ...MethodOptionFunc) (*UpdateReserveResp, *Response, error)) {
	r.mockVCUpdateReserve = f
}

func (r *Mock) UnMockVCUpdateReserve() {
	r.mockVCUpdateReserve = nil
}

type UpdateReserveReq struct {
	UserIDType      *IDType                          `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	ReserveID       string                           `path:"reserve_id" json:"-"`        // 预约ID, 示例值："6911188411932033028"
	EndTime         *string                          `json:"end_time,omitempty"`         // 预约到期时间（unix时间，单位sec）, 示例值："1608888867"
	MeetingSettings *UpdateReserveReqMeetingSettings `json:"meeting_settings,omitempty"` // 会议设置
}

type UpdateReserveReqMeetingSettings struct {
	Topic              *string                                            `json:"topic,omitempty"`                // 会议主题, 示例值："my meeting"
	ActionPermissions  []*UpdateReserveReqMeetingSettingsActionPermission `json:"action_permissions,omitempty"`   // 会议权限配置列表，如果存在相同的权限配置项则它们之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
	MeetingInitialType *int                                               `json:"meeting_initial_type,omitempty"` // 会议初始类型, 示例值：1, 可选值有: `1`：多人会议, `2`：1v1呼叫
	CallSetting        *UpdateReserveReqMeetingSettingsCallSetting        `json:"call_setting,omitempty"`         // 1v1呼叫相关参数
}

type UpdateReserveReqMeetingSettingsActionPermission struct {
	Permission         int                                                                 `json:"permission,omitempty"`          // 权限项, 示例值：1, 可选值有: `1`：是否能成为主持人, `2`：是否能邀请参会人, `3`：是否能加入会议
	PermissionCheckers []*UpdateReserveReqMeetingSettingsActionPermissionPermissionChecker `json:"permission_checkers,omitempty"` // 权限检查器列表，权限检查器之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
}

type UpdateReserveReqMeetingSettingsActionPermissionPermissionChecker struct {
	CheckField int      `json:"check_field,omitempty"` // 检查字段类型, 示例值：1, 可选值有: `1`：用户ID, `2`：用户类型, `3`：租户ID
	CheckMode  int      `json:"check_mode,omitempty"`  // 检查方式, 示例值：1, 可选值有: `1`：在check_list中为有权限（白名单）, `2`：不在check_list中为有权限（黑名单）
	CheckList  []string `json:"check_list,omitempty"`  // 检查字段列表
}

type UpdateReserveReqMeetingSettingsCallSetting struct {
	Callee *UpdateReserveReqMeetingSettingsCallSettingCallee `json:"callee,omitempty"` // 被呼叫的用户
}

type UpdateReserveReqMeetingSettingsCallSettingCallee struct {
	ID          *string                                                      `json:"id,omitempty"`            // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType    int                                                          `json:"user_type,omitempty"`     // 用户类型, 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
	PstnSipInfo *UpdateReserveReqMeetingSettingsCallSettingCalleePstnSipInfo `json:"pstn_sip_info,omitempty"` // pstn/sip信息
}

type UpdateReserveReqMeetingSettingsCallSettingCalleePstnSipInfo struct {
	Nickname    *string `json:"nickname,omitempty"`     // 给pstn/sip用户设置的临时昵称, 示例值："dodo"
	MainAddress string  `json:"main_address,omitempty"` // pstn/sip主机号, 示例值："1234"
}

type updateReserveResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateReserveResp `json:"data,omitempty"` //
}

type UpdateReserveResp struct {
	Reserve *UpdateReserveRespReserve `json:"reserve,omitempty"` // 预约数据
}

type UpdateReserveRespReserve struct {
	ID           string `json:"id,omitempty"`            // 预约ID
	MeetingNo    string `json:"meeting_no,omitempty"`    // 9位会议号
	URL          string `json:"url,omitempty"`           // 会议链接
	EndTime      string `json:"end_time,omitempty"`      // 预约到期时间（unix时间，单位sec）
	ExpireStatus int    `json:"expire_status,omitempty"` // 过期状态, 可选值有: `1`：未过期, `2`：已过期
}
