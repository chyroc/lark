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

// UpdateVCRoomLevel 该接口可以用来更新某个会议室层级的信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/room_level/patch
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/room_level/patch
func (r *VCService) UpdateVCRoomLevel(ctx context.Context, request *UpdateVCRoomLevelReq, options ...MethodOptionFunc) (*UpdateVCRoomLevelResp, *Response, error) {
	if r.cli.mock.mockVCUpdateVCRoomLevel != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateVCRoomLevel mock enable")
		return r.cli.mock.mockVCUpdateVCRoomLevel(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "UpdateVCRoomLevel",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/room_levels/:room_level_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateVCRoomLevelResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCUpdateVCRoomLevel mock VCUpdateVCRoomLevel method
func (r *Mock) MockVCUpdateVCRoomLevel(f func(ctx context.Context, request *UpdateVCRoomLevelReq, options ...MethodOptionFunc) (*UpdateVCRoomLevelResp, *Response, error)) {
	r.mockVCUpdateVCRoomLevel = f
}

// UnMockVCUpdateVCRoomLevel un-mock VCUpdateVCRoomLevel method
func (r *Mock) UnMockVCUpdateVCRoomLevel() {
	r.mockVCUpdateVCRoomLevel = nil
}

// UpdateVCRoomLevelReq ...
type UpdateVCRoomLevelReq struct {
	RoomLevelID   string  `path:"room_level_id" json:"-"`    // 层级ID, 示例值: "omb_4ad1a2c7a2fbc5fc9570f38456931293"
	Name          string  `json:"name,omitempty"`            // 层级名称, 示例值: "测试层级"
	ParentID      string  `json:"parent_id,omitempty"`       // 父层级ID, 示例值: "omb_4ad1a2c7a2fbc5fc9570f38456931293"
	CustomGroupID *string `json:"custom_group_id,omitempty"` // 自定义层级ID, 示例值: "10000"
}

// UpdateVCRoomLevelResp ...
type UpdateVCRoomLevelResp struct {
}

// updateVCRoomLevelResp ...
type updateVCRoomLevelResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *UpdateVCRoomLevelResp `json:"data,omitempty"`
}
