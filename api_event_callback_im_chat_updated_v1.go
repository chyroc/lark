package lark

import (
	"context"
)

// EventIMChatUpdatedV1
//
// 群组配置被修改后触发此事件，包含：
// - 群主转移
// - 群基本信息修改(群头像/群名称/群描述/群国际化名称)
// - 群权限修改(加人入群权限/群编辑权限/at所有人权限/群分享权限)
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [群配置修改] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/updated
func (r *EventCallbackAPI) HandlerEventIMChatUpdatedV1(f eventIMChatUpdatedV1Handler) {
	r.cli.eventHandler.eventIMChatUpdatedV1Handler = f
}

type eventIMChatUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMChatUpdatedV1) (string, error)

type EventIMChatUpdatedV1 struct {
	ChatID        string                             `json:"chat_id,omitempty"`        // 群组 ID
	OperatorID    *EventIMChatUpdatedV1OperatorID    `json:"operator_id,omitempty"`    // 用户 ID
	AfterChange   *EventIMChatUpdatedV1AfterChange   `json:"after_change,omitempty"`   // 更新后的群信息
	BeforeChange  *EventIMChatUpdatedV1BeforeChange  `json:"before_change,omitempty"`  // 更新前的群信息
	ModeratorList *EventIMChatUpdatedV1ModeratorList `json:"moderator_list,omitempty"` // 群可发言成员名单的变更信息
}

type EventIMChatUpdatedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatUpdatedV1AfterChange struct {
	Avatar                 string                                  `json:"avatar,omitempty"`                   // 群头像
	Name                   string                                  `json:"name,omitempty"`                     // 群名称
	Description            string                                  `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                              `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    string                                  `json:"add_member_permission,omitempty"`    // 加人入群权限(all_members/only_owner/unknown)
	ShareCardPermission    string                                  `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed/unknown)
	AtAllPermission        string                                  `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner/unknown)
	EditPermission         string                                  `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner/unknown)
	MembershipApproval     string                                  `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	JoinMessageVisibility  string                                  `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility string                                  `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	ModerationPermission   string                                  `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner)
	OwnerID                *EventIMChatUpdatedV1AfterChangeOwnerID `json:"owner_id,omitempty"`                 // 用户 ID
}

type EventIMChatUpdatedV1AfterChangeOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatUpdatedV1BeforeChange struct {
	Avatar                 string                                   `json:"avatar,omitempty"`                   // 群头像
	Name                   string                                   `json:"name,omitempty"`                     // 群名称
	Description            string                                   `json:"description,omitempty"`              // 群描述
	I18nNames              *I18nNames                               `json:"i18n_names,omitempty"`               // 群国际化名称
	AddMemberPermission    string                                   `json:"add_member_permission,omitempty"`    // 加人入群权限(all_members/only_owner/unknown)
	ShareCardPermission    string                                   `json:"share_card_permission,omitempty"`    // 群分享权限(allowed/not_allowed/unknown)
	AtAllPermission        string                                   `json:"at_all_permission,omitempty"`        // at 所有人权限(all_members/only_owner/unknown)
	EditPermission         string                                   `json:"edit_permission,omitempty"`          // 群编辑权限(all_members/only_owner/unknown)
	MembershipApproval     string                                   `json:"membership_approval,omitempty"`      // 加群审批(no_approval_required/approval_required)
	JoinMessageVisibility  string                                   `json:"join_message_visibility,omitempty"`  // 入群消息可见性(only_owner/all_members/not_anyone)
	LeaveMessageVisibility string                                   `json:"leave_message_visibility,omitempty"` // 出群消息可见性(only_owner/all_members/not_anyone)
	ModerationPermission   string                                   `json:"moderation_permission,omitempty"`    // 发言权限(all_members/only_owner)
	OwnerID                *EventIMChatUpdatedV1BeforeChangeOwnerID `json:"owner_id,omitempty"`                 // 用户 ID
}

type EventIMChatUpdatedV1BeforeChangeOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatUpdatedV1ModeratorList struct {
	AddedMemberList   []*EventIMChatUpdatedV1ModeratorListAddedMember   `json:"added_member_list,omitempty"`   // 被添加进可发言名单的用户列表（列表中一定会有owner）
	RemovedMemberList []*EventIMChatUpdatedV1ModeratorListRemovedMember `json:"removed_member_list,omitempty"` // 被移除出可发言名单的用户列表
}

type EventIMChatUpdatedV1ModeratorListAddedMember struct {
	UserID *EventIMChatUpdatedV1ModeratorListAddedMemberUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventIMChatUpdatedV1ModeratorListAddedMemberUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatUpdatedV1ModeratorListRemovedMember struct {
	UserID *EventIMChatUpdatedV1ModeratorListRemovedMemberUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventIMChatUpdatedV1ModeratorListRemovedMemberUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
