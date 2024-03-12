/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the KubernetesAlertDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAlertDto{}

// KubernetesAlertDto struct for KubernetesAlertDto
type KubernetesAlertDto struct {
	Id *int32 `json:"id,omitempty"`
	Labels interface{} `json:"labels,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Severity NullableString `json:"severity,omitempty"`
	Fingerprint NullableString `json:"fingerprint,omitempty"`
	Status NullableString `json:"status,omitempty"`
	StartsAt NullableString `json:"startsAt,omitempty"`
	EndAt NullableString `json:"endAt,omitempty"`
	IsSolved *bool `json:"isSolved,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	IsSilenced *bool `json:"isSilenced,omitempty"`
	SilenceReason NullableString `json:"silenceReason,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
}

// NewKubernetesAlertDto instantiates a new KubernetesAlertDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAlertDto() *KubernetesAlertDto {
	this := KubernetesAlertDto{}
	return &this
}

// NewKubernetesAlertDtoWithDefaults instantiates a new KubernetesAlertDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAlertDtoWithDefaults() *KubernetesAlertDto {
	this := KubernetesAlertDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesAlertDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesAlertDto) SetId(v int32) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetLabelsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given interface{} and assigns it to the Labels field.
func (o *KubernetesAlertDto) SetLabels(v interface{}) {
	o.Labels = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *KubernetesAlertDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *KubernetesAlertDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *KubernetesAlertDto) UnsetDescription() {
	o.Description.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *KubernetesAlertDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *KubernetesAlertDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *KubernetesAlertDto) UnsetTitle() {
	o.Title.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetSeverity() string {
	if o == nil || IsNil(o.Severity.Get()) {
		var ret string
		return ret
	}
	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// HasSeverity returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasSeverity() bool {
	if o != nil && o.Severity.IsSet() {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given NullableString and assigns it to the Severity field.
func (o *KubernetesAlertDto) SetSeverity(v string) {
	o.Severity.Set(&v)
}
// SetSeverityNil sets the value for Severity to be an explicit nil
func (o *KubernetesAlertDto) SetSeverityNil() {
	o.Severity.Set(nil)
}

// UnsetSeverity ensures that no value is present for Severity, not even an explicit nil
func (o *KubernetesAlertDto) UnsetSeverity() {
	o.Severity.Unset()
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// HasFingerprint returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasFingerprint() bool {
	if o != nil && o.Fingerprint.IsSet() {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given NullableString and assigns it to the Fingerprint field.
func (o *KubernetesAlertDto) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}
// SetFingerprintNil sets the value for Fingerprint to be an explicit nil
func (o *KubernetesAlertDto) SetFingerprintNil() {
	o.Fingerprint.Set(nil)
}

// UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
func (o *KubernetesAlertDto) UnsetFingerprint() {
	o.Fingerprint.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *KubernetesAlertDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesAlertDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesAlertDto) UnsetStatus() {
	o.Status.Unset()
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetStartsAt() string {
	if o == nil || IsNil(o.StartsAt.Get()) {
		var ret string
		return ret
	}
	return *o.StartsAt.Get()
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetStartsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartsAt.Get(), o.StartsAt.IsSet()
}

// HasStartsAt returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasStartsAt() bool {
	if o != nil && o.StartsAt.IsSet() {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given NullableString and assigns it to the StartsAt field.
func (o *KubernetesAlertDto) SetStartsAt(v string) {
	o.StartsAt.Set(&v)
}
// SetStartsAtNil sets the value for StartsAt to be an explicit nil
func (o *KubernetesAlertDto) SetStartsAtNil() {
	o.StartsAt.Set(nil)
}

// UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
func (o *KubernetesAlertDto) UnsetStartsAt() {
	o.StartsAt.Unset()
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetEndAt() string {
	if o == nil || IsNil(o.EndAt.Get()) {
		var ret string
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetEndAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given NullableString and assigns it to the EndAt field.
func (o *KubernetesAlertDto) SetEndAt(v string) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *KubernetesAlertDto) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *KubernetesAlertDto) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetIsSolved returns the IsSolved field value if set, zero value otherwise.
func (o *KubernetesAlertDto) GetIsSolved() bool {
	if o == nil || IsNil(o.IsSolved) {
		var ret bool
		return ret
	}
	return *o.IsSolved
}

// GetIsSolvedOk returns a tuple with the IsSolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDto) GetIsSolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSolved) {
		return nil, false
	}
	return o.IsSolved, true
}

// HasIsSolved returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasIsSolved() bool {
	if o != nil && !IsNil(o.IsSolved) {
		return true
	}

	return false
}

// SetIsSolved gets a reference to the given bool and assigns it to the IsSolved field.
func (o *KubernetesAlertDto) SetIsSolved(v bool) {
	o.IsSolved = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *KubernetesAlertDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *KubernetesAlertDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *KubernetesAlertDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *KubernetesAlertDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *KubernetesAlertDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetIsSilenced returns the IsSilenced field value if set, zero value otherwise.
func (o *KubernetesAlertDto) GetIsSilenced() bool {
	if o == nil || IsNil(o.IsSilenced) {
		var ret bool
		return ret
	}
	return *o.IsSilenced
}

// GetIsSilencedOk returns a tuple with the IsSilenced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDto) GetIsSilencedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSilenced) {
		return nil, false
	}
	return o.IsSilenced, true
}

// HasIsSilenced returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasIsSilenced() bool {
	if o != nil && !IsNil(o.IsSilenced) {
		return true
	}

	return false
}

// SetIsSilenced gets a reference to the given bool and assigns it to the IsSilenced field.
func (o *KubernetesAlertDto) SetIsSilenced(v bool) {
	o.IsSilenced = &v
}

// GetSilenceReason returns the SilenceReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetSilenceReason() string {
	if o == nil || IsNil(o.SilenceReason.Get()) {
		var ret string
		return ret
	}
	return *o.SilenceReason.Get()
}

// GetSilenceReasonOk returns a tuple with the SilenceReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetSilenceReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SilenceReason.Get(), o.SilenceReason.IsSet()
}

// HasSilenceReason returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasSilenceReason() bool {
	if o != nil && o.SilenceReason.IsSet() {
		return true
	}

	return false
}

// SetSilenceReason gets a reference to the given NullableString and assigns it to the SilenceReason field.
func (o *KubernetesAlertDto) SetSilenceReason(v string) {
	o.SilenceReason.Set(&v)
}
// SetSilenceReasonNil sets the value for SilenceReason to be an explicit nil
func (o *KubernetesAlertDto) SetSilenceReasonNil() {
	o.SilenceReason.Set(nil)
}

// UnsetSilenceReason ensures that no value is present for SilenceReason, not even an explicit nil
func (o *KubernetesAlertDto) UnsetSilenceReason() {
	o.SilenceReason.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesAlertDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *KubernetesAlertDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *KubernetesAlertDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *KubernetesAlertDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *KubernetesAlertDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

func (o KubernetesAlertDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAlertDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Severity.IsSet() {
		toSerialize["severity"] = o.Severity.Get()
	}
	if o.Fingerprint.IsSet() {
		toSerialize["fingerprint"] = o.Fingerprint.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StartsAt.IsSet() {
		toSerialize["startsAt"] = o.StartsAt.Get()
	}
	if o.EndAt.IsSet() {
		toSerialize["endAt"] = o.EndAt.Get()
	}
	if !IsNil(o.IsSolved) {
		toSerialize["isSolved"] = o.IsSolved
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.IsSilenced) {
		toSerialize["isSilenced"] = o.IsSilenced
	}
	if o.SilenceReason.IsSet() {
		toSerialize["silenceReason"] = o.SilenceReason.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	return toSerialize, nil
}

type NullableKubernetesAlertDto struct {
	value *KubernetesAlertDto
	isSet bool
}

func (v NullableKubernetesAlertDto) Get() *KubernetesAlertDto {
	return v.value
}

func (v *NullableKubernetesAlertDto) Set(val *KubernetesAlertDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAlertDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAlertDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAlertDto(val *KubernetesAlertDto) *NullableKubernetesAlertDto {
	return &NullableKubernetesAlertDto{value: val, isSet: true}
}

func (v NullableKubernetesAlertDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAlertDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


