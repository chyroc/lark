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

API 总数: 924, 事件总数: 124

<details>
  <summary>
    点击查看所有接口和事件
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
  - DetectFaceAttributes
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
  - OpenScan
  - OpenWorkbench
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
- Application
  - SearchApplicationWorkplaceAccessData
  - SearchApplicationCustomWorkplaceAccessData
  - SearchApplicationWorkplaceBlockAccessData
  - GetApplicationFavourite
  - GetApplicationRecommend
  - GetApplicationRecommendRuleList
  - IsApplicationUserAdmin
  - GetApplicationUserAdminScope
  - GetApplicationAppVisibility
  - GetApplicationContactsRangeConfiguration
  - UpdateApplicationContactsRangeConfiguration
  - GetApplicationUserVisibleApp
  - CheckApplicationVisibleWhiteBlackList
  - GetApplicationAppList
  - UpdateApplicationAppVisibility
  - UpdateApplicationAppVisibilityV6
  - GetApplicationAppAdminUserList
  - CheckUserIsInApplicationPaidScope
  - GetApplicationOrderList
  - GetApplicationOrder
  - GetApplicationUnderAuditList
  - GetApplication
  - GetApplicationVersion
  - GetApplicationVersionList
  - GetApplicationVersionContactsRangeSuggest
  - UpdateApplicationVersion
  - UpdateApplication
  - GetApplicationUsageDepartmentOverview
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
  - GetApprovalExternalApproval
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
  - TransformApprovalUserID
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
  - UpdateAttendanceLeaveAccrualRecord
  - GetAttendanceLeaveEmployExpireRecord
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
  - CopyBitableDashboard
  - GetBitableDashboardList
  - UpdateBitableView
  - GetBitableView
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
  - CreateBitableTable
  - BatchCreateBitableTable
  - DeleteBitableTable
  - BatchDeleteBitableTable
  - UpdateBitableTable
  - GetBitableTableList
  - UpdateBitableTableForm
  - GetBitableTableForm
  - UpdateBitableTableFormField
  - GetBitableTableFormFieldList
  - CopyBitableApp
  - CreateBitableApp
  - GetBitableMeta
  - UpdateBitableMeta
- Bot
  - GetBotInfo
  - AddBotToChat
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
  - GetChatOld
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
  - CreateChatMenuTree
  - DeleteChatMenuTree
  - UpdateChatMenuTree
  - SortChatMenuTree
  - GetChatMenuTree
- Contact
  - CreateUser
  - DeleteUser
  - ResurrectUser
  - GetUser
  - BatchGetUser
  - GetUserList
  - GetUserListOld
  - UpdateUserPatch
  - UpdateUser
  - BatchGetUserByID
  - BatchGetUserByIDOld
  - SearchUserOld
  - UpdateUserID
  - CreateDepartment
  - GetDepartment
  - BatchGetDepartment
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
  - CreateContactFunctionalRole
  - DeleteContactFunctionalRole
  - UpdateContactFunctionalRole
  - BatchCreateContactFunctionalRoleMember
  - BatchDeleteContactFunctionalRoleMember
  - UpdateContactFunctionalRoleMemberScope
  - GetContactFunctionalRoleMemberScope
  - GetContactFunctionalRoleMember
  - CreateContactJobLevel
  - DeleteContactJobLevel
  - UpdateContactJobLevel
  - GetContactJobLevel
  - GetContactJobLevelList
  - CreateContactJobFamily
  - DeleteContactJobFamily
  - UpdateContactJobFamily
  - GetContactJobFamily
  - GetContactJobFamilyList
  - GetContactJobTitle
  - GetContactJobTitleList
  - GetContactWorkCity
  - GetContactWorkCityList
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
- CoreHR
  - GetCoreHRCustomField
  - GetCoreHRCustomFieldList
  - GetCoreHRCustomFieldObjectApiNameList
  - GetCoreHRCountryRegion
  - GetCoreHRCountryRegionList
  - GetCoreHRSubdivision
  - GetCoreHRSubdivisionList
  - GetCoreHRSubregion
  - GetCoreHRSubregionList
  - CreateCoreHREmployeeType
  - DeleteCoreHREmployeeType
  - UpdateCoreHREmployeeType
  - GetCoreHREmployeeType
  - GetCoreHREmployeeTypeList
  - CreateCoreHRNationalIDType
  - DeleteCoreHRNationalIDType
  - UpdateCoreHRNationalIDType
  - GetCoreHRNationalIDType
  - GetCoreHRNationalIDTypeList
  - CreateCoreHRWorkingHoursType
  - DeleteCoreHRWorkingHoursType
  - UpdateCoreHRWorkingHoursType
  - GetCoreHRWorkingHoursType
  - GetCoreHRWorkingHoursTypeList
  - GetCoreHRCurrency
  - GetCoreHRCurrencyList
  - BatchGetCoreHREmployee
  - SearchCoreHREmployee
  - CreateCoreHREmployment
  - DeleteCoreHREmployment
  - UpdateCoreHREmployment
  - CreateCoreHRPerson
  - UpdateCoreHRPerson
  - DeleteCoreHRPerson
  - UploadCoreHRPersonFile
  - DownloadCoreHRPersonFile
  - GetCoreHRPerson
  - CreateCoreHRJobData
  - DeleteCoreHRJobData
  - UpdateCoreHRJobData
  - GetCoreHRJobData
  - BatchGetCoreHRJobData
  - GetCoreHRJobDataList
  - GetCoreHRDepartmentParentList
  - SearchCoreHRDepartment
  - CreateCoreHRDepartment
  - DeleteCoreHRDepartment
  - UpdateCoreHRDepartment
  - GetCoreHRDepartment
  - GetCoreHRDepartmentList
  - BatchGetCoreHRLocation
  - CreateCoreHRLocation
  - DeleteCoreHRLocation
  - GetCoreHRLocation
  - GetCoreHRLocationList
  - BatchGetCoreHRCompany
  - CreateCoreHRCompany
  - DeleteCoreHRCompany
  - GetCoreHRCompany
  - GetCoreHRCompanyList
  - BatchGetCoreHRJobLevel
  - CreateCoreHRJobLevel
  - DeleteCoreHRJobLevel
  - UpdateCoreHRJobLevel
  - GetCoreHRJobLevel
  - GetCoreHRJobLevelList
  - BatchGetCoreHRJobFamily
  - CreateCoreHRJobFamily
  - DeleteCoreHRJobFamily
  - UpdateCoreHRJobFamily
  - GetCoreHRJobFamily
  - GetCoreHRJobFamilyList
  - CreateCoreHRJob
  - DeleteCoreHRJob
  - UpdateCoreHRJob
  - GetCoreHRJobV2
  - GetCoreHRJobListV2
  - GetCoreHRJob
  - GetCoreHRJobList
  - CreateCoreHRPreHire
  - DeleteCoreHRPreHire
  - UpdateCoreHRPreHire
  - GetCoreHRPreHire
  - GetCoreHRPreHireList
  - SearchCoreHRContract
  - CreateCoreHRContract
  - DeleteCoreHRContract
  - UpdateCoreHRContract
  - GetCoreHRContract
  - GetCoreHRContractList
  - GetCoreHRTransferReasonList
  - GetCoreHRTransferTypeList
  - CreateCoreHRJobChange
  - SearchCoreHRJobChange
  - GetCoreHROffboardingList
  - CreateCoreHROffboarding
  - SearchCoreHROffboarding
  - CreateCoreHRLeaveGrantingRecord
  - DeleteCoreHRLeaveGrantingRecord
  - GetCoreHRLeaveTypeList
  - GetCoreHRLeaveBalanceList
  - GetCoreHRLeaveRequestHistoryList
  - SearchCoreHRAssignedUser
  - GetCoreHRSecurityGroupBp
  - GetCoreHRAuthorization
  - GetCoreHRAuthorizationList
  - GetCoreHRSecurityGroupList
  - GetCoreHRProcessFormVariableData
- Drive
  - GetDriveRootFolderMeta
  - GetDriveFileList
  - GetDriveFolderMeta
  - CreateDriveFolder
  - GetDriveFileMeta
  - CreateDriveFile
  - CopyDriveFile
  - MoveDriveFile
  - DeleteDriveFile
  - GetDriveFileStatistics
  - CreateDriveFileShortcut
  - GetDriveFileTask
  - UploadDriveMedia
  - DownloadDriveMedia
  - PrepareUploadDriveMedia
  - PartUploadDriveMedia
  - FinishUploadDriveMedia
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
  - UnsubscribeDriveFile
  - GetSubscribeDriveFile
  - SearchDriveFile
  - DeleteDriveSheetFile
  - GetDriveFolderChildren
  - CreateDriveExportTask
  - GetDriveExportTask
  - DownloadDriveExportTask
  - GetDriveFileViewRecordList
  - CreateDriveFileVersion
  - DeleteDriveFileVersion
  - GetDriveFileVersion
  - GetDriveFileVersionList
  - DownloadDriveFile
  - UploadDriveFile
  - PrepareUploadDriveFile
  - PartUploadDriveFile
  - FinishUploadDriveFile
  - CreateDriveMemberPermissionOld
  - TransferDriveOwnerPermission
  - CheckDriveMemberPermission
  - GetDriveMemberPermissionList
  - GetDriveMemberPermissionListOld
  - CreateDriveMemberPermission
  - UpdateDriveMemberPermission
  - DeleteDriveMemberPermission
  - GetDrivePublicPermissionOld
  - CreateDrivePermissionPublicPassword
  - UpdateDrivePermissionPublicPassword
  - DeleteDrivePermissionPublicPassword
  - GetDrivePublicPermission
  - UpdateDrivePublicPermission
  - GetDrivePublicPermissionV2
  - UpdateDrivePublicPermissionV2
  - BatchGetDriveMediaTmpDownloadURL
  - GetDriveCommentList
  - CreateDriveComment
  - GetDriveComment
  - BatchGetDriveComment
  - GetDriveCommentReplyList
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
  - DeleteDriveMemberPermissionOld
  - UpdateDriveMemberPermissionOld
  - TransferDriveMemberPermission
  - CheckDriveMemberPermissionOld
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
  - EventV1LeaveApprovalRevert
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
  - EventV1AddBot
  - EventV1RemoveBot
  - EventV1P2PChatCreate
  - EventV1ReceiveMessage
  - EventV1AddUserToChat
  - EventV1RemoveUserFromChat
  - EventV1RevokeAddUserFromChat
  - EventV1ChatDisband
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
  - EventV2VCMeetingAllMeetingStartedV1
  - EventV2VCMeetingAllMeetingEndedV1
  - EventV2VCRoomLevelCreatedV1
  - EventV2VCRoomLevelDeletedV1
  - EventV2VCRoomLevelUpdatedV1
  - EventV2VCReserveConfigUpdatedV1
  - EventV2ApplicationBotMenuV6
  - EventV2HireOfferStatusChangedV1
  - EventV2CorehrEmploymentCreatedV1
  - EventV2CorehrEmploymentUpdatedV1
  - EventV2CorehrEmploymentDeletedV1
  - EventV2CorehrPersonCreatedV1
  - EventV2CorehrPersonUpdatedV1
  - EventV2CorehrPersonDeletedV1
  - EventV2CorehrDepartmentCreatedV1
  - EventV2CorehrDepartmentUpdatedV1
  - EventV2CorehrDepartmentDeletedV1
  - EventV2CorehrJobDataEmployedV1
  - EventV2CorehrContractCreatedV1
  - EventV2CorehrEmploymentConvertedV1
  - EventV2CorehrJobChangeUpdatedV1
  - EventV2CorehrJobDataChangedV1
  - EventV2CorehrOffboardingUpdatedV1
  - EventV2CorehrEmploymentResignedV1
  - EventV2CorehrOrgRoleAuthorizationUpdatedV1
  - EventV2HireEHRImportTaskForInternshipOfferImportedV1
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
  - GetHireJobConfig
  - CreateHireJob
  - UpdateHireJob
  - UpdateHireJobConfig
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
  - GetHireReferralWebsiteJobPostList
  - GetHireReferralWebsiteJobPost
  - GetHireReferralByApplication
  - CreateHireExternalApplication
  - UpdateHireExternalApplication
  - DeleteHireExternalApplication
  - CreateHireExternalInterview
  - CreateHireExternalInterviewAssessment
  - CreateHireExternalBackgroundCheck
  - AddHireTalentToFolder
  - GetHireTalentFolderList
  - BatchGetHireTalent
  - GetHireJobProcessList
  - CreateHireApplication
  - TerminateHireApplication
  - GetHireApplication
  - GetHireApplicationList
  - GetHireEvaluationList
  - GetHireQuestionnaireList
  - GetHireInterviewList
  - UpdateHireOfferStatus
  - CreateHireOffer
  - UpdateHireOffer
  - UpdateHireOfferInternStatus
  - GetHireOfferByApplication
  - GetHireOffer
  - GetHireOfferList
  - GetHireApplicationInterviewList
  - GetHireOfferSchema
  - MakeHireTransferOnboardByApplication
  - UpdateHireEmployee
  - GetHireEmployeeByApplication
  - GetHireEmployee
  - UpdateHireEHRImportTask
- HumanAuth
  - GetFaceVerifyAuthResult
  - UploadFaceVerifyImage
  - CropFaceVerifyImage
  - CreateIdentity
- Jssdk
  - GetJssdkTicket
- Lingo
  - CreateLingoDraft
  - UpdateLingoDraft
  - CreateLingoEntity
  - UpdateLingoEntity
  - GetLingoEntity
  - GetLingoEntityList
  - MatchLingoEntity
  - SearchLingoEntity
  - HighlightLingoEntity
  - ExtractLingoEntity
  - GetLingoClassificationList
  - GetLingoRepoList
  - UploadLingoFile
  - DownloadLingoFile
- Mail
  - GetMailUser
  - CreateMailGroup
  - DeleteMailGroup
  - UpdateMailGroupPatch
  - UpdateMailGroup
  - GetMailGroup
  - GetMailGroupList
  - BatchCreateMailGroupManager
  - BatchDeleteMailGroupManager
  - GetMailGroupManagerList
  - CreateMailGroupMember
  - DeleteMailGroupMember
  - GetMailGroupMember
  - GetMailGroupMemberList
  - BatchCreateMailGroupMember
  - BatchDeleteMailGroupMember
  - CreateMailGroupPermissionMember
  - DeleteMailGroupPermissionMember
  - GetMailGroupPermissionMember
  - GetMailGroupPermissionMemberList
  - BatchCreateMailGroupPermissionMember
  - BatchDeleteMailGroupPermissionMember
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
  - BatchCreatePublicMailboxMember
  - BatchDeletePublicMailboxMember
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
  - UpdateMessageEdit
  - DeleteMessage
  - ForwardMessage
  - MergeForwardMessage
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
  - GetMessageSpecialFocusList
  - GetMessageSpecialFocusUnread
- Mina
  - MinaCodeToSession
- Minutes
  - GetMinutesStatistics
  - GetMinutesMinute
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
  - BatchUpdateOKRMetricSourceTableItem
  - UpdateOKRMetricSourceTableItem
  - GetOKRMetricSourceTableItem
  - GetOKRMetricSourceTableItemList
- Passport
  - GetPassportSession
- Performance
  - GetPerformanceSemesterList
  - GetPerformanceStageTaskByUser
  - GetPerformanceStageTaskByPage
- PersonalSettings
  - CreatePersonalSettingsSystemStatus
  - DeletePersonalSettingsSystemStatus
  - UpdatePersonalSettingsSystemStatus
  - GetPersonalSettingsSystemStatusList
  - BatchOpenPersonalSettingsSystemStatus
  - BatchClosePersonalSettingsSystemStatus
- Search
  - SearchMessage
  - SearchApp
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
  - GetTenantProductAssignInfo
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
  - GetVCReserveConfigForm
  - UpdateVCReserveConfigForm
  - GetVCReserveConfigAdmin
  - UpdateVCReserveConfigAdmin
  - GetVCReserveConfigDisableInform
  - UpdateVCReserveConfigDisableInform
  - ExportVCMeetingList
  - ExportVCParticipantList
  - ExportVCParticipantQualityList
  - ExportVCResourceReservationList
  - GetVCExportTask
  - DownloadVCExportFile
  - GetVCAlertList
  - GetVCMeetingList
  - GetVCParticipantList
  - GetVCParticipantQualityList
  - GetVCResourceReservationList
- Verification
  - GetVerification


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
    Name: ptrString("<CHAT_NAME_1>"),
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
    Name: ptrString("<CHAT_NAME>"),
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
Summary: ptrString("<SUMMARY>"),
})
fmt.Println(resp, err)
```
