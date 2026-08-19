# KubeConfigRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewKubeConfigRoleResponse

`func NewKubeConfigRoleResponse(data []CommonDropdownDto, totalCount int32, ) *KubeConfigRoleResponse`

NewKubeConfigRoleResponse instantiates a new KubeConfigRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigRoleResponseWithDefaults

`func NewKubeConfigRoleResponseWithDefaults() *KubeConfigRoleResponse`

NewKubeConfigRoleResponseWithDefaults instantiates a new KubeConfigRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *KubeConfigRoleResponse) GetData() []CommonDropdownDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KubeConfigRoleResponse) GetDataOk() (*[]CommonDropdownDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KubeConfigRoleResponse) SetData(v []CommonDropdownDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *KubeConfigRoleResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *KubeConfigRoleResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *KubeConfigRoleResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KubeConfigRoleResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KubeConfigRoleResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


