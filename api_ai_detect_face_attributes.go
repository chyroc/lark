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

// DetectFaceAttributes 检测图片中的人脸属性和质量等信息
//
// 注意：返回值为 -1 表示该功能还暂未实现
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/face_detection-v1/image/detect_face_attributes
func (r *AIService) DetectFaceAttributes(ctx context.Context, request *DetectFaceAttributesReq, options ...MethodOptionFunc) (*DetectFaceAttributesResp, *Response, error) {
	if r.cli.mock.mockAIDetectFaceAttributes != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#DetectFaceAttributes mock enable")
		return r.cli.mock.mockAIDetectFaceAttributes(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "DetectFaceAttributes",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/face_detection/v1/image/detect_face_attributes",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(detectFaceAttributesResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIDetectFaceAttributes mock AIDetectFaceAttributes method
func (r *Mock) MockAIDetectFaceAttributes(f func(ctx context.Context, request *DetectFaceAttributesReq, options ...MethodOptionFunc) (*DetectFaceAttributesResp, *Response, error)) {
	r.mockAIDetectFaceAttributes = f
}

// UnMockAIDetectFaceAttributes un-mock AIDetectFaceAttributes method
func (r *Mock) UnMockAIDetectFaceAttributes() {
	r.mockAIDetectFaceAttributes = nil
}

// DetectFaceAttributesReq ...
type DetectFaceAttributesReq struct {
	Image *string `json:"image,omitempty"` // 图片 base64 数据, 示例值："图片 base64 后的字符串"
}

// detectFaceAttributesResp ...
type detectFaceAttributesResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *DetectFaceAttributesResp `json:"data,omitempty"`
}

// DetectFaceAttributesResp ...
type DetectFaceAttributesResp struct {
	ImageInfo *DetectFaceAttributesRespImageInfo  `json:"image_info,omitempty"` // 图片信息
	FaceInfos []*DetectFaceAttributesRespFaceInfo `json:"face_infos,omitempty"` // 人脸信息列表
}

// DetectFaceAttributesRespImageInfo ...
type DetectFaceAttributesRespImageInfo struct {
	Width  int64 `json:"width,omitempty"`  // 图片的宽度
	Height int64 `json:"height,omitempty"` // 图片的高度
}

// DetectFaceAttributesRespFaceInfo ...
type DetectFaceAttributesRespFaceInfo struct {
	Position  *DetectFaceAttributesRespFaceInfoPosition  `json:"position,omitempty"`  // 人脸位置信息
	Attribute *DetectFaceAttributesRespFaceInfoAttribute `json:"attribute,omitempty"` // 人脸属性信息
	Quality   *DetectFaceAttributesRespFaceInfoQuality   `json:"quality,omitempty"`   // 人脸质量信息
}

// DetectFaceAttributesRespFaceInfoPosition ...
type DetectFaceAttributesRespFaceInfoPosition struct {
	UpperLeft  *DetectFaceAttributesRespFaceInfoPositionUpperLeft  `json:"upper_left,omitempty"`  // 人脸框的左上角坐标
	LowerRight *DetectFaceAttributesRespFaceInfoPositionLowerRight `json:"lower_right,omitempty"` // 人脸框的右下角坐标
}

// DetectFaceAttributesRespFaceInfoPositionUpperLeft ...
type DetectFaceAttributesRespFaceInfoPositionUpperLeft struct {
	X float64 `json:"x,omitempty"` // 横轴坐标
	Y float64 `json:"y,omitempty"` // 纵轴坐标
}

// DetectFaceAttributesRespFaceInfoPositionLowerRight ...
type DetectFaceAttributesRespFaceInfoPositionLowerRight struct {
	X float64 `json:"x,omitempty"` // 横轴坐标
	Y float64 `json:"y,omitempty"` // 纵轴坐标
}

// DetectFaceAttributesRespFaceInfoAttribute ...
type DetectFaceAttributesRespFaceInfoAttribute struct {
	Gender  *DetectFaceAttributesRespFaceInfoAttributeGender  `json:"gender,omitempty"`  // 性别信息：0 男性，1 女性
	Age     int64                                             `json:"age,omitempty"`     // 年龄大小
	Emotion *DetectFaceAttributesRespFaceInfoAttributeEmotion `json:"emotion,omitempty"` // 情绪：0 自然, 1 高兴，2 惊讶，3 害怕，4 悲伤，5 生气, 6 厌恶
	Beauty  int64                                             `json:"beauty,omitempty"`  // 颜值打分：[0, 100]
	Pose    *DetectFaceAttributesRespFaceInfoAttributePose    `json:"pose,omitempty"`    // 人脸姿态
	Hat     *DetectFaceAttributesRespFaceInfoAttributeHat     `json:"hat,omitempty"`     // 帽子：0 未戴帽子，1 戴帽子
	Glass   *DetectFaceAttributesRespFaceInfoAttributeGlass   `json:"glass,omitempty"`   // 眼镜：0 未戴眼镜，1 戴眼镜
	Mask    *DetectFaceAttributesRespFaceInfoAttributeMask    `json:"mask,omitempty"`    // 口罩：0 未戴口罩，1 戴口罩
}

// DetectFaceAttributesRespFaceInfoAttributeGender ...
type DetectFaceAttributesRespFaceInfoAttributeGender struct {
	Type        int64   `json:"type,omitempty"`        // 属性
	Probability float64 `json:"probability,omitempty"` // 识别置信度，[0, 1]，代表判断正确的概率
}

// DetectFaceAttributesRespFaceInfoAttributeEmotion ...
type DetectFaceAttributesRespFaceInfoAttributeEmotion struct {
	Type        int64   `json:"type,omitempty"`        // 属性
	Probability float64 `json:"probability,omitempty"` // 识别置信度，[0, 1]，代表判断正确的概率
}

// DetectFaceAttributesRespFaceInfoAttributePose ...
type DetectFaceAttributesRespFaceInfoAttributePose struct {
	Pitch int64 `json:"pitch,omitempty"` // 脸部上下偏移 [-90, 90]
	Yaw   int64 `json:"yaw,omitempty"`   // 脸部左右偏移 [-90, 90]
	Roll  int64 `json:"roll,omitempty"`  // 平面旋转 [-90, 90]
}

// DetectFaceAttributesRespFaceInfoAttributeHat ...
type DetectFaceAttributesRespFaceInfoAttributeHat struct {
	Type        int64   `json:"type,omitempty"`        // 属性
	Probability float64 `json:"probability,omitempty"` // 识别置信度，[0, 1]，代表判断正确的概率
}

// DetectFaceAttributesRespFaceInfoAttributeGlass ...
type DetectFaceAttributesRespFaceInfoAttributeGlass struct {
	Type        int64   `json:"type,omitempty"`        // 属性
	Probability float64 `json:"probability,omitempty"` // 识别置信度，[0, 1]，代表判断正确的概率
}

// DetectFaceAttributesRespFaceInfoAttributeMask ...
type DetectFaceAttributesRespFaceInfoAttributeMask struct {
	Type        int64   `json:"type,omitempty"`        // 属性
	Probability float64 `json:"probability,omitempty"` // 识别置信度，[0, 1]，代表判断正确的概率
}

// DetectFaceAttributesRespFaceInfoQuality ...
type DetectFaceAttributesRespFaceInfoQuality struct {
	Sharpness  float64                                         `json:"sharpness,omitempty"`  // 清晰度，值越高越清晰
	Brightness float64                                         `json:"brightness,omitempty"` // 亮度
	Occlude    *DetectFaceAttributesRespFaceInfoQualityOcclude `json:"occlude,omitempty"`    // 面部遮挡属性
}

// DetectFaceAttributesRespFaceInfoQualityOcclude ...
type DetectFaceAttributesRespFaceInfoQualityOcclude struct {
	Eyebrow  float64 `json:"eyebrow,omitempty"`   // 眉毛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Nose     float64 `json:"nose,omitempty"`      // 鼻子被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Cheek    float64 `json:"cheek,omitempty"`     // 脸颊被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Mouth    float64 `json:"mouth,omitempty"`     // 嘴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	Chin     float64 `json:"chin,omitempty"`      // 下巴被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	LeftEye  float64 `json:"left_eye,omitempty"`  // 左眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
	RightEye float64 `json:"right_eye,omitempty"` // 右眼睛被遮挡情况：[0, 1] 值越大被遮挡的概率越高
}