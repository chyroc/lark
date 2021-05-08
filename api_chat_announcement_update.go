package lark

import (
	"context"
)

// UpdateAnnouncement 更新会话中的群公告信息，更新公告信息的格式和更新[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 当授权用户或机器人是群主时，可更新群公告信息
// - 当授权用户或机器人非群主，但群主未设置 [仅群主可编辑群信息] 时，可更新群公告信息
// - 当授权用户或机器人非群主，但群主设置了 [仅群主可编辑群信息] 时，无法更新公告信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/patch
func (r *ChatAPI) UpdateAnnouncement(ctx context.Context, request *UpdateAnnouncementReq) (*UpdateAnnouncementResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateAnnouncementResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Chat", "UpdateAnnouncement", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateAnnouncementReq struct {
	ChatID   string   `path:"chat_id" json:"-"`   // 待修改公告的群 ID, 示例值："oc_5ad11d72b830411d72b836c20"
	Revision string   `json:"revision,omitempty"` // 文档当前版本号 int64 类型，get 接口会返回, 示例值："12"
	Requests []string `json:"requests,omitempty"` // 修改文档请求的序列化字段
}

type updateAnnouncementResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateAnnouncementResp `json:"data,omitempty"`
}

type UpdateAnnouncementResp struct{}
