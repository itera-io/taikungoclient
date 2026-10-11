# ConfigMapSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CommonSearchKubernetesResponseData**](CommonSearchKubernetesResponseData.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewConfigMapSearchList

`func NewConfigMapSearchList(data []CommonSearchKubernetesResponseData, totalCount int32, ) *ConfigMapSearchList`

NewConfigMapSearchList instantiates a new ConfigMapSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMapSearchListWithDefaults

`func NewConfigMapSearchListWithDefaults() *ConfigMapSearchList`

NewConfigMapSearchListWithDefaults instantiates a new ConfigMapSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ConfigMapSearchList) GetData() []CommonSearchKubernetesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigMapSearchList) GetDataOk() (*[]CommonSearchKubernetesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigMapSearchList) SetData(v []CommonSearchKubernetesResponseData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ConfigMapSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ConfigMapSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *ConfigMapSearchList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ConfigMapSearchList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ConfigMapSearchList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


