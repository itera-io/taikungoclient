# CreateVirtualClusterCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | **NullableInt32** |  | 
**ProjectId** | **int32** |  | 
**Name** | **string** |  | 
**ExpiredAt** | **NullableTime** |  | 
**DeleteOnExpiration** | **bool** |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**ExposeHostname** | **string** |  | 
**WorkloadResources** | [**CreateVirtualClusterWorkloadResources**](CreateVirtualClusterWorkloadResources.md) |  | 
**CpuLimits** | [**CreateVirtualClusterResourceLimits**](CreateVirtualClusterResourceLimits.md) |  | 
**RamLimits** | [**CreateVirtualClusterResourceLimits**](CreateVirtualClusterResourceLimits.md) |  | 
**EphemeralStorageLimits** | [**CreateVirtualClusterResourceLimits**](CreateVirtualClusterResourceLimits.md) |  | 

## Methods

### NewCreateVirtualClusterCommand

`func NewCreateVirtualClusterCommand(cloudId NullableInt32, projectId int32, name string, expiredAt NullableTime, deleteOnExpiration bool, alertingProfileId NullableInt32, exposeHostname string, workloadResources CreateVirtualClusterWorkloadResources, cpuLimits CreateVirtualClusterResourceLimits, ramLimits CreateVirtualClusterResourceLimits, ephemeralStorageLimits CreateVirtualClusterResourceLimits, ) *CreateVirtualClusterCommand`

NewCreateVirtualClusterCommand instantiates a new CreateVirtualClusterCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualClusterCommandWithDefaults

`func NewCreateVirtualClusterCommandWithDefaults() *CreateVirtualClusterCommand`

NewCreateVirtualClusterCommandWithDefaults instantiates a new CreateVirtualClusterCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *CreateVirtualClusterCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *CreateVirtualClusterCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *CreateVirtualClusterCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *CreateVirtualClusterCommand) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *CreateVirtualClusterCommand) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
### GetProjectId

`func (o *CreateVirtualClusterCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateVirtualClusterCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateVirtualClusterCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *CreateVirtualClusterCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVirtualClusterCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVirtualClusterCommand) SetName(v string)`

SetName sets Name field to given value.


### GetExpiredAt

`func (o *CreateVirtualClusterCommand) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CreateVirtualClusterCommand) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CreateVirtualClusterCommand) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *CreateVirtualClusterCommand) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *CreateVirtualClusterCommand) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetDeleteOnExpiration

`func (o *CreateVirtualClusterCommand) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *CreateVirtualClusterCommand) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *CreateVirtualClusterCommand) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.


### GetAlertingProfileId

`func (o *CreateVirtualClusterCommand) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *CreateVirtualClusterCommand) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *CreateVirtualClusterCommand) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *CreateVirtualClusterCommand) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *CreateVirtualClusterCommand) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetExposeHostname

`func (o *CreateVirtualClusterCommand) GetExposeHostname() string`

GetExposeHostname returns the ExposeHostname field if non-nil, zero value otherwise.

### GetExposeHostnameOk

`func (o *CreateVirtualClusterCommand) GetExposeHostnameOk() (*string, bool)`

GetExposeHostnameOk returns a tuple with the ExposeHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeHostname

`func (o *CreateVirtualClusterCommand) SetExposeHostname(v string)`

SetExposeHostname sets ExposeHostname field to given value.


### GetWorkloadResources

`func (o *CreateVirtualClusterCommand) GetWorkloadResources() CreateVirtualClusterWorkloadResources`

GetWorkloadResources returns the WorkloadResources field if non-nil, zero value otherwise.

### GetWorkloadResourcesOk

`func (o *CreateVirtualClusterCommand) GetWorkloadResourcesOk() (*CreateVirtualClusterWorkloadResources, bool)`

GetWorkloadResourcesOk returns a tuple with the WorkloadResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadResources

`func (o *CreateVirtualClusterCommand) SetWorkloadResources(v CreateVirtualClusterWorkloadResources)`

SetWorkloadResources sets WorkloadResources field to given value.


### GetCpuLimits

`func (o *CreateVirtualClusterCommand) GetCpuLimits() CreateVirtualClusterResourceLimits`

GetCpuLimits returns the CpuLimits field if non-nil, zero value otherwise.

### GetCpuLimitsOk

`func (o *CreateVirtualClusterCommand) GetCpuLimitsOk() (*CreateVirtualClusterResourceLimits, bool)`

GetCpuLimitsOk returns a tuple with the CpuLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimits

`func (o *CreateVirtualClusterCommand) SetCpuLimits(v CreateVirtualClusterResourceLimits)`

SetCpuLimits sets CpuLimits field to given value.


### GetRamLimits

`func (o *CreateVirtualClusterCommand) GetRamLimits() CreateVirtualClusterResourceLimits`

GetRamLimits returns the RamLimits field if non-nil, zero value otherwise.

### GetRamLimitsOk

`func (o *CreateVirtualClusterCommand) GetRamLimitsOk() (*CreateVirtualClusterResourceLimits, bool)`

GetRamLimitsOk returns a tuple with the RamLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamLimits

`func (o *CreateVirtualClusterCommand) SetRamLimits(v CreateVirtualClusterResourceLimits)`

SetRamLimits sets RamLimits field to given value.


### GetEphemeralStorageLimits

`func (o *CreateVirtualClusterCommand) GetEphemeralStorageLimits() CreateVirtualClusterResourceLimits`

GetEphemeralStorageLimits returns the EphemeralStorageLimits field if non-nil, zero value otherwise.

### GetEphemeralStorageLimitsOk

`func (o *CreateVirtualClusterCommand) GetEphemeralStorageLimitsOk() (*CreateVirtualClusterResourceLimits, bool)`

GetEphemeralStorageLimitsOk returns a tuple with the EphemeralStorageLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorageLimits

`func (o *CreateVirtualClusterCommand) SetEphemeralStorageLimits(v CreateVirtualClusterResourceLimits)`

SetEphemeralStorageLimits sets EphemeralStorageLimits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


