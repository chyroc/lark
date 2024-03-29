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

// DeleteCoreHRProbationAssessment 删除试用期的考核结果
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/probation-assessment/delete
func (r *CoreHRService) DeleteCoreHRProbationAssessment(ctx context.Context, request *DeleteCoreHRProbationAssessmentReq, options ...MethodOptionFunc) (*DeleteCoreHRProbationAssessmentResp, *Response, error) {
	if r.cli.mock.mockCoreHRDeleteCoreHRProbationAssessment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#DeleteCoreHRProbationAssessment mock enable")
		return r.cli.mock.mockCoreHRDeleteCoreHRProbationAssessment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "DeleteCoreHRProbationAssessment",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/probation/assessments/:assessment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHRProbationAssessmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRDeleteCoreHRProbationAssessment mock CoreHRDeleteCoreHRProbationAssessment method
func (r *Mock) MockCoreHRDeleteCoreHRProbationAssessment(f func(ctx context.Context, request *DeleteCoreHRProbationAssessmentReq, options ...MethodOptionFunc) (*DeleteCoreHRProbationAssessmentResp, *Response, error)) {
	r.mockCoreHRDeleteCoreHRProbationAssessment = f
}

// UnMockCoreHRDeleteCoreHRProbationAssessment un-mock CoreHRDeleteCoreHRProbationAssessment method
func (r *Mock) UnMockCoreHRDeleteCoreHRProbationAssessment() {
	r.mockCoreHRDeleteCoreHRProbationAssessment = nil
}

// DeleteCoreHRProbationAssessmentReq ...
type DeleteCoreHRProbationAssessmentReq struct {
	AssessmentID string `path:"assessment_id" json:"-"` // 考核结果 ID, 示例值: "7140964208476371331"
}

// DeleteCoreHRProbationAssessmentResp ...
type DeleteCoreHRProbationAssessmentResp struct {
}

// deleteCoreHRProbationAssessmentResp ...
type deleteCoreHRProbationAssessmentResp struct {
	Code  int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                               `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteCoreHRProbationAssessmentResp `json:"data,omitempty"`
	Error *ErrorDetail                         `json:"error,omitempty"`
}
