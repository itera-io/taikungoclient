# CreateKubeConfigCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**IsAccessibleForAll** | Pointer to **bool** |  | [optional] 
**IsAccessibleForManager** | Pointer to **bool** |  | [optional] 
**KubeConfigRoleId** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Namespace** | Pointer to **NullableString** |  | [optional] 
**Ttl** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateKubeConfigCommand

`func NewCreateKubeConfigCommand() *CreateKubeConfigCommand`

NewCreateKubeConfigCommand instantiates a new CreateKubeConfigCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKubeConfigCommandWithDefaults

`func NewCreateKubeConfigCommandWithDefaults() *CreateKubeConfigCommand`

NewCreateKubeConfigCommandWithDefaults instantiates a new CreateKubeConfigCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKubeConfigCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKubeConfigCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKubeConfigCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateKubeConfigCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateKubeConfigCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateKubeConfigCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *CreateKubeConfigCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateKubeConfigCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateKubeConfigCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateKubeConfigCommand) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetIsAccessibleForAll

`func (o *CreateKubeConfigCommand) GetIsAccessibleForAll() bool`

GetIsAccessibleForAll returns the IsAccessibleForAll field if non-nil, zero value otherwise.

### GetIsAccessibleForAllOk

`func (o *CreateKubeConfigCommand) GetIsAccessibleForAllOk() (*bool, bool)`

GetIsAccessibleForAllOk returns a tuple with the IsAccessibleForAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleForAll

`func (o *CreateKubeConfigCommand) SetIsAccessibleForAll(v bool)`

SetIsAccessibleForAll sets IsAccessibleForAll field to given value.

### HasIsAccessibleForAll

`func (o *CreateKubeConfigCommand) HasIsAccessibleForAll() bool`

HasIsAccessibleForAll returns a boolean if a field has been set.

### GetIsAccessibleForManager

`func (o *CreateKubeConfigCommand) GetIsAccessibleForManager() bool`

GetIsAccessibleForManager returns the IsAccessibleForManager field if non-nil, zero value otherwise.

### GetIsAccessibleForManagerOk

`func (o *CreateKubeConfigCommand) GetIsAccessibleForManagerOk() (*bool, bool)`

GetIsAccessibleForManagerOk returns a tuple with the IsAccessibleForManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessibleForManager

`func (o *CreateKubeConfigCommand) SetIsAccessibleForManager(v bool)`

SetIsAccessibleForManager sets IsAccessibleForManager field to given value.

### HasIsAccessibleForManager

`func (o *CreateKubeConfigCommand) HasIsAccessibleForManager() bool`

HasIsAccessibleForManager returns a boolean if a field has been set.

### GetKubeConfigRoleId

`func (o *CreateKubeConfigCommand) GetKubeConfigRoleId() int32`

GetKubeConfigRoleId returns the KubeConfigRoleId field if non-nil, zero value otherwise.

### GetKubeConfigRoleIdOk

`func (o *CreateKubeConfigCommand) GetKubeConfigRoleIdOk() (*int32, bool)`

GetKubeConfigRoleIdOk returns a tuple with the KubeConfigRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigRoleId

`func (o *CreateKubeConfigCommand) SetKubeConfigRoleId(v int32)`

SetKubeConfigRoleId sets KubeConfigRoleId field to given value.

### HasKubeConfigRoleId

`func (o *CreateKubeConfigCommand) HasKubeConfigRoleId() bool`

HasKubeConfigRoleId returns a boolean if a field has been set.

### GetUserId

`func (o *CreateKubeConfigCommand) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateKubeConfigCommand) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateKubeConfigCommand) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateKubeConfigCommand) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateKubeConfigCommand) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateKubeConfigCommand) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetNamespace

`func (o *CreateKubeConfigCommand) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateKubeConfigCommand) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateKubeConfigCommand) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateKubeConfigCommand) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### SetNamespaceNil

`func (o *CreateKubeConfigCommand) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *CreateKubeConfigCommand) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetTtl

`func (o *CreateKubeConfigCommand) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *CreateKubeConfigCommand) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *CreateKubeConfigCommand) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *CreateKubeConfigCommand) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *CreateKubeConfigCommand) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *CreateKubeConfigCommand) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


