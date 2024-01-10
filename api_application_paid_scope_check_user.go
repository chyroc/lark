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

// CheckUserIsInApplicationPaidScope 当付费套餐是按人数收费 或者 限制最大使用人数时, 开放平台会引导企业管理员设置“付费功能开通范围”。  但是受开通范围限制, 部分用户就无法使用对应的付费功能。  可以通过此接口, 在付费功能点入口判断是否允许某个用户进入使用。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN
// new doc: https://open.feishu.cn/document/server-docs/application-v6/appstore-paid-info/query-a-user's-app-access
func (r *ApplicationService) CheckUserIsInApplicationPaidScope(ctx context.Context, request *CheckUserIsInApplicationPaidScopeReq, options ...MethodOptionFunc) (*CheckUserIsInApplicationPaidScopeResp, *Response, error) {
	if r.cli.mock.mockApplicationCheckUserIsInApplicationPaidScope != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#CheckUserIsInApplicationPaidScope mock enable")
		return r.cli.mock.mockApplicationCheckUserIsInApplicationPaidScope(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "CheckUserIsInApplicationPaidScope",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/pay/v1/paid_scope/check_user",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(checkUserIsInApplicationPaidScopeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationCheckUserIsInApplicationPaidScope mock ApplicationCheckUserIsInApplicationPaidScope method
func (r *Mock) MockApplicationCheckUserIsInApplicationPaidScope(f func(ctx context.Context, request *CheckUserIsInApplicationPaidScopeReq, options ...MethodOptionFunc) (*CheckUserIsInApplicationPaidScopeResp, *Response, error)) {
	r.mockApplicationCheckUserIsInApplicationPaidScope = f
}

// UnMockApplicationCheckUserIsInApplicationPaidScope un-mock ApplicationCheckUserIsInApplicationPaidScope method
func (r *Mock) UnMockApplicationCheckUserIsInApplicationPaidScope() {
	r.mockApplicationCheckUserIsInApplicationPaidScope = nil
}

// CheckUserIsInApplicationPaidScopeReq ...
type CheckUserIsInApplicationPaidScopeReq struct {
	OpenID *string `query:"open_id" json:"-"` // 用户 open_id, open_id 和 user_id 两个参数必须包含其一, 若同时传入取 open_id
	UserID *string `query:"user_id" json:"-"` // 用户 user_id, user_id 和 open_id 两个参数必须包含其一, 若同时传入取 open_id
}

// CheckUserIsInApplicationPaidScopeResp ...
type CheckUserIsInApplicationPaidScopeResp struct {
	Status          string `json:"status,omitempty"`            // 用户是否在开通范围中, "valid" -该用户在开通范围中, "not_in_scope"-该用户不在开通范围中, "no_active_license"-企业未购买任何价格方案或价格方案已过期, "exceeds_maximum_limit"-企业当前配置的付费功能开通范围人数超出限制, 需提醒管理员调整
	PricePlanID     string `json:"price_plan_id,omitempty"`     // 租户当前使用的「价格方案ID」, 对应开发者后台中「付费方案配置」中的「价格方案ID」
	IsTrial         bool   `json:"is_trial,omitempty"`          // 是否为试用版本, true-是试用版本；false-非试用版本
	ServiceStopTime string `json:"service_stop_time,omitempty"` // 租户当前有生效价格方案时表示价格方案的到期时间, 为时间unix时间戳
}

// checkUserIsInApplicationPaidScopeResp ...
type checkUserIsInApplicationPaidScopeResp struct {
	Code int64                                  `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                                 `json:"msg,omitempty"`  // 返回码的描述
	Data *CheckUserIsInApplicationPaidScopeResp `json:"data,omitempty"` // 返回的业务信息
}
