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
	"bytes"
	"fmt"
)

// checks if the VClusterResourceLimitsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VClusterResourceLimitsListDto{}

// VClusterResourceLimitsListDto struct for VClusterResourceLimitsListDto
type VClusterResourceLimitsListDto struct {
	ResourceUnit NullableString `json:"resourceUnit"`
	MaxClusterRequests float64 `json:"maxClusterRequests"`
	MaxClusterLimits float64 `json:"maxClusterLimits"`
	DefaultContainerLimit float64 `json:"defaultContainerLimit"`
	DefaultContainerRequest float64 `json:"defaultContainerRequest"`
}

type _VClusterResourceLimitsListDto VClusterResourceLimitsListDto

// NewVClusterResourceLimitsListDto instantiates a new VClusterResourceLimitsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVClusterResourceLimitsListDto(resourceUnit NullableString, maxClusterRequests float64, maxClusterLimits float64, defaultContainerLimit float64, defaultContainerRequest float64) *VClusterResourceLimitsListDto {
	this := VClusterResourceLimitsListDto{}
	this.ResourceUnit = resourceUnit
	this.MaxClusterRequests = maxClusterRequests
	this.MaxClusterLimits = maxClusterLimits
	this.DefaultContainerLimit = defaultContainerLimit
	this.DefaultContainerRequest = defaultContainerRequest
	return &this
}

// NewVClusterResourceLimitsListDtoWithDefaults instantiates a new VClusterResourceLimitsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVClusterResourceLimitsListDtoWithDefaults() *VClusterResourceLimitsListDto {
	this := VClusterResourceLimitsListDto{}
	return &this
}

// GetResourceUnit returns the ResourceUnit field value
// If the value is explicit nil, the zero value for string will be returned
func (o *VClusterResourceLimitsListDto) GetResourceUnit() string {
	if o == nil || o.ResourceUnit.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResourceUnit.Get()
}

// GetResourceUnitOk returns a tuple with the ResourceUnit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VClusterResourceLimitsListDto) GetResourceUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceUnit.Get(), o.ResourceUnit.IsSet()
}

// SetResourceUnit sets field value
func (o *VClusterResourceLimitsListDto) SetResourceUnit(v string) {
	o.ResourceUnit.Set(&v)
}

// GetMaxClusterRequests returns the MaxClusterRequests field value
func (o *VClusterResourceLimitsListDto) GetMaxClusterRequests() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MaxClusterRequests
}

// GetMaxClusterRequestsOk returns a tuple with the MaxClusterRequests field value
// and a boolean to check if the value has been set.
func (o *VClusterResourceLimitsListDto) GetMaxClusterRequestsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxClusterRequests, true
}

// SetMaxClusterRequests sets field value
func (o *VClusterResourceLimitsListDto) SetMaxClusterRequests(v float64) {
	o.MaxClusterRequests = v
}

// GetMaxClusterLimits returns the MaxClusterLimits field value
func (o *VClusterResourceLimitsListDto) GetMaxClusterLimits() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MaxClusterLimits
}

// GetMaxClusterLimitsOk returns a tuple with the MaxClusterLimits field value
// and a boolean to check if the value has been set.
func (o *VClusterResourceLimitsListDto) GetMaxClusterLimitsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxClusterLimits, true
}

// SetMaxClusterLimits sets field value
func (o *VClusterResourceLimitsListDto) SetMaxClusterLimits(v float64) {
	o.MaxClusterLimits = v
}

// GetDefaultContainerLimit returns the DefaultContainerLimit field value
func (o *VClusterResourceLimitsListDto) GetDefaultContainerLimit() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DefaultContainerLimit
}

// GetDefaultContainerLimitOk returns a tuple with the DefaultContainerLimit field value
// and a boolean to check if the value has been set.
func (o *VClusterResourceLimitsListDto) GetDefaultContainerLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultContainerLimit, true
}

// SetDefaultContainerLimit sets field value
func (o *VClusterResourceLimitsListDto) SetDefaultContainerLimit(v float64) {
	o.DefaultContainerLimit = v
}

// GetDefaultContainerRequest returns the DefaultContainerRequest field value
func (o *VClusterResourceLimitsListDto) GetDefaultContainerRequest() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DefaultContainerRequest
}

// GetDefaultContainerRequestOk returns a tuple with the DefaultContainerRequest field value
// and a boolean to check if the value has been set.
func (o *VClusterResourceLimitsListDto) GetDefaultContainerRequestOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultContainerRequest, true
}

// SetDefaultContainerRequest sets field value
func (o *VClusterResourceLimitsListDto) SetDefaultContainerRequest(v float64) {
	o.DefaultContainerRequest = v
}

func (o VClusterResourceLimitsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VClusterResourceLimitsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceUnit"] = o.ResourceUnit.Get()
	toSerialize["maxClusterRequests"] = o.MaxClusterRequests
	toSerialize["maxClusterLimits"] = o.MaxClusterLimits
	toSerialize["defaultContainerLimit"] = o.DefaultContainerLimit
	toSerialize["defaultContainerRequest"] = o.DefaultContainerRequest
	return toSerialize, nil
}

func (o *VClusterResourceLimitsListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceUnit",
		"maxClusterRequests",
		"maxClusterLimits",
		"defaultContainerLimit",
		"defaultContainerRequest",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVClusterResourceLimitsListDto := _VClusterResourceLimitsListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVClusterResourceLimitsListDto)

	if err != nil {
		return err
	}

	*o = VClusterResourceLimitsListDto(varVClusterResourceLimitsListDto)

	return err
}

type NullableVClusterResourceLimitsListDto struct {
	value *VClusterResourceLimitsListDto
	isSet bool
}

func (v NullableVClusterResourceLimitsListDto) Get() *VClusterResourceLimitsListDto {
	return v.value
}

func (v *NullableVClusterResourceLimitsListDto) Set(val *VClusterResourceLimitsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableVClusterResourceLimitsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableVClusterResourceLimitsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVClusterResourceLimitsListDto(val *VClusterResourceLimitsListDto) *NullableVClusterResourceLimitsListDto {
	return &NullableVClusterResourceLimitsListDto{value: val, isSet: true}
}

func (v NullableVClusterResourceLimitsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVClusterResourceLimitsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


