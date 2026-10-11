# VClusterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VClusterListDto**](VClusterListDto.md) |  | 
**TotalCount** | **int64** |  | 
**Project** | [**ProjectDetailsForVmsDto**](ProjectDetailsForVmsDto.md) |  | 

## Methods

### NewVClusterList

`func NewVClusterList(data []VClusterListDto, totalCount int64, project ProjectDetailsForVmsDto, ) *VClusterList`

NewVClusterList instantiates a new VClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVClusterListWithDefaults

`func NewVClusterListWithDefaults() *VClusterList`

NewVClusterListWithDefaults instantiates a new VClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VClusterList) GetData() []VClusterListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VClusterList) GetDataOk() (*[]VClusterListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VClusterList) SetData(v []VClusterListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *VClusterList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *VClusterList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *VClusterList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *VClusterList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *VClusterList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetProject

`func (o *VClusterList) GetProject() ProjectDetailsForVmsDto`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VClusterList) GetProjectOk() (*ProjectDetailsForVmsDto, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VClusterList) SetProject(v ProjectDetailsForVmsDto)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


