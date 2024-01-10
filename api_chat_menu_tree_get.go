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

// GetChatMenuTree 通过群 ID 获取群内菜单。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
// - 机器人必须在群里。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-menu_tree/get
// new doc: https://open.feishu.cn/document/server-docs/group/chat-menu_tree/get
func (r *ChatService) GetChatMenuTree(ctx context.Context, request *GetChatMenuTreeReq, options ...MethodOptionFunc) (*GetChatMenuTreeResp, *Response, error) {
	if r.cli.mock.mockChatGetChatMenuTree != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Chat#GetChatMenuTree mock enable")
		return r.cli.mock.mockChatGetChatMenuTree(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "GetChatMenuTree",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/menu_tree",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getChatMenuTreeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatGetChatMenuTree mock ChatGetChatMenuTree method
func (r *Mock) MockChatGetChatMenuTree(f func(ctx context.Context, request *GetChatMenuTreeReq, options ...MethodOptionFunc) (*GetChatMenuTreeResp, *Response, error)) {
	r.mockChatGetChatMenuTree = f
}

// UnMockChatGetChatMenuTree un-mock ChatGetChatMenuTree method
func (r *Mock) UnMockChatGetChatMenuTree() {
	r.mockChatGetChatMenuTree = nil
}

// GetChatMenuTreeReq ...
type GetChatMenuTreeReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 仅支持群模式为`group`的群ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
}

// GetChatMenuTreeResp ...
type GetChatMenuTreeResp struct {
	MenuTree *GetChatMenuTreeRespMenuTree `json:"menu_tree,omitempty"` // 群内所有菜单
}

// GetChatMenuTreeRespMenuTree ...
type GetChatMenuTreeRespMenuTree struct {
	ChatMenuTopLevels []*GetChatMenuTreeRespMenuTreeChatMenuTopLevel `json:"chat_menu_top_levels,omitempty"` // 一级菜单列表
}

// GetChatMenuTreeRespMenuTreeChatMenuTopLevel ...
type GetChatMenuTreeRespMenuTreeChatMenuTopLevel struct {
	ChatMenuTopLevelID string                                                   `json:"chat_menu_top_level_id,omitempty"` // 一级菜单ID
	ChatMenuItem       *GetChatMenuTreeRespMenuTreeChatMenuTopLevelChatMenuItem `json:"chat_menu_item,omitempty"`         // 一级菜单信息
	Children           []*GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildren   `json:"children,omitempty"`               // 二级菜单列表
}

// GetChatMenuTreeRespMenuTreeChatMenuTopLevelChatMenuItem ...
type GetChatMenuTreeRespMenuTreeChatMenuTopLevelChatMenuItem struct {
	ActionType   string                                                               `json:"action_type,omitempty"`   // 菜单类型, 注意, 如果一级菜单有二级菜单时, 则此一级菜单的值必须为NONE, 可选值有: NONE: 无类型, REDIRECT_LINK: 跳转链接类型
	RedirectLink *GetChatMenuTreeRespMenuTreeChatMenuTopLevelChatMenuItemRedirectLink `json:"redirect_link,omitempty"` // 跳转链接
	ImageKey     string                                                               `json:"image_key,omitempty"`     // 图片的key值。通过 [上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create) 接口上传message类型图片获取image_key, 注意, 如果一级菜单有二级菜单, 则此一级菜单不能有图标。
	Name         string                                                               `json:"name,omitempty"`          // 菜单名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
	I18nNames    *I18nNames                                                           `json:"i18n_names,omitempty"`    // 菜单国际化名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
}

// GetChatMenuTreeRespMenuTreeChatMenuTopLevelChatMenuItemRedirectLink ...
type GetChatMenuTreeRespMenuTreeChatMenuTopLevelChatMenuItemRedirectLink struct {
	CommonURL  string `json:"common_url,omitempty"`  // 公用跳转链接, 必须以http开头。
	IosURL     string `json:"ios_url,omitempty"`     // IOS端跳转链接, 当该字段不设置时, IOS端会使用common_url。必须以http开头。
	AndroidURL string `json:"android_url,omitempty"` // Android端跳转链接, 当该字段不设置时, Android端会使用common_url。必须以http开头。
	PcURL      string `json:"pc_url,omitempty"`      // PC端跳转链接, 当该字段不设置时, PC端会使用common_url。必须以http开头。在PC端点击群菜单后, 如果需要url对应的页面在飞书侧边栏展开, 可以在url前加上https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=, 比如https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=https://open.feishu.cn/
	WebURL     string `json:"web_url,omitempty"`     // Web端跳转链接, 当该字段不设置时, Web端会使用common_url。必须以http开头。
}

// GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildren ...
type GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildren struct {
	ChatMenuSecondLevelID string                                                           `json:"chat_menu_second_level_id,omitempty"` // 二级菜单ID
	ChatMenuItem          *GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildrenChatMenuItem `json:"chat_menu_item,omitempty"`            // 二级菜单信息
}

// GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildrenChatMenuItem ...
type GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildrenChatMenuItem struct {
	ActionType   string                                                                       `json:"action_type,omitempty"`   // 菜单类型, 注意, 如果一级菜单有二级菜单时, 则此一级菜单的值必须为NONE, 可选值有: NONE: 无类型, REDIRECT_LINK: 跳转链接类型
	RedirectLink *GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildrenChatMenuItemRedirectLink `json:"redirect_link,omitempty"` // 跳转链接
	ImageKey     string                                                                       `json:"image_key,omitempty"`     // 图片的key值。通过 [上传图片](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/image/create) 接口上传message类型图片获取image_key, 注意, 如果一级菜单有二级菜单, 则此一级菜单不能有图标。
	Name         string                                                                       `json:"name,omitempty"`          // 菜单名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
	I18nNames    *I18nNames                                                                   `json:"i18n_names,omitempty"`    // 菜单国际化名称, 注意, 一级、二级菜单名称字符数要在1到120范围内
}

// GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildrenChatMenuItemRedirectLink ...
type GetChatMenuTreeRespMenuTreeChatMenuTopLevelChildrenChatMenuItemRedirectLink struct {
	CommonURL  string `json:"common_url,omitempty"`  // 公用跳转链接, 必须以http开头。
	IosURL     string `json:"ios_url,omitempty"`     // IOS端跳转链接, 当该字段不设置时, IOS端会使用common_url。必须以http开头。
	AndroidURL string `json:"android_url,omitempty"` // Android端跳转链接, 当该字段不设置时, Android端会使用common_url。必须以http开头。
	PcURL      string `json:"pc_url,omitempty"`      // PC端跳转链接, 当该字段不设置时, PC端会使用common_url。必须以http开头。在PC端点击群菜单后, 如果需要url对应的页面在飞书侧边栏展开, 可以在url前加上https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=, 比如https://applink.feishu.cn/client/web_url/open?mode=sidebar-semi&url=https://open.feishu.cn/
	WebURL     string `json:"web_url,omitempty"`     // Web端跳转链接, 当该字段不设置时, Web端会使用common_url。必须以http开头。
}

// getChatMenuTreeResp ...
type getChatMenuTreeResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetChatMenuTreeResp `json:"data,omitempty"`
}
