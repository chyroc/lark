// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// UpdateChatTab 更新会话标签页。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群里
// - 只允许更新类型为`doc`和`url`的会话标签页
// - 更新doc类型时, 操作者（access token对应的身份）需要拥有对应文档的权限
// - 在开启 [仅群主和管理员可管理标签页] 的设置时, 仅群主和群管理员可以更新会话标签页
// - 操作内部群时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-tab/update_tabs
// new doc: https://open.feishu.cn/document/server-docs/group/chat-tab/update_tabs
func (r *ChatService) UpdateChatTab(ctx context.Context, request *UpdateChatTabReq, options ...MethodOptionFunc) (*UpdateChatTabResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChatTab != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChatTab mock enable")
		return r.cli.mock.mockChatUpdateChatTab(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChatTab",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/chat_tabs/update_tabs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatTabResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChatTab mock ChatUpdateChatTab method
func (r *Mock) MockChatUpdateChatTab(f func(ctx context.Context, request *UpdateChatTabReq, options ...MethodOptionFunc) (*UpdateChatTabResp, *Response, error)) {
	r.mockChatUpdateChatTab = f
}

// UnMockChatUpdateChatTab un-mock ChatUpdateChatTab method
func (r *Mock) UnMockChatUpdateChatTab() {
	r.mockChatUpdateChatTab = nil
}

// UpdateChatTabReq ...
type UpdateChatTabReq struct {
	ChatID   string                     `path:"chat_id" json:"-"`    // 群ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 支持群模式为`p2p`与`group`的群ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	ChatTabs []*UpdateChatTabReqChatTab `json:"chat_tabs,omitempty"` // 会话标签页
}

// UpdateChatTabReqChatTab ...
type UpdateChatTabReqChatTab struct {
	TabID      *string                            `json:"tab_id,omitempty"`      // Tab ID, 示例值: "7101214603622940671"
	TabName    *string                            `json:"tab_name,omitempty"`    // Tab名称, 注意: 会话标签页的名称不能超过30个字符, 示例值: "文档"
	TabType    string                             `json:"tab_type,omitempty"`    // Tab类型, 示例值: "doc", 可选值有: message: 消息类型, doc_list: 云文档列表, doc: 文档, pin: Pin, meeting_minute: 会议纪要, chat_announcement: 群公告, url: URL, file: 文件
	TabContent *UpdateChatTabReqChatTabTabContent `json:"tab_content,omitempty"` // Tab内容
	TabConfig  *UpdateChatTabReqChatTabTabConfig  `json:"tab_config,omitempty"`  // Tab的配置
}

// UpdateChatTabReqChatTabTabConfig ...
type UpdateChatTabReqChatTabTabConfig struct {
	IconKey   *string `json:"icon_key,omitempty"`    // 群Tab图标, 示例值: "img_v2_b99741-7628-4abd-aad0-b881e4db83ig"
	IsBuiltIn *bool   `json:"is_built_in,omitempty"` // 群tab是否App内嵌打开, 示例值: false
}

// UpdateChatTabReqChatTabTabContent ...
type UpdateChatTabReqChatTabTabContent struct {
	URL           *string `json:"url,omitempty"`            // URL类型, 示例值: "https://www.feishu.cn"
	Doc           *string `json:"doc,omitempty"`            // Doc链接, 示例值: "https://bytedance.feishu.cn/wiki/wikcnPIcqWjJQwkwDzrB9t40123xz"
	MeetingMinute *string `json:"meeting_minute,omitempty"` // 会议纪要, 示例值: "https://bytedance.feishu.cn/docs/doccnvIXbV22i6hSD3utar4123dx"
}

// UpdateChatTabResp ...
type UpdateChatTabResp struct {
	ChatTabs []*UpdateChatTabRespChatTab `json:"chat_tabs,omitempty"` // 群标签
}

// UpdateChatTabRespChatTab ...
type UpdateChatTabRespChatTab struct {
	TabID      string                              `json:"tab_id,omitempty"`      // Tab ID
	TabName    string                              `json:"tab_name,omitempty"`    // Tab名称, 注意: 会话标签页的名称不能超过30个字符
	TabType    string                              `json:"tab_type,omitempty"`    // Tab类型, 可选值有: message: 消息类型, doc_list: 云文档列表, doc: 文档, pin: Pin, meeting_minute: 会议纪要, chat_announcement: 群公告, url: URL, file: 文件
	TabContent *UpdateChatTabRespChatTabTabContent `json:"tab_content,omitempty"` // Tab内容
	TabConfig  *UpdateChatTabRespChatTabTabConfig  `json:"tab_config,omitempty"`  // Tab的配置
}

// UpdateChatTabRespChatTabTabConfig ...
type UpdateChatTabRespChatTabTabConfig struct {
	IconKey   string `json:"icon_key,omitempty"`    // 群Tab图标
	IsBuiltIn bool   `json:"is_built_in,omitempty"` // 群tab是否App内嵌打开
}

// UpdateChatTabRespChatTabTabContent ...
type UpdateChatTabRespChatTabTabContent struct {
	URL           string `json:"url,omitempty"`            // URL类型
	Doc           string `json:"doc,omitempty"`            // Doc链接
	MeetingMinute string `json:"meeting_minute,omitempty"` // 会议纪要
}

// updateChatTabResp ...
type updateChatTabResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatTabResp `json:"data,omitempty"`
}
