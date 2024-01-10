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

// SearchHelpdeskFAQ 该接口用于搜索服务台知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/search
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/faq-management/faq/search
func (r *HelpdeskService) SearchHelpdeskFAQ(ctx context.Context, request *SearchHelpdeskFAQReq, options ...MethodOptionFunc) (*SearchHelpdeskFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskSearchHelpdeskFAQ != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#SearchHelpdeskFAQ mock enable")
		return r.cli.mock.mockHelpdeskSearchHelpdeskFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "SearchHelpdeskFAQ",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(searchHelpdeskFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskSearchHelpdeskFAQ mock HelpdeskSearchHelpdeskFAQ method
func (r *Mock) MockHelpdeskSearchHelpdeskFAQ(f func(ctx context.Context, request *SearchHelpdeskFAQReq, options ...MethodOptionFunc) (*SearchHelpdeskFAQResp, *Response, error)) {
	r.mockHelpdeskSearchHelpdeskFAQ = f
}

// UnMockHelpdeskSearchHelpdeskFAQ un-mock HelpdeskSearchHelpdeskFAQ method
func (r *Mock) UnMockHelpdeskSearchHelpdeskFAQ() {
	r.mockHelpdeskSearchHelpdeskFAQ = nil
}

// SearchHelpdeskFAQReq ...
type SearchHelpdeskFAQReq struct {
	Query     string  `query:"query" json:"-"`      // 搜索query, query内容如果不是英文, 包含中文空格等有两种编码策略: 1. url编码 2. base64编码, 同时加上base64=true参数, 示例值: wifi
	Base64    *string `query:"base64" json:"-"`     // 是否转换为base64, 输入true表示是, 不填写表示否, 示例值: true
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6936004780707807251
	PageSize  *int64  `query:"page_size" json:"-"`  // 示例值: 10, 默认值: `20`, 最大值: `100`
}

// SearchHelpdeskFAQResp ...
type SearchHelpdeskFAQResp struct {
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                       `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*SearchHelpdeskFAQRespItem `json:"items,omitempty"`      // 知识库列表
}

// SearchHelpdeskFAQRespItem ...
type SearchHelpdeskFAQRespItem struct {
	FAQID          string                                     `json:"faq_id,omitempty"`          // 知识库ID
	ID             string                                     `json:"id,omitempty"`              // 知识库旧版ID, 请使用faq_id
	HelpdeskID     string                                     `json:"helpdesk_id,omitempty"`     // 服务台ID
	Question       string                                     `json:"question,omitempty"`        // 问题
	Answer         string                                     `json:"answer,omitempty"`          // 答案
	AnswerRichtext []*SearchHelpdeskFAQRespItemAnswerRichtext `json:"answer_richtext,omitempty"` // 富文本答案
	CreateTime     int64                                      `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int64                                      `json:"update_time,omitempty"`     // 修改时间
	Categories     []*HelpdeskCategory                        `json:"categories,omitempty"`      // 分类
	Tags           []string                                   `json:"tags,omitempty"`            // 相似问题列表
	ExpireTime     int64                                      `json:"expire_time,omitempty"`     // 失效时间
	UpdateUser     *SearchHelpdeskFAQRespItemUpdateUser       `json:"update_user,omitempty"`     // 更新用户
	CreateUser     *SearchHelpdeskFAQRespItemCreateUser       `json:"create_user,omitempty"`     // 创建用户
}

// SearchHelpdeskFAQRespItemAnswerRichtext ...
type SearchHelpdeskFAQRespItemAnswerRichtext struct {
	Content string `json:"content,omitempty"` // 内容
	Type    string `json:"type,omitempty"`    // 类型
}

// SearchHelpdeskFAQRespItemCreateUser ...
type SearchHelpdeskFAQRespItemCreateUser struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// SearchHelpdeskFAQRespItemUpdateUser ...
type SearchHelpdeskFAQRespItemUpdateUser struct {
	ID         string `json:"id,omitempty"`         // 用户ID
	AvatarURL  string `json:"avatar_url,omitempty"` // 用户头像url
	Name       string `json:"name,omitempty"`       // 用户名
	Department string `json:"department,omitempty"` // 所在部门名称
	City       string `json:"city,omitempty"`       // 城市
	Country    string `json:"country,omitempty"`    // 国家代号(CountryCode), 参考: http://www.mamicode.com/info-detail-2186501.html
}

// searchHelpdeskFAQResp ...
type searchHelpdeskFAQResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *SearchHelpdeskFAQResp `json:"data,omitempty"`
}
