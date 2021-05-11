// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// UploadAttendanceFile
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//rule/file_upload
func (r *AttendanceAPI) UploadAttendanceFile(ctx context.Context, request *UploadAttendanceFileReq, options ...MethodOptionFunc) (*UploadAttendanceFileResp, *Response, error) {
	if r.cli.mock.mockAttendanceUploadAttendanceFile != nil {
		return r.cli.mock.mockAttendanceUploadAttendanceFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/files/upload",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(uploadAttendanceFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Attendance", "UploadAttendanceFile", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceUploadAttendanceFile(f func(ctx context.Context, request *UploadAttendanceFileReq, options ...MethodOptionFunc) (*UploadAttendanceFileResp, *Response, error)) {
	r.mockAttendanceUploadAttendanceFile = f
}

func (r *Mock) UnMockAttendanceUploadAttendanceFile() {
	r.mockAttendanceUploadAttendanceFile = nil
}

type UploadAttendanceFileReq struct {
	FileName string    `query:"file_name" json:"-"` // 文件名
	File     io.Reader `json:"file,omitempty"`      // 文件
}

type uploadAttendanceFileResp struct {
	Code int                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *UploadAttendanceFileResp `json:"data,omitempty"` //
}

type UploadAttendanceFileResp struct {
	File *UploadAttendanceFileRespFile `json:"file,omitempty"` // 文件
}

type UploadAttendanceFileRespFile struct {
	FileID string `json:"file_id,omitempty"` // 文件id
}