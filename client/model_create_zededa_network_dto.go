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

// checks if the CreateZededaNetworkDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateZededaNetworkDto{}

// CreateZededaNetworkDto struct for CreateZededaNetworkDto
type CreateZededaNetworkDto struct {
	Interface NullableString `json:"interface,omitempty"`
	Gateway NullableString `json:"gateway,omitempty"`
	IpAddress NullableString `json:"ipAddress,omitempty"`
	NetMask *int32 `json:"netMask,omitempty"`
	VlanId NullableInt32 `json:"vlanId,omitempty"`
	BeginAllocationRange NullableString `json:"beginAllocationRange,omitempty"`
	EndAllocationRange NullableString `json:"endAllocationRange,omitempty"`
}

// NewCreateZededaNetworkDto instantiates a new CreateZededaNetworkDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateZededaNetworkDto() *CreateZededaNetworkDto {
	this := CreateZededaNetworkDto{}
	return &this
}

// NewCreateZededaNetworkDtoWithDefaults instantiates a new CreateZededaNetworkDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateZededaNetworkDtoWithDefaults() *CreateZededaNetworkDto {
	this := CreateZededaNetworkDto{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaNetworkDto) GetInterface() string {
	if o == nil || IsNil(o.Interface.Get()) {
		var ret string
		return ret
	}
	return *o.Interface.Get()
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaNetworkDto) GetInterfaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interface.Get(), o.Interface.IsSet()
}

// HasInterface returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasInterface() bool {
	if o != nil && o.Interface.IsSet() {
		return true
	}

	return false
}

// SetInterface gets a reference to the given NullableString and assigns it to the Interface field.
func (o *CreateZededaNetworkDto) SetInterface(v string) {
	o.Interface.Set(&v)
}
// SetInterfaceNil sets the value for Interface to be an explicit nil
func (o *CreateZededaNetworkDto) SetInterfaceNil() {
	o.Interface.Set(nil)
}

// UnsetInterface ensures that no value is present for Interface, not even an explicit nil
func (o *CreateZededaNetworkDto) UnsetInterface() {
	o.Interface.Unset()
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaNetworkDto) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaNetworkDto) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *CreateZededaNetworkDto) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *CreateZededaNetworkDto) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *CreateZededaNetworkDto) UnsetGateway() {
	o.Gateway.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaNetworkDto) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaNetworkDto) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *CreateZededaNetworkDto) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *CreateZededaNetworkDto) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *CreateZededaNetworkDto) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetNetMask returns the NetMask field value if set, zero value otherwise.
func (o *CreateZededaNetworkDto) GetNetMask() int32 {
	if o == nil || IsNil(o.NetMask) {
		var ret int32
		return ret
	}
	return *o.NetMask
}

// GetNetMaskOk returns a tuple with the NetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZededaNetworkDto) GetNetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.NetMask) {
		return nil, false
	}
	return o.NetMask, true
}

// HasNetMask returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasNetMask() bool {
	if o != nil && !IsNil(o.NetMask) {
		return true
	}

	return false
}

// SetNetMask gets a reference to the given int32 and assigns it to the NetMask field.
func (o *CreateZededaNetworkDto) SetNetMask(v int32) {
	o.NetMask = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaNetworkDto) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId.Get()) {
		var ret int32
		return ret
	}
	return *o.VlanId.Get()
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaNetworkDto) GetVlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanId.Get(), o.VlanId.IsSet()
}

// HasVlanId returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasVlanId() bool {
	if o != nil && o.VlanId.IsSet() {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given NullableInt32 and assigns it to the VlanId field.
func (o *CreateZededaNetworkDto) SetVlanId(v int32) {
	o.VlanId.Set(&v)
}
// SetVlanIdNil sets the value for VlanId to be an explicit nil
func (o *CreateZededaNetworkDto) SetVlanIdNil() {
	o.VlanId.Set(nil)
}

// UnsetVlanId ensures that no value is present for VlanId, not even an explicit nil
func (o *CreateZededaNetworkDto) UnsetVlanId() {
	o.VlanId.Unset()
}

// GetBeginAllocationRange returns the BeginAllocationRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaNetworkDto) GetBeginAllocationRange() string {
	if o == nil || IsNil(o.BeginAllocationRange.Get()) {
		var ret string
		return ret
	}
	return *o.BeginAllocationRange.Get()
}

// GetBeginAllocationRangeOk returns a tuple with the BeginAllocationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaNetworkDto) GetBeginAllocationRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BeginAllocationRange.Get(), o.BeginAllocationRange.IsSet()
}

// HasBeginAllocationRange returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasBeginAllocationRange() bool {
	if o != nil && o.BeginAllocationRange.IsSet() {
		return true
	}

	return false
}

// SetBeginAllocationRange gets a reference to the given NullableString and assigns it to the BeginAllocationRange field.
func (o *CreateZededaNetworkDto) SetBeginAllocationRange(v string) {
	o.BeginAllocationRange.Set(&v)
}
// SetBeginAllocationRangeNil sets the value for BeginAllocationRange to be an explicit nil
func (o *CreateZededaNetworkDto) SetBeginAllocationRangeNil() {
	o.BeginAllocationRange.Set(nil)
}

// UnsetBeginAllocationRange ensures that no value is present for BeginAllocationRange, not even an explicit nil
func (o *CreateZededaNetworkDto) UnsetBeginAllocationRange() {
	o.BeginAllocationRange.Unset()
}

// GetEndAllocationRange returns the EndAllocationRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaNetworkDto) GetEndAllocationRange() string {
	if o == nil || IsNil(o.EndAllocationRange.Get()) {
		var ret string
		return ret
	}
	return *o.EndAllocationRange.Get()
}

// GetEndAllocationRangeOk returns a tuple with the EndAllocationRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaNetworkDto) GetEndAllocationRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAllocationRange.Get(), o.EndAllocationRange.IsSet()
}

// HasEndAllocationRange returns a boolean if a field has been set.
func (o *CreateZededaNetworkDto) HasEndAllocationRange() bool {
	if o != nil && o.EndAllocationRange.IsSet() {
		return true
	}

	return false
}

// SetEndAllocationRange gets a reference to the given NullableString and assigns it to the EndAllocationRange field.
func (o *CreateZededaNetworkDto) SetEndAllocationRange(v string) {
	o.EndAllocationRange.Set(&v)
}
// SetEndAllocationRangeNil sets the value for EndAllocationRange to be an explicit nil
func (o *CreateZededaNetworkDto) SetEndAllocationRangeNil() {
	o.EndAllocationRange.Set(nil)
}

// UnsetEndAllocationRange ensures that no value is present for EndAllocationRange, not even an explicit nil
func (o *CreateZededaNetworkDto) UnsetEndAllocationRange() {
	o.EndAllocationRange.Unset()
}

func (o CreateZededaNetworkDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateZededaNetworkDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Interface.IsSet() {
		toSerialize["interface"] = o.Interface.Get()
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if !IsNil(o.NetMask) {
		toSerialize["netMask"] = o.NetMask
	}
	if o.VlanId.IsSet() {
		toSerialize["vlanId"] = o.VlanId.Get()
	}
	if o.BeginAllocationRange.IsSet() {
		toSerialize["beginAllocationRange"] = o.BeginAllocationRange.Get()
	}
	if o.EndAllocationRange.IsSet() {
		toSerialize["endAllocationRange"] = o.EndAllocationRange.Get()
	}
	return toSerialize, nil
}

type NullableCreateZededaNetworkDto struct {
	value *CreateZededaNetworkDto
	isSet bool
}

func (v NullableCreateZededaNetworkDto) Get() *CreateZededaNetworkDto {
	return v.value
}

func (v *NullableCreateZededaNetworkDto) Set(val *CreateZededaNetworkDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateZededaNetworkDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateZededaNetworkDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateZededaNetworkDto(val *CreateZededaNetworkDto) *NullableCreateZededaNetworkDto {
	return &NullableCreateZededaNetworkDto{value: val, isSet: true}
}

func (v NullableCreateZededaNetworkDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateZededaNetworkDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


