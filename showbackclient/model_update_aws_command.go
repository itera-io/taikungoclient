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

// checks if the UpdateAwsCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAwsCommand{}

// UpdateAwsCommand struct for UpdateAwsCommand
type UpdateAwsCommand struct {
	Id                 *int32         `json:"id,omitempty"`
	Name               NullableString `json:"name,omitempty"`
	AwsSecretAccessKey NullableString `json:"awsSecretAccessKey,omitempty"`
	AwsAccessKeyId     NullableString `json:"awsAccessKeyId,omitempty"`
}

// NewUpdateAwsCommand instantiates a new UpdateAwsCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAwsCommand() *UpdateAwsCommand {
	this := UpdateAwsCommand{}
	return &this
}

// NewUpdateAwsCommandWithDefaults instantiates a new UpdateAwsCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAwsCommandWithDefaults() *UpdateAwsCommand {
	this := UpdateAwsCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateAwsCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAwsCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateAwsCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateAwsCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAwsCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAwsCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAwsCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateAwsCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateAwsCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateAwsCommand) UnsetName() {
	o.Name.Unset()
}

// GetAwsSecretAccessKey returns the AwsSecretAccessKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAwsCommand) GetAwsSecretAccessKey() string {
	if o == nil || IsNil(o.AwsSecretAccessKey.Get()) {
		var ret string
		return ret
	}
	return *o.AwsSecretAccessKey.Get()
}

// GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAwsCommand) GetAwsSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsSecretAccessKey.Get(), o.AwsSecretAccessKey.IsSet()
}

// HasAwsSecretAccessKey returns a boolean if a field has been set.
func (o *UpdateAwsCommand) HasAwsSecretAccessKey() bool {
	if o != nil && o.AwsSecretAccessKey.IsSet() {
		return true
	}

	return false
}

// SetAwsSecretAccessKey gets a reference to the given NullableString and assigns it to the AwsSecretAccessKey field.
func (o *UpdateAwsCommand) SetAwsSecretAccessKey(v string) {
	o.AwsSecretAccessKey.Set(&v)
}

// SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil
func (o *UpdateAwsCommand) SetAwsSecretAccessKeyNil() {
	o.AwsSecretAccessKey.Set(nil)
}

// UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
func (o *UpdateAwsCommand) UnsetAwsSecretAccessKey() {
	o.AwsSecretAccessKey.Unset()
}

// GetAwsAccessKeyId returns the AwsAccessKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateAwsCommand) GetAwsAccessKeyId() string {
	if o == nil || IsNil(o.AwsAccessKeyId.Get()) {
		var ret string
		return ret
	}
	return *o.AwsAccessKeyId.Get()
}

// GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateAwsCommand) GetAwsAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsAccessKeyId.Get(), o.AwsAccessKeyId.IsSet()
}

// HasAwsAccessKeyId returns a boolean if a field has been set.
func (o *UpdateAwsCommand) HasAwsAccessKeyId() bool {
	if o != nil && o.AwsAccessKeyId.IsSet() {
		return true
	}

	return false
}

// SetAwsAccessKeyId gets a reference to the given NullableString and assigns it to the AwsAccessKeyId field.
func (o *UpdateAwsCommand) SetAwsAccessKeyId(v string) {
	o.AwsAccessKeyId.Set(&v)
}

// SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil
func (o *UpdateAwsCommand) SetAwsAccessKeyIdNil() {
	o.AwsAccessKeyId.Set(nil)
}

// UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
func (o *UpdateAwsCommand) UnsetAwsAccessKeyId() {
	o.AwsAccessKeyId.Unset()
}

func (o UpdateAwsCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAwsCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AwsSecretAccessKey.IsSet() {
		toSerialize["awsSecretAccessKey"] = o.AwsSecretAccessKey.Get()
	}
	if o.AwsAccessKeyId.IsSet() {
		toSerialize["awsAccessKeyId"] = o.AwsAccessKeyId.Get()
	}
	return toSerialize, nil
}

type NullableUpdateAwsCommand struct {
	value *UpdateAwsCommand
	isSet bool
}

func (v NullableUpdateAwsCommand) Get() *UpdateAwsCommand {
	return v.value
}

func (v *NullableUpdateAwsCommand) Set(val *UpdateAwsCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAwsCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAwsCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAwsCommand(val *UpdateAwsCommand) *NullableUpdateAwsCommand {
	return &NullableUpdateAwsCommand{value: val, isSet: true}
}

func (v NullableUpdateAwsCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAwsCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
