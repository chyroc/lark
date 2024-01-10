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

// DeleteAttendanceGroup 通过班次 ID 删除班次。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/delete
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/group/delete
func (r *AttendanceService) DeleteAttendanceGroup(ctx context.Context, request *DeleteAttendanceGroupReq, options ...MethodOptionFunc) (*DeleteAttendanceGroupResp, *Response, error) {
	if r.cli.mock.mockAttendanceDeleteAttendanceGroup != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Attendance#DeleteAttendanceGroup mock enable")
		return r.cli.mock.mockAttendanceDeleteAttendanceGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "DeleteAttendanceGroup",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/groups/:group_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteAttendanceGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceDeleteAttendanceGroup mock AttendanceDeleteAttendanceGroup method
func (r *Mock) MockAttendanceDeleteAttendanceGroup(f func(ctx context.Context, request *DeleteAttendanceGroupReq, options ...MethodOptionFunc) (*DeleteAttendanceGroupResp, *Response, error)) {
	r.mockAttendanceDeleteAttendanceGroup = f
}

// UnMockAttendanceDeleteAttendanceGroup un-mock AttendanceDeleteAttendanceGroup method
func (r *Mock) UnMockAttendanceDeleteAttendanceGroup() {
	r.mockAttendanceDeleteAttendanceGroup = nil
}

// DeleteAttendanceGroupReq ...
type DeleteAttendanceGroupReq struct {
	GroupID string `path:"group_id" json:"-"` // 考勤组 ID, 获取方式: 1）[创建或修改考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create) 2）[按名称查询考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search) 3）[获取打卡结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query), 示例值: "6919358128597097404"
}

// DeleteAttendanceGroupResp ...
type DeleteAttendanceGroupResp struct {
}

// deleteAttendanceGroupResp ...
type deleteAttendanceGroupResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteAttendanceGroupResp `json:"data,omitempty"`
}
