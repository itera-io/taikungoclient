# EditAllowedHostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**IpAddress** | **NullableString** |  | 
**MaskBits** | **int32** |  | 

## Methods

### NewEditAllowedHostDto

`func NewEditAllowedHostDto(description NullableString, ipAddress NullableString, maskBits int32, ) *EditAllowedHostDto`

NewEditAllowedHostDto instantiates a new EditAllowedHostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAllowedHostDtoWithDefaults

`func NewEditAllowedHostDtoWithDefaults() *EditAllowedHostDto`

NewEditAllowedHostDtoWithDefaults instantiates a new EditAllowedHostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EditAllowedHostDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditAllowedHostDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditAllowedHostDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EditAllowedHostDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditAllowedHostDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *EditAllowedHostDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *EditAllowedHostDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *EditAllowedHostDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *EditAllowedHostDto) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *EditAllowedHostDto) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetMaskBits

`func (o *EditAllowedHostDto) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *EditAllowedHostDto) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *EditAllowedHostDto) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


