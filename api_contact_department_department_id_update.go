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

// UpdateDepartmentID 此接口可用户更新部门ID(department_id)。新的部门ID(department_id)需要确认在企业内未被占用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/update_department_id
func (r *ContactService) UpdateDepartmentID(ctx context.Context, request *UpdateDepartmentIDReq, options ...MethodOptionFunc) (*UpdateDepartmentIDResp, *Response, error) {
	if r.cli.mock.mockContactUpdateDepartmentID != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#UpdateDepartmentID mock enable")
		return r.cli.mock.mockContactUpdateDepartmentID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateDepartmentID",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/departments/:department_id/update_department_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateDepartmentIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateDepartmentID mock ContactUpdateDepartmentID method
func (r *Mock) MockContactUpdateDepartmentID(f func(ctx context.Context, request *UpdateDepartmentIDReq, options ...MethodOptionFunc) (*UpdateDepartmentIDResp, *Response, error)) {
	r.mockContactUpdateDepartmentID = f
}

// UnMockContactUpdateDepartmentID un-mock ContactUpdateDepartmentID method
func (r *Mock) UnMockContactUpdateDepartmentID() {
	r.mockContactUpdateDepartmentID = nil
}

// UpdateDepartmentIDReq ...
type UpdateDepartmentIDReq struct {
	DepartmentID     string            `path:"department_id" json:"-"`       // 需要更新的部门ID, ID类型与department_id_type中一致, 示例值: "od-d6b83d25c129775723a36f52495c4f81", 最大长度: `64` 字符, 正则校验: `^[a-zA-Z0-9][a-zA-Z0-9_\-@.]{0, 63}$`
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值: open_department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门
	NewDepartmentID  string            `json:"new_department_id,omitempty"`  // 新的部门自定义ID(department_id), 示例值: "NewDevDepartID", 最大长度: `128` 字符, 正则校验: `^0|[^od][A-Za-z0-9]*`
}

// UpdateDepartmentIDResp ...
type UpdateDepartmentIDResp struct {
}

// updateDepartmentIDResp ...
type updateDepartmentIDResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateDepartmentIDResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
