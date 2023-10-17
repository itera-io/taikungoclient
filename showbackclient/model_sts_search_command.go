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

// checks if the StsSearchCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StsSearchCommand{}

// StsSearchCommand struct for StsSearchCommand
type StsSearchCommand struct {
	Limit      NullableInt32  `json:"limit,omitempty"`
	Offset     NullableInt32  `json:"offset,omitempty"`
	SearchTerm NullableString `json:"searchTerm,omitempty"`
}

// NewStsSearchCommand instantiates a new StsSearchCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStsSearchCommand() *StsSearchCommand {
	this := StsSearchCommand{}
	return &this
}

// NewStsSearchCommandWithDefaults instantiates a new StsSearchCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStsSearchCommandWithDefaults() *StsSearchCommand {
	this := StsSearchCommand{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StsSearchCommand) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StsSearchCommand) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *StsSearchCommand) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *StsSearchCommand) SetLimit(v int32) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil
func (o *StsSearchCommand) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *StsSearchCommand) UnsetLimit() {
	o.Limit.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StsSearchCommand) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StsSearchCommand) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *StsSearchCommand) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *StsSearchCommand) SetOffset(v int32) {
	o.Offset.Set(&v)
}

// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *StsSearchCommand) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *StsSearchCommand) UnsetOffset() {
	o.Offset.Unset()
}

// GetSearchTerm returns the SearchTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StsSearchCommand) GetSearchTerm() string {
	if o == nil || IsNil(o.SearchTerm.Get()) {
		var ret string
		return ret
	}
	return *o.SearchTerm.Get()
}

// GetSearchTermOk returns a tuple with the SearchTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StsSearchCommand) GetSearchTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchTerm.Get(), o.SearchTerm.IsSet()
}

// HasSearchTerm returns a boolean if a field has been set.
func (o *StsSearchCommand) HasSearchTerm() bool {
	if o != nil && o.SearchTerm.IsSet() {
		return true
	}

	return false
}

// SetSearchTerm gets a reference to the given NullableString and assigns it to the SearchTerm field.
func (o *StsSearchCommand) SetSearchTerm(v string) {
	o.SearchTerm.Set(&v)
}

// SetSearchTermNil sets the value for SearchTerm to be an explicit nil
func (o *StsSearchCommand) SetSearchTermNil() {
	o.SearchTerm.Set(nil)
}

// UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
func (o *StsSearchCommand) UnsetSearchTerm() {
	o.SearchTerm.Unset()
}

func (o StsSearchCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StsSearchCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.SearchTerm.IsSet() {
		toSerialize["searchTerm"] = o.SearchTerm.Get()
	}
	return toSerialize, nil
}

type NullableStsSearchCommand struct {
	value *StsSearchCommand
	isSet bool
}

func (v NullableStsSearchCommand) Get() *StsSearchCommand {
	return v.value
}

func (v *NullableStsSearchCommand) Set(val *StsSearchCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableStsSearchCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableStsSearchCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStsSearchCommand(val *StsSearchCommand) *NullableStsSearchCommand {
	return &NullableStsSearchCommand{value: val, isSet: true}
}

func (v NullableStsSearchCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStsSearchCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
