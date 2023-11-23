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

// checks if the UserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDto{}

// UserDto struct for UserDto
type UserDto struct {
	UserId NullableString `json:"userId,omitempty"`
	UserName NullableString `json:"userName,omitempty"`
}

// NewUserDto instantiates a new UserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDto() *UserDto {
	this := UserDto{}
	return &this
}

// NewUserDtoWithDefaults instantiates a new UserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDtoWithDefaults() *UserDto {
	this := UserDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *UserDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *UserDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *UserDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *UserDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *UserDto) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *UserDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *UserDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *UserDto) UnsetUserName() {
	o.UserName.Unset()
}

func (o UserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	return toSerialize, nil
}

type NullableUserDto struct {
	value *UserDto
	isSet bool
}

func (v NullableUserDto) Get() *UserDto {
	return v.value
}

func (v *NullableUserDto) Set(val *UserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDto(val *UserDto) *NullableUserDto {
	return &NullableUserDto{value: val, isSet: true}
}

func (v NullableUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


