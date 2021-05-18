// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateFolder 该接口用于根据 folderToken 在该 folder 下创建文件夹。
//
// 该接口不支持并发创建，且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukTNzUjL5UzM14SO1MTN
func (r *DriveService) CreateFolder(ctx context.Context, request *CreateFolderReq, options ...MethodOptionFunc) (*CreateFolderResp, *Response, error) {
	if r.cli.mock.mockDriveCreateFolder != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateFolder mock enable")
		return r.cli.mock.mockDriveCreateFolder(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#CreateFolder call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateFolder request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/drive/explorer/v2/folder/{folderToken}",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(createFolderResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#CreateFolder POST https://open.feishu.cn/open-apis/drive/explorer/v2/folder/{folderToken} failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#CreateFolder POST https://open.feishu.cn/open-apis/drive/explorer/v2/folder/{folderToken} failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "CreateFolder", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateFolder success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveCreateFolder(f func(ctx context.Context, request *CreateFolderReq, options ...MethodOptionFunc) (*CreateFolderResp, *Response, error)) {
	r.mockDriveCreateFolder = f
}

func (r *Mock) UnMockDriveCreateFolder() {
	r.mockDriveCreateFolder = nil
}

type CreateFolderReq struct {
	FolderToken string `path:"folderToken" json:"-"` // 文件夹的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 3 项
	Title       string `json:"title,omitempty"`      // 文件夹标题
}

type createFolderResp struct {
	Code int               `json:"code,omitempty"`
	Msg  string            `json:"msg,omitempty"`
	Data *CreateFolderResp `json:"data,omitempty"`
}

type CreateFolderResp struct {
	URL      string `json:"url,omitempty"`      // 新创建文件夹的 url
	Revision int    `json:"revision,omitempty"` // 新创建文件夹的版本号
	Token    string `json:"token,omitempty"`    // 新创建文件夹的 token
}
