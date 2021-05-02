package lark

import (
	"context"
)

// CreateChat 创建群并设置群头像、群名、群描述等。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/create
func (r *ChatAPI) CreateChat(ctx context.Context, request *CreateChatReq) (*CreateChatResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
	}
	resp := new(createChatResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "CreateChat", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateChatReq struct {
	Avatar                 string                  `json:"avatar,omitempty"`                   // 群头像对应的 Image Key，可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意：上传图片的 ==image_type== 需要指定为 ==avatar==）,**示例值**："default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   string                  `json:"name,omitempty"`                     // 群名称,**示例值**："测试群名称"
	Description            string                  `json:"description,omitempty"`              // 群描述,**示例值**："测试群描述"
	I18nNames              *CreateChatReqI18nNames `json:"i18n_names,omitempty"`               // 群国际化名称
	ChatMode               string                  `json:"chat_mode,omitempty"`                // 群模式(group),**示例值**："group"
	ChatType               string                  `json:"chat_type,omitempty"`                // 群类型(private/public),**示例值**："private"
	JoinMessageVisibility  string                  `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone),**示例值**："all_members"
	LeaveMessageVisibility string                  `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone),**示例值**："all_members"
	MembershipApproval     string                  `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required),**示例值**："no_approval_required"
}

type CreateChatReqI18nNames struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名,**示例值**："群聊"
	EnUs string `json:"en_us,omitempty"` // 英文名,**示例值**："group chat"
	JaJp string `json:"ja_jp,omitempty"` // 日文名,**示例值**："グループチャット"
}

type createChatResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateChatResp `json:"data,omitempty"` // \-
}

type CreateChatResp struct {
	ChatID                 string                   `json:"chat_id,omitempty"`                  // 群 ID
	Avatar                 string                   `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string                   `json:"name,omitempty"`                     // 群名称
	Description            string                   `json:"description,omitempty"`              // 群描述
	I18nNames              *CreateChatRespI18nNames `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    string                   `json:"add_member_permission,omitempty"`    // 加 user/bot 入群权限(all_members/only_owner)
	ShareCardPermission    string                   `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed)
	AtAllPermission        string                   `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner)
	EditPermission         string                   `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner)
	ChatMode               string                   `json:"chat_mode,omitempty"`                // 群模式(group)
	ChatType               string                   `json:"chat_type,omitempty"`                // 群类型(private/public)
	ChatTag                string                   `json:"chat_tag,omitempty"`                 // 优先级最高的一个群 tag（inner/tenant/department/edu/meeting/customer_service）
	JoinMessageVisibility  string                   `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility string                   `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	MembershipApproval     string                   `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	ModerationPermission   string                   `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner/moderator_list)
}

type CreateChatRespI18nNames struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文名
	EnUs string `json:"en_us,omitempty"` // 英文名
	JaJp string `json:"ja_jp,omitempty"` // 日文名
}
