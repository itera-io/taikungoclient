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

// checks if the StandAloneProfileForDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfileForDetailsDto{}

// StandAloneProfileForDetailsDto struct for StandAloneProfileForDetailsDto
type StandAloneProfileForDetailsDto struct {
	Id             *int32                                        `json:"id,omitempty"`
	Name           NullableString                                `json:"name,omitempty"`
	PublicKey      NullableString                                `json:"publicKey,omitempty"`
	SecurityGroups []StandAloneProfileSecurityGroupForDetailsDto `json:"securityGroups,omitempty"`
}

// NewStandAloneProfileForDetailsDto instantiates a new StandAloneProfileForDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfileForDetailsDto() *StandAloneProfileForDetailsDto {
	this := StandAloneProfileForDetailsDto{}
	return &this
}

// NewStandAloneProfileForDetailsDtoWithDefaults instantiates a new StandAloneProfileForDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfileForDetailsDtoWithDefaults() *StandAloneProfileForDetailsDto {
	this := StandAloneProfileForDetailsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneProfileForDetailsDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileForDetailsDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneProfileForDetailsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneProfileForDetailsDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileForDetailsDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileForDetailsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneProfileForDetailsDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StandAloneProfileForDetailsDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *StandAloneProfileForDetailsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StandAloneProfileForDetailsDto) UnsetName() {
	o.Name.Unset()
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileForDetailsDto) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileForDetailsDto) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *StandAloneProfileForDetailsDto) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *StandAloneProfileForDetailsDto) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}

// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *StandAloneProfileForDetailsDto) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *StandAloneProfileForDetailsDto) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileForDetailsDto) GetSecurityGroups() []StandAloneProfileSecurityGroupForDetailsDto {
	if o == nil {
		var ret []StandAloneProfileSecurityGroupForDetailsDto
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileForDetailsDto) GetSecurityGroupsOk() ([]StandAloneProfileSecurityGroupForDetailsDto, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *StandAloneProfileForDetailsDto) HasSecurityGroups() bool {
	if o != nil && IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []StandAloneProfileSecurityGroupForDetailsDto and assigns it to the SecurityGroups field.
func (o *StandAloneProfileForDetailsDto) SetSecurityGroups(v []StandAloneProfileSecurityGroupForDetailsDto) {
	o.SecurityGroups = v
}

func (o StandAloneProfileForDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfileForDetailsDto) ToMap() (map[string]interface{}, error) {
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
	if o.SecurityGroups != nil {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	return toSerialize, nil
}

type NullableStandAloneProfileForDetailsDto struct {
	value *StandAloneProfileForDetailsDto
	isSet bool
}

func (v NullableStandAloneProfileForDetailsDto) Get() *StandAloneProfileForDetailsDto {
	return v.value
}

func (v *NullableStandAloneProfileForDetailsDto) Set(val *StandAloneProfileForDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfileForDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfileForDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfileForDetailsDto(val *StandAloneProfileForDetailsDto) *NullableStandAloneProfileForDetailsDto {
	return &NullableStandAloneProfileForDetailsDto{value: val, isSet: true}
}

func (v NullableStandAloneProfileForDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfileForDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
