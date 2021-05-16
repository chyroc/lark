# lark

[![codecov](https://codecov.io/gh/chyroc/lark/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark "go report card")](https://goreportcard.com/report/github.com/chyroc/lark)
[![test status](https://github.com/chyroc/lark/actions/workflows/go.yml/badge.svg)](https://github.com/chyroc/lark/actions)
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

## Install

```shell
go get github.com/chyroc/lark
```

## Docs

https://godoc.org/github.com/chyroc/lark

## Support APIs

API Count: 213, Event Count: 28

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
- Message
  - SendRawMessage
  - ReplyRawMessage
  - DeleteMessage
  - UpdateMessage
  - GetMessageReadUserList
  - GetMessageList
  - GetMessageFile
  - GetMessage
- Chat
  - CreateChat
  - GetChat
  - UpdateChat
  - DeleteChat
  - GetChatListOfSelf
  - SearchChat
  - GetMemberList
  - IsInChat
  - AddMember
  - DeleteMember
  - GetAnnouncement
  - UpdateAnnouncement
- Bot
  - GetBotInfo
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
  - GetSheetMeta
  - UpdateSheetProperty
  - ImportSheet
  - PrependSheetValue
  - AppendSheetValue
  - CreateMemberPermission
  - TransferMemberPermission
  - UpdatePublicPermission
  - GetMemberPermissionList
  - DeleteMemberPermission
  - UpdateMemberPermission
  - CheckMemberPermission
  - UpdatePublicPermissionV2
  - GetPublicPermissionV2
- MeetingRoom
  - BatchGetSummary
  - GetBuildingList
  - BatchGetBuilding
  - GetRoomList
  - BatchGetRoom
  - BatchGetFreebusy
  - ReplyInstance
  - CreateBuilding
  - UpdateBuilding
  - DeleteBuilding
  - BatchGetBuildingID
  - CreateRoom
  - UpdateRoom
  - DeleteRoom
  - BatchGetRoomID
  - GetCountryList
  - GetDistrictList
- VC
  - ApplyReserve
  - UpdateReserve
  - DeleteReserve
  - GetReserveActiveMeeting
  - GetMeeting
  - InviteMeeting
  - SetHostMeeting
  - EndMeeting
  - StartMeetingRecording
  - StopMeetingRecording
  - GetMeetingRecording
  - SetPermissionMeetingRecording
  - GetDailyReport
  - GetTopUserReport
  - QueryRoomConfig
  - SetRoomConfig
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
  - GetInstanceList
- Helpdesk
  - StartService
  - GetTicket
  - UpdateTicket
  - GetTicketList
  - DownloadTicketImage
  - GetTicketMessageList
  - SendTicketMessage
  - GetTicketCustomizedFieldList
  - DeleteTicketCustomizedField
  - UpdateTicketCustomizedField
  - CreateTicketCustomizedField
  - GetTicketCustomizedField
  - CreateCategory
  - GetCategory
  - UpdateCategory
  - DeleteCategory
  - GetCategoryList
  - CreateFAQ
  - GetFAQ
  - UpdateFAQ
  - DeleteFAQ
  - GetFAQList
  - GetFAQImage
  - SearchFAQ
  - SubscribeEvent
  - UnsubscribeEvent
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
- Attendance
  - UpdateUserSettings
  - UploadAttendanceFile
  - CreateUpdateGroup
  - DeleteGroup
  - GetGroup
  - CreateShift
  - DeleteShift
  - GetShiftByID
  - GetShiftByName
  - GetStatisticsData
  - GetStatisticsHeader
  - UpdateUserStatisticsSettings
  - GetUserStatisticsSettings
  - GetUserDailyShift
  - GetUserTask
  - GetUserFlow
  - BatchGetUserFlow
  - BatchCreateUserFlow
  - GetUserTaskRemedy
  - CreateUpdateUserDailyShift
  - GetUserApproval
  - CreateUserApproval
- File
  - UploadImage
  - DownloadImage
  - UploadFile
  - DownloadFile
- EventCallback
  - EventV2ContactUserUpdatedV3
  - EventV2ContactUserCreatedV3
  - EventV2ContactScopeUpdatedV3
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
  - EventV1AddUserToChat
  - EventV1RemoveUserFromChat
  - EventV1RevokeAddUserFromChat
  - EventV1ChatDisband
- OKR
  - GetPeriodList
  - BatchGetOKR
  - GetUserOKRList
- EHR
  - GetEmployeeList
  - DownloadAttachments
- Tenant
  - QueryTenant


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
    lark.WithEventCallbackVerify("<ENCRYPY_KEY>", "<VERIFICATION_TOKEN>"),
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

for more about event callback example, see [./examples/event_callback.go](./examples/event_callback.go) .

handle message callback example:

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithEventCallbackVerify("<ENCRYPY_KEY>", "<VERIFICATION_TOKEN>"),
)

// handle message callback
cli.EventCallback().HandlerEventIMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventV2Header, event *lark.EventV2IMMessageReceiveV1) (string, error) {
    _, _, err := cli.Message().Reply(event.Message.MessageID).SendText(ctx, "hi, "+event.Message.Content)
    return "", err
})

http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    cli.EventCallback().ListenCallback(r.Context(), r.Body, w)
})

fmt.Println("start server ...")
log.Fatal(http.ListenAndServe(":9726", nil))
```

### Example: ISV APP

for more about isv example, see [./examples/isv.go](./examples/isv.go) .

create isv chat:

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithISV(true),
    lark.WithStore("<NEW_STORE>"),
)

tenantKey1Cli := cli.WithTenant("<TENANT_KEY_1>")
resp, _, err := tenantKey1Cli.Chat().CreateChat(ctx, &lark.CreateChatReq{
    Name: ptr.String("<CHAT_NAME_1>"),
})
fmt.Println(resp, err)
```

### Example: get bot info

for more about bot example, see [./examples/bot.go](./examples/bot.go) .

get bot info example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Bot().GetBotInfo(ctx, &lark.GetBotInfoReq{})
fmt.Println(resp, err)
```

### Example: send message

for more about send message example, see [./examples/send_message.go](./examples/send_message.go) .

send text message example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message().Send().ToChatID("<CHAT_ID>").SendText(ctx, "<TEXT>")
fmt.Println(resp, err)
```

### Example: other message

for more about other message example, see [./examples/other_message.go](./examples/other_message.go) .

send delete message example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message().DeleteMessage(ctx, &lark.DeleteMessageReq{
    MessageID: "<MESSAGE_ID>",
})
fmt.Println(resp, err)
```

### Example: chat

for more about chat example, see [./examples/chat.go](./examples/chat.go) .

create chat example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Chat().CreateChat(ctx, &lark.CreateChatReq{
    Name: ptr.String("<CHAT_NAME>"),
})
fmt.Println(resp, err)
```

### Example: file

for more about file example, see [./examples/file.go](./examples/file.go) .

upload image example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

f, err := os.Open("<FILE_PATH>")
if err != nil {
    panic(err)
}
resp, _, err := cli.File().UploadImage(ctx, &lark.UploadImageReq{
    ImageType: lark.ImageTypeMessage,
    Image:     f,
})
fmt.Println(resp, err)
```

### Example: calendar

for more about calendar example, see [./examples/calendar.go](./examples/calendar.go) .

create calendar example:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Calendar().CreateCalendar(ctx, &lark.CreateCalendarReq{
    Summary: ptr.String("<SUMMARY>"),
})
fmt.Println(resp, err)
```
