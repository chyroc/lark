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

API Count: 667, Event Count: 90

<details>
  <summary>
    Click This to See ALL
  </summary>

- ACS
  - GetACSAccessRecordPhoto
  - GetACSAccessRecordList
  - GetACSDeviceList
  - GetACSUserFace
  - UpdateACSUserFace
  - GetACSUser
  - UpdateACSUser
  - GetACSUserList
- AI
  - RecognizeBasicImage
  - RecognizeSpeechStream
  - RecognizeSpeechFile
  - TranslateText
  - DetectTextLanguage
- Admin
  - AdminResetPassword
  - GetAdminDeptStats
  - GetAdminUserStats
  - UploadAdminBadgeImage
  - CreateAdminBadge
  - UpdateAdminBadge
  - GetAdminBadgeList
  - GetAdminBadge
  - CreateAdminBadgeGrant
  - UpdateAdminBadgeGrant
  - GetAdminBadgeGrantList
  - GetAdminBadgeGrant
  - DeleteAdminBadgeGrant
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
  - OpenTask
  - OpenTaskCreate
  - OpenTaskDetail
  - OpenTaskTab
  - OpenScan
- Application
  - GetApplicationRecommendRuleList
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
  - GetApplicationUnderAuditList
  - GetApplication
  - GetApplicationVersion
  - GetApplicationVersionList
  - UpdateApplicationVersion
  - UpdateApplication
  - GetApplicationUsageOverview
  - GetApplicationUsageTrend
  - UpdateApplicationFeedback
  - GetApplicationFeedbackList
- Approval
  - CreateApproval
  - GetApproval
  - GetApprovalList
  - SubscribeApprovalSubscription
  - UnsubscribeApprovalSubscription
  - CreateApprovalInstance
  - GetApprovalInstance
  - GetApprovalInstanceList
  - CancelApprovalInstance
  - CreateApprovalCarbonCopy
  - PreviewApprovalInstance
  - ApproveApprovalInstance
  - RejectApprovalInstance
  - TransferApprovalInstance
  - ResubmitApprovalInstanceTask
  - RollbackApprovalInstance
  - AddApprovalInstanceSign
  - CreateApprovalComment
  - GetApprovalComment
  - DeleteApprovalComment
  - RemoveApprovalComment
  - CreateApprovalExternalApproval
  - CreateApprovalExternalInstance
  - CheckApprovalExternalInstance
  - GetApprovalExternalList
  - UploadApprovalFile
  - UpdateApprovalMessage
  - SendApprovalMessage
  - SearchApprovalInstance
  - SearchApprovalCarbonCopy
  - SearchApprovalTask
  - GetApprovalUserTaskList
- Attendance
  - GetAttendanceGroupList
  - CreateAttendanceGroup
  - SearchAttendanceGroup
  - GetAttendanceGroup
  - DeleteAttendanceGroup
  - GetAttendanceShiftList
  - GetAttendanceShift
  - GetAttendanceShiftDetail
  - DeleteAttendanceShift
  - CreateAttendanceShift
  - GetAttendanceUserDailyShift
  - BatchCreateAttendanceUserDailyShift
  - GetAttendanceUserStatsField
  - GetAttendanceUserStatsView
  - UpdateAttendanceUserStatsView
  - GetAttendanceUserStatsData
  - GetAttendanceUserApproval
  - CreateAttendanceUserApproval
  - UpdateAttendanceRemedyApproval
  - BatchGetAttendanceUserFlow
  - GetAttendanceUserFlow
  - GetAttendanceUserTask
  - BatchCreateAttendanceUserFlow
  - GetAttendanceUserTaskRemedyAllowedRemedyList
  - GetAttendanceUserTaskRemedy
  - CreateAttendanceUserTaskRemedy
  - GetAttendanceUserSettingList
  - UpdateAttendanceUserSetting
  - DownloadAttendanceFile
  - UploadAttendanceFile
- Auth
  - ResendAppTicket
  - GetAccessToken
  - RefreshAccessToken
  - GetUserInfo
- Baike
  - CreateBaikeDraft
  - CreateBaikeUpdate
  - CreateBaikeEntity
  - UpdateBaikeEntity
  - GetBaikeEntity
  - GetBaikeEntityList
  - MatchBaikeEntity
  - SearchBaikeEntity
  - HighlightBaikeEntity
  - ExtractBaikeEntity
  - GetBaikeClassificationList
  - UploadBaikeImage
  - DownloadBaikeImage
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
  - GetBitableAppRoleList
  - CreateBitableAppRole
  - DeleteBitableAppRole
  - UpdateBitableAppRole
  - BatchDeleteBitableAppRoleMember
  - BatchCreateBitableAppRoleMember
  - GetBitableAppRoleMemberList
  - CreateBitableAppRoleMember
  - DeleteBitableAppRoleMember
  - GetBitableTableList
  - CreateBitableTable
  - BatchCreateBitableTable
  - DeleteBitableTable
  - BatchDeleteBitableTable
  - GetBitableDashboardList
  - UpdateBitableTableForm
  - GetBitableTableForm
  - UpdateBitableTableFormField
  - GetBitableTableFormFieldList
  - UpdateBitableMeta
  - GetBitableMeta
- Bot
  - GetBotInfo
- Calendar
  - CreateCalendarACL
  - DeleteCalendarACL
  - GetCalendarACLList
  - SubscribeCalendarACL
  - UnsubscribeCalendarACL
  - GetPrimaryCalendar
  - CreateCalendar
  - DeleteCalendar
  - GetCalendar
  - GetCalendarList
  - UpdateCalendar
  - SearchCalendar
  - SubscribeCalendar
  - UnsubscribeCalendar
  - SubscribeCalendarChangeEvent
  - UnsubscribeCalendarChangeEvent
  - CreateCalendarEvent
  - DeleteCalendarEvent
  - GetCalendarEvent
  - GetCalendarEventList
  - UpdateCalendarEvent
  - SearchCalendarEvent
  - SubscribeCalendarEvent
  - UnsubscribeCalendarEvent
  - CreateCalendarEventAttendee
  - GetCalendarEventAttendeeList
  - DeleteCalendarEventAttendee
  - GetCalendarEventAttendeeChatMemberList
  - GetCalendarFreeBusyList
  - CreateCalendarTimeoffEvent
  - DeleteCalendarTimeoffEvent
  - GenerateCaldavConf
  - CreateCalendarExchangeBinding
  - GetCalendarExchangeBinding
  - DeleteCalendarExchangeBinding
- Chat
  - CreateChat
  - GetChat
  - UpdateChat
  - DeleteChat
  - GetChatListOfSelf
  - SearchChat
  - GetChatMemberList
  - IsInChat
  - CreateChatManager
  - DeleteChatManager
  - AddChatMember
  - DeleteChatMember
  - JoinChat
  - GetChatModeration
  - UpdateChatModeration
  - UpdateChatTopNotice
  - DeleteChatTopNotice
  - GenChatShareLink
  - GetChatAnnouncement
  - UpdateChatAnnouncement
  - CreateChatTab
  - DeleteChatTab
  - GetChatTabList
  - UpdateChatTab
  - SortChatTab
- Contact
  - SearchUserOld
  - CreateUser
  - DeleteUser
  - GetUser
  - GetUserList
  - GetUserListOld
  - BatchGetUser
  - UpdateUserPatch
  - UpdateUser
  - BatchGetUserByID
  - BatchGetUserByIDOld
  - CreateDepartment
  - GetDepartment
  - GetDepartmentList
  - GetDepartmentListOld
  - GetParentDepartment
  - SearchDepartment
  - UpdateDepartmentPatch
  - UpdateDepartment
  - DeleteDepartment
  - UnbindDepartmentChat
  - CreateContactGroup
  - UpdateContactGroup
  - DeleteContactGroup
  - GetContactGroup
  - GetContactGroupList
  - GetContactMemberGroupList
  - AddContactGroupMember
  - BatchAddContactGroupMember
  - DeleteContactGroupMember
  - BatchDeleteContactGroupMember
  - GetContactGroupMember
  - GetEmployeeTypeEnumList
  - UpdateEmployeeTypeEnumPatch
  - DeleteEmployeeTypeEnum
  - CreateEmployeeTypeEnum
  - GetContactCustomAttrList
  - CreateContactUnit
  - UpdateContactUnit
  - DeleteContactUnit
  - GetContactUnit
  - GetContactUnitList
  - BindContactUnitDepartment
  - UnbindContactUnitDepartment
  - GetContactUnitDepartmentList
  - GetContactScopeList
- Drive
  - GetDocxDocument
  - GetDocxDocumentRawContent
  - GetDocxBlockListOfDocument
  - CreateDocx
  - GetDocxBlock
  - CreateDocxBlock
  - UpdateDocxBlock
  - BatchDeleteDocxBlock
  - GetDocxBlockListOfBlock
  - SubscribeDriveFile
  - SearchDriveFile
  - GetDriveFileMeta
  - CreateDriveFile
  - DeleteDriveFile
  - DeleteDriveSheetFile
  - GetDriveFileList
  - GetDriveRootFolderMeta
  - GetDriveFolderMeta
  - GetDriveFolderChildren
  - GetDriveFileStatistics
  - GetDriveFileTask
  - CreateDriveExportTask
  - GetDriveExportTask
  - DownloadDriveExportTask
  - DownloadDriveFile
  - CopyDriveFile
  - CreateDriveFolder
  - MoveDriveFile
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
  - DeleteDriveMemberPermissionOld
  - UpdateDriveMemberPermissionOld
  - UpdateDriveMemberPermission
  - CheckDriveMemberPermission
  - GetDrivePublicPermissionOld
  - GetDrivePublicPermission
  - UpdateDrivePublicPermission
  - BatchGetDriveMediaTmpDownloadURL
  - GetDriveCommentList
  - GetDriveComment
  - CreateDriveComment
  - UpdateDriveComment
  - DeleteDriveComment
  - UpdateDriveCommentPatch
  - CreateDriveFileSubscription
  - GetDriveFileSubscription
  - UpdateDriveFileSubscription
  - CreateDriveDoc
  - GetDriveDocContent
  - UpdateDriveDocContent
  - GetDriveDocRawContent
  - GetDriveDocMeta
  - UpdateSpreadsheet
  - GetSpreadsheet
  - CreateSpreadsheet
  - GetSheetMeta
  - UpdateSheetProperty
  - GetSheet
  - GetSheetList
  - BatchUpdateSheet
  - ImportSheet
  - CreateDriveImportTask
  - GetDriveImportTask
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
  - CreateWikiSpace
  - GetWikiSpaceList
  - GetWikiSpace
  - UpdateWikiSpaceSetting
  - DeleteWikiSpaceMember
  - AddWikiSpaceMember
  - CreateWikiNode
  - GetWikiNodeList
  - MoveWikiNode
  - UpdateWikiNodeTitle
  - CopyWikiNode
  - GetWikiNode
  - MoveDocsToWiki
  - GetWikiTask
- EHR
  - GetEHREmployeeList
  - DownloadEHRAttachments
- Event
  - GetEventOutboundIpList
- EventCallback
  - EventV2ApplicationApplicationAppVersionAuditV6
  - EventV2ApplicationApplicationAppVersionPublishApplyV6
  - EventV2ApplicationApplicationAppVersionPublishRevokeV6
  - EventV2ApplicationApplicationCreatedV6
  - EventV2ContactCustomAttrEventUpdatedV3
  - EventV2DriveFileBitableRecordChangedV1
  - EventV2DriveFileTitleUpdatedV1
  - EventV2DriveFileReadV1
  - EventV2DriveFileEditV1
  - EventV1AppOpen
  - EventV1ShiftApproval
  - EventV1LeaveApprovalV2
  - EventV1OutApproval
  - EventV1WorkApproval
  - EventV2DriveFilePermissionMemberAddedV1
  - EventV2DriveFileTrashedV1
  - EventV2DriveFileDeletedV1
  - EventV2DriveFilePermissionMemberRemovedV1
  - EventV2ApprovalApprovalUpdatedV4
  - EventV1RemedyApproval
  - EventV1ThirdPartyMeetingRoomEventUpdated
  - EventV1ThirdPartyMeetingRoomEventDeleted
  - EventV2MeetingRoomMeetingRoomCreatedV1
  - EventV2MeetingRoomMeetingRoomUpdatedV1
  - EventV2MeetingRoomMeetingRoomStatusChangedV1
  - EventV2MeetingRoomMeetingRoomDeletedV1
  - EventV1ThirdPartyMeetingRoomEventCreated
  - EventV1OrderPaid
  - EventV1AppTicket
  - EventV1AppUninstalled
  - EventV1AppStatusChange
  - EventV2ApplicationApplicationVisibilityAddedV6
  - EventV2ApplicationApplicationFeedbackCreatedV6
  - EventV2ApplicationApplicationFeedbackUpdatedV6
  - EventV2TaskTaskUpdateTenantV1
  - EventV2TaskTaskUpdatedV1
  - EventV2TaskTaskCommentUpdatedV1
  - EventV2HelpdeskTicketMessageCreatedV1
  - EventV2HelpdeskTicketCreatedV1
  - EventV2HelpdeskTicketUpdatedV1
  - EventV2HelpdeskNotificationApproveV1
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
  - EventV2IMMessageReactionDeletedV1
  - EventV2IMMessageReactionCreatedV1
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
  - EventV2ACSAccessRecordCreatedV1
  - EventV2ACSUserUpdatedV1
  - EventV2CalendarCalendarACLCreatedV4
  - EventV2CalendarCalendarACLDeletedV4
  - EventV2CalendarCalendarEventChangedV4
  - EventV2CalendarCalendarChangedV4
  - EventV1P2PChatCreate
  - EventV1ApprovalInstance
  - EventV1ApprovalTask
  - EventV1ApprovalCc
  - EventV2AttendanceUserTaskUpdatedV1
  - EventV2AttendanceUserFlowCreatedV1
  - EventV2IMMessageRecalledV1
  - EventV2VCRoomCreatedV1
  - EventV2VCRoomDeletedV1
  - EventV2VCRoomUpdatedV1
  - EventV2DriveFileBitableFieldChangedV1
- File
  - UploadImage
  - DownloadImage
  - UploadFile
  - DownloadFile
- Helpdesk
  - CreateHelpdeskNotification
  - UpdateHelpdeskNotification
  - GetHelpdeskNotification
  - PreviewHelpdeskNotification
  - SubmitApproveHelpdeskNotification
  - CancelApproveHelpdeskNotification
  - ExecuteSendHelpdeskNotification
  - CancelSendHelpdeskNotification
  - StartHelpdeskService
  - GetHelpdeskTicket
  - UpdateHelpdeskTicket
  - GetHelpdeskTicketList
  - DownloadHelpdeskTicketImage
  - AnswerHelpdeskTicketUserQuery
  - GetHelpdeskTicketCustomizedFields
  - GetHelpdeskTicketMessageList
  - SendHelpdeskTicketMessage
  - SendHelpdeskMessage
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
  - UpdateHelpdeskAgent
  - GetHelpdeskAgentEmail
  - CreateHelpdeskAgentSchedule
  - DeleteHelpdeskAgentSchedule
  - UpdateHelpdeskAgentSchedule
  - GetHelpdeskAgentSchedule
  - GetHelpdeskAgentScheduleList
  - CreateHelpdeskAgentSkill
  - GetHelpdeskAgentSkill
  - UpdateHelpdeskAgentSkill
  - DeleteHelpdeskAgentSkill
  - GetHelpdeskAgentSkillList
  - GetHelpdeskAgentSkillRuleList
  - SubscribeHelpdeskEvent
  - UnsubscribeHelpdeskEvent
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
- HumanAuth
  - GetFaceVerifyAuthResult
  - UploadFaceVerifyImage
  - CropFaceVerifyImage
  - CreateIdentity
- Jssdk
  - GetJssdkTicket
- Mail
  - GetMailUser
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
  - CreateMailGroupAlias
  - GetMailGroupAliasList
  - DeleteMailGroupAlias
  - CreatePublicMailbox
  - GetPublicMailbox
  - GetPublicMailboxList
  - UpdatePublicMailboxPatch
  - UpdatePublicMailbox
  - DeletePublicMailbox
  - CreatePublicMailboxMember
  - GetPublicMailboxMember
  - GetPublicMailboxMemberList
  - DeletePublicMailboxMember
  - ClearPublicMailboxMember
  - CreateMailPublicMailboxAlias
  - GetMailPublicMailboxAliasList
  - DeleteMailPublicMailboxAlias
  - CreateMailUserMailboxAlias
  - DeleteMailUserMailboxAlias
  - GetMailUserMailboxAliasList
  - DeleteMailUserMailbox
- Message
  - SendEphemeralMessage
  - SendUrgentAppMessage
  - SendUrgentSmsMessage
  - SendUrgentPhoneMessage
  - SendRawMessage
  - SendRawMessageOld
  - BatchSendOldRawMessage
  - ReplyRawMessage
  - DeleteMessage
  - BatchDeleteMessage
  - UpdateMessage
  - UpdateMessageDelay
  - GetMessageReadUserList
  - GetBatchSentMessageReadUser
  - GetBatchSentMessageProgress
  - GetMessageList
  - GetMessageFile
  - GetMessage
  - DeleteEphemeralMessage
  - CreateMessageReaction
  - GetMessageReactionList
  - DeleteMessageReaction
  - CreateMessagePin
  - DeleteMessagePin
  - GetMessagePinList
- Mina
  - MinaCodeToSession
- OKR
  - GetOKRPeriodList
  - BatchGetOKR
  - GetUserOKRList
  - DeleteOKRProgressRecord
  - UpdateOKRProgressRecord
  - GetOKRProgressRecord
  - CreateOKRProgressRecord
  - UploadOKRImage
  - GetOKRMetricSourceList
  - GetOKRMetricSourceTableList
  - GetOKRMetricSourceTableList
  - BatchUpdateOKRMetricSourceTableItem
  - UpdateOKRMetricSourceTableItem
  - GetOKRMetricSourceTableItem
  - GetOKRMetricSourceTableItemList
- Passport
  - GetPassportSession
- Search
  - CreateSearchDataSource
  - GetSearchDataSource
  - UpdateSearchDataSource
  - GetSearchDataSourceList
  - DeleteSearchDataSource
  - BatchCreateSearchDataSourceItem
  - CreateSearchDataSourceItem
  - GetSearchDataSourceItem
  - DeleteSearchDataSourceItem
  - UpdateSearchSchema
  - DeleteSearchSchema
  - GetSearchSchema
  - CreateSearchSchema
- Task
  - CreateTaskFollower
  - DeleteTaskFollower
  - BatchDeleteTaskFollower
  - GetTaskFollowerList
  - CreateTaskCollaborator
  - DeleteTaskCollaborator
  - BatchDeleteTaskCollaborator
  - GetTaskCollaboratorList
  - CreateTaskReminder
  - GetTaskReminderList
  - DeleteTaskReminder
  - CreateTask
  - GetTask
  - GetTaskList
  - DeleteTask
  - UpdateTask
  - CompleteTask
  - UncompleteTask
  - CreateTaskComment
  - GetTaskComment
  - GetTaskCommentList
  - DeleteTaskComment
  - UpdateTaskComment
- Tenant
  - GetTenant
- VC
  - ApplyVCReserve
  - UpdateVCReserve
  - DeleteVCReserve
  - GetVCReserve
  - GetVCReserveActiveMeeting
  - GetVCMeeting
  - ListVCMeetingByNo
  - InviteVCMeeting
  - KickoutVCMeeting
  - SetVCHostMeeting
  - EndVCMeeting
  - StartVCMeetingRecording
  - StopVCMeetingRecording
  - GetVCMeetingRecording
  - SetVCPermissionMeetingRecording
  - GetVCDailyReport
  - GetVCTopUserReport
  - GetVCRoomList
  - GetVCRoom
  - BatchGetVCRoom
  - CreateVCRoom
  - UpdateVCRoom
  - DeleteVCRoom
  - SearchVCRoom
  - GetVCRoomLevelList
  - GetVCRoomLevel
  - BatchGetVCRoomLevel
  - CreateVCRoomLevel
  - UpdateVCRoomLevel
  - DeleteVCRoomLevel
  - SearchVCRoomLevel
  - SetVCScopeConfig
  - GetVCScopeConfig
  - GetVCReserveConfig
  - UpdateVCReserveConfig
  - ExportVCMeetingList
  - ExportVCParticipantList
  - ExportVCParticipantQualityList
  - ExportVCResourceReservationList
  - GetVCExportTask
  - DownloadVCExportFile
  - GetVCAlertList
- Verification
  - GetVerification


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
