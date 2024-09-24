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
	"bytes"
	"fmt"
)

// checks if the KubernetesAlertDtoForPoller type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAlertDtoForPoller{}

// KubernetesAlertDtoForPoller struct for KubernetesAlertDtoForPoller
type KubernetesAlertDtoForPoller struct {
	Id int32 `json:"id"`
	Labels interface{} `json:"labels"`
	Description NullableString `json:"description"`
	Title NullableString `json:"title"`
	Severity NullableString `json:"severity"`
	Fingerprint NullableString `json:"fingerprint"`
	Status NullableString `json:"status"`
	StartsAt NullableString `json:"startsAt"`
	EndAt NullableString `json:"endAt"`
	IsSolved bool `json:"isSolved"`
	ProjectId int32 `json:"projectId"`
	ProjectName NullableString `json:"projectName"`
	IsSilenced bool `json:"isSilenced"`
	SilenceReason NullableString `json:"silenceReason"`
	LastModifiedBy NullableString `json:"lastModifiedBy"`
}

type _KubernetesAlertDtoForPoller KubernetesAlertDtoForPoller

// NewKubernetesAlertDtoForPoller instantiates a new KubernetesAlertDtoForPoller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAlertDtoForPoller(id int32, labels interface{}, description NullableString, title NullableString, severity NullableString, fingerprint NullableString, status NullableString, startsAt NullableString, endAt NullableString, isSolved bool, projectId int32, projectName NullableString, isSilenced bool, silenceReason NullableString, lastModifiedBy NullableString) *KubernetesAlertDtoForPoller {
	this := KubernetesAlertDtoForPoller{}
	this.Id = id
	this.Labels = labels
	this.Description = description
	this.Title = title
	this.Severity = severity
	this.Fingerprint = fingerprint
	this.Status = status
	this.StartsAt = startsAt
	this.EndAt = endAt
	this.IsSolved = isSolved
	this.ProjectId = projectId
	this.ProjectName = projectName
	this.IsSilenced = isSilenced
	this.SilenceReason = silenceReason
	this.LastModifiedBy = lastModifiedBy
	return &this
}

// NewKubernetesAlertDtoForPollerWithDefaults instantiates a new KubernetesAlertDtoForPoller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAlertDtoForPollerWithDefaults() *KubernetesAlertDtoForPoller {
	this := KubernetesAlertDtoForPoller{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesAlertDtoForPoller) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDtoForPoller) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesAlertDtoForPoller) SetId(v int32) {
	o.Id = v
}

// GetLabels returns the Labels field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *KubernetesAlertDtoForPoller) GetLabels() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetLabelsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *KubernetesAlertDtoForPoller) SetLabels(v interface{}) {
	o.Labels = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *KubernetesAlertDtoForPoller) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *KubernetesAlertDtoForPoller) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetSeverity returns the Severity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetSeverity() string {
	if o == nil || o.Severity.Get() == nil {
		var ret string
		return ret
	}

	return *o.Severity.Get()
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Severity.Get(), o.Severity.IsSet()
}

// SetSeverity sets field value
func (o *KubernetesAlertDtoForPoller) SetSeverity(v string) {
	o.Severity.Set(&v)
}

// GetFingerprint returns the Fingerprint field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetFingerprint() string {
	if o == nil || o.Fingerprint.Get() == nil {
		var ret string
		return ret
	}

	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// SetFingerprint sets field value
func (o *KubernetesAlertDtoForPoller) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *KubernetesAlertDtoForPoller) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetStartsAt returns the StartsAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetStartsAt() string {
	if o == nil || o.StartsAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartsAt.Get()
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetStartsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartsAt.Get(), o.StartsAt.IsSet()
}

// SetStartsAt sets field value
func (o *KubernetesAlertDtoForPoller) SetStartsAt(v string) {
	o.StartsAt.Set(&v)
}

// GetEndAt returns the EndAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetEndAt() string {
	if o == nil || o.EndAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetEndAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// SetEndAt sets field value
func (o *KubernetesAlertDtoForPoller) SetEndAt(v string) {
	o.EndAt.Set(&v)
}

// GetIsSolved returns the IsSolved field value
func (o *KubernetesAlertDtoForPoller) GetIsSolved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSolved
}

// GetIsSolvedOk returns a tuple with the IsSolved field value
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDtoForPoller) GetIsSolvedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSolved, true
}

// SetIsSolved sets field value
func (o *KubernetesAlertDtoForPoller) SetIsSolved(v bool) {
	o.IsSolved = v
}

// GetProjectId returns the ProjectId field value
func (o *KubernetesAlertDtoForPoller) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDtoForPoller) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *KubernetesAlertDtoForPoller) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetProjectName() string {
	if o == nil || o.ProjectName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// SetProjectName sets field value
func (o *KubernetesAlertDtoForPoller) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// GetIsSilenced returns the IsSilenced field value
func (o *KubernetesAlertDtoForPoller) GetIsSilenced() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSilenced
}

// GetIsSilencedOk returns a tuple with the IsSilenced field value
// and a boolean to check if the value has been set.
func (o *KubernetesAlertDtoForPoller) GetIsSilencedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSilenced, true
}

// SetIsSilenced sets field value
func (o *KubernetesAlertDtoForPoller) SetIsSilenced(v bool) {
	o.IsSilenced = v
}

// GetSilenceReason returns the SilenceReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetSilenceReason() string {
	if o == nil || o.SilenceReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.SilenceReason.Get()
}

// GetSilenceReasonOk returns a tuple with the SilenceReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetSilenceReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SilenceReason.Get(), o.SilenceReason.IsSet()
}

// SetSilenceReason sets field value
func (o *KubernetesAlertDtoForPoller) SetSilenceReason(v string) {
	o.SilenceReason.Set(&v)
}

// GetLastModifiedBy returns the LastModifiedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesAlertDtoForPoller) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesAlertDtoForPoller) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// SetLastModifiedBy sets field value
func (o *KubernetesAlertDtoForPoller) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}

func (o KubernetesAlertDtoForPoller) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAlertDtoForPoller) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["description"] = o.Description.Get()
	toSerialize["title"] = o.Title.Get()
	toSerialize["severity"] = o.Severity.Get()
	toSerialize["fingerprint"] = o.Fingerprint.Get()
	toSerialize["status"] = o.Status.Get()
	toSerialize["startsAt"] = o.StartsAt.Get()
	toSerialize["endAt"] = o.EndAt.Get()
	toSerialize["isSolved"] = o.IsSolved
	toSerialize["projectId"] = o.ProjectId
	toSerialize["projectName"] = o.ProjectName.Get()
	toSerialize["isSilenced"] = o.IsSilenced
	toSerialize["silenceReason"] = o.SilenceReason.Get()
	toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	return toSerialize, nil
}

func (o *KubernetesAlertDtoForPoller) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"labels",
		"description",
		"title",
		"severity",
		"fingerprint",
		"status",
		"startsAt",
		"endAt",
		"isSolved",
		"projectId",
		"projectName",
		"isSilenced",
		"silenceReason",
		"lastModifiedBy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubernetesAlertDtoForPoller := _KubernetesAlertDtoForPoller{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesAlertDtoForPoller)

	if err != nil {
		return err
	}

	*o = KubernetesAlertDtoForPoller(varKubernetesAlertDtoForPoller)

	return err
}

type NullableKubernetesAlertDtoForPoller struct {
	value *KubernetesAlertDtoForPoller
	isSet bool
}

func (v NullableKubernetesAlertDtoForPoller) Get() *KubernetesAlertDtoForPoller {
	return v.value
}

func (v *NullableKubernetesAlertDtoForPoller) Set(val *KubernetesAlertDtoForPoller) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAlertDtoForPoller) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAlertDtoForPoller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAlertDtoForPoller(val *KubernetesAlertDtoForPoller) *NullableKubernetesAlertDtoForPoller {
	return &NullableKubernetesAlertDtoForPoller{value: val, isSet: true}
}

func (v NullableKubernetesAlertDtoForPoller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAlertDtoForPoller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


