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
	"encoding/json"
	"fmt"
	"strings"
)

// ApprovalWidgetType 审批挂件类型
type ApprovalWidgetType string

// ApprovalWidgetTypeInput ...
const (
	ApprovalWidgetTypeInput        ApprovalWidgetType = "input"        // 单行文本
	ApprovalWidgetTypeTextarea     ApprovalWidgetType = "textarea"     // 多行文本
	ApprovalWidgetTypeText         ApprovalWidgetType = "text"         // 说明
	ApprovalWidgetTypeNumber       ApprovalWidgetType = "number"       // 数字
	ApprovalWidgetTypeAmount       ApprovalWidgetType = "amount"       // 金额
	ApprovalWidgetTypeDate         ApprovalWidgetType = "date"         // 日期
	ApprovalWidgetTypeDateInterval ApprovalWidgetType = "dateInterval" // 日期区间
	ApprovalWidgetTypeFieldList    ApprovalWidgetType = "fieldList"    // 明细
	ApprovalWidgetTypeFormula      ApprovalWidgetType = "formula"      // 计算公式
	ApprovalWidgetTypeRadio        ApprovalWidgetType = "radio"        // 单选
	ApprovalWidgetTypeRadioV2      ApprovalWidgetType = "radioV2"      // 单选
	ApprovalWidgetTypeCheckbox     ApprovalWidgetType = "checkbox"     // 多选
	ApprovalWidgetTypeCheckboxV2   ApprovalWidgetType = "checkboxV2"   // 多选
	ApprovalWidgetTypeAttachment   ApprovalWidgetType = "attachement"  // 附件
	ApprovalWidgetTypeAttachmentV2 ApprovalWidgetType = "attachmentV2" // 附件
	ApprovalWidgetTypeDepartment   ApprovalWidgetType = "department"   // 部门
	ApprovalWidgetTypeImage        ApprovalWidgetType = "image"        // 图片
	ApprovalWidgetTypeImageV2      ApprovalWidgetType = "imageV2"      // 图片
	ApprovalWidgetTypeContact      ApprovalWidgetType = "contact"      // 联系人
	ApprovalWidgetTypeConnect      ApprovalWidgetType = "connect"      // 关联审批
	ApprovalWidgetTypeAddress      ApprovalWidgetType = "address"      // 地址
	ApprovalWidgetTypeLeaveGroup   ApprovalWidgetType = "leaveGroup"   // 请假控件组
	ApprovalWidgetTypeWorkGroup    ApprovalWidgetType = "workGroup"    // 加班控件组
	ApprovalWidgetTypeShiftGroup   ApprovalWidgetType = "shiftGroup"   // 换班控件组
	ApprovalWidgetTypeRemedyGroup  ApprovalWidgetType = "remedyGroup"  // 补卡控件组
	ApprovalWidgetTypeTripGroup    ApprovalWidgetType = "tripGroup"    // 出差控件组
	ApprovalWidgetTypeOutGroup     ApprovalWidgetType = "outGroup"     // 外出控件组
)

// ApprovalWidgetList 审批挂件列表
type ApprovalWidgetList []*ApprovalWidget

// ApprovalWidget 审批挂件
type ApprovalWidget struct {
	EnableDefaultValue bool                   `json:"enable_default_value,omitempty"` // 此控件是否启用了默认值
	ID                 string                 `json:"id,omitempty"`                   //	控件 ID
	WidgetDefaultValue string                 `json:"widget_default_value,omitempty"` //	控件的默认值
	CustomID           string                 `json:"custom_id,omitempty"`            // 控件自定义 ID
	DefaultValueType   string                 `json:"default_value_type,omitempty"`   // 控件的默认值类型
	Name               string                 `json:"name,omitempty"`                 // 控件名称
	Type               ApprovalWidgetType     `json:"type,omitempty"`                 //	控件类型
	Value              interface{}            `json:"value"`
	Option             *ApprovalWidgetOptions `json:"option,omitempty"`
	Children           []*ApprovalWidget      `json:"children,omitempty"`
	//	控件显隐条件
}

// UnmarshalJSON ...
func (r *ApprovalWidgetList) UnmarshalJSON(bs []byte) (err error) {
	if len(bs) == 0 || string(bs) == `""` {
		return nil
	}
	if strings.HasPrefix(string(bs), `"`) {
		s := ""
		if err = json.Unmarshal(bs, &s); err != nil {
			return err
		}
		bs = []byte(s)
	}

	var awList []*ApprovalWidget
	if err = json.Unmarshal(bs, &awList); err != nil {
		return err
	}
	*r = awList
	return nil
}

// MarshalJSON ...
func (r ApprovalWidgetList) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte(`""`), nil
	}
	bs, err := json.Marshal([]*ApprovalWidget(r))
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%q", bs)), nil
}

// ApprovalWidgetOptions 审批挂件参数
type ApprovalWidgetOptions struct {
	IsList  bool
	Option  *ApprovalWidgetOption
	Options []*ApprovalWidgetOption
}

// UnmarshalJSON ...
func (r *ApprovalWidgetOptions) UnmarshalJSON(bs []byte) (err error) {
	s := string(bs)
	if strings.HasPrefix(s, "[") {
		options := []*ApprovalWidgetOption{}
		if err = json.Unmarshal(bs, &options); err != nil {
			return err
		}
		r.IsList = true
		r.Options = options
	} else {
		option := new(ApprovalWidgetOption)
		if err = json.Unmarshal(bs, option); err != nil {
			return err
		}
		r.IsList = false
		r.Option = option
	}
	return nil
}

// MarshalJSON ...
func (r ApprovalWidgetOptions) MarshalJSON() ([]byte, error) {
	if r.IsList {
		return json.Marshal(r.Options)
	}
	return json.Marshal(r.Option)
}

// ApprovalWidgetOption 审批挂件参数
type ApprovalWidgetOption struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Text  string `json:"text,omitempty"`
}

// UnmarshalJSON ...
func (r *GetApprovalInstanceRespTimelineExt) UnmarshalJSON(bs []byte) (err error) {
	if len(bs) == 0 || string(bs) == `""` {
		return nil
	}
	if strings.HasPrefix(string(bs), `"`) {
		s := ""
		if err = json.Unmarshal(bs, &s); err != nil {
			return err
		}
		bs = []byte(s)
	}

	res := new(getApprovalInstanceRespTimelineExt)
	if err = json.Unmarshal(bs, res); err != nil {
		return err
	}
	*r = GetApprovalInstanceRespTimelineExt{
		UserIDList: res.UserIDList,
		OpenIDList: res.OpenIDList,
		UserID:     res.UserID,
		OpenID:     res.OpenID,
	}
	return nil
}

// MarshalJSON ...
func (r GetApprovalInstanceRespTimelineExt) MarshalJSON() ([]byte, error) {
	bs, err := json.Marshal(getApprovalInstanceRespTimelineExt{
		UserIDList: r.UserIDList,
		OpenIDList: r.OpenIDList,
		UserID:     r.UserID,
		OpenID:     r.OpenID,
	})
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%q", bs)), nil
}

type getApprovalInstanceRespTimelineExt struct {
	UserIDList []string `json:"user_id_list,omitempty"` // **type类型** - **user_id_list 含义**<br>TRANSFER - 被转交人 <br>ADD_APPROVER_BEFORE  -  被加签人<br>ADD_APPROVER -   被加签人<br>ADD_APPROVER_AFTER -   被加签人 <br>DELETE_APPROVER  - 被减签人
	OpenIDList []string `json:"open_id_list,omitempty"` // user_id_list 对应的 open id
	UserID     *string  `json:"user_id,omitempty"`      // **type类型** - **user_id 含义**<br>CC - 抄送人
	OpenID     *string  `json:"open_id,omitempty"`      // user_id 对应的 open_id
}
