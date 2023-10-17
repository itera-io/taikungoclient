/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the CommonAvailabilityDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonAvailabilityDto{}

// CommonAvailabilityDto struct for CommonAvailabilityDto
type CommonAvailabilityDto struct {
	Id   *bool          `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewCommonAvailabilityDto instantiates a new CommonAvailabilityDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAvailabilityDto() *CommonAvailabilityDto {
	this := CommonAvailabilityDto{}
	return &this
}

// NewCommonAvailabilityDtoWithDefaults instantiates a new CommonAvailabilityDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAvailabilityDtoWithDefaults() *CommonAvailabilityDto {
	this := CommonAvailabilityDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommonAvailabilityDto) GetId() bool {
	if o == nil || IsNil(o.Id) {
		var ret bool
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAvailabilityDto) GetIdOk() (*bool, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommonAvailabilityDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given bool and assigns it to the Id field.
func (o *CommonAvailabilityDto) SetId(v bool) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonAvailabilityDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonAvailabilityDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CommonAvailabilityDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CommonAvailabilityDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CommonAvailabilityDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CommonAvailabilityDto) UnsetName() {
	o.Name.Unset()
}

func (o CommonAvailabilityDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonAvailabilityDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableCommonAvailabilityDto struct {
	value *CommonAvailabilityDto
	isSet bool
}

func (v NullableCommonAvailabilityDto) Get() *CommonAvailabilityDto {
	return v.value
}

func (v *NullableCommonAvailabilityDto) Set(val *CommonAvailabilityDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAvailabilityDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAvailabilityDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAvailabilityDto(val *CommonAvailabilityDto) *NullableCommonAvailabilityDto {
	return &NullableCommonAvailabilityDto{value: val, isSet: true}
}

func (v NullableCommonAvailabilityDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAvailabilityDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
