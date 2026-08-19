# OpenstackNetworkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkLimit** | Pointer to **int64** |  | [optional] 
**SubnetLimit** | Pointer to **int64** |  | [optional] 
**FloatingIpLimit** | Pointer to **int64** |  | [optional] 
**RouterLimit** | Pointer to **int64** |  | [optional] 
**SecurityGroupLimit** | Pointer to **int64** |  | [optional] 
**SecurityGroupRuleLimit** | Pointer to **int64** |  | [optional] 
**PortLimit** | Pointer to **int64** |  | [optional] 
**NetworkUsed** | Pointer to **int64** |  | [optional] 
**SubnetUsed** | Pointer to **int64** |  | [optional] 
**FloatingIpUsed** | Pointer to **int64** |  | [optional] 
**RouterUsed** | Pointer to **int64** |  | [optional] 
**SecurityGroupUsed** | Pointer to **int64** |  | [optional] 
**PortUsed** | Pointer to **int64** |  | [optional] 
**SecurityGroupRuleUsed** | Pointer to **int64** |  | [optional] 

## Methods

### NewOpenstackNetworkDto

`func NewOpenstackNetworkDto() *OpenstackNetworkDto`

NewOpenstackNetworkDto instantiates a new OpenstackNetworkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenstackNetworkDtoWithDefaults

`func NewOpenstackNetworkDtoWithDefaults() *OpenstackNetworkDto`

NewOpenstackNetworkDtoWithDefaults instantiates a new OpenstackNetworkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkLimit

`func (o *OpenstackNetworkDto) GetNetworkLimit() int64`

GetNetworkLimit returns the NetworkLimit field if non-nil, zero value otherwise.

### GetNetworkLimitOk

`func (o *OpenstackNetworkDto) GetNetworkLimitOk() (*int64, bool)`

GetNetworkLimitOk returns a tuple with the NetworkLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLimit

`func (o *OpenstackNetworkDto) SetNetworkLimit(v int64)`

SetNetworkLimit sets NetworkLimit field to given value.

### HasNetworkLimit

`func (o *OpenstackNetworkDto) HasNetworkLimit() bool`

HasNetworkLimit returns a boolean if a field has been set.

### GetSubnetLimit

`func (o *OpenstackNetworkDto) GetSubnetLimit() int64`

GetSubnetLimit returns the SubnetLimit field if non-nil, zero value otherwise.

### GetSubnetLimitOk

`func (o *OpenstackNetworkDto) GetSubnetLimitOk() (*int64, bool)`

GetSubnetLimitOk returns a tuple with the SubnetLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetLimit

`func (o *OpenstackNetworkDto) SetSubnetLimit(v int64)`

SetSubnetLimit sets SubnetLimit field to given value.

### HasSubnetLimit

`func (o *OpenstackNetworkDto) HasSubnetLimit() bool`

HasSubnetLimit returns a boolean if a field has been set.

### GetFloatingIpLimit

`func (o *OpenstackNetworkDto) GetFloatingIpLimit() int64`

GetFloatingIpLimit returns the FloatingIpLimit field if non-nil, zero value otherwise.

### GetFloatingIpLimitOk

`func (o *OpenstackNetworkDto) GetFloatingIpLimitOk() (*int64, bool)`

GetFloatingIpLimitOk returns a tuple with the FloatingIpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpLimit

`func (o *OpenstackNetworkDto) SetFloatingIpLimit(v int64)`

SetFloatingIpLimit sets FloatingIpLimit field to given value.

### HasFloatingIpLimit

`func (o *OpenstackNetworkDto) HasFloatingIpLimit() bool`

HasFloatingIpLimit returns a boolean if a field has been set.

### GetRouterLimit

`func (o *OpenstackNetworkDto) GetRouterLimit() int64`

GetRouterLimit returns the RouterLimit field if non-nil, zero value otherwise.

### GetRouterLimitOk

`func (o *OpenstackNetworkDto) GetRouterLimitOk() (*int64, bool)`

GetRouterLimitOk returns a tuple with the RouterLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterLimit

`func (o *OpenstackNetworkDto) SetRouterLimit(v int64)`

SetRouterLimit sets RouterLimit field to given value.

### HasRouterLimit

`func (o *OpenstackNetworkDto) HasRouterLimit() bool`

HasRouterLimit returns a boolean if a field has been set.

### GetSecurityGroupLimit

`func (o *OpenstackNetworkDto) GetSecurityGroupLimit() int64`

GetSecurityGroupLimit returns the SecurityGroupLimit field if non-nil, zero value otherwise.

### GetSecurityGroupLimitOk

`func (o *OpenstackNetworkDto) GetSecurityGroupLimitOk() (*int64, bool)`

GetSecurityGroupLimitOk returns a tuple with the SecurityGroupLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupLimit

`func (o *OpenstackNetworkDto) SetSecurityGroupLimit(v int64)`

SetSecurityGroupLimit sets SecurityGroupLimit field to given value.

### HasSecurityGroupLimit

`func (o *OpenstackNetworkDto) HasSecurityGroupLimit() bool`

HasSecurityGroupLimit returns a boolean if a field has been set.

### GetSecurityGroupRuleLimit

`func (o *OpenstackNetworkDto) GetSecurityGroupRuleLimit() int64`

GetSecurityGroupRuleLimit returns the SecurityGroupRuleLimit field if non-nil, zero value otherwise.

### GetSecurityGroupRuleLimitOk

`func (o *OpenstackNetworkDto) GetSecurityGroupRuleLimitOk() (*int64, bool)`

GetSecurityGroupRuleLimitOk returns a tuple with the SecurityGroupRuleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRuleLimit

`func (o *OpenstackNetworkDto) SetSecurityGroupRuleLimit(v int64)`

SetSecurityGroupRuleLimit sets SecurityGroupRuleLimit field to given value.

### HasSecurityGroupRuleLimit

`func (o *OpenstackNetworkDto) HasSecurityGroupRuleLimit() bool`

HasSecurityGroupRuleLimit returns a boolean if a field has been set.

### GetPortLimit

`func (o *OpenstackNetworkDto) GetPortLimit() int64`

GetPortLimit returns the PortLimit field if non-nil, zero value otherwise.

### GetPortLimitOk

`func (o *OpenstackNetworkDto) GetPortLimitOk() (*int64, bool)`

GetPortLimitOk returns a tuple with the PortLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLimit

`func (o *OpenstackNetworkDto) SetPortLimit(v int64)`

SetPortLimit sets PortLimit field to given value.

### HasPortLimit

`func (o *OpenstackNetworkDto) HasPortLimit() bool`

HasPortLimit returns a boolean if a field has been set.

### GetNetworkUsed

`func (o *OpenstackNetworkDto) GetNetworkUsed() int64`

GetNetworkUsed returns the NetworkUsed field if non-nil, zero value otherwise.

### GetNetworkUsedOk

`func (o *OpenstackNetworkDto) GetNetworkUsedOk() (*int64, bool)`

GetNetworkUsedOk returns a tuple with the NetworkUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUsed

`func (o *OpenstackNetworkDto) SetNetworkUsed(v int64)`

SetNetworkUsed sets NetworkUsed field to given value.

### HasNetworkUsed

`func (o *OpenstackNetworkDto) HasNetworkUsed() bool`

HasNetworkUsed returns a boolean if a field has been set.

### GetSubnetUsed

`func (o *OpenstackNetworkDto) GetSubnetUsed() int64`

GetSubnetUsed returns the SubnetUsed field if non-nil, zero value otherwise.

### GetSubnetUsedOk

`func (o *OpenstackNetworkDto) GetSubnetUsedOk() (*int64, bool)`

GetSubnetUsedOk returns a tuple with the SubnetUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetUsed

`func (o *OpenstackNetworkDto) SetSubnetUsed(v int64)`

SetSubnetUsed sets SubnetUsed field to given value.

### HasSubnetUsed

`func (o *OpenstackNetworkDto) HasSubnetUsed() bool`

HasSubnetUsed returns a boolean if a field has been set.

### GetFloatingIpUsed

`func (o *OpenstackNetworkDto) GetFloatingIpUsed() int64`

GetFloatingIpUsed returns the FloatingIpUsed field if non-nil, zero value otherwise.

### GetFloatingIpUsedOk

`func (o *OpenstackNetworkDto) GetFloatingIpUsedOk() (*int64, bool)`

GetFloatingIpUsedOk returns a tuple with the FloatingIpUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpUsed

`func (o *OpenstackNetworkDto) SetFloatingIpUsed(v int64)`

SetFloatingIpUsed sets FloatingIpUsed field to given value.

### HasFloatingIpUsed

`func (o *OpenstackNetworkDto) HasFloatingIpUsed() bool`

HasFloatingIpUsed returns a boolean if a field has been set.

### GetRouterUsed

`func (o *OpenstackNetworkDto) GetRouterUsed() int64`

GetRouterUsed returns the RouterUsed field if non-nil, zero value otherwise.

### GetRouterUsedOk

`func (o *OpenstackNetworkDto) GetRouterUsedOk() (*int64, bool)`

GetRouterUsedOk returns a tuple with the RouterUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterUsed

`func (o *OpenstackNetworkDto) SetRouterUsed(v int64)`

SetRouterUsed sets RouterUsed field to given value.

### HasRouterUsed

`func (o *OpenstackNetworkDto) HasRouterUsed() bool`

HasRouterUsed returns a boolean if a field has been set.

### GetSecurityGroupUsed

`func (o *OpenstackNetworkDto) GetSecurityGroupUsed() int64`

GetSecurityGroupUsed returns the SecurityGroupUsed field if non-nil, zero value otherwise.

### GetSecurityGroupUsedOk

`func (o *OpenstackNetworkDto) GetSecurityGroupUsedOk() (*int64, bool)`

GetSecurityGroupUsedOk returns a tuple with the SecurityGroupUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupUsed

`func (o *OpenstackNetworkDto) SetSecurityGroupUsed(v int64)`

SetSecurityGroupUsed sets SecurityGroupUsed field to given value.

### HasSecurityGroupUsed

`func (o *OpenstackNetworkDto) HasSecurityGroupUsed() bool`

HasSecurityGroupUsed returns a boolean if a field has been set.

### GetPortUsed

`func (o *OpenstackNetworkDto) GetPortUsed() int64`

GetPortUsed returns the PortUsed field if non-nil, zero value otherwise.

### GetPortUsedOk

`func (o *OpenstackNetworkDto) GetPortUsedOk() (*int64, bool)`

GetPortUsedOk returns a tuple with the PortUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUsed

`func (o *OpenstackNetworkDto) SetPortUsed(v int64)`

SetPortUsed sets PortUsed field to given value.

### HasPortUsed

`func (o *OpenstackNetworkDto) HasPortUsed() bool`

HasPortUsed returns a boolean if a field has been set.

### GetSecurityGroupRuleUsed

`func (o *OpenstackNetworkDto) GetSecurityGroupRuleUsed() int64`

GetSecurityGroupRuleUsed returns the SecurityGroupRuleUsed field if non-nil, zero value otherwise.

### GetSecurityGroupRuleUsedOk

`func (o *OpenstackNetworkDto) GetSecurityGroupRuleUsedOk() (*int64, bool)`

GetSecurityGroupRuleUsedOk returns a tuple with the SecurityGroupRuleUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRuleUsed

`func (o *OpenstackNetworkDto) SetSecurityGroupRuleUsed(v int64)`

SetSecurityGroupRuleUsed sets SecurityGroupRuleUsed field to given value.

### HasSecurityGroupRuleUsed

`func (o *OpenstackNetworkDto) HasSecurityGroupRuleUsed() bool`

HasSecurityGroupRuleUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


