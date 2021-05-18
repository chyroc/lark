// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteDepartment 该接口用于向通讯录中删除部门。
//
// 应用需要同时拥有待删除部门及其父部门的通讯录授权。应用商店应用无权限调用该接口。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/delete
func (r *ContactService) DeleteDepartment(ctx context.Context, request *DeleteDepartmentReq, options ...MethodOptionFunc) (*DeleteDepartmentResp, *Response, error) {
	if r.cli.mock.mockContactDeleteDepartment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteDepartment mock enable")
		return r.cli.mock.mockContactDeleteDepartment(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Contact#DeleteDepartment call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteDepartment request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/departments/:department_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Contact#DeleteDepartment DELETE https://open.feishu.cn/open-apis/contact/v3/departments/:department_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Contact#DeleteDepartment DELETE https://open.feishu.cn/open-apis/contact/v3/departments/:department_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Contact", "DeleteDepartment", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteDepartment success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockContactDeleteDepartment(f func(ctx context.Context, request *DeleteDepartmentReq, options ...MethodOptionFunc) (*DeleteDepartmentResp, *Response, error)) {
	r.mockContactDeleteDepartment = f
}

func (r *Mock) UnMockContactDeleteDepartment() {
	r.mockContactDeleteDepartment = nil
}

type DeleteDepartmentReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	DepartmentID     string            `path:"department_id" json:"-"`       // 部门ID，需要与查询参数中传入的department_id_type类型保持一致。, 示例值："od-4e6ac4d14bcd5071a37a39de902c7141", 最大长度：`128` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`
}

type deleteDepartmentResp struct {
	Code int                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *DeleteDepartmentResp `json:"data,omitempty"`
}

type DeleteDepartmentResp struct{}
