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

// checks if the StandaloneVmListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandaloneVmListDto{}

// StandaloneVmListDto struct for StandaloneVmListDto
type StandaloneVmListDto struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	FlavorId string `json:"flavorId"`
	VolumeSize int64 `json:"volumeSize"`
	OrganizationName string `json:"organizationName"`
	OrganizationId int32 `json:"organizationId"`
	Ram int64 `json:"ram"`
	Cpu int32 `json:"cpu"`
	VolumeType NullableString `json:"volumeType"`
	PublicIpEnabled bool `json:"publicIpEnabled"`
	PublicIp NullableString `json:"publicIp"`
	IpAddress NullableString `json:"ipAddress"`
	CloudType CloudType `json:"cloudType"`
	ImageName NullableString `json:"imageName"`
	Revision int32 `json:"revision"`
	IsWindows bool `json:"isWindows"`
	Status StandAloneVmStatus `json:"status"`
	ProjectName string `json:"projectName"`
	ProjectId int32 `json:"projectId"`
	StandAloneProfile StandaloneProfileListDto `json:"standAloneProfile"`
	CreatedAt NullableString `json:"createdAt"`
	CreatedBy NullableString `json:"createdBy"`
	LastModified NullableString `json:"lastModified"`
}

type _StandaloneVmListDto StandaloneVmListDto

// NewStandaloneVmListDto instantiates a new StandaloneVmListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandaloneVmListDto(id int32, name string, flavorId string, volumeSize int64, organizationName string, organizationId int32, ram int64, cpu int32, volumeType NullableString, publicIpEnabled bool, publicIp NullableString, ipAddress NullableString, cloudType CloudType, imageName NullableString, revision int32, isWindows bool, status StandAloneVmStatus, projectName string, projectId int32, standAloneProfile StandaloneProfileListDto, createdAt NullableString, createdBy NullableString, lastModified NullableString) *StandaloneVmListDto {
	this := StandaloneVmListDto{}
	this.Id = id
	this.Name = name
	this.FlavorId = flavorId
	this.VolumeSize = volumeSize
	this.OrganizationName = organizationName
	this.OrganizationId = organizationId
	this.Ram = ram
	this.Cpu = cpu
	this.VolumeType = volumeType
	this.PublicIpEnabled = publicIpEnabled
	this.PublicIp = publicIp
	this.IpAddress = ipAddress
	this.CloudType = cloudType
	this.ImageName = imageName
	this.Revision = revision
	this.IsWindows = isWindows
	this.Status = status
	this.ProjectName = projectName
	this.ProjectId = projectId
	this.StandAloneProfile = standAloneProfile
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.LastModified = lastModified
	return &this
}

// NewStandaloneVmListDtoWithDefaults instantiates a new StandaloneVmListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandaloneVmListDtoWithDefaults() *StandaloneVmListDto {
	this := StandaloneVmListDto{}
	return &this
}

// GetId returns the Id field value
func (o *StandaloneVmListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StandaloneVmListDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *StandaloneVmListDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StandaloneVmListDto) SetName(v string) {
	o.Name = v
}

// GetFlavorId returns the FlavorId field value
func (o *StandaloneVmListDto) GetFlavorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetFlavorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlavorId, true
}

// SetFlavorId sets field value
func (o *StandaloneVmListDto) SetFlavorId(v string) {
	o.FlavorId = v
}

// GetVolumeSize returns the VolumeSize field value
func (o *StandaloneVmListDto) GetVolumeSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetVolumeSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeSize, true
}

// SetVolumeSize sets field value
func (o *StandaloneVmListDto) SetVolumeSize(v int64) {
	o.VolumeSize = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *StandaloneVmListDto) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *StandaloneVmListDto) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *StandaloneVmListDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *StandaloneVmListDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetRam returns the Ram field value
func (o *StandaloneVmListDto) GetRam() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetRamOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *StandaloneVmListDto) SetRam(v int64) {
	o.Ram = v
}

// GetCpu returns the Cpu field value
func (o *StandaloneVmListDto) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *StandaloneVmListDto) SetCpu(v int32) {
	o.Cpu = v
}

// GetVolumeType returns the VolumeType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetVolumeType() string {
	if o == nil || o.VolumeType.Get() == nil {
		var ret string
		return ret
	}

	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// SetVolumeType sets field value
func (o *StandaloneVmListDto) SetVolumeType(v string) {
	o.VolumeType.Set(&v)
}

// GetPublicIpEnabled returns the PublicIpEnabled field value
func (o *StandaloneVmListDto) GetPublicIpEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PublicIpEnabled
}

// GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetPublicIpEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicIpEnabled, true
}

// SetPublicIpEnabled sets field value
func (o *StandaloneVmListDto) SetPublicIpEnabled(v bool) {
	o.PublicIpEnabled = v
}

// GetPublicIp returns the PublicIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetPublicIp() string {
	if o == nil || o.PublicIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicIp.Get()
}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetPublicIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicIp.Get(), o.PublicIp.IsSet()
}

// SetPublicIp sets field value
func (o *StandaloneVmListDto) SetPublicIp(v string) {
	o.PublicIp.Set(&v)
}

// GetIpAddress returns the IpAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// SetIpAddress sets field value
func (o *StandaloneVmListDto) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// GetCloudType returns the CloudType field value
func (o *StandaloneVmListDto) GetCloudType() CloudType {
	if o == nil {
		var ret CloudType
		return ret
	}

	return o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetCloudTypeOk() (*CloudType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudType, true
}

// SetCloudType sets field value
func (o *StandaloneVmListDto) SetCloudType(v CloudType) {
	o.CloudType = v
}

// GetImageName returns the ImageName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetImageName() string {
	if o == nil || o.ImageName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ImageName.Get()
}

// GetImageNameOk returns a tuple with the ImageName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageName.Get(), o.ImageName.IsSet()
}

// SetImageName sets field value
func (o *StandaloneVmListDto) SetImageName(v string) {
	o.ImageName.Set(&v)
}

// GetRevision returns the Revision field value
func (o *StandaloneVmListDto) GetRevision() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetRevisionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *StandaloneVmListDto) SetRevision(v int32) {
	o.Revision = v
}

// GetIsWindows returns the IsWindows field value
func (o *StandaloneVmListDto) GetIsWindows() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWindows
}

// GetIsWindowsOk returns a tuple with the IsWindows field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetIsWindowsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWindows, true
}

// SetIsWindows sets field value
func (o *StandaloneVmListDto) SetIsWindows(v bool) {
	o.IsWindows = v
}

// GetStatus returns the Status field value
func (o *StandaloneVmListDto) GetStatus() StandAloneVmStatus {
	if o == nil {
		var ret StandAloneVmStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetStatusOk() (*StandAloneVmStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StandaloneVmListDto) SetStatus(v StandAloneVmStatus) {
	o.Status = v
}

// GetProjectName returns the ProjectName field value
func (o *StandaloneVmListDto) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *StandaloneVmListDto) SetProjectName(v string) {
	o.ProjectName = v
}

// GetProjectId returns the ProjectId field value
func (o *StandaloneVmListDto) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *StandaloneVmListDto) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetStandAloneProfile returns the StandAloneProfile field value
func (o *StandaloneVmListDto) GetStandAloneProfile() StandaloneProfileListDto {
	if o == nil {
		var ret StandaloneProfileListDto
		return ret
	}

	return o.StandAloneProfile
}

// GetStandAloneProfileOk returns a tuple with the StandAloneProfile field value
// and a boolean to check if the value has been set.
func (o *StandaloneVmListDto) GetStandAloneProfileOk() (*StandaloneProfileListDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StandAloneProfile, true
}

// SetStandAloneProfile sets field value
func (o *StandaloneVmListDto) SetStandAloneProfile(v StandaloneProfileListDto) {
	o.StandAloneProfile = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *StandaloneVmListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// SetCreatedBy sets field value
func (o *StandaloneVmListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StandaloneVmListDto) GetLastModified() string {
	if o == nil || o.LastModified.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// SetLastModified sets field value
func (o *StandaloneVmListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}

func (o StandaloneVmListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandaloneVmListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["flavorId"] = o.FlavorId
	toSerialize["volumeSize"] = o.VolumeSize
	toSerialize["organizationName"] = o.OrganizationName
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["ram"] = o.Ram
	toSerialize["cpu"] = o.Cpu
	toSerialize["volumeType"] = o.VolumeType.Get()
	toSerialize["publicIpEnabled"] = o.PublicIpEnabled
	toSerialize["publicIp"] = o.PublicIp.Get()
	toSerialize["ipAddress"] = o.IpAddress.Get()
	toSerialize["cloudType"] = o.CloudType
	toSerialize["imageName"] = o.ImageName.Get()
	toSerialize["revision"] = o.Revision
	toSerialize["isWindows"] = o.IsWindows
	toSerialize["status"] = o.Status
	toSerialize["projectName"] = o.ProjectName
	toSerialize["projectId"] = o.ProjectId
	toSerialize["standAloneProfile"] = o.StandAloneProfile
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["createdBy"] = o.CreatedBy.Get()
	toSerialize["lastModified"] = o.LastModified.Get()
	return toSerialize, nil
}

func (o *StandaloneVmListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"flavorId",
		"volumeSize",
		"organizationName",
		"organizationId",
		"ram",
		"cpu",
		"volumeType",
		"publicIpEnabled",
		"publicIp",
		"ipAddress",
		"cloudType",
		"imageName",
		"revision",
		"isWindows",
		"status",
		"projectName",
		"projectId",
		"standAloneProfile",
		"createdAt",
		"createdBy",
		"lastModified",
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

	varStandaloneVmListDto := _StandaloneVmListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStandaloneVmListDto)

	if err != nil {
		return err
	}

	*o = StandaloneVmListDto(varStandaloneVmListDto)

	return err
}

type NullableStandaloneVmListDto struct {
	value *StandaloneVmListDto
	isSet bool
}

func (v NullableStandaloneVmListDto) Get() *StandaloneVmListDto {
	return v.value
}

func (v *NullableStandaloneVmListDto) Set(val *StandaloneVmListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandaloneVmListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandaloneVmListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandaloneVmListDto(val *StandaloneVmListDto) *NullableStandaloneVmListDto {
	return &NullableStandaloneVmListDto{value: val, isSet: true}
}

func (v NullableStandaloneVmListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandaloneVmListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


