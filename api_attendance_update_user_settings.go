package lark

import (
	"context"
)

// UpdateUserSettings
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance/rule/user-setting-modify
func (r *AttendanceAPI) UpdateUserSettings(ctx context.Context, request *UpdateUserSettingsReq) (*UpdateUserSettingsResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_settings/modify",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(updateUserSettingsResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Attendance", "UpdateUserSettings", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateUserSettingsReq struct {
	EmployeeType string                            `query:"employee_type" json:"-"` // 用户类型,**可选值有**：,- `employee_id`： 员工id,- `employee_no`： 员工工号
	UserSetting  *UpdateUserSettingsReqUserSetting `json:"user_setting,omitempty"`  // 用户信息
}

type UpdateUserSettingsReqUserSetting struct {
	UserID  string `json:"user_id,omitempty"`  // 用户id
	FaceKey string `json:"face_key,omitempty"` // 人脸照片key（通过文件上传接口得到）
}

type updateUserSettingsResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateUserSettingsResp `json:"data,omitempty"` // \-
}

type UpdateUserSettingsResp struct {
	UserSetting *UpdateUserSettingsRespUserSetting `json:"user_setting,omitempty"` // 用户设置
}

type UpdateUserSettingsRespUserSetting struct {
	UserID  string `json:"user_id,omitempty"`  // 用户id
	FaceKey string `json:"face_key,omitempty"` // 人脸照片key
}
