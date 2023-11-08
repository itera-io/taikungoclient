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

// checks if the CloudCredentialsDropdownRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudCredentialsDropdownRecordDto{}

// CloudCredentialsDropdownRecordDto struct for CloudCredentialsDropdownRecordDto
type CloudCredentialsDropdownRecordDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	CloudType NullableString `json:"cloudType,omitempty"`
	Projects []ProjectWithFlavorsAndImagesDto `json:"projects,omitempty"`
}

// NewCloudCredentialsDropdownRecordDto instantiates a new CloudCredentialsDropdownRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCredentialsDropdownRecordDto() *CloudCredentialsDropdownRecordDto {
	this := CloudCredentialsDropdownRecordDto{}
	return &this
}

// NewCloudCredentialsDropdownRecordDtoWithDefaults instantiates a new CloudCredentialsDropdownRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCredentialsDropdownRecordDtoWithDefaults() *CloudCredentialsDropdownRecordDto {
	this := CloudCredentialsDropdownRecordDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudCredentialsDropdownRecordDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCredentialsDropdownRecordDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudCredentialsDropdownRecordDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CloudCredentialsDropdownRecordDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsDropdownRecordDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsDropdownRecordDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CloudCredentialsDropdownRecordDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CloudCredentialsDropdownRecordDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CloudCredentialsDropdownRecordDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CloudCredentialsDropdownRecordDto) UnsetName() {
	o.Name.Unset()
}

// GetCloudType returns the CloudType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsDropdownRecordDto) GetCloudType() string {
	if o == nil || IsNil(o.CloudType.Get()) {
		var ret string
		return ret
	}
	return *o.CloudType.Get()
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsDropdownRecordDto) GetCloudTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudType.Get(), o.CloudType.IsSet()
}

// HasCloudType returns a boolean if a field has been set.
func (o *CloudCredentialsDropdownRecordDto) HasCloudType() bool {
	if o != nil && o.CloudType.IsSet() {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given NullableString and assigns it to the CloudType field.
func (o *CloudCredentialsDropdownRecordDto) SetCloudType(v string) {
	o.CloudType.Set(&v)
}
// SetCloudTypeNil sets the value for CloudType to be an explicit nil
func (o *CloudCredentialsDropdownRecordDto) SetCloudTypeNil() {
	o.CloudType.Set(nil)
}

// UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
func (o *CloudCredentialsDropdownRecordDto) UnsetCloudType() {
	o.CloudType.Unset()
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsDropdownRecordDto) GetProjects() []ProjectWithFlavorsAndImagesDto {
	if o == nil {
		var ret []ProjectWithFlavorsAndImagesDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsDropdownRecordDto) GetProjectsOk() ([]ProjectWithFlavorsAndImagesDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *CloudCredentialsDropdownRecordDto) HasProjects() bool {
	if o != nil && IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectWithFlavorsAndImagesDto and assigns it to the Projects field.
func (o *CloudCredentialsDropdownRecordDto) SetProjects(v []ProjectWithFlavorsAndImagesDto) {
	o.Projects = v
}

func (o CloudCredentialsDropdownRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudCredentialsDropdownRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CloudType.IsSet() {
		toSerialize["cloudType"] = o.CloudType.Get()
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableCloudCredentialsDropdownRecordDto struct {
	value *CloudCredentialsDropdownRecordDto
	isSet bool
}

func (v NullableCloudCredentialsDropdownRecordDto) Get() *CloudCredentialsDropdownRecordDto {
	return v.value
}

func (v *NullableCloudCredentialsDropdownRecordDto) Set(val *CloudCredentialsDropdownRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCredentialsDropdownRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCredentialsDropdownRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCredentialsDropdownRecordDto(val *CloudCredentialsDropdownRecordDto) *NullableCloudCredentialsDropdownRecordDto {
	return &NullableCloudCredentialsDropdownRecordDto{value: val, isSet: true}
}

func (v NullableCloudCredentialsDropdownRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCredentialsDropdownRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


