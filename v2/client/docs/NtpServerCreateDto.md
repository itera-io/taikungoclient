# NtpServerCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNtpServerCreateDto

`func NewNtpServerCreateDto() *NtpServerCreateDto`

NewNtpServerCreateDto instantiates a new NtpServerCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpServerCreateDtoWithDefaults

`func NewNtpServerCreateDtoWithDefaults() *NtpServerCreateDto`

NewNtpServerCreateDtoWithDefaults instantiates a new NtpServerCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NtpServerCreateDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NtpServerCreateDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NtpServerCreateDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NtpServerCreateDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *NtpServerCreateDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NtpServerCreateDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


