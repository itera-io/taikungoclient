# ServiceListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Type** | **NullableString** |  | 
**ClusterIp** | **NullableString** |  | 
**ExternalIp** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 

## Methods

### NewServiceListDto

`func NewServiceListDto(name NullableString, namespace NullableString, type_ NullableString, clusterIp NullableString, externalIp NullableString, createdAt NullableString, ) *ServiceListDto`

NewServiceListDto instantiates a new ServiceListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceListDtoWithDefaults

`func NewServiceListDtoWithDefaults() *ServiceListDto`

NewServiceListDtoWithDefaults instantiates a new ServiceListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ServiceListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ServiceListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamespace

`func (o *ServiceListDto) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ServiceListDto) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ServiceListDto) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *ServiceListDto) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ServiceListDto) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetType

`func (o *ServiceListDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceListDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceListDto) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ServiceListDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ServiceListDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetClusterIp

`func (o *ServiceListDto) GetClusterIp() string`

GetClusterIp returns the ClusterIp field if non-nil, zero value otherwise.

### GetClusterIpOk

`func (o *ServiceListDto) GetClusterIpOk() (*string, bool)`

GetClusterIpOk returns a tuple with the ClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIp

`func (o *ServiceListDto) SetClusterIp(v string)`

SetClusterIp sets ClusterIp field to given value.


### SetClusterIpNil

`func (o *ServiceListDto) SetClusterIpNil(b bool)`

 SetClusterIpNil sets the value for ClusterIp to be an explicit nil

### UnsetClusterIp
`func (o *ServiceListDto) UnsetClusterIp()`

UnsetClusterIp ensures that no value is present for ClusterIp, not even an explicit nil
### GetExternalIp

`func (o *ServiceListDto) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ServiceListDto) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ServiceListDto) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.


### SetExternalIpNil

`func (o *ServiceListDto) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *ServiceListDto) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetCreatedAt

`func (o *ServiceListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ServiceListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


