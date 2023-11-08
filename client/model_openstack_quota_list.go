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

// checks if the OpenstackQuotaList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackQuotaList{}

// OpenstackQuotaList struct for OpenstackQuotaList
type OpenstackQuotaList struct {
	Compute *OpenstackComputeQuotaDto `json:"compute,omitempty"`
	Volume  *OpenstackVolumeQuotaDto  `json:"volume,omitempty"`
	Network *OpenstackNetworkDto      `json:"network,omitempty"`
	IsInfra *bool                     `json:"isInfra,omitempty"`
}

// NewOpenstackQuotaList instantiates a new OpenstackQuotaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackQuotaList() *OpenstackQuotaList {
	this := OpenstackQuotaList{}
	return &this
}

// NewOpenstackQuotaListWithDefaults instantiates a new OpenstackQuotaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackQuotaListWithDefaults() *OpenstackQuotaList {
	this := OpenstackQuotaList{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise.
func (o *OpenstackQuotaList) GetCompute() OpenstackComputeQuotaDto {
	if o == nil || IsNil(o.Compute) {
		var ret OpenstackComputeQuotaDto
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackQuotaList) GetComputeOk() (*OpenstackComputeQuotaDto, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}
	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *OpenstackQuotaList) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given OpenstackComputeQuotaDto and assigns it to the Compute field.
func (o *OpenstackQuotaList) SetCompute(v OpenstackComputeQuotaDto) {
	o.Compute = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *OpenstackQuotaList) GetVolume() OpenstackVolumeQuotaDto {
	if o == nil || IsNil(o.Volume) {
		var ret OpenstackVolumeQuotaDto
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackQuotaList) GetVolumeOk() (*OpenstackVolumeQuotaDto, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *OpenstackQuotaList) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given OpenstackVolumeQuotaDto and assigns it to the Volume field.
func (o *OpenstackQuotaList) SetVolume(v OpenstackVolumeQuotaDto) {
	o.Volume = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OpenstackQuotaList) GetNetwork() OpenstackNetworkDto {
	if o == nil || IsNil(o.Network) {
		var ret OpenstackNetworkDto
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackQuotaList) GetNetworkOk() (*OpenstackNetworkDto, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OpenstackQuotaList) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OpenstackNetworkDto and assigns it to the Network field.
func (o *OpenstackQuotaList) SetNetwork(v OpenstackNetworkDto) {
	o.Network = &v
}

// GetIsInfra returns the IsInfra field value if set, zero value otherwise.
func (o *OpenstackQuotaList) GetIsInfra() bool {
	if o == nil || IsNil(o.IsInfra) {
		var ret bool
		return ret
	}
	return *o.IsInfra
}

// GetIsInfraOk returns a tuple with the IsInfra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackQuotaList) GetIsInfraOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInfra) {
		return nil, false
	}
	return o.IsInfra, true
}

// HasIsInfra returns a boolean if a field has been set.
func (o *OpenstackQuotaList) HasIsInfra() bool {
	if o != nil && !IsNil(o.IsInfra) {
		return true
	}

	return false
}

// SetIsInfra gets a reference to the given bool and assigns it to the IsInfra field.
func (o *OpenstackQuotaList) SetIsInfra(v bool) {
	o.IsInfra = &v
}

func (o OpenstackQuotaList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackQuotaList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.IsInfra) {
		toSerialize["isInfra"] = o.IsInfra
	}
	return toSerialize, nil
}

type NullableOpenstackQuotaList struct {
	value *OpenstackQuotaList
	isSet bool
}

func (v NullableOpenstackQuotaList) Get() *OpenstackQuotaList {
	return v.value
}

func (v *NullableOpenstackQuotaList) Set(val *OpenstackQuotaList) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackQuotaList) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackQuotaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackQuotaList(val *OpenstackQuotaList) *NullableOpenstackQuotaList {
	return &NullableOpenstackQuotaList{value: val, isSet: true}
}

func (v NullableOpenstackQuotaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackQuotaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
