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

// checks if the ServerCommonRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerCommonRecordDto{}

// ServerCommonRecordDto struct for ServerCommonRecordDto
type ServerCommonRecordDto struct {
	ProjectId *int32 `json:"projectId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	Names []string `json:"names,omitempty"`
}

// NewServerCommonRecordDto instantiates a new ServerCommonRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerCommonRecordDto() *ServerCommonRecordDto {
	this := ServerCommonRecordDto{}
	return &this
}

// NewServerCommonRecordDtoWithDefaults instantiates a new ServerCommonRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerCommonRecordDtoWithDefaults() *ServerCommonRecordDto {
	this := ServerCommonRecordDto{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ServerCommonRecordDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerCommonRecordDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ServerCommonRecordDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ServerCommonRecordDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerCommonRecordDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCommonRecordDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *ServerCommonRecordDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *ServerCommonRecordDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *ServerCommonRecordDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *ServerCommonRecordDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetNames returns the Names field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerCommonRecordDto) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerCommonRecordDto) GetNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ServerCommonRecordDto) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *ServerCommonRecordDto) SetNames(v []string) {
	o.Names = v
}

func (o ServerCommonRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerCommonRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	return toSerialize, nil
}

type NullableServerCommonRecordDto struct {
	value *ServerCommonRecordDto
	isSet bool
}

func (v NullableServerCommonRecordDto) Get() *ServerCommonRecordDto {
	return v.value
}

func (v *NullableServerCommonRecordDto) Set(val *ServerCommonRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServerCommonRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServerCommonRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerCommonRecordDto(val *ServerCommonRecordDto) *NullableServerCommonRecordDto {
	return &NullableServerCommonRecordDto{value: val, isSet: true}
}

func (v NullableServerCommonRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerCommonRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


