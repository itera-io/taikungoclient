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
	"bytes"
	"fmt"
)

// checks if the TransferList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferList{}

// TransferList struct for TransferList
type TransferList struct {
	UserId NullableString `json:"userId"`
	UserName NullableString `json:"userName"`
	Role UserRole `json:"role"`
}

type _TransferList TransferList

// NewTransferList instantiates a new TransferList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferList(userId NullableString, userName NullableString, role UserRole) *TransferList {
	this := TransferList{}
	this.UserId = userId
	this.UserName = userName
	this.Role = role
	return &this
}

// NewTransferListWithDefaults instantiates a new TransferList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferListWithDefaults() *TransferList {
	this := TransferList{}
	return &this
}

// GetUserId returns the UserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferList) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferList) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// SetUserId sets field value
func (o *TransferList) SetUserId(v string) {
	o.UserId.Set(&v)
}

// GetUserName returns the UserName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferList) GetUserName() string {
	if o == nil || o.UserName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferList) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// SetUserName sets field value
func (o *TransferList) SetUserName(v string) {
	o.UserName.Set(&v)
}

// GetRole returns the Role field value
func (o *TransferList) GetRole() UserRole {
	if o == nil {
		var ret UserRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *TransferList) GetRoleOk() (*UserRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *TransferList) SetRole(v UserRole) {
	o.Role = v
}

func (o TransferList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId.Get()
	toSerialize["userName"] = o.UserName.Get()
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *TransferList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"userName",
		"role",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTransferList := _TransferList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferList)

	if err != nil {
		return err
	}

	*o = TransferList(varTransferList)

	return err
}

type NullableTransferList struct {
	value *TransferList
	isSet bool
}

func (v NullableTransferList) Get() *TransferList {
	return v.value
}

func (v *NullableTransferList) Set(val *TransferList) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferList) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferList(val *TransferList) *NullableTransferList {
	return &NullableTransferList{value: val, isSet: true}
}

func (v NullableTransferList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


