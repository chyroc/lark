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

// CreateHireExternalBackgroundCheck 导入来自其他系统的背调信息, 创建为外部背调。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/external_background_check/create
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/get-candidates/import-external-system-information/create-2
func (r *HireService) CreateHireExternalBackgroundCheck(ctx context.Context, request *CreateHireExternalBackgroundCheckReq, options ...MethodOptionFunc) (*CreateHireExternalBackgroundCheckResp, *Response, error) {
	if r.cli.mock.mockHireCreateHireExternalBackgroundCheck != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#CreateHireExternalBackgroundCheck mock enable")
		return r.cli.mock.mockHireCreateHireExternalBackgroundCheck(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "CreateHireExternalBackgroundCheck",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/external_background_checks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createHireExternalBackgroundCheckResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireCreateHireExternalBackgroundCheck mock HireCreateHireExternalBackgroundCheck method
func (r *Mock) MockHireCreateHireExternalBackgroundCheck(f func(ctx context.Context, request *CreateHireExternalBackgroundCheckReq, options ...MethodOptionFunc) (*CreateHireExternalBackgroundCheckResp, *Response, error)) {
	r.mockHireCreateHireExternalBackgroundCheck = f
}

// UnMockHireCreateHireExternalBackgroundCheck un-mock HireCreateHireExternalBackgroundCheck method
func (r *Mock) UnMockHireCreateHireExternalBackgroundCheck() {
	r.mockHireCreateHireExternalBackgroundCheck = nil
}

// CreateHireExternalBackgroundCheckReq ...
type CreateHireExternalBackgroundCheckReq struct {
	ExternalID            *string  `json:"external_id,omitempty"`             // 外部系统背调主键 （仅用于幂等）, 示例值: "123"
	ExternalApplicationID string   `json:"external_application_id,omitempty"` // 外部投递 ID, 示例值: "1234111"
	Date                  *int64   `json:"date,omitempty"`                    // 背调日期, 示例值: 1626602069393
	Name                  *string  `json:"name,omitempty"`                    // 背调名字, 示例值: "测试.pdf"
	Result                *string  `json:"result,omitempty"`                  // 背调结果, 示例值: "1"
	AttachmentIDList      []string `json:"attachment_id_list,omitempty"`      // 背调附件ID列表, 示例值: ["6989181065243969836"]
}

// CreateHireExternalBackgroundCheckResp ...
type CreateHireExternalBackgroundCheckResp struct {
	ExternalBackgroundCheck *CreateHireExternalBackgroundCheckRespExternalBackgroundCheck `json:"external_background_check,omitempty"` // 外部背调信息
}

// CreateHireExternalBackgroundCheckRespExternalBackgroundCheck ...
type CreateHireExternalBackgroundCheckRespExternalBackgroundCheck struct {
	ID                    string                                                                    `json:"id,omitempty"`                      // 外部背调 ID
	ExternalID            string                                                                    `json:"external_id,omitempty"`             // 外部系统背调主键 （仅用于幂等）
	ExternalApplicationID string                                                                    `json:"external_application_id,omitempty"` // 外部投递 ID
	Date                  int64                                                                     `json:"date,omitempty"`                    // 背调日期
	Name                  string                                                                    `json:"name,omitempty"`                    // 背调名字
	Result                string                                                                    `json:"result,omitempty"`                  // 背调结果
	AttachmentIDList      []string                                                                  `json:"attachment_id_list,omitempty"`      // 背调附件ID列表
	AttachmentList        []*CreateHireExternalBackgroundCheckRespExternalBackgroundCheckAttachment `json:"attachment_list,omitempty"`         // 背调附件
}

// CreateHireExternalBackgroundCheckRespExternalBackgroundCheckAttachment ...
type CreateHireExternalBackgroundCheckRespExternalBackgroundCheckAttachment struct {
	ID   string `json:"id,omitempty"`   // 附件 ID
	Name string `json:"name,omitempty"` // 附件名字
	Size int64  `json:"size,omitempty"` // 附件大小
}

// createHireExternalBackgroundCheckResp ...
type createHireExternalBackgroundCheckResp struct {
	Code int64                                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateHireExternalBackgroundCheckResp `json:"data,omitempty"`
}
