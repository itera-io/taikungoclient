# DeleteVmDiskCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandaloneVmId** | **int32** |  | 
**VmDiskIds** | **[]int32** |  | 

## Methods

### NewDeleteVmDiskCommand

`func NewDeleteVmDiskCommand(standaloneVmId int32, vmDiskIds []int32, ) *DeleteVmDiskCommand`

NewDeleteVmDiskCommand instantiates a new DeleteVmDiskCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteVmDiskCommandWithDefaults

`func NewDeleteVmDiskCommandWithDefaults() *DeleteVmDiskCommand`

NewDeleteVmDiskCommandWithDefaults instantiates a new DeleteVmDiskCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandaloneVmId

`func (o *DeleteVmDiskCommand) GetStandaloneVmId() int32`

GetStandaloneVmId returns the StandaloneVmId field if non-nil, zero value otherwise.

### GetStandaloneVmIdOk

`func (o *DeleteVmDiskCommand) GetStandaloneVmIdOk() (*int32, bool)`

GetStandaloneVmIdOk returns a tuple with the StandaloneVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVmId

`func (o *DeleteVmDiskCommand) SetStandaloneVmId(v int32)`

SetStandaloneVmId sets StandaloneVmId field to given value.


### GetVmDiskIds

`func (o *DeleteVmDiskCommand) GetVmDiskIds() []int32`

GetVmDiskIds returns the VmDiskIds field if non-nil, zero value otherwise.

### GetVmDiskIdsOk

`func (o *DeleteVmDiskCommand) GetVmDiskIdsOk() (*[]int32, bool)`

GetVmDiskIdsOk returns a tuple with the VmDiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmDiskIds

`func (o *DeleteVmDiskCommand) SetVmDiskIds(v []int32)`

SetVmDiskIds sets VmDiskIds field to given value.


### SetVmDiskIdsNil

`func (o *DeleteVmDiskCommand) SetVmDiskIdsNil(b bool)`

 SetVmDiskIdsNil sets the value for VmDiskIds to be an explicit nil

### UnsetVmDiskIds
`func (o *DeleteVmDiskCommand) UnsetVmDiskIds()`

UnsetVmDiskIds ensures that no value is present for VmDiskIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


