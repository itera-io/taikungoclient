# OrganizationSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]OrganizationsSearchResponseData**](OrganizationsSearchResponseData.md) |  | 
**TotalCount** | **int64** |  | 

## Methods

### NewOrganizationSearchList

`func NewOrganizationSearchList(data []OrganizationsSearchResponseData, totalCount int64, ) *OrganizationSearchList`

NewOrganizationSearchList instantiates a new OrganizationSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSearchListWithDefaults

`func NewOrganizationSearchListWithDefaults() *OrganizationSearchList`

NewOrganizationSearchListWithDefaults instantiates a new OrganizationSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OrganizationSearchList) GetData() []OrganizationsSearchResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationSearchList) GetDataOk() (*[]OrganizationsSearchResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationSearchList) SetData(v []OrganizationsSearchResponseData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *OrganizationSearchList) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *OrganizationSearchList) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *OrganizationSearchList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrganizationSearchList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrganizationSearchList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


