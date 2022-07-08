# lark

[![codecov](https://codecov.io/gh/chyroc/lark/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark "go report card")](https://goreportcard.com/report/github.com/chyroc/lark)
[![test status](https://github.com/chyroc/lark/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/lark/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Flark.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Flark)
[![Used by](https://github-used-by.chyroc.cn/chyroc/lark.svg)]()

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

API Count: 607, Event Count: 91

<details>
  <summary>
    Click This to See ALL
  </summary>

- ACS
  - GetACSAccessRecordList
  - GetACSAccessRecordPhoto
  - GetACSDeviceList
  - GetACSUser
  - GetACSUserFace
  - GetACSUserList
  - UpdateACSUser
  - UpdateACSUserFace
- AI
  - DetectTextLanguage
  - RecognizeBasicImage
  - RecognizeSpeechFile
  - RecognizeSpeechStream
  - TranslateText
- Admin
  - AdminResetPassword
  - GetAdminDeptStats
  - GetAdminUserStats
- AppLink
  - OpenBot
  - OpenCalender
  - OpenCalenderAccount
  - OpenCalenderEventCreate
  - OpenCalenderView
  - OpenChat
  - OpenDocs
  - OpenLark
  - OpenMiniProgram
  - OpenSSOLogin
  - OpenScan
  - OpenTask
  - OpenTaskCreate
  - OpenTaskDetail
  - OpenTaskTab
  - OpenWebApp
  - OpenWebURL
- Application
  - CheckUserIsInApplicationPaidScope
  - GetApplication
  - GetApplicationAppAdminUserList
  - GetApplicationAppList
  - GetApplicationAppVisibility
  - GetApplicationFeedbackList
  - GetApplicationOrder
  - GetApplicationOrderList
  - GetApplicationUnderAuditList
  - GetApplicationUsageOverview
  - GetApplicationUsageTrend
  - GetApplicationUserAdminScope
  - GetApplicationUserVisibleApp
  - GetApplicationVersion
  - IsApplicationUserAdmin
  - UpdateApplication
  - UpdateApplicationAppVisibility
  - UpdateApplicationFeedback
  - UpdateApplicationVersion
- Approval
  - AddApprovalInstanceSign
  - ApproveApprovalInstance
  - CancelApprovalInstance
  - CheckApprovalExternalInstance
  - CreateApproval
  - CreateApprovalCarbonCopy
  - CreateApprovalExternalApproval
  - CreateApprovalExternalInstance
  - CreateApprovalInstance
  - GetApproval
  - GetApprovalExternalList
  - GetApprovalInstance
  - GetApprovalInstanceList
  - GetApprovalUserTaskList
  - PreviewApprovalInstance
  - RejectApprovalInstance
  - RollbackApprovalInstance
  - SearchApprovalCarbonCopy
  - SearchApprovalInstance
  - SearchApprovalTask
  - SendApprovalMessage
  - SubscribeApprovalSubscription
  - TransferApprovalInstance
  - UnsubscribeApprovalSubscription
  - UpdateApprovalMessage
  - UploadApprovalFile
- Attendance
  - BatchCreateAttendanceUserDailyShift
  - BatchCreateAttendanceUserFlow
  - BatchGetAttendanceUserFlow
  - CreateAttendanceGroup
  - CreateAttendanceShift
  - CreateAttendanceUserApproval
  - CreateAttendanceUserTaskRemedy
  - DeleteAttendanceGroup
  - DeleteAttendanceShift
  - DownloadAttendanceFile
  - GetAttendanceGroup
  - GetAttendanceGroupList
  - GetAttendanceShift
  - GetAttendanceShiftDetail
  - GetAttendanceShiftList
  - GetAttendanceUserApproval
  - GetAttendanceUserDailyShift
  - GetAttendanceUserFlow
  - GetAttendanceUserSettingList
  - GetAttendanceUserStatsData
  - GetAttendanceUserStatsField
  - GetAttendanceUserStatsView
  - GetAttendanceUserTask
  - GetAttendanceUserTaskRemedy
  - GetAttendanceUserTaskRemedyAllowedRemedyList
  - SearchAttendanceGroup
  - UpdateAttendanceRemedyApproval
  - UpdateAttendanceUserSetting
  - UpdateAttendanceUserStatsView
  - UploadAttendanceFile
- Auth
  - GetAccessToken
  - GetUserInfo
  - RefreshAccessToken
  - ResendAppTicket
- Baike
  - CreateBaikeDraft
  - CreateBaikeEntity
  - CreateBaikeUpdate
  - GetBaikeClassificationList
  - GetBaikeEntity
  - GetBaikeEntityList
  - HighlightBaikeEntity
  - MatchBaikeEntity
  - SearchBaikeEntity
  - UpdateBaikeEntity
- Bitable
  - BatchCreateBitableAppRoleMember
  - BatchCreateBitableRecord
  - BatchCreateBitableTable
  - BatchDeleteBitableAppRoleMember
  - BatchDeleteBitableRecord
  - BatchDeleteBitableTable
  - BatchUpdateBitableRecord
  - CreateBitableAppRole
  - CreateBitableAppRoleMember
  - CreateBitableField
  - CreateBitableRecord
  - CreateBitableTable
  - CreateBitableView
  - DeleteBitableAppRole
  - DeleteBitableAppRoleMember
  - DeleteBitableField
  - DeleteBitableRecord
  - DeleteBitableTable
  - DeleteBitableView
  - GetBitableAppRoleList
  - GetBitableAppRoleMemberList
  - GetBitableFieldList
  - GetBitableMeta
  - GetBitableRecord
  - GetBitableRecordList
  - GetBitableTableFormFieldList
  - GetBitableTableList
  - GetBitableViewList
  - UpdateBitableAppRole
  - UpdateBitableField
  - UpdateBitableMeta
  - UpdateBitableRecord
  - UpdateBitableTableFormField
- Bot
  - AddBotToChat
  - GetBotInfo
- Calendar
  - CreateCalendar
  - CreateCalendarACL
  - CreateCalendarEvent
  - CreateCalendarEventAttendee
  - CreateCalendarExchangeBinding
  - CreateCalendarTimeoffEvent
  - DeleteCalendar
  - DeleteCalendarACL
  - DeleteCalendarEvent
  - DeleteCalendarEventAttendee
  - DeleteCalendarExchangeBinding
  - DeleteCalendarTimeoffEvent
  - GenerateCaldavConf
  - GetCalendar
  - GetCalendarACLList
  - GetCalendarEvent
  - GetCalendarEventAttendeeChatMemberList
  - GetCalendarEventAttendeeList
  - GetCalendarEventList
  - GetCalendarExchangeBinding
  - GetCalendarFreeBusyList
  - GetCalendarList
  - GetPrimaryCalendar
  - SearchCalendar
  - SearchCalendarEvent
  - SubscribeCalendar
  - SubscribeCalendarACL
  - SubscribeCalendarChangeEvent
  - SubscribeCalendarEvent
  - UnsubscribeCalendar
  - UpdateCalendar
  - UpdateCalendarEvent
- Chat
  - AddChatMember
  - CreateChat
  - CreateChatManager
  - CreateChatTab
  - DeleteChat
  - DeleteChatManager
  - DeleteChatMember
  - DeleteChatTab
  - DeleteChatTopNotice
  - GetChat
  - GetChatAnnouncement
  - GetChatListOfSelf
  - GetChatMemberList
  - GetChatModeration
  - GetChatOld
  - GetChatTabList
  - IsInChat
  - JoinChat
  - SearchChat
  - SortChatTab
  - UpdateChat
  - UpdateChatAnnouncement
  - UpdateChatModeration
  - UpdateChatTab
  - UpdateChatTopNotice
- Contact
  - AddContactGroupMember
  - BatchAddContactGroupMember
  - BatchDeleteContactGroupMember
  - BatchGetUser
  - BatchGetUserByID
  - BatchGetUserByIDOld
  - BindContactUnitDepartment
  - CreateContactGroup
  - CreateContactUnit
  - CreateDepartment
  - CreateEmployeeTypeEnum
  - CreateUser
  - DeleteContactGroup
  - DeleteContactGroupMember
  - DeleteContactUnit
  - DeleteDepartment
  - DeleteEmployeeTypeEnum
  - DeleteUser
  - GetContactCustomAttrList
  - GetContactGroup
  - GetContactGroupList
  - GetContactGroupMember
  - GetContactMemberGroupList
  - GetContactScopeList
  - GetContactUnit
  - GetContactUnitDepartmentList
  - GetContactUnitList
  - GetDepartment
  - GetDepartmentList
  - GetDepartmentListOld
  - GetEmployeeTypeEnumList
  - GetParentDepartment
  - GetUser
  - GetUserList
  - GetUserListOld
  - SearchDepartment
  - SearchUserOld
  - UnbindContactUnitDepartment
  - UnbindDepartmentChat
  - UpdateContactGroup
  - UpdateContactUnit
  - UpdateDepartment
  - UpdateDepartmentPatch
  - UpdateEmployeeTypeEnumPatch
  - UpdateUser
  - UpdateUserPatch
- Drive
  - AddSheetDimensionRange
  - AddWikiSpaceMember
  - AppendSheetValue
  - BatchDeleteDocxBlock
  - BatchGetDriveMediaTmpDownloadURL
  - BatchGetSheetValue
  - BatchSetSheetStyle
  - BatchSetSheetValue
  - BatchUpdateSheet
  - CheckDriveMemberPermission
  - CopyDriveFile
  - CopyWikiNode
  - CreateDocx
  - CreateDocxBlock
  - CreateDriveComment
  - CreateDriveDoc
  - CreateDriveExportTask
  - CreateDriveFile
  - CreateDriveFileSubscription
  - CreateDriveFolder
  - CreateDriveImportTask
  - CreateDriveMemberPermission
  - CreateDriveMemberPermissionOld
  - CreateSheet
  - CreateSheetConditionFormat
  - CreateSheetDataValidationDropdown
  - CreateSheetFilter
  - CreateSheetFilterView
  - CreateSheetFilterViewCondition
  - CreateSheetFloatImage
  - CreateSheetProtectedDimension
  - CreateWikiNode
  - CreateWikiSpace
  - DeleteDriveComment
  - DeleteDriveFile
  - DeleteDriveMemberPermission
  - DeleteDriveMemberPermissionOld
  - DeleteDriveSheetFile
  - DeleteSheetConditionFormat
  - DeleteSheetDataValidationDropdown
  - DeleteSheetDimensionRange
  - DeleteSheetFilter
  - DeleteSheetFilterView
  - DeleteSheetFilterViewCondition
  - DeleteSheetFloatImage
  - DeleteSheetProtectedDimension
  - DeleteWikiSpaceMember
  - DownloadDriveExportTask
  - DownloadDriveFile
  - DownloadDriveMedia
  - FindSheet
  - FinishUploadDriveFile
  - FinishUploadDriveMedia
  - GetDocxBlock
  - GetDocxBlockListOfBlock
  - GetDocxBlockListOfDocument
  - GetDocxDocument
  - GetDocxDocumentRawContent
  - GetDriveComment
  - GetDriveCommentList
  - GetDriveDocContent
  - GetDriveDocMeta
  - GetDriveDocRawContent
  - GetDriveExportTask
  - GetDriveFileList
  - GetDriveFileMeta
  - GetDriveFileStatistics
  - GetDriveFileSubscription
  - GetDriveFileTask
  - GetDriveFolderChildren
  - GetDriveFolderMeta
  - GetDriveImportTask
  - GetDriveMemberPermissionList
  - GetDrivePublicPermission
  - GetDriveRootFolderMeta
  - GetSheetConditionFormat
  - GetSheetDataValidationDropdown
  - GetSheetFilter
  - GetSheetFilterView
  - GetSheetFilterViewCondition
  - GetSheetFloatImage
  - GetSheetMeta
  - GetSheetProtectedDimension
  - GetSheetValue
  - GetWikiNode
  - GetWikiNodeList
  - GetWikiSpace
  - GetWikiSpaceList
  - GetWikiTask
  - ImportSheet
  - InsertSheetDimensionRange
  - MergeSheetCell
  - MoveDocsToWiki
  - MoveDriveFile
  - MoveSheetDimension
  - MoveWikiNode
  - PartUploadDriveFile
  - PartUploadDriveMedia
  - PrepareUploadDriveFile
  - PrepareUploadDriveMedia
  - PrependSheetValue
  - QuerySheetFilterView
  - QuerySheetFilterViewCondition
  - QuerySheetFloatImage
  - ReplaceSheet
  - SearchDriveFile
  - SetSheetStyle
  - SetSheetValue
  - SetSheetValueImage
  - SubscribeDriveFile
  - TransferDriveMemberPermission
  - UnmergeSheetCell
  - UpdateDocxBlock
  - UpdateDriveComment
  - UpdateDriveCommentPatch
  - UpdateDriveDocContent
  - UpdateDriveFileSubscription
  - UpdateDriveMemberPermission
  - UpdateDriveMemberPermissionOld
  - UpdateDrivePublicPermission
  - UpdateSheetConditionFormat
  - UpdateSheetDataValidationDropdown
  - UpdateSheetDimensionRange
  - UpdateSheetFilter
  - UpdateSheetFilterView
  - UpdateSheetFilterViewCondition
  - UpdateSheetFloatImage
  - UpdateSheetProperty
  - UpdateSheetProtectedDimension
  - UpdateWikiNodeTitle
  - UpdateWikiSpaceSetting
  - UploadDriveFile
  - UploadDriveMedia
- EHR
  - DownloadEHRAttachments
  - GetEHREmployeeList
- Event
  - GetEventOutboundIpList
- EventCallback
  - EventV1AddBot
  - EventV1AddUserToChat
  - EventV1AppOpen
  - EventV1AppStatusChange
  - EventV1AppTicket
  - EventV1AppUninstalled
  - EventV1ApprovalCc
  - EventV1ApprovalInstance
  - EventV1ApprovalTask
  - EventV1ChatDisband
  - EventV1LeaveApprovalV2
  - EventV1OrderPaid
  - EventV1OutApproval
  - EventV1P2PChatCreate
  - EventV1ReceiveMessage
  - EventV1RemedyApproval
  - EventV1RemoveBot
  - EventV1RemoveUserFromChat
  - EventV1RevokeAddUserFromChat
  - EventV1ShiftApproval
  - EventV1ThirdPartyMeetingRoomEventCreated
  - EventV1ThirdPartyMeetingRoomEventDeleted
  - EventV1ThirdPartyMeetingRoomEventUpdated
  - EventV1TripApproval
  - EventV1WorkApproval
  - EventV2ACSAccessRecordCreatedV1
  - EventV2ACSUserUpdatedV1
  - EventV2ApplicationApplicationAppVersionAuditV6
  - EventV2ApplicationApplicationAppVersionPublishApplyV6
  - EventV2ApplicationApplicationAppVersionPublishRevokeV6
  - EventV2ApplicationApplicationCreatedV6
  - EventV2ApplicationApplicationFeedbackCreatedV6
  - EventV2ApplicationApplicationFeedbackUpdatedV6
  - EventV2ApplicationApplicationVisibilityAddedV6
  - EventV2ApprovalApprovalUpdatedV4
  - EventV2CalendarCalendarACLCreatedV4
  - EventV2CalendarCalendarACLDeletedV4
  - EventV2CalendarCalendarChangedV4
  - EventV2CalendarCalendarEventChangedV4
  - EventV2ContactCustomAttrEventUpdatedV3
  - EventV2ContactDepartmentCreatedV3
  - EventV2ContactDepartmentDeletedV3
  - EventV2ContactDepartmentUpdatedV3
  - EventV2ContactEmployeeTypeEnumActivedV3
  - EventV2ContactEmployeeTypeEnumCreatedV3
  - EventV2ContactEmployeeTypeEnumDeactivatedV3
  - EventV2ContactEmployeeTypeEnumDeletedV3
  - EventV2ContactEmployeeTypeEnumUpdatedV3
  - EventV2ContactScopeUpdatedV3
  - EventV2ContactUserCreatedV3
  - EventV2ContactUserDeletedV3
  - EventV2ContactUserUpdatedV3
  - EventV2DriveFileBitableRecordChangedV1
  - EventV2DriveFileDeletedV1
  - EventV2DriveFileEditV1
  - EventV2DriveFilePermissionMemberAddedV1
  - EventV2DriveFilePermissionMemberRemovedV1
  - EventV2DriveFileReadV1
  - EventV2DriveFileTitleUpdatedV1
  - EventV2DriveFileTrashedV1
  - EventV2HelpdeskNotificationApproveV1
  - EventV2HelpdeskTicketCreatedV1
  - EventV2HelpdeskTicketMessageCreatedV1
  - EventV2HelpdeskTicketUpdatedV1
  - EventV2IMChatDisbandedV1
  - EventV2IMChatMemberBotAddedV1
  - EventV2IMChatMemberBotDeletedV1
  - EventV2IMChatMemberUserAddedV1
  - EventV2IMChatMemberUserDeletedV1
  - EventV2IMChatMemberUserWithdrawnV1
  - EventV2IMChatUpdatedV1
  - EventV2IMMessageReactionCreatedV1
  - EventV2IMMessageReactionDeletedV1
  - EventV2IMMessageReadV1
  - EventV2IMMessageReceiveV1
  - EventV2MeetingRoomMeetingRoomCreatedV1
  - EventV2MeetingRoomMeetingRoomDeletedV1
  - EventV2MeetingRoomMeetingRoomStatusChangedV1
  - EventV2MeetingRoomMeetingRoomUpdatedV1
  - EventV2TaskTaskCommentUpdatedV1
  - EventV2TaskTaskUpdateTenantV1
  - EventV2TaskTaskUpdatedV1
  - EventV2VCMeetingJoinMeetingV1
  - EventV2VCMeetingLeaveMeetingV1
  - EventV2VCMeetingMeetingEndedV1
  - EventV2VCMeetingMeetingStartedV1
  - EventV2VCMeetingRecordingEndedV1
  - EventV2VCMeetingRecordingReadyV1
  - EventV2VCMeetingRecordingStartedV1
  - EventV2VCMeetingShareEndedV1
  - EventV2VCMeetingShareStartedV1
- File
  - DownloadFile
  - DownloadImage
  - UploadFile
  - UploadImage
- Helpdesk
  - AnswerHelpdeskTicketUserQuery
  - CancelApproveHelpdeskNotification
  - CancelSendHelpdeskNotification
  - CreateHelpdeskAgentSchedule
  - CreateHelpdeskAgentSkill
  - CreateHelpdeskCategory
  - CreateHelpdeskFAQ
  - CreateHelpdeskNotification
  - CreateHelpdeskTicketCustomizedField
  - DeleteHelpdeskAgentSchedule
  - DeleteHelpdeskAgentSkill
  - DeleteHelpdeskCategory
  - DeleteHelpdeskFAQ
  - DeleteHelpdeskTicketCustomizedField
  - DownloadHelpdeskTicketImage
  - ExecuteSendHelpdeskNotification
  - GetHelpdeskAgentEmail
  - GetHelpdeskAgentSchedule
  - GetHelpdeskAgentScheduleList
  - GetHelpdeskAgentSkill
  - GetHelpdeskAgentSkillList
  - GetHelpdeskAgentSkillRuleList
  - GetHelpdeskCategory
  - GetHelpdeskCategoryList
  - GetHelpdeskFAQ
  - GetHelpdeskFAQImage
  - GetHelpdeskFAQList
  - GetHelpdeskNotification
  - GetHelpdeskTicket
  - GetHelpdeskTicketCustomizedField
  - GetHelpdeskTicketCustomizedFieldList
  - GetHelpdeskTicketCustomizedFields
  - GetHelpdeskTicketList
  - GetHelpdeskTicketMessageList
  - PreviewHelpdeskNotification
  - SearchHelpdeskFAQ
  - SendHelpdeskMessage
  - SendHelpdeskTicketMessage
  - StartHelpdeskService
  - SubmitApproveHelpdeskNotification
  - SubscribeHelpdeskEvent
  - UnsubscribeHelpdeskEvent
  - UpdateHelpdeskAgent
  - UpdateHelpdeskAgentSchedule
  - UpdateHelpdeskAgentSkill
  - UpdateHelpdeskCategory
  - UpdateHelpdeskFAQ
  - UpdateHelpdeskNotification
  - UpdateHelpdeskTicket
  - UpdateHelpdeskTicketCustomizedField
- Hire
  - CreateHireApplication
  - CreateHireNote
  - GetHireApplication
  - GetHireApplicationInterviewList
  - GetHireApplicationList
  - GetHireAttachment
  - GetHireAttachmentPreview
  - GetHireEmployee
  - GetHireEmployeeByApplication
  - GetHireJob
  - GetHireJobManager
  - GetHireJobProcessList
  - GetHireNote
  - GetHireNoteList
  - GetHireOfferByApplication
  - GetHireOfferSchema
  - GetHireReferralByApplication
  - GetHireResumeSource
  - GetHireTalent
  - MakeHireTransferOnboardByApplication
  - TerminateHireApplication
  - UpdateHireEmployee
  - UpdateHireNote
- HumanAuth
  - CreateIdentity
  - CropFaceVerifyImage
  - GetFaceVerifyAuthResult
  - UploadFaceVerifyImage
- Jssdk
  - GetJssdkTicket
- Mail
  - ClearPublicMailboxMember
  - CreateMailGroup
  - CreateMailGroupAlias
  - CreateMailGroupMember
  - CreateMailGroupPermissionMember
  - CreateMailPublicMailboxAlias
  - CreateMailUserMailboxAlias
  - CreatePublicMailbox
  - CreatePublicMailboxMember
  - DeleteMailGroup
  - DeleteMailGroupAlias
  - DeleteMailGroupMember
  - DeleteMailGroupPermissionMember
  - DeleteMailPublicMailboxAlias
  - DeleteMailUserMailbox
  - DeleteMailUserMailboxAlias
  - DeletePublicMailbox
  - DeletePublicMailboxMember
  - GetMailGroup
  - GetMailGroupAliasList
  - GetMailGroupList
  - GetMailGroupMember
  - GetMailGroupMemberList
  - GetMailGroupPermissionMember
  - GetMailGroupPermissionMemberList
  - GetMailPublicMailboxAliasList
  - GetMailUser
  - GetMailUserMailboxAliasList
  - GetPublicMailbox
  - GetPublicMailboxList
  - GetPublicMailboxMember
  - GetPublicMailboxMemberList
  - UpdateMailGroup
  - UpdateMailGroupPatch
  - UpdatePublicMailbox
  - UpdatePublicMailboxPatch
- MeetingRoom
  - BatchGetMeetingRoomBuilding
  - BatchGetMeetingRoomBuildingID
  - BatchGetMeetingRoomFreebusy
  - BatchGetMeetingRoomRoom
  - BatchGetMeetingRoomRoomID
  - BatchGetMeetingRoomSummary
  - CreateMeetingRoomBuilding
  - CreateMeetingRoomRoom
  - DeleteMeetingRoomBuilding
  - DeleteMeetingRoomRoom
  - GetMeetingRoomBuildingList
  - GetMeetingRoomCountryList
  - GetMeetingRoomCustomization
  - GetMeetingRoomDistrictList
  - GetMeetingRoomRoomList
  - ReplyMeetingRoomInstance
  - UpdateMeetingRoomBuilding
  - UpdateMeetingRoomRoom
- Message
  - BatchDeleteMessage
  - BatchSendOldRawMessage
  - CreateMessageReaction
  - DeleteEphemeralMessage
  - DeleteMessage
  - DeleteMessageReaction
  - GetBatchSentMessageProgress
  - GetBatchSentMessageReadUser
  - GetMessage
  - GetMessageFile
  - GetMessageList
  - GetMessageReactionList
  - GetMessageReadUserList
  - ReplyRawMessage
  - SendEphemeralMessage
  - SendRawMessage
  - SendRawMessageOld
  - SendUrgentAppMessage
  - SendUrgentPhoneMessage
  - SendUrgentSmsMessage
  - UpdateMessage
  - UpdateMessageDelay
- OKR
  - BatchGetOKR
  - GetOKRPeriodList
  - GetUserOKRList
- Passport
  - GetPassportSession
- Search
  - CreateSearchDataSource
  - CreateSearchDataSourceItem
  - DeleteSearchDataSource
  - DeleteSearchDataSourceItem
  - GetSearchDataSource
  - GetSearchDataSourceItem
  - GetSearchDataSourceList
  - UpdateSearchDataSource
- Task
  - CompleteTask
  - CreateTask
  - CreateTaskCollaborator
  - CreateTaskComment
  - CreateTaskFollower
  - CreateTaskReminder
  - DeleteTask
  - DeleteTaskCollaborator
  - DeleteTaskComment
  - DeleteTaskFollower
  - DeleteTaskReminder
  - GetTask
  - GetTaskCollaboratorList
  - GetTaskComment
  - GetTaskCommentList
  - GetTaskFollowerList
  - GetTaskList
  - GetTaskReminderList
  - UncompleteTask
  - UpdateTask
  - UpdateTaskComment
- Tenant
  - GetTenant
- VC
  - ApplyVCReserve
  - DeleteVCReserve
  - EndVCMeeting
  - GetVCDailyReport
  - GetVCMeeting
  - GetVCMeetingRecording
  - GetVCReserve
  - GetVCReserveActiveMeeting
  - GetVCRoomConfig
  - GetVCTopUserReport
  - InviteVCMeeting
  - KickoutVCMeeting
  - ListVCMeetingByNo
  - SetVCHostMeeting
  - SetVCPermissionMeetingRecording
  - SetVCRoomConfig
  - StartVCMeetingRecording
  - StopVCMeetingRecording
  - UpdateVCReserve


</details>

## Usage

### Example: create lark client

- for sample bot and app:

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))
```

- processing larksuite (non-China region) request

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithOpenBaseURL("https://open.larksuite.com"),
    lark.WithWWWBaseURL("https://www.larksuite.com"),
)
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
