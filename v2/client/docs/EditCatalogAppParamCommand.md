# EditCatalogAppParamCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogAppId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to [**[]CatalogAppParamsDto**](CatalogAppParamsDto.md) |  | [optional] 

## Methods

### NewEditCatalogAppParamCommand

`func NewEditCatalogAppParamCommand() *EditCatalogAppParamCommand`

NewEditCatalogAppParamCommand instantiates a new EditCatalogAppParamCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditCatalogAppParamCommandWithDefaults

`func NewEditCatalogAppParamCommandWithDefaults() *EditCatalogAppParamCommand`

NewEditCatalogAppParamCommandWithDefaults instantiates a new EditCatalogAppParamCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogAppId

`func (o *EditCatalogAppParamCommand) GetCatalogAppId() int32`

GetCatalogAppId returns the CatalogAppId field if non-nil, zero value otherwise.

### GetCatalogAppIdOk

`func (o *EditCatalogAppParamCommand) GetCatalogAppIdOk() (*int32, bool)`

GetCatalogAppIdOk returns a tuple with the CatalogAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppId

`func (o *EditCatalogAppParamCommand) SetCatalogAppId(v int32)`

SetCatalogAppId sets CatalogAppId field to given value.

### HasCatalogAppId

`func (o *EditCatalogAppParamCommand) HasCatalogAppId() bool`

HasCatalogAppId returns a boolean if a field has been set.

### GetParameters

`func (o *EditCatalogAppParamCommand) GetParameters() []CatalogAppParamsDto`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EditCatalogAppParamCommand) GetParametersOk() (*[]CatalogAppParamsDto, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EditCatalogAppParamCommand) SetParameters(v []CatalogAppParamsDto)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EditCatalogAppParamCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *EditCatalogAppParamCommand) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *EditCatalogAppParamCommand) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


