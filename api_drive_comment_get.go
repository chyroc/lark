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

// GetDriveComment 获取云文档中的某条全文评论, 暂时不支持局部评论
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/get
// new doc: https://open.feishu.cn/document/server-docs/docs/CommentAPI/get
func (r *DriveService) GetDriveComment(ctx context.Context, request *GetDriveCommentReq, options ...MethodOptionFunc) (*GetDriveCommentResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveComment mock enable")
		return r.cli.mock.mockDriveGetDriveComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveComment",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveComment mock DriveGetDriveComment method
func (r *Mock) MockDriveGetDriveComment(f func(ctx context.Context, request *GetDriveCommentReq, options ...MethodOptionFunc) (*GetDriveCommentResp, *Response, error)) {
	r.mockDriveGetDriveComment = f
}

// UnMockDriveGetDriveComment un-mock DriveGetDriveComment method
func (r *Mock) UnMockDriveGetDriveComment() {
	r.mockDriveGetDriveComment = nil
}

// GetDriveCommentReq ...
type GetDriveCommentReq struct {
	FileToken  string   `path:"file_token" json:"-"`    // 文档token, 示例值: "doccnHh7U87HOFpii5u5G*"
	CommentID  string   `path:"comment_id" json:"-"`    // 评论ID, 示例值: "6916106822734578184"
	FileType   FileType `query:"file_type" json:"-"`    // 文档类型, 示例值: "doc", 可选值有: doc: 文档, sheet: 表格, file: 文件, docx: 新版文档
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetDriveCommentResp ...
type GetDriveCommentResp struct {
	CommentID    string                        `json:"comment_id,omitempty"`     // 评论ID（创建新评论可不填；如填写, 则视为回复已有评论）
	UserID       string                        `json:"user_id,omitempty"`        // 用户ID
	CreateTime   int64                         `json:"create_time,omitempty"`    // 创建时间
	UpdateTime   int64                         `json:"update_time,omitempty"`    // 更新时间
	IsSolved     bool                          `json:"is_solved,omitempty"`      // 是否已解决
	SolvedTime   int64                         `json:"solved_time,omitempty"`    // 解决评论时间
	SolverUserID string                        `json:"solver_user_id,omitempty"` // 解决评论者的用户ID
	HasMore      bool                          `json:"has_more,omitempty"`       // 是否还有更多项
	PageToken    string                        `json:"page_token,omitempty"`     // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	IsWhole      bool                          `json:"is_whole,omitempty"`       // 是否是全文评论
	Quote        string                        `json:"quote,omitempty"`          // 如果是局部评论, 引用字段
	ReplyList    *GetDriveCommentRespReplyList `json:"reply_list,omitempty"`     // 评论里的回复列表
}

// GetDriveCommentRespReplyList ...
type GetDriveCommentRespReplyList struct {
	Replies []*GetDriveCommentRespReplyListReply `json:"replies,omitempty"` // 回复列表
}

// GetDriveCommentRespReplyListReply ...
type GetDriveCommentRespReplyListReply struct {
	ReplyID    string                                    `json:"reply_id,omitempty"`    // 回复ID
	UserID     string                                    `json:"user_id,omitempty"`     // 用户ID
	CreateTime int64                                     `json:"create_time,omitempty"` // 创建时间
	UpdateTime int64                                     `json:"update_time,omitempty"` // 更新时间
	Content    *GetDriveCommentRespReplyListReplyContent `json:"content,omitempty"`     // 回复内容
	Extra      *GetDriveCommentRespReplyListReplyExtra   `json:"extra,omitempty"`       // 回复的其他内容, 图片token等
}

// GetDriveCommentRespReplyListReplyContent ...
type GetDriveCommentRespReplyListReplyContent struct {
	Elements []*GetDriveCommentRespReplyListReplyContentElement `json:"elements,omitempty"` // 回复的内容
}

// GetDriveCommentRespReplyListReplyContentElement ...
type GetDriveCommentRespReplyListReplyContentElement struct {
	Type     string                                                   `json:"type,omitempty"`      // 回复的内容元素, 可选值有: text_run: 普通文本, docs_link: at 云文档链接, person: at 联系人
	TextRun  *GetDriveCommentRespReplyListReplyContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *GetDriveCommentRespReplyListReplyContentElementDocsLink `json:"docs_link,omitempty"` // 文本内容
	Person   *GetDriveCommentRespReplyListReplyContentElementPerson   `json:"person,omitempty"`    // 文本内容
}

// GetDriveCommentRespReplyListReplyContentElementDocsLink ...
type GetDriveCommentRespReplyListReplyContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at云文档
}

// GetDriveCommentRespReplyListReplyContentElementPerson ...
type GetDriveCommentRespReplyListReplyContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 回复 at联系人
}

// GetDriveCommentRespReplyListReplyContentElementTextRun ...
type GetDriveCommentRespReplyListReplyContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本
}

// GetDriveCommentRespReplyListReplyExtra ...
type GetDriveCommentRespReplyListReplyExtra struct {
	ImageList []string `json:"image_list,omitempty"` // 评论中的图片token list
}

// getDriveCommentResp ...
type getDriveCommentResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetDriveCommentResp `json:"data,omitempty"`
}
