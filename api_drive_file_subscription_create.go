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

// CreateDriveFileSubscription 订阅文档中的变更事件, 当前支持文档评论订阅, 订阅后文档评论更新会有“云文档助手”推送给订阅的用户
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-subscription/create
// new doc: https://open.feishu.cn/document/server-docs/docs/docs-assistant/file-subscription/create
func (r *DriveService) CreateDriveFileSubscription(ctx context.Context, request *CreateDriveFileSubscriptionReq, options ...MethodOptionFunc) (*CreateDriveFileSubscriptionResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveFileSubscription != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveFileSubscription mock enable")
		return r.cli.mock.mockDriveCreateDriveFileSubscription(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "CreateDriveFileSubscription",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/subscriptions",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(createDriveFileSubscriptionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveFileSubscription mock DriveCreateDriveFileSubscription method
func (r *Mock) MockDriveCreateDriveFileSubscription(f func(ctx context.Context, request *CreateDriveFileSubscriptionReq, options ...MethodOptionFunc) (*CreateDriveFileSubscriptionResp, *Response, error)) {
	r.mockDriveCreateDriveFileSubscription = f
}

// UnMockDriveCreateDriveFileSubscription un-mock DriveCreateDriveFileSubscription method
func (r *Mock) UnMockDriveCreateDriveFileSubscription() {
	r.mockDriveCreateDriveFileSubscription = nil
}

// CreateDriveFileSubscriptionReq ...
type CreateDriveFileSubscriptionReq struct {
	FileToken        string   `path:"file_token" json:"-"`         // 文档token, 示例值: "doxcnxxxxxxxxxxxxxxxxxxxxxx"
	SubscriptionID   *string  `json:"subscription_id,omitempty"`   // 订阅关系ID, 示例值: "1234567890987654321"
	SubscriptionType string   `json:"subscription_type,omitempty"` // 订阅类型, 示例值: "comment_update", 可选值有: comment_update: 评论更新
	IsSubcribe       *bool    `json:"is_subcribe,omitempty"`       // 是否订阅, 示例值: true
	FileType         FileType `json:"file_type,omitempty"`         // 文档类型, 示例值: "doc", 可选值有: doc: 文档, docx: 新版文档, wiki: 知识库wiki
}

// CreateDriveFileSubscriptionResp ...
type CreateDriveFileSubscriptionResp struct {
	SubscriptionID   string   `json:"subscription_id,omitempty"`   // 订阅关系ID
	SubscriptionType string   `json:"subscription_type,omitempty"` // 订阅类型, 可选值有: comment_update: 评论更新
	IsSubcribe       bool     `json:"is_subcribe,omitempty"`       // 是否订阅
	FileType         FileType `json:"file_type,omitempty"`         // 文档类型, 可选值有: doc: 旧版文档, docx: 新版文档, wiki: 知识库
}

// createDriveFileSubscriptionResp ...
type createDriveFileSubscriptionResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveFileSubscriptionResp `json:"data,omitempty"`
}
