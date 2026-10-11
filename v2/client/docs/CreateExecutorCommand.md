# CreateExecutorCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Enabled** | **bool** |  | 
**LocalCluster** | **bool** |  | 
**Continent** | **NullableString** |  | 
**KubeConfig** | **NullableString** |  | 

## Methods

### NewCreateExecutorCommand

`func NewCreateExecutorCommand(name NullableString, enabled bool, localCluster bool, continent NullableString, kubeConfig NullableString, ) *CreateExecutorCommand`

NewCreateExecutorCommand instantiates a new CreateExecutorCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExecutorCommandWithDefaults

`func NewCreateExecutorCommandWithDefaults() *CreateExecutorCommand`

NewCreateExecutorCommandWithDefaults instantiates a new CreateExecutorCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateExecutorCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExecutorCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExecutorCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateExecutorCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateExecutorCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *CreateExecutorCommand) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateExecutorCommand) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateExecutorCommand) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLocalCluster

`func (o *CreateExecutorCommand) GetLocalCluster() bool`

GetLocalCluster returns the LocalCluster field if non-nil, zero value otherwise.

### GetLocalClusterOk

`func (o *CreateExecutorCommand) GetLocalClusterOk() (*bool, bool)`

GetLocalClusterOk returns a tuple with the LocalCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCluster

`func (o *CreateExecutorCommand) SetLocalCluster(v bool)`

SetLocalCluster sets LocalCluster field to given value.


### GetContinent

`func (o *CreateExecutorCommand) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *CreateExecutorCommand) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *CreateExecutorCommand) SetContinent(v string)`

SetContinent sets Continent field to given value.


### SetContinentNil

`func (o *CreateExecutorCommand) SetContinentNil(b bool)`

 SetContinentNil sets the value for Continent to be an explicit nil

### UnsetContinent
`func (o *CreateExecutorCommand) UnsetContinent()`

UnsetContinent ensures that no value is present for Continent, not even an explicit nil
### GetKubeConfig

`func (o *CreateExecutorCommand) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *CreateExecutorCommand) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *CreateExecutorCommand) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.


### SetKubeConfigNil

`func (o *CreateExecutorCommand) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *CreateExecutorCommand) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


