# UpdateOrganizationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**FullName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateOrganizationCommand

`func NewUpdateOrganizationCommand() *UpdateOrganizationCommand`

NewUpdateOrganizationCommand instantiates a new UpdateOrganizationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationCommandWithDefaults

`func NewUpdateOrganizationCommandWithDefaults() *UpdateOrganizationCommand`

NewUpdateOrganizationCommandWithDefaults instantiates a new UpdateOrganizationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOrganizationCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrganizationCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrganizationCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOrganizationCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateOrganizationCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateOrganizationCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateOrganizationCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFullName

`func (o *UpdateOrganizationCommand) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UpdateOrganizationCommand) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UpdateOrganizationCommand) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UpdateOrganizationCommand) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *UpdateOrganizationCommand) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *UpdateOrganizationCommand) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


