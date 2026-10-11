# NtpServerListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Address** | **NullableString** |  | 

## Methods

### NewNtpServerListDto

`func NewNtpServerListDto(id int32, address NullableString, ) *NtpServerListDto`

NewNtpServerListDto instantiates a new NtpServerListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtpServerListDtoWithDefaults

`func NewNtpServerListDtoWithDefaults() *NtpServerListDto`

NewNtpServerListDtoWithDefaults instantiates a new NtpServerListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NtpServerListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NtpServerListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NtpServerListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetAddress

`func (o *NtpServerListDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NtpServerListDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NtpServerListDto) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *NtpServerListDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NtpServerListDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


