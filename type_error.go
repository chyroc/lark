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
	"fmt"
	"strings"
)

// Error ...
type Error struct {
	Scope       string
	FuncName    string
	Code        int64
	Msg         string
	ErrorDetail *ErrorDetail
}

type ErrorDetail struct {
	FieldViolations      []*ErrorFieldViolation      `json:"field_violations"`
	PermissionViolations []*ErrorPermissionViolation `json:"permission_violations"`
	Helps                []*ErrorHelper              `json:"helps"`
	Message              string                      `json:"message"`
	LogID                string                      `json:"log_id"`
}

func (r *ErrorDetail) String() string {
	if r == nil {
		return ""
	}
	s := []string{}

	// message
	{
		if r.Message != "" {
			s = append(s, r.Message)
		}
	}
	// field validation
	{
		for _, v := range r.FieldViolations {
			s = append(s, v.Field+": "+v.Description)
		}
	}
	// permission
	{
		// action_scope_required: [docx:document, docx:document:readonly]
		for _, v := range r.groupByPermissionType() {
			s = append(s, fmt.Sprintf("%s: [%s]", v.permissionAction, strings.Join(v.permissionSubjects, ", ")))
		}
	}
	// help
	{
		// Learn more about scopes and how to add them: [docx:document,docx:document:readonly]: https://open.feishu.cn/app/xxx/auth?q=docx:document,docx:document:readonly&op_from=openapi
		for _, v := range r.Helps {
			s = append(s, v.Description+": "+v.URL)
		}
	}

	return strings.Join(s, "\n")
}

type ErrorFieldViolation struct {
	Field       string `json:"field"`
	Description string `json:"description"`
}

type ErrorPermissionViolation struct {
	Type    string `json:"type"`
	Subject string `json:"subject"`
}

type ErrorHelper struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Error ...
func (r *Error) Error() string {
	if r.Code == 0 {
		return ""
	}
	return fmt.Sprintf("request %s#%s failed: code: %d, msg: %s", r.Scope, r.FuncName, r.Code, r.Msg)
}

// NewError ...
func NewError(scope, funcName string, code int64, msg string) error {
	return &Error{
		Scope:    scope,
		FuncName: funcName,
		Code:     code,
		Msg:      msg,
	}
}

// GetErrorCode ...
func GetErrorCode(err error) int64 {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.Code
		}
		return -1
	}
	return 0
}

// GetErrorDetail ...
func GetErrorDetail(err error) *ErrorDetail {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.ErrorDetail
		}
		return nil
	}
	return nil
}

type permissionType struct {
	permissionAction   string
	permissionSubjects []string
}

func (r *ErrorDetail) groupByPermissionType() []permissionType {
	permissionActions := []string{}
	permissionSubjects := map[string][]string{}
	for _, v := range r.PermissionViolations {
		if len(permissionSubjects[v.Type]) == 0 {
			permissionActions = append(permissionActions, v.Type)
		}
		permissionSubjects[v.Type] = append(permissionSubjects[v.Type], v.Subject)
	}
	permissionTypes := []permissionType{}
	for _, action := range permissionActions {
		permissionTypes = append(permissionTypes, permissionType{
			permissionAction:   action,
			permissionSubjects: permissionSubjects[action],
		})
	}
	return permissionTypes
}
