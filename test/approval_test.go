package test

import (
	"context"
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_GetApproval(t *testing.T) {
	as := assert.New(t)

	ctx := context.TODO()
	cli := AppALLPermission.Ins()

	resp, _, err := cli.Approval.GetApproval(ctx, &lark.GetApprovalReq{
		ApprovalCode: ApprovalALLField.Code,
		Locale:       nil,
	})
	as.Nil(err)

	as.Equal("sdk-demo", resp.ApprovalName)
	widgets := []*lark.ApprovalWidget{
		{Type: lark.ApprovalWidgetTypeInput, Name: "单行文本"},
		{Type: lark.ApprovalWidgetTypeTextarea, Name: "多行文本"},
		{Type: lark.ApprovalWidgetTypeText, Name: "说明 1"},
		{Type: lark.ApprovalWidgetTypeNumber, Name: "数字"},
		{Type: lark.ApprovalWidgetTypeAmount, Name: "金额"},
		{Type: lark.ApprovalWidgetTypeFormula, Name: "计算公式"},
		{Type: lark.ApprovalWidgetTypeRadioV2, Name: "单选", Option: []*lark.ApprovalWidgetOption{{Text: "1"}, {Text: "2"}}},
		{Type: lark.ApprovalWidgetTypeCheckboxV2, Name: "多选", Option: []*lark.ApprovalWidgetOption{{Text: "1"}, {Text: "2"}, {Text: "3"}}},
		{Type: lark.ApprovalWidgetTypeDate, Name: "日期"},
		{Type: lark.ApprovalWidgetTypeDateInterval, Name: "DateInterval"},
		{Type: lark.ApprovalWidgetTypeFieldList, Name: "明细", Children: []*lark.ApprovalWidget{{Name: "单行文本"}, {Name: "多行文本"}}},
		{Type: lark.ApprovalWidgetTypeImage, Name: "图片"},
		{Type: lark.ApprovalWidgetTypeAttachmentV2, Name: "附件"},
		{Type: lark.ApprovalWidgetTypeDepartment, Name: "部门"},
		{Type: lark.ApprovalWidgetTypeContact, Name: "联系人"},
		{Type: lark.ApprovalWidgetTypeConnect, Name: "关联审批"},
		{Type: lark.ApprovalWidgetTypeAddress, Name: "地址"},
	}
	as.Equal(len(widgets), len(resp.Form))
	for idx, widget := range resp.Form {
		expectWidget := widgets[idx]
		as.Equal(expectWidget.Name, widget.Name)
		as.Equal(expectWidget.Type, widget.Type)
		as.Equal(len(expectWidget.Children), len(widget.Children))
		as.Equal(len(expectWidget.Option), len(widget.Option))
		for optionIdx, option := range widget.Option {
			expectOption := expectWidget.Option[optionIdx]
			as.Equal(expectOption.Text, option.Text)
		}
		for childrenIdx, children := range widget.Children {
			expectChildren := expectWidget.Children[childrenIdx]
			as.Equal(expectChildren.Name, children.Name)
		}
	}
}

// func Test_CreateApproval(t *testing.T) {
// 	as := assert.New(t)
//
// 	ctx := context.TODO()
// 	cli := AppALLPermission.Ins()
//
// 	var widgetDefine lark.ApprovalWidgetList
// 	{
// 		resp, _, err := cli.Approval.GetApproval(ctx, &lark.GetApprovalReq{
// 			ApprovalCode: ApprovalALLField.Code,
// 			Locale:       nil,
// 		})
// 		as.Nil(err)
// 		widgetDefine = resp.Form
// 	}
//
// 	v := lark.ApprovalWidgetList{
// 		{
// 			ID:    widgetDefine[0].ID,
// 			Type:  lark.ApprovalWidgetTypeInput,
// 			Value: "test",
// 		},
// 	}
//
// 	resp, _, err := cli.Approval.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{
// 		ApprovalCode:           ApprovalALLField.Code,
// 		UserID:                 &UserAdmin.UserID,
// 		OpenID:                 "",
// 		DepartmentID:           nil,
// 		Form:                   v,
// 		NodeApproverUserIDList: nil,
// 		NodeApproverOpenIDList: nil,
// 		UUID:                   nil,
// 	})
// 	as.Nil(err)
// 	fmt.Println(resp.InstanceCode)
// }
