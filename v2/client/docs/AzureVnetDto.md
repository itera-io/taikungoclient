# AzureVnetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**AddressSpace** | **NullableString** |  | 
**IpMode** | [**IpMode**](IpMode.md) |  | 

## Methods

### NewAzureVnetDto

`func NewAzureVnetDto(name NullableString, addressSpace NullableString, ipMode IpMode, ) *AzureVnetDto`

NewAzureVnetDto instantiates a new AzureVnetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVnetDtoWithDefaults

`func NewAzureVnetDtoWithDefaults() *AzureVnetDto`

NewAzureVnetDtoWithDefaults instantiates a new AzureVnetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureVnetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureVnetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureVnetDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AzureVnetDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureVnetDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddressSpace

`func (o *AzureVnetDto) GetAddressSpace() string`

GetAddressSpace returns the AddressSpace field if non-nil, zero value otherwise.

### GetAddressSpaceOk

`func (o *AzureVnetDto) GetAddressSpaceOk() (*string, bool)`

GetAddressSpaceOk returns a tuple with the AddressSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSpace

`func (o *AzureVnetDto) SetAddressSpace(v string)`

SetAddressSpace sets AddressSpace field to given value.


### SetAddressSpaceNil

`func (o *AzureVnetDto) SetAddressSpaceNil(b bool)`

 SetAddressSpaceNil sets the value for AddressSpace to be an explicit nil

### UnsetAddressSpace
`func (o *AzureVnetDto) UnsetAddressSpace()`

UnsetAddressSpace ensures that no value is present for AddressSpace, not even an explicit nil
### GetIpMode

`func (o *AzureVnetDto) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AzureVnetDto) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AzureVnetDto) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


