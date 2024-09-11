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

// checks if the ServerTemplateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerTemplateDto{}

// ServerTemplateDto struct for ServerTemplateDto
type ServerTemplateDto struct {
	Role *CloudRole `json:"role,omitempty"`
	Flavor NullableString `json:"flavor,omitempty"`
	DiskSize *float64 `json:"diskSize,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewServerTemplateDto instantiates a new ServerTemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerTemplateDto() *ServerTemplateDto {
	this := ServerTemplateDto{}
	return &this
}

// NewServerTemplateDtoWithDefaults instantiates a new ServerTemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerTemplateDtoWithDefaults() *ServerTemplateDto {
	this := ServerTemplateDto{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ServerTemplateDto) GetRole() CloudRole {
	if o == nil || IsNil(o.Role) {
		var ret CloudRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerTemplateDto) GetRoleOk() (*CloudRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ServerTemplateDto) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given CloudRole and assigns it to the Role field.
func (o *ServerTemplateDto) SetRole(v CloudRole) {
	o.Role = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerTemplateDto) GetFlavor() string {
	if o == nil || IsNil(o.Flavor.Get()) {
		var ret string
		return ret
	}
	return *o.Flavor.Get()
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerTemplateDto) GetFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flavor.Get(), o.Flavor.IsSet()
}

// HasFlavor returns a boolean if a field has been set.
func (o *ServerTemplateDto) HasFlavor() bool {
	if o != nil && o.Flavor.IsSet() {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given NullableString and assigns it to the Flavor field.
func (o *ServerTemplateDto) SetFlavor(v string) {
	o.Flavor.Set(&v)
}
// SetFlavorNil sets the value for Flavor to be an explicit nil
func (o *ServerTemplateDto) SetFlavorNil() {
	o.Flavor.Set(nil)
}

// UnsetFlavor ensures that no value is present for Flavor, not even an explicit nil
func (o *ServerTemplateDto) UnsetFlavor() {
	o.Flavor.Unset()
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *ServerTemplateDto) GetDiskSize() float64 {
	if o == nil || IsNil(o.DiskSize) {
		var ret float64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerTemplateDto) GetDiskSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *ServerTemplateDto) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given float64 and assigns it to the DiskSize field.
func (o *ServerTemplateDto) SetDiskSize(v float64) {
	o.DiskSize = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ServerTemplateDto) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerTemplateDto) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ServerTemplateDto) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ServerTemplateDto) SetCount(v int32) {
	o.Count = &v
}

func (o ServerTemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerTemplateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.Flavor.IsSet() {
		toSerialize["flavor"] = o.Flavor.Get()
	}
	if !IsNil(o.DiskSize) {
		toSerialize["diskSize"] = o.DiskSize
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableServerTemplateDto struct {
	value *ServerTemplateDto
	isSet bool
}

func (v NullableServerTemplateDto) Get() *ServerTemplateDto {
	return v.value
}

func (v *NullableServerTemplateDto) Set(val *ServerTemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServerTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServerTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerTemplateDto(val *ServerTemplateDto) *NullableServerTemplateDto {
	return &NullableServerTemplateDto{value: val, isSet: true}
}

func (v NullableServerTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


