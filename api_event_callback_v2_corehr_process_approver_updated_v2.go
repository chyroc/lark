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

// EventV2CorehrProcessApproverUpdatedV2 审批任务（approver_id 是唯一标识）, 比如一个多人会签节点, 会分别生成多人的审批任务。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=corehr&version=v2&resource=process.approver&event=updated)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/process-approver/events/updated
func (r *EventCallbackService) HandlerEventV2CorehrProcessApproverUpdatedV2(f EventV2CorehrProcessApproverUpdatedV2Handler) {
	r.cli.eventHandler.eventV2CorehrProcessApproverUpdatedV2Handler = f
}

// EventV2CorehrProcessApproverUpdatedV2Handler event EventV2CorehrProcessApproverUpdatedV2 handler
type EventV2CorehrProcessApproverUpdatedV2Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CorehrProcessApproverUpdatedV2) (string, error)

// EventV2CorehrProcessApproverUpdatedV2 ...
type EventV2CorehrProcessApproverUpdatedV2 struct {
	ProcessID  string `json:"process_id,omitempty"`  // 流程实例ID
	ApproverID string `json:"approver_id,omitempty"` // 单据ID
	Type       int64  `json:"type,omitempty"`        // 单据类型, 可选值有: 1: 审批单, 5: 表单
	Status     int64  `json:"status,omitempty"`      // 单据状态, 可选值有: 1: 待办, 3: 已完成, 2: 拒绝, 4: 取消
	BizType    string `json:"biz_type,omitempty"`    // 业务类型, 详情请查看[业务类型](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/process-approver/events/biz-type), 长度范围: `1` ～ `200` 字符
}
