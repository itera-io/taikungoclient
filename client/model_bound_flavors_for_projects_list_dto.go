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

// checks if the BoundFlavorsForProjectsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoundFlavorsForProjectsListDto{}

// BoundFlavorsForProjectsListDto struct for BoundFlavorsForProjectsListDto
type BoundFlavorsForProjectsListDto struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Cpu int32 `json:"cpu"`
	Ram float64 `json:"ram"`
	ProjectId NullableInt32 `json:"projectId"`
	ProjectName string `json:"projectName"`
	MaxDataDiskCount NullableInt32 `json:"maxDataDiskCount"`
	ExistsLinuxSpotPrice bool `json:"existsLinuxSpotPrice"`
	ExistsWindowsSpotPrice bool `json:"existsWindowsSpotPrice"`
	LinuxSpotPrice NullableString `json:"linuxSpotPrice"`
	LinuxPrice NullableString `json:"linuxPrice"`
	WindowsSpotPrice NullableString `json:"windowsSpotPrice"`
	WindowsPrice NullableString `json:"windowsPrice"`
	CloudType CloudType `json:"cloudType"`
	LocalDiskSize *int32 `json:"localDiskSize,omitempty"`
}

type _BoundFlavorsForProjectsListDto BoundFlavorsForProjectsListDto

// NewBoundFlavorsForProjectsListDto instantiates a new BoundFlavorsForProjectsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoundFlavorsForProjectsListDto(id int32, name string, cpu int32, ram float64, projectId NullableInt32, projectName string, maxDataDiskCount NullableInt32, existsLinuxSpotPrice bool, existsWindowsSpotPrice bool, linuxSpotPrice NullableString, linuxPrice NullableString, windowsSpotPrice NullableString, windowsPrice NullableString, cloudType CloudType) *BoundFlavorsForProjectsListDto {
	this := BoundFlavorsForProjectsListDto{}
	this.Id = id
	this.Name = name
	this.Cpu = cpu
	this.Ram = ram
	this.ProjectId = projectId
	this.ProjectName = projectName
	this.MaxDataDiskCount = maxDataDiskCount
	this.ExistsLinuxSpotPrice = existsLinuxSpotPrice
	this.ExistsWindowsSpotPrice = existsWindowsSpotPrice
	this.LinuxSpotPrice = linuxSpotPrice
	this.LinuxPrice = linuxPrice
	this.WindowsSpotPrice = windowsSpotPrice
	this.WindowsPrice = windowsPrice
	this.CloudType = cloudType
	return &this
}

// NewBoundFlavorsForProjectsListDtoWithDefaults instantiates a new BoundFlavorsForProjectsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoundFlavorsForProjectsListDtoWithDefaults() *BoundFlavorsForProjectsListDto {
	this := BoundFlavorsForProjectsListDto{}
	return &this
}

// GetId returns the Id field value
func (o *BoundFlavorsForProjectsListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BoundFlavorsForProjectsListDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BoundFlavorsForProjectsListDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BoundFlavorsForProjectsListDto) SetName(v string) {
	o.Name = v
}

// GetCpu returns the Cpu field value
func (o *BoundFlavorsForProjectsListDto) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *BoundFlavorsForProjectsListDto) SetCpu(v int32) {
	o.Cpu = v
}

// GetRam returns the Ram field value
func (o *BoundFlavorsForProjectsListDto) GetRam() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Ram
}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetRamOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ram, true
}

// SetRam sets field value
func (o *BoundFlavorsForProjectsListDto) SetRam(v float64) {
	o.Ram = v
}

// GetProjectId returns the ProjectId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BoundFlavorsForProjectsListDto) GetProjectId() int32 {
	if o == nil || o.ProjectId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// SetProjectId sets field value
func (o *BoundFlavorsForProjectsListDto) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}

// GetProjectName returns the ProjectName field value
func (o *BoundFlavorsForProjectsListDto) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *BoundFlavorsForProjectsListDto) SetProjectName(v string) {
	o.ProjectName = v
}

// GetMaxDataDiskCount returns the MaxDataDiskCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCount() int32 {
	if o == nil || o.MaxDataDiskCount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxDataDiskCount.Get()
}

// GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDataDiskCount.Get(), o.MaxDataDiskCount.IsSet()
}

// SetMaxDataDiskCount sets field value
func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCount(v int32) {
	o.MaxDataDiskCount.Set(&v)
}

// GetExistsLinuxSpotPrice returns the ExistsLinuxSpotPrice field value
func (o *BoundFlavorsForProjectsListDto) GetExistsLinuxSpotPrice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExistsLinuxSpotPrice
}

// GetExistsLinuxSpotPriceOk returns a tuple with the ExistsLinuxSpotPrice field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetExistsLinuxSpotPriceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExistsLinuxSpotPrice, true
}

// SetExistsLinuxSpotPrice sets field value
func (o *BoundFlavorsForProjectsListDto) SetExistsLinuxSpotPrice(v bool) {
	o.ExistsLinuxSpotPrice = v
}

// GetExistsWindowsSpotPrice returns the ExistsWindowsSpotPrice field value
func (o *BoundFlavorsForProjectsListDto) GetExistsWindowsSpotPrice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExistsWindowsSpotPrice
}

// GetExistsWindowsSpotPriceOk returns a tuple with the ExistsWindowsSpotPrice field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetExistsWindowsSpotPriceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExistsWindowsSpotPrice, true
}

// SetExistsWindowsSpotPrice sets field value
func (o *BoundFlavorsForProjectsListDto) SetExistsWindowsSpotPrice(v bool) {
	o.ExistsWindowsSpotPrice = v
}

// GetLinuxSpotPrice returns the LinuxSpotPrice field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPrice() string {
	if o == nil || o.LinuxSpotPrice.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinuxSpotPrice.Get()
}

// GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinuxSpotPrice.Get(), o.LinuxSpotPrice.IsSet()
}

// SetLinuxSpotPrice sets field value
func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPrice(v string) {
	o.LinuxSpotPrice.Set(&v)
}

// GetLinuxPrice returns the LinuxPrice field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BoundFlavorsForProjectsListDto) GetLinuxPrice() string {
	if o == nil || o.LinuxPrice.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinuxPrice.Get()
}

// GetLinuxPriceOk returns a tuple with the LinuxPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetLinuxPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinuxPrice.Get(), o.LinuxPrice.IsSet()
}

// SetLinuxPrice sets field value
func (o *BoundFlavorsForProjectsListDto) SetLinuxPrice(v string) {
	o.LinuxPrice.Set(&v)
}

// GetWindowsSpotPrice returns the WindowsSpotPrice field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPrice() string {
	if o == nil || o.WindowsSpotPrice.Get() == nil {
		var ret string
		return ret
	}

	return *o.WindowsSpotPrice.Get()
}

// GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsSpotPrice.Get(), o.WindowsSpotPrice.IsSet()
}

// SetWindowsSpotPrice sets field value
func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPrice(v string) {
	o.WindowsSpotPrice.Set(&v)
}

// GetWindowsPrice returns the WindowsPrice field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BoundFlavorsForProjectsListDto) GetWindowsPrice() string {
	if o == nil || o.WindowsPrice.Get() == nil {
		var ret string
		return ret
	}

	return *o.WindowsPrice.Get()
}

// GetWindowsPriceOk returns a tuple with the WindowsPrice field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetWindowsPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsPrice.Get(), o.WindowsPrice.IsSet()
}

// SetWindowsPrice sets field value
func (o *BoundFlavorsForProjectsListDto) SetWindowsPrice(v string) {
	o.WindowsPrice.Set(&v)
}

// GetCloudType returns the CloudType field value
func (o *BoundFlavorsForProjectsListDto) GetCloudType() CloudType {
	if o == nil {
		var ret CloudType
		return ret
	}

	return o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetCloudTypeOk() (*CloudType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudType, true
}

// SetCloudType sets field value
func (o *BoundFlavorsForProjectsListDto) SetCloudType(v CloudType) {
	o.CloudType = v
}

// GetLocalDiskSize returns the LocalDiskSize field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetLocalDiskSize() int32 {
	if o == nil || IsNil(o.LocalDiskSize) {
		var ret int32
		return ret
	}
	return *o.LocalDiskSize
}

// GetLocalDiskSizeOk returns a tuple with the LocalDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetLocalDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalDiskSize) {
		return nil, false
	}
	return o.LocalDiskSize, true
}

// HasLocalDiskSize returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasLocalDiskSize() bool {
	if o != nil && !IsNil(o.LocalDiskSize) {
		return true
	}

	return false
}

// SetLocalDiskSize gets a reference to the given int32 and assigns it to the LocalDiskSize field.
func (o *BoundFlavorsForProjectsListDto) SetLocalDiskSize(v int32) {
	o.LocalDiskSize = &v
}

func (o BoundFlavorsForProjectsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoundFlavorsForProjectsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["cpu"] = o.Cpu
	toSerialize["ram"] = o.Ram
	toSerialize["projectId"] = o.ProjectId.Get()
	toSerialize["projectName"] = o.ProjectName
	toSerialize["maxDataDiskCount"] = o.MaxDataDiskCount.Get()
	toSerialize["existsLinuxSpotPrice"] = o.ExistsLinuxSpotPrice
	toSerialize["existsWindowsSpotPrice"] = o.ExistsWindowsSpotPrice
	toSerialize["linuxSpotPrice"] = o.LinuxSpotPrice.Get()
	toSerialize["linuxPrice"] = o.LinuxPrice.Get()
	toSerialize["windowsSpotPrice"] = o.WindowsSpotPrice.Get()
	toSerialize["windowsPrice"] = o.WindowsPrice.Get()
	toSerialize["cloudType"] = o.CloudType
	if !IsNil(o.LocalDiskSize) {
		toSerialize["localDiskSize"] = o.LocalDiskSize
	}
	return toSerialize, nil
}

func (o *BoundFlavorsForProjectsListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"cpu",
		"ram",
		"projectId",
		"projectName",
		"maxDataDiskCount",
		"existsLinuxSpotPrice",
		"existsWindowsSpotPrice",
		"linuxSpotPrice",
		"linuxPrice",
		"windowsSpotPrice",
		"windowsPrice",
		"cloudType",
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

	varBoundFlavorsForProjectsListDto := _BoundFlavorsForProjectsListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBoundFlavorsForProjectsListDto)

	if err != nil {
		return err
	}

	*o = BoundFlavorsForProjectsListDto(varBoundFlavorsForProjectsListDto)

	return err
}

type NullableBoundFlavorsForProjectsListDto struct {
	value *BoundFlavorsForProjectsListDto
	isSet bool
}

func (v NullableBoundFlavorsForProjectsListDto) Get() *BoundFlavorsForProjectsListDto {
	return v.value
}

func (v *NullableBoundFlavorsForProjectsListDto) Set(val *BoundFlavorsForProjectsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBoundFlavorsForProjectsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBoundFlavorsForProjectsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoundFlavorsForProjectsListDto(val *BoundFlavorsForProjectsListDto) *NullableBoundFlavorsForProjectsListDto {
	return &NullableBoundFlavorsForProjectsListDto{value: val, isSet: true}
}

func (v NullableBoundFlavorsForProjectsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoundFlavorsForProjectsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


