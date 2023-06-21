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

// DeleteCoreHrJob 删除职务。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job/delete
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/job-management/job/delete
func (r *CoreHrService) DeleteCoreHrJob(ctx context.Context, request *DeleteCoreHrJobReq, options ...MethodOptionFunc) (*DeleteCoreHrJobResp, *Response, error) {
	if r.cli.mock.mockCoreHrDeleteCoreHrJob != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#DeleteCoreHrJob mock enable")
		return r.cli.mock.mockCoreHrDeleteCoreHrJob(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "DeleteCoreHrJob",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/jobs/:job_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHrJobResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrDeleteCoreHrJob mock CoreHrDeleteCoreHrJob method
func (r *Mock) MockCoreHrDeleteCoreHrJob(f func(ctx context.Context, request *DeleteCoreHrJobReq, options ...MethodOptionFunc) (*DeleteCoreHrJobResp, *Response, error)) {
	r.mockCoreHrDeleteCoreHrJob = f
}

// UnMockCoreHrDeleteCoreHrJob un-mock CoreHrDeleteCoreHrJob method
func (r *Mock) UnMockCoreHrDeleteCoreHrJob() {
	r.mockCoreHrDeleteCoreHrJob = nil
}

// DeleteCoreHrJobReq ...
type DeleteCoreHrJobReq struct {
	JobID string `path:"job_id" json:"-"` // 需要删除的职务 ID, 示例值: "67163716371"
}

// DeleteCoreHrJobResp ...
type DeleteCoreHrJobResp struct {
}

// deleteCoreHrJobResp ...
type deleteCoreHrJobResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCoreHrJobResp `json:"data,omitempty"`
}
