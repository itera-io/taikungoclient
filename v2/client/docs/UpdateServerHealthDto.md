# UpdateServerHealthDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**ServerHealth** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateServerHealthDto

`func NewUpdateServerHealthDto() *UpdateServerHealthDto`

NewUpdateServerHealthDto instantiates a new UpdateServerHealthDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerHealthDtoWithDefaults

`func NewUpdateServerHealthDtoWithDefaults() *UpdateServerHealthDto`

NewUpdateServerHealthDtoWithDefaults instantiates a new UpdateServerHealthDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *UpdateServerHealthDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateServerHealthDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateServerHealthDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateServerHealthDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *UpdateServerHealthDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *UpdateServerHealthDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetServerHealth

`func (o *UpdateServerHealthDto) GetServerHealth() string`

GetServerHealth returns the ServerHealth field if non-nil, zero value otherwise.

### GetServerHealthOk

`func (o *UpdateServerHealthDto) GetServerHealthOk() (*string, bool)`

GetServerHealthOk returns a tuple with the ServerHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHealth

`func (o *UpdateServerHealthDto) SetServerHealth(v string)`

SetServerHealth sets ServerHealth field to given value.

### HasServerHealth

`func (o *UpdateServerHealthDto) HasServerHealth() bool`

HasServerHealth returns a boolean if a field has been set.

### SetServerHealthNil

`func (o *UpdateServerHealthDto) SetServerHealthNil(b bool)`

 SetServerHealthNil sets the value for ServerHealth to be an explicit nil

### UnsetServerHealth
`func (o *UpdateServerHealthDto) UnsetServerHealth()`

UnsetServerHealth ensures that no value is present for ServerHealth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


