# AzureSubnetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**AddressPrefix** | **NullableString** |  | 
**IsPublic** | **bool** |  | 
**IpMode** | [**IpMode**](IpMode.md) |  | 

## Methods

### NewAzureSubnetDto

`func NewAzureSubnetDto(name NullableString, addressPrefix NullableString, isPublic bool, ipMode IpMode, ) *AzureSubnetDto`

NewAzureSubnetDto instantiates a new AzureSubnetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubnetDtoWithDefaults

`func NewAzureSubnetDtoWithDefaults() *AzureSubnetDto`

NewAzureSubnetDtoWithDefaults instantiates a new AzureSubnetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureSubnetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureSubnetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureSubnetDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AzureSubnetDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureSubnetDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddressPrefix

`func (o *AzureSubnetDto) GetAddressPrefix() string`

GetAddressPrefix returns the AddressPrefix field if non-nil, zero value otherwise.

### GetAddressPrefixOk

`func (o *AzureSubnetDto) GetAddressPrefixOk() (*string, bool)`

GetAddressPrefixOk returns a tuple with the AddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPrefix

`func (o *AzureSubnetDto) SetAddressPrefix(v string)`

SetAddressPrefix sets AddressPrefix field to given value.


### SetAddressPrefixNil

`func (o *AzureSubnetDto) SetAddressPrefixNil(b bool)`

 SetAddressPrefixNil sets the value for AddressPrefix to be an explicit nil

### UnsetAddressPrefix
`func (o *AzureSubnetDto) UnsetAddressPrefix()`

UnsetAddressPrefix ensures that no value is present for AddressPrefix, not even an explicit nil
### GetIsPublic

`func (o *AzureSubnetDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AzureSubnetDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AzureSubnetDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetIpMode

`func (o *AzureSubnetDto) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AzureSubnetDto) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AzureSubnetDto) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


