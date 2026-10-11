# ProjectActionVisibilityDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**Repair** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**Upgrade** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableMonitoring** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableMonitoring** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableBackup** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableBackup** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableOpa** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableOpa** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableAutoscaler** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableAutoscaler** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableGatewayApi** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableGatewayApi** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**VmRepair** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**VmCommit** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**Lock** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**Unlock** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableSpotWorker** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableSpotWorker** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableFullSpot** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableFullSpot** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableSpotVm** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableSpotVm** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**AttachAlertingProfile** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DetachAlertingProfile** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableAi** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableAi** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**AiAssistant** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**ProjectMaintenanceMode** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**AddServer** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**AddVm** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableDnsCert** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableDnsCert** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableCertManager** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableCertManager** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableExternalDns** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableExternalDns** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableTrustManager** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableTrustManager** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**EnableCTProfiles** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**DisableCTProfiles** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 
**LetsEncryptCertificateProfileBlockedReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewProjectActionVisibilityDto

`func NewProjectActionVisibilityDto(commit ButtonStatusDto, repair ButtonStatusDto, upgrade ButtonStatusDto, enableMonitoring ButtonStatusDto, disableMonitoring ButtonStatusDto, enableBackup ButtonStatusDto, disableBackup ButtonStatusDto, enableOpa ButtonStatusDto, disableOpa ButtonStatusDto, enableAutoscaler ButtonStatusDto, disableAutoscaler ButtonStatusDto, enableGatewayApi ButtonStatusDto, disableGatewayApi ButtonStatusDto, vmRepair ButtonStatusDto, vmCommit ButtonStatusDto, lock ButtonStatusDto, unlock ButtonStatusDto, enableSpotWorker ButtonStatusDto, disableSpotWorker ButtonStatusDto, enableFullSpot ButtonStatusDto, disableFullSpot ButtonStatusDto, enableSpotVm ButtonStatusDto, disableSpotVm ButtonStatusDto, attachAlertingProfile ButtonStatusDto, detachAlertingProfile ButtonStatusDto, enableAi ButtonStatusDto, disableAi ButtonStatusDto, aiAssistant ButtonStatusDto, projectMaintenanceMode ButtonStatusDto, addServer ButtonStatusDto, addVm ButtonStatusDto, enableDnsCert ButtonStatusDto, disableDnsCert ButtonStatusDto, enableCertManager ButtonStatusDto, disableCertManager ButtonStatusDto, enableExternalDns ButtonStatusDto, disableExternalDns ButtonStatusDto, enableTrustManager ButtonStatusDto, disableTrustManager ButtonStatusDto, enableCTProfiles ButtonStatusDto, disableCTProfiles ButtonStatusDto, ) *ProjectActionVisibilityDto`

NewProjectActionVisibilityDto instantiates a new ProjectActionVisibilityDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectActionVisibilityDtoWithDefaults

`func NewProjectActionVisibilityDtoWithDefaults() *ProjectActionVisibilityDto`

NewProjectActionVisibilityDtoWithDefaults instantiates a new ProjectActionVisibilityDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *ProjectActionVisibilityDto) GetCommit() ButtonStatusDto`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ProjectActionVisibilityDto) GetCommitOk() (*ButtonStatusDto, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ProjectActionVisibilityDto) SetCommit(v ButtonStatusDto)`

SetCommit sets Commit field to given value.


### GetRepair

`func (o *ProjectActionVisibilityDto) GetRepair() ButtonStatusDto`

GetRepair returns the Repair field if non-nil, zero value otherwise.

### GetRepairOk

`func (o *ProjectActionVisibilityDto) GetRepairOk() (*ButtonStatusDto, bool)`

GetRepairOk returns a tuple with the Repair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepair

`func (o *ProjectActionVisibilityDto) SetRepair(v ButtonStatusDto)`

SetRepair sets Repair field to given value.


### GetUpgrade

`func (o *ProjectActionVisibilityDto) GetUpgrade() ButtonStatusDto`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *ProjectActionVisibilityDto) GetUpgradeOk() (*ButtonStatusDto, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *ProjectActionVisibilityDto) SetUpgrade(v ButtonStatusDto)`

SetUpgrade sets Upgrade field to given value.


### GetEnableMonitoring

`func (o *ProjectActionVisibilityDto) GetEnableMonitoring() ButtonStatusDto`

GetEnableMonitoring returns the EnableMonitoring field if non-nil, zero value otherwise.

### GetEnableMonitoringOk

`func (o *ProjectActionVisibilityDto) GetEnableMonitoringOk() (*ButtonStatusDto, bool)`

GetEnableMonitoringOk returns a tuple with the EnableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMonitoring

`func (o *ProjectActionVisibilityDto) SetEnableMonitoring(v ButtonStatusDto)`

SetEnableMonitoring sets EnableMonitoring field to given value.


### GetDisableMonitoring

`func (o *ProjectActionVisibilityDto) GetDisableMonitoring() ButtonStatusDto`

GetDisableMonitoring returns the DisableMonitoring field if non-nil, zero value otherwise.

### GetDisableMonitoringOk

`func (o *ProjectActionVisibilityDto) GetDisableMonitoringOk() (*ButtonStatusDto, bool)`

GetDisableMonitoringOk returns a tuple with the DisableMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMonitoring

`func (o *ProjectActionVisibilityDto) SetDisableMonitoring(v ButtonStatusDto)`

SetDisableMonitoring sets DisableMonitoring field to given value.


### GetEnableBackup

`func (o *ProjectActionVisibilityDto) GetEnableBackup() ButtonStatusDto`

GetEnableBackup returns the EnableBackup field if non-nil, zero value otherwise.

### GetEnableBackupOk

`func (o *ProjectActionVisibilityDto) GetEnableBackupOk() (*ButtonStatusDto, bool)`

GetEnableBackupOk returns a tuple with the EnableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBackup

`func (o *ProjectActionVisibilityDto) SetEnableBackup(v ButtonStatusDto)`

SetEnableBackup sets EnableBackup field to given value.


### GetDisableBackup

`func (o *ProjectActionVisibilityDto) GetDisableBackup() ButtonStatusDto`

GetDisableBackup returns the DisableBackup field if non-nil, zero value otherwise.

### GetDisableBackupOk

`func (o *ProjectActionVisibilityDto) GetDisableBackupOk() (*ButtonStatusDto, bool)`

GetDisableBackupOk returns a tuple with the DisableBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableBackup

`func (o *ProjectActionVisibilityDto) SetDisableBackup(v ButtonStatusDto)`

SetDisableBackup sets DisableBackup field to given value.


### GetEnableOpa

`func (o *ProjectActionVisibilityDto) GetEnableOpa() ButtonStatusDto`

GetEnableOpa returns the EnableOpa field if non-nil, zero value otherwise.

### GetEnableOpaOk

`func (o *ProjectActionVisibilityDto) GetEnableOpaOk() (*ButtonStatusDto, bool)`

GetEnableOpaOk returns a tuple with the EnableOpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOpa

`func (o *ProjectActionVisibilityDto) SetEnableOpa(v ButtonStatusDto)`

SetEnableOpa sets EnableOpa field to given value.


### GetDisableOpa

`func (o *ProjectActionVisibilityDto) GetDisableOpa() ButtonStatusDto`

GetDisableOpa returns the DisableOpa field if non-nil, zero value otherwise.

### GetDisableOpaOk

`func (o *ProjectActionVisibilityDto) GetDisableOpaOk() (*ButtonStatusDto, bool)`

GetDisableOpaOk returns a tuple with the DisableOpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOpa

`func (o *ProjectActionVisibilityDto) SetDisableOpa(v ButtonStatusDto)`

SetDisableOpa sets DisableOpa field to given value.


### GetEnableAutoscaler

`func (o *ProjectActionVisibilityDto) GetEnableAutoscaler() ButtonStatusDto`

GetEnableAutoscaler returns the EnableAutoscaler field if non-nil, zero value otherwise.

### GetEnableAutoscalerOk

`func (o *ProjectActionVisibilityDto) GetEnableAutoscalerOk() (*ButtonStatusDto, bool)`

GetEnableAutoscalerOk returns a tuple with the EnableAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoscaler

`func (o *ProjectActionVisibilityDto) SetEnableAutoscaler(v ButtonStatusDto)`

SetEnableAutoscaler sets EnableAutoscaler field to given value.


### GetDisableAutoscaler

`func (o *ProjectActionVisibilityDto) GetDisableAutoscaler() ButtonStatusDto`

GetDisableAutoscaler returns the DisableAutoscaler field if non-nil, zero value otherwise.

### GetDisableAutoscalerOk

`func (o *ProjectActionVisibilityDto) GetDisableAutoscalerOk() (*ButtonStatusDto, bool)`

GetDisableAutoscalerOk returns a tuple with the DisableAutoscaler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoscaler

`func (o *ProjectActionVisibilityDto) SetDisableAutoscaler(v ButtonStatusDto)`

SetDisableAutoscaler sets DisableAutoscaler field to given value.


### GetEnableGatewayApi

`func (o *ProjectActionVisibilityDto) GetEnableGatewayApi() ButtonStatusDto`

GetEnableGatewayApi returns the EnableGatewayApi field if non-nil, zero value otherwise.

### GetEnableGatewayApiOk

`func (o *ProjectActionVisibilityDto) GetEnableGatewayApiOk() (*ButtonStatusDto, bool)`

GetEnableGatewayApiOk returns a tuple with the EnableGatewayApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGatewayApi

`func (o *ProjectActionVisibilityDto) SetEnableGatewayApi(v ButtonStatusDto)`

SetEnableGatewayApi sets EnableGatewayApi field to given value.


### GetDisableGatewayApi

`func (o *ProjectActionVisibilityDto) GetDisableGatewayApi() ButtonStatusDto`

GetDisableGatewayApi returns the DisableGatewayApi field if non-nil, zero value otherwise.

### GetDisableGatewayApiOk

`func (o *ProjectActionVisibilityDto) GetDisableGatewayApiOk() (*ButtonStatusDto, bool)`

GetDisableGatewayApiOk returns a tuple with the DisableGatewayApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableGatewayApi

`func (o *ProjectActionVisibilityDto) SetDisableGatewayApi(v ButtonStatusDto)`

SetDisableGatewayApi sets DisableGatewayApi field to given value.


### GetVmRepair

`func (o *ProjectActionVisibilityDto) GetVmRepair() ButtonStatusDto`

GetVmRepair returns the VmRepair field if non-nil, zero value otherwise.

### GetVmRepairOk

`func (o *ProjectActionVisibilityDto) GetVmRepairOk() (*ButtonStatusDto, bool)`

GetVmRepairOk returns a tuple with the VmRepair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRepair

`func (o *ProjectActionVisibilityDto) SetVmRepair(v ButtonStatusDto)`

SetVmRepair sets VmRepair field to given value.


### GetVmCommit

`func (o *ProjectActionVisibilityDto) GetVmCommit() ButtonStatusDto`

GetVmCommit returns the VmCommit field if non-nil, zero value otherwise.

### GetVmCommitOk

`func (o *ProjectActionVisibilityDto) GetVmCommitOk() (*ButtonStatusDto, bool)`

GetVmCommitOk returns a tuple with the VmCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCommit

`func (o *ProjectActionVisibilityDto) SetVmCommit(v ButtonStatusDto)`

SetVmCommit sets VmCommit field to given value.


### GetLock

`func (o *ProjectActionVisibilityDto) GetLock() ButtonStatusDto`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *ProjectActionVisibilityDto) GetLockOk() (*ButtonStatusDto, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *ProjectActionVisibilityDto) SetLock(v ButtonStatusDto)`

SetLock sets Lock field to given value.


### GetUnlock

`func (o *ProjectActionVisibilityDto) GetUnlock() ButtonStatusDto`

GetUnlock returns the Unlock field if non-nil, zero value otherwise.

### GetUnlockOk

`func (o *ProjectActionVisibilityDto) GetUnlockOk() (*ButtonStatusDto, bool)`

GetUnlockOk returns a tuple with the Unlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlock

`func (o *ProjectActionVisibilityDto) SetUnlock(v ButtonStatusDto)`

SetUnlock sets Unlock field to given value.


### GetEnableSpotWorker

`func (o *ProjectActionVisibilityDto) GetEnableSpotWorker() ButtonStatusDto`

GetEnableSpotWorker returns the EnableSpotWorker field if non-nil, zero value otherwise.

### GetEnableSpotWorkerOk

`func (o *ProjectActionVisibilityDto) GetEnableSpotWorkerOk() (*ButtonStatusDto, bool)`

GetEnableSpotWorkerOk returns a tuple with the EnableSpotWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotWorker

`func (o *ProjectActionVisibilityDto) SetEnableSpotWorker(v ButtonStatusDto)`

SetEnableSpotWorker sets EnableSpotWorker field to given value.


### GetDisableSpotWorker

`func (o *ProjectActionVisibilityDto) GetDisableSpotWorker() ButtonStatusDto`

GetDisableSpotWorker returns the DisableSpotWorker field if non-nil, zero value otherwise.

### GetDisableSpotWorkerOk

`func (o *ProjectActionVisibilityDto) GetDisableSpotWorkerOk() (*ButtonStatusDto, bool)`

GetDisableSpotWorkerOk returns a tuple with the DisableSpotWorker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpotWorker

`func (o *ProjectActionVisibilityDto) SetDisableSpotWorker(v ButtonStatusDto)`

SetDisableSpotWorker sets DisableSpotWorker field to given value.


### GetEnableFullSpot

`func (o *ProjectActionVisibilityDto) GetEnableFullSpot() ButtonStatusDto`

GetEnableFullSpot returns the EnableFullSpot field if non-nil, zero value otherwise.

### GetEnableFullSpotOk

`func (o *ProjectActionVisibilityDto) GetEnableFullSpotOk() (*ButtonStatusDto, bool)`

GetEnableFullSpotOk returns a tuple with the EnableFullSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFullSpot

`func (o *ProjectActionVisibilityDto) SetEnableFullSpot(v ButtonStatusDto)`

SetEnableFullSpot sets EnableFullSpot field to given value.


### GetDisableFullSpot

`func (o *ProjectActionVisibilityDto) GetDisableFullSpot() ButtonStatusDto`

GetDisableFullSpot returns the DisableFullSpot field if non-nil, zero value otherwise.

### GetDisableFullSpotOk

`func (o *ProjectActionVisibilityDto) GetDisableFullSpotOk() (*ButtonStatusDto, bool)`

GetDisableFullSpotOk returns a tuple with the DisableFullSpot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFullSpot

`func (o *ProjectActionVisibilityDto) SetDisableFullSpot(v ButtonStatusDto)`

SetDisableFullSpot sets DisableFullSpot field to given value.


### GetEnableSpotVm

`func (o *ProjectActionVisibilityDto) GetEnableSpotVm() ButtonStatusDto`

GetEnableSpotVm returns the EnableSpotVm field if non-nil, zero value otherwise.

### GetEnableSpotVmOk

`func (o *ProjectActionVisibilityDto) GetEnableSpotVmOk() (*ButtonStatusDto, bool)`

GetEnableSpotVmOk returns a tuple with the EnableSpotVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpotVm

`func (o *ProjectActionVisibilityDto) SetEnableSpotVm(v ButtonStatusDto)`

SetEnableSpotVm sets EnableSpotVm field to given value.


### GetDisableSpotVm

`func (o *ProjectActionVisibilityDto) GetDisableSpotVm() ButtonStatusDto`

GetDisableSpotVm returns the DisableSpotVm field if non-nil, zero value otherwise.

### GetDisableSpotVmOk

`func (o *ProjectActionVisibilityDto) GetDisableSpotVmOk() (*ButtonStatusDto, bool)`

GetDisableSpotVmOk returns a tuple with the DisableSpotVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSpotVm

`func (o *ProjectActionVisibilityDto) SetDisableSpotVm(v ButtonStatusDto)`

SetDisableSpotVm sets DisableSpotVm field to given value.


### GetAttachAlertingProfile

`func (o *ProjectActionVisibilityDto) GetAttachAlertingProfile() ButtonStatusDto`

GetAttachAlertingProfile returns the AttachAlertingProfile field if non-nil, zero value otherwise.

### GetAttachAlertingProfileOk

`func (o *ProjectActionVisibilityDto) GetAttachAlertingProfileOk() (*ButtonStatusDto, bool)`

GetAttachAlertingProfileOk returns a tuple with the AttachAlertingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachAlertingProfile

`func (o *ProjectActionVisibilityDto) SetAttachAlertingProfile(v ButtonStatusDto)`

SetAttachAlertingProfile sets AttachAlertingProfile field to given value.


### GetDetachAlertingProfile

`func (o *ProjectActionVisibilityDto) GetDetachAlertingProfile() ButtonStatusDto`

GetDetachAlertingProfile returns the DetachAlertingProfile field if non-nil, zero value otherwise.

### GetDetachAlertingProfileOk

`func (o *ProjectActionVisibilityDto) GetDetachAlertingProfileOk() (*ButtonStatusDto, bool)`

GetDetachAlertingProfileOk returns a tuple with the DetachAlertingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachAlertingProfile

`func (o *ProjectActionVisibilityDto) SetDetachAlertingProfile(v ButtonStatusDto)`

SetDetachAlertingProfile sets DetachAlertingProfile field to given value.


### GetEnableAi

`func (o *ProjectActionVisibilityDto) GetEnableAi() ButtonStatusDto`

GetEnableAi returns the EnableAi field if non-nil, zero value otherwise.

### GetEnableAiOk

`func (o *ProjectActionVisibilityDto) GetEnableAiOk() (*ButtonStatusDto, bool)`

GetEnableAiOk returns a tuple with the EnableAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAi

`func (o *ProjectActionVisibilityDto) SetEnableAi(v ButtonStatusDto)`

SetEnableAi sets EnableAi field to given value.


### GetDisableAi

`func (o *ProjectActionVisibilityDto) GetDisableAi() ButtonStatusDto`

GetDisableAi returns the DisableAi field if non-nil, zero value otherwise.

### GetDisableAiOk

`func (o *ProjectActionVisibilityDto) GetDisableAiOk() (*ButtonStatusDto, bool)`

GetDisableAiOk returns a tuple with the DisableAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAi

`func (o *ProjectActionVisibilityDto) SetDisableAi(v ButtonStatusDto)`

SetDisableAi sets DisableAi field to given value.


### GetAiAssistant

`func (o *ProjectActionVisibilityDto) GetAiAssistant() ButtonStatusDto`

GetAiAssistant returns the AiAssistant field if non-nil, zero value otherwise.

### GetAiAssistantOk

`func (o *ProjectActionVisibilityDto) GetAiAssistantOk() (*ButtonStatusDto, bool)`

GetAiAssistantOk returns a tuple with the AiAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAssistant

`func (o *ProjectActionVisibilityDto) SetAiAssistant(v ButtonStatusDto)`

SetAiAssistant sets AiAssistant field to given value.


### GetProjectMaintenanceMode

`func (o *ProjectActionVisibilityDto) GetProjectMaintenanceMode() ButtonStatusDto`

GetProjectMaintenanceMode returns the ProjectMaintenanceMode field if non-nil, zero value otherwise.

### GetProjectMaintenanceModeOk

`func (o *ProjectActionVisibilityDto) GetProjectMaintenanceModeOk() (*ButtonStatusDto, bool)`

GetProjectMaintenanceModeOk returns a tuple with the ProjectMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMaintenanceMode

`func (o *ProjectActionVisibilityDto) SetProjectMaintenanceMode(v ButtonStatusDto)`

SetProjectMaintenanceMode sets ProjectMaintenanceMode field to given value.


### GetAddServer

`func (o *ProjectActionVisibilityDto) GetAddServer() ButtonStatusDto`

GetAddServer returns the AddServer field if non-nil, zero value otherwise.

### GetAddServerOk

`func (o *ProjectActionVisibilityDto) GetAddServerOk() (*ButtonStatusDto, bool)`

GetAddServerOk returns a tuple with the AddServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddServer

`func (o *ProjectActionVisibilityDto) SetAddServer(v ButtonStatusDto)`

SetAddServer sets AddServer field to given value.


### GetAddVm

`func (o *ProjectActionVisibilityDto) GetAddVm() ButtonStatusDto`

GetAddVm returns the AddVm field if non-nil, zero value otherwise.

### GetAddVmOk

`func (o *ProjectActionVisibilityDto) GetAddVmOk() (*ButtonStatusDto, bool)`

GetAddVmOk returns a tuple with the AddVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVm

`func (o *ProjectActionVisibilityDto) SetAddVm(v ButtonStatusDto)`

SetAddVm sets AddVm field to given value.


### GetEnableDnsCert

`func (o *ProjectActionVisibilityDto) GetEnableDnsCert() ButtonStatusDto`

GetEnableDnsCert returns the EnableDnsCert field if non-nil, zero value otherwise.

### GetEnableDnsCertOk

`func (o *ProjectActionVisibilityDto) GetEnableDnsCertOk() (*ButtonStatusDto, bool)`

GetEnableDnsCertOk returns a tuple with the EnableDnsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsCert

`func (o *ProjectActionVisibilityDto) SetEnableDnsCert(v ButtonStatusDto)`

SetEnableDnsCert sets EnableDnsCert field to given value.


### GetDisableDnsCert

`func (o *ProjectActionVisibilityDto) GetDisableDnsCert() ButtonStatusDto`

GetDisableDnsCert returns the DisableDnsCert field if non-nil, zero value otherwise.

### GetDisableDnsCertOk

`func (o *ProjectActionVisibilityDto) GetDisableDnsCertOk() (*ButtonStatusDto, bool)`

GetDisableDnsCertOk returns a tuple with the DisableDnsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDnsCert

`func (o *ProjectActionVisibilityDto) SetDisableDnsCert(v ButtonStatusDto)`

SetDisableDnsCert sets DisableDnsCert field to given value.


### GetEnableCertManager

`func (o *ProjectActionVisibilityDto) GetEnableCertManager() ButtonStatusDto`

GetEnableCertManager returns the EnableCertManager field if non-nil, zero value otherwise.

### GetEnableCertManagerOk

`func (o *ProjectActionVisibilityDto) GetEnableCertManagerOk() (*ButtonStatusDto, bool)`

GetEnableCertManagerOk returns a tuple with the EnableCertManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCertManager

`func (o *ProjectActionVisibilityDto) SetEnableCertManager(v ButtonStatusDto)`

SetEnableCertManager sets EnableCertManager field to given value.


### GetDisableCertManager

`func (o *ProjectActionVisibilityDto) GetDisableCertManager() ButtonStatusDto`

GetDisableCertManager returns the DisableCertManager field if non-nil, zero value otherwise.

### GetDisableCertManagerOk

`func (o *ProjectActionVisibilityDto) GetDisableCertManagerOk() (*ButtonStatusDto, bool)`

GetDisableCertManagerOk returns a tuple with the DisableCertManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCertManager

`func (o *ProjectActionVisibilityDto) SetDisableCertManager(v ButtonStatusDto)`

SetDisableCertManager sets DisableCertManager field to given value.


### GetEnableExternalDns

`func (o *ProjectActionVisibilityDto) GetEnableExternalDns() ButtonStatusDto`

GetEnableExternalDns returns the EnableExternalDns field if non-nil, zero value otherwise.

### GetEnableExternalDnsOk

`func (o *ProjectActionVisibilityDto) GetEnableExternalDnsOk() (*ButtonStatusDto, bool)`

GetEnableExternalDnsOk returns a tuple with the EnableExternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalDns

`func (o *ProjectActionVisibilityDto) SetEnableExternalDns(v ButtonStatusDto)`

SetEnableExternalDns sets EnableExternalDns field to given value.


### GetDisableExternalDns

`func (o *ProjectActionVisibilityDto) GetDisableExternalDns() ButtonStatusDto`

GetDisableExternalDns returns the DisableExternalDns field if non-nil, zero value otherwise.

### GetDisableExternalDnsOk

`func (o *ProjectActionVisibilityDto) GetDisableExternalDnsOk() (*ButtonStatusDto, bool)`

GetDisableExternalDnsOk returns a tuple with the DisableExternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableExternalDns

`func (o *ProjectActionVisibilityDto) SetDisableExternalDns(v ButtonStatusDto)`

SetDisableExternalDns sets DisableExternalDns field to given value.


### GetEnableTrustManager

`func (o *ProjectActionVisibilityDto) GetEnableTrustManager() ButtonStatusDto`

GetEnableTrustManager returns the EnableTrustManager field if non-nil, zero value otherwise.

### GetEnableTrustManagerOk

`func (o *ProjectActionVisibilityDto) GetEnableTrustManagerOk() (*ButtonStatusDto, bool)`

GetEnableTrustManagerOk returns a tuple with the EnableTrustManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrustManager

`func (o *ProjectActionVisibilityDto) SetEnableTrustManager(v ButtonStatusDto)`

SetEnableTrustManager sets EnableTrustManager field to given value.


### GetDisableTrustManager

`func (o *ProjectActionVisibilityDto) GetDisableTrustManager() ButtonStatusDto`

GetDisableTrustManager returns the DisableTrustManager field if non-nil, zero value otherwise.

### GetDisableTrustManagerOk

`func (o *ProjectActionVisibilityDto) GetDisableTrustManagerOk() (*ButtonStatusDto, bool)`

GetDisableTrustManagerOk returns a tuple with the DisableTrustManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTrustManager

`func (o *ProjectActionVisibilityDto) SetDisableTrustManager(v ButtonStatusDto)`

SetDisableTrustManager sets DisableTrustManager field to given value.


### GetEnableCTProfiles

`func (o *ProjectActionVisibilityDto) GetEnableCTProfiles() ButtonStatusDto`

GetEnableCTProfiles returns the EnableCTProfiles field if non-nil, zero value otherwise.

### GetEnableCTProfilesOk

`func (o *ProjectActionVisibilityDto) GetEnableCTProfilesOk() (*ButtonStatusDto, bool)`

GetEnableCTProfilesOk returns a tuple with the EnableCTProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCTProfiles

`func (o *ProjectActionVisibilityDto) SetEnableCTProfiles(v ButtonStatusDto)`

SetEnableCTProfiles sets EnableCTProfiles field to given value.


### GetDisableCTProfiles

`func (o *ProjectActionVisibilityDto) GetDisableCTProfiles() ButtonStatusDto`

GetDisableCTProfiles returns the DisableCTProfiles field if non-nil, zero value otherwise.

### GetDisableCTProfilesOk

`func (o *ProjectActionVisibilityDto) GetDisableCTProfilesOk() (*ButtonStatusDto, bool)`

GetDisableCTProfilesOk returns a tuple with the DisableCTProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCTProfiles

`func (o *ProjectActionVisibilityDto) SetDisableCTProfiles(v ButtonStatusDto)`

SetDisableCTProfiles sets DisableCTProfiles field to given value.


### GetLetsEncryptCertificateProfileBlockedReason

`func (o *ProjectActionVisibilityDto) GetLetsEncryptCertificateProfileBlockedReason() string`

GetLetsEncryptCertificateProfileBlockedReason returns the LetsEncryptCertificateProfileBlockedReason field if non-nil, zero value otherwise.

### GetLetsEncryptCertificateProfileBlockedReasonOk

`func (o *ProjectActionVisibilityDto) GetLetsEncryptCertificateProfileBlockedReasonOk() (*string, bool)`

GetLetsEncryptCertificateProfileBlockedReasonOk returns a tuple with the LetsEncryptCertificateProfileBlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLetsEncryptCertificateProfileBlockedReason

`func (o *ProjectActionVisibilityDto) SetLetsEncryptCertificateProfileBlockedReason(v string)`

SetLetsEncryptCertificateProfileBlockedReason sets LetsEncryptCertificateProfileBlockedReason field to given value.

### HasLetsEncryptCertificateProfileBlockedReason

`func (o *ProjectActionVisibilityDto) HasLetsEncryptCertificateProfileBlockedReason() bool`

HasLetsEncryptCertificateProfileBlockedReason returns a boolean if a field has been set.

### SetLetsEncryptCertificateProfileBlockedReasonNil

`func (o *ProjectActionVisibilityDto) SetLetsEncryptCertificateProfileBlockedReasonNil(b bool)`

 SetLetsEncryptCertificateProfileBlockedReasonNil sets the value for LetsEncryptCertificateProfileBlockedReason to be an explicit nil

### UnsetLetsEncryptCertificateProfileBlockedReason
`func (o *ProjectActionVisibilityDto) UnsetLetsEncryptCertificateProfileBlockedReason()`

UnsetLetsEncryptCertificateProfileBlockedReason ensures that no value is present for LetsEncryptCertificateProfileBlockedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


