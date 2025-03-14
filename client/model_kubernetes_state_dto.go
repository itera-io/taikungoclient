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

// checks if the KubernetesStateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesStateDto{}

// KubernetesStateDto struct for KubernetesStateDto
type KubernetesStateDto struct {
	Status NullableString `json:"status,omitempty"`
	Reason *EKubernetesState `json:"reason,omitempty"`
}

// NewKubernetesStateDto instantiates a new KubernetesStateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesStateDto() *KubernetesStateDto {
	this := KubernetesStateDto{}
	return &this
}

// NewKubernetesStateDtoWithDefaults instantiates a new KubernetesStateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesStateDtoWithDefaults() *KubernetesStateDto {
	this := KubernetesStateDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesStateDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesStateDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *KubernetesStateDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *KubernetesStateDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *KubernetesStateDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *KubernetesStateDto) UnsetStatus() {
	o.Status.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *KubernetesStateDto) GetReason() EKubernetesState {
	if o == nil || IsNil(o.Reason) {
		var ret EKubernetesState
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStateDto) GetReasonOk() (*EKubernetesState, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *KubernetesStateDto) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given EKubernetesState and assigns it to the Reason field.
func (o *KubernetesStateDto) SetReason(v EKubernetesState) {
	o.Reason = &v
}

func (o KubernetesStateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesStateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableKubernetesStateDto struct {
	value *KubernetesStateDto
	isSet bool
}

func (v NullableKubernetesStateDto) Get() *KubernetesStateDto {
	return v.value
}

func (v *NullableKubernetesStateDto) Set(val *KubernetesStateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesStateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesStateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesStateDto(val *KubernetesStateDto) *NullableKubernetesStateDto {
	return &NullableKubernetesStateDto{value: val, isSet: true}
}

func (v NullableKubernetesStateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesStateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


