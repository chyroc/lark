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

// BatchGetMeetingRoomBuildingID 该接口用于根据租户自定义建筑 ID 查询建筑 ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQzMxYjL0MTM24CNzEjN
func (r *MeetingRoomService) BatchGetMeetingRoomBuildingID(ctx context.Context, request *BatchGetMeetingRoomBuildingIDReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomBuildingIDResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomBuildingID != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomBuildingID mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomBuildingID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomBuildingID",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/meeting_room/building/batch_get_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomBuildingIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMeetingRoomBatchGetMeetingRoomBuildingID mock MeetingRoomBatchGetMeetingRoomBuildingID method
func (r *Mock) MockMeetingRoomBatchGetMeetingRoomBuildingID(f func(ctx context.Context, request *BatchGetMeetingRoomBuildingIDReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomBuildingIDResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomBuildingID = f
}

// UnMockMeetingRoomBatchGetMeetingRoomBuildingID un-mock MeetingRoomBatchGetMeetingRoomBuildingID method
func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomBuildingID() {
	r.mockMeetingRoomBatchGetMeetingRoomBuildingID = nil
}

// BatchGetMeetingRoomBuildingIDReq ...
type BatchGetMeetingRoomBuildingIDReq struct {
	CustomBuildingIDs []string `query:"custom_building_ids" json:"-"` // 用于查询指定建筑物的租户自定义建筑ID
}

// BatchGetMeetingRoomBuildingIDResp ...
type BatchGetMeetingRoomBuildingIDResp struct {
	Buildings []*BatchGetMeetingRoomBuildingIDRespBuilding `json:"buildings,omitempty"` // 建筑列表
}

// BatchGetMeetingRoomBuildingIDRespBuilding ...
type BatchGetMeetingRoomBuildingIDRespBuilding struct {
	BuildingID       string `json:"building_id,omitempty"`        // 建筑物ID
	CustomBuildingID string `json:"custom_building_id,omitempty"` // 租户自定义建筑物ID
}

// batchGetMeetingRoomBuildingIDResp ...
type batchGetMeetingRoomBuildingIDResp struct {
	Code int64                              `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 返回码的描述, "success" 表示成功, 其他为错误提示信息
	Data *BatchGetMeetingRoomBuildingIDResp `json:"data,omitempty"` // 返回业务信息
}
