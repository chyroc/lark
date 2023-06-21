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

// UnbindContactUnitDepartment 通过该接口解除部门与单位的绑定关系, 需更新单位的权限, 需对应部门的通讯录权限。由于单位是旗舰版付费功能, 企业需开通相关功能, 否则会解绑失败。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/unit/unbind_department
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/unit/unbind_department
func (r *ContactService) UnbindContactUnitDepartment(ctx context.Context, request *UnbindContactUnitDepartmentReq, options ...MethodOptionFunc) (*UnbindContactUnitDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactUnbindContactUnitDepartment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UnbindContactUnitDepartment mock enable")
		return r.cli.mock.mockContactUnbindContactUnitDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UnbindContactUnitDepartment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/unit/unbind_department",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(unbindContactUnitDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUnbindContactUnitDepartment mock ContactUnbindContactUnitDepartment method
func (r *Mock) MockContactUnbindContactUnitDepartment(f func(ctx context.Context, request *UnbindContactUnitDepartmentReq, options ...MethodOptionFunc) (*UnbindContactUnitDepartmentResp, *Response, error)) {
	r.mockContactUnbindContactUnitDepartment = f
}

// UnMockContactUnbindContactUnitDepartment un-mock ContactUnbindContactUnitDepartment method
func (r *Mock) UnMockContactUnbindContactUnitDepartment() {
	r.mockContactUnbindContactUnitDepartment = nil
}

// UnbindContactUnitDepartmentReq ...
type UnbindContactUnitDepartmentReq struct {
	UnitID           string            `json:"unit_id,omitempty"`            // 单位ID, 示例值: "BU121"
	DepartmentID     string            `json:"department_id,omitempty"`      // 预解除关联的部门ID, 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	DepartmentIDType *DepartmentIDType `json:"department_id_type,omitempty"` // 此次调用中使用的部门ID的类型, 示例值: "open_department_id", 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
}

// UnbindContactUnitDepartmentResp ...
type UnbindContactUnitDepartmentResp struct {
}

// unbindContactUnitDepartmentResp ...
type unbindContactUnitDepartmentResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UnbindContactUnitDepartmentResp `json:"data,omitempty"`
}
