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

// checks if the ProjectQuotaListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectQuotaListDto{}

// ProjectQuotaListDto struct for ProjectQuotaListDto
type ProjectQuotaListDto struct {
	ServerCpu int64 `json:"serverCpu"`
	ServerRam float64 `json:"serverRam"`
	ServerDiskSize float64 `json:"serverDiskSize"`
	VmCpu int64 `json:"vmCpu"`
	VmRam float64 `json:"vmRam"`
	VmVolumeSize float64 `json:"vmVolumeSize"`
	ProjectId int32 `json:"projectId"`
	ProjectName NullableString `json:"projectName"`
}

type _ProjectQuotaListDto ProjectQuotaListDto

// NewProjectQuotaListDto instantiates a new ProjectQuotaListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectQuotaListDto(serverCpu int64, serverRam float64, serverDiskSize float64, vmCpu int64, vmRam float64, vmVolumeSize float64, projectId int32, projectName NullableString) *ProjectQuotaListDto {
	this := ProjectQuotaListDto{}
	this.ServerCpu = serverCpu
	this.ServerRam = serverRam
	this.ServerDiskSize = serverDiskSize
	this.VmCpu = vmCpu
	this.VmRam = vmRam
	this.VmVolumeSize = vmVolumeSize
	this.ProjectId = projectId
	this.ProjectName = projectName
	return &this
}

// NewProjectQuotaListDtoWithDefaults instantiates a new ProjectQuotaListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectQuotaListDtoWithDefaults() *ProjectQuotaListDto {
	this := ProjectQuotaListDto{}
	return &this
}

// GetServerCpu returns the ServerCpu field value
func (o *ProjectQuotaListDto) GetServerCpu() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ServerCpu
}

// GetServerCpuOk returns a tuple with the ServerCpu field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetServerCpuOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerCpu, true
}

// SetServerCpu sets field value
func (o *ProjectQuotaListDto) SetServerCpu(v int64) {
	o.ServerCpu = v
}

// GetServerRam returns the ServerRam field value
func (o *ProjectQuotaListDto) GetServerRam() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ServerRam
}

// GetServerRamOk returns a tuple with the ServerRam field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetServerRamOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerRam, true
}

// SetServerRam sets field value
func (o *ProjectQuotaListDto) SetServerRam(v float64) {
	o.ServerRam = v
}

// GetServerDiskSize returns the ServerDiskSize field value
func (o *ProjectQuotaListDto) GetServerDiskSize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ServerDiskSize
}

// GetServerDiskSizeOk returns a tuple with the ServerDiskSize field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetServerDiskSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerDiskSize, true
}

// SetServerDiskSize sets field value
func (o *ProjectQuotaListDto) SetServerDiskSize(v float64) {
	o.ServerDiskSize = v
}

// GetVmCpu returns the VmCpu field value
func (o *ProjectQuotaListDto) GetVmCpu() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VmCpu
}

// GetVmCpuOk returns a tuple with the VmCpu field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetVmCpuOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmCpu, true
}

// SetVmCpu sets field value
func (o *ProjectQuotaListDto) SetVmCpu(v int64) {
	o.VmCpu = v
}

// GetVmRam returns the VmRam field value
func (o *ProjectQuotaListDto) GetVmRam() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.VmRam
}

// GetVmRamOk returns a tuple with the VmRam field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetVmRamOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmRam, true
}

// SetVmRam sets field value
func (o *ProjectQuotaListDto) SetVmRam(v float64) {
	o.VmRam = v
}

// GetVmVolumeSize returns the VmVolumeSize field value
func (o *ProjectQuotaListDto) GetVmVolumeSize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.VmVolumeSize
}

// GetVmVolumeSizeOk returns a tuple with the VmVolumeSize field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetVmVolumeSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmVolumeSize, true
}

// SetVmVolumeSize sets field value
func (o *ProjectQuotaListDto) SetVmVolumeSize(v float64) {
	o.VmVolumeSize = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectQuotaListDto) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectQuotaListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectQuotaListDto) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectQuotaListDto) GetProjectName() string {
	if o == nil || o.ProjectName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectQuotaListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// SetProjectName sets field value
func (o *ProjectQuotaListDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
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
	toSerialize["serverCpu"] = o.ServerCpu
	toSerialize["serverRam"] = o.ServerRam
	toSerialize["serverDiskSize"] = o.ServerDiskSize
	toSerialize["vmCpu"] = o.VmCpu
	toSerialize["vmRam"] = o.VmRam
	toSerialize["vmVolumeSize"] = o.VmVolumeSize
	toSerialize["projectId"] = o.ProjectId
	toSerialize["projectName"] = o.ProjectName.Get()
	return toSerialize, nil
}

func (o *ProjectQuotaListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serverCpu",
		"serverRam",
		"serverDiskSize",
		"vmCpu",
		"vmRam",
		"vmVolumeSize",
		"projectId",
		"projectName",
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

	varProjectQuotaListDto := _ProjectQuotaListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectQuotaListDto)

	if err != nil {
		return err
	}

	*o = ProjectQuotaListDto(varProjectQuotaListDto)

	return err
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


