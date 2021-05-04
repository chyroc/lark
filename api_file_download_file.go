package lark

import (
	"context"
	"io"
)

// DownloadFile 下载文件接口，只能下载应用自己上传的文件
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 只能下载机器人自己上传的文件
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/file/get
func (r *FileAPI) DownloadFile(ctx context.Context, request *DownloadFileReq) (*DownloadFileResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/files/:file_key",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(downloadFileResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("File", "DownloadFile", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DownloadFileReq struct {
	FileKey string `path:"file_key" json:"-"` // 文件的key,**示例值**："file_456a92d6-c6ea-4de4-ac3f-7afcf44ac78g"
}

type downloadFileResp struct {
	IsFile bool              `json:"is_file,omitempty"`
	Code   int               `json:"code,omitempty"`
	Msg    string            `json:"msg,omitempty"`
	Data   *DownloadFileResp `json:"data,omitempty"`
}

func (r *downloadFileResp) IsFileType() bool {
	return r.IsFile
}

func (r *downloadFileResp) SetFile(file io.Reader) {
	r.Data = &DownloadFileResp{
		File: file,
	}
}

type DownloadFileResp struct {
	File io.Reader `json:"file,omitempty"`
}
