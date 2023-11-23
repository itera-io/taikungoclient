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

// checks if the OpenshiftImagesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenshiftImagesData{}

// OpenshiftImagesData struct for OpenshiftImagesData
type OpenshiftImagesData struct {
	Name NullableString `json:"name,omitempty"`
	PvcName NullableString `json:"pvcName,omitempty"`
	PvcCapacity NullableString `json:"pvcCapacity,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Status NullableString `json:"status,omitempty"`
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
}

// NewOpenshiftImagesData instantiates a new OpenshiftImagesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenshiftImagesData() *OpenshiftImagesData {
	this := OpenshiftImagesData{}
	return &this
}

// NewOpenshiftImagesDataWithDefaults instantiates a new OpenshiftImagesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenshiftImagesDataWithDefaults() *OpenshiftImagesData {
	this := OpenshiftImagesData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftImagesData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftImagesData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OpenshiftImagesData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OpenshiftImagesData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OpenshiftImagesData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OpenshiftImagesData) UnsetName() {
	o.Name.Unset()
}

// GetPvcName returns the PvcName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftImagesData) GetPvcName() string {
	if o == nil || IsNil(o.PvcName.Get()) {
		var ret string
		return ret
	}
	return *o.PvcName.Get()
}

// GetPvcNameOk returns a tuple with the PvcName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftImagesData) GetPvcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PvcName.Get(), o.PvcName.IsSet()
}

// HasPvcName returns a boolean if a field has been set.
func (o *OpenshiftImagesData) HasPvcName() bool {
	if o != nil && o.PvcName.IsSet() {
		return true
	}

	return false
}

// SetPvcName gets a reference to the given NullableString and assigns it to the PvcName field.
func (o *OpenshiftImagesData) SetPvcName(v string) {
	o.PvcName.Set(&v)
}
// SetPvcNameNil sets the value for PvcName to be an explicit nil
func (o *OpenshiftImagesData) SetPvcNameNil() {
	o.PvcName.Set(nil)
}

// UnsetPvcName ensures that no value is present for PvcName, not even an explicit nil
func (o *OpenshiftImagesData) UnsetPvcName() {
	o.PvcName.Unset()
}

// GetPvcCapacity returns the PvcCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftImagesData) GetPvcCapacity() string {
	if o == nil || IsNil(o.PvcCapacity.Get()) {
		var ret string
		return ret
	}
	return *o.PvcCapacity.Get()
}

// GetPvcCapacityOk returns a tuple with the PvcCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftImagesData) GetPvcCapacityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PvcCapacity.Get(), o.PvcCapacity.IsSet()
}

// HasPvcCapacity returns a boolean if a field has been set.
func (o *OpenshiftImagesData) HasPvcCapacity() bool {
	if o != nil && o.PvcCapacity.IsSet() {
		return true
	}

	return false
}

// SetPvcCapacity gets a reference to the given NullableString and assigns it to the PvcCapacity field.
func (o *OpenshiftImagesData) SetPvcCapacity(v string) {
	o.PvcCapacity.Set(&v)
}
// SetPvcCapacityNil sets the value for PvcCapacity to be an explicit nil
func (o *OpenshiftImagesData) SetPvcCapacityNil() {
	o.PvcCapacity.Set(nil)
}

// UnsetPvcCapacity ensures that no value is present for PvcCapacity, not even an explicit nil
func (o *OpenshiftImagesData) UnsetPvcCapacity() {
	o.PvcCapacity.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftImagesData) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftImagesData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *OpenshiftImagesData) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *OpenshiftImagesData) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *OpenshiftImagesData) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *OpenshiftImagesData) UnsetType() {
	o.Type.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenshiftImagesData) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftImagesData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *OpenshiftImagesData) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *OpenshiftImagesData) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *OpenshiftImagesData) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *OpenshiftImagesData) UnsetStatus() {
	o.Status.Unset()
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *OpenshiftImagesData) GetCreationTimestamp() time.Time {
	if o == nil || IsNil(o.CreationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenshiftImagesData) GetCreationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *OpenshiftImagesData) HasCreationTimestamp() bool {
	if o != nil && !IsNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given time.Time and assigns it to the CreationTimestamp field.
func (o *OpenshiftImagesData) SetCreationTimestamp(v time.Time) {
	o.CreationTimestamp = &v
}

func (o OpenshiftImagesData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenshiftImagesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PvcName.IsSet() {
		toSerialize["pvcName"] = o.PvcName.Get()
	}
	if o.PvcCapacity.IsSet() {
		toSerialize["pvcCapacity"] = o.PvcCapacity.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.CreationTimestamp) {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	return toSerialize, nil
}

type NullableOpenshiftImagesData struct {
	value *OpenshiftImagesData
	isSet bool
}

func (v NullableOpenshiftImagesData) Get() *OpenshiftImagesData {
	return v.value
}

func (v *NullableOpenshiftImagesData) Set(val *OpenshiftImagesData) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenshiftImagesData) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenshiftImagesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenshiftImagesData(val *OpenshiftImagesData) *NullableOpenshiftImagesData {
	return &NullableOpenshiftImagesData{value: val, isSet: true}
}

func (v NullableOpenshiftImagesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenshiftImagesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


