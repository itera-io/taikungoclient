# PodsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Ready** | **NullableString** |  | 
**RestartCount** | **int32** |  | 
**CreatedAt** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Node** | **NullableString** |  | 
**Ip** | **NullableString** |  | 

## Methods

### NewPodsListDto

`func NewPodsListDto(state NullableString, name NullableString, ready NullableString, restartCount int32, createdAt NullableString, namespace NullableString, node NullableString, ip NullableString, ) *PodsListDto`

NewPodsListDto instantiates a new PodsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodsListDtoWithDefaults

`func NewPodsListDtoWithDefaults() *PodsListDto`

NewPodsListDtoWithDefaults instantiates a new PodsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PodsListDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PodsListDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PodsListDto) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *PodsListDto) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *PodsListDto) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetName

`func (o *PodsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodsListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PodsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PodsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetReady

`func (o *PodsListDto) GetReady() string`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *PodsListDto) GetReadyOk() (*string, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *PodsListDto) SetReady(v string)`

SetReady sets Ready field to given value.


### SetReadyNil

`func (o *PodsListDto) SetReadyNil(b bool)`

 SetReadyNil sets the value for Ready to be an explicit nil

### UnsetReady
`func (o *PodsListDto) UnsetReady()`

UnsetReady ensures that no value is present for Ready, not even an explicit nil
### GetRestartCount

`func (o *PodsListDto) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *PodsListDto) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *PodsListDto) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.


### GetCreatedAt

`func (o *PodsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PodsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PodsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PodsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PodsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetNamespace

`func (o *PodsListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PodsListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PodsListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *PodsListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *PodsListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetNode

`func (o *PodsListDto) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *PodsListDto) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *PodsListDto) SetNode(v string)`

SetNode sets Node field to given value.


### SetNodeNil

`func (o *PodsListDto) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *PodsListDto) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetIp

`func (o *PodsListDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PodsListDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PodsListDto) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *PodsListDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *PodsListDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


