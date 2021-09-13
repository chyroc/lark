package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_GetApproval(t *testing.T) {
	as := assert.New(t)

	cli := AppAllPermission.Ins()

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
		{Type: lark.ApprovalWidgetTypeRadioV2, Name: "单选", Option: &lark.ApprovalWidgetOptions{IsList: true, Options: []*lark.ApprovalWidgetOption{{Text: "1"}, {Text: "2"}}}},
		{Type: lark.ApprovalWidgetTypeCheckboxV2, Name: "多选", Option: &lark.ApprovalWidgetOptions{IsList: true, Options: []*lark.ApprovalWidgetOption{{Text: "1"}, {Text: "2"}, {Text: "3"}}}},
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
		if expectWidget.Option != nil {
			as.Equal(len(expectWidget.Option.Options), len(widget.Option.Options))
			for optionIdx, option := range widget.Option.Options {
				expectOption := expectWidget.Option.Options[optionIdx]
				as.Equal(expectOption.Text, option.Text)
			}
			for childrenIdx, children := range widget.Children {
				expectChildren := expectWidget.Children[childrenIdx]
				as.Equal(expectChildren.Name, children.Name)
			}
		}
	}
}

func Test_UnmarshalGetApprovalInstance(t *testing.T) {
	as := assert.New(t)

	s := `{
    "approval_code": "96F7DAD7-5F10-416B-A50C-C3331A2B3010",
    "approval_name": "sdk-demo",
    "comment_list": [],
    "department_id": "",
    "end_time": 0,
    "form": "[{\"id\":\"widget3\",\"name\":\"单行文本\",\"type\":\"input\",\"ext\":null,\"value\":\"test\"},{\"id\":\"widget16215623422760001\",\"name\":\"说明 1\",\"type\":\"text\",\"ext\":null,\"value\":\"说明\"}]",
    "open_id": "ou_6e2bbe38e6551b5b2731a409ef50e8ce",
    "serial_number": "202105230010",
    "start_time": 1621742291784,
    "status": "PENDING",
    "task_list": [
      {
        "end_time": 0,
        "id": "6965330108705980419",
        "node_id": "6db614baa6d5cb208decf9efa2e3eee3",
        "node_name": "直接主管",
        "open_id": "ou_6e2bbe38e6551b5b2731a409ef50e8ce",
        "start_time": 1621742292164,
        "status": "PENDING",
        "type": "AND",
        "user_id": "3gg4cf3g"
      }
    ],
    "timeline": [
      {
        "create_time": 1621742291784,
        "ext": "{\"user_id\":\"1\"}",
        "open_id": "ou_6e2bbe38e6551b5b2731a409ef50e8ce",
        "type": "START",
        "user_id": "3gg4cf3g"
      }
    ],
    "user_id": "3gg4cf3g",
    "uuid": ""
  }
`
	res := new(lark.GetApprovalInstanceResp)
	err := json.Unmarshal([]byte(s), res)
	as.Nil(err)
	as.NotNil(res)
	as.Equal("1", ptr.ValueString(res.Timeline[0].Ext.UserID))
	printData(res)
}

func Test_Create_CancelApproval(t *testing.T) {
	as := assert.New(t)

	cli := AppAllPermission.Ins()

	t.Run("cancel", func(t *testing.T) {
		instanceCode, _ := testCreateApproval(t, cli, ApprovalALLField.Code, UserAdmin.UserID)
		_, _, err := cli.Approval.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{
			ApprovalCode: ApprovalALLField.Code,
			InstanceCode: instanceCode,
			UserID:       UserAdmin.UserID,
		})
		as.Nil(err)
	})

	t.Run("approve-reject", func(t *testing.T) {
		t.SkipNow()
		taskDone := map[string]bool{}
		instanceCode, instance := testCreateApproval(t, cli, ApprovalALLField.Code, UserAdmin.UserID)
		for taskIdx, task := range instance.TaskList {
			if taskDone[task.ID] {
				continue
			}
			taskDone[task.ID] = true
			fmt.Println("task", task.ID, task.NodeID)
			_, _, err := cli.Approval.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{
				ApprovalCode: ApprovalALLField.Code,
				InstanceCode: instanceCode,
				UserID:       UserAdmin.UserID,
				TaskID:       task.ID,
				Comment:      ptr.String(fmt.Sprintf("task: %d, approve", taskIdx)),
			})
			as.Nil(err)
		}

		resp, _, err := cli.Approval.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{
			InstanceCode: instanceCode,
			Locale:       nil,
		})
		as.Nil(err)
		for taskIdx, task := range resp.TaskList {
			if taskDone[task.ID] {
				continue
			}
			taskDone[task.ID] = true
			fmt.Println("task", task.ID, task.NodeID)
			_, _, err := cli.Approval.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{
				ApprovalCode: ApprovalALLField.Code,
				InstanceCode: instanceCode,
				UserID:       UserAdmin.UserID,
				TaskID:       task.ID,
				Comment:      ptr.String(fmt.Sprintf("task: %d, approve", taskIdx)),
			})
			as.Nil(err)
		}
	})
}

func testCreateApproval(t *testing.T, cli *lark.Lark, approvalCode, userID string) (string, *lark.GetApprovalInstanceResp) {
	as := assert.New(t)

	var widgetDefine lark.ApprovalWidgetList
	{
		resp, _, err := cli.Approval.GetApproval(ctx, &lark.GetApprovalReq{
			ApprovalCode: approvalCode,
			Locale:       nil,
		})
		as.Nil(err)
		widgetDefine = resp.Form
	}

	v := lark.ApprovalWidgetList{
		{
			ID:    widgetDefine[0].ID,
			Type:  lark.ApprovalWidgetTypeInput,
			Value: "test",
		},
	}

	instanceCode := ""
	{
		resp, _, err := cli.Approval.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{
			ApprovalCode:           approvalCode,
			UserID:                 &userID,
			OpenID:                 "",
			DepartmentID:           nil,
			Form:                   v,
			NodeApproverUserIDList: nil,
			NodeApproverOpenIDList: nil,
			UUID:                   nil,
		})
		as.Nil(err)
		instanceCode = resp.InstanceCode
	}

	resp, _, err := cli.Approval.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{
		InstanceCode: instanceCode,
		Locale:       nil,
	})
	as.Nil(err)

	return instanceCode, resp
}
