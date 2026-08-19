# AlertLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alertname** | **NullableString** |  | 
**Condition** | **NullableString** |  | 
**Container** | **NullableString** |  | 
**Endpoint** | **NullableString** |  | 
**Instance** | **NullableString** |  | 
**Job** | **NullableString** |  | 
**Namespace** | **NullableString** |  | 
**Node** | **NullableString** |  | 
**Pod** | **NullableString** |  | 
**Service** | **NullableString** |  | 
**Severity** | **NullableString** |  | 
**Status** | **NullableString** |  | 
**Daemonset** | **NullableString** |  | 

## Methods

### NewAlertLabels

`func NewAlertLabels(alertname NullableString, condition NullableString, container NullableString, endpoint NullableString, instance NullableString, job NullableString, namespace NullableString, node NullableString, pod NullableString, service NullableString, severity NullableString, status NullableString, daemonset NullableString, ) *AlertLabels`

NewAlertLabels instantiates a new AlertLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertLabelsWithDefaults

`func NewAlertLabelsWithDefaults() *AlertLabels`

NewAlertLabelsWithDefaults instantiates a new AlertLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertname

`func (o *AlertLabels) GetAlertname() string`

GetAlertname returns the Alertname field if non-nil, zero value otherwise.

### GetAlertnameOk

`func (o *AlertLabels) GetAlertnameOk() (*string, bool)`

GetAlertnameOk returns a tuple with the Alertname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertname

`func (o *AlertLabels) SetAlertname(v string)`

SetAlertname sets Alertname field to given value.


### SetAlertnameNil

`func (o *AlertLabels) SetAlertnameNil(b bool)`

 SetAlertnameNil sets the value for Alertname to be an explicit nil

### UnsetAlertname
`func (o *AlertLabels) UnsetAlertname()`

UnsetAlertname ensures that no value is present for Alertname, not even an explicit nil
### GetCondition

`func (o *AlertLabels) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AlertLabels) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AlertLabels) SetCondition(v string)`

SetCondition sets Condition field to given value.


### SetConditionNil

`func (o *AlertLabels) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *AlertLabels) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetContainer

`func (o *AlertLabels) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AlertLabels) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AlertLabels) SetContainer(v string)`

SetContainer sets Container field to given value.


### SetContainerNil

`func (o *AlertLabels) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *AlertLabels) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetEndpoint

`func (o *AlertLabels) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AlertLabels) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AlertLabels) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### SetEndpointNil

`func (o *AlertLabels) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *AlertLabels) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetInstance

`func (o *AlertLabels) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AlertLabels) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AlertLabels) SetInstance(v string)`

SetInstance sets Instance field to given value.


### SetInstanceNil

`func (o *AlertLabels) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *AlertLabels) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetJob

`func (o *AlertLabels) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *AlertLabels) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *AlertLabels) SetJob(v string)`

SetJob sets Job field to given value.


### SetJobNil

`func (o *AlertLabels) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *AlertLabels) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetNamespace

`func (o *AlertLabels) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AlertLabels) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AlertLabels) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *AlertLabels) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *AlertLabels) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetNode

`func (o *AlertLabels) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *AlertLabels) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *AlertLabels) SetNode(v string)`

SetNode sets Node field to given value.


### SetNodeNil

`func (o *AlertLabels) SetNodeNil(b bool)`

 SetNodeNil sets the value for Node to be an explicit nil

### UnsetNode
`func (o *AlertLabels) UnsetNode()`

UnsetNode ensures that no value is present for Node, not even an explicit nil
### GetPod

`func (o *AlertLabels) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *AlertLabels) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *AlertLabels) SetPod(v string)`

SetPod sets Pod field to given value.


### SetPodNil

`func (o *AlertLabels) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *AlertLabels) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetService

`func (o *AlertLabels) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AlertLabels) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AlertLabels) SetService(v string)`

SetService sets Service field to given value.


### SetServiceNil

`func (o *AlertLabels) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *AlertLabels) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetSeverity

`func (o *AlertLabels) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertLabels) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertLabels) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### SetSeverityNil

`func (o *AlertLabels) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *AlertLabels) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
### GetStatus

`func (o *AlertLabels) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertLabels) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertLabels) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *AlertLabels) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AlertLabels) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDaemonset

`func (o *AlertLabels) GetDaemonset() string`

GetDaemonset returns the Daemonset field if non-nil, zero value otherwise.

### GetDaemonsetOk

`func (o *AlertLabels) GetDaemonsetOk() (*string, bool)`

GetDaemonsetOk returns a tuple with the Daemonset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonset

`func (o *AlertLabels) SetDaemonset(v string)`

SetDaemonset sets Daemonset field to given value.


### SetDaemonsetNil

`func (o *AlertLabels) SetDaemonsetNil(b bool)`

 SetDaemonsetNil sets the value for Daemonset to be an explicit nil

### UnsetDaemonset
`func (o *AlertLabels) UnsetDaemonset()`

UnsetDaemonset ensures that no value is present for Daemonset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


