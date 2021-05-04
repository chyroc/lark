package lark

import (
	"context"
	"io"
)

// DownloadImage 下载图片资源，只能下载应用自己上传且图片类型为message的图片
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 只能下载机器人自己上传且图片类型为message的图片，avatar类型暂不支持下载；
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/get
func (r *FileAPI) DownloadImage(ctx context.Context, request *DownloadImageReq) (*DownloadImageResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/images/:image_key",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(downloadImageResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("File", "DownloadImage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DownloadImageReq struct {
	ImageKey string `path:"image_key" json:"-"` // 图片的key,**示例值**："img_8d5181ca-0aed-40f0-b0d1-b1452132afbg"
}

type downloadImageResp struct {
	IsFile bool               `json:"is_file,omitempty"`
	Code   int                `json:"code,omitempty"`
	Msg    string             `json:"msg,omitempty"`
	Data   *DownloadImageResp `json:"data,omitempty"`
}

func (r *downloadImageResp) IsFileType() bool {
	return r.IsFile
}

func (r *downloadImageResp) SetFile(file io.Reader) {
	r.Data = &DownloadImageResp{
		File: file,
	}
}

type DownloadImageResp struct {
	File io.Reader `json:"file,omitempty"`
}
