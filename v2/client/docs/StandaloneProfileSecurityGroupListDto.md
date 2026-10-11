# StandaloneProfileSecurityGroupListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Protocol** | **NullableString** |  | 
**PortMinRange** | **int32** |  | 
**PortMaxRange** | **int32** |  | 
**RemoteIpPrefix** | **NullableString** |  | 

## Methods

### NewStandaloneProfileSecurityGroupListDto

`func NewStandaloneProfileSecurityGroupListDto(id int32, name NullableString, protocol NullableString, portMinRange int32, portMaxRange int32, remoteIpPrefix NullableString, ) *StandaloneProfileSecurityGroupListDto`

NewStandaloneProfileSecurityGroupListDto instantiates a new StandaloneProfileSecurityGroupListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneProfileSecurityGroupListDtoWithDefaults

`func NewStandaloneProfileSecurityGroupListDtoWithDefaults() *StandaloneProfileSecurityGroupListDto`

NewStandaloneProfileSecurityGroupListDtoWithDefaults instantiates a new StandaloneProfileSecurityGroupListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandaloneProfileSecurityGroupListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandaloneProfileSecurityGroupListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandaloneProfileSecurityGroupListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandaloneProfileSecurityGroupListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandaloneProfileSecurityGroupListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandaloneProfileSecurityGroupListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StandaloneProfileSecurityGroupListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandaloneProfileSecurityGroupListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocol

`func (o *StandaloneProfileSecurityGroupListDto) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StandaloneProfileSecurityGroupListDto) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StandaloneProfileSecurityGroupListDto) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### SetProtocolNil

`func (o *StandaloneProfileSecurityGroupListDto) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *StandaloneProfileSecurityGroupListDto) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetPortMinRange

`func (o *StandaloneProfileSecurityGroupListDto) GetPortMinRange() int32`

GetPortMinRange returns the PortMinRange field if non-nil, zero value otherwise.

### GetPortMinRangeOk

`func (o *StandaloneProfileSecurityGroupListDto) GetPortMinRangeOk() (*int32, bool)`

GetPortMinRangeOk returns a tuple with the PortMinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMinRange

`func (o *StandaloneProfileSecurityGroupListDto) SetPortMinRange(v int32)`

SetPortMinRange sets PortMinRange field to given value.


### GetPortMaxRange

`func (o *StandaloneProfileSecurityGroupListDto) GetPortMaxRange() int32`

GetPortMaxRange returns the PortMaxRange field if non-nil, zero value otherwise.

### GetPortMaxRangeOk

`func (o *StandaloneProfileSecurityGroupListDto) GetPortMaxRangeOk() (*int32, bool)`

GetPortMaxRangeOk returns a tuple with the PortMaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaxRange

`func (o *StandaloneProfileSecurityGroupListDto) SetPortMaxRange(v int32)`

SetPortMaxRange sets PortMaxRange field to given value.


### GetRemoteIpPrefix

`func (o *StandaloneProfileSecurityGroupListDto) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *StandaloneProfileSecurityGroupListDto) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *StandaloneProfileSecurityGroupListDto) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.


### SetRemoteIpPrefixNil

`func (o *StandaloneProfileSecurityGroupListDto) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *StandaloneProfileSecurityGroupListDto) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


