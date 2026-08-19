# CScheduleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**MetadataName** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 
**Schedule** | **NullableString** |  | 
**Ttl** | **NullableString** |  | 
**LastBackup** | **NullableTime** |  | 
**Phase** | **NullableString** |  | 
**ExcludedNamespaces** | **[]string** |  | 
**IncludedNamespaces** | **[]string** |  | 

## Methods

### NewCScheduleDto

`func NewCScheduleDto(metadataName NullableString, namespace NullableString, createdAt NullableTime, schedule NullableString, ttl NullableString, lastBackup NullableTime, phase NullableString, excludedNamespaces []string, includedNamespaces []string, ) *CScheduleDto`

NewCScheduleDto instantiates a new CScheduleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCScheduleDtoWithDefaults

`func NewCScheduleDtoWithDefaults() *CScheduleDto`

NewCScheduleDtoWithDefaults instantiates a new CScheduleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CScheduleDto) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CScheduleDto) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CScheduleDto) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CScheduleDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadataName

`func (o *CScheduleDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CScheduleDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CScheduleDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.


### SetMetadataNameNil

`func (o *CScheduleDto) SetMetadataNameNil(b bool)`

 SetMetadataNameNil sets the value for MetadataName to be an explicit nil

### UnsetMetadataName
`func (o *CScheduleDto) UnsetMetadataName()`

UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
### GetNamespace

`func (o *CScheduleDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CScheduleDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CScheduleDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *CScheduleDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *CScheduleDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetCreatedAt

`func (o *CScheduleDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CScheduleDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CScheduleDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CScheduleDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CScheduleDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetSchedule

`func (o *CScheduleDto) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CScheduleDto) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CScheduleDto) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### SetScheduleNil

`func (o *CScheduleDto) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CScheduleDto) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTtl

`func (o *CScheduleDto) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CScheduleDto) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CScheduleDto) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### SetTtlNil

`func (o *CScheduleDto) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *CScheduleDto) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetLastBackup

`func (o *CScheduleDto) GetLastBackup() time.Time`

GetLastBackup returns the LastBackup field if non-nil, zero value otherwise.

### GetLastBackupOk

`func (o *CScheduleDto) GetLastBackupOk() (*time.Time, bool)`

GetLastBackupOk returns a tuple with the LastBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackup

`func (o *CScheduleDto) SetLastBackup(v time.Time)`

SetLastBackup sets LastBackup field to given value.


### SetLastBackupNil

`func (o *CScheduleDto) SetLastBackupNil(b bool)`

 SetLastBackupNil sets the value for LastBackup to be an explicit nil

### UnsetLastBackup
`func (o *CScheduleDto) UnsetLastBackup()`

UnsetLastBackup ensures that no value is present for LastBackup, not even an explicit nil
### GetPhase

`func (o *CScheduleDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CScheduleDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CScheduleDto) SetPhase(v string)`

SetPhase sets Phase field to given value.


### SetPhaseNil

`func (o *CScheduleDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *CScheduleDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetExcludedNamespaces

`func (o *CScheduleDto) GetExcludedNamespaces() []string`

GetExcludedNamespaces returns the ExcludedNamespaces field if non-nil, zero value otherwise.

### GetExcludedNamespacesOk

`func (o *CScheduleDto) GetExcludedNamespacesOk() (*[]string, bool)`

GetExcludedNamespacesOk returns a tuple with the ExcludedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedNamespaces

`func (o *CScheduleDto) SetExcludedNamespaces(v []string)`

SetExcludedNamespaces sets ExcludedNamespaces field to given value.


### SetExcludedNamespacesNil

`func (o *CScheduleDto) SetExcludedNamespacesNil(b bool)`

 SetExcludedNamespacesNil sets the value for ExcludedNamespaces to be an explicit nil

### UnsetExcludedNamespaces
`func (o *CScheduleDto) UnsetExcludedNamespaces()`

UnsetExcludedNamespaces ensures that no value is present for ExcludedNamespaces, not even an explicit nil
### GetIncludedNamespaces

`func (o *CScheduleDto) GetIncludedNamespaces() []string`

GetIncludedNamespaces returns the IncludedNamespaces field if non-nil, zero value otherwise.

### GetIncludedNamespacesOk

`func (o *CScheduleDto) GetIncludedNamespacesOk() (*[]string, bool)`

GetIncludedNamespacesOk returns a tuple with the IncludedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedNamespaces

`func (o *CScheduleDto) SetIncludedNamespaces(v []string)`

SetIncludedNamespaces sets IncludedNamespaces field to given value.


### SetIncludedNamespacesNil

`func (o *CScheduleDto) SetIncludedNamespacesNil(b bool)`

 SetIncludedNamespacesNil sets the value for IncludedNamespaces to be an explicit nil

### UnsetIncludedNamespaces
`func (o *CScheduleDto) UnsetIncludedNamespaces()`

UnsetIncludedNamespaces ensures that no value is present for IncludedNamespaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


