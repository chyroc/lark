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

// CropFaceVerifyImage 无源人脸比对流程, 开发者后台通过调用此接口对基准图片做规范校验及处理。
//
// ::: note
// 无源人脸比对接入需申请白名单, 接入前请联系飞书开放平台工作人员, 邮箱: open-platform@bytedance.com。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/human_authentication-v1/face/facial-image-cropping
// new doc: https://open.feishu.cn/document/server-docs/human_authentication-v1/facial-image-cropping
func (r *HumanAuthService) CropFaceVerifyImage(ctx context.Context, request *CropFaceVerifyImageReq, options ...MethodOptionFunc) (*CropFaceVerifyImageResp, *Response, error) {
	if r.cli.mock.mockHumanAuthCropFaceVerifyImage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] HumanAuth#CropFaceVerifyImage mock enable")
		return r.cli.mock.mockHumanAuthCropFaceVerifyImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "HumanAuth",
		API:                   "CropFaceVerifyImage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/face_verify/v1/crop_face_image",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(cropFaceVerifyImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHumanAuthCropFaceVerifyImage mock HumanAuthCropFaceVerifyImage method
func (r *Mock) MockHumanAuthCropFaceVerifyImage(f func(ctx context.Context, request *CropFaceVerifyImageReq, options ...MethodOptionFunc) (*CropFaceVerifyImageResp, *Response, error)) {
	r.mockHumanAuthCropFaceVerifyImage = f
}

// UnMockHumanAuthCropFaceVerifyImage un-mock HumanAuthCropFaceVerifyImage method
func (r *Mock) UnMockHumanAuthCropFaceVerifyImage() {
	r.mockHumanAuthCropFaceVerifyImage = nil
}

// CropFaceVerifyImageReq ...
type CropFaceVerifyImageReq struct {
	RawImage io.Reader `json:"raw_image,omitempty"` // 带有头像的人脸照片文件名称
}

// CropFaceVerifyImageResp ...
type CropFaceVerifyImageResp struct {
	FaceImage string `json:"face_image,omitempty"` // BASE64(裁剪后的人脸基准图片), code为0时返回
}

// cropFaceVerifyImageResp ...
type cropFaceVerifyImageResp struct {
	Code int64                    `json:"code,omitempty"` // 返回码, 非0为失败
	Msg  string                   `json:"msg,omitempty"`  // 返回信息, 返回码的描述
	Data *CropFaceVerifyImageResp `json:"data,omitempty"` // 业务数据
}
