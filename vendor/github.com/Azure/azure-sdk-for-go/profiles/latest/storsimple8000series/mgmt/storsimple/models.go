// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storsimple

import original "github.com/Azure/azure-sdk-for-go/services/storsimple8000series/mgmt/2017-06-01/storsimple"

type AccessControlRecordsClient = original.AccessControlRecordsClient

func NewAccessControlRecordsClient(subscriptionID string) AccessControlRecordsClient {
	return original.NewAccessControlRecordsClient(subscriptionID)
}
func NewAccessControlRecordsClientWithBaseURI(baseURI string, subscriptionID string) AccessControlRecordsClient {
	return original.NewAccessControlRecordsClientWithBaseURI(baseURI, subscriptionID)
}

type AlertsClient = original.AlertsClient

func NewAlertsClient(subscriptionID string) AlertsClient {
	return original.NewAlertsClient(subscriptionID)
}
func NewAlertsClientWithBaseURI(baseURI string, subscriptionID string) AlertsClient {
	return original.NewAlertsClientWithBaseURI(baseURI, subscriptionID)
}

type BackupPoliciesClient = original.BackupPoliciesClient

func NewBackupPoliciesClient(subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClient(subscriptionID)
}
func NewBackupPoliciesClientWithBaseURI(baseURI string, subscriptionID string) BackupPoliciesClient {
	return original.NewBackupPoliciesClientWithBaseURI(baseURI, subscriptionID)
}

type BackupsClient = original.BackupsClient

func NewBackupsClient(subscriptionID string) BackupsClient {
	return original.NewBackupsClient(subscriptionID)
}
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return original.NewBackupsClientWithBaseURI(baseURI, subscriptionID)
}

type BackupSchedulesClient = original.BackupSchedulesClient

func NewBackupSchedulesClient(subscriptionID string) BackupSchedulesClient {
	return original.NewBackupSchedulesClient(subscriptionID)
}
func NewBackupSchedulesClientWithBaseURI(baseURI string, subscriptionID string) BackupSchedulesClient {
	return original.NewBackupSchedulesClientWithBaseURI(baseURI, subscriptionID)
}

type BandwidthSettingsClient = original.BandwidthSettingsClient

func NewBandwidthSettingsClient(subscriptionID string) BandwidthSettingsClient {
	return original.NewBandwidthSettingsClient(subscriptionID)
}
func NewBandwidthSettingsClientWithBaseURI(baseURI string, subscriptionID string) BandwidthSettingsClient {
	return original.NewBandwidthSettingsClientWithBaseURI(baseURI, subscriptionID)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}

type CloudAppliancesClient = original.CloudAppliancesClient

func NewCloudAppliancesClient(subscriptionID string) CloudAppliancesClient {
	return original.NewCloudAppliancesClient(subscriptionID)
}
func NewCloudAppliancesClientWithBaseURI(baseURI string, subscriptionID string) CloudAppliancesClient {
	return original.NewCloudAppliancesClientWithBaseURI(baseURI, subscriptionID)
}

type DevicesClient = original.DevicesClient

func NewDevicesClient(subscriptionID string) DevicesClient {
	return original.NewDevicesClient(subscriptionID)
}
func NewDevicesClientWithBaseURI(baseURI string, subscriptionID string) DevicesClient {
	return original.NewDevicesClientWithBaseURI(baseURI, subscriptionID)
}

type DeviceSettingsClient = original.DeviceSettingsClient

func NewDeviceSettingsClient(subscriptionID string) DeviceSettingsClient {
	return original.NewDeviceSettingsClient(subscriptionID)
}
func NewDeviceSettingsClientWithBaseURI(baseURI string, subscriptionID string) DeviceSettingsClient {
	return original.NewDeviceSettingsClientWithBaseURI(baseURI, subscriptionID)
}

type HardwareComponentGroupsClient = original.HardwareComponentGroupsClient

func NewHardwareComponentGroupsClient(subscriptionID string) HardwareComponentGroupsClient {
	return original.NewHardwareComponentGroupsClient(subscriptionID)
}
func NewHardwareComponentGroupsClientWithBaseURI(baseURI string, subscriptionID string) HardwareComponentGroupsClient {
	return original.NewHardwareComponentGroupsClientWithBaseURI(baseURI, subscriptionID)
}

type JobsClient = original.JobsClient

func NewJobsClient(subscriptionID string) JobsClient {
	return original.NewJobsClient(subscriptionID)
}
func NewJobsClientWithBaseURI(baseURI string, subscriptionID string) JobsClient {
	return original.NewJobsClientWithBaseURI(baseURI, subscriptionID)
}

type ManagersClient = original.ManagersClient

func NewManagersClient(subscriptionID string) ManagersClient {
	return original.NewManagersClient(subscriptionID)
}
func NewManagersClientWithBaseURI(baseURI string, subscriptionID string) ManagersClient {
	return original.NewManagersClientWithBaseURI(baseURI, subscriptionID)
}

type AlertEmailNotificationStatus = original.AlertEmailNotificationStatus

const (
	Disabled AlertEmailNotificationStatus = original.Disabled
	Enabled  AlertEmailNotificationStatus = original.Enabled
)

func PossibleAlertEmailNotificationStatusValues() []AlertEmailNotificationStatus {
	return original.PossibleAlertEmailNotificationStatusValues()
}

type AlertScope = original.AlertScope

const (
	AlertScopeDevice   AlertScope = original.AlertScopeDevice
	AlertScopeResource AlertScope = original.AlertScopeResource
)

func PossibleAlertScopeValues() []AlertScope {
	return original.PossibleAlertScopeValues()
}

type AlertSeverity = original.AlertSeverity

const (
	Critical      AlertSeverity = original.Critical
	Informational AlertSeverity = original.Informational
	Warning       AlertSeverity = original.Warning
)

func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}

type AlertSourceType = original.AlertSourceType

const (
	AlertSourceTypeDevice   AlertSourceType = original.AlertSourceTypeDevice
	AlertSourceTypeResource AlertSourceType = original.AlertSourceTypeResource
)

func PossibleAlertSourceTypeValues() []AlertSourceType {
	return original.PossibleAlertSourceTypeValues()
}

type AlertStatus = original.AlertStatus

const (
	Active  AlertStatus = original.Active
	Cleared AlertStatus = original.Cleared
)

func PossibleAlertStatusValues() []AlertStatus {
	return original.PossibleAlertStatusValues()
}

type AuthenticationType = original.AuthenticationType

const (
	Basic   AuthenticationType = original.Basic
	Invalid AuthenticationType = original.Invalid
	None    AuthenticationType = original.None
	NTLM    AuthenticationType = original.NTLM
)

func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}

type AuthorizationEligibility = original.AuthorizationEligibility

const (
	Eligible   AuthorizationEligibility = original.Eligible
	InEligible AuthorizationEligibility = original.InEligible
)

func PossibleAuthorizationEligibilityValues() []AuthorizationEligibility {
	return original.PossibleAuthorizationEligibilityValues()
}

type AuthorizationStatus = original.AuthorizationStatus

const (
	AuthorizationStatusDisabled AuthorizationStatus = original.AuthorizationStatusDisabled
	AuthorizationStatusEnabled  AuthorizationStatus = original.AuthorizationStatusEnabled
)

func PossibleAuthorizationStatusValues() []AuthorizationStatus {
	return original.PossibleAuthorizationStatusValues()
}

type BackupJobCreationType = original.BackupJobCreationType

const (
	Adhoc      BackupJobCreationType = original.Adhoc
	BySchedule BackupJobCreationType = original.BySchedule
	BySSM      BackupJobCreationType = original.BySSM
)

func PossibleBackupJobCreationTypeValues() []BackupJobCreationType {
	return original.PossibleBackupJobCreationTypeValues()
}

type BackupPolicyCreationType = original.BackupPolicyCreationType

const (
	BackupPolicyCreationTypeBySaaS BackupPolicyCreationType = original.BackupPolicyCreationTypeBySaaS
	BackupPolicyCreationTypeBySSM  BackupPolicyCreationType = original.BackupPolicyCreationTypeBySSM
)

func PossibleBackupPolicyCreationTypeValues() []BackupPolicyCreationType {
	return original.PossibleBackupPolicyCreationTypeValues()
}

type BackupStatus = original.BackupStatus

const (
	BackupStatusDisabled BackupStatus = original.BackupStatusDisabled
	BackupStatusEnabled  BackupStatus = original.BackupStatusEnabled
)

func PossibleBackupStatusValues() []BackupStatus {
	return original.PossibleBackupStatusValues()
}

type BackupType = original.BackupType

const (
	CloudSnapshot BackupType = original.CloudSnapshot
	LocalSnapshot BackupType = original.LocalSnapshot
)

func PossibleBackupTypeValues() []BackupType {
	return original.PossibleBackupTypeValues()
}

type ControllerID = original.ControllerID

const (
	ControllerIDController0 ControllerID = original.ControllerIDController0
	ControllerIDController1 ControllerID = original.ControllerIDController1
	ControllerIDNone        ControllerID = original.ControllerIDNone
	ControllerIDUnknown     ControllerID = original.ControllerIDUnknown
)

func PossibleControllerIDValues() []ControllerID {
	return original.PossibleControllerIDValues()
}

type ControllerPowerStateAction = original.ControllerPowerStateAction

const (
	Restart  ControllerPowerStateAction = original.Restart
	Shutdown ControllerPowerStateAction = original.Shutdown
	Start    ControllerPowerStateAction = original.Start
)

func PossibleControllerPowerStateActionValues() []ControllerPowerStateAction {
	return original.PossibleControllerPowerStateActionValues()
}

type ControllerStatus = original.ControllerStatus

const (
	ControllerStatusFailure    ControllerStatus = original.ControllerStatusFailure
	ControllerStatusNotPresent ControllerStatus = original.ControllerStatusNotPresent
	ControllerStatusOk         ControllerStatus = original.ControllerStatusOk
	ControllerStatusPoweredOff ControllerStatus = original.ControllerStatusPoweredOff
	ControllerStatusRecovering ControllerStatus = original.ControllerStatusRecovering
	ControllerStatusWarning    ControllerStatus = original.ControllerStatusWarning
)

func PossibleControllerStatusValues() []ControllerStatus {
	return original.PossibleControllerStatusValues()
}

type DayOfWeek = original.DayOfWeek

const (
	Friday    DayOfWeek = original.Friday
	Monday    DayOfWeek = original.Monday
	Saturday  DayOfWeek = original.Saturday
	Sunday    DayOfWeek = original.Sunday
	Thursday  DayOfWeek = original.Thursday
	Tuesday   DayOfWeek = original.Tuesday
	Wednesday DayOfWeek = original.Wednesday
)

func PossibleDayOfWeekValues() []DayOfWeek {
	return original.PossibleDayOfWeekValues()
}

type DeviceConfigurationStatus = original.DeviceConfigurationStatus

const (
	Complete DeviceConfigurationStatus = original.Complete
	Pending  DeviceConfigurationStatus = original.Pending
)

func PossibleDeviceConfigurationStatusValues() []DeviceConfigurationStatus {
	return original.PossibleDeviceConfigurationStatusValues()
}

type DeviceStatus = original.DeviceStatus

const (
	Creating          DeviceStatus = original.Creating
	Deactivated       DeviceStatus = original.Deactivated
	Deactivating      DeviceStatus = original.Deactivating
	Deleted           DeviceStatus = original.Deleted
	MaintenanceMode   DeviceStatus = original.MaintenanceMode
	Offline           DeviceStatus = original.Offline
	Online            DeviceStatus = original.Online
	Provisioning      DeviceStatus = original.Provisioning
	ReadyToSetup      DeviceStatus = original.ReadyToSetup
	RequiresAttention DeviceStatus = original.RequiresAttention
	Unknown           DeviceStatus = original.Unknown
)

func PossibleDeviceStatusValues() []DeviceStatus {
	return original.PossibleDeviceStatusValues()
}

type DeviceType = original.DeviceType

const (
	DeviceTypeInvalid                     DeviceType = original.DeviceTypeInvalid
	DeviceTypeSeries8000PhysicalAppliance DeviceType = original.DeviceTypeSeries8000PhysicalAppliance
	DeviceTypeSeries8000VirtualAppliance  DeviceType = original.DeviceTypeSeries8000VirtualAppliance
)

func PossibleDeviceTypeValues() []DeviceType {
	return original.PossibleDeviceTypeValues()
}

type EncryptionAlgorithm = original.EncryptionAlgorithm

const (
	EncryptionAlgorithmAES256        EncryptionAlgorithm = original.EncryptionAlgorithmAES256
	EncryptionAlgorithmNone          EncryptionAlgorithm = original.EncryptionAlgorithmNone
	EncryptionAlgorithmRSAESPKCS1V15 EncryptionAlgorithm = original.EncryptionAlgorithmRSAESPKCS1V15
)

func PossibleEncryptionAlgorithmValues() []EncryptionAlgorithm {
	return original.PossibleEncryptionAlgorithmValues()
}

type EncryptionStatus = original.EncryptionStatus

const (
	EncryptionStatusDisabled EncryptionStatus = original.EncryptionStatusDisabled
	EncryptionStatusEnabled  EncryptionStatus = original.EncryptionStatusEnabled
)

func PossibleEncryptionStatusValues() []EncryptionStatus {
	return original.PossibleEncryptionStatusValues()
}

type FeatureSupportStatus = original.FeatureSupportStatus

const (
	NotAvailable             FeatureSupportStatus = original.NotAvailable
	Supported                FeatureSupportStatus = original.Supported
	UnsupportedDeviceVersion FeatureSupportStatus = original.UnsupportedDeviceVersion
)

func PossibleFeatureSupportStatusValues() []FeatureSupportStatus {
	return original.PossibleFeatureSupportStatusValues()
}

type HardwareComponentStatus = original.HardwareComponentStatus

const (
	HardwareComponentStatusFailure    HardwareComponentStatus = original.HardwareComponentStatusFailure
	HardwareComponentStatusNotPresent HardwareComponentStatus = original.HardwareComponentStatusNotPresent
	HardwareComponentStatusOk         HardwareComponentStatus = original.HardwareComponentStatusOk
	HardwareComponentStatusPoweredOff HardwareComponentStatus = original.HardwareComponentStatusPoweredOff
	HardwareComponentStatusRecovering HardwareComponentStatus = original.HardwareComponentStatusRecovering
	HardwareComponentStatusUnknown    HardwareComponentStatus = original.HardwareComponentStatusUnknown
	HardwareComponentStatusWarning    HardwareComponentStatus = original.HardwareComponentStatusWarning
)

func PossibleHardwareComponentStatusValues() []HardwareComponentStatus {
	return original.PossibleHardwareComponentStatusValues()
}

type InEligibilityCategory = original.InEligibilityCategory

const (
	DeviceNotOnline       InEligibilityCategory = original.DeviceNotOnline
	NotSupportedAppliance InEligibilityCategory = original.NotSupportedAppliance
	RolloverPending       InEligibilityCategory = original.RolloverPending
)

func PossibleInEligibilityCategoryValues() []InEligibilityCategory {
	return original.PossibleInEligibilityCategoryValues()
}

type ISCSIAndCloudStatus = original.ISCSIAndCloudStatus

const (
	ISCSIAndCloudStatusCloudEnabled         ISCSIAndCloudStatus = original.ISCSIAndCloudStatusCloudEnabled
	ISCSIAndCloudStatusDisabled             ISCSIAndCloudStatus = original.ISCSIAndCloudStatusDisabled
	ISCSIAndCloudStatusIscsiAndCloudEnabled ISCSIAndCloudStatus = original.ISCSIAndCloudStatusIscsiAndCloudEnabled
	ISCSIAndCloudStatusIscsiEnabled         ISCSIAndCloudStatus = original.ISCSIAndCloudStatusIscsiEnabled
)

func PossibleISCSIAndCloudStatusValues() []ISCSIAndCloudStatus {
	return original.PossibleISCSIAndCloudStatusValues()
}

type JobStatus = original.JobStatus

const (
	Canceled  JobStatus = original.Canceled
	Failed    JobStatus = original.Failed
	Running   JobStatus = original.Running
	Succeeded JobStatus = original.Succeeded
)

func PossibleJobStatusValues() []JobStatus {
	return original.PossibleJobStatusValues()
}

type JobType = original.JobType

const (
	CloneVolume               JobType = original.CloneVolume
	CreateCloudAppliance      JobType = original.CreateCloudAppliance
	CreateLocallyPinnedVolume JobType = original.CreateLocallyPinnedVolume
	FailoverVolumeContainers  JobType = original.FailoverVolumeContainers
	InstallUpdates            JobType = original.InstallUpdates
	ManualBackup              JobType = original.ManualBackup
	ModifyVolume              JobType = original.ModifyVolume
	RestoreBackup             JobType = original.RestoreBackup
	ScheduledBackup           JobType = original.ScheduledBackup
	SupportPackageLogs        JobType = original.SupportPackageLogs
)

func PossibleJobTypeValues() []JobType {
	return original.PossibleJobTypeValues()
}

type KeyRolloverStatus = original.KeyRolloverStatus

const (
	NotRequired KeyRolloverStatus = original.NotRequired
	Required    KeyRolloverStatus = original.Required
)

func PossibleKeyRolloverStatusValues() []KeyRolloverStatus {
	return original.PossibleKeyRolloverStatusValues()
}

type Kind = original.Kind

const (
	Series8000 Kind = original.Series8000
)

func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}

type ManagerType = original.ManagerType

const (
	GardaV1    ManagerType = original.GardaV1
	HelsinkiV1 ManagerType = original.HelsinkiV1
)

func PossibleManagerTypeValues() []ManagerType {
	return original.PossibleManagerTypeValues()
}

type MetricAggregationType = original.MetricAggregationType

const (
	MetricAggregationTypeAverage MetricAggregationType = original.MetricAggregationTypeAverage
	MetricAggregationTypeLast    MetricAggregationType = original.MetricAggregationTypeLast
	MetricAggregationTypeMaximum MetricAggregationType = original.MetricAggregationTypeMaximum
	MetricAggregationTypeMinimum MetricAggregationType = original.MetricAggregationTypeMinimum
	MetricAggregationTypeNone    MetricAggregationType = original.MetricAggregationTypeNone
	MetricAggregationTypeTotal   MetricAggregationType = original.MetricAggregationTypeTotal
)

func PossibleMetricAggregationTypeValues() []MetricAggregationType {
	return original.PossibleMetricAggregationTypeValues()
}

type MetricUnit = original.MetricUnit

const (
	Bytes          MetricUnit = original.Bytes
	BytesPerSecond MetricUnit = original.BytesPerSecond
	Count          MetricUnit = original.Count
	CountPerSecond MetricUnit = original.CountPerSecond
	Percent        MetricUnit = original.Percent
	Seconds        MetricUnit = original.Seconds
)

func PossibleMetricUnitValues() []MetricUnit {
	return original.PossibleMetricUnitValues()
}

type MonitoringStatus = original.MonitoringStatus

const (
	MonitoringStatusDisabled MonitoringStatus = original.MonitoringStatusDisabled
	MonitoringStatusEnabled  MonitoringStatus = original.MonitoringStatusEnabled
)

func PossibleMonitoringStatusValues() []MonitoringStatus {
	return original.PossibleMonitoringStatusValues()
}

type NetInterfaceID = original.NetInterfaceID

const (
	NetInterfaceIDData0   NetInterfaceID = original.NetInterfaceIDData0
	NetInterfaceIDData1   NetInterfaceID = original.NetInterfaceIDData1
	NetInterfaceIDData2   NetInterfaceID = original.NetInterfaceIDData2
	NetInterfaceIDData3   NetInterfaceID = original.NetInterfaceIDData3
	NetInterfaceIDData4   NetInterfaceID = original.NetInterfaceIDData4
	NetInterfaceIDData5   NetInterfaceID = original.NetInterfaceIDData5
	NetInterfaceIDInvalid NetInterfaceID = original.NetInterfaceIDInvalid
)

func PossibleNetInterfaceIDValues() []NetInterfaceID {
	return original.PossibleNetInterfaceIDValues()
}

type NetInterfaceStatus = original.NetInterfaceStatus

const (
	NetInterfaceStatusDisabled NetInterfaceStatus = original.NetInterfaceStatusDisabled
	NetInterfaceStatusEnabled  NetInterfaceStatus = original.NetInterfaceStatusEnabled
)

func PossibleNetInterfaceStatusValues() []NetInterfaceStatus {
	return original.PossibleNetInterfaceStatusValues()
}

type NetworkMode = original.NetworkMode

const (
	NetworkModeBOTH    NetworkMode = original.NetworkModeBOTH
	NetworkModeInvalid NetworkMode = original.NetworkModeInvalid
	NetworkModeIPV4    NetworkMode = original.NetworkModeIPV4
	NetworkModeIPV6    NetworkMode = original.NetworkModeIPV6
)

func PossibleNetworkModeValues() []NetworkMode {
	return original.PossibleNetworkModeValues()
}

type OperationStatus = original.OperationStatus

const (
	OperationStatusDeleting  OperationStatus = original.OperationStatusDeleting
	OperationStatusNone      OperationStatus = original.OperationStatusNone
	OperationStatusRestoring OperationStatus = original.OperationStatusRestoring
	OperationStatusUpdating  OperationStatus = original.OperationStatusUpdating
)

func PossibleOperationStatusValues() []OperationStatus {
	return original.PossibleOperationStatusValues()
}

type OwnerShipStatus = original.OwnerShipStatus

const (
	NotOwned OwnerShipStatus = original.NotOwned
	Owned    OwnerShipStatus = original.Owned
)

func PossibleOwnerShipStatusValues() []OwnerShipStatus {
	return original.PossibleOwnerShipStatusValues()
}

type RecurrenceType = original.RecurrenceType

const (
	Daily   RecurrenceType = original.Daily
	Hourly  RecurrenceType = original.Hourly
	Minutes RecurrenceType = original.Minutes
	Weekly  RecurrenceType = original.Weekly
)

func PossibleRecurrenceTypeValues() []RecurrenceType {
	return original.PossibleRecurrenceTypeValues()
}

type RemoteManagementModeConfiguration = original.RemoteManagementModeConfiguration

const (
	RemoteManagementModeConfigurationDisabled            RemoteManagementModeConfiguration = original.RemoteManagementModeConfigurationDisabled
	RemoteManagementModeConfigurationHTTPSAndHTTPEnabled RemoteManagementModeConfiguration = original.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled
	RemoteManagementModeConfigurationHTTPSEnabled        RemoteManagementModeConfiguration = original.RemoteManagementModeConfigurationHTTPSEnabled
	RemoteManagementModeConfigurationUnknown             RemoteManagementModeConfiguration = original.RemoteManagementModeConfigurationUnknown
)

func PossibleRemoteManagementModeConfigurationValues() []RemoteManagementModeConfiguration {
	return original.PossibleRemoteManagementModeConfigurationValues()
}

type ScheduledBackupStatus = original.ScheduledBackupStatus

const (
	ScheduledBackupStatusDisabled ScheduledBackupStatus = original.ScheduledBackupStatusDisabled
	ScheduledBackupStatusEnabled  ScheduledBackupStatus = original.ScheduledBackupStatusEnabled
)

func PossibleScheduledBackupStatusValues() []ScheduledBackupStatus {
	return original.PossibleScheduledBackupStatusValues()
}

type ScheduleStatus = original.ScheduleStatus

const (
	ScheduleStatusDisabled ScheduleStatus = original.ScheduleStatusDisabled
	ScheduleStatusEnabled  ScheduleStatus = original.ScheduleStatusEnabled
)

func PossibleScheduleStatusValues() []ScheduleStatus {
	return original.PossibleScheduleStatusValues()
}

type SslStatus = original.SslStatus

const (
	SslStatusDisabled SslStatus = original.SslStatusDisabled
	SslStatusEnabled  SslStatus = original.SslStatusEnabled
)

func PossibleSslStatusValues() []SslStatus {
	return original.PossibleSslStatusValues()
}

type TargetEligibilityResultCode = original.TargetEligibilityResultCode

const (
	LocalToTieredVolumesConversionWarning     TargetEligibilityResultCode = original.LocalToTieredVolumesConversionWarning
	TargetAndSourceCannotBeSameError          TargetEligibilityResultCode = original.TargetAndSourceCannotBeSameError
	TargetInsufficientCapacityError           TargetEligibilityResultCode = original.TargetInsufficientCapacityError
	TargetInsufficientLocalVolumeMemoryError  TargetEligibilityResultCode = original.TargetInsufficientLocalVolumeMemoryError
	TargetInsufficientTieredVolumeMemoryError TargetEligibilityResultCode = original.TargetInsufficientTieredVolumeMemoryError
	TargetIsNotOnlineError                    TargetEligibilityResultCode = original.TargetIsNotOnlineError
	TargetSourceIncompatibleVersionError      TargetEligibilityResultCode = original.TargetSourceIncompatibleVersionError
)

func PossibleTargetEligibilityResultCodeValues() []TargetEligibilityResultCode {
	return original.PossibleTargetEligibilityResultCodeValues()
}

type TargetEligibilityStatus = original.TargetEligibilityStatus

const (
	TargetEligibilityStatusEligible    TargetEligibilityStatus = original.TargetEligibilityStatusEligible
	TargetEligibilityStatusNotEligible TargetEligibilityStatus = original.TargetEligibilityStatusNotEligible
)

func PossibleTargetEligibilityStatusValues() []TargetEligibilityStatus {
	return original.PossibleTargetEligibilityStatusValues()
}

type VirtualMachineAPIType = original.VirtualMachineAPIType

const (
	Arm     VirtualMachineAPIType = original.Arm
	Classic VirtualMachineAPIType = original.Classic
)

func PossibleVirtualMachineAPITypeValues() []VirtualMachineAPIType {
	return original.PossibleVirtualMachineAPITypeValues()
}

type VolumeStatus = original.VolumeStatus

const (
	VolumeStatusOffline VolumeStatus = original.VolumeStatusOffline
	VolumeStatusOnline  VolumeStatus = original.VolumeStatusOnline
)

func PossibleVolumeStatusValues() []VolumeStatus {
	return original.PossibleVolumeStatusValues()
}

type VolumeType = original.VolumeType

const (
	Archival      VolumeType = original.Archival
	LocallyPinned VolumeType = original.LocallyPinned
	Tiered        VolumeType = original.Tiered
)

func PossibleVolumeTypeValues() []VolumeType {
	return original.PossibleVolumeTypeValues()
}

type AccessControlRecord = original.AccessControlRecord
type AccessControlRecordList = original.AccessControlRecordList
type AccessControlRecordProperties = original.AccessControlRecordProperties
type AccessControlRecordsCreateOrUpdateFuture = original.AccessControlRecordsCreateOrUpdateFuture
type AccessControlRecordsDeleteFuture = original.AccessControlRecordsDeleteFuture
type AcsConfiguration = original.AcsConfiguration
type Alert = original.Alert
type AlertErrorDetails = original.AlertErrorDetails
type AlertFilter = original.AlertFilter
type AlertList = original.AlertList
type AlertListIterator = original.AlertListIterator
type AlertListPage = original.AlertListPage
type AlertNotificationProperties = original.AlertNotificationProperties
type AlertProperties = original.AlertProperties
type AlertSettings = original.AlertSettings
type AlertSource = original.AlertSource
type AsymmetricEncryptedSecret = original.AsymmetricEncryptedSecret
type AvailableProviderOperation = original.AvailableProviderOperation
type AvailableProviderOperationDisplay = original.AvailableProviderOperationDisplay
type AvailableProviderOperationList = original.AvailableProviderOperationList
type AvailableProviderOperationListIterator = original.AvailableProviderOperationListIterator
type AvailableProviderOperationListPage = original.AvailableProviderOperationListPage
type Backup = original.Backup
type BackupElement = original.BackupElement
type BackupFilter = original.BackupFilter
type BackupList = original.BackupList
type BackupListIterator = original.BackupListIterator
type BackupListPage = original.BackupListPage
type BackupPoliciesBackupNowFuture = original.BackupPoliciesBackupNowFuture
type BackupPoliciesCreateOrUpdateFuture = original.BackupPoliciesCreateOrUpdateFuture
type BackupPoliciesDeleteFuture = original.BackupPoliciesDeleteFuture
type BackupPolicy = original.BackupPolicy
type BackupPolicyList = original.BackupPolicyList
type BackupPolicyProperties = original.BackupPolicyProperties
type BackupProperties = original.BackupProperties
type BackupSchedule = original.BackupSchedule
type BackupScheduleList = original.BackupScheduleList
type BackupScheduleProperties = original.BackupScheduleProperties
type BackupSchedulesCreateOrUpdateFuture = original.BackupSchedulesCreateOrUpdateFuture
type BackupSchedulesDeleteFuture = original.BackupSchedulesDeleteFuture
type BackupsCloneFuture = original.BackupsCloneFuture
type BackupsDeleteFuture = original.BackupsDeleteFuture
type BackupsRestoreFuture = original.BackupsRestoreFuture
type BandwidthRateSettingProperties = original.BandwidthRateSettingProperties
type BandwidthSchedule = original.BandwidthSchedule
type BandwidthSetting = original.BandwidthSetting
type BandwidthSettingList = original.BandwidthSettingList
type BandwidthSettingsCreateOrUpdateFuture = original.BandwidthSettingsCreateOrUpdateFuture
type BandwidthSettingsDeleteFuture = original.BandwidthSettingsDeleteFuture
type BaseModel = original.BaseModel
type ChapSettings = original.ChapSettings
type ClearAlertRequest = original.ClearAlertRequest
type CloneRequest = original.CloneRequest
type CloudAppliance = original.CloudAppliance
type CloudApplianceConfiguration = original.CloudApplianceConfiguration
type CloudApplianceConfigurationList = original.CloudApplianceConfigurationList
type CloudApplianceConfigurationProperties = original.CloudApplianceConfigurationProperties
type CloudApplianceSettings = original.CloudApplianceSettings
type CloudAppliancesProvisionFuture = original.CloudAppliancesProvisionFuture
type ConfigureDeviceRequest = original.ConfigureDeviceRequest
type ConfigureDeviceRequestProperties = original.ConfigureDeviceRequestProperties
type ControllerPowerStateChangeRequest = original.ControllerPowerStateChangeRequest
type ControllerPowerStateChangeRequestProperties = original.ControllerPowerStateChangeRequestProperties
type DataStatistics = original.DataStatistics
type Device = original.Device
type DeviceDetails = original.DeviceDetails
type DeviceList = original.DeviceList
type DevicePatch = original.DevicePatch
type DevicePatchProperties = original.DevicePatchProperties
type DeviceProperties = original.DeviceProperties
type DeviceRolloverDetails = original.DeviceRolloverDetails
type DevicesConfigureFuture = original.DevicesConfigureFuture
type DevicesDeactivateFuture = original.DevicesDeactivateFuture
type DevicesDeleteFuture = original.DevicesDeleteFuture
type DeviceSettingsCreateOrUpdateAlertSettingsFuture = original.DeviceSettingsCreateOrUpdateAlertSettingsFuture
type DeviceSettingsCreateOrUpdateTimeSettingsFuture = original.DeviceSettingsCreateOrUpdateTimeSettingsFuture
type DeviceSettingsSyncRemotemanagementCertificateFuture = original.DeviceSettingsSyncRemotemanagementCertificateFuture
type DeviceSettingsUpdateNetworkSettingsFuture = original.DeviceSettingsUpdateNetworkSettingsFuture
type DeviceSettingsUpdateSecuritySettingsFuture = original.DeviceSettingsUpdateSecuritySettingsFuture
type DevicesFailoverFuture = original.DevicesFailoverFuture
type DevicesInstallUpdatesFuture = original.DevicesInstallUpdatesFuture
type DevicesScanForUpdatesFuture = original.DevicesScanForUpdatesFuture
type DimensionFilter = original.DimensionFilter
type DNSSettings = original.DNSSettings
type EncryptionSettings = original.EncryptionSettings
type EncryptionSettingsProperties = original.EncryptionSettingsProperties
type FailoverRequest = original.FailoverRequest
type FailoverSet = original.FailoverSet
type FailoverSetEligibilityResult = original.FailoverSetEligibilityResult
type FailoverSetsList = original.FailoverSetsList
type FailoverTarget = original.FailoverTarget
type FailoverTargetsList = original.FailoverTargetsList
type Feature = original.Feature
type FeatureFilter = original.FeatureFilter
type FeatureList = original.FeatureList
type HardwareComponent = original.HardwareComponent
type HardwareComponentGroup = original.HardwareComponentGroup
type HardwareComponentGroupList = original.HardwareComponentGroupList
type HardwareComponentGroupProperties = original.HardwareComponentGroupProperties
type HardwareComponentGroupsChangeControllerPowerStateFuture = original.HardwareComponentGroupsChangeControllerPowerStateFuture
type Job = original.Job
type JobErrorDetails = original.JobErrorDetails
type JobErrorItem = original.JobErrorItem
type JobFilter = original.JobFilter
type JobList = original.JobList
type JobListIterator = original.JobListIterator
type JobListPage = original.JobListPage
type JobProperties = original.JobProperties
type JobsCancelFuture = original.JobsCancelFuture
type JobStage = original.JobStage
type Key = original.Key
type ListFailoverTargetsRequest = original.ListFailoverTargetsRequest
type Manager = original.Manager
type ManagerExtendedInfo = original.ManagerExtendedInfo
type ManagerExtendedInfoProperties = original.ManagerExtendedInfoProperties
type ManagerIntrinsicSettings = original.ManagerIntrinsicSettings
type ManagerList = original.ManagerList
type ManagerPatch = original.ManagerPatch
type ManagerProperties = original.ManagerProperties
type ManagerSku = original.ManagerSku
type MetricAvailablity = original.MetricAvailablity
type MetricData = original.MetricData
type MetricDefinition = original.MetricDefinition
type MetricDefinitionList = original.MetricDefinitionList
type MetricDimension = original.MetricDimension
type MetricFilter = original.MetricFilter
type MetricList = original.MetricList
type MetricName = original.MetricName
type MetricNameFilter = original.MetricNameFilter
type Metrics = original.Metrics
type NetworkAdapterList = original.NetworkAdapterList
type NetworkAdapters = original.NetworkAdapters
type NetworkInterfaceData0Settings = original.NetworkInterfaceData0Settings
type NetworkSettings = original.NetworkSettings
type NetworkSettingsPatch = original.NetworkSettingsPatch
type NetworkSettingsPatchProperties = original.NetworkSettingsPatchProperties
type NetworkSettingsProperties = original.NetworkSettingsProperties
type NicIPv4 = original.NicIPv4
type NicIPv6 = original.NicIPv6
type PublicKey = original.PublicKey
type RemoteManagementSettings = original.RemoteManagementSettings
type RemoteManagementSettingsPatch = original.RemoteManagementSettingsPatch
type Resource = original.Resource
type ScheduleRecurrence = original.ScheduleRecurrence
type SecondaryDNSSettings = original.SecondaryDNSSettings
type SecuritySettings = original.SecuritySettings
type SecuritySettingsPatch = original.SecuritySettingsPatch
type SecuritySettingsPatchProperties = original.SecuritySettingsPatchProperties
type SecuritySettingsProperties = original.SecuritySettingsProperties
type SendTestAlertEmailRequest = original.SendTestAlertEmailRequest
type StorageAccountCredential = original.StorageAccountCredential
type StorageAccountCredentialList = original.StorageAccountCredentialList
type StorageAccountCredentialProperties = original.StorageAccountCredentialProperties
type StorageAccountCredentialsCreateOrUpdateFuture = original.StorageAccountCredentialsCreateOrUpdateFuture
type StorageAccountCredentialsDeleteFuture = original.StorageAccountCredentialsDeleteFuture
type SymmetricEncryptedSecret = original.SymmetricEncryptedSecret
type TargetEligibilityErrorMessage = original.TargetEligibilityErrorMessage
type TargetEligibilityResult = original.TargetEligibilityResult
type Time = original.Time
type TimeSettings = original.TimeSettings
type TimeSettingsProperties = original.TimeSettingsProperties
type Updates = original.Updates
type UpdatesProperties = original.UpdatesProperties
type VMImage = original.VMImage
type Volume = original.Volume
type VolumeContainer = original.VolumeContainer
type VolumeContainerFailoverMetadata = original.VolumeContainerFailoverMetadata
type VolumeContainerList = original.VolumeContainerList
type VolumeContainerProperties = original.VolumeContainerProperties
type VolumeContainersCreateOrUpdateFuture = original.VolumeContainersCreateOrUpdateFuture
type VolumeContainersDeleteFuture = original.VolumeContainersDeleteFuture
type VolumeFailoverMetadata = original.VolumeFailoverMetadata
type VolumeList = original.VolumeList
type VolumeProperties = original.VolumeProperties
type VolumesCreateOrUpdateFuture = original.VolumesCreateOrUpdateFuture
type VolumesDeleteFuture = original.VolumesDeleteFuture
type WebproxySettings = original.WebproxySettings
type OperationsClient = original.OperationsClient

func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}

type StorageAccountCredentialsClient = original.StorageAccountCredentialsClient

func NewStorageAccountCredentialsClient(subscriptionID string) StorageAccountCredentialsClient {
	return original.NewStorageAccountCredentialsClient(subscriptionID)
}
func NewStorageAccountCredentialsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountCredentialsClient {
	return original.NewStorageAccountCredentialsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}

type VolumeContainersClient = original.VolumeContainersClient

func NewVolumeContainersClient(subscriptionID string) VolumeContainersClient {
	return original.NewVolumeContainersClient(subscriptionID)
}
func NewVolumeContainersClientWithBaseURI(baseURI string, subscriptionID string) VolumeContainersClient {
	return original.NewVolumeContainersClientWithBaseURI(baseURI, subscriptionID)
}

type VolumesClient = original.VolumesClient

func NewVolumesClient(subscriptionID string) VolumesClient {
	return original.NewVolumesClient(subscriptionID)
}
func NewVolumesClientWithBaseURI(baseURI string, subscriptionID string) VolumesClient {
	return original.NewVolumesClientWithBaseURI(baseURI, subscriptionID)
}