package lark

import (
	"context"
	"io"
)

// UploadImage 上传图片接口，可以上传 JPEG、PNG、WEBP 格式图片
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create
func (r *FileAPI) UploadImage(ctx context.Context, request *UploadImageReq) (*UploadImageResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/images",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                true,
	}
	resp := new(uploadImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("File", "UploadImage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UploadImageReq struct {
	ImageType ImageType `json:"image_type,omitempty"` // 图片类型, 示例值："message", 可选值有: `message`：用于发送消息, `avatar`：用于设置头像
	Image     io.Reader `json:"image,omitempty"`      // 图片内容, 示例值：二进流
}

type uploadImageResp struct {
	Code int              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *UploadImageResp `json:"data,omitempty"` //
}

type UploadImageResp struct {
	ImageKey string `json:"image_key,omitempty"` // 图片的key
}
