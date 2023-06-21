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

// GetVCDailyReport 获取一段时间内组织的每日会议使用报告。
//
// 支持最近90天内的数据查询
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/report/get_daily
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/report/get_daily
func (r *VCService) GetVCDailyReport(ctx context.Context, request *GetVCDailyReportReq, options ...MethodOptionFunc) (*GetVCDailyReportResp, *Response, error) {
	if r.cli.mock.mockVCGetVCDailyReport != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#GetVCDailyReport mock enable")
		return r.cli.mock.mockVCGetVCDailyReport(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "GetVCDailyReport",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reports/get_daily",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getVCDailyReportResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCGetVCDailyReport mock VCGetVCDailyReport method
func (r *Mock) MockVCGetVCDailyReport(f func(ctx context.Context, request *GetVCDailyReportReq, options ...MethodOptionFunc) (*GetVCDailyReportResp, *Response, error)) {
	r.mockVCGetVCDailyReport = f
}

// UnMockVCGetVCDailyReport un-mock VCGetVCDailyReport method
func (r *Mock) UnMockVCGetVCDailyReport() {
	r.mockVCGetVCDailyReport = nil
}

// GetVCDailyReportReq ...
type GetVCDailyReportReq struct {
	StartTime string `query:"start_time" json:"-"` // 开始时间（unix时间, 单位sec）, 示例值: "1608888867"
	EndTime   string `query:"end_time" json:"-"`   // 结束时间（unix时间, 单位sec）, 示例值: "1608888966"
}

// GetVCDailyReportResp ...
type GetVCDailyReportResp struct {
	MeetingReport *GetVCDailyReportRespMeetingReport `json:"meeting_report,omitempty"` // 会议报告
}

// GetVCDailyReportRespMeetingReport ...
type GetVCDailyReportRespMeetingReport struct {
	TotalMeetingCount     int64                                           `json:"total_meeting_count,omitempty"`     // 总会议数量
	TotalMeetingDuration  int64                                           `json:"total_meeting_duration,omitempty"`  // 总会议时长（单位sec）
	TotalParticipantCount int64                                           `json:"total_participant_count,omitempty"` // 总参会人数
	DailyReport           []*GetVCDailyReportRespMeetingReportDailyReport `json:"daily_report,omitempty"`            // 每日会议报告列表
}

// GetVCDailyReportRespMeetingReportDailyReport ...
type GetVCDailyReportRespMeetingReportDailyReport struct {
	Date             string `json:"date,omitempty"`              // 日期（unix时间, 单位sec）
	MeetingCount     string `json:"meeting_count,omitempty"`     // 会议数量
	MeetingDuration  string `json:"meeting_duration,omitempty"`  // 会议时长（单位sec）
	ParticipantCount string `json:"participant_count,omitempty"` // 参会人数
}

// getVCDailyReportResp ...
type getVCDailyReportResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetVCDailyReportResp `json:"data,omitempty"`
}
