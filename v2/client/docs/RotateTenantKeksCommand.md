# RotateTenantKeksCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchSize** | Pointer to **int32** |  | [optional] [default to 200]
**Mode** | Pointer to [**KeyRotationMode**](KeyRotationMode.md) |  | [optional] 

## Methods

### NewRotateTenantKeksCommand

`func NewRotateTenantKeksCommand() *RotateTenantKeksCommand`

NewRotateTenantKeksCommand instantiates a new RotateTenantKeksCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateTenantKeksCommandWithDefaults

`func NewRotateTenantKeksCommandWithDefaults() *RotateTenantKeksCommand`

NewRotateTenantKeksCommandWithDefaults instantiates a new RotateTenantKeksCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchSize

`func (o *RotateTenantKeksCommand) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *RotateTenantKeksCommand) GetBatchSizeOk() (*int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *RotateTenantKeksCommand) SetBatchSize(v int32)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *RotateTenantKeksCommand) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetMode

`func (o *RotateTenantKeksCommand) GetMode() KeyRotationMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RotateTenantKeksCommand) GetModeOk() (*KeyRotationMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RotateTenantKeksCommand) SetMode(v KeyRotationMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RotateTenantKeksCommand) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


