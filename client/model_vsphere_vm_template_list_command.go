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

// checks if the VsphereVmTemplateListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsphereVmTemplateListCommand{}

// VsphereVmTemplateListCommand struct for VsphereVmTemplateListCommand
type VsphereVmTemplateListCommand struct {
	Url NullableString `json:"url,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	Datacenter NullableString `json:"datacenter,omitempty"`
}

// NewVsphereVmTemplateListCommand instantiates a new VsphereVmTemplateListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsphereVmTemplateListCommand() *VsphereVmTemplateListCommand {
	this := VsphereVmTemplateListCommand{}
	return &this
}

// NewVsphereVmTemplateListCommandWithDefaults instantiates a new VsphereVmTemplateListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsphereVmTemplateListCommandWithDefaults() *VsphereVmTemplateListCommand {
	this := VsphereVmTemplateListCommand{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereVmTemplateListCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereVmTemplateListCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *VsphereVmTemplateListCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *VsphereVmTemplateListCommand) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *VsphereVmTemplateListCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *VsphereVmTemplateListCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereVmTemplateListCommand) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereVmTemplateListCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *VsphereVmTemplateListCommand) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *VsphereVmTemplateListCommand) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *VsphereVmTemplateListCommand) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *VsphereVmTemplateListCommand) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereVmTemplateListCommand) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereVmTemplateListCommand) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *VsphereVmTemplateListCommand) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *VsphereVmTemplateListCommand) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *VsphereVmTemplateListCommand) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *VsphereVmTemplateListCommand) UnsetPassword() {
	o.Password.Unset()
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereVmTemplateListCommand) GetDatacenter() string {
	if o == nil || IsNil(o.Datacenter.Get()) {
		var ret string
		return ret
	}
	return *o.Datacenter.Get()
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereVmTemplateListCommand) GetDatacenterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datacenter.Get(), o.Datacenter.IsSet()
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VsphereVmTemplateListCommand) HasDatacenter() bool {
	if o != nil && o.Datacenter.IsSet() {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given NullableString and assigns it to the Datacenter field.
func (o *VsphereVmTemplateListCommand) SetDatacenter(v string) {
	o.Datacenter.Set(&v)
}
// SetDatacenterNil sets the value for Datacenter to be an explicit nil
func (o *VsphereVmTemplateListCommand) SetDatacenterNil() {
	o.Datacenter.Set(nil)
}

// UnsetDatacenter ensures that no value is present for Datacenter, not even an explicit nil
func (o *VsphereVmTemplateListCommand) UnsetDatacenter() {
	o.Datacenter.Unset()
}

func (o VsphereVmTemplateListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VsphereVmTemplateListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Datacenter.IsSet() {
		toSerialize["datacenter"] = o.Datacenter.Get()
	}
	return toSerialize, nil
}

type NullableVsphereVmTemplateListCommand struct {
	value *VsphereVmTemplateListCommand
	isSet bool
}

func (v NullableVsphereVmTemplateListCommand) Get() *VsphereVmTemplateListCommand {
	return v.value
}

func (v *NullableVsphereVmTemplateListCommand) Set(val *VsphereVmTemplateListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableVsphereVmTemplateListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableVsphereVmTemplateListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsphereVmTemplateListCommand(val *VsphereVmTemplateListCommand) *NullableVsphereVmTemplateListCommand {
	return &NullableVsphereVmTemplateListCommand{value: val, isSet: true}
}

func (v NullableVsphereVmTemplateListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsphereVmTemplateListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

