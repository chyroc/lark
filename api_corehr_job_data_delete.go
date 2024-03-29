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

// DeleteCoreHRJobData 删除人员的任职信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_data/delete
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/employee/job_data/delete
func (r *CoreHRService) DeleteCoreHRJobData(ctx context.Context, request *DeleteCoreHRJobDataReq, options ...MethodOptionFunc) (*DeleteCoreHRJobDataResp, *Response, error) {
	if r.cli.mock.mockCoreHRDeleteCoreHRJobData != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#DeleteCoreHRJobData mock enable")
		return r.cli.mock.mockCoreHRDeleteCoreHRJobData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "DeleteCoreHRJobData",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/job_datas/:job_data_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHRJobDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRDeleteCoreHRJobData mock CoreHRDeleteCoreHRJobData method
func (r *Mock) MockCoreHRDeleteCoreHRJobData(f func(ctx context.Context, request *DeleteCoreHRJobDataReq, options ...MethodOptionFunc) (*DeleteCoreHRJobDataResp, *Response, error)) {
	r.mockCoreHRDeleteCoreHRJobData = f
}

// UnMockCoreHRDeleteCoreHRJobData un-mock CoreHRDeleteCoreHRJobData method
func (r *Mock) UnMockCoreHRDeleteCoreHRJobData() {
	r.mockCoreHRDeleteCoreHRJobData = nil
}

// DeleteCoreHRJobDataReq ...
type DeleteCoreHRJobDataReq struct {
	JobDataID string `path:"job_data_id" json:"-"` // 需要删除的任职信息 ID, 示例值: "467642764726472"
}

// DeleteCoreHRJobDataResp ...
type DeleteCoreHRJobDataResp struct {
}

// deleteCoreHRJobDataResp ...
type deleteCoreHRJobDataResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteCoreHRJobDataResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
