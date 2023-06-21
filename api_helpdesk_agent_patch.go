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

// UpdateHelpdeskAgent 更新客服状态等信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent/patch
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/agent-function/agent/patch
func (r *HelpdeskService) UpdateHelpdeskAgent(ctx context.Context, request *UpdateHelpdeskAgentReq, options ...MethodOptionFunc) (*UpdateHelpdeskAgentResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskAgent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskAgent mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskAgent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskAgent",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/agents/:agent_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskAgentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskUpdateHelpdeskAgent mock HelpdeskUpdateHelpdeskAgent method
func (r *Mock) MockHelpdeskUpdateHelpdeskAgent(f func(ctx context.Context, request *UpdateHelpdeskAgentReq, options ...MethodOptionFunc) (*UpdateHelpdeskAgentResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskAgent = f
}

// UnMockHelpdeskUpdateHelpdeskAgent un-mock HelpdeskUpdateHelpdeskAgent method
func (r *Mock) UnMockHelpdeskUpdateHelpdeskAgent() {
	r.mockHelpdeskUpdateHelpdeskAgent = nil
}

// UpdateHelpdeskAgentReq ...
type UpdateHelpdeskAgentReq struct {
	AgentID string `path:"agent_id" json:"-"` // 客服id, 示例值: "ou_14777d82ffef0f707de5a8c7ff2c5ebe"
	Status  *int64 `json:"status,omitempty"`  // agent status, 示例值: 1: 在线；2: 离线
}

// UpdateHelpdeskAgentResp ...
type UpdateHelpdeskAgentResp struct {
}

// updateHelpdeskAgentResp ...
type updateHelpdeskAgentResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHelpdeskAgentResp `json:"data,omitempty"`
}
