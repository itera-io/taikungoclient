# KubernetesOverviewDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**ProjectName** | **string** |  | 
**HealthyPods** | **int32** |  | 
**UnhealthyPods** | **int32** |  | 
**HealthyNodes** | **int32** |  | 
**UnhealthyNodes** | **int32** |  | 
**AlertsCount** | **int32** |  | 
**KubernetesHealth** | **NullableString** |  | 

## Methods

### NewKubernetesOverviewDto

`func NewKubernetesOverviewDto(projectId int32, projectName string, healthyPods int32, unhealthyPods int32, healthyNodes int32, unhealthyNodes int32, alertsCount int32, kubernetesHealth NullableString, ) *KubernetesOverviewDto`

NewKubernetesOverviewDto instantiates a new KubernetesOverviewDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesOverviewDtoWithDefaults

`func NewKubernetesOverviewDtoWithDefaults() *KubernetesOverviewDto`

NewKubernetesOverviewDtoWithDefaults instantiates a new KubernetesOverviewDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *KubernetesOverviewDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KubernetesOverviewDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KubernetesOverviewDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *KubernetesOverviewDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *KubernetesOverviewDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *KubernetesOverviewDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetHealthyPods

`func (o *KubernetesOverviewDto) GetHealthyPods() int32`

GetHealthyPods returns the HealthyPods field if non-nil, zero value otherwise.

### GetHealthyPodsOk

`func (o *KubernetesOverviewDto) GetHealthyPodsOk() (*int32, bool)`

GetHealthyPodsOk returns a tuple with the HealthyPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyPods

`func (o *KubernetesOverviewDto) SetHealthyPods(v int32)`

SetHealthyPods sets HealthyPods field to given value.


### GetUnhealthyPods

`func (o *KubernetesOverviewDto) GetUnhealthyPods() int32`

GetUnhealthyPods returns the UnhealthyPods field if non-nil, zero value otherwise.

### GetUnhealthyPodsOk

`func (o *KubernetesOverviewDto) GetUnhealthyPodsOk() (*int32, bool)`

GetUnhealthyPodsOk returns a tuple with the UnhealthyPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyPods

`func (o *KubernetesOverviewDto) SetUnhealthyPods(v int32)`

SetUnhealthyPods sets UnhealthyPods field to given value.


### GetHealthyNodes

`func (o *KubernetesOverviewDto) GetHealthyNodes() int32`

GetHealthyNodes returns the HealthyNodes field if non-nil, zero value otherwise.

### GetHealthyNodesOk

`func (o *KubernetesOverviewDto) GetHealthyNodesOk() (*int32, bool)`

GetHealthyNodesOk returns a tuple with the HealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyNodes

`func (o *KubernetesOverviewDto) SetHealthyNodes(v int32)`

SetHealthyNodes sets HealthyNodes field to given value.


### GetUnhealthyNodes

`func (o *KubernetesOverviewDto) GetUnhealthyNodes() int32`

GetUnhealthyNodes returns the UnhealthyNodes field if non-nil, zero value otherwise.

### GetUnhealthyNodesOk

`func (o *KubernetesOverviewDto) GetUnhealthyNodesOk() (*int32, bool)`

GetUnhealthyNodesOk returns a tuple with the UnhealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyNodes

`func (o *KubernetesOverviewDto) SetUnhealthyNodes(v int32)`

SetUnhealthyNodes sets UnhealthyNodes field to given value.


### GetAlertsCount

`func (o *KubernetesOverviewDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *KubernetesOverviewDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *KubernetesOverviewDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetKubernetesHealth

`func (o *KubernetesOverviewDto) GetKubernetesHealth() string`

GetKubernetesHealth returns the KubernetesHealth field if non-nil, zero value otherwise.

### GetKubernetesHealthOk

`func (o *KubernetesOverviewDto) GetKubernetesHealthOk() (*string, bool)`

GetKubernetesHealthOk returns a tuple with the KubernetesHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesHealth

`func (o *KubernetesOverviewDto) SetKubernetesHealth(v string)`

SetKubernetesHealth sets KubernetesHealth field to given value.


### SetKubernetesHealthNil

`func (o *KubernetesOverviewDto) SetKubernetesHealthNil(b bool)`

 SetKubernetesHealthNil sets the value for KubernetesHealth to be an explicit nil

### UnsetKubernetesHealth
`func (o *KubernetesOverviewDto) UnsetKubernetesHealth()`

UnsetKubernetesHealth ensures that no value is present for KubernetesHealth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


