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

// EventV2TaskTaskCommentUpdatedV1 当 APP 创建的任务评论信息发生变更时触发此事件, 包括任务评论的创建、回复、更新、删除。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=task&version=v1&resource=task.comment&event=updated)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-comment/events/updated
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task-comment/events/updated
func (r *EventCallbackService) HandlerEventV2TaskTaskCommentUpdatedV1(f EventV2TaskTaskCommentUpdatedV1Handler) {
	r.cli.eventHandler.eventV2TaskTaskCommentUpdatedV1Handler = f
}

// EventV2TaskTaskCommentUpdatedV1Handler event EventV2TaskTaskCommentUpdatedV1 handler
type EventV2TaskTaskCommentUpdatedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2TaskTaskCommentUpdatedV1) (string, error)

// EventV2TaskTaskCommentUpdatedV1 ...
type EventV2TaskTaskCommentUpdatedV1 struct {
	TaskID    string `json:"task_id,omitempty"`    // 任务ID
	CommentID string `json:"comment_id,omitempty"` // 任务评论ID
	ParentID  string `json:"parent_id,omitempty"`  // 任务评论父ID
	ObjType   int64  `json:"obj_type,omitempty"`   // 通知类型（1: 创建评论, 2: 回复评论, 3: 更新评论, 4: 删除评论）
}
