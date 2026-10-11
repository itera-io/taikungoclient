# HelmReleasesSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CommonSearchKubernetesResponseData**](CommonSearchKubernetesResponseData.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewHelmReleasesSearchList

`func NewHelmReleasesSearchList(data []CommonSearchKubernetesResponseData, totalCount int32, ) *HelmReleasesSearchList`

NewHelmReleasesSearchList instantiates a new HelmReleasesSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmReleasesSearchListWithDefaults

`func NewHelmReleasesSearchListWithDefaults() *HelmReleasesSearchList`

NewHelmReleasesSearchListWithDefaults instantiates a new HelmReleasesSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *HelmReleasesSearchList) GetData() []CommonSearchKubernetesResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HelmReleasesSearchList) GetDataOk() (*[]CommonSearchKubernetesResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HelmReleasesSearchList) SetData(v []CommonSearchKubernetesResponseData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *HelmReleasesSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *HelmReleasesSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *HelmReleasesSearchList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *HelmReleasesSearchList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *HelmReleasesSearchList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


