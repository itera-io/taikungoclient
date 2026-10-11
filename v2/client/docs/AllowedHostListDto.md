# AllowedHostListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Description** | **NullableString** |  | 
**IpAddress** | **string** |  | 
**MaskBits** | **int32** |  | 
**AccessProfileId** | **int32** |  | 
**AccessProfileName** | **string** |  | 

## Methods

### NewAllowedHostListDto

`func NewAllowedHostListDto(id int32, description NullableString, ipAddress string, maskBits int32, accessProfileId int32, accessProfileName string, ) *AllowedHostListDto`

NewAllowedHostListDto instantiates a new AllowedHostListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedHostListDtoWithDefaults

`func NewAllowedHostListDtoWithDefaults() *AllowedHostListDto`

NewAllowedHostListDtoWithDefaults instantiates a new AllowedHostListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedHostListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedHostListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedHostListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetDescription

`func (o *AllowedHostListDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowedHostListDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowedHostListDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AllowedHostListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AllowedHostListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *AllowedHostListDto) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AllowedHostListDto) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AllowedHostListDto) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetMaskBits

`func (o *AllowedHostListDto) GetMaskBits() int32`

GetMaskBits returns the MaskBits field if non-nil, zero value otherwise.

### GetMaskBitsOk

`func (o *AllowedHostListDto) GetMaskBitsOk() (*int32, bool)`

GetMaskBitsOk returns a tuple with the MaskBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskBits

`func (o *AllowedHostListDto) SetMaskBits(v int32)`

SetMaskBits sets MaskBits field to given value.


### GetAccessProfileId

`func (o *AllowedHostListDto) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *AllowedHostListDto) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *AllowedHostListDto) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.


### GetAccessProfileName

`func (o *AllowedHostListDto) GetAccessProfileName() string`

GetAccessProfileName returns the AccessProfileName field if non-nil, zero value otherwise.

### GetAccessProfileNameOk

`func (o *AllowedHostListDto) GetAccessProfileNameOk() (*string, bool)`

GetAccessProfileNameOk returns a tuple with the AccessProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileName

`func (o *AllowedHostListDto) SetAccessProfileName(v string)`

SetAccessProfileName sets AccessProfileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


