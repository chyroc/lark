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

// UpdateContactGroup 使用该接口更新用户组信息, 请注意更新用户组时应用的通讯录权限范围需为“全部员工”, 否则会更新失败。[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/patch
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/group/patch
func (r *ContactService) UpdateContactGroup(ctx context.Context, request *UpdateContactGroupReq, options ...MethodOptionFunc) (*UpdateContactGroupResp, *Response, error) {
	if r.cli.mock.mockContactUpdateContactGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UpdateContactGroup mock enable")
		return r.cli.mock.mockContactUpdateContactGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateContactGroup",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateContactGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateContactGroup mock ContactUpdateContactGroup method
func (r *Mock) MockContactUpdateContactGroup(f func(ctx context.Context, request *UpdateContactGroupReq, options ...MethodOptionFunc) (*UpdateContactGroupResp, *Response, error)) {
	r.mockContactUpdateContactGroup = f
}

// UnMockContactUpdateContactGroup un-mock ContactUpdateContactGroup method
func (r *Mock) UnMockContactUpdateContactGroup() {
	r.mockContactUpdateContactGroup = nil
}

// UpdateContactGroupReq ...
type UpdateContactGroupReq struct {
	GroupID     string  `path:"group_id" json:"-"`     // 用户组ID, 示例值: "g187131"
	Name        *string `json:"name,omitempty"`        // 用户组的名字, 企业内唯一, 最大长度: 100 字符, 示例值: "外包 IT 用户组"
	Description *string `json:"description,omitempty"` // 用户组描述信息, 最大长度: 500 字, 示例值: "IT 外包用户组, 需要进行细粒度权限管控"
}

// UpdateContactGroupResp ...
type UpdateContactGroupResp struct {
}

// updateContactGroupResp ...
type updateContactGroupResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateContactGroupResp `json:"data,omitempty"`
}
