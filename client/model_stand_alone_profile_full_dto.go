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

// checks if the StandAloneProfileFullDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfileFullDto{}

// StandAloneProfileFullDto struct for StandAloneProfileFullDto
type StandAloneProfileFullDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	PublicKey NullableString `json:"publicKey,omitempty"`
	Revision *int32 `json:"revision,omitempty"`
	StandAloneProfileSecurityGroups []StandAloneProfileSecurityGroupFullDto `json:"standAloneProfileSecurityGroups,omitempty"`
}

// NewStandAloneProfileFullDto instantiates a new StandAloneProfileFullDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfileFullDto() *StandAloneProfileFullDto {
	this := StandAloneProfileFullDto{}
	return &this
}

// NewStandAloneProfileFullDtoWithDefaults instantiates a new StandAloneProfileFullDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfileFullDtoWithDefaults() *StandAloneProfileFullDto {
	this := StandAloneProfileFullDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneProfileFullDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileFullDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneProfileFullDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneProfileFullDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileFullDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileFullDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneProfileFullDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StandAloneProfileFullDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *StandAloneProfileFullDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StandAloneProfileFullDto) UnsetName() {
	o.Name.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileFullDto) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileFullDto) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *StandAloneProfileFullDto) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *StandAloneProfileFullDto) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}
// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *StandAloneProfileFullDto) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *StandAloneProfileFullDto) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *StandAloneProfileFullDto) GetRevision() int32 {
	if o == nil || IsNil(o.Revision) {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileFullDto) GetRevisionOk() (*int32, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *StandAloneProfileFullDto) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *StandAloneProfileFullDto) SetRevision(v int32) {
	o.Revision = &v
}

// GetStandAloneProfileSecurityGroups returns the StandAloneProfileSecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileFullDto) GetStandAloneProfileSecurityGroups() []StandAloneProfileSecurityGroupFullDto {
	if o == nil {
		var ret []StandAloneProfileSecurityGroupFullDto
		return ret
	}
	return o.StandAloneProfileSecurityGroups
}

// GetStandAloneProfileSecurityGroupsOk returns a tuple with the StandAloneProfileSecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileFullDto) GetStandAloneProfileSecurityGroupsOk() ([]StandAloneProfileSecurityGroupFullDto, bool) {
	if o == nil || IsNil(o.StandAloneProfileSecurityGroups) {
		return nil, false
	}
	return o.StandAloneProfileSecurityGroups, true
}

// HasStandAloneProfileSecurityGroups returns a boolean if a field has been set.
func (o *StandAloneProfileFullDto) HasStandAloneProfileSecurityGroups() bool {
	if o != nil && IsNil(o.StandAloneProfileSecurityGroups) {
		return true
	}

	return false
}

// SetStandAloneProfileSecurityGroups gets a reference to the given []StandAloneProfileSecurityGroupFullDto and assigns it to the StandAloneProfileSecurityGroups field.
func (o *StandAloneProfileFullDto) SetStandAloneProfileSecurityGroups(v []StandAloneProfileSecurityGroupFullDto) {
	o.StandAloneProfileSecurityGroups = v
}

func (o StandAloneProfileFullDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfileFullDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PublicKey.IsSet() {
		toSerialize["publicKey"] = o.PublicKey.Get()
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if o.StandAloneProfileSecurityGroups != nil {
		toSerialize["standAloneProfileSecurityGroups"] = o.StandAloneProfileSecurityGroups
	}
	return toSerialize, nil
}

type NullableStandAloneProfileFullDto struct {
	value *StandAloneProfileFullDto
	isSet bool
}

func (v NullableStandAloneProfileFullDto) Get() *StandAloneProfileFullDto {
	return v.value
}

func (v *NullableStandAloneProfileFullDto) Set(val *StandAloneProfileFullDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfileFullDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfileFullDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfileFullDto(val *StandAloneProfileFullDto) *NullableStandAloneProfileFullDto {
	return &NullableStandAloneProfileFullDto{value: val, isSet: true}
}

func (v NullableStandAloneProfileFullDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfileFullDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


