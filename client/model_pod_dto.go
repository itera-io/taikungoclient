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

// checks if the PodDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PodDto{}

// PodDto struct for PodDto
type PodDto struct {
	MetadataName NullableString `json:"metadataName,omitempty"`
	Status       NullableString `json:"status,omitempty"`
	RestartCount *int32         `json:"restartCount,omitempty"`
	Namespace    NullableString `json:"namespace,omitempty"`
	Age          NullableTime   `json:"age,omitempty"`
	Node         NullableString `json:"node,omitempty"`
	Phase        NullableString `json:"phase,omitempty"`
}

// NewPodDto instantiates a new PodDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodDto() *PodDto {
	this := PodDto{}
	return &this
}

// NewPodDtoWithDefaults instantiates a new PodDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodDtoWithDefaults() *PodDto {
	this := PodDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDto) GetMetadataName() string {
	if o == nil || IsNil(o.MetadataName.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataName.Get()
}

// GetMetadataNameOk returns a tuple with the MetadataName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataName.Get(), o.MetadataName.IsSet()
}

// HasMetadataName returns a boolean if a field has been set.
func (o *PodDto) HasMetadataName() bool {
	if o != nil && o.MetadataName.IsSet() {
		return true
	}

	return false
}

// SetMetadataName gets a reference to the given NullableString and assigns it to the MetadataName field.
func (o *PodDto) SetMetadataName(v string) {
	o.MetadataName.Set(&v)
}

// SetMetadataNameNil sets the value for MetadataName to be an explicit nil
func (o *PodDto) SetMetadataNameNil() {
	o.MetadataName.Set(nil)
}

// UnsetMetadataName ensures that no value is present for MetadataName, not even an explicit nil
func (o *PodDto) UnsetMetadataName() {
	o.MetadataName.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PodDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PodDto) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *PodDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PodDto) UnsetStatus() {
	o.Status.Unset()
}

// GetRestartCount returns the RestartCount field value if set, zero value otherwise.
func (o *PodDto) GetRestartCount() int32 {
	if o == nil || IsNil(o.RestartCount) {
		var ret int32
		return ret
	}
	return *o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodDto) GetRestartCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RestartCount) {
		return nil, false
	}
	return o.RestartCount, true
}

// HasRestartCount returns a boolean if a field has been set.
func (o *PodDto) HasRestartCount() bool {
	if o != nil && !IsNil(o.RestartCount) {
		return true
	}

	return false
}

// SetRestartCount gets a reference to the given int32 and assigns it to the RestartCount field.
func (o *PodDto) SetRestartCount(v int32) {
	o.RestartCount = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *PodDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *PodDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *PodDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *PodDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetAge returns the Age field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDto) GetAge() time.Time {
	if o == nil || IsNil(o.Age.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDto) GetAgeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// HasAge returns a boolean if a field has been set.
func (o *PodDto) HasAge() bool {
	if o != nil && o.Age.IsSet() {
		return true
	}

	return false
}

// SetAge gets a reference to the given NullableTime and assigns it to the Age field.
func (o *PodDto) SetAge(v time.Time) {
	o.Age.Set(&v)
}

// SetAgeNil sets the value for Age to be an explicit nil
func (o *PodDto) SetAgeNil() {
	o.Age.Set(nil)
}

// UnsetAge ensures that no value is present for Age, not even an explicit nil
func (o *PodDto) UnsetAge() {
	o.Age.Unset()
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDto) GetNode() string {
	if o == nil || IsNil(o.Node.Get()) {
		var ret string
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDto) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *PodDto) HasNode() bool {
	if o != nil && o.Node.IsSet() {
		return true
	}

	return false
}

// SetNode gets a reference to the given NullableString and assigns it to the Node field.
func (o *PodDto) SetNode(v string) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil
func (o *PodDto) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil
func (o *PodDto) UnsetNode() {
	o.Node.Unset()
}

// GetPhase returns the Phase field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodDto) GetPhase() string {
	if o == nil || IsNil(o.Phase.Get()) {
		var ret string
		return ret
	}
	return *o.Phase.Get()
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDto) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase.Get(), o.Phase.IsSet()
}

// HasPhase returns a boolean if a field has been set.
func (o *PodDto) HasPhase() bool {
	if o != nil && o.Phase.IsSet() {
		return true
	}

	return false
}

// SetPhase gets a reference to the given NullableString and assigns it to the Phase field.
func (o *PodDto) SetPhase(v string) {
	o.Phase.Set(&v)
}

// SetPhaseNil sets the value for Phase to be an explicit nil
func (o *PodDto) SetPhaseNil() {
	o.Phase.Set(nil)
}

// UnsetPhase ensures that no value is present for Phase, not even an explicit nil
func (o *PodDto) UnsetPhase() {
	o.Phase.Unset()
}

func (o PodDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataName.IsSet() {
		toSerialize["metadataName"] = o.MetadataName.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.RestartCount) {
		toSerialize["restartCount"] = o.RestartCount
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Age.IsSet() {
		toSerialize["age"] = o.Age.Get()
	}
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}
	if o.Phase.IsSet() {
		toSerialize["phase"] = o.Phase.Get()
	}
	return toSerialize, nil
}

type NullablePodDto struct {
	value *PodDto
	isSet bool
}

func (v NullablePodDto) Get() *PodDto {
	return v.value
}

func (v *NullablePodDto) Set(val *PodDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePodDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePodDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodDto(val *PodDto) *NullablePodDto {
	return &NullablePodDto{value: val, isSet: true}
}

func (v NullablePodDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
