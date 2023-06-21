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

// GetApplicationVersion 根据应用 ID 和应用版本 ID 来获取同租户下的应用版本的信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_version/get
// new doc: https://open.feishu.cn/document/server-docs/application-v6/application/get-2
func (r *ApplicationService) GetApplicationVersion(ctx context.Context, request *GetApplicationVersionReq, options ...MethodOptionFunc) (*GetApplicationVersionResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationVersion != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Application#GetApplicationVersion mock enable")
		return r.cli.mock.mockApplicationGetApplicationVersion(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationVersion",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/app_versions/:version_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationVersionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationVersion mock ApplicationGetApplicationVersion method
func (r *Mock) MockApplicationGetApplicationVersion(f func(ctx context.Context, request *GetApplicationVersionReq, options ...MethodOptionFunc) (*GetApplicationVersionResp, *Response, error)) {
	r.mockApplicationGetApplicationVersion = f
}

// UnMockApplicationGetApplicationVersion un-mock ApplicationGetApplicationVersion method
func (r *Mock) UnMockApplicationGetApplicationVersion() {
	r.mockApplicationGetApplicationVersion = nil
}

// GetApplicationVersionReq ...
type GetApplicationVersionReq struct {
	AppID      string  `path:"app_id" json:"-"`        // 应用的 app_id, 需要查询其他应用版本信息时, 必须申请[获取应用版本信息](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)权限, 仅查询本应用版本信息时, 可填入 "me" 或者应用自身 app_id, 示例值: "cli_9f3ca975326b501b"
	VersionID  string  `path:"version_id" json:"-"`    // 唯一标识应用版本的 ID, 示例值: "oav_d317f090b7258ad0372aa53963cda70d"
	Lang       string  `query:"lang" json:"-"`         // 应用信息的语言版本, 示例值: "zh_cn", 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetApplicationVersionResp ...
type GetApplicationVersionResp struct {
	AppVersion *GetApplicationVersionRespAppVersion `json:"app_version,omitempty"` // 应用版本信息
}

// GetApplicationVersionRespAppVersion ...
type GetApplicationVersionRespAppVersion struct {
	AppID            string                                      `json:"app_id,omitempty"`            // 应用 id
	Version          string                                      `json:"version,omitempty"`           // 在开发者后台填入的应用版本号
	VersionID        string                                      `json:"version_id,omitempty"`        // 唯一标识应用版本的 ID
	AppName          string                                      `json:"app_name,omitempty"`          // 应用默认名称
	AvatarURL        string                                      `json:"avatar_url,omitempty"`        // 应用头像 url
	Description      string                                      `json:"description,omitempty"`       // 应用默认描述
	Scopes           []*GetApplicationVersionRespAppVersionScope `json:"scopes,omitempty"`            // 应用权限列表
	BackHomeURL      string                                      `json:"back_home_url,omitempty"`     // 后台主页地址
	I18n             []*GetApplicationVersionRespAppVersionI18n  `json:"i18n,omitempty"`              // 应用的国际化信息列表
	CommonCategories []string                                    `json:"common_categories,omitempty"` // 应用分类的国际化描述
	Events           []string                                    `json:"events,omitempty"`            // 应用已订阅开放平台事件列表
	Status           int64                                       `json:"status,omitempty"`            // 版本状态, 可选值有: 0: 未知状态, 1: 审核通过, 2: 审核拒绝, 3: 审核中, 4: 未提交审核
	CreateTime       string                                      `json:"create_time,omitempty"`       // 版本创建时间（单位: s）
	PublishTime      string                                      `json:"publish_time,omitempty"`      // 版本发布时间（单位: s）
	Ability          *GetApplicationVersionRespAppVersionAbility `json:"ability,omitempty"`           // 当前版本下应用开启的能力
	Remark           *GetApplicationVersionRespAppVersionRemark  `json:"remark,omitempty"`            // 跟随应用版本的信息
}

// GetApplicationVersionRespAppVersionAbility ...
type GetApplicationVersionRespAppVersionAbility struct {
	Gadget           *GetApplicationVersionRespAppVersionAbilityGadget            `json:"gadget,omitempty"`            // 小程序能力
	WebApp           *GetApplicationVersionRespAppVersionAbilityWebApp            `json:"web_app,omitempty"`           // 网页能力
	Bot              *GetApplicationVersionRespAppVersionAbilityBot               `json:"bot,omitempty"`               // 机器人能力
	WorkplaceWidgets []*GetApplicationVersionRespAppVersionAbilityWorkplaceWidget `json:"workplace_widgets,omitempty"` // 小组件能力
	Navigate         *GetApplicationVersionRespAppVersionAbilityNavigate          `json:"navigate,omitempty"`          // 主导航小程序
	CloudDoc         *GetApplicationVersionRespAppVersionAbilityCloudDoc          `json:"cloud_doc,omitempty"`         // 云文档应用
	DocsBlocks       []*GetApplicationVersionRespAppVersionAbilityDocsBlock       `json:"docs_blocks,omitempty"`       // 云文档小组件
	MessageAction    *GetApplicationVersionRespAppVersionAbilityMessageAction     `json:"message_action,omitempty"`    // 消息快捷操作
	PlusMenu         *GetApplicationVersionRespAppVersionAbilityPlusMenu          `json:"plus_menu,omitempty"`         // 加号菜单
}

// GetApplicationVersionRespAppVersionAbilityBot ...
type GetApplicationVersionRespAppVersionAbilityBot struct {
	CardRequestURL string `json:"card_request_url,omitempty"` // 消息卡片回调地址
}

// GetApplicationVersionRespAppVersionAbilityCloudDoc ...
type GetApplicationVersionRespAppVersionAbilityCloudDoc struct {
	SpaceURL string                                                    `json:"space_url,omitempty"` // 云空间重定向 url
	I18n     []*GetApplicationVersionRespAppVersionAbilityCloudDocI18n `json:"i18n,omitempty"`      // 国际化信息
	IconURL  string                                                    `json:"icon_url,omitempty"`  // 图标链接
	Mode     int64                                                     `json:"mode,omitempty"`      // 云文档支持模式, 可选值有: 0: 未知, 1: 移动端
}

// GetApplicationVersionRespAppVersionAbilityCloudDocI18n ...
type GetApplicationVersionRespAppVersionAbilityCloudDocI18n struct {
	I18nKey          string `json:"i18n_key,omitempty"`          // 国际化语言的 key, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	Name             string `json:"name,omitempty"`              // 云文档国际化名称
	ReadDescription  string `json:"read_description,omitempty"`  // 云文档国际化读权限说明
	WriteDescription string `json:"write_description,omitempty"` // 云文档国际化写权限说明
}

// GetApplicationVersionRespAppVersionAbilityDocsBlock ...
type GetApplicationVersionRespAppVersionAbilityDocsBlock struct {
	BlockTypeID   string                                                     `json:"block_type_id,omitempty"`   // BlockTypeID
	I18n          []*GetApplicationVersionRespAppVersionAbilityDocsBlockI18n `json:"i18n,omitempty"`            // block 的国际化信息
	MobileIconURL string                                                     `json:"mobile_icon_url,omitempty"` // 移动端 icon 链接
	PcIconURL     string                                                     `json:"pc_icon_url,omitempty"`     // pc 端口 icon 链接
}

// GetApplicationVersionRespAppVersionAbilityDocsBlockI18n ...
type GetApplicationVersionRespAppVersionAbilityDocsBlockI18n struct {
	I18nKey string `json:"i18n_key,omitempty"` // 国际化语言的 key, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	Name    string `json:"name,omitempty"`     // 名称
}

// GetApplicationVersionRespAppVersionAbilityGadget ...
type GetApplicationVersionRespAppVersionAbilityGadget struct {
	EnablePcMode         int64    `json:"enable_pc_mode,omitempty"`          // pc 支持的小程序模式, bit 位表示, 可选值有: 1: sidebar 模式, 2: pc 模式, 4: 主导航模式
	SchemaURLs           []string `json:"schema_urls,omitempty"`             // schema url 列表
	PcUseMobilePkg       bool     `json:"pc_use_mobile_pkg,omitempty"`       // pc 端是否使用小程序版本
	PcVersion            string   `json:"pc_version,omitempty"`              // pc 的小程序版本号
	MobileVersion        string   `json:"mobile_version,omitempty"`          // 移动端小程序版本号
	MobileMinLarkVersion string   `json:"mobile_min_lark_version,omitempty"` // 移动端兼容的最低飞书版本
	PcMinLarkVersion     string   `json:"pc_min_lark_version,omitempty"`     // pc 端兼容的最低飞书版本
}

// GetApplicationVersionRespAppVersionAbilityMessageAction ...
type GetApplicationVersionRespAppVersionAbilityMessageAction struct {
	PcAppLink     string                                                         `json:"pc_app_link,omitempty"`     // pc 端链接
	MobileAppLink string                                                         `json:"mobile_app_link,omitempty"` // 移动端链接
	I18n          []*GetApplicationVersionRespAppVersionAbilityMessageActionI18n `json:"i18n,omitempty"`            // 国际化信息
}

// GetApplicationVersionRespAppVersionAbilityMessageActionI18n ...
type GetApplicationVersionRespAppVersionAbilityMessageActionI18n struct {
	I18nKey string `json:"i18n_key,omitempty"` // 国际化语言的 key, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	Name    string `json:"name,omitempty"`     // 国际化名称
}

// GetApplicationVersionRespAppVersionAbilityNavigate ...
type GetApplicationVersionRespAppVersionAbilityNavigate struct {
	Pc     *GetApplicationVersionRespAppVersionAbilityNavigatePc     `json:"pc,omitempty"`     // pc 端主导航信息
	Mobile *GetApplicationVersionRespAppVersionAbilityNavigateMobile `json:"mobile,omitempty"` // 移动端主导航信息
}

// GetApplicationVersionRespAppVersionAbilityNavigateMobile ...
type GetApplicationVersionRespAppVersionAbilityNavigateMobile struct {
	Version       string `json:"version,omitempty"`         // 主导航小程序版本号
	ImageURL      string `json:"image_url,omitempty"`       // 默认图片 url
	HoverImageURL string `json:"hover_image_url,omitempty"` // 选中态图片 url
}

// GetApplicationVersionRespAppVersionAbilityNavigatePc ...
type GetApplicationVersionRespAppVersionAbilityNavigatePc struct {
	Version       string `json:"version,omitempty"`         // 主导航小程序版本号
	ImageURL      string `json:"image_url,omitempty"`       // 默认图片 url
	HoverImageURL string `json:"hover_image_url,omitempty"` // 选中态图片 url
}

// GetApplicationVersionRespAppVersionAbilityPlusMenu ...
type GetApplicationVersionRespAppVersionAbilityPlusMenu struct {
	PcAppLink     string `json:"pc_app_link,omitempty"`     // pc 端链接
	MobileAppLink string `json:"mobile_app_link,omitempty"` // 移动端链接
}

// GetApplicationVersionRespAppVersionAbilityWebApp ...
type GetApplicationVersionRespAppVersionAbilityWebApp struct {
	PcURL     string `json:"pc_url,omitempty"`     // pc 端 url
	MobileURL string `json:"mobile_url,omitempty"` // 移动端 url
}

// GetApplicationVersionRespAppVersionAbilityWorkplaceWidget ...
type GetApplicationVersionRespAppVersionAbilityWorkplaceWidget struct {
	MinLarkVersion string `json:"min_lark_version,omitempty"` // 最低兼容飞书版本号
}

// GetApplicationVersionRespAppVersionI18n ...
type GetApplicationVersionRespAppVersionI18n struct {
	I18nKey     string `json:"i18n_key,omitempty"`    // 国际化语言的 key, 可选值有: zh_cn: 中文, en_us: 英文, ja_jp: 日文
	Name        string `json:"name,omitempty"`        // 应用国际化名称
	Description string `json:"description,omitempty"` // 应用国际化描述（副标题）
	HelpUse     string `json:"help_use,omitempty"`    // 国际化帮助文档链接
}

// GetApplicationVersionRespAppVersionRemark ...
type GetApplicationVersionRespAppVersionRemark struct {
	Remark       string                                               `json:"remark,omitempty"`        // 备注说明
	UpdateRemark string                                               `json:"update_remark,omitempty"` // 更新说明
	Visibility   *GetApplicationVersionRespAppVersionRemarkVisibility `json:"visibility,omitempty"`    // 应用当前版本开发者编辑的可见性建议, 若开发者未编辑可见性建议, 则该字段无内容
}

// GetApplicationVersionRespAppVersionRemarkVisibility ...
type GetApplicationVersionRespAppVersionRemarkVisibility struct {
	IsAll         bool                                                              `json:"is_all,omitempty"`         // 是否全员可见
	VisibleList   *GetApplicationVersionRespAppVersionRemarkVisibilityVisibleList   `json:"visible_list,omitempty"`   // 可见名单
	InvisibleList *GetApplicationVersionRespAppVersionRemarkVisibilityInvisibleList `json:"invisible_list,omitempty"` // 不可见名单
}

// GetApplicationVersionRespAppVersionRemarkVisibilityInvisibleList ...
type GetApplicationVersionRespAppVersionRemarkVisibilityInvisibleList struct {
	OpenIDs       []string `json:"open_ids,omitempty"`       // 不可见性成员 open_id 列表
	DepartmentIDs []string `json:"department_ids,omitempty"` // 不可见性部门的 id 列表
	GroupIDs      []string `json:"group_ids,omitempty"`      // 不可见性成员 group_id 列表
}

// GetApplicationVersionRespAppVersionRemarkVisibilityVisibleList ...
type GetApplicationVersionRespAppVersionRemarkVisibilityVisibleList struct {
	OpenIDs       []string `json:"open_ids,omitempty"`       // 可见性成员 open_id 列表
	DepartmentIDs []string `json:"department_ids,omitempty"` // 可见性部门的 id 列表
	GroupIDs      []string `json:"group_ids,omitempty"`      // 可见性成员 group_id 列表
}

// GetApplicationVersionRespAppVersionScope ...
type GetApplicationVersionRespAppVersionScope struct {
	Scope       string `json:"scope,omitempty"`       // 应用权限
	Description string `json:"description,omitempty"` // 应用权限的国际化描述
	Level       int64  `json:"level,omitempty"`       // 权限等级描述, 可选值有: 1: 普通权限, 2: 高级权限, 3: 超敏感权限, 0: 未知等级
}

// getApplicationVersionResp ...
type getApplicationVersionResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetApplicationVersionResp `json:"data,omitempty"`
}
