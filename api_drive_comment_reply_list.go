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

// GetDriveCommentReplyList 该接口用于根据评论 ID, 获取该条评论对应的回复信息, 包括回复 ID、回复内容、回复人的用户 ID 等。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment-reply/list
// new doc: https://open.feishu.cn/document/server-docs/docs/CommentAPI/list-2
func (r *DriveService) GetDriveCommentReplyList(ctx context.Context, request *GetDriveCommentReplyListReq, options ...MethodOptionFunc) (*GetDriveCommentReplyListResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveCommentReplyList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetDriveCommentReplyList mock enable")
		return r.cli.mock.mockDriveGetDriveCommentReplyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveCommentReplyList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/comments/:comment_id/replies",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveCommentReplyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveCommentReplyList mock DriveGetDriveCommentReplyList method
func (r *Mock) MockDriveGetDriveCommentReplyList(f func(ctx context.Context, request *GetDriveCommentReplyListReq, options ...MethodOptionFunc) (*GetDriveCommentReplyListResp, *Response, error)) {
	r.mockDriveGetDriveCommentReplyList = f
}

// UnMockDriveGetDriveCommentReplyList un-mock DriveGetDriveCommentReplyList method
func (r *Mock) UnMockDriveGetDriveCommentReplyList() {
	r.mockDriveGetDriveCommentReplyList = nil
}

// GetDriveCommentReplyListReq ...
type GetDriveCommentReplyListReq struct {
	FileToken  string   `path:"file_token" json:"-"`    // 文档 Token, 示例值: "doxbcdl03Vsxhm7Qmnj110abcef"
	CommentID  string   `path:"comment_id" json:"-"`    // 评论 ID, 示例值: "1654857036541812356"
	PageSize   *int64   `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 最大值: `1 ～ 100`
	PageToken  *string  `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 1654857036541812356
	FileType   FileType `query:"file_type" json:"-"`    // 文档类型, 示例值: docx, 可选值有: doc: 文档类型, sheet: 电子表格类型, file: 文件夹类型, docx: 新版文档类型
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetDriveCommentReplyListResp ...
type GetDriveCommentReplyListResp struct {
	Items     []*GetDriveCommentReplyListRespItem `json:"items,omitempty"`      // 回复列表
	PageToken string                              `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetDriveCommentReplyListRespItem ...
type GetDriveCommentReplyListRespItem struct {
	ReplyID    string                                   `json:"reply_id,omitempty"`    // 回复 ID
	UserID     string                                   `json:"user_id,omitempty"`     // 用户 ID
	CreateTime int64                                    `json:"create_time,omitempty"` // 创建时间
	UpdateTime int64                                    `json:"update_time,omitempty"` // 更新时间
	Content    *GetDriveCommentReplyListRespItemContent `json:"content,omitempty"`     // 回复内容
	Extra      *GetDriveCommentReplyListRespItemExtra   `json:"extra,omitempty"`       // 回复的其他内容, 图片 Token 等
}

// GetDriveCommentReplyListRespItemContent ...
type GetDriveCommentReplyListRespItemContent struct {
	Elements []*GetDriveCommentReplyListRespItemContentElement `json:"elements,omitempty"` // 回复的内容
}

// GetDriveCommentReplyListRespItemContentElement ...
type GetDriveCommentReplyListRespItemContentElement struct {
	Type     string                                                  `json:"type,omitempty"`      // 回复的内容元素, 可选值有: text_run: 普通文本, docs_link: at 云文档链接, person: at 联系人
	TextRun  *GetDriveCommentReplyListRespItemContentElementTextRun  `json:"text_run,omitempty"`  // 文本内容
	DocsLink *GetDriveCommentReplyListRespItemContentElementDocsLink `json:"docs_link,omitempty"` // 添加云文档链接
	Person   *GetDriveCommentReplyListRespItemContentElementPerson   `json:"person,omitempty"`    // 添加用户的 user_id
}

// GetDriveCommentReplyListRespItemContentElementDocsLink ...
type GetDriveCommentReplyListRespItemContentElementDocsLink struct {
	URL string `json:"url,omitempty"` // 回复 at 云文档
}

// GetDriveCommentReplyListRespItemContentElementPerson ...
type GetDriveCommentReplyListRespItemContentElementPerson struct {
	UserID string `json:"user_id,omitempty"` // 添加用户的 user_id 以@用户
}

// GetDriveCommentReplyListRespItemContentElementTextRun ...
type GetDriveCommentReplyListRespItemContentElementTextRun struct {
	Text string `json:"text,omitempty"` // 回复 普通文本
}

// GetDriveCommentReplyListRespItemExtra ...
type GetDriveCommentReplyListRespItemExtra struct {
	ImageList []string `json:"image_list,omitempty"` // 评论中的图片 Token list
}

// getDriveCommentReplyListResp ...
type getDriveCommentReplyListResp struct {
	Code  int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                        `json:"msg,omitempty"`  // 错误描述
	Data  *GetDriveCommentReplyListResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}
