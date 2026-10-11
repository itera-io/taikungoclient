# KubernetesActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | **[]string** |  | 
**Failed** | **[]string** |  | 

## Methods

### NewKubernetesActionResponse

`func NewKubernetesActionResponse(succeeded []string, failed []string, ) *KubernetesActionResponse`

NewKubernetesActionResponse instantiates a new KubernetesActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesActionResponseWithDefaults

`func NewKubernetesActionResponseWithDefaults() *KubernetesActionResponse`

NewKubernetesActionResponseWithDefaults instantiates a new KubernetesActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *KubernetesActionResponse) GetSucceeded() []string`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *KubernetesActionResponse) GetSucceededOk() (*[]string, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *KubernetesActionResponse) SetSucceeded(v []string)`

SetSucceeded sets Succeeded field to given value.


### SetSucceededNil

`func (o *KubernetesActionResponse) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *KubernetesActionResponse) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetFailed

`func (o *KubernetesActionResponse) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *KubernetesActionResponse) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *KubernetesActionResponse) SetFailed(v []string)`

SetFailed sets Failed field to given value.


### SetFailedNil

`func (o *KubernetesActionResponse) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *KubernetesActionResponse) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


