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

// EventV2HireEHRImportTaskImportedV1 了解事件订阅的使用场景和配置流程, 请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)。
//
// 注意该事件仅通知变更相关 ID, 需要配合另外的查询接口反查实际的数据, 当导入完成后, 需调用[更新 e-HR 导入任务](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/ehr_import_task/patch) 完成任务。
// 在用户点击「导入 e-HR」后, 推送候选人信息至订阅系统。
// - 依赖权限: [更新导入 e-HR 任务]
// - 搭配使用: [获取投递信息](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/get) [更新 e-HR 导入任务](https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/ehr_import_task/patch)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/event/import-ehr
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/candidate-management/delivery-process-management/onboard/events/import-ehr
func (r *EventCallbackService) HandlerEventV2HireEHRImportTaskImportedV1(f EventV2HireEHRImportTaskImportedV1Handler) {
	r.cli.eventHandler.eventV2HireEHRImportTaskImportedV1Handler = f
}

// EventV2HireEHRImportTaskImportedV1Handler event EventV2HireEHRImportTaskImportedV1 handler
type EventV2HireEHRImportTaskImportedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2HireEHRImportTaskImportedV1) (string, error)

// EventV2HireEHRImportTaskImportedV1 ...
type EventV2HireEHRImportTaskImportedV1 struct {
	TaskID           string                 `json:"task_id,omitempty"`            // 导入任务 ID. 如: 6914551145542568199
	ApplicationID    string                 `json:"application_id,omitempty"`     // 投递 ID. 如: 6694104661676296462
	EHRDepartmentID  string                 `json:"ehr_department_id,omitempty"`  // 导入部门 ID. 如: 6694104661676263694
	EHRDepartment    map[string]interface{} `json:"ehr_department,omitempty"`     // 导入部门的飞书 ID
	EHRRequirementID string                 `json:"ehr_requirement_id,omitempty"` // 招聘需求 ID. 如: 6960663240925956636
	OperatorID       string                 `json:"operator_id,omitempty"`        // 操作人的飞书招聘 user_id. 如: 6887868781834536462
	OperatorUserID   map[string]interface{} `json:"operator_user_id,omitempty"`   // 操作人的飞书 user_id
}
