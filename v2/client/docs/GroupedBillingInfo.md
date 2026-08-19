# GroupedBillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GroupedBillings**](GroupedBillings.md) |  | [optional] 
**ProjectId** | Pointer to **int32** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGroupedBillingInfo

`func NewGroupedBillingInfo() *GroupedBillingInfo`

NewGroupedBillingInfo instantiates a new GroupedBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupedBillingInfoWithDefaults

`func NewGroupedBillingInfoWithDefaults() *GroupedBillingInfo`

NewGroupedBillingInfoWithDefaults instantiates a new GroupedBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GroupedBillingInfo) GetData() []GroupedBillings`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupedBillingInfo) GetDataOk() (*[]GroupedBillings, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupedBillingInfo) SetData(v []GroupedBillings)`

SetData sets Data field to given value.

### HasData

`func (o *GroupedBillingInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GroupedBillingInfo) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GroupedBillingInfo) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetProjectId

`func (o *GroupedBillingInfo) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GroupedBillingInfo) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GroupedBillingInfo) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GroupedBillingInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *GroupedBillingInfo) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *GroupedBillingInfo) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *GroupedBillingInfo) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *GroupedBillingInfo) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *GroupedBillingInfo) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *GroupedBillingInfo) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


