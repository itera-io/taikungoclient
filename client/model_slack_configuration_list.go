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

// checks if the SlackConfigurationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackConfigurationList{}

// SlackConfigurationList struct for SlackConfigurationList
type SlackConfigurationList struct {
	Data []SlackConfigurationDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewSlackConfigurationList instantiates a new SlackConfigurationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackConfigurationList() *SlackConfigurationList {
	this := SlackConfigurationList{}
	return &this
}

// NewSlackConfigurationListWithDefaults instantiates a new SlackConfigurationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackConfigurationListWithDefaults() *SlackConfigurationList {
	this := SlackConfigurationList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SlackConfigurationList) GetData() []SlackConfigurationDto {
	if o == nil {
		var ret []SlackConfigurationDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SlackConfigurationList) GetDataOk() ([]SlackConfigurationDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SlackConfigurationList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SlackConfigurationDto and assigns it to the Data field.
func (o *SlackConfigurationList) SetData(v []SlackConfigurationDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SlackConfigurationList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackConfigurationList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SlackConfigurationList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *SlackConfigurationList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o SlackConfigurationList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackConfigurationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableSlackConfigurationList struct {
	value *SlackConfigurationList
	isSet bool
}

func (v NullableSlackConfigurationList) Get() *SlackConfigurationList {
	return v.value
}

func (v *NullableSlackConfigurationList) Set(val *SlackConfigurationList) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackConfigurationList) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackConfigurationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackConfigurationList(val *SlackConfigurationList) *NullableSlackConfigurationList {
	return &NullableSlackConfigurationList{value: val, isSet: true}
}

func (v NullableSlackConfigurationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackConfigurationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


