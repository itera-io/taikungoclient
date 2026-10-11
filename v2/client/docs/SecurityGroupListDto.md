# SecurityGroupListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Protocol** | **NullableString** |  | 
**PortMinRange** | **int32** |  | 
**PortMaxRange** | **int32** |  | 
**RemoteIpPrefix** | **NullableString** |  | 
**ProfileName** | **NullableString** |  | 

## Methods

### NewSecurityGroupListDto

`func NewSecurityGroupListDto(id int32, name NullableString, protocol NullableString, portMinRange int32, portMaxRange int32, remoteIpPrefix NullableString, profileName NullableString, ) *SecurityGroupListDto`

NewSecurityGroupListDto instantiates a new SecurityGroupListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupListDtoWithDefaults

`func NewSecurityGroupListDtoWithDefaults() *SecurityGroupListDto`

NewSecurityGroupListDtoWithDefaults instantiates a new SecurityGroupListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityGroupListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *SecurityGroupListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityGroupListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocol

`func (o *SecurityGroupListDto) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SecurityGroupListDto) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SecurityGroupListDto) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *SecurityGroupListDto) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *SecurityGroupListDto) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPortMinRange

`func (o *SecurityGroupListDto) GetPortMinRange() int32`

GetPortMinRange returns the PortMinRange field if non-nil, zero value otherwise.

### GetPortMinRangeOk

`func (o *SecurityGroupListDto) GetPortMinRangeOk() (*int32, bool)`

GetPortMinRangeOk returns a tuple with the PortMinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMinRange

`func (o *SecurityGroupListDto) SetPortMinRange(v int32)`

SetPortMinRange sets PortMinRange field to given value.


### GetPortMaxRange

`func (o *SecurityGroupListDto) GetPortMaxRange() int32`

GetPortMaxRange returns the PortMaxRange field if non-nil, zero value otherwise.

### GetPortMaxRangeOk

`func (o *SecurityGroupListDto) GetPortMaxRangeOk() (*int32, bool)`

GetPortMaxRangeOk returns a tuple with the PortMaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaxRange

`func (o *SecurityGroupListDto) SetPortMaxRange(v int32)`

SetPortMaxRange sets PortMaxRange field to given value.


### GetRemoteIpPrefix

`func (o *SecurityGroupListDto) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *SecurityGroupListDto) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *SecurityGroupListDto) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.


### SetRemoteIpPrefixNil

`func (o *SecurityGroupListDto) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *SecurityGroupListDto) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetProfileName

`func (o *SecurityGroupListDto) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *SecurityGroupListDto) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *SecurityGroupListDto) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### SetProfileNameNil

`func (o *SecurityGroupListDto) SetProfileNameNil(b bool)`

 SetProfileNameNil sets the value for ProfileName to be an explicit nil

### UnsetProfileName
`func (o *SecurityGroupListDto) UnsetProfileName()`

UnsetProfileName ensures that no value is present for ProfileName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


