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
	ID       string                  `json:"id,omitempty"`
	Name     string                  `json:"name,omitempty"`
	Type     ApprovalWidgetType      `json:"type,omitempty"`
	Value    interface{}             `json:"value"`
	Option   []*ApprovalWidgetOption `json:"option,omitempty"`
	Children []*ApprovalWidget       `json:"children,omitempty"`
}

func (r *ApprovalWidgetList) UnmarshalJSON(bs []byte) (err error) {
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

func (r *ApprovalWidgetList) MarshalJSON() ([]byte,  error) {
	bs, err := json.Marshal([]*ApprovalWidget(*r))
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%q", bs)), nil
}

type ApprovalWidgetOption struct {
	Value string `json:"value,omitempty"`
	Text  string `json:"text,omitempty"`
}
