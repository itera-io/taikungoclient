# NodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataName** | **map[string]interface{}** |  | 
**KubeletReady** | **map[string]interface{}** |  | 
**KubeletSufficient** | **map[string]interface{}** |  | 
**KubeletDiskPressure** | **map[string]interface{}** |  | 
**KubeletMemory** | **map[string]interface{}** |  | 

## Methods

### NewNodeDto

`func NewNodeDto(metadataName map[string]interface{}, kubeletReady map[string]interface{}, kubeletSufficient map[string]interface{}, kubeletDiskPressure map[string]interface{}, kubeletMemory map[string]interface{}, ) *NodeDto`

NewNodeDto instantiates a new NodeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDtoWithDefaults

`func NewNodeDtoWithDefaults() *NodeDto`

NewNodeDtoWithDefaults instantiates a new NodeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataName

`func (o *NodeDto) GetMetadataName() map[string]interface{}`

GetMetadataName returns the MetadataName field if non-nil, zero value otherwise.

### GetMetadataNameOk

`func (o *NodeDto) GetMetadataNameOk() (*map[string]interface{}, bool)`

GetMetadataNameOk returns a tuple with the MetadataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataName

`func (o *NodeDto) SetMetadataName(v map[string]interface{})`

SetMetadataName sets MetadataName field to given value.


### GetKubeletReady

`func (o *NodeDto) GetKubeletReady() map[string]interface{}`

GetKubeletReady returns the KubeletReady field if non-nil, zero value otherwise.

### GetKubeletReadyOk

`func (o *NodeDto) GetKubeletReadyOk() (*map[string]interface{}, bool)`

GetKubeletReadyOk returns a tuple with the KubeletReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletReady

`func (o *NodeDto) SetKubeletReady(v map[string]interface{})`

SetKubeletReady sets KubeletReady field to given value.


### GetKubeletSufficient

`func (o *NodeDto) GetKubeletSufficient() map[string]interface{}`

GetKubeletSufficient returns the KubeletSufficient field if non-nil, zero value otherwise.

### GetKubeletSufficientOk

`func (o *NodeDto) GetKubeletSufficientOk() (*map[string]interface{}, bool)`

GetKubeletSufficientOk returns a tuple with the KubeletSufficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletSufficient

`func (o *NodeDto) SetKubeletSufficient(v map[string]interface{})`

SetKubeletSufficient sets KubeletSufficient field to given value.


### GetKubeletDiskPressure

`func (o *NodeDto) GetKubeletDiskPressure() map[string]interface{}`

GetKubeletDiskPressure returns the KubeletDiskPressure field if non-nil, zero value otherwise.

### GetKubeletDiskPressureOk

`func (o *NodeDto) GetKubeletDiskPressureOk() (*map[string]interface{}, bool)`

GetKubeletDiskPressureOk returns a tuple with the KubeletDiskPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletDiskPressure

`func (o *NodeDto) SetKubeletDiskPressure(v map[string]interface{})`

SetKubeletDiskPressure sets KubeletDiskPressure field to given value.


### GetKubeletMemory

`func (o *NodeDto) GetKubeletMemory() map[string]interface{}`

GetKubeletMemory returns the KubeletMemory field if non-nil, zero value otherwise.

### GetKubeletMemoryOk

`func (o *NodeDto) GetKubeletMemoryOk() (*map[string]interface{}, bool)`

GetKubeletMemoryOk returns a tuple with the KubeletMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletMemory

`func (o *NodeDto) SetKubeletMemory(v map[string]interface{})`

SetKubeletMemory sets KubeletMemory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


