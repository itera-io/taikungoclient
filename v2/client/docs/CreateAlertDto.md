# CreateAlertDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]KubernetesAlertCreateDto**](KubernetesAlertCreateDto.md) |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateAlertDto

`func NewCreateAlertDto() *CreateAlertDto`

NewCreateAlertDto instantiates a new CreateAlertDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertDtoWithDefaults

`func NewCreateAlertDtoWithDefaults() *CreateAlertDto`

NewCreateAlertDtoWithDefaults instantiates a new CreateAlertDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *CreateAlertDto) GetAlerts() []KubernetesAlertCreateDto`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *CreateAlertDto) GetAlertsOk() (*[]KubernetesAlertCreateDto, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *CreateAlertDto) SetAlerts(v []KubernetesAlertCreateDto)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *CreateAlertDto) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### SetAlertsNil

`func (o *CreateAlertDto) SetAlertsNil(b bool)`

 SetAlertsNil sets the value for Alerts to be an explicit nil

### UnsetAlerts
`func (o *CreateAlertDto) UnsetAlerts()`

UnsetAlerts ensures that no value is present for Alerts, not even an explicit nil
### GetStatus

`func (o *CreateAlertDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateAlertDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateAlertDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateAlertDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CreateAlertDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CreateAlertDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


