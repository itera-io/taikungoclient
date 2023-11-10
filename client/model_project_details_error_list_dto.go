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

// checks if the ProjectDetailsErrorListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectDetailsErrorListDto{}

// ProjectDetailsErrorListDto struct for ProjectDetailsErrorListDto
type ProjectDetailsErrorListDto struct {
	Type *ProjectDetailsErrorType `json:"type,omitempty"`
	Message []string `json:"message,omitempty"`
}

// NewProjectDetailsErrorListDto instantiates a new ProjectDetailsErrorListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDetailsErrorListDto() *ProjectDetailsErrorListDto {
	this := ProjectDetailsErrorListDto{}
	return &this
}

// NewProjectDetailsErrorListDtoWithDefaults instantiates a new ProjectDetailsErrorListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDetailsErrorListDtoWithDefaults() *ProjectDetailsErrorListDto {
	this := ProjectDetailsErrorListDto{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProjectDetailsErrorListDto) GetType() ProjectDetailsErrorType {
	if o == nil || IsNil(o.Type) {
		var ret ProjectDetailsErrorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetailsErrorListDto) GetTypeOk() (*ProjectDetailsErrorType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProjectDetailsErrorListDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProjectDetailsErrorType and assigns it to the Type field.
func (o *ProjectDetailsErrorListDto) SetType(v ProjectDetailsErrorType) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectDetailsErrorListDto) GetMessage() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectDetailsErrorListDto) GetMessageOk() ([]string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProjectDetailsErrorListDto) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given []string and assigns it to the Message field.
func (o *ProjectDetailsErrorListDto) SetMessage(v []string) {
	o.Message = v
}

func (o ProjectDetailsErrorListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectDetailsErrorListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableProjectDetailsErrorListDto struct {
	value *ProjectDetailsErrorListDto
	isSet bool
}

func (v NullableProjectDetailsErrorListDto) Get() *ProjectDetailsErrorListDto {
	return v.value
}

func (v *NullableProjectDetailsErrorListDto) Set(val *ProjectDetailsErrorListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetailsErrorListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetailsErrorListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetailsErrorListDto(val *ProjectDetailsErrorListDto) *NullableProjectDetailsErrorListDto {
	return &NullableProjectDetailsErrorListDto{value: val, isSet: true}
}

func (v NullableProjectDetailsErrorListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetailsErrorListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


