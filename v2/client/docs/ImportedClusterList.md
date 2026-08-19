# ImportedClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ImportedClusterListDto**](ImportedClusterListDto.md) |  | 
**Project** | [**ImportedClusterDetailsDto**](ImportedClusterDetailsDto.md) |  | 
**TotalCount** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewImportedClusterList

`func NewImportedClusterList(data []ImportedClusterListDto, project ImportedClusterDetailsDto, ) *ImportedClusterList`

NewImportedClusterList instantiates a new ImportedClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedClusterListWithDefaults

`func NewImportedClusterListWithDefaults() *ImportedClusterList`

NewImportedClusterListWithDefaults instantiates a new ImportedClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ImportedClusterList) GetData() []ImportedClusterListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ImportedClusterList) GetDataOk() (*[]ImportedClusterListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ImportedClusterList) SetData(v []ImportedClusterListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *ImportedClusterList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ImportedClusterList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetProject

`func (o *ImportedClusterList) GetProject() ImportedClusterDetailsDto`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ImportedClusterList) GetProjectOk() (*ImportedClusterDetailsDto, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ImportedClusterList) SetProject(v ImportedClusterDetailsDto)`

SetProject sets Project field to given value.


### GetTotalCount

`func (o *ImportedClusterList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ImportedClusterList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ImportedClusterList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ImportedClusterList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *ImportedClusterList) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *ImportedClusterList) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


