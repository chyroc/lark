package lark

import (
	"context"
)

// UpdateChat 更新群头像、群名称、群描述、群配置、转让群主等。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 若群未开启==仅群主可编辑群信息== 配置：
// - 群主 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可更新所有信息
// - 不满足上述条件的群成员或机器人，仅可更新群头像、群名称、群描述、群国际化名称信息
// - 若群开启了==仅群主可编辑群信息==配置：
// - 群主 或 创建群组且具备==更新应用所创建群的群信息==权限的机器人，可更新所有信息
// - 不满足上述条件的群成员或者机器人，任何群信息都不能修改
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/update
func (r *ChatAPI) UpdateChat(ctx context.Context, request *UpdateChatReq) (*UpdateChatResp, *Response, error) {
	req := &requestParam{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
	}
	resp := new(updateChatResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "UpdateChat", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateChatReq struct {
	UserIDType             *IDType    `query:"user_id_type" json:"-"`             // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	ChatID                 string     `path:"chat_id" json:"-"`                   // 群 ID,**示例值**："oc_a0553eda9014c201e6969b478895c230"
	Avatar                 *string    `json:"avatar,omitempty"`                   // 群头像对应的 Image Key，可通过[上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create)获取（注意：上传图片的 ==image_type== 需要指定为 ==avatar==）,**示例值**："default-avatar_44ae0ca3-e140-494b-956f-78091e348435"
	Name                   *string    `json:"name,omitempty"`                     // 群名称,**示例值**："测试群名称"
	Description            *string    `json:"description,omitempty"`              // 群描述,**示例值**："测试群描述"
	I18nNames              *I18nNames `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    *string    `json:"add_member_permission,omitempty"`    // 加 user/bot 入群权限(all_members/only_owner),**示例值**："all_members"
	ShareCardPermission    *string    `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed),**示例值**："allowed"
	AtAllPermission        *string    `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner),**示例值**："all_members"
	EditPermission         *string    `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner),**示例值**："all_members"
	OwnerID                *string    `json:"owner_id,omitempty"`                 // 新群主 ID,**示例值**："4d7a3c6g"
	JoinMessageVisibility  *string    `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone),**示例值**："only_owner"
	LeaveMessageVisibility *string    `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone),**示例值**："only_owner"
	MembershipApproval     *string    `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required),**示例值**："no_approval_required"
}

type updateChatResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatResp `json:"data,omitempty"`
}

type UpdateChatResp struct{}
