# CreateNtpServerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **NullableString** | Ntp address example: 8.8.8.8 | 
**AccessProfileId** | **int32** | Id should be positive integer | 

## Methods

### NewCreateNtpServerCommand

`func NewCreateNtpServerCommand(address NullableString, accessProfileId int32, ) *CreateNtpServerCommand`

NewCreateNtpServerCommand instantiates a new CreateNtpServerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNtpServerCommandWithDefaults

`func NewCreateNtpServerCommandWithDefaults() *CreateNtpServerCommand`

NewCreateNtpServerCommandWithDefaults instantiates a new CreateNtpServerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateNtpServerCommand) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateNtpServerCommand) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateNtpServerCommand) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *CreateNtpServerCommand) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CreateNtpServerCommand) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAccessProfileId

`func (o *CreateNtpServerCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *CreateNtpServerCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *CreateNtpServerCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


