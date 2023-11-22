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

// checks if the StandAloneProfileSecurityGroupForDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfileSecurityGroupForDetailsDto{}

// StandAloneProfileSecurityGroupForDetailsDto struct for StandAloneProfileSecurityGroupForDetailsDto
type StandAloneProfileSecurityGroupForDetailsDto struct {
	Id               *int32         `json:"id,omitempty"`
	Name             NullableString `json:"name,omitempty"`
	Protocol         NullableString `json:"protocol,omitempty"`
	PortMinRange     *int32         `json:"portMinRange,omitempty"`
	PortMaxRange     *int32         `json:"portMaxRange,omitempty"`
	RemoteIpPrefix   NullableString `json:"remoteIpPrefix,omitempty"`
	IsRdpPortEnabled *bool          `json:"isRdpPortEnabled,omitempty"`
}

// NewStandAloneProfileSecurityGroupForDetailsDto instantiates a new StandAloneProfileSecurityGroupForDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfileSecurityGroupForDetailsDto() *StandAloneProfileSecurityGroupForDetailsDto {
	this := StandAloneProfileSecurityGroupForDetailsDto{}
	return &this
}

// NewStandAloneProfileSecurityGroupForDetailsDtoWithDefaults instantiates a new StandAloneProfileSecurityGroupForDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfileSecurityGroupForDetailsDtoWithDefaults() *StandAloneProfileSecurityGroupForDetailsDto {
	this := StandAloneProfileSecurityGroupForDetailsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StandAloneProfileSecurityGroupForDetailsDto) UnsetName() {
	o.Name.Unset()
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetProtocol(v string) {
	o.Protocol.Set(&v)
}

// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *StandAloneProfileSecurityGroupForDetailsDto) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetPortMinRange returns the PortMinRange field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMinRange() int32 {
	if o == nil || IsNil(o.PortMinRange) {
		var ret int32
		return ret
	}
	return *o.PortMinRange
}

// GetPortMinRangeOk returns a tuple with the PortMinRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMinRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMinRange) {
		return nil, false
	}
	return o.PortMinRange, true
}

// HasPortMinRange returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasPortMinRange() bool {
	if o != nil && !IsNil(o.PortMinRange) {
		return true
	}

	return false
}

// SetPortMinRange gets a reference to the given int32 and assigns it to the PortMinRange field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetPortMinRange(v int32) {
	o.PortMinRange = &v
}

// GetPortMaxRange returns the PortMaxRange field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMaxRange() int32 {
	if o == nil || IsNil(o.PortMaxRange) {
		var ret int32
		return ret
	}
	return *o.PortMaxRange
}

// GetPortMaxRangeOk returns a tuple with the PortMaxRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetPortMaxRangeOk() (*int32, bool) {
	if o == nil || IsNil(o.PortMaxRange) {
		return nil, false
	}
	return o.PortMaxRange, true
}

// HasPortMaxRange returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasPortMaxRange() bool {
	if o != nil && !IsNil(o.PortMaxRange) {
		return true
	}

	return false
}

// SetPortMaxRange gets a reference to the given int32 and assigns it to the PortMaxRange field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetPortMaxRange(v int32) {
	o.PortMaxRange = &v
}

// GetRemoteIpPrefix returns the RemoteIpPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetRemoteIpPrefix() string {
	if o == nil || IsNil(o.RemoteIpPrefix.Get()) {
		var ret string
		return ret
	}
	return *o.RemoteIpPrefix.Get()
}

// GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetRemoteIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteIpPrefix.Get(), o.RemoteIpPrefix.IsSet()
}

// HasRemoteIpPrefix returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasRemoteIpPrefix() bool {
	if o != nil && o.RemoteIpPrefix.IsSet() {
		return true
	}

	return false
}

// SetRemoteIpPrefix gets a reference to the given NullableString and assigns it to the RemoteIpPrefix field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetRemoteIpPrefix(v string) {
	o.RemoteIpPrefix.Set(&v)
}

// SetRemoteIpPrefixNil sets the value for RemoteIpPrefix to be an explicit nil
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetRemoteIpPrefixNil() {
	o.RemoteIpPrefix.Set(nil)
}

// UnsetRemoteIpPrefix ensures that no value is present for RemoteIpPrefix, not even an explicit nil
func (o *StandAloneProfileSecurityGroupForDetailsDto) UnsetRemoteIpPrefix() {
	o.RemoteIpPrefix.Unset()
}

// GetIsRdpPortEnabled returns the IsRdpPortEnabled field value if set, zero value otherwise.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetIsRdpPortEnabled() bool {
	if o == nil || IsNil(o.IsRdpPortEnabled) {
		var ret bool
		return ret
	}
	return *o.IsRdpPortEnabled
}

// GetIsRdpPortEnabledOk returns a tuple with the IsRdpPortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) GetIsRdpPortEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRdpPortEnabled) {
		return nil, false
	}
	return o.IsRdpPortEnabled, true
}

// HasIsRdpPortEnabled returns a boolean if a field has been set.
func (o *StandAloneProfileSecurityGroupForDetailsDto) HasIsRdpPortEnabled() bool {
	if o != nil && !IsNil(o.IsRdpPortEnabled) {
		return true
	}

	return false
}

// SetIsRdpPortEnabled gets a reference to the given bool and assigns it to the IsRdpPortEnabled field.
func (o *StandAloneProfileSecurityGroupForDetailsDto) SetIsRdpPortEnabled(v bool) {
	o.IsRdpPortEnabled = &v
}

func (o StandAloneProfileSecurityGroupForDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfileSecurityGroupForDetailsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
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
	if !IsNil(o.IsRdpPortEnabled) {
		toSerialize["isRdpPortEnabled"] = o.IsRdpPortEnabled
	}
	return toSerialize, nil
}

type NullableStandAloneProfileSecurityGroupForDetailsDto struct {
	value *StandAloneProfileSecurityGroupForDetailsDto
	isSet bool
}

func (v NullableStandAloneProfileSecurityGroupForDetailsDto) Get() *StandAloneProfileSecurityGroupForDetailsDto {
	return v.value
}

func (v *NullableStandAloneProfileSecurityGroupForDetailsDto) Set(val *StandAloneProfileSecurityGroupForDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfileSecurityGroupForDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfileSecurityGroupForDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfileSecurityGroupForDetailsDto(val *StandAloneProfileSecurityGroupForDetailsDto) *NullableStandAloneProfileSecurityGroupForDetailsDto {
	return &NullableStandAloneProfileSecurityGroupForDetailsDto{value: val, isSet: true}
}

func (v NullableStandAloneProfileSecurityGroupForDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfileSecurityGroupForDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
