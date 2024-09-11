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

// checks if the ProjectQuotaListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectQuotaListDto{}

// ProjectQuotaListDto struct for ProjectQuotaListDto
type ProjectQuotaListDto struct {
	ServerCpu *int64 `json:"serverCpu,omitempty"`
	ServerRam *float64 `json:"serverRam,omitempty"`
	ServerDiskSize *float64 `json:"serverDiskSize,omitempty"`
	VmCpu *int64 `json:"vmCpu,omitempty"`
	VmRam *float64 `json:"vmRam,omitempty"`
	VmVolumeSize *float64 `json:"vmVolumeSize,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
}

// NewProjectQuotaListDto instantiates a new ProjectQuotaListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectQuotaListDto() *ProjectQuotaListDto {
	this := ProjectQuotaListDto{}
	return &this
}

// NewProjectQuotaListDtoWithDefaults instantiates a new ProjectQuotaListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectQuotaListDtoWithDefaults() *ProjectQuotaListDto {
	this := ProjectQuotaListDto{}
	return &this
}

// GetServerCpu returns the ServerCpu field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetServerCpu() int64 {
	if o == nil || IsNil(o.ServerCpu) {
		var ret int64
		return ret
	}
	return *o.ServerCpu
}

// GetServerCpuOk returns a tuple with the ServerCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetServerCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerCpu) {
		return nil, false
	}
	return o.ServerCpu, true
}

// HasServerCpu returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasServerCpu() bool {
	if o != nil && !IsNil(o.ServerCpu) {
		return true
	}

	return false
}

// SetServerCpu gets a reference to the given int64 and assigns it to the ServerCpu field.
func (o *ProjectQuotaListDto) SetServerCpu(v int64) {
	o.ServerCpu = &v
}

// GetServerRam returns the ServerRam field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetServerRam() float64 {
	if o == nil || IsNil(o.ServerRam) {
		var ret float64
		return ret
	}
	return *o.ServerRam
}

// GetServerRamOk returns a tuple with the ServerRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetServerRamOk() (*float64, bool) {
	if o == nil || IsNil(o.ServerRam) {
		return nil, false
	}
	return o.ServerRam, true
}

// HasServerRam returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasServerRam() bool {
	if o != nil && !IsNil(o.ServerRam) {
		return true
	}

	return false
}

// SetServerRam gets a reference to the given float64 and assigns it to the ServerRam field.
func (o *ProjectQuotaListDto) SetServerRam(v float64) {
	o.ServerRam = &v
}

// GetServerDiskSize returns the ServerDiskSize field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetServerDiskSize() float64 {
	if o == nil || IsNil(o.ServerDiskSize) {
		var ret float64
		return ret
	}
	return *o.ServerDiskSize
}

// GetServerDiskSizeOk returns a tuple with the ServerDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetServerDiskSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.ServerDiskSize) {
		return nil, false
	}
	return o.ServerDiskSize, true
}

// HasServerDiskSize returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasServerDiskSize() bool {
	if o != nil && !IsNil(o.ServerDiskSize) {
		return true
	}

	return false
}

// SetServerDiskSize gets a reference to the given float64 and assigns it to the ServerDiskSize field.
func (o *ProjectQuotaListDto) SetServerDiskSize(v float64) {
	o.ServerDiskSize = &v
}

// GetVmCpu returns the VmCpu field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetVmCpu() int64 {
	if o == nil || IsNil(o.VmCpu) {
		var ret int64
		return ret
	}
	return *o.VmCpu
}

// GetVmCpuOk returns a tuple with the VmCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetVmCpuOk() (*int64, bool) {
	if o == nil || IsNil(o.VmCpu) {
		return nil, false
	}
	return o.VmCpu, true
}

// HasVmCpu returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasVmCpu() bool {
	if o != nil && !IsNil(o.VmCpu) {
		return true
	}

	return false
}

// SetVmCpu gets a reference to the given int64 and assigns it to the VmCpu field.
func (o *ProjectQuotaListDto) SetVmCpu(v int64) {
	o.VmCpu = &v
}

// GetVmRam returns the VmRam field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetVmRam() float64 {
	if o == nil || IsNil(o.VmRam) {
		var ret float64
		return ret
	}
	return *o.VmRam
}

// GetVmRamOk returns a tuple with the VmRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetVmRamOk() (*float64, bool) {
	if o == nil || IsNil(o.VmRam) {
		return nil, false
	}
	return o.VmRam, true
}

// HasVmRam returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasVmRam() bool {
	if o != nil && !IsNil(o.VmRam) {
		return true
	}

	return false
}

// SetVmRam gets a reference to the given float64 and assigns it to the VmRam field.
func (o *ProjectQuotaListDto) SetVmRam(v float64) {
	o.VmRam = &v
}

// GetVmVolumeSize returns the VmVolumeSize field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetVmVolumeSize() float64 {
	if o == nil || IsNil(o.VmVolumeSize) {
		var ret float64
		return ret
	}
	return *o.VmVolumeSize
}

// GetVmVolumeSizeOk returns a tuple with the VmVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetVmVolumeSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.VmVolumeSize) {
		return nil, false
	}
	return o.VmVolumeSize, true
}

// HasVmVolumeSize returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasVmVolumeSize() bool {
	if o != nil && !IsNil(o.VmVolumeSize) {
		return true
	}

	return false
}

// SetVmVolumeSize gets a reference to the given float64 and assigns it to the VmVolumeSize field.
func (o *ProjectQuotaListDto) SetVmVolumeSize(v float64) {
	o.VmVolumeSize = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ProjectQuotaListDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *ProjectQuotaListDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *ProjectQuotaListDto) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *ProjectQuotaListDto) SetProjectName(v string) {
	o.ProjectName = &v
}

func (o ProjectQuotaListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectQuotaListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerCpu) {
		toSerialize["serverCpu"] = o.ServerCpu
	}
	if !IsNil(o.ServerRam) {
		toSerialize["serverRam"] = o.ServerRam
	}
	if !IsNil(o.ServerDiskSize) {
		toSerialize["serverDiskSize"] = o.ServerDiskSize
	}
	if !IsNil(o.VmCpu) {
		toSerialize["vmCpu"] = o.VmCpu
	}
	if !IsNil(o.VmRam) {
		toSerialize["vmRam"] = o.VmRam
	}
	if !IsNil(o.VmVolumeSize) {
		toSerialize["vmVolumeSize"] = o.VmVolumeSize
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ProjectName) {
		toSerialize["projectName"] = o.ProjectName
	}
	return toSerialize, nil
}

type NullableProjectQuotaListDto struct {
	value *ProjectQuotaListDto
	isSet bool
}

func (v NullableProjectQuotaListDto) Get() *ProjectQuotaListDto {
	return v.value
}

func (v *NullableProjectQuotaListDto) Set(val *ProjectQuotaListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectQuotaListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectQuotaListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectQuotaListDto(val *ProjectQuotaListDto) *NullableProjectQuotaListDto {
	return &NullableProjectQuotaListDto{value: val, isSet: true}
}

func (v NullableProjectQuotaListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectQuotaListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


