# CreateTrustedRegistriesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registry** | **NullableString** |  | 
**AccessProfileId** | **int32** |  | 

## Methods

### NewCreateTrustedRegistriesCommand

`func NewCreateTrustedRegistriesCommand(registry NullableString, accessProfileId int32, ) *CreateTrustedRegistriesCommand`

NewCreateTrustedRegistriesCommand instantiates a new CreateTrustedRegistriesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTrustedRegistriesCommandWithDefaults

`func NewCreateTrustedRegistriesCommandWithDefaults() *CreateTrustedRegistriesCommand`

NewCreateTrustedRegistriesCommandWithDefaults instantiates a new CreateTrustedRegistriesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistry

`func (o *CreateTrustedRegistriesCommand) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *CreateTrustedRegistriesCommand) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *CreateTrustedRegistriesCommand) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### SetRegistryNil

`func (o *CreateTrustedRegistriesCommand) SetRegistryNil(b bool)`

 SetRegistryNil sets the value for Registry to be an explicit nil

### UnsetRegistry
`func (o *CreateTrustedRegistriesCommand) UnsetRegistry()`

UnsetRegistry ensures that no value is present for Registry, not even an explicit nil
### GetAccessProfileId

`func (o *CreateTrustedRegistriesCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *CreateTrustedRegistriesCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *CreateTrustedRegistriesCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


