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

// ExportVCMeetingList 导出会议明细, 具体权限要求请参考资源介绍。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/meeting_list
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/export/meeting_list
func (r *VCService) ExportVCMeetingList(ctx context.Context, request *ExportVCMeetingListReq, options ...MethodOptionFunc) (*ExportVCMeetingListResp, *Response, error) {
	if r.cli.mock.mockVCExportVCMeetingList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#ExportVCMeetingList mock enable")
		return r.cli.mock.mockVCExportVCMeetingList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "ExportVCMeetingList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/exports/meeting_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(exportVCMeetingListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCExportVCMeetingList mock VCExportVCMeetingList method
func (r *Mock) MockVCExportVCMeetingList(f func(ctx context.Context, request *ExportVCMeetingListReq, options ...MethodOptionFunc) (*ExportVCMeetingListResp, *Response, error)) {
	r.mockVCExportVCMeetingList = f
}

// UnMockVCExportVCMeetingList un-mock VCExportVCMeetingList method
func (r *Mock) UnMockVCExportVCMeetingList() {
	r.mockVCExportVCMeetingList = nil
}

// ExportVCMeetingListReq ...
type ExportVCMeetingListReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	StartTime  string  `json:"start_time,omitempty"`   // 查询开始时间（unix时间, 单位sec）, 示例值: "1655276858"
	EndTime    string  `json:"end_time,omitempty"`     // 查询结束时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingNo  *string `json:"meeting_no,omitempty"`   // 按9位会议号筛选（最多一个筛选条件）, 示例值: "123456789"
	UserID     *string `json:"user_id,omitempty"`      // 按参会Lark用户筛选（最多一个筛选条件）, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	RoomID     *string `json:"room_id,omitempty"`      // 按参会Rooms筛选（最多一个筛选条件）, 示例值: "omm_eada1d61a550955240c28757e7dec3af"
}

// ExportVCMeetingListResp ...
type ExportVCMeetingListResp struct {
	TaskID string `json:"task_id,omitempty"` // 任务id
}

// exportVCMeetingListResp ...
type exportVCMeetingListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *ExportVCMeetingListResp `json:"data,omitempty"`
}
