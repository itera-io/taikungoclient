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

// checks if the ProxmoxHypervisorDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProxmoxHypervisorDto{}

// ProxmoxHypervisorDto struct for ProxmoxHypervisorDto
type ProxmoxHypervisorDto struct {
	Name NullableString `json:"name,omitempty"`
	IsBound *bool `json:"isBound,omitempty"`
	UsedByServer *bool `json:"usedByServer,omitempty"`
}

// NewProxmoxHypervisorDto instantiates a new ProxmoxHypervisorDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxmoxHypervisorDto() *ProxmoxHypervisorDto {
	this := ProxmoxHypervisorDto{}
	return &this
}

// NewProxmoxHypervisorDtoWithDefaults instantiates a new ProxmoxHypervisorDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxmoxHypervisorDtoWithDefaults() *ProxmoxHypervisorDto {
	this := ProxmoxHypervisorDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProxmoxHypervisorDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProxmoxHypervisorDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ProxmoxHypervisorDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ProxmoxHypervisorDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ProxmoxHypervisorDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ProxmoxHypervisorDto) UnsetName() {
	o.Name.Unset()
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *ProxmoxHypervisorDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxHypervisorDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *ProxmoxHypervisorDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *ProxmoxHypervisorDto) SetIsBound(v bool) {
	o.IsBound = &v
}

// GetUsedByServer returns the UsedByServer field value if set, zero value otherwise.
func (o *ProxmoxHypervisorDto) GetUsedByServer() bool {
	if o == nil || IsNil(o.UsedByServer) {
		var ret bool
		return ret
	}
	return *o.UsedByServer
}

// GetUsedByServerOk returns a tuple with the UsedByServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxmoxHypervisorDto) GetUsedByServerOk() (*bool, bool) {
	if o == nil || IsNil(o.UsedByServer) {
		return nil, false
	}
	return o.UsedByServer, true
}

// HasUsedByServer returns a boolean if a field has been set.
func (o *ProxmoxHypervisorDto) HasUsedByServer() bool {
	if o != nil && !IsNil(o.UsedByServer) {
		return true
	}

	return false
}

// SetUsedByServer gets a reference to the given bool and assigns it to the UsedByServer field.
func (o *ProxmoxHypervisorDto) SetUsedByServer(v bool) {
	o.UsedByServer = &v
}

func (o ProxmoxHypervisorDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProxmoxHypervisorDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	if !IsNil(o.UsedByServer) {
		toSerialize["usedByServer"] = o.UsedByServer
	}
	return toSerialize, nil
}

type NullableProxmoxHypervisorDto struct {
	value *ProxmoxHypervisorDto
	isSet bool
}

func (v NullableProxmoxHypervisorDto) Get() *ProxmoxHypervisorDto {
	return v.value
}

func (v *NullableProxmoxHypervisorDto) Set(val *ProxmoxHypervisorDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProxmoxHypervisorDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProxmoxHypervisorDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxmoxHypervisorDto(val *ProxmoxHypervisorDto) *NullableProxmoxHypervisorDto {
	return &NullableProxmoxHypervisorDto{value: val, isSet: true}
}

func (v NullableProxmoxHypervisorDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxmoxHypervisorDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


