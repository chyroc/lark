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

// WithdrawHireReferralAccount 支持通过账号 ID 全额提取内推账号下的积分或现金奖励。调用前, 请确认已完成[「注册外部系统内推账户」](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral_account/create)并获取到账号 ID。提现后, 内推人的对应积分或现金余额将变为 0, 扣减后对应奖励将在招聘系统同步标记为「已发放」
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/referral_account/withdraw
func (r *HireService) WithdrawHireReferralAccount(ctx context.Context, request *WithdrawHireReferralAccountReq, options ...MethodOptionFunc) (*WithdrawHireReferralAccountResp, *Response, error) {
	if r.cli.mock.mockHireWithdrawHireReferralAccount != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#WithdrawHireReferralAccount mock enable")
		return r.cli.mock.mockHireWithdrawHireReferralAccount(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "WithdrawHireReferralAccount",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/referral_account/:referral_account_id/withdraw",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(withdrawHireReferralAccountResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireWithdrawHireReferralAccount mock HireWithdrawHireReferralAccount method
func (r *Mock) MockHireWithdrawHireReferralAccount(f func(ctx context.Context, request *WithdrawHireReferralAccountReq, options ...MethodOptionFunc) (*WithdrawHireReferralAccountResp, *Response, error)) {
	r.mockHireWithdrawHireReferralAccount = f
}

// UnMockHireWithdrawHireReferralAccount un-mock HireWithdrawHireReferralAccount method
func (r *Mock) UnMockHireWithdrawHireReferralAccount() {
	r.mockHireWithdrawHireReferralAccount = nil
}

// WithdrawHireReferralAccountReq ...
type WithdrawHireReferralAccountReq struct {
	ReferralAccountID string  `path:"referral_account_id" json:"-"`  // 账户ID, 示例值: "6942778198054125570"
	WithdrawBonusType []int64 `json:"withdraw_bonus_type,omitempty"` // 请求提现的奖励类型, 示例值: [1], 可选值有: 1: 积分
	ExternalOrderID   *string `json:"external_order_id,omitempty"`   // 提现单ID, 请求时由请求方提供, 后续关于本次提现操作的交互都以此提现单ID为标识进行, 需要保证唯一, 用于保证提现的幂等性, 传入重复ID会返回对应提现单提取的金额明细, 示例值: "6942778198054125570"
}

// WithdrawHireReferralAccountResp ...
type WithdrawHireReferralAccountResp struct {
	ExternalOrderID   string                                            `json:"external_order_id,omitempty"`  // 请求时传入的提现单ID
	TransTime         string                                            `json:"trans_time,omitempty"`         // 交易时间戳, 需要保存, 用于统一交易时间, 方便对账
	WithdrawalDetails *WithdrawHireReferralAccountRespWithdrawalDetails `json:"withdrawal_details,omitempty"` // 本次提现金额明细
}

// WithdrawHireReferralAccountRespWithdrawalDetails ...
type WithdrawHireReferralAccountRespWithdrawalDetails struct {
	PointBonus int64 `json:"point_bonus,omitempty"` // 积分奖励
}

// withdrawHireReferralAccountResp ...
type withdrawHireReferralAccountResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *WithdrawHireReferralAccountResp `json:"data,omitempty"`
}
