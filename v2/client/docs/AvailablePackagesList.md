# AvailablePackagesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AvailablePackagesDto**](AvailablePackagesDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAvailablePackagesList

`func NewAvailablePackagesList(data []AvailablePackagesDto, totalCount int32, ) *AvailablePackagesList`

NewAvailablePackagesList instantiates a new AvailablePackagesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePackagesListWithDefaults

`func NewAvailablePackagesListWithDefaults() *AvailablePackagesList`

NewAvailablePackagesListWithDefaults instantiates a new AvailablePackagesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AvailablePackagesList) GetData() []AvailablePackagesDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AvailablePackagesList) GetDataOk() (*[]AvailablePackagesDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AvailablePackagesList) SetData(v []AvailablePackagesDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *AvailablePackagesList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AvailablePackagesList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AvailablePackagesList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AvailablePackagesList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AvailablePackagesList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


