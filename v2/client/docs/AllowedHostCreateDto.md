# AllowedHostCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**MaskBits** | Pointer to **int32** |  | [optional] 

## Methods

### NewAllowedHostCreateDto

`func NewAllowedHostCreateDto() *AllowedHostCreateDto`

NewAllowedHostCreateDto instantiates a new AllowedHostCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedHostCreateDtoWithDefaults

`func NewAllowedHostCreateDtoWithDefaults() *AllowedHostCreateDto`

NewAllowedHostCreateDtoWithDefaults instantiates a new AllowedHostCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AllowedHostCreateDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowedHostCreateDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowedHostCreateDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllowedHostCreateDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AllowedHostCreateDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AllowedHostCreateDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *AllowedHostCreateDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AllowedHostCreateDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AllowedHostCreateDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AllowedHostCreateDto) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *AllowedHostCreateDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *AllowedHostCreateDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMaskBits

`func (o *AllowedHostCreateDto) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *AllowedHostCreateDto) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *AllowedHostCreateDto) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.

### HasMaskBits

`func (o *AllowedHostCreateDto) HasMaskBits() bool`

HasMaskBits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


