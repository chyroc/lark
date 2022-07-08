# lark

[![codecov](https://codecov.io/gh/chyroc/lark/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark "go report card")](https://goreportcard.com/report/github.com/chyroc/lark)
[![test status](https://github.com/chyroc/lark/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/lark/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Flark.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Flark)
[![Used by](https://github-used-by.chyroc.cn/chyroc/lark.svg)]()

[English README](./README.md)

飞书/Lark 的开放接口 Go SDK，支持所有的开放接口，和事件回调。

支持的功能

- 非常多的接口和事件
- 支持 Mock 以支持测试
- 支持 ISV 和自建 App
- 支持 Logger 接口
- 支持 UserAccessToken
- 使用代码生成创建，接口和文档更新及时

## 讨论

⁣点击 [飞书群聊](https://applink.feishu.cn/client/chat/chatter/add_by_link?link_token=985n4cf0-70d7-444c-909f-98885892c233) 一起讨论。

## 安装

```shell
go get github.com/chyroc/lark
```

## 文档

https://godoc.org/github.com/chyroc/lark

## 支持的接口

API 总数: 607, 事件总数: 91

<details>
  <summary>
    点击查看所有接口和事件
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

## 使用

### 例子: 创建 Lark 客户端

- 使用简单的 App

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))
```

- 处理 larksuite（非中国地区）的请求域名

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithOpenBaseURL("https://open.larksuite.com"),
    lark.WithWWWBaseURL("https://www.larksuite.com"),
)
```

- 处理事件回调

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithEventCallbackVerify("<ENCRYPT_KEY>", "<VERIFICATION_TOKEN>"),
)
```

- 服务台

```go
cli := lark.New(
    lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"),
    lark.WithHelpdeskCredential("<HELPDESK_ID>", "HELPDESK_TOKEN"),
)
```

### 例子：处理事件回调

如果需要更多的例子，可以参考：[./_examples/event_callback.go](./_examples/event_callback.go) 。

处理消息的事件回调：

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

### 例子: ISV APP

如果需要更多的例子，可以参考： [./_examples/isv.go](./_examples/isv.go) .

ISV 创建群聊:

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

### 例子: 获取机器人信息

如果需要更多的例子，可以参考： [./_examples/bot.go](./_examples/bot.go) 。

获取机器人信息：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Bot.GetBotInfo(ctx, &lark.GetBotInfoReq{})
fmt.Println(resp, err)
```

### 例子: 发送消息

如果需要更多的例子，可以参考： [./_examples/send_message.go](./_examples/send_message.go) 。

发送文本消息：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message.Send().ToChatID("<CHAT_ID>").SendText(ctx, "<TEXT>")
fmt.Println(resp, err)
```

### 例子: 其他消息

如果需要更多的例子，可以参考： [./_examples/other_message.go](./_examples/other_message.go) 。

撤回消息：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Message.DeleteMessage(ctx, &lark.DeleteMessageReq{
    MessageID: "<MESSAGE_ID>",
})
fmt.Println(resp, err)
```

### 例子: 群聊

如果需要更多的例子，可以参考： [./_examples/chat.go](./_examples/chat.go) 。

创建群聊：

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
    Name: ptr.String("<CHAT_NAME>"),
})
fmt.Println(resp, err)
```

### 例子: 文件

如果需要更多的例子，可以参考： [./_examples/file.go](./_examples/file.go) 。

上传图片

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

### 例子: 日历

如果需要更多的例子，可以参考： [./_examples/calendar.go](./_examples/calendar.go) 。

创建日历

```go
cli := lark.New(lark.WithAppCredential("<APP_ID>", "<APP_SECRET>"))

resp, _, err := cli.Calendar.CreateCalendar(ctx, &lark.CreateCalendarReq{
Summary: ptr.String("<SUMMARY>"),
})
fmt.Println(resp, err)
```
