# AllFlavorsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FlavorsListDto**](FlavorsListDto.md) |  | 
**TotalCount** | **int32** |  | 
**CloudType** | **NullableString** |  | 

## Methods

### NewAllFlavorsList

`func NewAllFlavorsList(data []FlavorsListDto, totalCount int32, cloudType NullableString, ) *AllFlavorsList`

NewAllFlavorsList instantiates a new AllFlavorsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllFlavorsListWithDefaults

`func NewAllFlavorsListWithDefaults() *AllFlavorsList`

NewAllFlavorsListWithDefaults instantiates a new AllFlavorsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AllFlavorsList) GetData() []FlavorsListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AllFlavorsList) GetDataOk() (*[]FlavorsListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AllFlavorsList) SetData(v []FlavorsListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *AllFlavorsList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AllFlavorsList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AllFlavorsList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AllFlavorsList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AllFlavorsList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetCloudType

`func (o *AllFlavorsList) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *AllFlavorsList) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *AllFlavorsList) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.


### SetCloudTypeNil

`func (o *AllFlavorsList) SetCloudTypeNil(b bool)`

 SetCloudTypeNil sets the value for CloudType to be an explicit nil

### UnsetCloudType
`func (o *AllFlavorsList) UnsetCloudType()`

UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


