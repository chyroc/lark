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

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Task_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Task

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskCollaborator(ctx, &lark.CreateTaskCollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskCollaboratorList(ctx, &lark.GetTaskCollaboratorListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskCollaborator(ctx, &lark.DeleteTaskCollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskFollower(ctx, &lark.CreateTaskFollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskFollowerList(ctx, &lark.GetTaskFollowerListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskFollower(ctx, &lark.DeleteTaskFollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskReminder(ctx, &lark.CreateTaskReminderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskReminderList(ctx, &lark.GetTaskReminderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskReminder(ctx, &lark.DeleteTaskReminderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTask(ctx, &lark.CreateTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTask(ctx, &lark.GetTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskList(ctx, &lark.GetTaskListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTask(ctx, &lark.DeleteTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateTask(ctx, &lark.UpdateTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CompleteTask(ctx, &lark.CompleteTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UncompleteTask(ctx, &lark.UncompleteTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskComment(ctx, &lark.CreateTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskComment(ctx, &lark.GetTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskComment(ctx, &lark.DeleteTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateTaskComment(ctx, &lark.UpdateTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Task

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskCreateTaskCollaborator(func(ctx context.Context, request *lark.CreateTaskCollaboratorReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskCollaboratorResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskCreateTaskCollaborator()

			_, _, err := moduleCli.CreateTaskCollaborator(ctx, &lark.CreateTaskCollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskGetTaskCollaboratorList(func(ctx context.Context, request *lark.GetTaskCollaboratorListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskCollaboratorListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskGetTaskCollaboratorList()

			_, _, err := moduleCli.GetTaskCollaboratorList(ctx, &lark.GetTaskCollaboratorListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskDeleteTaskCollaborator(func(ctx context.Context, request *lark.DeleteTaskCollaboratorReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskCollaboratorResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskDeleteTaskCollaborator()

			_, _, err := moduleCli.DeleteTaskCollaborator(ctx, &lark.DeleteTaskCollaboratorReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskCreateTaskFollower(func(ctx context.Context, request *lark.CreateTaskFollowerReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskFollowerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskCreateTaskFollower()

			_, _, err := moduleCli.CreateTaskFollower(ctx, &lark.CreateTaskFollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskGetTaskFollowerList(func(ctx context.Context, request *lark.GetTaskFollowerListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskFollowerListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskGetTaskFollowerList()

			_, _, err := moduleCli.GetTaskFollowerList(ctx, &lark.GetTaskFollowerListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskDeleteTaskFollower(func(ctx context.Context, request *lark.DeleteTaskFollowerReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskFollowerResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskDeleteTaskFollower()

			_, _, err := moduleCli.DeleteTaskFollower(ctx, &lark.DeleteTaskFollowerReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskCreateTaskReminder(func(ctx context.Context, request *lark.CreateTaskReminderReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskReminderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskCreateTaskReminder()

			_, _, err := moduleCli.CreateTaskReminder(ctx, &lark.CreateTaskReminderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskGetTaskReminderList(func(ctx context.Context, request *lark.GetTaskReminderListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskReminderListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskGetTaskReminderList()

			_, _, err := moduleCli.GetTaskReminderList(ctx, &lark.GetTaskReminderListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskDeleteTaskReminder(func(ctx context.Context, request *lark.DeleteTaskReminderReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskReminderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskDeleteTaskReminder()

			_, _, err := moduleCli.DeleteTaskReminder(ctx, &lark.DeleteTaskReminderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskCreateTask(func(ctx context.Context, request *lark.CreateTaskReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskCreateTask()

			_, _, err := moduleCli.CreateTask(ctx, &lark.CreateTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskGetTask(func(ctx context.Context, request *lark.GetTaskReq, options ...lark.MethodOptionFunc) (*lark.GetTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskGetTask()

			_, _, err := moduleCli.GetTask(ctx, &lark.GetTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskGetTaskList(func(ctx context.Context, request *lark.GetTaskListReq, options ...lark.MethodOptionFunc) (*lark.GetTaskListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskGetTaskList()

			_, _, err := moduleCli.GetTaskList(ctx, &lark.GetTaskListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskDeleteTask(func(ctx context.Context, request *lark.DeleteTaskReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskDeleteTask()

			_, _, err := moduleCli.DeleteTask(ctx, &lark.DeleteTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskUpdateTask(func(ctx context.Context, request *lark.UpdateTaskReq, options ...lark.MethodOptionFunc) (*lark.UpdateTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskUpdateTask()

			_, _, err := moduleCli.UpdateTask(ctx, &lark.UpdateTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskCompleteTask(func(ctx context.Context, request *lark.CompleteTaskReq, options ...lark.MethodOptionFunc) (*lark.CompleteTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskCompleteTask()

			_, _, err := moduleCli.CompleteTask(ctx, &lark.CompleteTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskUncompleteTask(func(ctx context.Context, request *lark.UncompleteTaskReq, options ...lark.MethodOptionFunc) (*lark.UncompleteTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskUncompleteTask()

			_, _, err := moduleCli.UncompleteTask(ctx, &lark.UncompleteTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskCreateTaskComment(func(ctx context.Context, request *lark.CreateTaskCommentReq, options ...lark.MethodOptionFunc) (*lark.CreateTaskCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskCreateTaskComment()

			_, _, err := moduleCli.CreateTaskComment(ctx, &lark.CreateTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskGetTaskComment(func(ctx context.Context, request *lark.GetTaskCommentReq, options ...lark.MethodOptionFunc) (*lark.GetTaskCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskGetTaskComment()

			_, _, err := moduleCli.GetTaskComment(ctx, &lark.GetTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskDeleteTaskComment(func(ctx context.Context, request *lark.DeleteTaskCommentReq, options ...lark.MethodOptionFunc) (*lark.DeleteTaskCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskDeleteTaskComment()

			_, _, err := moduleCli.DeleteTaskComment(ctx, &lark.DeleteTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockTaskUpdateTaskComment(func(ctx context.Context, request *lark.UpdateTaskCommentReq, options ...lark.MethodOptionFunc) (*lark.UpdateTaskCommentResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockTaskUpdateTaskComment()

			_, _, err := moduleCli.UpdateTaskComment(ctx, &lark.UpdateTaskCommentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Task

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskCollaborator(ctx, &lark.CreateTaskCollaboratorReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskCollaboratorList(ctx, &lark.GetTaskCollaboratorListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskCollaborator(ctx, &lark.DeleteTaskCollaboratorReq{
				TaskID:         "x",
				CollaboratorID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskFollower(ctx, &lark.CreateTaskFollowerReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskFollowerList(ctx, &lark.GetTaskFollowerListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskFollower(ctx, &lark.DeleteTaskFollowerReq{
				TaskID:     "x",
				FollowerID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskReminder(ctx, &lark.CreateTaskReminderReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskReminderList(ctx, &lark.GetTaskReminderListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskReminder(ctx, &lark.DeleteTaskReminderReq{
				TaskID:     "x",
				ReminderID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTask(ctx, &lark.CreateTaskReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTask(ctx, &lark.GetTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskList(ctx, &lark.GetTaskListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTask(ctx, &lark.DeleteTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateTask(ctx, &lark.UpdateTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CompleteTask(ctx, &lark.CompleteTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UncompleteTask(ctx, &lark.UncompleteTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskComment(ctx, &lark.CreateTaskCommentReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskComment(ctx, &lark.GetTaskCommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskComment(ctx, &lark.DeleteTaskCommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateTaskComment(ctx, &lark.UpdateTaskCommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Task
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskCollaborator(ctx, &lark.CreateTaskCollaboratorReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskCollaboratorList(ctx, &lark.GetTaskCollaboratorListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskCollaborator(ctx, &lark.DeleteTaskCollaboratorReq{
				TaskID:         "x",
				CollaboratorID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskFollower(ctx, &lark.CreateTaskFollowerReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskFollowerList(ctx, &lark.GetTaskFollowerListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskFollower(ctx, &lark.DeleteTaskFollowerReq{
				TaskID:     "x",
				FollowerID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskReminder(ctx, &lark.CreateTaskReminderReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskReminderList(ctx, &lark.GetTaskReminderListReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskReminder(ctx, &lark.DeleteTaskReminderReq{
				TaskID:     "x",
				ReminderID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTask(ctx, &lark.CreateTaskReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTask(ctx, &lark.GetTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskList(ctx, &lark.GetTaskListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTask(ctx, &lark.DeleteTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateTask(ctx, &lark.UpdateTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CompleteTask(ctx, &lark.CompleteTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UncompleteTask(ctx, &lark.UncompleteTaskReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateTaskComment(ctx, &lark.CreateTaskCommentReq{
				TaskID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetTaskComment(ctx, &lark.GetTaskCommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteTaskComment(ctx, &lark.DeleteTaskCommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateTaskComment(ctx, &lark.UpdateTaskCommentReq{
				TaskID:    "x",
				CommentID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})
	})
}
