# KubeConfigForUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]KubeConfigForUserDto**](KubeConfigForUserDto.md) |  | 
**TotalCount** | **int32** |  | 
**CanAdd** | [**ButtonStatusDto**](ButtonStatusDto.md) |  | 

## Methods

### NewKubeConfigForUserList

`func NewKubeConfigForUserList(data []KubeConfigForUserDto, totalCount int32, canAdd ButtonStatusDto, ) *KubeConfigForUserList`

NewKubeConfigForUserList instantiates a new KubeConfigForUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigForUserListWithDefaults

`func NewKubeConfigForUserListWithDefaults() *KubeConfigForUserList`

NewKubeConfigForUserListWithDefaults instantiates a new KubeConfigForUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *KubeConfigForUserList) GetData() []KubeConfigForUserDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KubeConfigForUserList) GetDataOk() (*[]KubeConfigForUserDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KubeConfigForUserList) SetData(v []KubeConfigForUserDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *KubeConfigForUserList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *KubeConfigForUserList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *KubeConfigForUserList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KubeConfigForUserList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KubeConfigForUserList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetCanAdd

`func (o *KubeConfigForUserList) GetCanAdd() ButtonStatusDto`

GetCanAdd returns the CanAdd field if non-nil, zero value otherwise.

### GetCanAddOk

`func (o *KubeConfigForUserList) GetCanAddOk() (*ButtonStatusDto, bool)`

GetCanAddOk returns a tuple with the CanAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdd

`func (o *KubeConfigForUserList) SetCanAdd(v ButtonStatusDto)`

SetCanAdd sets CanAdd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


