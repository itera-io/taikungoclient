# KubernetesProfilesSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CommonSearchResponseData**](CommonSearchResponseData.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewKubernetesProfilesSearchList

`func NewKubernetesProfilesSearchList(data []CommonSearchResponseData, totalCount int32, ) *KubernetesProfilesSearchList`

NewKubernetesProfilesSearchList instantiates a new KubernetesProfilesSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProfilesSearchListWithDefaults

`func NewKubernetesProfilesSearchListWithDefaults() *KubernetesProfilesSearchList`

NewKubernetesProfilesSearchListWithDefaults instantiates a new KubernetesProfilesSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *KubernetesProfilesSearchList) GetData() []CommonSearchResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KubernetesProfilesSearchList) GetDataOk() (*[]CommonSearchResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KubernetesProfilesSearchList) SetData(v []CommonSearchResponseData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *KubernetesProfilesSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *KubernetesProfilesSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *KubernetesProfilesSearchList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KubernetesProfilesSearchList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KubernetesProfilesSearchList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


