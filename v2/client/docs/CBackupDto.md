# CBackupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | **string** |  | 
**CreatedAt** | **NullableTime** |  | 
**Expiration** | **NullableTime** |  | 
**ScheduleName** | **NullableString** |  | 
**Namespace** | **string** |  | 
**Location** | **NullableString** |  | 
**Phase** | **NullableString** |  | 

## Methods

### NewCBackupDto

`func NewCBackupDto(metadataName string, createdAt NullableTime, expiration NullableTime, scheduleName NullableString, namespace string, location NullableString, phase NullableString, ) *CBackupDto`

NewCBackupDto instantiates a new CBackupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCBackupDtoWithDefaults

`func NewCBackupDtoWithDefaults() *CBackupDto`

NewCBackupDtoWithDefaults instantiates a new CBackupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *CBackupDto) GetMetadataName() string`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *CBackupDto) GetMetadataNameOk() (*string, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *CBackupDto) SetMetadataName(v string)`

SetMetadataName sets MetadataName field to given value.


### GetCreatedAt

`func (o *CBackupDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CBackupDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CBackupDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CBackupDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CBackupDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiration

`func (o *CBackupDto) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CBackupDto) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CBackupDto) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### SetExpirationNil

`func (o *CBackupDto) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *CBackupDto) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetScheduleName

`func (o *CBackupDto) GetScheduleName() string`

GetScheduleName returns the ScheduleName field if non-nil, zero value otherwise.

### GetScheduleNameOk

`func (o *CBackupDto) GetScheduleNameOk() (*string, bool)`

GetScheduleNameOk returns a tuple with the ScheduleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleName

`func (o *CBackupDto) SetScheduleName(v string)`

SetScheduleName sets ScheduleName field to given value.


### SetScheduleNameNil

`func (o *CBackupDto) SetScheduleNameNil(b bool)`

 SetScheduleNameNil sets the value for ScheduleName to be an explicit nil

### UnsetScheduleName
`func (o *CBackupDto) UnsetScheduleName()`

UnsetScheduleName ensures that no value is present for ScheduleName, not even an explicit nil
### GetNamespace

`func (o *CBackupDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CBackupDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CBackupDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetLocation

`func (o *CBackupDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CBackupDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CBackupDto) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *CBackupDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *CBackupDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetPhase

`func (o *CBackupDto) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *CBackupDto) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *CBackupDto) SetPhase(v string)`

SetPhase sets Phase field to given value.


### SetPhaseNil

`func (o *CBackupDto) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *CBackupDto) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


