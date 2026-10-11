# AzureQuotaListRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCores** | **int64** |  | 
**CurrentUsage** | **int32** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewAzureQuotaListRecordDto

`func NewAzureQuotaListRecordDto(totalCores int64, currentUsage int32, name NullableString, ) *AzureQuotaListRecordDto`

NewAzureQuotaListRecordDto instantiates a new AzureQuotaListRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureQuotaListRecordDtoWithDefaults

`func NewAzureQuotaListRecordDtoWithDefaults() *AzureQuotaListRecordDto`

NewAzureQuotaListRecordDtoWithDefaults instantiates a new AzureQuotaListRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCores

`func (o *AzureQuotaListRecordDto) GetTotalCores() int64`

GetTotalCores returns the TotalCores field if non-nil, zero value otherwise.

### GetTotalCoresOk

`func (o *AzureQuotaListRecordDto) GetTotalCoresOk() (*int64, bool)`

GetTotalCoresOk returns a tuple with the TotalCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCores

`func (o *AzureQuotaListRecordDto) SetTotalCores(v int64)`

SetTotalCores sets TotalCores field to given value.


### GetCurrentUsage

`func (o *AzureQuotaListRecordDto) GetCurrentUsage() int32`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *AzureQuotaListRecordDto) GetCurrentUsageOk() (*int32, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *AzureQuotaListRecordDto) SetCurrentUsage(v int32)`

SetCurrentUsage sets CurrentUsage field to given value.


### GetName

`func (o *AzureQuotaListRecordDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureQuotaListRecordDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureQuotaListRecordDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AzureQuotaListRecordDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureQuotaListRecordDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


