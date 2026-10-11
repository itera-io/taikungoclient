# ProjectChartDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**Updating** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**Purging** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**Deleting** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**Importing** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 
**FailedToImport** | Pointer to [**[]ProjectCommonRecordDto**](ProjectCommonRecordDto.md) |  | [optional] 

## Methods

### NewProjectChartDto

`func NewProjectChartDto() *ProjectChartDto`

NewProjectChartDto instantiates a new ProjectChartDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectChartDtoWithDefaults

`func NewProjectChartDtoWithDefaults() *ProjectChartDto`

NewProjectChartDtoWithDefaults instantiates a new ProjectChartDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *ProjectChartDto) GetSucceeded() []ProjectCommonRecordDto`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *ProjectChartDto) GetSucceededOk() (*[]ProjectCommonRecordDto, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *ProjectChartDto) SetSucceeded(v []ProjectCommonRecordDto)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *ProjectChartDto) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *ProjectChartDto) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *ProjectChartDto) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetUpdating

`func (o *ProjectChartDto) GetUpdating() []ProjectCommonRecordDto`

GetUpdating returns the Updating field if non-nil, zero value otherwise.

### GetUpdatingOk

`func (o *ProjectChartDto) GetUpdatingOk() (*[]ProjectCommonRecordDto, bool)`

GetUpdatingOk returns a tuple with the Updating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdating

`func (o *ProjectChartDto) SetUpdating(v []ProjectCommonRecordDto)`

SetUpdating sets Updating field to given value.

### HasUpdating

`func (o *ProjectChartDto) HasUpdating() bool`

HasUpdating returns a boolean if a field has been set.

### SetUpdatingNil

`func (o *ProjectChartDto) SetUpdatingNil(b bool)`

 SetUpdatingNil sets the value for Updating to be an explicit nil

### UnsetUpdating
`func (o *ProjectChartDto) UnsetUpdating()`

UnsetUpdating ensures that no value is present for Updating, not even an explicit nil
### GetTotalCount

`func (o *ProjectChartDto) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectChartDto) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectChartDto) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProjectChartDto) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetFailed

`func (o *ProjectChartDto) GetFailed() []ProjectCommonRecordDto`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProjectChartDto) GetFailedOk() (*[]ProjectCommonRecordDto, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProjectChartDto) SetFailed(v []ProjectCommonRecordDto)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ProjectChartDto) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *ProjectChartDto) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *ProjectChartDto) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetPurging

`func (o *ProjectChartDto) GetPurging() []ProjectCommonRecordDto`

GetPurging returns the Purging field if non-nil, zero value otherwise.

### GetPurgingOk

`func (o *ProjectChartDto) GetPurgingOk() (*[]ProjectCommonRecordDto, bool)`

GetPurgingOk returns a tuple with the Purging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurging

`func (o *ProjectChartDto) SetPurging(v []ProjectCommonRecordDto)`

SetPurging sets Purging field to given value.

### HasPurging

`func (o *ProjectChartDto) HasPurging() bool`

HasPurging returns a boolean if a field has been set.

### SetPurgingNil

`func (o *ProjectChartDto) SetPurgingNil(b bool)`

 SetPurgingNil sets the value for Purging to be an explicit nil

### UnsetPurging
`func (o *ProjectChartDto) UnsetPurging()`

UnsetPurging ensures that no value is present for Purging, not even an explicit nil
### GetDeleting

`func (o *ProjectChartDto) GetDeleting() []ProjectCommonRecordDto`

GetDeleting returns the Deleting field if non-nil, zero value otherwise.

### GetDeletingOk

`func (o *ProjectChartDto) GetDeletingOk() (*[]ProjectCommonRecordDto, bool)`

GetDeletingOk returns a tuple with the Deleting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleting

`func (o *ProjectChartDto) SetDeleting(v []ProjectCommonRecordDto)`

SetDeleting sets Deleting field to given value.

### HasDeleting

`func (o *ProjectChartDto) HasDeleting() bool`

HasDeleting returns a boolean if a field has been set.

### SetDeletingNil

`func (o *ProjectChartDto) SetDeletingNil(b bool)`

 SetDeletingNil sets the value for Deleting to be an explicit nil

### UnsetDeleting
`func (o *ProjectChartDto) UnsetDeleting()`

UnsetDeleting ensures that no value is present for Deleting, not even an explicit nil
### GetImporting

`func (o *ProjectChartDto) GetImporting() []ProjectCommonRecordDto`

GetImporting returns the Importing field if non-nil, zero value otherwise.

### GetImportingOk

`func (o *ProjectChartDto) GetImportingOk() (*[]ProjectCommonRecordDto, bool)`

GetImportingOk returns a tuple with the Importing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporting

`func (o *ProjectChartDto) SetImporting(v []ProjectCommonRecordDto)`

SetImporting sets Importing field to given value.

### HasImporting

`func (o *ProjectChartDto) HasImporting() bool`

HasImporting returns a boolean if a field has been set.

### SetImportingNil

`func (o *ProjectChartDto) SetImportingNil(b bool)`

 SetImportingNil sets the value for Importing to be an explicit nil

### UnsetImporting
`func (o *ProjectChartDto) UnsetImporting()`

UnsetImporting ensures that no value is present for Importing, not even an explicit nil
### GetFailedToImport

`func (o *ProjectChartDto) GetFailedToImport() []ProjectCommonRecordDto`

GetFailedToImport returns the FailedToImport field if non-nil, zero value otherwise.

### GetFailedToImportOk

`func (o *ProjectChartDto) GetFailedToImportOk() (*[]ProjectCommonRecordDto, bool)`

GetFailedToImportOk returns a tuple with the FailedToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToImport

`func (o *ProjectChartDto) SetFailedToImport(v []ProjectCommonRecordDto)`

SetFailedToImport sets FailedToImport field to given value.

### HasFailedToImport

`func (o *ProjectChartDto) HasFailedToImport() bool`

HasFailedToImport returns a boolean if a field has been set.

### SetFailedToImportNil

`func (o *ProjectChartDto) SetFailedToImportNil(b bool)`

 SetFailedToImportNil sets the value for FailedToImport to be an explicit nil

### UnsetFailedToImport
`func (o *ProjectChartDto) UnsetFailedToImport()`

UnsetFailedToImport ensures that no value is present for FailedToImport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


