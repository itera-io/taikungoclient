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

// checks if the OpenstackCredentialList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenstackCredentialList{}

// OpenstackCredentialList struct for OpenstackCredentialList
type OpenstackCredentialList struct {
	Data []OpenstackCredentialsListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewOpenstackCredentialList instantiates a new OpenstackCredentialList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenstackCredentialList() *OpenstackCredentialList {
	this := OpenstackCredentialList{}
	return &this
}

// NewOpenstackCredentialListWithDefaults instantiates a new OpenstackCredentialList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenstackCredentialListWithDefaults() *OpenstackCredentialList {
	this := OpenstackCredentialList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OpenstackCredentialList) GetData() []OpenstackCredentialsListDto {
	if o == nil || IsNil(o.Data) {
		var ret []OpenstackCredentialsListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialList) GetDataOk() ([]OpenstackCredentialsListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OpenstackCredentialList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []OpenstackCredentialsListDto and assigns it to the Data field.
func (o *OpenstackCredentialList) SetData(v []OpenstackCredentialsListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OpenstackCredentialList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenstackCredentialList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OpenstackCredentialList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *OpenstackCredentialList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o OpenstackCredentialList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenstackCredentialList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableOpenstackCredentialList struct {
	value *OpenstackCredentialList
	isSet bool
}

func (v NullableOpenstackCredentialList) Get() *OpenstackCredentialList {
	return v.value
}

func (v *NullableOpenstackCredentialList) Set(val *OpenstackCredentialList) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenstackCredentialList) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenstackCredentialList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenstackCredentialList(val *OpenstackCredentialList) *NullableOpenstackCredentialList {
	return &NullableOpenstackCredentialList{value: val, isSet: true}
}

func (v NullableOpenstackCredentialList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenstackCredentialList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


