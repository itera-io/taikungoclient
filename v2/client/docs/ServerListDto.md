# ServerListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ProjectName** | **string** |  | 
**OrganizationName** | **string** |  | 
**OrganizationId** | **int32** |  | 
**ProjectId** | **int32** |  | 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**DiskSize** | **float64** |  | 
**KubernetesHealth** | **NullableString** |  | 
**Cpu** | **int32** |  | 
**Ram** | **float64** |  | 
**Role** | [**CloudRole**](CloudRole.md) |  | 
**Status** | **string** |  | 
**CreatedAt** | **NullableString** |  | 
**CloudType** | **map[string]interface{}** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**SpotPrice** | **float64** |  | 
**SpotInstance** | **bool** |  | 
**ShutOff** | **bool** |  | 
**UseLocalDisk** | Pointer to **bool** |  | [optional] 
**AutoscalingGroup** | **NullableString** |  | 
**ProviderID** | **NullableString** |  | 
**InstanceId** | **NullableString** |  | 
**AwsHostName** | **NullableString** |  | 
**AvailabilityZone** | **NullableString** |  | 
**Hypervisor** | **NullableString** |  | 
**HypervisorId** | **NullableString** |  | 
**ActionButtons** | Pointer to [**ServerActionButtonVisibilityDto**](ServerActionButtonVisibilityDto.md) |  | [optional] 
**KubernetesNodeLabels** | [**[]KubernetesNodeLabelsDto**](KubernetesNodeLabelsDto.md) |  | 
**ReplicaCount** | **NullableInt32** |  | 
**WasmEnabled** | **bool** |  | 
**Flavor** | **NullableString** |  | 

## Methods

### NewServerListDto

`func NewServerListDto(id int32, name string, projectName string, organizationName string, organizationId int32, projectId int32, diskSize float64, kubernetesHealth NullableString, cpu int32, ram float64, role CloudRole, status string, createdAt NullableString, cloudType map[string]interface{}, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, spotPrice float64, spotInstance bool, shutOff bool, autoscalingGroup NullableString, providerID NullableString, instanceId NullableString, awsHostName NullableString, availabilityZone NullableString, hypervisor NullableString, hypervisorId NullableString, kubernetesNodeLabels []KubernetesNodeLabelsDto, replicaCount NullableInt32, wasmEnabled bool, flavor NullableString, ) *ServerListDto`

NewServerListDto instantiates a new ServerListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerListDtoWithDefaults

`func NewServerListDtoWithDefaults() *ServerListDto`

NewServerListDtoWithDefaults instantiates a new ServerListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ServerListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerListDto) SetName(v string)`

SetName sets Name field to given value.


### GetProjectName

`func (o *ServerListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ServerListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ServerListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetOrganizationName

`func (o *ServerListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ServerListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ServerListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *ServerListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServerListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServerListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetProjectId

`func (o *ServerListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ServerListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ServerListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetIpAddress

`func (o *ServerListDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ServerListDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ServerListDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ServerListDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ServerListDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ServerListDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetDiskSize

`func (o *ServerListDto) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *ServerListDto) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *ServerListDto) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.


### GetKubernetesHealth

`func (o *ServerListDto) GetKubernetesHealth() string`

GetKubernetesHealth returns the KubernetesHealth field if non-nil, zero value otherwise.

### GetKubernetesHealthOk

`func (o *ServerListDto) GetKubernetesHealthOk() (*string, bool)`

GetKubernetesHealthOk returns a tuple with the KubernetesHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesHealth

`func (o *ServerListDto) SetKubernetesHealth(v string)`

SetKubernetesHealth sets KubernetesHealth field to given value.


### SetKubernetesHealthNil

`func (o *ServerListDto) SetKubernetesHealthNil(b bool)`

 SetKubernetesHealthNil sets the value for KubernetesHealth to be an explicit nil

### UnsetKubernetesHealth
`func (o *ServerListDto) UnsetKubernetesHealth()`

UnsetKubernetesHealth ensures that no value is present for KubernetesHealth, not even an explicit nil
### GetCpu

`func (o *ServerListDto) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServerListDto) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServerListDto) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetRam

`func (o *ServerListDto) GetRam() float64`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerListDto) GetRamOk() (*float64, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerListDto) SetRam(v float64)`

SetRam sets Ram field to given value.


### GetRole

`func (o *ServerListDto) GetRole() CloudRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ServerListDto) GetRoleOk() (*CloudRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ServerListDto) SetRole(v CloudRole)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ServerListDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerListDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerListDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ServerListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ServerListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServerListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCloudType

`func (o *ServerListDto) GetCloudType() map[string]interface{}`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ServerListDto) GetCloudTypeOk() (*map[string]interface{}, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ServerListDto) SetCloudType(v map[string]interface{})`

SetCloudType sets CloudType field to given value.


### GetCreatedBy

`func (o *ServerListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServerListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServerListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *ServerListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ServerListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ServerListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *ServerListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ServerListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *ServerListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ServerListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ServerListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetSpotPrice

`func (o *ServerListDto) GetSpotPrice() float64`

GetSpotPrice returns the SpotPrice field if non-nil, zero value otherwise.

### GetSpotPriceOk

`func (o *ServerListDto) GetSpotPriceOk() (*float64, bool)`

GetSpotPriceOk returns a tuple with the SpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPrice

`func (o *ServerListDto) SetSpotPrice(v float64)`

SetSpotPrice sets SpotPrice field to given value.


### GetSpotInstance

`func (o *ServerListDto) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *ServerListDto) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *ServerListDto) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.


### GetShutOff

`func (o *ServerListDto) GetShutOff() bool`

GetShutOff returns the ShutOff field if non-nil, zero value otherwise.

### GetShutOffOk

`func (o *ServerListDto) GetShutOffOk() (*bool, bool)`

GetShutOffOk returns a tuple with the ShutOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutOff

`func (o *ServerListDto) SetShutOff(v bool)`

SetShutOff sets ShutOff field to given value.


### GetUseLocalDisk

`func (o *ServerListDto) GetUseLocalDisk() bool`

GetUseLocalDisk returns the UseLocalDisk field if non-nil, zero value otherwise.

### GetUseLocalDiskOk

`func (o *ServerListDto) GetUseLocalDiskOk() (*bool, bool)`

GetUseLocalDiskOk returns a tuple with the UseLocalDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalDisk

`func (o *ServerListDto) SetUseLocalDisk(v bool)`

SetUseLocalDisk sets UseLocalDisk field to given value.

### HasUseLocalDisk

`func (o *ServerListDto) HasUseLocalDisk() bool`

HasUseLocalDisk returns a boolean if a field has been set.

### GetAutoscalingGroup

`func (o *ServerListDto) GetAutoscalingGroup() string`

GetAutoscalingGroup returns the AutoscalingGroup field if non-nil, zero value otherwise.

### GetAutoscalingGroupOk

`func (o *ServerListDto) GetAutoscalingGroupOk() (*string, bool)`

GetAutoscalingGroupOk returns a tuple with the AutoscalingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingGroup

`func (o *ServerListDto) SetAutoscalingGroup(v string)`

SetAutoscalingGroup sets AutoscalingGroup field to given value.


### SetAutoscalingGroupNil

`func (o *ServerListDto) SetAutoscalingGroupNil(b bool)`

 SetAutoscalingGroupNil sets the value for AutoscalingGroup to be an explicit nil

### UnsetAutoscalingGroup
`func (o *ServerListDto) UnsetAutoscalingGroup()`

UnsetAutoscalingGroup ensures that no value is present for AutoscalingGroup, not even an explicit nil
### GetProviderID

`func (o *ServerListDto) GetProviderID() string`

GetProviderID returns the ProviderID field if non-nil, zero value otherwise.

### GetProviderIDOk

`func (o *ServerListDto) GetProviderIDOk() (*string, bool)`

GetProviderIDOk returns a tuple with the ProviderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderID

`func (o *ServerListDto) SetProviderID(v string)`

SetProviderID sets ProviderID field to given value.


### SetProviderIDNil

`func (o *ServerListDto) SetProviderIDNil(b bool)`

 SetProviderIDNil sets the value for ProviderID to be an explicit nil

### UnsetProviderID
`func (o *ServerListDto) UnsetProviderID()`

UnsetProviderID ensures that no value is present for ProviderID, not even an explicit nil
### GetInstanceId

`func (o *ServerListDto) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ServerListDto) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ServerListDto) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### SetInstanceIdNil

`func (o *ServerListDto) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ServerListDto) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetAwsHostName

`func (o *ServerListDto) GetAwsHostName() string`

GetAwsHostName returns the AwsHostName field if non-nil, zero value otherwise.

### GetAwsHostNameOk

`func (o *ServerListDto) GetAwsHostNameOk() (*string, bool)`

GetAwsHostNameOk returns a tuple with the AwsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsHostName

`func (o *ServerListDto) SetAwsHostName(v string)`

SetAwsHostName sets AwsHostName field to given value.


### SetAwsHostNameNil

`func (o *ServerListDto) SetAwsHostNameNil(b bool)`

 SetAwsHostNameNil sets the value for AwsHostName to be an explicit nil

### UnsetAwsHostName
`func (o *ServerListDto) UnsetAwsHostName()`

UnsetAwsHostName ensures that no value is present for AwsHostName, not even an explicit nil
### GetAvailabilityZone

`func (o *ServerListDto) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ServerListDto) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ServerListDto) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### SetAvailabilityZoneNil

`func (o *ServerListDto) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *ServerListDto) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetHypervisor

`func (o *ServerListDto) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *ServerListDto) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *ServerListDto) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.


### SetHypervisorNil

`func (o *ServerListDto) SetHypervisorNil(b bool)`

 SetHypervisorNil sets the value for Hypervisor to be an explicit nil

### UnsetHypervisor
`func (o *ServerListDto) UnsetHypervisor()`

UnsetHypervisor ensures that no value is present for Hypervisor, not even an explicit nil
### GetHypervisorId

`func (o *ServerListDto) GetHypervisorId() string`

GetHypervisorId returns the HypervisorId field if non-nil, zero value otherwise.

### GetHypervisorIdOk

`func (o *ServerListDto) GetHypervisorIdOk() (*string, bool)`

GetHypervisorIdOk returns a tuple with the HypervisorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorId

`func (o *ServerListDto) SetHypervisorId(v string)`

SetHypervisorId sets HypervisorId field to given value.


### SetHypervisorIdNil

`func (o *ServerListDto) SetHypervisorIdNil(b bool)`

 SetHypervisorIdNil sets the value for HypervisorId to be an explicit nil

### UnsetHypervisorId
`func (o *ServerListDto) UnsetHypervisorId()`

UnsetHypervisorId ensures that no value is present for HypervisorId, not even an explicit nil
### GetActionButtons

`func (o *ServerListDto) GetActionButtons() ServerActionButtonVisibilityDto`

GetActionButtons returns the ActionButtons field if non-nil, zero value otherwise.

### GetActionButtonsOk

`func (o *ServerListDto) GetActionButtonsOk() (*ServerActionButtonVisibilityDto, bool)`

GetActionButtonsOk returns a tuple with the ActionButtons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionButtons

`func (o *ServerListDto) SetActionButtons(v ServerActionButtonVisibilityDto)`

SetActionButtons sets ActionButtons field to given value.

### HasActionButtons

`func (o *ServerListDto) HasActionButtons() bool`

HasActionButtons returns a boolean if a field has been set.

### GetKubernetesNodeLabels

`func (o *ServerListDto) GetKubernetesNodeLabels() []KubernetesNodeLabelsDto`

GetKubernetesNodeLabels returns the KubernetesNodeLabels field if non-nil, zero value otherwise.

### GetKubernetesNodeLabelsOk

`func (o *ServerListDto) GetKubernetesNodeLabelsOk() (*[]KubernetesNodeLabelsDto, bool)`

GetKubernetesNodeLabelsOk returns a tuple with the KubernetesNodeLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesNodeLabels

`func (o *ServerListDto) SetKubernetesNodeLabels(v []KubernetesNodeLabelsDto)`

SetKubernetesNodeLabels sets KubernetesNodeLabels field to given value.


### GetReplicaCount

`func (o *ServerListDto) GetReplicaCount() int32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ServerListDto) GetReplicaCountOk() (*int32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ServerListDto) SetReplicaCount(v int32)`

SetReplicaCount sets ReplicaCount field to given value.


### SetReplicaCountNil

`func (o *ServerListDto) SetReplicaCountNil(b bool)`

 SetReplicaCountNil sets the value for ReplicaCount to be an explicit nil

### UnsetReplicaCount
`func (o *ServerListDto) UnsetReplicaCount()`

UnsetReplicaCount ensures that no value is present for ReplicaCount, not even an explicit nil
### GetWasmEnabled

`func (o *ServerListDto) GetWasmEnabled() bool`

GetWasmEnabled returns the WasmEnabled field if non-nil, zero value otherwise.

### GetWasmEnabledOk

`func (o *ServerListDto) GetWasmEnabledOk() (*bool, bool)`

GetWasmEnabledOk returns a tuple with the WasmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasmEnabled

`func (o *ServerListDto) SetWasmEnabled(v bool)`

SetWasmEnabled sets WasmEnabled field to given value.


### GetFlavor

`func (o *ServerListDto) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *ServerListDto) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *ServerListDto) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.


### SetFlavorNil

`func (o *ServerListDto) SetFlavorNil(b bool)`

 SetFlavorNil sets the value for Flavor to be an explicit nil

### UnsetFlavor
`func (o *ServerListDto) UnsetFlavor()`

UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


