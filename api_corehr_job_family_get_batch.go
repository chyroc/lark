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

// BatchGetCoreHRJobFamily 通过序列 ID 批量获取序列信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/job_family/batch_get
func (r *CoreHRService) BatchGetCoreHRJobFamily(ctx context.Context, request *BatchGetCoreHRJobFamilyReq, options ...MethodOptionFunc) (*BatchGetCoreHRJobFamilyResp, *Response, error) {
	if r.cli.mock.mockCoreHRBatchGetCoreHRJobFamily != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#BatchGetCoreHRJobFamily mock enable")
		return r.cli.mock.mockCoreHRBatchGetCoreHRJobFamily(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "BatchGetCoreHRJobFamily",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/job_families/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetCoreHRJobFamilyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRBatchGetCoreHRJobFamily mock CoreHRBatchGetCoreHRJobFamily method
func (r *Mock) MockCoreHRBatchGetCoreHRJobFamily(f func(ctx context.Context, request *BatchGetCoreHRJobFamilyReq, options ...MethodOptionFunc) (*BatchGetCoreHRJobFamilyResp, *Response, error)) {
	r.mockCoreHRBatchGetCoreHRJobFamily = f
}

// UnMockCoreHRBatchGetCoreHRJobFamily un-mock CoreHRBatchGetCoreHRJobFamily method
func (r *Mock) UnMockCoreHRBatchGetCoreHRJobFamily() {
	r.mockCoreHRBatchGetCoreHRJobFamily = nil
}

// BatchGetCoreHRJobFamilyReq ...
type BatchGetCoreHRJobFamilyReq struct {
	JobFamilyIDs []string `json:"job_family_ids,omitempty"` // 序列 ID 列表, 示例值: ["1554548"], 长度范围: `1` ～ `100`
}

// BatchGetCoreHRJobFamilyResp ...
type BatchGetCoreHRJobFamilyResp struct {
	Items []*BatchGetCoreHRJobFamilyRespItem `json:"items,omitempty"` // 查询的序列信息
}

// BatchGetCoreHRJobFamilyRespItem ...
type BatchGetCoreHRJobFamilyRespItem struct {
	JobFamilyID    string                                        `json:"job_family_id,omitempty"`   // 序列 ID
	Name           []*BatchGetCoreHRJobFamilyRespItemName        `json:"name,omitempty"`            // 名称
	Active         bool                                          `json:"active,omitempty"`          // 启用
	ParentID       string                                        `json:"parent_id,omitempty"`       // 上级序列
	EffectiveTime  string                                        `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                        `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                        `json:"code,omitempty"`            // 编码
	CustomFields   []*BatchGetCoreHRJobFamilyRespItemCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// BatchGetCoreHRJobFamilyRespItemCustomField ...
type BatchGetCoreHRJobFamilyRespItemCustomField struct {
	CustomApiName string                                          `json:"custom_api_name,omitempty"` // 自定义字段 apiname, 即自定义字段的唯一标识
	Name          *BatchGetCoreHRJobFamilyRespItemCustomFieldName `json:"name,omitempty"`            // 自定义字段名称
	Type          int64                                           `json:"type,omitempty"`            // 自定义字段类型
	Value         string                                          `json:"value,omitempty"`           // 字段值, 是 json 转义后的字符串, 根据元数据定义不同, 字段格式不同（如 123, 123.23, "true", ["id1", "id2"], "2006-01-02 15:04:05"）
}

// BatchGetCoreHRJobFamilyRespItemCustomFieldName ...
type BatchGetCoreHRJobFamilyRespItemCustomFieldName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// BatchGetCoreHRJobFamilyRespItemName ...
type BatchGetCoreHRJobFamilyRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// batchGetCoreHRJobFamilyResp ...
type batchGetCoreHRJobFamilyResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *BatchGetCoreHRJobFamilyResp `json:"data,omitempty"`
}
