# KubernetesDashboardDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pods** | [**[]PodDto**](PodDto.md) |  | 
**Nodes** | [**[]NodeDto**](NodeDto.md) |  | 

## Methods

### NewKubernetesDashboardDto

`func NewKubernetesDashboardDto(pods []PodDto, nodes []NodeDto, ) *KubernetesDashboardDto`

NewKubernetesDashboardDto instantiates a new KubernetesDashboardDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDashboardDtoWithDefaults

`func NewKubernetesDashboardDtoWithDefaults() *KubernetesDashboardDto`

NewKubernetesDashboardDtoWithDefaults instantiates a new KubernetesDashboardDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPods

`func (o *KubernetesDashboardDto) GetPods() []PodDto`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *KubernetesDashboardDto) GetPodsOk() (*[]PodDto, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *KubernetesDashboardDto) SetPods(v []PodDto)`

SetPods sets Pods field to given value.


### SetPodsNil

`func (o *KubernetesDashboardDto) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *KubernetesDashboardDto) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetNodes

`func (o *KubernetesDashboardDto) GetNodes() []NodeDto`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *KubernetesDashboardDto) GetNodesOk() (*[]NodeDto, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *KubernetesDashboardDto) SetNodes(v []NodeDto)`

SetNodes sets Nodes field to given value.


### SetNodesNil

`func (o *KubernetesDashboardDto) SetNodesNil(b bool)`

 SetNodesNil sets the value for Nodes to be an explicit nil

### UnsetNodes
`func (o *KubernetesDashboardDto) UnsetNodes()`

UnsetNodes ensures that no value is present for Nodes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


