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

// checks if the StandAloneProfileSecurityGroupDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfileSecurityGroupDto{}

// StandAloneProfileSecurityGroupDto struct for StandAloneProfileSecurityGroupDto
type StandAloneProfileSecurityGroupDto struct {
	Name NullableString `json:"name,omitempty"`
	Protocol *SecurityGroupProtocol `json:"protocol,omitempty"`
	PortMinRange *int32 `json:"portMinRange,omitempty"`
	PortMaxRange *int32 `json:"portMaxRange,omitempty"`
	RemoteIpPrefix NullableString `json:"remoteIpPrefix,omitempty"`
}

// NewStandAloneProfileSecurityGroupDto instantiates a new StandAloneProfileSecurityGroupDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfileSecurityGroupDto() *StandAloneProfileSecurityGroupDto {
	this := StandAloneProfileSecurityGroupDto{}
	return &this
}

// NewStandAloneProfileSecurityGroupDtoWithDefaults instantiates a new StandAloneProfileSecurityGroupDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfileSecurityGroupDtoWithDefaults() *StandAloneProfileSecurityGroupDto {
	this := StandAloneProfileSecurityGroupDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileSecurityGroupDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileSecurityGroupDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StandAloneProfileSecurityGroupDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *StandAloneProfileSecurityGroupDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StandAloneProfileSecurityGroupDto) UnsetName() {
	o.Name.Unset()
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupDto) GetProtocol() SecurityGroupProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret SecurityGroupProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupDto) GetProtocolOk() (*SecurityGroupProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupDto) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given SecurityGroupProtocol and assigns it to the Protocol field.
func (o *StandAloneProfileSecurityGroupDto) SetProtocol(v SecurityGroupProtocol) {
	o.Protocol = &v
}

// GetPortMinRange returns the PortMinRange field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupDto) GetPortMinRange() int32 {
	if o == nil || IsNil(o.PortMinRange) {
		var ret int32
		return ret
	}
	return *o.PortMinRange
}

// GetPortMinRangeOk returns a tuple with the PortMinRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupDto) GetPortMinRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMinRange) {
		return nil, false
	}
	return o.PortMinRange, true
}

// HasPortMinRange returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupDto) HasPortMinRange() bool {
	if o != nil && !IsNil(o.PortMinRange) {
		return true
	}

	return false
}

// SetPortMinRange gets a reference to the given int32 and assigns it to the PortMinRange field.
func (o *StandAloneProfileSecurityGroupDto) SetPortMinRange(v int32) {
	o.PortMinRange = &v
}

// GetPortMaxRange returns the PortMaxRange field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupDto) GetPortMaxRange() int32 {
	if o == nil || IsNil(o.PortMaxRange) {
		var ret int32
		return ret
	}
	return *o.PortMaxRange
}

// GetPortMaxRangeOk returns a tuple with the PortMaxRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupDto) GetPortMaxRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMaxRange) {
		return nil, false
	}
	return o.PortMaxRange, true
}

// HasPortMaxRange returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupDto) HasPortMaxRange() bool {
	if o != nil && !IsNil(o.PortMaxRange) {
		return true
	}

	return false
}

// SetPortMaxRange gets a reference to the given int32 and assigns it to the PortMaxRange field.
func (o *StandAloneProfileSecurityGroupDto) SetPortMaxRange(v int32) {
	o.PortMaxRange = &v
}

// GetRemoteIpPrefix returns the RemoteIpPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileSecurityGroupDto) GetRemoteIpPrefix() string {
	if o == nil || IsNil(o.RemoteIpPrefix.Get()) {
		var ret string
		return ret
	}
	return *o.RemoteIpPrefix.Get()
}

// GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileSecurityGroupDto) GetRemoteIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteIpPrefix.Get(), o.RemoteIpPrefix.IsSet()
}

// HasRemoteIpPrefix returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupDto) HasRemoteIpPrefix() bool {
	if o != nil && o.RemoteIpPrefix.IsSet() {
		return true
	}

	return false
}

// SetRemoteIpPrefix gets a reference to the given NullableString and assigns it to the RemoteIpPrefix field.
func (o *StandAloneProfileSecurityGroupDto) SetRemoteIpPrefix(v string) {
	o.RemoteIpPrefix.Set(&v)
}
// SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil
func (o *StandAloneProfileSecurityGroupDto) SetRemoteIpPrefixNil() {
	o.RemoteIpPrefix.Set(nil)
}

// UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
func (o *StandAloneProfileSecurityGroupDto) UnsetRemoteIpPrefix() {
	o.RemoteIpPrefix.Unset()
}

func (o StandAloneProfileSecurityGroupDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfileSecurityGroupDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.PortMinRange) {
		toSerialize["portMinRange"] = o.PortMinRange
	}
	if !IsNil(o.PortMaxRange) {
		toSerialize["portMaxRange"] = o.PortMaxRange
	}
	if o.RemoteIpPrefix.IsSet() {
		toSerialize["remoteIpPrefix"] = o.RemoteIpPrefix.Get()
	}
	return toSerialize, nil
}

type NullableStandAloneProfileSecurityGroupDto struct {
	value *StandAloneProfileSecurityGroupDto
	isSet bool
}

func (v NullableStandAloneProfileSecurityGroupDto) Get() *StandAloneProfileSecurityGroupDto {
	return v.value
}

func (v *NullableStandAloneProfileSecurityGroupDto) Set(val *StandAloneProfileSecurityGroupDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfileSecurityGroupDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfileSecurityGroupDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfileSecurityGroupDto(val *StandAloneProfileSecurityGroupDto) *NullableStandAloneProfileSecurityGroupDto {
	return &NullableStandAloneProfileSecurityGroupDto{value: val, isSet: true}
}

func (v NullableStandAloneProfileSecurityGroupDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfileSecurityGroupDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


