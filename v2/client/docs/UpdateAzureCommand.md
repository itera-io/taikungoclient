# UpdateAzureCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**AzureClientSecret** | **NullableString** |  | 
**AzureClientId** | **NullableString** |  | 

## Methods

### NewUpdateAzureCommand

`func NewUpdateAzureCommand(id int32, name NullableString, azureClientSecret NullableString, azureClientId NullableString, ) *UpdateAzureCommand`

NewUpdateAzureCommand instantiates a new UpdateAzureCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAzureCommandWithDefaults

`func NewUpdateAzureCommandWithDefaults() *UpdateAzureCommand`

NewUpdateAzureCommandWithDefaults instantiates a new UpdateAzureCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAzureCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAzureCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAzureCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAzureCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAzureCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAzureCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateAzureCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAzureCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAzureClientSecret

`func (o *UpdateAzureCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *UpdateAzureCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *UpdateAzureCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### SetAzureClientSecretNil

`func (o *UpdateAzureCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *UpdateAzureCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetAzureClientId

`func (o *UpdateAzureCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *UpdateAzureCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *UpdateAzureCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### SetAzureClientIdNil

`func (o *UpdateAzureCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *UpdateAzureCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


