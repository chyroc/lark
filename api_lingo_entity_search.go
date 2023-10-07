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

// SearchLingoEntity 传入关键词, 与词条名、别名、释义等信息进行模糊匹配, 返回搜到的词条信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/search
func (r *LingoService) SearchLingoEntity(ctx context.Context, request *SearchLingoEntityReq, options ...MethodOptionFunc) (*SearchLingoEntityResp, *Response, error) {
	if r.cli.mock.mockLingoSearchLingoEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Lingo#SearchLingoEntity mock enable")
		return r.cli.mock.mockLingoSearchLingoEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Lingo",
		API:                   "SearchLingoEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/lingo/v1/entities/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(searchLingoEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockLingoSearchLingoEntity mock LingoSearchLingoEntity method
func (r *Mock) MockLingoSearchLingoEntity(f func(ctx context.Context, request *SearchLingoEntityReq, options ...MethodOptionFunc) (*SearchLingoEntityResp, *Response, error)) {
	r.mockLingoSearchLingoEntity = f
}

// UnMockLingoSearchLingoEntity un-mock LingoSearchLingoEntity method
func (r *Mock) UnMockLingoSearchLingoEntity() {
	r.mockLingoSearchLingoEntity = nil
}

// SearchLingoEntityReq ...
type SearchLingoEntityReq struct {
	PageSize             *int64                                    `query:"page_size" json:"-"`             // 每页返回的词条量, 示例值: 20, 默认值: `20`, 取值范围: `1` ～ `100`
	PageToken            *string                                   `query:"page_token" json:"-"`            // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: b152fa6e6f62a291019a04c3a93f365f8ac641910506ff15ff4cad6534e087cb4ed8fa2c
	RepoID               *string                                   `query:"repo_id" json:"-"`               // 词库ID(不传时默认在全员词库内搜索), 如以应用身份搜索非全员词库中的词条, 需要在“词库设置”页面添加应用；若以用户身份搜索非全员词库中的词条, 该用户需要拥有对应词库的可见权限, 示例值: 7202510112396640276
	UserIDType           *IDType                                   `query:"user_id_type" json:"-"`          // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Query                *string                                   `json:"query,omitempty"`                 // 搜索关键词, 示例值: "飞书词典", 长度范围: `1` ～ `100` 字符
	ClassificationFilter *SearchLingoEntityReqClassificationFilter `json:"classification_filter,omitempty"` // 分类筛选
	Sources              []int64                                   `json:"sources,omitempty"`               // 词条的创建来源, 1: 用户主动创建, 2: 批量导入, 3: 官方词, 4: OpenAPI 创建, 示例值: [1]
	Creators             []string                                  `json:"creators,omitempty"`              // 创建者, 示例值: ["ou_30b07b63089ea46518789914dac63d36"]
}

// SearchLingoEntityReqClassificationFilter ...
type SearchLingoEntityReqClassificationFilter struct {
	Include []string `json:"include,omitempty"` // 需要获取的分类, 示例值: ["7195482254012055580"]
	Exclude []string `json:"exclude,omitempty"` // 需要排除的分类, 示例值: ["7195482254012055581"]
}

// SearchLingoEntityResp ...
type SearchLingoEntityResp struct {
	Entities  []*SearchLingoEntityRespEntitie `json:"entities,omitempty"`   // 搜索结果
	PageToken string                          `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// SearchLingoEntityRespEntitie ...
type SearchLingoEntityRespEntitie struct {
	ID          string                                   `json:"id,omitempty"`           // 词条 ID
	MainKeys    []*SearchLingoEntityRespEntitieMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*SearchLingoEntityRespEntitieAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                   `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001
	Creator     string                                   `json:"creator,omitempty"`      // 创建者
	CreateTime  string                                   `json:"create_time,omitempty"`  // 词条创建时间（秒级时间戳）
	Updater     string                                   `json:"updater,omitempty"`      // 最近一次更新者
	UpdateTime  string                                   `json:"update_time,omitempty"`  // 最近一次更新实体词时间（秒级时间戳）
	RelatedMeta *SearchLingoEntityRespEntitieRelatedMeta `json:"related_meta,omitempty"` // 词条的相关资源
	Statistics  *SearchLingoEntityRespEntitieStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *SearchLingoEntityRespEntitieOuterInfo   `json:"outer_info,omitempty"`   // 词条在外部系统中的信息
	RichText    string                                   `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[企业百科指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
	Source      int64                                    `json:"source,omitempty"`       // 词条的创建来源, 1: 用户主动创建, 2: 批量导入, 3: 官方词, 4: OpenAPI 创建
}

// SearchLingoEntityRespEntitieAliase ...
type SearchLingoEntityRespEntitieAliase struct {
	Key           string                                           `json:"key,omitempty"`            // 名称
	DisplayStatus *SearchLingoEntityRespEntitieAliaseDisplayStatus `json:"display_status,omitempty"` // 展示状态
}

// SearchLingoEntityRespEntitieAliaseDisplayStatus ...
type SearchLingoEntityRespEntitieAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 是否允许在 IM 和 Doc 等场景进行高亮提示
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// SearchLingoEntityRespEntitieMainKey ...
type SearchLingoEntityRespEntitieMainKey struct {
	Key           string                                            `json:"key,omitempty"`            // 名称
	DisplayStatus *SearchLingoEntityRespEntitieMainKeyDisplayStatus `json:"display_status,omitempty"` // 展示状态
}

// SearchLingoEntityRespEntitieMainKeyDisplayStatus ...
type SearchLingoEntityRespEntitieMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 是否允许在 IM 和 Doc 等场景进行高亮提示
	AllowSearch    bool `json:"allow_search,omitempty"`    // 是否允许在飞书中被搜索到
}

// SearchLingoEntityRespEntitieOuterInfo ...
type SearchLingoEntityRespEntitieOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 数据提供方（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统的唯一标识, 可用来和其他平台的内容进行绑定。需保证和百科词条唯一对应（不能包含中横线 "-"）
}

// SearchLingoEntityRespEntitieRelatedMeta ...
type SearchLingoEntityRespEntitieRelatedMeta struct {
	Users           []*SearchLingoEntityRespEntitieRelatedMetaUser           `json:"users,omitempty"`           // 相关人信息
	Chats           []*SearchLingoEntityRespEntitieRelatedMetaChat           `json:"chats,omitempty"`           // 相关公开群信息
	Docs            []*SearchLingoEntityRespEntitieRelatedMetaDoc            `json:"docs,omitempty"`            // 关联文档信息
	Oncalls         []*SearchLingoEntityRespEntitieRelatedMetaOncall         `json:"oncalls,omitempty"`         // 关联服务台信息
	Links           []*SearchLingoEntityRespEntitieRelatedMetaLink           `json:"links,omitempty"`           // 关联链接信息
	Abbreviations   []*SearchLingoEntityRespEntitieRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条信息
	Classifications []*SearchLingoEntityRespEntitieRelatedMetaClassification `json:"classifications,omitempty"` // 所属分类信息（不支持传入一级分类。词条不可同时属于同一个一级分类下的多个二级分类, 一级分类下的二级分类互斥）
	Images          []*SearchLingoEntityRespEntitieRelatedMetaImage          `json:"images,omitempty"`          // 上传的相关图片
}

// SearchLingoEntityRespEntitieRelatedMetaAbbreviation ...
type SearchLingoEntityRespEntitieRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关词条 id
}

// SearchLingoEntityRespEntitieRelatedMetaChat ...
type SearchLingoEntityRespEntitieRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 数据 id
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// SearchLingoEntityRespEntitieRelatedMetaClassification ...
type SearchLingoEntityRespEntitieRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 唯一分类 ID
	Name     string `json:"name,omitempty"`      // 分类名称
	FatherID string `json:"father_id,omitempty"` // 父级分类的 ID
}

// SearchLingoEntityRespEntitieRelatedMetaDoc ...
type SearchLingoEntityRespEntitieRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// SearchLingoEntityRespEntitieRelatedMetaImage ...
type SearchLingoEntityRespEntitieRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传后的图片 token
}

// SearchLingoEntityRespEntitieRelatedMetaLink ...
type SearchLingoEntityRespEntitieRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// SearchLingoEntityRespEntitieRelatedMetaOncall ...
type SearchLingoEntityRespEntitieRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 数据 id
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// SearchLingoEntityRespEntitieRelatedMetaUser ...
type SearchLingoEntityRespEntitieRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 数据 id
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// SearchLingoEntityRespEntitieStatistics ...
type SearchLingoEntityRespEntitieStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 点赞数量
	DislikeCount int64 `json:"dislike_count,omitempty"` // 点踩数量
}

// searchLingoEntityResp ...
type searchLingoEntityResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *SearchLingoEntityResp `json:"data,omitempty"`
}