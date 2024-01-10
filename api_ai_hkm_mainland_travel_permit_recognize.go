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
	"io"
)

// RecognizeAIHkmMainlandTravelPermit 港澳居民来往内地通行证识别接口, 支持JPG/JPEG/PNG/BMP四种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/hkm_mainland_travel_permit/recognize
func (r *AIService) RecognizeAIHkmMainlandTravelPermit(ctx context.Context, request *RecognizeAIHkmMainlandTravelPermitReq, options ...MethodOptionFunc) (*RecognizeAIHkmMainlandTravelPermitResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAIHkmMainlandTravelPermit != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] AI#RecognizeAIHkmMainlandTravelPermit mock enable")
		return r.cli.mock.mockAIRecognizeAIHkmMainlandTravelPermit(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAIHkmMainlandTravelPermit",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/hkm_mainland_travel_permit/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(recognizeAIHkmMainlandTravelPermitResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAIHkmMainlandTravelPermit mock AIRecognizeAIHkmMainlandTravelPermit method
func (r *Mock) MockAIRecognizeAIHkmMainlandTravelPermit(f func(ctx context.Context, request *RecognizeAIHkmMainlandTravelPermitReq, options ...MethodOptionFunc) (*RecognizeAIHkmMainlandTravelPermitResp, *Response, error)) {
	r.mockAIRecognizeAIHkmMainlandTravelPermit = f
}

// UnMockAIRecognizeAIHkmMainlandTravelPermit un-mock AIRecognizeAIHkmMainlandTravelPermit method
func (r *Mock) UnMockAIRecognizeAIHkmMainlandTravelPermit() {
	r.mockAIRecognizeAIHkmMainlandTravelPermit = nil
}

// RecognizeAIHkmMainlandTravelPermitReq ...
type RecognizeAIHkmMainlandTravelPermitReq struct {
	File io.Reader `json:"file,omitempty"` // 识别的港澳居民来往内地通行证源文件, 示例值: file binary
}

// RecognizeAIHkmMainlandTravelPermitResp ...
type RecognizeAIHkmMainlandTravelPermitResp struct {
	HkmMainlandTravelPermit *RecognizeAIHkmMainlandTravelPermitRespHkmMainlandTravelPermit `json:"hkm_mainland_travel_permit,omitempty"` // 港澳居民来往内地通行证信息
}

// RecognizeAIHkmMainlandTravelPermitRespHkmMainlandTravelPermit ...
type RecognizeAIHkmMainlandTravelPermitRespHkmMainlandTravelPermit struct {
	Entities []*RecognizeAIHkmMainlandTravelPermitRespHkmMainlandTravelPermitEntitie `json:"entities,omitempty"` // 识别出的实体类型
}

// RecognizeAIHkmMainlandTravelPermitRespHkmMainlandTravelPermitEntitie ...
type RecognizeAIHkmMainlandTravelPermitRespHkmMainlandTravelPermitEntitie struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: full_name_cn: 中文姓名, full_name_en: 英文格式姓名, date_of_birth: 出生日期, date_of_expiry: 有效期至, card_number: 证件号码
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAIHkmMainlandTravelPermitResp ...
type recognizeAIHkmMainlandTravelPermitResp struct {
	Code int64                                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                  `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeAIHkmMainlandTravelPermitResp `json:"data,omitempty"`
}
