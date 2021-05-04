package lark

import (
	"context"
)

// GetChat 获取群名称、群描述、群头像、群主 ID 等群基本信息。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 机器人或授权用户必须在群里
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get
func (r *ChatAPI) GetChat(ctx context.Context, request *GetChatReq) (*GetChatResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		IsFile:                false,
	}
	resp := new(getChatResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "GetChat", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetChatReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	ChatID     string  `path:"chat_id" json:"-"`       // 群 ID,**示例值**："oc_a0553eda9014c201e6969b478895c230"
}

type getChatResp struct {
	Code int          `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string       `json:"msg,omitempty"`  // 错误描述
	Data *GetChatResp `json:"data,omitempty"` // \-
}

type GetChatResp struct {
	Avatar                 string     `json:"avatar,omitempty"`                   // 群头像 URL
	Name                   string     `json:"name,omitempty"`                     // 群名称
	Description            string     `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    string     `json:"add_member_permission,omitempty"`    // 群成员添加权限(all_members/only_owner)
	ShareCardPermission    string     `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed)
	AtAllPermission        string     `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner)
	EditPermission         string     `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner)
	OwnerIDType            IDType     `json:"owner_id_type,omitempty"`            // 群主 ID 的类型(open_id/user_id/union_id)，群主是机器人时，不返回该字段。
	OwnerID                string     `json:"owner_id,omitempty"`                 // 群主 ID，群主是机器人时，不返回该字段。
	ChatMode               string     `json:"chat_mode,omitempty"`                // 群模式(group/topic/p2p)
	ChatType               string     `json:"chat_type,omitempty"`                // 群类型(private/public)
	ChatTag                string     `json:"chat_tag,omitempty"`                 // 优先级最高的一个群tag(inner/tenant/department/edu/meeting/customer_service)
	JoinMessageVisibility  string     `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility string     `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	MembershipApproval     string     `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	ModerationPermission   string     `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner/moderator_list)
}
