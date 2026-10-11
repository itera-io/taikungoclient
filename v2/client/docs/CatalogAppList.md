# CatalogAppList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CatalogAppListDto**](CatalogAppListDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewCatalogAppList

`func NewCatalogAppList(data []CatalogAppListDto, totalCount int32, ) *CatalogAppList`

NewCatalogAppList instantiates a new CatalogAppList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogAppListWithDefaults

`func NewCatalogAppListWithDefaults() *CatalogAppList`

NewCatalogAppListWithDefaults instantiates a new CatalogAppList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CatalogAppList) GetData() []CatalogAppListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CatalogAppList) GetDataOk() (*[]CatalogAppListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CatalogAppList) SetData(v []CatalogAppListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *CatalogAppList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CatalogAppList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *CatalogAppList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CatalogAppList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CatalogAppList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


