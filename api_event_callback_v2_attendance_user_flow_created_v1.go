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

// EventV2AttendanceUserFlowCreatedV1 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
//
// 事件
// 用户打卡成功后, 推送该用户的打卡流水消息。
// 事件体
// |名称|类型|描述|
// |---|---|---|
// |schema|string|事件模式|
// |header|event_header|事件头|
// |-event_id|string|事件 ID|
// |-event_type|string|事件类型|
// |-create_time|string|事件创建时间戳（单位: 毫秒）|
// |-token|string|事件 Token|
// |-app_id|string|应用 ID|
// |-tenant_key|string|租户 Key|
// |event|-|事件体|
// |-employee_id|string|员工 ID|
// |-employee_no|string|员工工号|
// |-location_name|string|打卡位置名称信息|
// |-check_time|string|打卡时间, 精确到秒的时间戳|
// |-comment|string|打卡备注|
// |-record_id|string|打卡记录 ID|
// |-longitude|float|打卡经度|
// |-latitude|float|打卡纬度|
// |-ssid|string|打卡 Wi-Fi 的 SSID|
// |-bssid|string|打卡 Wi-Fi 的 MAC 地址|
// |-is_field|boolean|是否为外勤打卡|
// |-is_wifi|boolean|是否为 Wi-Fi 打卡|
// |-type|int|记录生成方式, 可用值: 【0（用户自己打卡）, 1（管理员修改）, 2（用户补卡）, 3（系统自动生成）, 4（下班免打卡）, 5（考勤机打卡）, 6（极速打卡）, 7（考勤开放平台导入）】|
// |-photo_urls|string\[\]|打卡照片列表|
// |-risk_result|int|疑似作弊打卡行为, 0: 无疑似作弊, 1: 疑似使用作弊软件, 2: 疑似使用他人的设备打卡, 3: 疑似多人使用同一设备打卡|
// 回调示例
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/event/user-attendance-records-event
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/event/user-attendance-records-event
func (r *EventCallbackService) HandlerEventV2AttendanceUserFlowCreatedV1(f EventV2AttendanceUserFlowCreatedV1Handler) {
	r.cli.eventHandler.eventV2AttendanceUserFlowCreatedV1Handler = f
}

// EventV2AttendanceUserFlowCreatedV1Handler event EventV2AttendanceUserFlowCreatedV1 handler
type EventV2AttendanceUserFlowCreatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2AttendanceUserFlowCreatedV1) (string, error)

// EventV2AttendanceUserFlowCreatedV1 ...
type EventV2AttendanceUserFlowCreatedV1 struct {
}
