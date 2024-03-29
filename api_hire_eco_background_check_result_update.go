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

// UpdateHireEcoBackgroundCheckResult 回传背调的最终结果, 终版报告。 回传背调结果后, 若租户未启用报告审批功能, 则背调订单状态将会变成已完成。 若启用报告审批功能, 则在管理员审批通过终版报告后, 订单状态流转为已完成。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_background_check/update_result
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_background_check/update_result
func (r *HireService) UpdateHireEcoBackgroundCheckResult(ctx context.Context, request *UpdateHireEcoBackgroundCheckResultReq, options ...MethodOptionFunc) (*UpdateHireEcoBackgroundCheckResultResp, *Response, error) {
	if r.cli.mock.mockHireUpdateHireEcoBackgroundCheckResult != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#UpdateHireEcoBackgroundCheckResult mock enable")
		return r.cli.mock.mockHireUpdateHireEcoBackgroundCheckResult(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "UpdateHireEcoBackgroundCheckResult",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/eco_background_checks/update_result",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateHireEcoBackgroundCheckResultResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireUpdateHireEcoBackgroundCheckResult mock HireUpdateHireEcoBackgroundCheckResult method
func (r *Mock) MockHireUpdateHireEcoBackgroundCheckResult(f func(ctx context.Context, request *UpdateHireEcoBackgroundCheckResultReq, options ...MethodOptionFunc) (*UpdateHireEcoBackgroundCheckResultResp, *Response, error)) {
	r.mockHireUpdateHireEcoBackgroundCheckResult = f
}

// UnMockHireUpdateHireEcoBackgroundCheckResult un-mock HireUpdateHireEcoBackgroundCheckResult method
func (r *Mock) UnMockHireUpdateHireEcoBackgroundCheckResult() {
	r.mockHireUpdateHireEcoBackgroundCheckResult = nil
}

// UpdateHireEcoBackgroundCheckResultReq ...
type UpdateHireEcoBackgroundCheckResultReq struct {
	BackgroundCheckID string                                             `json:"background_check_id,omitempty"` // 背调 ID, 示例值: "6931286400470354183"
	Result            string                                             `json:"result,omitempty"`              // 背调结果, 示例值: "无差异"
	ResultTime        string                                             `json:"result_time,omitempty"`         // 背调结果时间, 示例值: "1660123456789"
	ReportFileList    []*UpdateHireEcoBackgroundCheckResultReqReportFile `json:"report_file_list,omitempty"`    // 报告列表
}

// UpdateHireEcoBackgroundCheckResultReqReportFile ...
type UpdateHireEcoBackgroundCheckResultReqReportFile struct {
	ReportName    string `json:"report_name,omitempty"`     // 报告名称, 示例值: "阶段报告.pdf"
	ReportURL     string `json:"report_url,omitempty"`      // 报告地址；当report_url_type 为空或为 1 时需为可下载 pdf 的链接；为 2 时为预览型链接, 示例值: "https://xxxxx/xxxxxx/xxxx.pdf"
	ReportURLType *int64 `json:"report_url_type,omitempty"` // 报告地址类型；枚举值 1 或为空时为可下载的 pdf 链接, 2 为预览型链接, 示例值: 1, 可选值有: 1: 可下载的链接, 2: 外链型链接
}

// UpdateHireEcoBackgroundCheckResultResp ...
type UpdateHireEcoBackgroundCheckResultResp struct {
}

// updateHireEcoBackgroundCheckResultResp ...
type updateHireEcoBackgroundCheckResultResp struct {
	Code  int64                                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                  `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateHireEcoBackgroundCheckResultResp `json:"data,omitempty"`
	Error *ErrorDetail                            `json:"error,omitempty"`
}
