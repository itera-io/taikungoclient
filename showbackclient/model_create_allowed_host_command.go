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

// checks if the CreateAllowedHostCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAllowedHostCommand{}

// CreateAllowedHostCommand struct for CreateAllowedHostCommand
type CreateAllowedHostCommand struct {
	AccessProfileId *int32         `json:"accessProfileId,omitempty"`
	Description     NullableString `json:"description,omitempty"`
	IpAddress       NullableString `json:"ipAddress,omitempty"`
	MaskBits        *int32         `json:"maskBits,omitempty"`
}

// NewCreateAllowedHostCommand instantiates a new CreateAllowedHostCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAllowedHostCommand() *CreateAllowedHostCommand {
	this := CreateAllowedHostCommand{}
	return &this
}

// NewCreateAllowedHostCommandWithDefaults instantiates a new CreateAllowedHostCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAllowedHostCommandWithDefaults() *CreateAllowedHostCommand {
	this := CreateAllowedHostCommand{}
	return &this
}

// GetAccessProfileId returns the AccessProfileId field value if set, zero value otherwise.
func (o *CreateAllowedHostCommand) GetAccessProfileId() int32 {
	if o == nil || IsNil(o.AccessProfileId) {
		var ret int32
		return ret
	}
	return *o.AccessProfileId
}

// GetAccessProfileIdOk returns a tuple with the AccessProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAllowedHostCommand) GetAccessProfileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessProfileId) {
		return nil, false
	}
	return o.AccessProfileId, true
}

// HasAccessProfileId returns a boolean if a field has been set.
func (o *CreateAllowedHostCommand) HasAccessProfileId() bool {
	if o != nil && !IsNil(o.AccessProfileId) {
		return true
	}

	return false
}

// SetAccessProfileId gets a reference to the given int32 and assigns it to the AccessProfileId field.
func (o *CreateAllowedHostCommand) SetAccessProfileId(v int32) {
	o.AccessProfileId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAllowedHostCommand) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAllowedHostCommand) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAllowedHostCommand) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CreateAllowedHostCommand) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CreateAllowedHostCommand) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CreateAllowedHostCommand) UnsetDescription() {
	o.Description.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAllowedHostCommand) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAllowedHostCommand) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CreateAllowedHostCommand) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *CreateAllowedHostCommand) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *CreateAllowedHostCommand) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *CreateAllowedHostCommand) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetMaskBits returns the MaskBits field value if set, zero value otherwise.
func (o *CreateAllowedHostCommand) GetMaskBits() int32 {
	if o == nil || IsNil(o.MaskBits) {
		var ret int32
		return ret
	}
	return *o.MaskBits
}

// GetMaskBitsOk returns a tuple with the MaskBits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAllowedHostCommand) GetMaskBitsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaskBits) {
		return nil, false
	}
	return o.MaskBits, true
}

// HasMaskBits returns a boolean if a field has been set.
func (o *CreateAllowedHostCommand) HasMaskBits() bool {
	if o != nil && !IsNil(o.MaskBits) {
		return true
	}

	return false
}

// SetMaskBits gets a reference to the given int32 and assigns it to the MaskBits field.
func (o *CreateAllowedHostCommand) SetMaskBits(v int32) {
	o.MaskBits = &v
}

func (o CreateAllowedHostCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAllowedHostCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessProfileId) {
		toSerialize["accessProfileId"] = o.AccessProfileId
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if !IsNil(o.MaskBits) {
		toSerialize["maskBits"] = o.MaskBits
	}
	return toSerialize, nil
}

type NullableCreateAllowedHostCommand struct {
	value *CreateAllowedHostCommand
	isSet bool
}

func (v NullableCreateAllowedHostCommand) Get() *CreateAllowedHostCommand {
	return v.value
}

func (v *NullableCreateAllowedHostCommand) Set(val *CreateAllowedHostCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAllowedHostCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAllowedHostCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAllowedHostCommand(val *CreateAllowedHostCommand) *NullableCreateAllowedHostCommand {
	return &NullableCreateAllowedHostCommand{value: val, isSet: true}
}

func (v NullableCreateAllowedHostCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAllowedHostCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
