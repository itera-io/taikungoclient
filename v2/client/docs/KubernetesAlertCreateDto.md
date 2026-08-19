# KubernetesAlertCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to [**KubernetesAnnotations**](KubernetesAnnotations.md) |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKubernetesAlertCreateDto

`func NewKubernetesAlertCreateDto() *KubernetesAlertCreateDto`

NewKubernetesAlertCreateDto instantiates a new KubernetesAlertCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesAlertCreateDtoWithDefaults

`func NewKubernetesAlertCreateDtoWithDefaults() *KubernetesAlertCreateDto`

NewKubernetesAlertCreateDtoWithDefaults instantiates a new KubernetesAlertCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *KubernetesAlertCreateDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesAlertCreateDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesAlertCreateDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesAlertCreateDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KubernetesAlertCreateDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KubernetesAlertCreateDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLabels

`func (o *KubernetesAlertCreateDto) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesAlertCreateDto) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesAlertCreateDto) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *KubernetesAlertCreateDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAnnotations

`func (o *KubernetesAlertCreateDto) GetAnnotations() KubernetesAnnotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesAlertCreateDto) GetAnnotationsOk() (*KubernetesAnnotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesAlertCreateDto) SetAnnotations(v KubernetesAnnotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *KubernetesAlertCreateDto) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetStartsAt

`func (o *KubernetesAlertCreateDto) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *KubernetesAlertCreateDto) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *KubernetesAlertCreateDto) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *KubernetesAlertCreateDto) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *KubernetesAlertCreateDto) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *KubernetesAlertCreateDto) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *KubernetesAlertCreateDto) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *KubernetesAlertCreateDto) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetFingerprint

`func (o *KubernetesAlertCreateDto) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *KubernetesAlertCreateDto) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *KubernetesAlertCreateDto) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *KubernetesAlertCreateDto) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *KubernetesAlertCreateDto) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *KubernetesAlertCreateDto) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


