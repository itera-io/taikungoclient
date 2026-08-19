# ShowbackProjectsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**MonitoringUrl** | Pointer to **NullableString** |  | [optional] 
**MonitoringUrlHeader** | Pointer to **NullableString** |  | [optional] 
**MonitoringUsername** | Pointer to **NullableString** |  | [optional] 
**MonitoringPassword** | Pointer to **NullableString** |  | [optional] 
**KubernetesCurrentVersion** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShowbackProjectsListDto

`func NewShowbackProjectsListDto() *ShowbackProjectsListDto`

NewShowbackProjectsListDto instantiates a new ShowbackProjectsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowbackProjectsListDtoWithDefaults

`func NewShowbackProjectsListDtoWithDefaults() *ShowbackProjectsListDto`

NewShowbackProjectsListDtoWithDefaults instantiates a new ShowbackProjectsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShowbackProjectsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShowbackProjectsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShowbackProjectsListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShowbackProjectsListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShowbackProjectsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShowbackProjectsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShowbackProjectsListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShowbackProjectsListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShowbackProjectsListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShowbackProjectsListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMonitoringUrl

`func (o *ShowbackProjectsListDto) GetMonitoringUrl() string`

GetMonitoringUrl returns the MonitoringUrl field if non-nil, zero value otherwise.

### GetMonitoringUrlOk

`func (o *ShowbackProjectsListDto) GetMonitoringUrlOk() (*string, bool)`

GetMonitoringUrlOk returns a tuple with the MonitoringUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUrl

`func (o *ShowbackProjectsListDto) SetMonitoringUrl(v string)`

SetMonitoringUrl sets MonitoringUrl field to given value.

### HasMonitoringUrl

`func (o *ShowbackProjectsListDto) HasMonitoringUrl() bool`

HasMonitoringUrl returns a boolean if a field has been set.

### SetMonitoringUrlNil

`func (o *ShowbackProjectsListDto) SetMonitoringUrlNil(b bool)`

 SetMonitoringUrlNil sets the value for MonitoringUrl to be an explicit nil

### UnsetMonitoringUrl
`func (o *ShowbackProjectsListDto) UnsetMonitoringUrl()`

UnsetMonitoringUrl ensures that no value is present for MonitoringUrl, not even an explicit nil
### GetMonitoringUrlHeader

`func (o *ShowbackProjectsListDto) GetMonitoringUrlHeader() string`

GetMonitoringUrlHeader returns the MonitoringUrlHeader field if non-nil, zero value otherwise.

### GetMonitoringUrlHeaderOk

`func (o *ShowbackProjectsListDto) GetMonitoringUrlHeaderOk() (*string, bool)`

GetMonitoringUrlHeaderOk returns a tuple with the MonitoringUrlHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUrlHeader

`func (o *ShowbackProjectsListDto) SetMonitoringUrlHeader(v string)`

SetMonitoringUrlHeader sets MonitoringUrlHeader field to given value.

### HasMonitoringUrlHeader

`func (o *ShowbackProjectsListDto) HasMonitoringUrlHeader() bool`

HasMonitoringUrlHeader returns a boolean if a field has been set.

### SetMonitoringUrlHeaderNil

`func (o *ShowbackProjectsListDto) SetMonitoringUrlHeaderNil(b bool)`

 SetMonitoringUrlHeaderNil sets the value for MonitoringUrlHeader to be an explicit nil

### UnsetMonitoringUrlHeader
`func (o *ShowbackProjectsListDto) UnsetMonitoringUrlHeader()`

UnsetMonitoringUrlHeader ensures that no value is present for MonitoringUrlHeader, not even an explicit nil
### GetMonitoringUsername

`func (o *ShowbackProjectsListDto) GetMonitoringUsername() string`

GetMonitoringUsername returns the MonitoringUsername field if non-nil, zero value otherwise.

### GetMonitoringUsernameOk

`func (o *ShowbackProjectsListDto) GetMonitoringUsernameOk() (*string, bool)`

GetMonitoringUsernameOk returns a tuple with the MonitoringUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringUsername

`func (o *ShowbackProjectsListDto) SetMonitoringUsername(v string)`

SetMonitoringUsername sets MonitoringUsername field to given value.

### HasMonitoringUsername

`func (o *ShowbackProjectsListDto) HasMonitoringUsername() bool`

HasMonitoringUsername returns a boolean if a field has been set.

### SetMonitoringUsernameNil

`func (o *ShowbackProjectsListDto) SetMonitoringUsernameNil(b bool)`

 SetMonitoringUsernameNil sets the value for MonitoringUsername to be an explicit nil

### UnsetMonitoringUsername
`func (o *ShowbackProjectsListDto) UnsetMonitoringUsername()`

UnsetMonitoringUsername ensures that no value is present for MonitoringUsername, not even an explicit nil
### GetMonitoringPassword

`func (o *ShowbackProjectsListDto) GetMonitoringPassword() string`

GetMonitoringPassword returns the MonitoringPassword field if non-nil, zero value otherwise.

### GetMonitoringPasswordOk

`func (o *ShowbackProjectsListDto) GetMonitoringPasswordOk() (*string, bool)`

GetMonitoringPasswordOk returns a tuple with the MonitoringPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringPassword

`func (o *ShowbackProjectsListDto) SetMonitoringPassword(v string)`

SetMonitoringPassword sets MonitoringPassword field to given value.

### HasMonitoringPassword

`func (o *ShowbackProjectsListDto) HasMonitoringPassword() bool`

HasMonitoringPassword returns a boolean if a field has been set.

### SetMonitoringPasswordNil

`func (o *ShowbackProjectsListDto) SetMonitoringPasswordNil(b bool)`

 SetMonitoringPasswordNil sets the value for MonitoringPassword to be an explicit nil

### UnsetMonitoringPassword
`func (o *ShowbackProjectsListDto) UnsetMonitoringPassword()`

UnsetMonitoringPassword ensures that no value is present for MonitoringPassword, not even an explicit nil
### GetKubernetesCurrentVersion

`func (o *ShowbackProjectsListDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ShowbackProjectsListDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ShowbackProjectsListDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.

### HasKubernetesCurrentVersion

`func (o *ShowbackProjectsListDto) HasKubernetesCurrentVersion() bool`

HasKubernetesCurrentVersion returns a boolean if a field has been set.

### SetKubernetesCurrentVersionNil

`func (o *ShowbackProjectsListDto) SetKubernetesCurrentVersionNil(b bool)`

 SetKubernetesCurrentVersionNil sets the value for KubernetesCurrentVersion to be an explicit nil

### UnsetKubernetesCurrentVersion
`func (o *ShowbackProjectsListDto) UnsetKubernetesCurrentVersion()`

UnsetKubernetesCurrentVersion ensures that no value is present for KubernetesCurrentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


