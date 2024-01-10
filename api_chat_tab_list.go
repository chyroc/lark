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

// GetChatTabList 拉取会话标签页。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群里
// - 操作内部群时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-tab/list_tabs
// new doc: https://open.feishu.cn/document/server-docs/group/chat-tab/list_tabs
func (r *ChatService) GetChatTabList(ctx context.Context, request *GetChatTabListReq, options ...MethodOptionFunc) (*GetChatTabListResp, *Response, error) {
	if r.cli.mock.mockChatGetChatTabList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Chat#GetChatTabList mock enable")
		return r.cli.mock.mockChatGetChatTabList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChatTabList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/chat_tabs/list_tabs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getChatTabListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatGetChatTabList mock ChatGetChatTabList method
func (r *Mock) MockChatGetChatTabList(f func(ctx context.Context, request *GetChatTabListReq, options ...MethodOptionFunc) (*GetChatTabListResp, *Response, error)) {
	r.mockChatGetChatTabList = f
}

// UnMockChatGetChatTabList un-mock ChatGetChatTabList method
func (r *Mock) UnMockChatGetChatTabList() {
	r.mockChatGetChatTabList = nil
}

// GetChatTabListReq ...
type GetChatTabListReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 支持群模式为`p2p`与`group`的群ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
}

// GetChatTabListResp ...
type GetChatTabListResp struct {
	ChatTabs []*GetChatTabListRespChatTab `json:"chat_tabs,omitempty"` // 会话标签页
}

// GetChatTabListRespChatTab ...
type GetChatTabListRespChatTab struct {
	TabID      string                               `json:"tab_id,omitempty"`      // Tab ID
	TabName    string                               `json:"tab_name,omitempty"`    // Tab名称, 注意: 会话标签页的名称不能超过30个字符
	TabType    string                               `json:"tab_type,omitempty"`    // Tab类型, 可选值有: message: 消息类型, doc_list: 云文档列表, doc: 文档, pin: Pin, meeting_minute: 会议纪要, chat_announcement: 群公告, url: URL, file: 文件
	TabContent *GetChatTabListRespChatTabTabContent `json:"tab_content,omitempty"` // Tab内容
	TabConfig  *GetChatTabListRespChatTabTabConfig  `json:"tab_config,omitempty"`  // Tab的配置
}

// GetChatTabListRespChatTabTabConfig ...
type GetChatTabListRespChatTabTabConfig struct {
	IconKey   string `json:"icon_key,omitempty"`    // 群Tab图标
	IsBuiltIn bool   `json:"is_built_in,omitempty"` // 群tab是否App内嵌打开
}

// GetChatTabListRespChatTabTabContent ...
type GetChatTabListRespChatTabTabContent struct {
	URL           string `json:"url,omitempty"`            // URL类型
	Doc           string `json:"doc,omitempty"`            // Doc链接
	MeetingMinute string `json:"meeting_minute,omitempty"` // 会议纪要
}

// getChatTabListResp ...
type getChatTabListResp struct {
	Code int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 错误描述
	Data *GetChatTabListResp `json:"data,omitempty"`
}
