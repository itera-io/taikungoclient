/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the ProblemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetails{}

// ProblemDetails struct for ProblemDetails
type ProblemDetails struct {
	Type                 NullableString `json:"type,omitempty"`
	Title                NullableString `json:"title,omitempty"`
	Status               NullableInt32  `json:"status,omitempty"`
	Detail               NullableString `json:"detail,omitempty"`
	Instance             NullableString `json:"instance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProblemDetails ProblemDetails

// NewProblemDetails instantiates a new ProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetails() *ProblemDetails {
	this := ProblemDetails{}
	return &this
}

// NewProblemDetailsWithDefaults instantiates a new ProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailsWithDefaults() *ProblemDetails {
	this := ProblemDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDetails) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ProblemDetails) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ProblemDetails) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *ProblemDetails) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ProblemDetails) UnsetType() {
	o.Type.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDetails) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ProblemDetails) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ProblemDetails) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil
func (o *ProblemDetails) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ProblemDetails) UnsetTitle() {
	o.Title.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDetails) GetStatus() int32 {
	if o == nil || IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDetails) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ProblemDetails) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *ProblemDetails) SetStatus(v int32) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *ProblemDetails) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ProblemDetails) UnsetStatus() {
	o.Status.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDetails) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDetails) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemDetails) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *ProblemDetails) SetDetail(v string) {
	o.Detail.Set(&v)
}

// SetDetailNil sets the value for Detail to be an explicit nil
func (o *ProblemDetails) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *ProblemDetails) UnsetDetail() {
	o.Detail.Unset()
}

// GetInstance returns the Instance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProblemDetails) GetInstance() string {
	if o == nil || IsNil(o.Instance.Get()) {
		var ret string
		return ret
	}
	return *o.Instance.Get()
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProblemDetails) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instance.Get(), o.Instance.IsSet()
}

// HasInstance returns a boolean if a field has been set.
func (o *ProblemDetails) HasInstance() bool {
	if o != nil && o.Instance.IsSet() {
		return true
	}

	return false
}

// SetInstance gets a reference to the given NullableString and assigns it to the Instance field.
func (o *ProblemDetails) SetInstance(v string) {
	o.Instance.Set(&v)
}

// SetInstanceNil sets the value for Instance to be an explicit nil
func (o *ProblemDetails) SetInstanceNil() {
	o.Instance.Set(nil)
}

// UnsetInstance ensures that no value is present for Instance, not even an explicit nil
func (o *ProblemDetails) UnsetInstance() {
	o.Instance.Unset()
}

func (o ProblemDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	if o.Instance.IsSet() {
		toSerialize["instance"] = o.Instance.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProblemDetails) UnmarshalJSON(bytes []byte) (err error) {
	varProblemDetails := _ProblemDetails{}

	err = json.Unmarshal(bytes, &varProblemDetails)

	if err != nil {
		return err
	}

	*o = ProblemDetails(varProblemDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "status")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "instance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProblemDetails struct {
	value *ProblemDetails
	isSet bool
}

func (v NullableProblemDetails) Get() *ProblemDetails {
	return v.value
}

func (v *NullableProblemDetails) Set(val *ProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetails(val *ProblemDetails) *NullableProblemDetails {
	return &NullableProblemDetails{value: val, isSet: true}
}

func (v NullableProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
