# InternalProjectListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Token** | **NullableString** |  | 
**Status** | **int32** |  | 
**OrganizationId** | **int32** |  | 
**Health** | **int32** |  | 
**ImportClusterType** | **int32** |  | 
**KubernetesVersion** | **NullableString** |  | 
**AccessIp** | **NullableString** |  | 
**IsMonitoringEnabled** | **bool** |  | 
**Username** | **NullableString** |  | 
**Password** | **NullableString** |  | 
**PrometheusUrl** | **NullableString** |  | 
**AlertManagerUrl** | **NullableString** |  | 
**LokiUrl** | **NullableString** |  | 

## Methods

### NewInternalProjectListDto

`func NewInternalProjectListDto(id int32, name string, token NullableString, status int32, organizationId int32, health int32, importClusterType int32, kubernetesVersion NullableString, accessIp NullableString, isMonitoringEnabled bool, username NullableString, password NullableString, prometheusUrl NullableString, alertManagerUrl NullableString, lokiUrl NullableString, ) *InternalProjectListDto`

NewInternalProjectListDto instantiates a new InternalProjectListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalProjectListDtoWithDefaults

`func NewInternalProjectListDtoWithDefaults() *InternalProjectListDto`

NewInternalProjectListDtoWithDefaults instantiates a new InternalProjectListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalProjectListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalProjectListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalProjectListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InternalProjectListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalProjectListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalProjectListDto) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *InternalProjectListDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InternalProjectListDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InternalProjectListDto) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *InternalProjectListDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *InternalProjectListDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetStatus

`func (o *InternalProjectListDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalProjectListDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalProjectListDto) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetOrganizationId

`func (o *InternalProjectListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *InternalProjectListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *InternalProjectListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetHealth

`func (o *InternalProjectListDto) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *InternalProjectListDto) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *InternalProjectListDto) SetHealth(v int32)`

SetHealth sets Health field to given value.


### GetImportClusterType

`func (o *InternalProjectListDto) GetImportClusterType() int32`

GetImportClusterType returns the ImportClusterType field if non-nil, zero value otherwise.

### GetImportClusterTypeOk

`func (o *InternalProjectListDto) GetImportClusterTypeOk() (*int32, bool)`

GetImportClusterTypeOk returns a tuple with the ImportClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportClusterType

`func (o *InternalProjectListDto) SetImportClusterType(v int32)`

SetImportClusterType sets ImportClusterType field to given value.


### GetKubernetesVersion

`func (o *InternalProjectListDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *InternalProjectListDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *InternalProjectListDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *InternalProjectListDto) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *InternalProjectListDto) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetAccessIp

`func (o *InternalProjectListDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *InternalProjectListDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *InternalProjectListDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.


### SetAccessIpNil

`func (o *InternalProjectListDto) SetAccessIpNil(b bool)`

 SetAccessIpNil sets the value for AccessIp to be an explicit nil

### UnsetAccessIp
`func (o *InternalProjectListDto) UnsetAccessIp()`

UnsetAccessIp ensures that no value is present for AccessIp, not even an explicit nil
### GetIsMonitoringEnabled

`func (o *InternalProjectListDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *InternalProjectListDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *InternalProjectListDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.


### GetUsername

`func (o *InternalProjectListDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InternalProjectListDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InternalProjectListDto) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *InternalProjectListDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *InternalProjectListDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *InternalProjectListDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InternalProjectListDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InternalProjectListDto) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *InternalProjectListDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *InternalProjectListDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPrometheusUrl

`func (o *InternalProjectListDto) GetPrometheusUrl() string`

GetPrometheusUrl returns the PrometheusUrl field if non-nil, zero value otherwise.

### GetPrometheusUrlOk

`func (o *InternalProjectListDto) GetPrometheusUrlOk() (*string, bool)`

GetPrometheusUrlOk returns a tuple with the PrometheusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheusUrl

`func (o *InternalProjectListDto) SetPrometheusUrl(v string)`

SetPrometheusUrl sets PrometheusUrl field to given value.


### SetPrometheusUrlNil

`func (o *InternalProjectListDto) SetPrometheusUrlNil(b bool)`

 SetPrometheusUrlNil sets the value for PrometheusUrl to be an explicit nil

### UnsetPrometheusUrl
`func (o *InternalProjectListDto) UnsetPrometheusUrl()`

UnsetPrometheusUrl ensures that no value is present for PrometheusUrl, not even an explicit nil
### GetAlertManagerUrl

`func (o *InternalProjectListDto) GetAlertManagerUrl() string`

GetAlertManagerUrl returns the AlertManagerUrl field if non-nil, zero value otherwise.

### GetAlertManagerUrlOk

`func (o *InternalProjectListDto) GetAlertManagerUrlOk() (*string, bool)`

GetAlertManagerUrlOk returns a tuple with the AlertManagerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertManagerUrl

`func (o *InternalProjectListDto) SetAlertManagerUrl(v string)`

SetAlertManagerUrl sets AlertManagerUrl field to given value.


### SetAlertManagerUrlNil

`func (o *InternalProjectListDto) SetAlertManagerUrlNil(b bool)`

 SetAlertManagerUrlNil sets the value for AlertManagerUrl to be an explicit nil

### UnsetAlertManagerUrl
`func (o *InternalProjectListDto) UnsetAlertManagerUrl()`

UnsetAlertManagerUrl ensures that no value is present for AlertManagerUrl, not even an explicit nil
### GetLokiUrl

`func (o *InternalProjectListDto) GetLokiUrl() string`

GetLokiUrl returns the LokiUrl field if non-nil, zero value otherwise.

### GetLokiUrlOk

`func (o *InternalProjectListDto) GetLokiUrlOk() (*string, bool)`

GetLokiUrlOk returns a tuple with the LokiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLokiUrl

`func (o *InternalProjectListDto) SetLokiUrl(v string)`

SetLokiUrl sets LokiUrl field to given value.


### SetLokiUrlNil

`func (o *InternalProjectListDto) SetLokiUrlNil(b bool)`

 SetLokiUrlNil sets the value for LokiUrl to be an explicit nil

### UnsetLokiUrl
`func (o *InternalProjectListDto) UnsetLokiUrl()`

UnsetLokiUrl ensures that no value is present for LokiUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


