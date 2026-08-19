# CreateSecurityGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**SecurityGroupProtocol**](SecurityGroupProtocol.md) |  | [optional] 
**PortMinRange** | Pointer to **int32** |  | [optional] 
**PortMaxRange** | Pointer to **int32** |  | [optional] 
**RemoteIpPrefix** | Pointer to **NullableString** |  | [optional] 
**StandAloneProfileId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateSecurityGroupCommand

`func NewCreateSecurityGroupCommand() *CreateSecurityGroupCommand`

NewCreateSecurityGroupCommand instantiates a new CreateSecurityGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupCommandWithDefaults

`func NewCreateSecurityGroupCommandWithDefaults() *CreateSecurityGroupCommand`

NewCreateSecurityGroupCommandWithDefaults instantiates a new CreateSecurityGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSecurityGroupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecurityGroupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecurityGroupCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSecurityGroupCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateSecurityGroupCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateSecurityGroupCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocol

`func (o *CreateSecurityGroupCommand) GetProtocol() SecurityGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateSecurityGroupCommand) GetProtocolOk() (*SecurityGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateSecurityGroupCommand) SetProtocol(v SecurityGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CreateSecurityGroupCommand) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPortMinRange

`func (o *CreateSecurityGroupCommand) GetPortMinRange() int32`

GetPortMinRange returns the PortMinRange field if non-nil, zero value otherwise.

### GetPortMinRangeOk

`func (o *CreateSecurityGroupCommand) GetPortMinRangeOk() (*int32, bool)`

GetPortMinRangeOk returns a tuple with the PortMinRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMinRange

`func (o *CreateSecurityGroupCommand) SetPortMinRange(v int32)`

SetPortMinRange sets PortMinRange field to given value.

### HasPortMinRange

`func (o *CreateSecurityGroupCommand) HasPortMinRange() bool`

HasPortMinRange returns a boolean if a field has been set.

### GetPortMaxRange

`func (o *CreateSecurityGroupCommand) GetPortMaxRange() int32`

GetPortMaxRange returns the PortMaxRange field if non-nil, zero value otherwise.

### GetPortMaxRangeOk

`func (o *CreateSecurityGroupCommand) GetPortMaxRangeOk() (*int32, bool)`

GetPortMaxRangeOk returns a tuple with the PortMaxRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaxRange

`func (o *CreateSecurityGroupCommand) SetPortMaxRange(v int32)`

SetPortMaxRange sets PortMaxRange field to given value.

### HasPortMaxRange

`func (o *CreateSecurityGroupCommand) HasPortMaxRange() bool`

HasPortMaxRange returns a boolean if a field has been set.

### GetRemoteIpPrefix

`func (o *CreateSecurityGroupCommand) GetRemoteIpPrefix() string`

GetRemoteIpPrefix returns the RemoteIpPrefix field if non-nil, zero value otherwise.

### GetRemoteIpPrefixOk

`func (o *CreateSecurityGroupCommand) GetRemoteIpPrefixOk() (*string, bool)`

GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIpPrefix

`func (o *CreateSecurityGroupCommand) SetRemoteIpPrefix(v string)`

SetRemoteIpPrefix sets RemoteIpPrefix field to given value.

### HasRemoteIpPrefix

`func (o *CreateSecurityGroupCommand) HasRemoteIpPrefix() bool`

HasRemoteIpPrefix returns a boolean if a field has been set.

### SetRemoteIpPrefixNil

`func (o *CreateSecurityGroupCommand) SetRemoteIpPrefixNil(b bool)`

 SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil

### UnsetRemoteIpPrefix
`func (o *CreateSecurityGroupCommand) UnsetRemoteIpPrefix()`

UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
### GetStandAloneProfileId

`func (o *CreateSecurityGroupCommand) GetStandAloneProfileId() int32`

GetStandAloneProfileId returns the StandAloneProfileId field if non-nil, zero value otherwise.

### GetStandAloneProfileIdOk

`func (o *CreateSecurityGroupCommand) GetStandAloneProfileIdOk() (*int32, bool)`

GetStandAloneProfileIdOk returns a tuple with the StandAloneProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandAloneProfileId

`func (o *CreateSecurityGroupCommand) SetStandAloneProfileId(v int32)`

SetStandAloneProfileId sets StandAloneProfileId field to given value.

### HasStandAloneProfileId

`func (o *CreateSecurityGroupCommand) HasStandAloneProfileId() bool`

HasStandAloneProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


