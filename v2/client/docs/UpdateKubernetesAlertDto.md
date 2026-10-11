# UpdateKubernetesAlertDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**IsSilenced** | Pointer to **NullableBool** |  | [optional] 
**SilenceReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateKubernetesAlertDto

`func NewUpdateKubernetesAlertDto() *UpdateKubernetesAlertDto`

NewUpdateKubernetesAlertDto instantiates a new UpdateKubernetesAlertDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKubernetesAlertDtoWithDefaults

`func NewUpdateKubernetesAlertDtoWithDefaults() *UpdateKubernetesAlertDto`

NewUpdateKubernetesAlertDtoWithDefaults instantiates a new UpdateKubernetesAlertDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateKubernetesAlertDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateKubernetesAlertDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateKubernetesAlertDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateKubernetesAlertDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateKubernetesAlertDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateKubernetesAlertDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLabels

`func (o *UpdateKubernetesAlertDto) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateKubernetesAlertDto) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateKubernetesAlertDto) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateKubernetesAlertDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetStartsAt

`func (o *UpdateKubernetesAlertDto) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateKubernetesAlertDto) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateKubernetesAlertDto) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateKubernetesAlertDto) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *UpdateKubernetesAlertDto) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *UpdateKubernetesAlertDto) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *UpdateKubernetesAlertDto) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *UpdateKubernetesAlertDto) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *UpdateKubernetesAlertDto) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *UpdateKubernetesAlertDto) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *UpdateKubernetesAlertDto) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *UpdateKubernetesAlertDto) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *UpdateKubernetesAlertDto) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *UpdateKubernetesAlertDto) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetIsSilenced

`func (o *UpdateKubernetesAlertDto) GetIsSilenced() bool`

GetIsSilenced returns the IsSilenced field if non-nil, zero value otherwise.

### GetIsSilencedOk

`func (o *UpdateKubernetesAlertDto) GetIsSilencedOk() (*bool, bool)`

GetIsSilencedOk returns a tuple with the IsSilenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSilenced

`func (o *UpdateKubernetesAlertDto) SetIsSilenced(v bool)`

SetIsSilenced sets IsSilenced field to given value.

### HasIsSilenced

`func (o *UpdateKubernetesAlertDto) HasIsSilenced() bool`

HasIsSilenced returns a boolean if a field has been set.

### SetIsSilencedNil

`func (o *UpdateKubernetesAlertDto) SetIsSilencedNil(b bool)`

 SetIsSilencedNil sets the value for IsSilenced to be an explicit nil

### UnsetIsSilenced
`func (o *UpdateKubernetesAlertDto) UnsetIsSilenced()`

UnsetIsSilenced ensures that no value is present for IsSilenced, not even an explicit nil
### GetSilenceReason

`func (o *UpdateKubernetesAlertDto) GetSilenceReason() string`

GetSilenceReason returns the SilenceReason field if non-nil, zero value otherwise.

### GetSilenceReasonOk

`func (o *UpdateKubernetesAlertDto) GetSilenceReasonOk() (*string, bool)`

GetSilenceReasonOk returns a tuple with the SilenceReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilenceReason

`func (o *UpdateKubernetesAlertDto) SetSilenceReason(v string)`

SetSilenceReason sets SilenceReason field to given value.

### HasSilenceReason

`func (o *UpdateKubernetesAlertDto) HasSilenceReason() bool`

HasSilenceReason returns a boolean if a field has been set.

### SetSilenceReasonNil

`func (o *UpdateKubernetesAlertDto) SetSilenceReasonNil(b bool)`

 SetSilenceReasonNil sets the value for SilenceReason to be an explicit nil

### UnsetSilenceReason
`func (o *UpdateKubernetesAlertDto) UnsetSilenceReason()`

UnsetSilenceReason ensures that no value is present for SilenceReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


