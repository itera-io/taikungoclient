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

// checks if the EditSshUserCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditSshUserCommand{}

// EditSshUserCommand struct for EditSshUserCommand
type EditSshUserCommand struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SshPublicKey *string `json:"sshPublicKey,omitempty"`
	AccessProfileId *int32 `json:"accessProfileId,omitempty"`
}

// NewEditSshUserCommand instantiates a new EditSshUserCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditSshUserCommand() *EditSshUserCommand {
	this := EditSshUserCommand{}
	return &this
}

// NewEditSshUserCommandWithDefaults instantiates a new EditSshUserCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditSshUserCommandWithDefaults() *EditSshUserCommand {
	this := EditSshUserCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EditSshUserCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EditSshUserCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EditSshUserCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditSshUserCommand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditSshUserCommand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditSshUserCommand) SetName(v string) {
	o.Name = &v
}

// GetSshPublicKey returns the SshPublicKey field value if set, zero value otherwise.
func (o *EditSshUserCommand) GetSshPublicKey() string {
	if o == nil || IsNil(o.SshPublicKey) {
		var ret string
		return ret
	}
	return *o.SshPublicKey
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetSshPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SshPublicKey) {
		return nil, false
	}
	return o.SshPublicKey, true
}

// HasSshPublicKey returns a boolean if a field has been set.
func (o *EditSshUserCommand) HasSshPublicKey() bool {
	if o != nil && !IsNil(o.SshPublicKey) {
		return true
	}

	return false
}

// SetSshPublicKey gets a reference to the given string and assigns it to the SshPublicKey field.
func (o *EditSshUserCommand) SetSshPublicKey(v string) {
	o.SshPublicKey = &v
}

// GetAccessProfileId returns the AccessProfileId field value if set, zero value otherwise.
func (o *EditSshUserCommand) GetAccessProfileId() int32 {
	if o == nil || IsNil(o.AccessProfileId) {
		var ret int32
		return ret
	}
	return *o.AccessProfileId
}

// GetAccessProfileIdOk returns a tuple with the AccessProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditSshUserCommand) GetAccessProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessProfileId) {
		return nil, false
	}
	return o.AccessProfileId, true
}

// HasAccessProfileId returns a boolean if a field has been set.
func (o *EditSshUserCommand) HasAccessProfileId() bool {
	if o != nil && !IsNil(o.AccessProfileId) {
		return true
	}

	return false
}

// SetAccessProfileId gets a reference to the given int32 and assigns it to the AccessProfileId field.
func (o *EditSshUserCommand) SetAccessProfileId(v int32) {
	o.AccessProfileId = &v
}

func (o EditSshUserCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditSshUserCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SshPublicKey) {
		toSerialize["sshPublicKey"] = o.SshPublicKey
	}
	if !IsNil(o.AccessProfileId) {
		toSerialize["accessProfileId"] = o.AccessProfileId
	}
	return toSerialize, nil
}

type NullableEditSshUserCommand struct {
	value *EditSshUserCommand
	isSet bool
}

func (v NullableEditSshUserCommand) Get() *EditSshUserCommand {
	return v.value
}

func (v *NullableEditSshUserCommand) Set(val *EditSshUserCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditSshUserCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditSshUserCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditSshUserCommand(val *EditSshUserCommand) *NullableEditSshUserCommand {
	return &NullableEditSshUserCommand{value: val, isSet: true}
}

func (v NullableEditSshUserCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditSshUserCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


