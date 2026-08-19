# AlertData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]Group**](Group.md) |  | 

## Methods

### NewAlertData

`func NewAlertData(groups []Group, ) *AlertData`

NewAlertData instantiates a new AlertData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertDataWithDefaults

`func NewAlertDataWithDefaults() *AlertData`

NewAlertDataWithDefaults instantiates a new AlertData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AlertData) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AlertData) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AlertData) SetGroups(v []Group)`

SetGroups sets Groups field to given value.


### SetGroupsNil

`func (o *AlertData) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *AlertData) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


