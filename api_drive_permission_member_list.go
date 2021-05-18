// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetMemberPermissionList 该接口用于根据 filetoken 查询协作者，目前包括人("user")和群("chat") 。
//
// 你能获取到协作者列表的前提是你对该文档有分享权限
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATN3UjLwUzN14CM1cTN
func (r *DriveService) GetMemberPermissionList(ctx context.Context, request *GetMemberPermissionListReq, options ...MethodOptionFunc) (*GetMemberPermissionListResp, *Response, error) {
	if r.cli.mock.mockDriveGetMemberPermissionList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetMemberPermissionList mock enable")
		return r.cli.mock.mockDriveGetMemberPermissionList(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#GetMemberPermissionList call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetMemberPermissionList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/member/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(getMemberPermissionListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#GetMemberPermissionList POST https://open.feishu.cn/open-apis/drive/permission/member/list failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#GetMemberPermissionList POST https://open.feishu.cn/open-apis/drive/permission/member/list failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "GetMemberPermissionList", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetMemberPermissionList success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveGetMemberPermissionList(f func(ctx context.Context, request *GetMemberPermissionListReq, options ...MethodOptionFunc) (*GetMemberPermissionListResp, *Response, error)) {
	r.mockDriveGetMemberPermissionList = f
}

func (r *Mock) UnMockDriveGetMemberPermissionList() {
	r.mockDriveGetMemberPermissionList = nil
}

type GetMemberPermissionListReq struct {
	Token string `json:"token,omitempty"` // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type  string `json:"type,omitempty"`  // 文档类型  "doc"  or  "sheet" or "file"
}

type getMemberPermissionListResp struct {
	Code int                          `json:"code,omitempty"`
	Msg  string                       `json:"msg,omitempty"`
	Data *GetMemberPermissionListResp `json:"data,omitempty"`
}

type GetMemberPermissionListResp struct {
	Members *GetMemberPermissionListRespMembers `json:"members,omitempty"` // 协作者列表
}

type GetMemberPermissionListRespMembers struct {
	MemberType   string `json:"member_type,omitempty"`    // 协作者类型 "user" or "chat"
	MemberOpenID string `json:"member_open_id,omitempty"` // 协作者openid
	MemberUserID string `json:"member_user_id,omitempty"` // 协作者userid(仅当member_type="user"时有效)
	Perm         string `json:"perm,omitempty"`           // 协作者权限 (注意: **有"edit"权限的协作者一定有"view"权限**)
}
