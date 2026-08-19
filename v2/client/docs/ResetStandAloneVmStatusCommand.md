# ResetStandAloneVmStatusCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**VmIds** | Pointer to **[]int32** |  | [optional] 
**Status** | Pointer to [**StandAloneVmStatus**](StandAloneVmStatus.md) |  | [optional] 

## Methods

### NewResetStandAloneVmStatusCommand

`func NewResetStandAloneVmStatusCommand() *ResetStandAloneVmStatusCommand`

NewResetStandAloneVmStatusCommand instantiates a new ResetStandAloneVmStatusCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetStandAloneVmStatusCommandWithDefaults

`func NewResetStandAloneVmStatusCommandWithDefaults() *ResetStandAloneVmStatusCommand`

NewResetStandAloneVmStatusCommandWithDefaults instantiates a new ResetStandAloneVmStatusCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ResetStandAloneVmStatusCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ResetStandAloneVmStatusCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ResetStandAloneVmStatusCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ResetStandAloneVmStatusCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetVmIds

`func (o *ResetStandAloneVmStatusCommand) GetVmIds() []int32`

GetVmIds returns the VmIds field if non-nil, zero value otherwise.

### GetVmIdsOk

`func (o *ResetStandAloneVmStatusCommand) GetVmIdsOk() (*[]int32, bool)`

GetVmIdsOk returns a tuple with the VmIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIds

`func (o *ResetStandAloneVmStatusCommand) SetVmIds(v []int32)`

SetVmIds sets VmIds field to given value.

### HasVmIds

`func (o *ResetStandAloneVmStatusCommand) HasVmIds() bool`

HasVmIds returns a boolean if a field has been set.

### SetVmIdsNil

`func (o *ResetStandAloneVmStatusCommand) SetVmIdsNil(b bool)`

 SetVmIdsNil sets the value for VmIds to be an explicit nil

### UnsetVmIds
`func (o *ResetStandAloneVmStatusCommand) UnsetVmIds()`

UnsetVmIds ensures that no value is present for VmIds, not even an explicit nil
### GetStatus

`func (o *ResetStandAloneVmStatusCommand) GetStatus() StandAloneVmStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResetStandAloneVmStatusCommand) GetStatusOk() (*StandAloneVmStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResetStandAloneVmStatusCommand) SetStatus(v StandAloneVmStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResetStandAloneVmStatusCommand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


