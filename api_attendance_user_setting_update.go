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

// UpdateAttendanceUserSetting 修改授权内员工的用户设置信息, 包括人脸照片文件 ID。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/modify
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_setting/modify
func (r *AttendanceService) UpdateAttendanceUserSetting(ctx context.Context, request *UpdateAttendanceUserSettingReq, options ...MethodOptionFunc) (*UpdateAttendanceUserSettingResp, *Response, error) {
	if r.cli.mock.mockAttendanceUpdateAttendanceUserSetting != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Attendance#UpdateAttendanceUserSetting mock enable")
		return r.cli.mock.mockAttendanceUpdateAttendanceUserSetting(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "UpdateAttendanceUserSetting",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_settings/modify",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateAttendanceUserSettingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceUpdateAttendanceUserSetting mock AttendanceUpdateAttendanceUserSetting method
func (r *Mock) MockAttendanceUpdateAttendanceUserSetting(f func(ctx context.Context, request *UpdateAttendanceUserSettingReq, options ...MethodOptionFunc) (*UpdateAttendanceUserSettingResp, *Response, error)) {
	r.mockAttendanceUpdateAttendanceUserSetting = f
}

// UnMockAttendanceUpdateAttendanceUserSetting un-mock AttendanceUpdateAttendanceUserSetting method
func (r *Mock) UnMockAttendanceUpdateAttendanceUserSetting() {
	r.mockAttendanceUpdateAttendanceUserSetting = nil
}

// UpdateAttendanceUserSettingReq ...
type UpdateAttendanceUserSettingReq struct {
	EmployeeType EmployeeType                               `query:"employee_type" json:"-"` // 请求体和响应体中的 user_id 的员工工号类型, 示例值: employee_id, 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	UserSetting  *UpdateAttendanceUserSettingReqUserSetting `json:"user_setting,omitempty"`  // 用户设置
}

// UpdateAttendanceUserSettingReqUserSetting ...
type UpdateAttendanceUserSettingReqUserSetting struct {
	UserID            string  `json:"user_id,omitempty"`              // 用户 ID, 示例值: "abd754f7"
	FaceKey           string  `json:"face_key,omitempty"`             // 人脸照片文件 ID, 获取方式: [文件上传](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload), 示例值: "xxxxxb306842b1c189bc5212eefxxxxx"
	FaceKeyUpdateTime *string `json:"face_key_update_time,omitempty"` // 人脸照片更新时间, 精确到秒的时间戳, 示例值: "1625681917"
}

// UpdateAttendanceUserSettingResp ...
type UpdateAttendanceUserSettingResp struct {
	UserSetting *UpdateAttendanceUserSettingRespUserSetting `json:"user_setting,omitempty"` // 用户设置
}

// UpdateAttendanceUserSettingRespUserSetting ...
type UpdateAttendanceUserSettingRespUserSetting struct {
	UserID            string `json:"user_id,omitempty"`              // 用户 ID
	FaceKey           string `json:"face_key,omitempty"`             // 人脸照片文件 ID, 获取方式: [文件上传](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload)
	FaceKeyUpdateTime string `json:"face_key_update_time,omitempty"` // 人脸照片更新时间, 精确到秒的时间戳
}

// updateAttendanceUserSettingResp ...
type updateAttendanceUserSettingResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateAttendanceUserSettingResp `json:"data,omitempty"`
}
