# ProjectMonitoringAlertsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **NullableString** |  | 
**Data** | [**AlertData**](AlertData.md) |  | 

## Methods

### NewProjectMonitoringAlertsDto

`func NewProjectMonitoringAlertsDto(status NullableString, data AlertData, ) *ProjectMonitoringAlertsDto`

NewProjectMonitoringAlertsDto instantiates a new ProjectMonitoringAlertsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMonitoringAlertsDtoWithDefaults

`func NewProjectMonitoringAlertsDtoWithDefaults() *ProjectMonitoringAlertsDto`

NewProjectMonitoringAlertsDtoWithDefaults instantiates a new ProjectMonitoringAlertsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectMonitoringAlertsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectMonitoringAlertsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectMonitoringAlertsDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *ProjectMonitoringAlertsDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProjectMonitoringAlertsDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetData

`func (o *ProjectMonitoringAlertsDto) GetData() AlertData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProjectMonitoringAlertsDto) GetDataOk() (*AlertData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProjectMonitoringAlertsDto) SetData(v AlertData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


