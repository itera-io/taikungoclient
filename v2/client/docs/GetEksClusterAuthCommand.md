# GetEksClusterAuthCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**KubeConfigName** | **NullableString** |  | 
**ConfigToken** | **NullableString** |  | 

## Methods

### NewGetEksClusterAuthCommand

`func NewGetEksClusterAuthCommand(projectId int32, kubeConfigName NullableString, configToken NullableString, ) *GetEksClusterAuthCommand`

NewGetEksClusterAuthCommand instantiates a new GetEksClusterAuthCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEksClusterAuthCommandWithDefaults

`func NewGetEksClusterAuthCommandWithDefaults() *GetEksClusterAuthCommand`

NewGetEksClusterAuthCommandWithDefaults instantiates a new GetEksClusterAuthCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetEksClusterAuthCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetEksClusterAuthCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetEksClusterAuthCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetKubeConfigName

`func (o *GetEksClusterAuthCommand) GetKubeConfigName() string`

GetKubeConfigName returns the KubeConfigName field if non-nil, zero value otherwise.

### GetKubeConfigNameOk

`func (o *GetEksClusterAuthCommand) GetKubeConfigNameOk() (*string, bool)`

GetKubeConfigNameOk returns a tuple with the KubeConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfigName

`func (o *GetEksClusterAuthCommand) SetKubeConfigName(v string)`

SetKubeConfigName sets KubeConfigName field to given value.


### SetKubeConfigNameNil

`func (o *GetEksClusterAuthCommand) SetKubeConfigNameNil(b bool)`

 SetKubeConfigNameNil sets the value for KubeConfigName to be an explicit nil

### UnsetKubeConfigName
`func (o *GetEksClusterAuthCommand) UnsetKubeConfigName()`

UnsetKubeConfigName ensures that no value is present for KubeConfigName, not even an explicit nil
### GetConfigToken

`func (o *GetEksClusterAuthCommand) GetConfigToken() string`

GetConfigToken returns the ConfigToken field if non-nil, zero value otherwise.

### GetConfigTokenOk

`func (o *GetEksClusterAuthCommand) GetConfigTokenOk() (*string, bool)`

GetConfigTokenOk returns a tuple with the ConfigToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigToken

`func (o *GetEksClusterAuthCommand) SetConfigToken(v string)`

SetConfigToken sets ConfigToken field to given value.


### SetConfigTokenNil

`func (o *GetEksClusterAuthCommand) SetConfigTokenNil(b bool)`

 SetConfigTokenNil sets the value for ConfigToken to be an explicit nil

### UnsetConfigToken
`func (o *GetEksClusterAuthCommand) UnsetConfigToken()`

UnsetConfigToken ensures that no value is present for ConfigToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


