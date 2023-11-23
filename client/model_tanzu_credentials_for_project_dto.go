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

// checks if the TanzuCredentialsForProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TanzuCredentialsForProjectDto{}

// TanzuCredentialsForProjectDto struct for TanzuCredentialsForProjectDto
type TanzuCredentialsForProjectDto struct {
	Username NullableString `json:"username,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Password NullableString `json:"password,omitempty"`
	VolumeType NullableString `json:"volumeType,omitempty"`
	Namespace NullableString `json:"namespace,omitempty"`
	Port NullableInt32 `json:"port,omitempty"`
}

// NewTanzuCredentialsForProjectDto instantiates a new TanzuCredentialsForProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTanzuCredentialsForProjectDto() *TanzuCredentialsForProjectDto {
	this := TanzuCredentialsForProjectDto{}
	return &this
}

// NewTanzuCredentialsForProjectDtoWithDefaults instantiates a new TanzuCredentialsForProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTanzuCredentialsForProjectDtoWithDefaults() *TanzuCredentialsForProjectDto {
	this := TanzuCredentialsForProjectDto{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsForProjectDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsForProjectDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *TanzuCredentialsForProjectDto) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *TanzuCredentialsForProjectDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *TanzuCredentialsForProjectDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *TanzuCredentialsForProjectDto) UnsetUsername() {
	o.Username.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsForProjectDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsForProjectDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *TanzuCredentialsForProjectDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *TanzuCredentialsForProjectDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *TanzuCredentialsForProjectDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *TanzuCredentialsForProjectDto) UnsetUrl() {
	o.Url.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsForProjectDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsForProjectDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *TanzuCredentialsForProjectDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *TanzuCredentialsForProjectDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *TanzuCredentialsForProjectDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *TanzuCredentialsForProjectDto) UnsetPassword() {
	o.Password.Unset()
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsForProjectDto) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType.Get()) {
		var ret string
		return ret
	}
	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsForProjectDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// HasVolumeType returns a boolean if a field has been set.
func (o *TanzuCredentialsForProjectDto) HasVolumeType() bool {
	if o != nil && o.VolumeType.IsSet() {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given NullableString and assigns it to the VolumeType field.
func (o *TanzuCredentialsForProjectDto) SetVolumeType(v string) {
	o.VolumeType.Set(&v)
}
// SetVolumeTypeNil sets the value for VolumeType to be an explicit nil
func (o *TanzuCredentialsForProjectDto) SetVolumeTypeNil() {
	o.VolumeType.Set(nil)
}

// UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
func (o *TanzuCredentialsForProjectDto) UnsetVolumeType() {
	o.VolumeType.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsForProjectDto) GetNamespace() string {
	if o == nil || IsNil(o.Namespace.Get()) {
		var ret string
		return ret
	}
	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsForProjectDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// HasNamespace returns a boolean if a field has been set.
func (o *TanzuCredentialsForProjectDto) HasNamespace() bool {
	if o != nil && o.Namespace.IsSet() {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given NullableString and assigns it to the Namespace field.
func (o *TanzuCredentialsForProjectDto) SetNamespace(v string) {
	o.Namespace.Set(&v)
}
// SetNamespaceNil sets the value for Namespace to be an explicit nil
func (o *TanzuCredentialsForProjectDto) SetNamespaceNil() {
	o.Namespace.Set(nil)
}

// UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
func (o *TanzuCredentialsForProjectDto) UnsetNamespace() {
	o.Namespace.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TanzuCredentialsForProjectDto) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsForProjectDto) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *TanzuCredentialsForProjectDto) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *TanzuCredentialsForProjectDto) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *TanzuCredentialsForProjectDto) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *TanzuCredentialsForProjectDto) UnsetPort() {
	o.Port.Unset()
}

func (o TanzuCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TanzuCredentialsForProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.VolumeType.IsSet() {
		toSerialize["volumeType"] = o.VolumeType.Get()
	}
	if o.Namespace.IsSet() {
		toSerialize["namespace"] = o.Namespace.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	return toSerialize, nil
}

type NullableTanzuCredentialsForProjectDto struct {
	value *TanzuCredentialsForProjectDto
	isSet bool
}

func (v NullableTanzuCredentialsForProjectDto) Get() *TanzuCredentialsForProjectDto {
	return v.value
}

func (v *NullableTanzuCredentialsForProjectDto) Set(val *TanzuCredentialsForProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTanzuCredentialsForProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTanzuCredentialsForProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTanzuCredentialsForProjectDto(val *TanzuCredentialsForProjectDto) *NullableTanzuCredentialsForProjectDto {
	return &NullableTanzuCredentialsForProjectDto{value: val, isSet: true}
}

func (v NullableTanzuCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTanzuCredentialsForProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


