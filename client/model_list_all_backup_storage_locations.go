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

// checks if the ListAllBackupStorageLocations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAllBackupStorageLocations{}

// ListAllBackupStorageLocations struct for ListAllBackupStorageLocations
type ListAllBackupStorageLocations struct {
	Data []BackupStorageLocationDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Projects []int32 `json:"projects,omitempty"`
}

// NewListAllBackupStorageLocations instantiates a new ListAllBackupStorageLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllBackupStorageLocations() *ListAllBackupStorageLocations {
	this := ListAllBackupStorageLocations{}
	return &this
}

// NewListAllBackupStorageLocationsWithDefaults instantiates a new ListAllBackupStorageLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllBackupStorageLocationsWithDefaults() *ListAllBackupStorageLocations {
	this := ListAllBackupStorageLocations{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAllBackupStorageLocations) GetData() []BackupStorageLocationDto {
	if o == nil {
		var ret []BackupStorageLocationDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAllBackupStorageLocations) GetDataOk() ([]BackupStorageLocationDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListAllBackupStorageLocations) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BackupStorageLocationDto and assigns it to the Data field.
func (o *ListAllBackupStorageLocations) SetData(v []BackupStorageLocationDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ListAllBackupStorageLocations) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAllBackupStorageLocations) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ListAllBackupStorageLocations) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ListAllBackupStorageLocations) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListAllBackupStorageLocations) GetProjects() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListAllBackupStorageLocations) GetProjectsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *ListAllBackupStorageLocations) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []int32 and assigns it to the Projects field.
func (o *ListAllBackupStorageLocations) SetProjects(v []int32) {
	o.Projects = v
}

func (o ListAllBackupStorageLocations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllBackupStorageLocations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableListAllBackupStorageLocations struct {
	value *ListAllBackupStorageLocations
	isSet bool
}

func (v NullableListAllBackupStorageLocations) Get() *ListAllBackupStorageLocations {
	return v.value
}

func (v *NullableListAllBackupStorageLocations) Set(val *ListAllBackupStorageLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllBackupStorageLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllBackupStorageLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllBackupStorageLocations(val *ListAllBackupStorageLocations) *NullableListAllBackupStorageLocations {
	return &NullableListAllBackupStorageLocations{value: val, isSet: true}
}

func (v NullableListAllBackupStorageLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllBackupStorageLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


