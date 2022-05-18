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
package larkext

import (
	"context"

	"github.com/chyroc/lark"
)

type Permission struct {
	client           *lark.Lark
	typ              string
	needNotification *bool
	token            string
	memberType       string // `email`, `openid`, `openchat`, `opendepartmentid`, `userid`
	memberID         string
	perm             string // 需要更新的权限，可选值有：, `view`: 可阅读, `edit`: 可编辑, `full_access`: 所有权限, 示例值："view"
	userAccessToken  string
	ops              []lark.MethodOptionFunc
}

func newPermission(client *lark.Lark, fileToken, typ string) *Permission {
	return &Permission{
		client: client,
		typ:    typ,
		token:  fileToken,
		perm:   "view",
	}
}

func (r *Permission) WithUserAccessToken(userAccessToken string) *Permission {
	r.userAccessToken = userAccessToken
	r.ops = []lark.MethodOptionFunc{lark.WithUserAccessToken(userAccessToken)}
	return r
}

func (r *Permission) ToEmail(email string) *Permission {
	r.memberType = "email"
	r.memberID = email
	return r
}

func (r *Permission) ToOpen(openID string) *Permission {
	r.memberType = "openid"
	r.memberID = openID
	return r
}

func (r *Permission) ToOpenChat(openChatID string) *Permission {
	r.memberType = "openchat"
	r.memberID = openChatID
	return r
}

func (r *Permission) ToOpenDepartment(openDepartmentID string) *Permission {
	r.memberType = "opendepartmentid"
	r.memberID = openDepartmentID
	return r
}

func (r *Permission) ToUserID(userID string) *Permission {
	r.memberType = "userid"
	r.memberID = userID
	return r
}

func (r *Permission) NeedNotification(needNotification bool) *Permission {
	r.needNotification = &needNotification
	return r
}

func (r *Permission) Perm(perm string) *Permission {
	r.perm = perm
	return r
}

func (r *Permission) ViewPerm() *Permission {
	r.perm = "view"
	return r
}

func (r *Permission) EditPerm() *Permission {
	r.perm = "edit"
	return r
}

func (r *Permission) FullAccessPerm() *Permission {
	r.perm = "full_access"
	return r
}

// Create 添加权限
func (r *Permission) GetList(ctx context.Context) ([]*lark.GetDriveMemberPermissionListRespMember, error) {
	resp, _, err := r.client.Drive.GetDriveMemberPermissionList(ctx, &lark.GetDriveMemberPermissionListReq{
		Type:  r.typ,
		Token: r.token,
	}, r.ops...)
	if err != nil {
		return nil, err
	}
	return resp.Members, nil
}

// TransferOwner 转移 owner
func (r *Permission) TransferOwner(ctx context.Context, removeOldOwner, cancelNotify *bool) error {
	_, _, err := r.client.Drive.TransferDriveMemberPermission(ctx, &lark.TransferDriveMemberPermissionReq{
		Token: r.token,
		Type:  r.typ,
		Owner: &lark.TransferDriveMemberPermissionReqOwner{
			MemberType: r.memberType,
			MemberID:   r.memberID,
		},
		RemoveOldOwner: removeOldOwner,
		CancelNotify:   cancelNotify,
	}, r.ops...)
	return err
}

// Create 添加权限
func (r *Permission) Create(ctx context.Context) (*lark.CreateDriveMemberPermissionRespMember, error) {
	resp, _, err := r.client.Drive.CreateDriveMemberPermission(ctx, &lark.CreateDriveMemberPermissionReq{
		Type:             r.typ,
		Token:            r.token,
		MemberType:       r.memberType,
		MemberID:         r.memberID,
		NeedNotification: r.needNotification,
		Perm:             r.perm,
	}, r.ops...)
	if err != nil {
		return nil, err
	}
	return resp.Member, nil
}

// Delete 删除权限
func (r *Permission) Delete(ctx context.Context) error {
	_, _, err := r.client.Drive.DeleteDriveMemberPermission(ctx, &lark.DeleteDriveMemberPermissionReq{
		Type:       r.typ,
		Token:      r.token,
		MemberType: r.memberType,
		MemberID:   r.memberID,
	}, r.ops...)
	return err
}

// Update 更新权限
func (r *Permission) Update(ctx context.Context) (*lark.UpdateDriveMemberPermissionRespMember, error) {
	resp, _, err := r.client.Drive.UpdateDriveMemberPermission(ctx, &lark.UpdateDriveMemberPermissionReq{
		Type:             r.typ,
		Token:            r.token,
		MemberType:       r.memberType,
		MemberID:         r.memberID,
		NeedNotification: r.needNotification,
		Perm:             r.perm,
	}, r.ops...)
	if err != nil {
		return nil, err
	}
	return resp.Member, nil
}

// Check 检查权限
func (r *Permission) Check(ctx context.Context) (bool, error) {
	resp, _, err := r.client.Drive.CheckDriveMemberPermission(ctx, &lark.CheckDriveMemberPermissionReq{
		Type:  r.typ,
		Token: r.token,
		Perm:  r.perm,
	}, r.ops...)
	if err != nil {
		return false, err
	}
	return resp.IsPermitted, nil
}

// GetPublic 获取公共权限
func (r *Permission) GetPublic(ctx context.Context) (*lark.GetDrivePublicPermissionRespPermissionPublic, error) {
	resp, _, err := r.client.Drive.GetDrivePublicPermission(ctx, &lark.GetDrivePublicPermissionReq{
		Type:  r.typ,
		Token: r.token,
	}, r.ops...)
	if err != nil {
		return nil, err
	}
	return resp.PermissionPublic, nil
}

// UpdatePublic 更新公共权限
func (r *Permission) UpdatePublic(ctx context.Context, req *lark.UpdateDrivePublicPermissionReq) error {
	_, _, err := r.client.Drive.UpdateDrivePublicPermission(ctx, &lark.UpdateDrivePublicPermissionReq{
		Type:            r.typ,
		Token:           r.token,
		ExternalAccess:  req.ExternalAccess,
		SecurityEntity:  req.SecurityEntity,
		CommentEntity:   req.CommentEntity,
		ShareEntity:     req.ShareEntity,
		LinkShareEntity: req.LinkShareEntity,
		InviteExternal:  req.InviteExternal,
	}, r.ops...)
	return err
}
