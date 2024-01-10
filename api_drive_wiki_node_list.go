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

// GetWikiNodeList 此接口用于分页获取Wiki节点的子节点列表。
//
// 此接口为分页接口。由于权限过滤, 可能返回列表为空, 但分页标记（has_more）为true, 可以继续分页请求。
// 知识库权限要求, 当前使用的 access token 所代表的应用或用户拥有:
// - 父节点阅读权限
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/list
// new doc: https://open.feishu.cn/document/server-docs/docs/wiki-v2/space-node/list
func (r *DriveService) GetWikiNodeList(ctx context.Context, request *GetWikiNodeListReq, options ...MethodOptionFunc) (*GetWikiNodeListResp, *Response, error) {
	if r.cli.mock.mockDriveGetWikiNodeList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetWikiNodeList mock enable")
		return r.cli.mock.mockDriveGetWikiNodeList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetWikiNodeList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/spaces/:space_id/nodes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getWikiNodeListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetWikiNodeList mock DriveGetWikiNodeList method
func (r *Mock) MockDriveGetWikiNodeList(f func(ctx context.Context, request *GetWikiNodeListReq, options ...MethodOptionFunc) (*GetWikiNodeListResp, *Response, error)) {
	r.mockDriveGetWikiNodeList = f
}

// UnMockDriveGetWikiNodeList un-mock DriveGetWikiNodeList method
func (r *Mock) UnMockDriveGetWikiNodeList() {
	r.mockDriveGetWikiNodeList = nil
}

// GetWikiNodeListReq ...
type GetWikiNodeListReq struct {
	SpaceID         string  `path:"space_id" json:"-"`           // 知识空间id, 示例值: "6946843325487906839"
	PageSize        *int64  `query:"page_size" json:"-"`         // 分页大小, 示例值: 10, 最大值: `50`
	PageToken       *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6946843325487456878
	ParentNodeToken *string `query:"parent_node_token" json:"-"` // 父节点token, 示例值: wikcnKQ1k3p**8Vabce
}

// GetWikiNodeListResp ...
type GetWikiNodeListResp struct {
	Items     []*GetWikiNodeListRespItem `json:"items,omitempty"`      // 数据列表
	PageToken string                     `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                       `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetWikiNodeListRespItem ...
type GetWikiNodeListRespItem struct {
	SpaceID         string `json:"space_id,omitempty"`          // 知识空间id, [获取方式](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-overview)
	NodeToken       string `json:"node_token,omitempty"`        // 节点token
	ObjToken        string `json:"obj_token,omitempty"`         // 对应文档类型的token, 可根据 obj_type 判断属于哪种文档类型。
	ObjType         string `json:"obj_type,omitempty"`          // 文档类型, 对于快捷方式, 该字段是对应的实体的obj_type, 可选值有: doc: 旧版文档, sheet: 表格, mindnote: 思维导图, bitable: 多维表格, file: 文件, docx: 新版文档
	ParentNodeToken string `json:"parent_node_token,omitempty"` // 父节点 token。若当前节点为一级节点, 父节点 token 为空。
	NodeType        string `json:"node_type,omitempty"`         // 节点类型, 可选值有: origin: 实体, shortcut: 快捷方式
	OriginNodeToken string `json:"origin_node_token,omitempty"` // 快捷方式对应的实体node_token, 当节点为快捷方式时, 该值不为空。
	OriginSpaceID   string `json:"origin_space_id,omitempty"`   // 快捷方式对应的实体所在的space id
	HasChild        bool   `json:"has_child,omitempty"`         // 是否有子节点
	Title           string `json:"title,omitempty"`             // 文档标题
	ObjCreateTime   string `json:"obj_create_time,omitempty"`   // 文档创建时间
	ObjEditTime     string `json:"obj_edit_time,omitempty"`     // 文档最近编辑时间
	NodeCreateTime  string `json:"node_create_time,omitempty"`  // 节点创建时间
	Creator         string `json:"creator,omitempty"`           // 节点创建者
	Owner           string `json:"owner,omitempty"`             // 节点所有者
}

// getWikiNodeListResp ...
type getWikiNodeListResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetWikiNodeListResp `json:"data,omitempty"`
}
