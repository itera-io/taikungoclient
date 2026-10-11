# ProjectDetailsForServersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertsCount** | **int32** |  | 
**Worker** | **int32** |  | 
**Bastion** | **int32** |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Name** | **string** |  | 
**AccessIp** | **string** |  | 
**Id** | **int32** |  | 
**MasterReady** | **int32** |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**CloudName** | **string** |  | 
**CloudId** | **int32** |  | 
**QuotaId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**OrganizationId** | **int32** |  | 
**KubernetesVersion** | **string** |  | 
**IsBackupEnabled** | **bool** |  | 
**AiEnabled** | **bool** |  | 
**IsLocked** | **bool** |  | 
**IsAutoUpgrade** | **bool** |  | 
**IsMonitoringEnabled** | **bool** |  | 
**IsOpaEnabled** | **bool** |  | 
**HasKubeConfigFile** | **bool** |  | 
**HasSelectedFlavors** | **bool** |  | 
**IsMaintenanceModeEnabled** | **bool** |  | 
**IsProjectMaintenanceModeEnabled** | **bool** |  | 
**IsDeprecated** | **bool** |  | 
**CpuLimit** | **int64** |  | 
**RamLimit** | **int64** |  | 
**DiskSizeLimit** | **int64** |  | 
**UsedCpu** | **int64** |  | 
**UsedRam** | **int64** |  | 
**UsedDiskSize** | **int64** |  | 
**VmCpuLimit** | **int64** |  | 
**VmRamLimit** | **int64** |  | 
**VmVolumeSizeLimit** | **int64** |  | 
**VmUsedCpu** | **int64** |  | 
**VmUsedRam** | **int64** |  | 
**VmUsedVolumeSize** | **int64** |  | 
**AccessProfileName** | **string** |  | 
**AccessProfileId** | **NullableInt32** |  | 
**KubernetesProfileName** | **string** |  | 
**KubernetesProfileId** | **NullableInt32** |  | 
**AlertingProfileName** | **string** |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**S3CredentialId** | **NullableInt32** |  | 
**AiCredentialId** | **NullableInt32** |  | 
**DnsCredentialId** | **NullableInt32** |  | 
**DnsCredentialName** | **string** |  | 
**CtProfileId** | Pointer to **NullableInt32** |  | [optional] 
**CtProfileName** | Pointer to **NullableString** |  | [optional] 
**ExpiredAt** | **string** |  | 
**CertificationExpiredAt** | **string** |  | 
**OpaProfileId** | **NullableInt32** |  | 
**OpaProfileName** | **string** |  | 
**AllowFullSpotKubernetes** | **bool** |  | 
**AllowSpotWorkers** | **bool** |  | 
**AllowSpotVMs** | **bool** |  | 
**AutoscalingGroupName** | **string** |  | 
**MinSize** | **NullableInt32** |  | 
**MaxSize** | **NullableInt32** |  | 
**DiskSize** | **NullableFloat64** |  | 
**Flavor** | **NullableString** |  | 
**SpotEnabled** | **NullableBool** |  | 
**IsAutoscalingEnabled** | **bool** |  | 
**IsAutoscalingSpotEnabled** | **bool** |  | 
**GatewayApiModern** | **bool** |  | 
**HasNfsServer** | **bool** |  | 
**WasmEnabled** | **bool** |  | 
**AvailabilityZones** | **[]string** |  | 
**Hypervisors** | **[]string** |  | 
**MaxSpotPrice** | **NullableFloat64** |  | 
**VpcMode** | Pointer to **map[string]interface{}** |  | [optional] 
**CloudSubnets** | Pointer to [**[]CloudSubnetDetailsDto**](CloudSubnetDetailsDto.md) |  | [optional] 
**PrivateOnly** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewProjectDetailsForServersDto

`func NewProjectDetailsForServersDto(alertsCount int32, worker int32, bastion int32, status ProjectStatus, name string, accessIp string, id int32, masterReady int32, cloudType ECloudCredentialType, cloudName string, cloudId int32, quotaId int32, organizationName string, organizationId int32, kubernetesVersion string, isBackupEnabled bool, aiEnabled bool, isLocked bool, isAutoUpgrade bool, isMonitoringEnabled bool, isOpaEnabled bool, hasKubeConfigFile bool, hasSelectedFlavors bool, isMaintenanceModeEnabled bool, isProjectMaintenanceModeEnabled bool, isDeprecated bool, cpuLimit int64, ramLimit int64, diskSizeLimit int64, usedCpu int64, usedRam int64, usedDiskSize int64, vmCpuLimit int64, vmRamLimit int64, vmVolumeSizeLimit int64, vmUsedCpu int64, vmUsedRam int64, vmUsedVolumeSize int64, accessProfileName string, accessProfileId NullableInt32, kubernetesProfileName string, kubernetesProfileId NullableInt32, alertingProfileName string, health ProjectHealth, alertingProfileId NullableInt32, s3CredentialId NullableInt32, aiCredentialId NullableInt32, dnsCredentialId NullableInt32, dnsCredentialName string, expiredAt string, certificationExpiredAt string, opaProfileId NullableInt32, opaProfileName string, allowFullSpotKubernetes bool, allowSpotWorkers bool, allowSpotVMs bool, autoscalingGroupName string, minSize NullableInt32, maxSize NullableInt32, diskSize NullableFloat64, flavor NullableString, spotEnabled NullableBool, isAutoscalingEnabled bool, isAutoscalingSpotEnabled bool, gatewayApiModern bool, hasNfsServer bool, wasmEnabled bool, availabilityZones []string, hypervisors []string, maxSpotPrice NullableFloat64, ) *ProjectDetailsForServersDto`

NewProjectDetailsForServersDto instantiates a new ProjectDetailsForServersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsForServersDtoWithDefaults

`func NewProjectDetailsForServersDtoWithDefaults() *ProjectDetailsForServersDto`

NewProjectDetailsForServersDtoWithDefaults instantiates a new ProjectDetailsForServersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertsCount

`func (o *ProjectDetailsForServersDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *ProjectDetailsForServersDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *ProjectDetailsForServersDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetWorker

`func (o *ProjectDetailsForServersDto) GetWorker() int32`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ProjectDetailsForServersDto) GetWorkerOk() (*int32, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ProjectDetailsForServersDto) SetWorker(v int32)`

SetWorker sets Worker field to given value.


### GetBastion

`func (o *ProjectDetailsForServersDto) GetBastion() int32`

GetBastion returns the Bastion field if non-nil, zero value otherwise.

### GetBastionOk

`func (o *ProjectDetailsForServersDto) GetBastionOk() (*int32, bool)`

GetBastionOk returns a tuple with the Bastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBastion

`func (o *ProjectDetailsForServersDto) SetBastion(v int32)`

SetBastion sets Bastion field to given value.


### GetStatus

`func (o *ProjectDetailsForServersDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectDetailsForServersDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectDetailsForServersDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetName

`func (o *ProjectDetailsForServersDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetailsForServersDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetailsForServersDto) SetName(v string)`

SetName sets Name field to given value.


### GetAccessIp

`func (o *ProjectDetailsForServersDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectDetailsForServersDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectDetailsForServersDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.


### GetId

`func (o *ProjectDetailsForServersDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDetailsForServersDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDetailsForServersDto) SetId(v int32)`

SetId sets Id field to given value.


### GetMasterReady

`func (o *ProjectDetailsForServersDto) GetMasterReady() int32`

GetMasterReady returns the MasterReady field if non-nil, zero value otherwise.

### GetMasterReadyOk

`func (o *ProjectDetailsForServersDto) GetMasterReadyOk() (*int32, bool)`

GetMasterReadyOk returns a tuple with the MasterReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterReady

`func (o *ProjectDetailsForServersDto) SetMasterReady(v int32)`

SetMasterReady sets MasterReady field to given value.


### GetCloudType

`func (o *ProjectDetailsForServersDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectDetailsForServersDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectDetailsForServersDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetCloudName

`func (o *ProjectDetailsForServersDto) GetCloudName() string`

GetCloudName returns the CloudName field if non-nil, zero value otherwise.

### GetCloudNameOk

`func (o *ProjectDetailsForServersDto) GetCloudNameOk() (*string, bool)`

GetCloudNameOk returns a tuple with the CloudName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudName

`func (o *ProjectDetailsForServersDto) SetCloudName(v string)`

SetCloudName sets CloudName field to given value.


### GetCloudId

`func (o *ProjectDetailsForServersDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ProjectDetailsForServersDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ProjectDetailsForServersDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### GetQuotaId

`func (o *ProjectDetailsForServersDto) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *ProjectDetailsForServersDto) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *ProjectDetailsForServersDto) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.


### GetOrganizationName

`func (o *ProjectDetailsForServersDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectDetailsForServersDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectDetailsForServersDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *ProjectDetailsForServersDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectDetailsForServersDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectDetailsForServersDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetKubernetesVersion

`func (o *ProjectDetailsForServersDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ProjectDetailsForServersDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ProjectDetailsForServersDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetIsBackupEnabled

`func (o *ProjectDetailsForServersDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectDetailsForServersDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.


### GetAiEnabled

`func (o *ProjectDetailsForServersDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *ProjectDetailsForServersDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *ProjectDetailsForServersDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.


### GetIsLocked

`func (o *ProjectDetailsForServersDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectDetailsForServersDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectDetailsForServersDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsAutoUpgrade

`func (o *ProjectDetailsForServersDto) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *ProjectDetailsForServersDto) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *ProjectDetailsForServersDto) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.


### GetIsMonitoringEnabled

`func (o *ProjectDetailsForServersDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectDetailsForServersDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.


### GetIsOpaEnabled

`func (o *ProjectDetailsForServersDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ProjectDetailsForServersDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.


### GetHasKubeConfigFile

`func (o *ProjectDetailsForServersDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectDetailsForServersDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectDetailsForServersDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.


### GetHasSelectedFlavors

`func (o *ProjectDetailsForServersDto) GetHasSelectedFlavors() bool`

GetHasSelectedFlavors returns the HasSelectedFlavors field if non-nil, zero value otherwise.

### GetHasSelectedFlavorsOk

`func (o *ProjectDetailsForServersDto) GetHasSelectedFlavorsOk() (*bool, bool)`

GetHasSelectedFlavorsOk returns a tuple with the HasSelectedFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSelectedFlavors

`func (o *ProjectDetailsForServersDto) SetHasSelectedFlavors(v bool)`

SetHasSelectedFlavors sets HasSelectedFlavors field to given value.


### GetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.


### GetIsProjectMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) GetIsProjectMaintenanceModeEnabled() bool`

GetIsProjectMaintenanceModeEnabled returns the IsProjectMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsProjectMaintenanceModeEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsProjectMaintenanceModeEnabledOk() (*bool, bool)`

GetIsProjectMaintenanceModeEnabledOk returns a tuple with the IsProjectMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectMaintenanceModeEnabled

`func (o *ProjectDetailsForServersDto) SetIsProjectMaintenanceModeEnabled(v bool)`

SetIsProjectMaintenanceModeEnabled sets IsProjectMaintenanceModeEnabled field to given value.


### GetIsDeprecated

`func (o *ProjectDetailsForServersDto) GetIsDeprecated() bool`

GetIsDeprecated returns the IsDeprecated field if non-nil, zero value otherwise.

### GetIsDeprecatedOk

`func (o *ProjectDetailsForServersDto) GetIsDeprecatedOk() (*bool, bool)`

GetIsDeprecatedOk returns a tuple with the IsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeprecated

`func (o *ProjectDetailsForServersDto) SetIsDeprecated(v bool)`

SetIsDeprecated sets IsDeprecated field to given value.


### GetCpuLimit

`func (o *ProjectDetailsForServersDto) GetCpuLimit() int64`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ProjectDetailsForServersDto) GetCpuLimitOk() (*int64, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ProjectDetailsForServersDto) SetCpuLimit(v int64)`

SetCpuLimit sets CpuLimit field to given value.


### GetRamLimit

`func (o *ProjectDetailsForServersDto) GetRamLimit() int64`

GetRamLimit returns the RamLimit field if non-nil, zero value otherwise.

### GetRamLimitOk

`func (o *ProjectDetailsForServersDto) GetRamLimitOk() (*int64, bool)`

GetRamLimitOk returns a tuple with the RamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamLimit

`func (o *ProjectDetailsForServersDto) SetRamLimit(v int64)`

SetRamLimit sets RamLimit field to given value.


### GetDiskSizeLimit

`func (o *ProjectDetailsForServersDto) GetDiskSizeLimit() int64`

GetDiskSizeLimit returns the DiskSizeLimit field if non-nil, zero value otherwise.

### GetDiskSizeLimitOk

`func (o *ProjectDetailsForServersDto) GetDiskSizeLimitOk() (*int64, bool)`

GetDiskSizeLimitOk returns a tuple with the DiskSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeLimit

`func (o *ProjectDetailsForServersDto) SetDiskSizeLimit(v int64)`

SetDiskSizeLimit sets DiskSizeLimit field to given value.


### GetUsedCpu

`func (o *ProjectDetailsForServersDto) GetUsedCpu() int64`

GetUsedCpu returns the UsedCpu field if non-nil, zero value otherwise.

### GetUsedCpuOk

`func (o *ProjectDetailsForServersDto) GetUsedCpuOk() (*int64, bool)`

GetUsedCpuOk returns a tuple with the UsedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCpu

`func (o *ProjectDetailsForServersDto) SetUsedCpu(v int64)`

SetUsedCpu sets UsedCpu field to given value.


### GetUsedRam

`func (o *ProjectDetailsForServersDto) GetUsedRam() int64`

GetUsedRam returns the UsedRam field if non-nil, zero value otherwise.

### GetUsedRamOk

`func (o *ProjectDetailsForServersDto) GetUsedRamOk() (*int64, bool)`

GetUsedRamOk returns a tuple with the UsedRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRam

`func (o *ProjectDetailsForServersDto) SetUsedRam(v int64)`

SetUsedRam sets UsedRam field to given value.


### GetUsedDiskSize

`func (o *ProjectDetailsForServersDto) GetUsedDiskSize() int64`

GetUsedDiskSize returns the UsedDiskSize field if non-nil, zero value otherwise.

### GetUsedDiskSizeOk

`func (o *ProjectDetailsForServersDto) GetUsedDiskSizeOk() (*int64, bool)`

GetUsedDiskSizeOk returns a tuple with the UsedDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDiskSize

`func (o *ProjectDetailsForServersDto) SetUsedDiskSize(v int64)`

SetUsedDiskSize sets UsedDiskSize field to given value.


### GetVmCpuLimit

`func (o *ProjectDetailsForServersDto) GetVmCpuLimit() int64`

GetVmCpuLimit returns the VmCpuLimit field if non-nil, zero value otherwise.

### GetVmCpuLimitOk

`func (o *ProjectDetailsForServersDto) GetVmCpuLimitOk() (*int64, bool)`

GetVmCpuLimitOk returns a tuple with the VmCpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCpuLimit

`func (o *ProjectDetailsForServersDto) SetVmCpuLimit(v int64)`

SetVmCpuLimit sets VmCpuLimit field to given value.


### GetVmRamLimit

`func (o *ProjectDetailsForServersDto) GetVmRamLimit() int64`

GetVmRamLimit returns the VmRamLimit field if non-nil, zero value otherwise.

### GetVmRamLimitOk

`func (o *ProjectDetailsForServersDto) GetVmRamLimitOk() (*int64, bool)`

GetVmRamLimitOk returns a tuple with the VmRamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRamLimit

`func (o *ProjectDetailsForServersDto) SetVmRamLimit(v int64)`

SetVmRamLimit sets VmRamLimit field to given value.


### GetVmVolumeSizeLimit

`func (o *ProjectDetailsForServersDto) GetVmVolumeSizeLimit() int64`

GetVmVolumeSizeLimit returns the VmVolumeSizeLimit field if non-nil, zero value otherwise.

### GetVmVolumeSizeLimitOk

`func (o *ProjectDetailsForServersDto) GetVmVolumeSizeLimitOk() (*int64, bool)`

GetVmVolumeSizeLimitOk returns a tuple with the VmVolumeSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVolumeSizeLimit

`func (o *ProjectDetailsForServersDto) SetVmVolumeSizeLimit(v int64)`

SetVmVolumeSizeLimit sets VmVolumeSizeLimit field to given value.


### GetVmUsedCpu

`func (o *ProjectDetailsForServersDto) GetVmUsedCpu() int64`

GetVmUsedCpu returns the VmUsedCpu field if non-nil, zero value otherwise.

### GetVmUsedCpuOk

`func (o *ProjectDetailsForServersDto) GetVmUsedCpuOk() (*int64, bool)`

GetVmUsedCpuOk returns a tuple with the VmUsedCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsedCpu

`func (o *ProjectDetailsForServersDto) SetVmUsedCpu(v int64)`

SetVmUsedCpu sets VmUsedCpu field to given value.


### GetVmUsedRam

`func (o *ProjectDetailsForServersDto) GetVmUsedRam() int64`

GetVmUsedRam returns the VmUsedRam field if non-nil, zero value otherwise.

### GetVmUsedRamOk

`func (o *ProjectDetailsForServersDto) GetVmUsedRamOk() (*int64, bool)`

GetVmUsedRamOk returns a tuple with the VmUsedRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsedRam

`func (o *ProjectDetailsForServersDto) SetVmUsedRam(v int64)`

SetVmUsedRam sets VmUsedRam field to given value.


### GetVmUsedVolumeSize

`func (o *ProjectDetailsForServersDto) GetVmUsedVolumeSize() int64`

GetVmUsedVolumeSize returns the VmUsedVolumeSize field if non-nil, zero value otherwise.

### GetVmUsedVolumeSizeOk

`func (o *ProjectDetailsForServersDto) GetVmUsedVolumeSizeOk() (*int64, bool)`

GetVmUsedVolumeSizeOk returns a tuple with the VmUsedVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmUsedVolumeSize

`func (o *ProjectDetailsForServersDto) SetVmUsedVolumeSize(v int64)`

SetVmUsedVolumeSize sets VmUsedVolumeSize field to given value.


### GetAccessProfileName

`func (o *ProjectDetailsForServersDto) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *ProjectDetailsForServersDto) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *ProjectDetailsForServersDto) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.


### GetAccessProfileId

`func (o *ProjectDetailsForServersDto) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *ProjectDetailsForServersDto) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *ProjectDetailsForServersDto) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.


### SetAccessProfileIdNil

`func (o *ProjectDetailsForServersDto) SetAccessProfileIdNil(b bool)`

 SetAccessProfileIdNil sets the value for AccessProfileId to be an explicit nil

### UnsetAccessProfileId
`func (o *ProjectDetailsForServersDto) UnsetAccessProfileId()`

UnsetAccessProfileId ensures that no value is present for AccessProfileId, not even an explicit nil
### GetKubernetesProfileName

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileName() string`

GetKubernetesProfileName returns the KubernetesProfileName field if non-nil, zero value otherwise.

### GetKubernetesProfileNameOk

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileNameOk() (*string, bool)`

GetKubernetesProfileNameOk returns a tuple with the KubernetesProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfileName

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileName(v string)`

SetKubernetesProfileName sets KubernetesProfileName field to given value.


### GetKubernetesProfileId

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileId() int32`

GetKubernetesProfileId returns the KubernetesProfileId field if non-nil, zero value otherwise.

### GetKubernetesProfileIdOk

`func (o *ProjectDetailsForServersDto) GetKubernetesProfileIdOk() (*int32, bool)`

GetKubernetesProfileIdOk returns a tuple with the KubernetesProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfileId

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileId(v int32)`

SetKubernetesProfileId sets KubernetesProfileId field to given value.


### SetKubernetesProfileIdNil

`func (o *ProjectDetailsForServersDto) SetKubernetesProfileIdNil(b bool)`

 SetKubernetesProfileIdNil sets the value for KubernetesProfileId to be an explicit nil

### UnsetKubernetesProfileId
`func (o *ProjectDetailsForServersDto) UnsetKubernetesProfileId()`

UnsetKubernetesProfileId ensures that no value is present for KubernetesProfileId, not even an explicit nil
### GetAlertingProfileName

`func (o *ProjectDetailsForServersDto) GetAlertingProfileName() string`

GetAlertingProfileName returns the AlertingProfileName field if non-nil, zero value otherwise.

### GetAlertingProfileNameOk

`func (o *ProjectDetailsForServersDto) GetAlertingProfileNameOk() (*string, bool)`

GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileName

`func (o *ProjectDetailsForServersDto) SetAlertingProfileName(v string)`

SetAlertingProfileName sets AlertingProfileName field to given value.


### GetHealth

`func (o *ProjectDetailsForServersDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectDetailsForServersDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectDetailsForServersDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetAlertingProfileId

`func (o *ProjectDetailsForServersDto) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *ProjectDetailsForServersDto) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *ProjectDetailsForServersDto) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *ProjectDetailsForServersDto) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *ProjectDetailsForServersDto) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetS3CredentialId

`func (o *ProjectDetailsForServersDto) GetS3CredentialId() int32`

GetS3CredentialId returns the S3CredentialId field if non-nil, zero value otherwise.

### GetS3CredentialIdOk

`func (o *ProjectDetailsForServersDto) GetS3CredentialIdOk() (*int32, bool)`

GetS3CredentialIdOk returns a tuple with the S3CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CredentialId

`func (o *ProjectDetailsForServersDto) SetS3CredentialId(v int32)`

SetS3CredentialId sets S3CredentialId field to given value.


### SetS3CredentialIdNil

`func (o *ProjectDetailsForServersDto) SetS3CredentialIdNil(b bool)`

 SetS3CredentialIdNil sets the value for S3CredentialId to be an explicit nil

### UnsetS3CredentialId
`func (o *ProjectDetailsForServersDto) UnsetS3CredentialId()`

UnsetS3CredentialId ensures that no value is present for S3CredentialId, not even an explicit nil
### GetAiCredentialId

`func (o *ProjectDetailsForServersDto) GetAiCredentialId() int32`

GetAiCredentialId returns the AiCredentialId field if non-nil, zero value otherwise.

### GetAiCredentialIdOk

`func (o *ProjectDetailsForServersDto) GetAiCredentialIdOk() (*int32, bool)`

GetAiCredentialIdOk returns a tuple with the AiCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiCredentialId

`func (o *ProjectDetailsForServersDto) SetAiCredentialId(v int32)`

SetAiCredentialId sets AiCredentialId field to given value.


### SetAiCredentialIdNil

`func (o *ProjectDetailsForServersDto) SetAiCredentialIdNil(b bool)`

 SetAiCredentialIdNil sets the value for AiCredentialId to be an explicit nil

### UnsetAiCredentialId
`func (o *ProjectDetailsForServersDto) UnsetAiCredentialId()`

UnsetAiCredentialId ensures that no value is present for AiCredentialId, not even an explicit nil
### GetDnsCredentialId

`func (o *ProjectDetailsForServersDto) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *ProjectDetailsForServersDto) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *ProjectDetailsForServersDto) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.


### SetDnsCredentialIdNil

`func (o *ProjectDetailsForServersDto) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *ProjectDetailsForServersDto) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetDnsCredentialName

`func (o *ProjectDetailsForServersDto) GetDnsCredentialName() string`

GetDnsCredentialName returns the DnsCredentialName field if non-nil, zero value otherwise.

### GetDnsCredentialNameOk

`func (o *ProjectDetailsForServersDto) GetDnsCredentialNameOk() (*string, bool)`

GetDnsCredentialNameOk returns a tuple with the DnsCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialName

`func (o *ProjectDetailsForServersDto) SetDnsCredentialName(v string)`

SetDnsCredentialName sets DnsCredentialName field to given value.


### GetCtProfileId

`func (o *ProjectDetailsForServersDto) GetCtProfileId() int32`

GetCtProfileId returns the CtProfileId field if non-nil, zero value otherwise.

### GetCtProfileIdOk

`func (o *ProjectDetailsForServersDto) GetCtProfileIdOk() (*int32, bool)`

GetCtProfileIdOk returns a tuple with the CtProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtProfileId

`func (o *ProjectDetailsForServersDto) SetCtProfileId(v int32)`

SetCtProfileId sets CtProfileId field to given value.

### HasCtProfileId

`func (o *ProjectDetailsForServersDto) HasCtProfileId() bool`

HasCtProfileId returns a boolean if a field has been set.

### SetCtProfileIdNil

`func (o *ProjectDetailsForServersDto) SetCtProfileIdNil(b bool)`

 SetCtProfileIdNil sets the value for CtProfileId to be an explicit nil

### UnsetCtProfileId
`func (o *ProjectDetailsForServersDto) UnsetCtProfileId()`

UnsetCtProfileId ensures that no value is present for CtProfileId, not even an explicit nil
### GetCtProfileName

`func (o *ProjectDetailsForServersDto) GetCtProfileName() string`

GetCtProfileName returns the CtProfileName field if non-nil, zero value otherwise.

### GetCtProfileNameOk

`func (o *ProjectDetailsForServersDto) GetCtProfileNameOk() (*string, bool)`

GetCtProfileNameOk returns a tuple with the CtProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtProfileName

`func (o *ProjectDetailsForServersDto) SetCtProfileName(v string)`

SetCtProfileName sets CtProfileName field to given value.

### HasCtProfileName

`func (o *ProjectDetailsForServersDto) HasCtProfileName() bool`

HasCtProfileName returns a boolean if a field has been set.

### SetCtProfileNameNil

`func (o *ProjectDetailsForServersDto) SetCtProfileNameNil(b bool)`

 SetCtProfileNameNil sets the value for CtProfileName to be an explicit nil

### UnsetCtProfileName
`func (o *ProjectDetailsForServersDto) UnsetCtProfileName()`

UnsetCtProfileName ensures that no value is present for CtProfileName, not even an explicit nil
### GetExpiredAt

`func (o *ProjectDetailsForServersDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectDetailsForServersDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectDetailsForServersDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### GetCertificationExpiredAt

`func (o *ProjectDetailsForServersDto) GetCertificationExpiredAt() string`

GetCertificationExpiredAt returns the CertificationExpiredAt field if non-nil, zero value otherwise.

### GetCertificationExpiredAtOk

`func (o *ProjectDetailsForServersDto) GetCertificationExpiredAtOk() (*string, bool)`

GetCertificationExpiredAtOk returns a tuple with the CertificationExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationExpiredAt

`func (o *ProjectDetailsForServersDto) SetCertificationExpiredAt(v string)`

SetCertificationExpiredAt sets CertificationExpiredAt field to given value.


### GetOpaProfileId

`func (o *ProjectDetailsForServersDto) GetOpaProfileId() int32`

GetOpaProfileId returns the OpaProfileId field if non-nil, zero value otherwise.

### GetOpaProfileIdOk

`func (o *ProjectDetailsForServersDto) GetOpaProfileIdOk() (*int32, bool)`

GetOpaProfileIdOk returns a tuple with the OpaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileId

`func (o *ProjectDetailsForServersDto) SetOpaProfileId(v int32)`

SetOpaProfileId sets OpaProfileId field to given value.


### SetOpaProfileIdNil

`func (o *ProjectDetailsForServersDto) SetOpaProfileIdNil(b bool)`

 SetOpaProfileIdNil sets the value for OpaProfileId to be an explicit nil

### UnsetOpaProfileId
`func (o *ProjectDetailsForServersDto) UnsetOpaProfileId()`

UnsetOpaProfileId ensures that no value is present for OpaProfileId, not even an explicit nil
### GetOpaProfileName

`func (o *ProjectDetailsForServersDto) GetOpaProfileName() string`

GetOpaProfileName returns the OpaProfileName field if non-nil, zero value otherwise.

### GetOpaProfileNameOk

`func (o *ProjectDetailsForServersDto) GetOpaProfileNameOk() (*string, bool)`

GetOpaProfileNameOk returns a tuple with the OpaProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileName

`func (o *ProjectDetailsForServersDto) SetOpaProfileName(v string)`

SetOpaProfileName sets OpaProfileName field to given value.


### GetAllowFullSpotKubernetes

`func (o *ProjectDetailsForServersDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectDetailsForServersDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectDetailsForServersDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.


### GetAllowSpotWorkers

`func (o *ProjectDetailsForServersDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectDetailsForServersDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectDetailsForServersDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.


### GetAllowSpotVMs

`func (o *ProjectDetailsForServersDto) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *ProjectDetailsForServersDto) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *ProjectDetailsForServersDto) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.


### GetAutoscalingGroupName

`func (o *ProjectDetailsForServersDto) GetAutoscalingGroupName() string`

GetAutoscalingGroupName returns the AutoscalingGroupName field if non-nil, zero value otherwise.

### GetAutoscalingGroupNameOk

`func (o *ProjectDetailsForServersDto) GetAutoscalingGroupNameOk() (*string, bool)`

GetAutoscalingGroupNameOk returns a tuple with the AutoscalingGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingGroupName

`func (o *ProjectDetailsForServersDto) SetAutoscalingGroupName(v string)`

SetAutoscalingGroupName sets AutoscalingGroupName field to given value.


### GetMinSize

`func (o *ProjectDetailsForServersDto) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *ProjectDetailsForServersDto) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *ProjectDetailsForServersDto) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.


### SetMinSizeNil

`func (o *ProjectDetailsForServersDto) SetMinSizeNil(b bool)`

 SetMinSizeNil sets the value for MinSize to be an explicit nil

### UnsetMinSize
`func (o *ProjectDetailsForServersDto) UnsetMinSize()`

UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
### GetMaxSize

`func (o *ProjectDetailsForServersDto) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *ProjectDetailsForServersDto) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *ProjectDetailsForServersDto) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.


### SetMaxSizeNil

`func (o *ProjectDetailsForServersDto) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *ProjectDetailsForServersDto) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetDiskSize

`func (o *ProjectDetailsForServersDto) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ProjectDetailsForServersDto) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ProjectDetailsForServersDto) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.


### SetDiskSizeNil

`func (o *ProjectDetailsForServersDto) SetDiskSizeNil(b bool)`

 SetDiskSizeNil sets the value for DiskSize to be an explicit nil

### UnsetDiskSize
`func (o *ProjectDetailsForServersDto) UnsetDiskSize()`

UnsetDiskSize ensures that no value is present for DiskSize, not even an explicit nil
### GetFlavor

`func (o *ProjectDetailsForServersDto) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ProjectDetailsForServersDto) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ProjectDetailsForServersDto) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### SetFlavorNil

`func (o *ProjectDetailsForServersDto) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *ProjectDetailsForServersDto) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
### GetSpotEnabled

`func (o *ProjectDetailsForServersDto) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *ProjectDetailsForServersDto) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *ProjectDetailsForServersDto) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.


### SetSpotEnabledNil

`func (o *ProjectDetailsForServersDto) SetSpotEnabledNil(b bool)`

 SetSpotEnabledNil sets the value for SpotEnabled to be an explicit nil

### UnsetSpotEnabled
`func (o *ProjectDetailsForServersDto) UnsetSpotEnabled()`

UnsetSpotEnabled ensures that no value is present for SpotEnabled, not even an explicit nil
### GetIsAutoscalingEnabled

`func (o *ProjectDetailsForServersDto) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ProjectDetailsForServersDto) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.


### GetIsAutoscalingSpotEnabled

`func (o *ProjectDetailsForServersDto) GetIsAutoscalingSpotEnabled() bool`

GetIsAutoscalingSpotEnabled returns the IsAutoscalingSpotEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingSpotEnabledOk

`func (o *ProjectDetailsForServersDto) GetIsAutoscalingSpotEnabledOk() (*bool, bool)`

GetIsAutoscalingSpotEnabledOk returns a tuple with the IsAutoscalingSpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingSpotEnabled

`func (o *ProjectDetailsForServersDto) SetIsAutoscalingSpotEnabled(v bool)`

SetIsAutoscalingSpotEnabled sets IsAutoscalingSpotEnabled field to given value.


### GetGatewayApiModern

`func (o *ProjectDetailsForServersDto) GetGatewayApiModern() bool`

GetGatewayApiModern returns the GatewayApiModern field if non-nil, zero value otherwise.

### GetGatewayApiModernOk

`func (o *ProjectDetailsForServersDto) GetGatewayApiModernOk() (*bool, bool)`

GetGatewayApiModernOk returns a tuple with the GatewayApiModern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayApiModern

`func (o *ProjectDetailsForServersDto) SetGatewayApiModern(v bool)`

SetGatewayApiModern sets GatewayApiModern field to given value.


### GetHasNfsServer

`func (o *ProjectDetailsForServersDto) GetHasNfsServer() bool`

GetHasNfsServer returns the HasNfsServer field if non-nil, zero value otherwise.

### GetHasNfsServerOk

`func (o *ProjectDetailsForServersDto) GetHasNfsServerOk() (*bool, bool)`

GetHasNfsServerOk returns a tuple with the HasNfsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNfsServer

`func (o *ProjectDetailsForServersDto) SetHasNfsServer(v bool)`

SetHasNfsServer sets HasNfsServer field to given value.


### GetWasmEnabled

`func (o *ProjectDetailsForServersDto) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *ProjectDetailsForServersDto) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *ProjectDetailsForServersDto) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.


### GetAvailabilityZones

`func (o *ProjectDetailsForServersDto) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *ProjectDetailsForServersDto) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *ProjectDetailsForServersDto) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.


### GetHypervisors

`func (o *ProjectDetailsForServersDto) GetHypervisors() []string`

GetHypervisors returns the Hypervisors field if non-nil, zero value otherwise.

### GetHypervisorsOk

`func (o *ProjectDetailsForServersDto) GetHypervisorsOk() (*[]string, bool)`

GetHypervisorsOk returns a tuple with the Hypervisors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisors

`func (o *ProjectDetailsForServersDto) SetHypervisors(v []string)`

SetHypervisors sets Hypervisors field to given value.


### GetMaxSpotPrice

`func (o *ProjectDetailsForServersDto) GetMaxSpotPrice() float64`

GetMaxSpotPrice returns the MaxSpotPrice field if non-nil, zero value otherwise.

### GetMaxSpotPriceOk

`func (o *ProjectDetailsForServersDto) GetMaxSpotPriceOk() (*float64, bool)`

GetMaxSpotPriceOk returns a tuple with the MaxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotPrice

`func (o *ProjectDetailsForServersDto) SetMaxSpotPrice(v float64)`

SetMaxSpotPrice sets MaxSpotPrice field to given value.


### SetMaxSpotPriceNil

`func (o *ProjectDetailsForServersDto) SetMaxSpotPriceNil(b bool)`

 SetMaxSpotPriceNil sets the value for MaxSpotPrice to be an explicit nil

### UnsetMaxSpotPrice
`func (o *ProjectDetailsForServersDto) UnsetMaxSpotPrice()`

UnsetMaxSpotPrice ensures that no value is present for MaxSpotPrice, not even an explicit nil
### GetVpcMode

`func (o *ProjectDetailsForServersDto) GetVpcMode() map[string]interface{}`

GetVpcMode returns the VpcMode field if non-nil, zero value otherwise.

### GetVpcModeOk

`func (o *ProjectDetailsForServersDto) GetVpcModeOk() (*map[string]interface{}, bool)`

GetVpcModeOk returns a tuple with the VpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMode

`func (o *ProjectDetailsForServersDto) SetVpcMode(v map[string]interface{})`

SetVpcMode sets VpcMode field to given value.

### HasVpcMode

`func (o *ProjectDetailsForServersDto) HasVpcMode() bool`

HasVpcMode returns a boolean if a field has been set.

### GetCloudSubnets

`func (o *ProjectDetailsForServersDto) GetCloudSubnets() []CloudSubnetDetailsDto`

GetCloudSubnets returns the CloudSubnets field if non-nil, zero value otherwise.

### GetCloudSubnetsOk

`func (o *ProjectDetailsForServersDto) GetCloudSubnetsOk() (*[]CloudSubnetDetailsDto, bool)`

GetCloudSubnetsOk returns a tuple with the CloudSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSubnets

`func (o *ProjectDetailsForServersDto) SetCloudSubnets(v []CloudSubnetDetailsDto)`

SetCloudSubnets sets CloudSubnets field to given value.

### HasCloudSubnets

`func (o *ProjectDetailsForServersDto) HasCloudSubnets() bool`

HasCloudSubnets returns a boolean if a field has been set.

### SetCloudSubnetsNil

`func (o *ProjectDetailsForServersDto) SetCloudSubnetsNil(b bool)`

 SetCloudSubnetsNil sets the value for CloudSubnets to be an explicit nil

### UnsetCloudSubnets
`func (o *ProjectDetailsForServersDto) UnsetCloudSubnets()`

UnsetCloudSubnets ensures that no value is present for CloudSubnets, not even an explicit nil
### GetPrivateOnly

`func (o *ProjectDetailsForServersDto) GetPrivateOnly() bool`

GetPrivateOnly returns the PrivateOnly field if non-nil, zero value otherwise.

### GetPrivateOnlyOk

`func (o *ProjectDetailsForServersDto) GetPrivateOnlyOk() (*bool, bool)`

GetPrivateOnlyOk returns a tuple with the PrivateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOnly

`func (o *ProjectDetailsForServersDto) SetPrivateOnly(v bool)`

SetPrivateOnly sets PrivateOnly field to given value.

### HasPrivateOnly

`func (o *ProjectDetailsForServersDto) HasPrivateOnly() bool`

HasPrivateOnly returns a boolean if a field has been set.

### SetPrivateOnlyNil

`func (o *ProjectDetailsForServersDto) SetPrivateOnlyNil(b bool)`

 SetPrivateOnlyNil sets the value for PrivateOnly to be an explicit nil

### UnsetPrivateOnly
`func (o *ProjectDetailsForServersDto) UnsetPrivateOnly()`

UnsetPrivateOnly ensures that no value is present for PrivateOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


