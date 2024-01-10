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

// MoveWikiNode 此方法用于在Wiki内移动节点, 支持跨知识空间移动。如果有子节点, 会携带子节点一起移动。
//
// 知识库权限要求:
// - 节点编辑权限
// - 原父节点容器编辑权限
// - 目的父节点容器编辑权限
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/move
// new doc: https://open.feishu.cn/document/server-docs/docs/wiki-v2/space-node/move
func (r *DriveService) MoveWikiNode(ctx context.Context, request *MoveWikiNodeReq, options ...MethodOptionFunc) (*MoveWikiNodeResp, *Response, error) {
	if r.cli.mock.mockDriveMoveWikiNode != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#MoveWikiNode mock enable")
		return r.cli.mock.mockDriveMoveWikiNode(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "MoveWikiNode",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/move",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(moveWikiNodeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveMoveWikiNode mock DriveMoveWikiNode method
func (r *Mock) MockDriveMoveWikiNode(f func(ctx context.Context, request *MoveWikiNodeReq, options ...MethodOptionFunc) (*MoveWikiNodeResp, *Response, error)) {
	r.mockDriveMoveWikiNode = f
}

// UnMockDriveMoveWikiNode un-mock DriveMoveWikiNode method
func (r *Mock) UnMockDriveMoveWikiNode() {
	r.mockDriveMoveWikiNode = nil
}

// MoveWikiNodeReq ...
type MoveWikiNodeReq struct {
	SpaceID           string  `path:"space_id" json:"-"`             // 知识空间id, 示例值: "7008061636015512345"
	NodeToken         string  `path:"node_token" json:"-"`           // 需要迁移的节点token, 示例值: "wikbcd6ydSUyOEzbdlt1BfpA5Yc"
	TargetParentToken *string `json:"target_parent_token,omitempty"` // 移动到的父节点token, 示例值: "wikbcd6ydSUyOEzbdlt1BfpA5Yc"
	TargetSpaceID     *string `json:"target_space_id,omitempty"`     // 移动到的知识空间ID, 示例值: "7008061636015512345"
}

// MoveWikiNodeResp ...
type MoveWikiNodeResp struct {
	Node *MoveWikiNodeRespNode `json:"node,omitempty"` // 移动后的节点信息
}

// MoveWikiNodeRespNode ...
type MoveWikiNodeRespNode struct {
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

// moveWikiNodeResp ...
type moveWikiNodeResp struct {
	Code int64             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *MoveWikiNodeResp `json:"data,omitempty"`
}
