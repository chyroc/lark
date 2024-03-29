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

// RecognizeAIDrivingLicense 驾驶证识别接口, 支持JPG/JPEG/PNG/BMP四种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/driving_license/recognize
func (r *AIService) RecognizeAIDrivingLicense(ctx context.Context, request *RecognizeAIDrivingLicenseReq, options ...MethodOptionFunc) (*RecognizeAIDrivingLicenseResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAIDrivingLicense != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] AI#RecognizeAIDrivingLicense mock enable")
		return r.cli.mock.mockAIRecognizeAIDrivingLicense(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAIDrivingLicense",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/driving_license/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(recognizeAIDrivingLicenseResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAIDrivingLicense mock AIRecognizeAIDrivingLicense method
func (r *Mock) MockAIRecognizeAIDrivingLicense(f func(ctx context.Context, request *RecognizeAIDrivingLicenseReq, options ...MethodOptionFunc) (*RecognizeAIDrivingLicenseResp, *Response, error)) {
	r.mockAIRecognizeAIDrivingLicense = f
}

// UnMockAIRecognizeAIDrivingLicense un-mock AIRecognizeAIDrivingLicense method
func (r *Mock) UnMockAIRecognizeAIDrivingLicense() {
	r.mockAIRecognizeAIDrivingLicense = nil
}

// RecognizeAIDrivingLicenseReq ...
type RecognizeAIDrivingLicenseReq struct {
	File io.Reader `json:"file,omitempty"` // 识别的驾驶证源文件, 示例值: file binary
}

// RecognizeAIDrivingLicenseResp ...
type RecognizeAIDrivingLicenseResp struct {
	DrivingLicense *RecognizeAIDrivingLicenseRespDrivingLicense `json:"driving_license,omitempty"` // 驾驶证信息
}

// RecognizeAIDrivingLicenseRespDrivingLicense ...
type RecognizeAIDrivingLicenseRespDrivingLicense struct {
	Entities []*RecognizeAIDrivingLicenseRespDrivingLicenseEntitie `json:"entities,omitempty"` // 识别出的实体类型
}

// RecognizeAIDrivingLicenseRespDrivingLicenseEntitie ...
type RecognizeAIDrivingLicenseRespDrivingLicenseEntitie struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: id_number: 证号, name: 姓名, sex: 性别, nationality: 国籍, address: 住址, date_of_birth: 出生日期, date_of_first_issue: 初次领证日期, class: 准驾车型, valid_begin: 有效期起, valid_end: 有效期止, license_issuing_authority: 发证机关, document_id: 档案编号, record: 记录, id_photo_location: 相片位置
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAIDrivingLicenseResp ...
type recognizeAIDrivingLicenseResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *RecognizeAIDrivingLicenseResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
