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

// checks if the VClusterList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VClusterList{}

// VClusterList struct for VClusterList
type VClusterList struct {
	Data []VClusterListDto `json:"data,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Project *ProjectDetailsForVmsDto `json:"project,omitempty"`
}

// NewVClusterList instantiates a new VClusterList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVClusterList() *VClusterList {
	this := VClusterList{}
	return &this
}

// NewVClusterListWithDefaults instantiates a new VClusterList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVClusterListWithDefaults() *VClusterList {
	this := VClusterList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VClusterList) GetData() []VClusterListDto {
	if o == nil {
		var ret []VClusterListDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VClusterList) GetDataOk() ([]VClusterListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VClusterList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []VClusterListDto and assigns it to the Data field.
func (o *VClusterList) SetData(v []VClusterListDto) {
	o.Data = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *VClusterList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VClusterList) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *VClusterList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *VClusterList) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *VClusterList) GetProject() ProjectDetailsForVmsDto {
	if o == nil || IsNil(o.Project) {
		var ret ProjectDetailsForVmsDto
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VClusterList) GetProjectOk() (*ProjectDetailsForVmsDto, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *VClusterList) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectDetailsForVmsDto and assigns it to the Project field.
func (o *VClusterList) SetProject(v ProjectDetailsForVmsDto) {
	o.Project = &v
}

func (o VClusterList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VClusterList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	return toSerialize, nil
}

type NullableVClusterList struct {
	value *VClusterList
	isSet bool
}

func (v NullableVClusterList) Get() *VClusterList {
	return v.value
}

func (v *NullableVClusterList) Set(val *VClusterList) {
	v.value = val
	v.isSet = true
}

func (v NullableVClusterList) IsSet() bool {
	return v.isSet
}

func (v *NullableVClusterList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVClusterList(val *VClusterList) *NullableVClusterList {
	return &NullableVClusterList{value: val, isSet: true}
}

func (v NullableVClusterList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVClusterList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


