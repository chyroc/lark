package lark

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ApprovalWidgetType string

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

type ApprovalWidgetList []*ApprovalWidget

type ApprovalWidget struct {
	ID       string                 `json:"id,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Type     ApprovalWidgetType     `json:"type,omitempty"`
	Value    interface{}            `json:"value"`
	Option   *ApprovalWidgetOptions `json:"option,omitempty"`
	Children []*ApprovalWidget      `json:"children,omitempty"`
}

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

func (r *ApprovalWidgetList) MarshalJSON() ([]byte, error) {
	if r == nil {
		return []byte(`""`), nil
	}
	bs, err := json.Marshal([]*ApprovalWidget(*r))
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%q", bs)), nil
}

type ApprovalWidgetOptions struct {
	IsList  bool
	Option  *ApprovalWidgetOption
	Options []*ApprovalWidgetOption
}

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

func (r *ApprovalWidgetOptions) MarshalJSON() ([]byte, error) {
	if r.IsList {
		return json.Marshal(r.Options)
	}
	return json.Marshal(r.Option)
}

type ApprovalWidgetOption struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Text  string `json:"text,omitempty"`
}

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
