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

// checks if the ImportRepoCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRepoCommand{}

// ImportRepoCommand struct for ImportRepoCommand
type ImportRepoCommand struct {
	Name        NullableString `json:"name,omitempty"`
	Url         NullableString `json:"url,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewImportRepoCommand instantiates a new ImportRepoCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRepoCommand() *ImportRepoCommand {
	this := ImportRepoCommand{}
	return &this
}

// NewImportRepoCommandWithDefaults instantiates a new ImportRepoCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRepoCommandWithDefaults() *ImportRepoCommand {
	this := ImportRepoCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportRepoCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportRepoCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ImportRepoCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ImportRepoCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ImportRepoCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ImportRepoCommand) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportRepoCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportRepoCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ImportRepoCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ImportRepoCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *ImportRepoCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ImportRepoCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportRepoCommand) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportRepoCommand) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ImportRepoCommand) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ImportRepoCommand) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ImportRepoCommand) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ImportRepoCommand) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o ImportRepoCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRepoCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return toSerialize, nil
}

type NullableImportRepoCommand struct {
	value *ImportRepoCommand
	isSet bool
}

func (v NullableImportRepoCommand) Get() *ImportRepoCommand {
	return v.value
}

func (v *NullableImportRepoCommand) Set(val *ImportRepoCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRepoCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRepoCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRepoCommand(val *ImportRepoCommand) *NullableImportRepoCommand {
	return &NullableImportRepoCommand{value: val, isSet: true}
}

func (v NullableImportRepoCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRepoCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
