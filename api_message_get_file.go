package lark

import (
	"context"
	"io"
)

// GetMessageFile 获取消息中的资源文件，包括音频，视频，图片和文件。当前仅支持 100M 以内的资源文件的下载。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 机器人和消息需要在同一会话中
// - 请求的 file_key 和 message_id 需要匹配
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-resource/get
func (r *MessageAPI) GetMessageFile(ctx context.Context, request *GetMessageFileReq) (*GetMessageFileResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id/resources/:file_key",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMessageFileResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "GetMessageFile", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMessageFileReq struct {
	Type      string `query:"type" json:"-"`      // 资源类型，可选"image, file“； image对应消息中的 图片资源。  file对应消息中的 文件、音频、视频、表情包 资源,**示例值**："image,"
	MessageID string `path:"message_id" json:"-"` // 待查询资源对应的消息ID,**示例值**："om_dc13264520392913993dd051dba21dcf"
	FileKey   string `path:"file_key" json:"-"`   // 待查询资源的key,**示例值**："file_456a92d6-c6ea-4de4-ac3f-7afcf44ac78g"
}

type getMessageFileResp struct {
	IsFile bool                `json:"is_file,omitempty"`
	Code   int                 `json:"code,omitempty"`
	Msg    string              `json:"msg,omitempty"`
	Data   *GetMessageFileResp `json:"data,omitempty"`
}

func (r *getMessageFileResp) IsFileType() bool {
	return r.IsFile
}

func (r *getMessageFileResp) SetFile(file io.Reader) {
	r.Data = &GetMessageFileResp{
		File: file,
	}
}

type GetMessageFileResp struct {
	File io.Reader `json:"file,omitempty"`
}
