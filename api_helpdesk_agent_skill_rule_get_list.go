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

// GetHelpdeskAgentSkillRuleList 该接口用于获取全部客服技能。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill_rule/list
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/agent-function/agent_skill_rule/list
func (r *HelpdeskService) GetHelpdeskAgentSkillRuleList(ctx context.Context, request *GetHelpdeskAgentSkillRuleListReq, options ...MethodOptionFunc) (*GetHelpdeskAgentSkillRuleListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskAgentSkillRuleList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskAgentSkillRuleList mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskAgentSkillRuleList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskAgentSkillRuleList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/agent_skill_rules",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskAgentSkillRuleListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskAgentSkillRuleList mock HelpdeskGetHelpdeskAgentSkillRuleList method
func (r *Mock) MockHelpdeskGetHelpdeskAgentSkillRuleList(f func(ctx context.Context, request *GetHelpdeskAgentSkillRuleListReq, options ...MethodOptionFunc) (*GetHelpdeskAgentSkillRuleListResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskAgentSkillRuleList = f
}

// UnMockHelpdeskGetHelpdeskAgentSkillRuleList un-mock HelpdeskGetHelpdeskAgentSkillRuleList method
func (r *Mock) UnMockHelpdeskGetHelpdeskAgentSkillRuleList() {
	r.mockHelpdeskGetHelpdeskAgentSkillRuleList = nil
}

// GetHelpdeskAgentSkillRuleListReq ...
type GetHelpdeskAgentSkillRuleListReq struct {
}

// GetHelpdeskAgentSkillRuleListResp ...
type GetHelpdeskAgentSkillRuleListResp struct {
	Rules []*GetHelpdeskAgentSkillRuleListRespRule `json:"rules,omitempty"` // rules列表
}

// GetHelpdeskAgentSkillRuleListRespRule ...
type GetHelpdeskAgentSkillRuleListRespRule struct {
	ID              string  `json:"id,omitempty"`               // rule id, 参考[获取客服技能rules](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_skill_rule/list) 用于获取rules options
	OperatorOptions []int64 `json:"operator_options,omitempty"` // rule操作数value, [客服技能及运算符](https://open.feishu.cn/document/ukTMukTMukTM/ucDOyYjL3gjM24yN4IjN/operator-options)
	Operand         string  `json:"operand,omitempty"`          // rule 操作数的值
	Category        int64   `json:"category,omitempty"`         // rule 类型, 1-知识库, 2-工单信息, 3-用户飞书信息
	DisplayName     string  `json:"display_name,omitempty"`     // rule 名
}

// getHelpdeskAgentSkillRuleListResp ...
type getHelpdeskAgentSkillRuleListResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskAgentSkillRuleListResp `json:"data,omitempty"`
}
