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

// BatchDeleteHireEcoExamPaper 删除指定帐号的指定试卷列表, 删除不影响已创建的笔试, 删除不存在的试卷时不会报错
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/eco_exam_paper/batch_delete
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/ecological-docking/eco_exam_paper/batch_delete
func (r *HireService) BatchDeleteHireEcoExamPaper(ctx context.Context, request *BatchDeleteHireEcoExamPaperReq, options ...MethodOptionFunc) (*BatchDeleteHireEcoExamPaperResp, *Response, error) {
	if r.cli.mock.mockHireBatchDeleteHireEcoExamPaper != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Hire#BatchDeleteHireEcoExamPaper mock enable")
		return r.cli.mock.mockHireBatchDeleteHireEcoExamPaper(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "BatchDeleteHireEcoExamPaper",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/eco_exam_papers/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchDeleteHireEcoExamPaperResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireBatchDeleteHireEcoExamPaper mock HireBatchDeleteHireEcoExamPaper method
func (r *Mock) MockHireBatchDeleteHireEcoExamPaper(f func(ctx context.Context, request *BatchDeleteHireEcoExamPaperReq, options ...MethodOptionFunc) (*BatchDeleteHireEcoExamPaperResp, *Response, error)) {
	r.mockHireBatchDeleteHireEcoExamPaper = f
}

// UnMockHireBatchDeleteHireEcoExamPaper un-mock HireBatchDeleteHireEcoExamPaper method
func (r *Mock) UnMockHireBatchDeleteHireEcoExamPaper() {
	r.mockHireBatchDeleteHireEcoExamPaper = nil
}

// BatchDeleteHireEcoExamPaperReq ...
type BatchDeleteHireEcoExamPaperReq struct {
	AccountID   string   `json:"account_id,omitempty"`    // 背调账号 ID, 可在「账号绑定」事件中获取, 示例值: "7147998241542539527"
	PaperIDList []string `json:"paper_id_list,omitempty"` // 试卷 ID 列表, 示例值: ["7147998241542539526"]
}

// BatchDeleteHireEcoExamPaperResp ...
type BatchDeleteHireEcoExamPaperResp struct {
}

// batchDeleteHireEcoExamPaperResp ...
type batchDeleteHireEcoExamPaperResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *BatchDeleteHireEcoExamPaperResp `json:"data,omitempty"`
}
