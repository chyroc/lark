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

// RecognizeAITaxiInvoice 出租车发票识别接口, 支持JPG/JPEG/PNG/PDF/OFD五种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/taxi_invoice/recognize
func (r *AIService) RecognizeAITaxiInvoice(ctx context.Context, request *RecognizeAITaxiInvoiceReq, options ...MethodOptionFunc) (*RecognizeAITaxiInvoiceResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAITaxiInvoice != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] AI#RecognizeAITaxiInvoice mock enable")
		return r.cli.mock.mockAIRecognizeAITaxiInvoice(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAITaxiInvoice",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/taxi_invoice/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(recognizeAITaxiInvoiceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAITaxiInvoice mock AIRecognizeAITaxiInvoice method
func (r *Mock) MockAIRecognizeAITaxiInvoice(f func(ctx context.Context, request *RecognizeAITaxiInvoiceReq, options ...MethodOptionFunc) (*RecognizeAITaxiInvoiceResp, *Response, error)) {
	r.mockAIRecognizeAITaxiInvoice = f
}

// UnMockAIRecognizeAITaxiInvoice un-mock AIRecognizeAITaxiInvoice method
func (r *Mock) UnMockAIRecognizeAITaxiInvoice() {
	r.mockAIRecognizeAITaxiInvoice = nil
}

// RecognizeAITaxiInvoiceReq ...
type RecognizeAITaxiInvoiceReq struct {
	File io.Reader `json:"file,omitempty"` // 识别的出租车票源文件, 示例值: file binary
}

// RecognizeAITaxiInvoiceResp ...
type RecognizeAITaxiInvoiceResp struct {
	TaxiInvoices []*RecognizeAITaxiInvoiceRespTaxiInvoice `json:"taxi_invoices,omitempty"` // 出租车票信息
}

// RecognizeAITaxiInvoiceRespTaxiInvoice ...
type RecognizeAITaxiInvoiceRespTaxiInvoice struct {
	Entities []*RecognizeAITaxiInvoiceRespTaxiInvoiceEntitie `json:"entities,omitempty"` // 识别出的实体类型
}

// RecognizeAITaxiInvoiceRespTaxiInvoiceEntitie ...
type RecognizeAITaxiInvoiceRespTaxiInvoiceEntitie struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: car_number: 车号, start_time: 上车时间, end_time: 下车时间, distance: 里程, start_date: 日期, total_amount: 出租车价格
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAITaxiInvoiceResp ...
type recognizeAITaxiInvoiceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeAITaxiInvoiceResp `json:"data,omitempty"`
}
