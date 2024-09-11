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
	"time"
)

// checks if the KubernetesEventCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesEventCreateDto{}

// KubernetesEventCreateDto struct for KubernetesEventCreateDto
type KubernetesEventCreateDto struct {
	Type *string `json:"type,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Source interface{} `json:"source,omitempty"`
	InvolvedObject interface{} `json:"involvedObject,omitempty"`
	FirstTimeStamp NullableTime `json:"firstTimeStamp,omitempty"`
	LastTimeStamp NullableTime `json:"lastTimeStamp,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewKubernetesEventCreateDto instantiates a new KubernetesEventCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEventCreateDto() *KubernetesEventCreateDto {
	this := KubernetesEventCreateDto{}
	return &this
}

// NewKubernetesEventCreateDtoWithDefaults instantiates a new KubernetesEventCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEventCreateDtoWithDefaults() *KubernetesEventCreateDto {
	this := KubernetesEventCreateDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesEventCreateDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEventCreateDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesEventCreateDto) SetType(v string) {
	o.Type = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *KubernetesEventCreateDto) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEventCreateDto) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *KubernetesEventCreateDto) SetReason(v string) {
	o.Reason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KubernetesEventCreateDto) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEventCreateDto) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KubernetesEventCreateDto) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEventCreateDto) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEventCreateDto) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *KubernetesEventCreateDto) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEventCreateDto) GetSource() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEventCreateDto) GetSourceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given interface{} and assigns it to the Source field.
func (o *KubernetesEventCreateDto) SetSource(v interface{}) {
	o.Source = v
}

// GetInvolvedObject returns the InvolvedObject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEventCreateDto) GetInvolvedObject() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InvolvedObject
}

// GetInvolvedObjectOk returns a tuple with the InvolvedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEventCreateDto) GetInvolvedObjectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.InvolvedObject) {
		return nil, false
	}
	return &o.InvolvedObject, true
}

// HasInvolvedObject returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasInvolvedObject() bool {
	if o != nil && !IsNil(o.InvolvedObject) {
		return true
	}

	return false
}

// SetInvolvedObject gets a reference to the given interface{} and assigns it to the InvolvedObject field.
func (o *KubernetesEventCreateDto) SetInvolvedObject(v interface{}) {
	o.InvolvedObject = v
}

// GetFirstTimeStamp returns the FirstTimeStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEventCreateDto) GetFirstTimeStamp() time.Time {
	if o == nil || IsNil(o.FirstTimeStamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FirstTimeStamp.Get()
}

// GetFirstTimeStampOk returns a tuple with the FirstTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEventCreateDto) GetFirstTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstTimeStamp.Get(), o.FirstTimeStamp.IsSet()
}

// HasFirstTimeStamp returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasFirstTimeStamp() bool {
	if o != nil && o.FirstTimeStamp.IsSet() {
		return true
	}

	return false
}

// SetFirstTimeStamp gets a reference to the given NullableTime and assigns it to the FirstTimeStamp field.
func (o *KubernetesEventCreateDto) SetFirstTimeStamp(v time.Time) {
	o.FirstTimeStamp.Set(&v)
}
// SetFirstTimeStampNil sets the value for FirstTimeStamp to be an explicit nil
func (o *KubernetesEventCreateDto) SetFirstTimeStampNil() {
	o.FirstTimeStamp.Set(nil)
}

// UnsetFirstTimeStamp ensures that no value is present for FirstTimeStamp, not even an explicit nil
func (o *KubernetesEventCreateDto) UnsetFirstTimeStamp() {
	o.FirstTimeStamp.Unset()
}

// GetLastTimeStamp returns the LastTimeStamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEventCreateDto) GetLastTimeStamp() time.Time {
	if o == nil || IsNil(o.LastTimeStamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastTimeStamp.Get()
}

// GetLastTimeStampOk returns a tuple with the LastTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEventCreateDto) GetLastTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTimeStamp.Get(), o.LastTimeStamp.IsSet()
}

// HasLastTimeStamp returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasLastTimeStamp() bool {
	if o != nil && o.LastTimeStamp.IsSet() {
		return true
	}

	return false
}

// SetLastTimeStamp gets a reference to the given NullableTime and assigns it to the LastTimeStamp field.
func (o *KubernetesEventCreateDto) SetLastTimeStamp(v time.Time) {
	o.LastTimeStamp.Set(&v)
}
// SetLastTimeStampNil sets the value for LastTimeStamp to be an explicit nil
func (o *KubernetesEventCreateDto) SetLastTimeStampNil() {
	o.LastTimeStamp.Set(nil)
}

// UnsetLastTimeStamp ensures that no value is present for LastTimeStamp, not even an explicit nil
func (o *KubernetesEventCreateDto) UnsetLastTimeStamp() {
	o.LastTimeStamp.Unset()
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *KubernetesEventCreateDto) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEventCreateDto) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *KubernetesEventCreateDto) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *KubernetesEventCreateDto) SetCount(v int32) {
	o.Count = &v
}

func (o KubernetesEventCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesEventCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.InvolvedObject != nil {
		toSerialize["involvedObject"] = o.InvolvedObject
	}
	if o.FirstTimeStamp.IsSet() {
		toSerialize["firstTimeStamp"] = o.FirstTimeStamp.Get()
	}
	if o.LastTimeStamp.IsSet() {
		toSerialize["lastTimeStamp"] = o.LastTimeStamp.Get()
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableKubernetesEventCreateDto struct {
	value *KubernetesEventCreateDto
	isSet bool
}

func (v NullableKubernetesEventCreateDto) Get() *KubernetesEventCreateDto {
	return v.value
}

func (v *NullableKubernetesEventCreateDto) Set(val *KubernetesEventCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEventCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEventCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEventCreateDto(val *KubernetesEventCreateDto) *NullableKubernetesEventCreateDto {
	return &NullableKubernetesEventCreateDto{value: val, isSet: true}
}

func (v NullableKubernetesEventCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEventCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


