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

// GetDriveCommentList 该接口用于根据文档 token 分页获取文档评论。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/list
func (r *DriveService) GetDriveCommentList(ctx context.Context, request *GetDriveCommentListReq, options ...MethodOptionFunc) (*GetDriveCommentListResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveCommentList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveCommentList mock enable")
		return r.cli.mock.mockDriveGetDriveCommentList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveCommentList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/comments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveCommentListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveCommentList mock DriveGetDriveCommentList method
func (r *Mock) MockDriveGetDriveCommentList(f func(ctx context.Context, request *GetDriveCommentListReq, options ...MethodOptionFunc) (*GetDriveCommentListResp, *Response, error)) {
	r.mockDriveGetDriveCommentList = f
}

// UnMockDriveGetDriveCommentList un-mock DriveGetDriveCommentList method
func (r *Mock) UnMockDriveGetDriveCommentList() {
	r.mockDriveGetDriveCommentList = nil
}

// GetDriveCommentListReq ...
type GetDriveCommentListReq struct {
	FileToken  string   `path:"file_token" json:"-"`    // 文档token, 示例值: "XIHSdYSI7oMEU1xrsnxc8fabcef"
	FileType   FileType `query:"file_type" json:"-"`    // 文档类型, 示例值: "doc", 可选值有: doc: 文档类型, sheet: 电子表格类型, file: 文件类型, docx: 新版文档类型
	IsWhole    *bool    `query:"is_whole" json:"-"`     // 是否全文评论, 示例值: false
	IsSolved   *bool    `query:"is_solved" json:"-"`    // 是否已解决（可选）, 示例值: false
	PageToken  *string  `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "7153511712153412356"
	PageSize   *int64   `query:"page_size" json:"-"`    // 分页大小, 示例值: "10", 最大值: `100`
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetDriveCommentListResp ...
type GetDriveCommentListResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetDriveCommentListRespItem `json:"items,omitempty"`      // 评论列表
}

// GetDriveCommentListRespItem ...
type GetDriveCommentListRespItem struct {
	CommentID    string                                `json:"comment_id,omitempty"`     // 评论ID（创建新评论可不填；如填写, 则视为回复已有评论）
	UserID       string                                `json:"user_id,omitempty"`        // 用户ID
	CreateTime   int64                                 `json:"create_time,omitempty"`    // 创建时间
	UpdateTime   int64                                 `json:"update_time,omitempty"`    // 更新时间
	IsSolved     bool                                  `json:"is_solved,omitempty"`      // 是否已解决
	SolvedTime   int64                                 `json:"solved_time,omitempty"`    // 解决评论时间
	SolverUserID string                                `json:"solver_user_id,omitempty"` // 解决评论者的用户ID
	HasMore      bool                                  `json:"has_more,omitempty"`       // 是否有更多回复
	PageToken    string                                `json:"page_token,omitempty"`     // 回复分页标记
	IsWhole      bool                                  `json:"is_whole,omitempty"`       // 是否是全文评论
	Quote        string                                `json:"quote,omitempty"`          // 如果是局部评论, 引用字段
	ReplyList    *GetDriveCommentListRespItemReplyList `json:"reply_list,omitempty"`     // 评论里的回复列表
}

// GetDriveCommentListRespItemReplyList ...
type GetDriveCommentListRespItemReplyList struct {
	Replies []*GetDriveCommentListRespItemReplyListReply `json:"replies,omitempty"` // 回复列表
}

// GetDriveCommentListRespItemReplyListReply ...
type GetDriveCommentListRespItemReplyListReply struct {
	ReplyID    string                                            `json:"reply_id,omitempty"`    // 回复ID
	UserID     string                                            `json:"user_id,omitempty"`     // 用户ID
	CreateTime int64                                             `json:"create_time,omitempty"` // 创建时间
	UpdateTime int64                                             `json:"update_time,omitempty"` // 更新时间
	Content    *GetDriveCommentListRespItemReplyListReplyContent `json:"content,omitempty"`     // 回复内容
	Extra      *GetDriveCommentListRespItemReplyListReplyExtra   `json:"extra,omitempty"`       // 回复的其他内容, 图片token等
}

// GetDriveCommentListRespItemReplyListReplyContent ...
type GetDriveCommentListRespItemReplyListReplyContent struct {
	Elements []*GetDriveCommentListRespItemReplyListReplyContentElement `json:"elements,omitempty"` // 回复的内容
}

// GetDriveCommentListRespItemReplyListReplyContentElement ...
type GetDriveCommentListRespItemReplyListReplyContentElement struct {
	Type     string                                                           `json:"type,omitempty"`      // 回复的内容元素, 可选值有: text_run: 普通文本, docs_link: at 云文档链接, person: at 联系人
	TextRun  *GetDriveCommentListRespItemReplyListReplyContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *GetDriveCommentListRespItemReplyListReplyContentElementDocsLink `json:"docs_link,omitempty"` // 文本内容
	Person   *GetDriveCommentListRespItemReplyListReplyContentElementPerson   `json:"person,omitempty"`    // 文本内容
}

// GetDriveCommentListRespItemReplyListReplyContentElementDocsLink ...
type GetDriveCommentListRespItemReplyListReplyContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at云文档
}

// GetDriveCommentListRespItemReplyListReplyContentElementPerson ...
type GetDriveCommentListRespItemReplyListReplyContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 回复 at联系人
}

// GetDriveCommentListRespItemReplyListReplyContentElementTextRun ...
type GetDriveCommentListRespItemReplyListReplyContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本
}

// GetDriveCommentListRespItemReplyListReplyExtra ...
type GetDriveCommentListRespItemReplyListReplyExtra struct {
	ImageList []string `json:"image_list,omitempty"` // 评论中的图片token list
}

// getDriveCommentListResp ...
type getDriveCommentListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetDriveCommentListResp `json:"data,omitempty"`
}
