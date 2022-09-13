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

// UpdateBaikeEntity 通过此接口更新已有的词条, 不需要百科管理员审核可直接写入词库, 请慎重使用【租户管理员请慎重审批】。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/update
func (r *BaikeService) UpdateBaikeEntity(ctx context.Context, request *UpdateBaikeEntityReq, options ...MethodOptionFunc) (*UpdateBaikeEntityResp, *Response, error) {
	if r.cli.mock.mockBaikeUpdateBaikeEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Baike#UpdateBaikeEntity mock enable")
		return r.cli.mock.mockBaikeUpdateBaikeEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "UpdateBaikeEntity",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/entities/:entity_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateBaikeEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeUpdateBaikeEntity mock BaikeUpdateBaikeEntity method
func (r *Mock) MockBaikeUpdateBaikeEntity(f func(ctx context.Context, request *UpdateBaikeEntityReq, options ...MethodOptionFunc) (*UpdateBaikeEntityResp, *Response, error)) {
	r.mockBaikeUpdateBaikeEntity = f
}

// UnMockBaikeUpdateBaikeEntity un-mock BaikeUpdateBaikeEntity method
func (r *Mock) UnMockBaikeUpdateBaikeEntity() {
	r.mockBaikeUpdateBaikeEntity = nil
}

// UpdateBaikeEntityReq ...
type UpdateBaikeEntityReq struct {
	EntityID    string                           `path:"entity_id" json:"-"`     // 实体词 ID, 示例值: "enterprise_40217521"
	UserIDType  *IDType                          `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MainKeys    []*UpdateBaikeEntityReqMainKey   `json:"main_keys,omitempty"`    // 词条名, 最大长度: `1`
	Aliases     []*UpdateBaikeEntityReqAliase    `json:"aliases,omitempty"`      // 别名, 最大长度: `10`
	Description *string                          `json:"description,omitempty"`  // 词条释义（纯文本格式）, 示例值: "企业百科是飞书提供的一款知识管理工具, 通过企业百科可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通", 最大长度: `5000` 字符
	RelatedMeta *UpdateBaikeEntityReqRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	OuterInfo   *UpdateBaikeEntityReqOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    *string                          `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[企业百科指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分, 示例值: "<b>加粗</b><i>斜体</i><p><a href=\"https://feishu.cn\">链接</a></p><p><span>企业百科是飞书提供的一款知识管理工具, 通过企业百科可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通</span></p>", 最大长度: `5000` 字符
}

// UpdateBaikeEntityReqAliase ...
type UpdateBaikeEntityReqAliase struct {
	Key           string                                   `json:"key,omitempty"`            // 名称的值, 示例值: "企业百科"
	DisplayStatus *UpdateBaikeEntityReqAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// UpdateBaikeEntityReqAliaseDisplayStatus ...
type UpdateBaikeEntityReqAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// UpdateBaikeEntityReqMainKey ...
type UpdateBaikeEntityReqMainKey struct {
	Key           string                                    `json:"key,omitempty"`            // 名称的值, 示例值: "企业百科"
	DisplayStatus *UpdateBaikeEntityReqMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// UpdateBaikeEntityReqMainKeyDisplayStatus ...
type UpdateBaikeEntityReqMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// UpdateBaikeEntityReqOuterInfo ...
type UpdateBaikeEntityReqOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）, 示例值: "星云", 长度范围: `2` ～ `32` 字符
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）, 示例值: "client_6539i3498d", 长度范围: `1` ～ `64` 字符
}

// UpdateBaikeEntityReqRelatedMeta ...
type UpdateBaikeEntityReqRelatedMeta struct {
	Users           []*UpdateBaikeEntityReqRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*UpdateBaikeEntityReqRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*UpdateBaikeEntityReqRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*UpdateBaikeEntityReqRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*UpdateBaikeEntityReqRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*UpdateBaikeEntityReqRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*UpdateBaikeEntityReqRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*UpdateBaikeEntityReqRelatedMetaImage          `json:"images,omitempty"`          // 上传的图片, 最大长度: `10`
}

// UpdateBaikeEntityReqRelatedMetaAbbreviation ...
type UpdateBaikeEntityReqRelatedMetaAbbreviation struct {
	ID *string `json:"id,omitempty"` // 相关词条 ID, 示例值: "enterprise_51587960"
}

// UpdateBaikeEntityReqRelatedMetaChat ...
type UpdateBaikeEntityReqRelatedMetaChat struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// UpdateBaikeEntityReqRelatedMetaClassification ...
type UpdateBaikeEntityReqRelatedMetaClassification struct {
	ID       string  `json:"id,omitempty"`        // 二级分类 ID, 示例值: "7049606926702837761"
	FatherID *string `json:"father_id,omitempty"` // 对应一级分类 ID, 示例值: "7049606926702837777"
}

// UpdateBaikeEntityReqRelatedMetaDoc ...
type UpdateBaikeEntityReqRelatedMetaDoc struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "企业百科帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// UpdateBaikeEntityReqRelatedMetaImage ...
type UpdateBaikeEntityReqRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token, 示例值: "boxbcEcmKiD3SGHvgqWTpvdc7jc"
}

// UpdateBaikeEntityReqRelatedMetaLink ...
type UpdateBaikeEntityReqRelatedMetaLink struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "企业百科帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// UpdateBaikeEntityReqRelatedMetaOncall ...
type UpdateBaikeEntityReqRelatedMetaOncall struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// UpdateBaikeEntityReqRelatedMetaUser ...
type UpdateBaikeEntityReqRelatedMetaUser struct {
	ID    string  `json:"id,omitempty"`    // 对应相关信息 ID, 示例值: "格式请看请求体示例"
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "企业百科帮助中心"
}

// UpdateBaikeEntityResp ...
type UpdateBaikeEntityResp struct {
	Entity *UpdateBaikeEntityRespEntity `json:"entity,omitempty"` // 词条信息
}

// UpdateBaikeEntityRespEntity ...
type UpdateBaikeEntityRespEntity struct {
	ID          string                                  `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）
	MainKeys    []*UpdateBaikeEntityRespEntityMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*UpdateBaikeEntityRespEntityAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                  `json:"description,omitempty"`  // 词条释义（纯文本格式）
	CreateTime  string                                  `json:"create_time,omitempty"`  // 词条创建时间
	UpdateTime  string                                  `json:"update_time,omitempty"`  // 词条最近更新时间
	RelatedMeta *UpdateBaikeEntityRespEntityRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	Statistics  *UpdateBaikeEntityRespEntityStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *UpdateBaikeEntityRespEntityOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    string                                  `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[企业百科指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
}

// UpdateBaikeEntityRespEntityAliase ...
type UpdateBaikeEntityRespEntityAliase struct {
	Key           string                                          `json:"key,omitempty"`            // 名称的值
	DisplayStatus *UpdateBaikeEntityRespEntityAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// UpdateBaikeEntityRespEntityAliaseDisplayStatus ...
type UpdateBaikeEntityRespEntityAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// UpdateBaikeEntityRespEntityMainKey ...
type UpdateBaikeEntityRespEntityMainKey struct {
	Key           string                                           `json:"key,omitempty"`            // 名称的值
	DisplayStatus *UpdateBaikeEntityRespEntityMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// UpdateBaikeEntityRespEntityMainKeyDisplayStatus ...
type UpdateBaikeEntityRespEntityMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// UpdateBaikeEntityRespEntityOuterInfo ...
type UpdateBaikeEntityRespEntityOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）
}

// UpdateBaikeEntityRespEntityRelatedMeta ...
type UpdateBaikeEntityRespEntityRelatedMeta struct {
	Users           []*UpdateBaikeEntityRespEntityRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*UpdateBaikeEntityRespEntityRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*UpdateBaikeEntityRespEntityRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*UpdateBaikeEntityRespEntityRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*UpdateBaikeEntityRespEntityRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*UpdateBaikeEntityRespEntityRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*UpdateBaikeEntityRespEntityRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*UpdateBaikeEntityRespEntityRelatedMetaImage          `json:"images,omitempty"`          // 上传的图片
}

// UpdateBaikeEntityRespEntityRelatedMetaAbbreviation ...
type UpdateBaikeEntityRespEntityRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关词条 ID
}

// UpdateBaikeEntityRespEntityRelatedMetaChat ...
type UpdateBaikeEntityRespEntityRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// UpdateBaikeEntityRespEntityRelatedMetaClassification ...
type UpdateBaikeEntityRespEntityRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	Name     string `json:"name,omitempty"`      // 二级分类名称
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// UpdateBaikeEntityRespEntityRelatedMetaDoc ...
type UpdateBaikeEntityRespEntityRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// UpdateBaikeEntityRespEntityRelatedMetaImage ...
type UpdateBaikeEntityRespEntityRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token
}

// UpdateBaikeEntityRespEntityRelatedMetaLink ...
type UpdateBaikeEntityRespEntityRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// UpdateBaikeEntityRespEntityRelatedMetaOncall ...
type UpdateBaikeEntityRespEntityRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// UpdateBaikeEntityRespEntityRelatedMetaUser ...
type UpdateBaikeEntityRespEntityRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// UpdateBaikeEntityRespEntityStatistics ...
type UpdateBaikeEntityRespEntityStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 累计点赞
	DislikeCount int64 `json:"dislike_count,omitempty"` // 当前词条版本收到的负反馈数量
}

// updateBaikeEntityResp ...
type updateBaikeEntityResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *UpdateBaikeEntityResp `json:"data,omitempty"`
}
