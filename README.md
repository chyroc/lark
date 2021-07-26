# lark

[![codecov](https://codecov.io/gh/chyroc/lark/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark "go report card")](https://goreportcard.com/report/github.com/chyroc/lark)
[![test status](https://github.com/chyroc/lark/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/lark/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark)

[中文版 README](./README_CN.md)

Feishu/Lark Open API Go Sdk, Support ALL Open API and Event Callback.

Supported Features

- Many APIs and events
- Support mock to support test
- Support isv and self-built apps
- Support Logger interface
- Support UserAccessToken
- Use code generation to create, api and document update timely

## Chat

⁣Click [Lark Chat Link](https://applink.feishu.cn/client/chat/chatter/add_by_link?link_token=985n4cf0-70d7-444c-909f-98885892c233) to discuss.

## Install

```shell
go get github.com/chyroc/lark
```

## Docs

https://godoc.org/github.com/chyroc/lark

## Support APIs

API Count: 407, Event Count: 41

<details>
  <summary>
    Click This to See ALL
  </summary>

- Auth
  - ResendAppTicket
  - GetAccessToken
  - RefreshAccessToken
  - GetUserInfo
- Contact
  - CreateUser
  - DeleteUser
  - GetUser
  - BatchGetUser
  - BatchGetUserByID
  - GetUserList
  - UpdateUserPatch
  - UpdateUser
  - CreateDepartment
  - GetDepartment
  - GetDepartmentList
  - GetParentDepartment
  - SearchDepartment
  - UpdateDepartmentPatch
  - UpdateDepartment
  - DeleteDepartment
  - GetEmployeeTypeEnumList
  - UpdateEmployeeTypeEnumPatch
  - DeleteEmployeeTypeEnum
  - CreateEmployeeTypeEnum
  - GetContactCustomAttrList
- Message
  - SendRawMessage
  - SendRawMessageOld
  - ReplyRawMessage
  - DeleteMessage
  - UpdateMessage
  - GetMessageReadUserList
  - GetMessageList
  - GetMessageFile
  - GetMessage
  - DeleteEphemeralMessage
- Chat
  - CreateChat
  - GetChat
  - UpdateChat
  - DeleteChat
  - GetChatListOfSelf
  - SearchChat
  - GetChatMemberList
  - IsInChat
  - AddChatMember
  - DeleteChatMember
  - JoinChat
  - GetChatAnnouncement
  - UpdateChatAnnouncement
- Bot
  - GetBotInfo
  - AddBotToChat
- Calendar
  - CreateCalendarACL
  - DeleteCalendarACL
  - GetCalendarACLList
  - SubscribeCalendarACL
  - CreateCalendar
  - DeleteCalendar
  - GetCalendar
  - GetCalendarList
  - UpdateCalendar
  - SearchCalendar
  - SubscribeCalendar
  - UnsubscribeCalendar
  - SubscribeCalendarChangeEvent
  - CreateCalendarEvent
  - DeleteCalendarEvent
  - GetCalendarEvent
  - GetCalendarEventList
  - UpdateCalendarEvent
  - SearchCalendarEvent
  - SubscribeCalendarEvent
  - CreateCalendarEventAttendee
  - GetCalendarEventAttendeeList
  - DeleteCalendarEventAttendee
  - GetCalendarEventAttendeeChatMemberList
  - GetCalendarFreeBusyList
  - CreateCalendarTimeoffEvent
  - DeleteCalendarTimeoffEvent
  - GenerateCaldavConf
- Drive
  - SearchDriveFile
  - GetDriveFileMeta
  - CreateDriveFile
  - CopyDriveFile
  - DeleteDriveFile
  - DeleteDriveSheetFile
  - CreateDriveFolder
  - GetDriveFolderMeta
  - GetDriveRootFolderMeta
  - GetDriveFolderChildren
  - DownloadDriveFile
  - UploadDriveFile
  - PrepareUploadDriveFile
  - PartUploadDriveFile
  - FinishUploadDriveFile
  - DownloadDriveMedia
  - UploadDriveMedia
  - PrepareUploadDriveMedia
  - PartUploadDriveMedia
  - FinishUploadDriveMedia
  - CreateDriveMemberPermissionOld
  - TransferDriveMemberPermission
  - GetDriveMemberPermissionList
  - CreateDriveMemberPermission
  - DeleteDriveMemberPermission
  - UpdateDriveMemberPermission
  - CheckDriveMemberPermission
  - UpdateDrivePublicPermissionV1Old
  - UpdateDrivePublicPermissionV2Old
  - GetDrivePublicPermissionV2
  - UpdateDrivePublicPermission
  - BatchGetDriveMediaTmpDownloadURL
  - GetDriveCommentList
  - GetDriveComment
  - CreateDriveComment
  - UpdateDriveComment
  - DeleteDriveComment
  - UpdateDriveCommentPatch
  - CreateDriveDoc
  - GetDriveDocContent
  - GetDriveDocRawContent
  - GetDriveDocMeta
  - CreateSheet
  - GetSheetMeta
  - UpdateSheetProperty
  - BatchUpdateSheet
  - ImportSheet
  - MoveSheetDimension
  - PrependSheetValue
  - AppendSheetValue
  - InsertSheetDimensionRange
  - AddSheetDimensionRange
  - UpdateSheetDimensionRange
  - DeleteSheetDimensionRange
  - GetSheetValue
  - BatchGetSheetValue
  - SetSheetValue
  - BatchSetSheetValue
  - SetSheetStyle
  - BatchSetSheetStyle
  - MergeSheetCell
  - UnmergeSheetCell
  - SetSheetValueImage
  - FindSheet
  - ReplaceSheet
  - CreateSheetConditionFormat
  - GetSheetConditionFormat
  - UpdateSheetConditionFormat
  - DeleteSheetConditionFormat
  - CreateSheetProtectedDimension
  - GetSheetProtectedDimension
  - UpdateSheetProtectedDimension
  - DeleteSheetProtectedDimension
  - CreateSheetDataValidationDropdown
  - DeleteSheetDataValidationDropdown
  - UpdateSheetDataValidationDropdown
  - GetSheetDataValidationDropdown
  - CreateSheetFilter
  - DeleteSheetFilter
  - UpdateSheetFilter
  - GetSheetFilter
  - CreateSheetFilterView
  - DeleteSheetFilterView
  - UpdateSheetFilterView
  - GetSheetFilterView
  - QuerySheetFilterView
  - CreateSheetFilterViewCondition
  - DeleteSheetFilterViewCondition
  - UpdateSheetFilterViewCondition
  - GetSheetFilterViewCondition
  - QuerySheetFilterViewCondition
  - CreateSheetFloatImage
  - DeleteSheetFloatImage
  - UpdateSheetFloatImage
  - GetSheetFloatImage
  - QuerySheetFloatImage
- Bitable
  - GetBitableViewList
  - CreateBitableView
  - DeleteBitableView
  - GetBitableRecordList
  - GetBitableRecord
  - CreateBitableRecord
  - BatchCreateBitableRecord
  - UpdateBitableRecord
  - BatchUpdateBitableRecord
  - DeleteBitableRecord
  - BatchDeleteBitableRecord
  - GetBitableFieldList
  - CreateBitableField
  - UpdateBitableField
  - DeleteBitableField
  - GetBitableTableList
  - CreateBitableTable
  - BatchCreateBitableTable
  - DeleteBitableTable
  - BatchDeleteBitableTable
  - GetBitableMeta
- MeetingRoom
  - BatchGetMeetingRoomSummary
  - GetMeetingRoomBuildingList
  - BatchGetMeetingRoomBuilding
  - GetMeetingRoomRoomList
  - BatchGetMeetingRoomRoom
  - BatchGetMeetingRoomFreebusy
  - ReplyMeetingRoomInstance
  - CreateMeetingRoomBuilding
  - UpdateMeetingRoomBuilding
  - DeleteMeetingRoomBuilding
  - BatchGetMeetingRoomBuildingID
  - CreateMeetingRoomRoom
  - UpdateMeetingRoomRoom
  - DeleteMeetingRoomRoom
  - BatchGetMeetingRoomRoomID
  - GetMeetingRoomCountryList
  - GetMeetingRoomDistrictList
- VC
  - ApplyVCReserve
  - UpdateVCReserve
  - DeleteVCReserve
  - GetVCReserve
  - GetVCReserveActiveMeeting
  - GetVCMeeting
  - InviteVCMeeting
  - SetVCHostMeeting
  - EndVCMeeting
  - StartVCMeetingRecording
  - StopVCMeetingRecording
  - GetVCMeetingRecording
  - SetVCPermissionMeetingRecording
  - GetVCDailyReport
  - GetVCTopUserReport
  - GetVCRoomConfig
  - SetVCRoomConfig
- Application
  - IsApplicationUserAdmin
  - GetApplicationUserAdminScope
  - GetApplicationAppVisibility
  - GetApplicationUserVisibleApp
  - GetApplicationAppList
  - UpdateApplicationAppVisibility
  - GetApplicationAppAdminUserList
  - CheckUserIsInApplicationPaidScope
  - GetApplicationOrderList
  - GetApplicationOrder
  - GetApplicationUsageOverview
  - GetApplicationUsageTrend
  - GetApplicationUsageDetail
  - GetApplicationMessageOverview
  - GetApplicationMessageTrend
  - GetApplicationMessageDetail
- Mail
  - CreateMailGroup
  - GetMailGroup
  - GetMailGroupList
  - UpdateMailGroupPatch
  - UpdateMailGroup
  - DeleteMailGroup
  - CreateMailGroupMember
  - GetMailGroupMember
  - GetMailGroupMemberList
  - DeleteMailGroupMember
  - CreateMailGroupPermissionMember
  - GetMailGroupPermissionMember
  - GetMailGroupPermissionMemberList
  - DeleteMailGroupPermissionMember
  - CreatePublicMailbox
  - GetPublicMailbox
  - GetPublicMailboxList
  - UpdatePublicMailboxPatch
  - UpdatePublicMailbox
  - CreatePublicMailboxMember
  - GetPublicMailboxMember
  - GetPublicMailboxMemberList
  - DeletePublicMailboxMember
  - ClearPublicMailboxMember
- Approval
  - GetApproval
  - GetApprovalInstanceList
  - GetApprovalInstance
  - CreateApprovalInstance
  - ApproveApprovalInstance
  - RejectApprovalInstance
  - TransferApprovalInstance
  - CancelApprovalInstance
  - UploadApprovalFile
  - SearchApprovalInstance
  - SearchApprovalTask
  - SearchApprovalCarbonCopy
  - CreateApprovalCarbonCopy
- Helpdesk
  - StartHelpdeskService
  - GetHelpdeskTicket
  - UpdateHelpdeskTicket
  - GetHelpdeskTicketList
  - DownloadHelpdeskTicketImage
  - AnswerHelpdeskTicketUserQuery
  - GetHelpdeskTicketMessageList
  - SendHelpdeskTicketMessage
  - GetHelpdeskTicketCustomizedFieldList
  - DeleteHelpdeskTicketCustomizedField
  - UpdateHelpdeskTicketCustomizedField
  - CreateHelpdeskTicketCustomizedField
  - GetHelpdeskTicketCustomizedField
  - CreateHelpdeskCategory
  - GetHelpdeskCategory
  - UpdateHelpdeskCategory
  - DeleteHelpdeskCategory
  - GetHelpdeskCategoryList
  - CreateHelpdeskFAQ
  - GetHelpdeskFAQ
  - UpdateHelpdeskFAQ
  - DeleteHelpdeskFAQ
  - GetHelpdeskFAQList
  - GetHelpdeskFAQImage
  - SearchHelpdeskFAQ
  - SubscribeHelpdeskEvent
  - UnsubscribeHelpdeskEvent
- Admin
  - GetAdminDeptStats
  - GetAdminUserStats
- HumanAuth
  - GetFaceVerifyAuthResult
  - UploadFaceVerifyImage
  - CropFaceVerifyImage
  - CreateIdentity
- AI
  - RecognizeBasicImage
  - RecognizeSpeechStream
  - RecognizeSpeechFile
  - TranslateText
  - DetectTextLanguage
  - DetectFaceAttributes
- Attendance
  - DownloadAttendanceFile
  - UploadAttendanceFile
  - QueryAttendanceUserSettings
  - UpdateAttendanceUserSettings
  - CreateUpdateAttendanceGroup
  - DeleteAttendanceGroup
  - GetAttendanceGroup
  - CreateAttendanceShift
  - DeleteAttendanceShift
  - GetAttendanceShiftByID
  - GetAttendanceShiftByName
  - GetAttendanceStatisticsData
  - GetAttendanceStatisticsHeader
  - UpdateAttendanceUserStatisticsSettings
  - GetAttendanceUserStatisticsSettings
  - GetAttendanceUserDailyShift
  - GetAttendanceUserTask
  - GetAttendanceUserFlow
  - BatchGetAttendanceUserFlow
  - BatchCreateAttendanceUserFlow
  - GetAttendanceUserTaskRemedy
  - CreateUpdateAttendanceUserDailyShift
  - GetAttendanceUserApproval
  - CreateAttendanceUserApproval
  - GetAttendanceUserAllowedRemedy
  - InitAttendanceRemedyApproval
  - UpdateAttendanceRemedyApproval
- File
  - UploadImage
  - DownloadImage
  - UploadFile
  - DownloadFile
- EventCallback
  - EventV2HelpdeskTicketMessageCreatedV1
  - EventV2HelpdeskTicketCreatedV1
  - EventV2HelpdeskTicketMessageUpdatedV1
  - EventV2ContactDepartmentCreatedV3
  - EventV2ContactDepartmentDeletedV3
  - EventV2ContactDepartmentUpdatedV3
  - EventV2ContactUserUpdatedV3
  - EventV2ContactUserCreatedV3
  - EventV2ContactUserDeletedV3
  - EventV2ContactScopeUpdatedV3
  - EventV2ContactEmployeeTypeEnumCreatedV3
  - EventV2ContactEmployeeTypeEnumActivedV3
  - EventV2ContactEmployeeTypeEnumDeactivatedV3
  - EventV2ContactEmployeeTypeEnumUpdatedV3
  - EventV2ContactEmployeeTypeEnumDeletedV3
  - EventV2IMMessageReceiveV1
  - EventV2IMMessageReadV1
  - EventV2IMChatDisbandedV1
  - EventV2IMChatUpdatedV1
  - EventV2IMChatMemberBotAddedV1
  - EventV2IMChatMemberBotDeletedV1
  - EventV2IMChatMemberUserAddedV1
  - EventV2IMChatMemberUserWithdrawnV1
  - EventV2IMChatMemberUserDeletedV1
  - EventV2VCMeetingMeetingStartedV1
  - EventV2VCMeetingMeetingEndedV1
  - EventV2VCMeetingJoinMeetingV1
  - EventV2VCMeetingLeaveMeetingV1
  - EventV2VCMeetingRecordingStartedV1
  - EventV2VCMeetingRecordingEndedV1
  - EventV2VCMeetingRecordingReadyV1
  - EventV2VCMeetingShareStartedV1
  - EventV2VCMeetingShareEndedV1
  - EventV1AddBot
  - EventV1RemoveBot
  - EventV1P2PChatCreate
  - EventV1ReceiveMessage
  - EventV1AddUserToChat
  - EventV1RemoveUserFromChat
  - EventV1RevokeAddUserFromChat
  - EventV1ChatDisband
- OKR
  - GetOKRPeriodList
  - BatchGetOKR
  - GetUserOKRList
- EHR
  - GetEHREmployeeList
  - DownloadEHRAttachments
- Tenant
  - GetTenant
- Search
  - CreateSearchDataSourceItem
  - GetSearchDataSourceItem
  - DeleteSearchDataSourceItem
  - CreateSearchDataSource
  - GetSearchDataSource
  - UpdateSearchDataSource
  - GetSearchDataSourceList
  - DeleteSearchDataSource
- Hire
  - GetHireJob
  - GetHireJobManager
  - GetHireTalent
  - GetHireAttachment
  - GetHireAttachmentPreview
  - GetHireResumeSource
  - CreateHireNote
  - UpdateHireNote
  - GetHireNote
  - GetHireNoteList
  - GetHireReferralByApplication
  - GetHireJobProcessList
  - CreateHireApplication
  - TerminateHireApplication
  - GetHireApplication
  - GetHireApplicationList
  - GetHireApplicationInterviewList
  - GetHireOfferByApplication
  - GetHireOfferSchema
  - MakeHireTransferOnboardByApplication
  - UpdateHireEmployee
  - GetHireEmployeeByApplication
  - GetHireEmployee
- AppLink
  - OpenLark
  - OpenMiniProgram
  - OpenWebApp
  - OpenChat
  - OpenCalender
  - OpenCalenderView
  - OpenCalenderEventCreate
  - OpenCalenderAccount
  - OpenDocs
  - OpenBot
  - OpenSSOLogin
  - OpenWebURL


</details>

## Usage

### Example: create lark client

- for sample bot and app:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))
```

- for need handle event callback:

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithEventCallbackVerify("<ENCRYPT_KEY>", "<VERIFICATION_TOKEN>"),
)
```

- for helpdesk app:

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithHelpdeskCredential("<HELPDESK_ID>", "HELPDESK_TOKEN"),
)
```

### Example: handle event callback

for more about event callback example, see [./_examples/event_callback.go](./_examples/event_callback.go) .

handle message callback example:

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithEventCallbackVerify("<ENCRYPT_KEY>", "<VERIFICATION_TOKEN>"),
)

// handle message callback
cli.EventCallback.HandlerEventIMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventV2Header, event *lark.EventV2IMMessageReceiveV1) (string, error) {
    _, _, err := cli.Message.Reply(event.Message.MessageID).SendText(ctx, "hi, "+event.Message.Content)
    return "", err
})

http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    cli.EventCallback.ListenCallback(r.Context(), r.Body, w)
})

fmt.Println("start server ...")
log.Fatal(http.ListenAndServe(":9726", nil))
```

### Example: ISV APP

for more about isv example, see [./_examples/isv.go](./_examples/isv.go) .

create isv chat:

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithISV(true),
    lark.WithStore("<NEW_STORE>"),
)

tenantKey1Cli := cli.WithTenant("<TENANT_KEY_1>")
resp, _, err := tenantKey1cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
    Name: ptr.String("<CHAT_NAME_1>"),
})
fmt.Println(resp, err)
```

### Example: get bot info

for more about bot example, see [./_examples/bot.go](./_examples/bot.go) .

get bot info example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Bot.GetBotInfo(ctx, &lark.GetBotInfoReq{})
fmt.Println(resp, err)
```

### Example: send message

for more about send message example, see [./_examples/send_message.go](./_examples/send_message.go) .

send text message example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendText(ctx, "<TEXT>")
fmt.Println(resp, err)
```

### Example: other message

for more about other message example, see [./_examples/other_message.go](./_examples/other_message.go) .

send delete message example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message.DeleteMessage(ctx, &lark.DeleteMessageReq{
    MessageID: "<MESSAGE_ID>",
})
fmt.Println(resp, err)
```

### Example: chat

for more about chat example, see [./_examples/chat.go](./_examples/chat.go) .

create chat example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
    Name: ptr.String("<CHAT_NAME>"),
})
fmt.Println(resp, err)
```

### Example: file

for more about file example, see [./_examples/file.go](./_examples/file.go) .

upload image example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

f, err := os.Open("<FILE_PATH>")
if err != nil {
    panic(err)
}
resp, _, err := cli.File.UploadImage(ctx, &lark.UploadImageReq{
    ImageType: lark.ImageTypeMessage,
    Image:     f,
})
fmt.Println(resp, err)
```

### Example: calendar

for more about calendar example, see [./_examples/calendar.go](./_examples/calendar.go) .

create calendar example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Calendar.CreateCalendar(ctx, &lark.CreateCalendarReq{
    Summary: ptr.String("<SUMMARY>"),
})
fmt.Println(resp, err)
```
