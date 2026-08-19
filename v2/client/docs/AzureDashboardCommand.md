# AzureDashboardCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | **int32** |  | 
**FilterBy** | **NullableString** |  | 

## Methods

### NewAzureDashboardCommand

`func NewAzureDashboardCommand(cloudId int32, filterBy NullableString, ) *AzureDashboardCommand`

NewAzureDashboardCommand instantiates a new AzureDashboardCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDashboardCommandWithDefaults

`func NewAzureDashboardCommandWithDefaults() *AzureDashboardCommand`

NewAzureDashboardCommandWithDefaults instantiates a new AzureDashboardCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *AzureDashboardCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AzureDashboardCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AzureDashboardCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### GetFilterBy

`func (o *AzureDashboardCommand) GetFilterBy() string`

GetFilterBy returns the FilterBy field if non-nil, zero value otherwise.

### GetFilterByOk

`func (o *AzureDashboardCommand) GetFilterByOk() (*string, bool)`

GetFilterByOk returns a tuple with the FilterBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBy

`func (o *AzureDashboardCommand) SetFilterBy(v string)`

SetFilterBy sets FilterBy field to given value.


### SetFilterByNil

`func (o *AzureDashboardCommand) SetFilterByNil(b bool)`

 SetFilterByNil sets the value for FilterBy to be an explicit nil

### UnsetFilterBy
`func (o *AzureDashboardCommand) UnsetFilterBy()`

UnsetFilterBy ensures that no value is present for FilterBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


