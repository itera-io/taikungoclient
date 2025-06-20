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
	"time"
)

// checks if the CreateVirtualClusterCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVirtualClusterCommand{}

// CreateVirtualClusterCommand struct for CreateVirtualClusterCommand
type CreateVirtualClusterCommand struct {
	CloudId NullableInt32 `json:"cloudId,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	Name *string `json:"name,omitempty"`
	ExpiredAt NullableTime `json:"expiredAt,omitempty"`
	DeleteOnExpiration *bool `json:"deleteOnExpiration,omitempty"`
	AlertingProfileId NullableInt32 `json:"alertingProfileId,omitempty"`
	ExposeHostname *string `json:"exposeHostname,omitempty"`
	WorkloadResources *CreateVirtualClusterWorkloadResources `json:"workloadResources,omitempty"`
	CpuLimits *CreateVirtualClusterResourceLimits `json:"cpuLimits,omitempty"`
	RamLimits *CreateVirtualClusterResourceLimits `json:"ramLimits,omitempty"`
	EphemeralStorageLimits *CreateVirtualClusterResourceLimits `json:"ephemeralStorageLimits,omitempty"`
}

// NewCreateVirtualClusterCommand instantiates a new CreateVirtualClusterCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVirtualClusterCommand() *CreateVirtualClusterCommand {
	this := CreateVirtualClusterCommand{}
	return &this
}

// NewCreateVirtualClusterCommandWithDefaults instantiates a new CreateVirtualClusterCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVirtualClusterCommandWithDefaults() *CreateVirtualClusterCommand {
	this := CreateVirtualClusterCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVirtualClusterCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId.Get()) {
		var ret int32
		return ret
	}
	return *o.CloudId.Get()
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVirtualClusterCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudId.Get(), o.CloudId.IsSet()
}

// HasCloudId returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasCloudId() bool {
	if o != nil && o.CloudId.IsSet() {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given NullableInt32 and assigns it to the CloudId field.
func (o *CreateVirtualClusterCommand) SetCloudId(v int32) {
	o.CloudId.Set(&v)
}
// SetCloudIdNil sets the value for CloudId to be an explicit nil
func (o *CreateVirtualClusterCommand) SetCloudIdNil() {
	o.CloudId.Set(nil)
}

// UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
func (o *CreateVirtualClusterCommand) UnsetCloudId() {
	o.CloudId.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CreateVirtualClusterCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateVirtualClusterCommand) SetName(v string) {
	o.Name = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVirtualClusterCommand) GetExpiredAt() time.Time {
	if o == nil || IsNil(o.ExpiredAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVirtualClusterCommand) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given NullableTime and assigns it to the ExpiredAt field.
func (o *CreateVirtualClusterCommand) SetExpiredAt(v time.Time) {
	o.ExpiredAt.Set(&v)
}
// SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil
func (o *CreateVirtualClusterCommand) SetExpiredAtNil() {
	o.ExpiredAt.Set(nil)
}

// UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
func (o *CreateVirtualClusterCommand) UnsetExpiredAt() {
	o.ExpiredAt.Unset()
}

// GetDeleteOnExpiration returns the DeleteOnExpiration field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetDeleteOnExpiration() bool {
	if o == nil || IsNil(o.DeleteOnExpiration) {
		var ret bool
		return ret
	}
	return *o.DeleteOnExpiration
}

// GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetDeleteOnExpirationOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteOnExpiration) {
		return nil, false
	}
	return o.DeleteOnExpiration, true
}

// HasDeleteOnExpiration returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasDeleteOnExpiration() bool {
	if o != nil && !IsNil(o.DeleteOnExpiration) {
		return true
	}

	return false
}

// SetDeleteOnExpiration gets a reference to the given bool and assigns it to the DeleteOnExpiration field.
func (o *CreateVirtualClusterCommand) SetDeleteOnExpiration(v bool) {
	o.DeleteOnExpiration = &v
}

// GetAlertingProfileId returns the AlertingProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVirtualClusterCommand) GetAlertingProfileId() int32 {
	if o == nil || IsNil(o.AlertingProfileId.Get()) {
		var ret int32
		return ret
	}
	return *o.AlertingProfileId.Get()
}

// GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVirtualClusterCommand) GetAlertingProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertingProfileId.Get(), o.AlertingProfileId.IsSet()
}

// HasAlertingProfileId returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasAlertingProfileId() bool {
	if o != nil && o.AlertingProfileId.IsSet() {
		return true
	}

	return false
}

// SetAlertingProfileId gets a reference to the given NullableInt32 and assigns it to the AlertingProfileId field.
func (o *CreateVirtualClusterCommand) SetAlertingProfileId(v int32) {
	o.AlertingProfileId.Set(&v)
}
// SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil
func (o *CreateVirtualClusterCommand) SetAlertingProfileIdNil() {
	o.AlertingProfileId.Set(nil)
}

// UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
func (o *CreateVirtualClusterCommand) UnsetAlertingProfileId() {
	o.AlertingProfileId.Unset()
}

// GetExposeHostname returns the ExposeHostname field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetExposeHostname() string {
	if o == nil || IsNil(o.ExposeHostname) {
		var ret string
		return ret
	}
	return *o.ExposeHostname
}

// GetExposeHostnameOk returns a tuple with the ExposeHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetExposeHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ExposeHostname) {
		return nil, false
	}
	return o.ExposeHostname, true
}

// HasExposeHostname returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasExposeHostname() bool {
	if o != nil && !IsNil(o.ExposeHostname) {
		return true
	}

	return false
}

// SetExposeHostname gets a reference to the given string and assigns it to the ExposeHostname field.
func (o *CreateVirtualClusterCommand) SetExposeHostname(v string) {
	o.ExposeHostname = &v
}

// GetWorkloadResources returns the WorkloadResources field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetWorkloadResources() CreateVirtualClusterWorkloadResources {
	if o == nil || IsNil(o.WorkloadResources) {
		var ret CreateVirtualClusterWorkloadResources
		return ret
	}
	return *o.WorkloadResources
}

// GetWorkloadResourcesOk returns a tuple with the WorkloadResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetWorkloadResourcesOk() (*CreateVirtualClusterWorkloadResources, bool) {
	if o == nil || IsNil(o.WorkloadResources) {
		return nil, false
	}
	return o.WorkloadResources, true
}

// HasWorkloadResources returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasWorkloadResources() bool {
	if o != nil && !IsNil(o.WorkloadResources) {
		return true
	}

	return false
}

// SetWorkloadResources gets a reference to the given CreateVirtualClusterWorkloadResources and assigns it to the WorkloadResources field.
func (o *CreateVirtualClusterCommand) SetWorkloadResources(v CreateVirtualClusterWorkloadResources) {
	o.WorkloadResources = &v
}

// GetCpuLimits returns the CpuLimits field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetCpuLimits() CreateVirtualClusterResourceLimits {
	if o == nil || IsNil(o.CpuLimits) {
		var ret CreateVirtualClusterResourceLimits
		return ret
	}
	return *o.CpuLimits
}

// GetCpuLimitsOk returns a tuple with the CpuLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetCpuLimitsOk() (*CreateVirtualClusterResourceLimits, bool) {
	if o == nil || IsNil(o.CpuLimits) {
		return nil, false
	}
	return o.CpuLimits, true
}

// HasCpuLimits returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasCpuLimits() bool {
	if o != nil && !IsNil(o.CpuLimits) {
		return true
	}

	return false
}

// SetCpuLimits gets a reference to the given CreateVirtualClusterResourceLimits and assigns it to the CpuLimits field.
func (o *CreateVirtualClusterCommand) SetCpuLimits(v CreateVirtualClusterResourceLimits) {
	o.CpuLimits = &v
}

// GetRamLimits returns the RamLimits field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetRamLimits() CreateVirtualClusterResourceLimits {
	if o == nil || IsNil(o.RamLimits) {
		var ret CreateVirtualClusterResourceLimits
		return ret
	}
	return *o.RamLimits
}

// GetRamLimitsOk returns a tuple with the RamLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetRamLimitsOk() (*CreateVirtualClusterResourceLimits, bool) {
	if o == nil || IsNil(o.RamLimits) {
		return nil, false
	}
	return o.RamLimits, true
}

// HasRamLimits returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasRamLimits() bool {
	if o != nil && !IsNil(o.RamLimits) {
		return true
	}

	return false
}

// SetRamLimits gets a reference to the given CreateVirtualClusterResourceLimits and assigns it to the RamLimits field.
func (o *CreateVirtualClusterCommand) SetRamLimits(v CreateVirtualClusterResourceLimits) {
	o.RamLimits = &v
}

// GetEphemeralStorageLimits returns the EphemeralStorageLimits field value if set, zero value otherwise.
func (o *CreateVirtualClusterCommand) GetEphemeralStorageLimits() CreateVirtualClusterResourceLimits {
	if o == nil || IsNil(o.EphemeralStorageLimits) {
		var ret CreateVirtualClusterResourceLimits
		return ret
	}
	return *o.EphemeralStorageLimits
}

// GetEphemeralStorageLimitsOk returns a tuple with the EphemeralStorageLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVirtualClusterCommand) GetEphemeralStorageLimitsOk() (*CreateVirtualClusterResourceLimits, bool) {
	if o == nil || IsNil(o.EphemeralStorageLimits) {
		return nil, false
	}
	return o.EphemeralStorageLimits, true
}

// HasEphemeralStorageLimits returns a boolean if a field has been set.
func (o *CreateVirtualClusterCommand) HasEphemeralStorageLimits() bool {
	if o != nil && !IsNil(o.EphemeralStorageLimits) {
		return true
	}

	return false
}

// SetEphemeralStorageLimits gets a reference to the given CreateVirtualClusterResourceLimits and assigns it to the EphemeralStorageLimits field.
func (o *CreateVirtualClusterCommand) SetEphemeralStorageLimits(v CreateVirtualClusterResourceLimits) {
	o.EphemeralStorageLimits = &v
}

func (o CreateVirtualClusterCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVirtualClusterCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudId.IsSet() {
		toSerialize["cloudId"] = o.CloudId.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.ExpiredAt.IsSet() {
		toSerialize["expiredAt"] = o.ExpiredAt.Get()
	}
	if !IsNil(o.DeleteOnExpiration) {
		toSerialize["deleteOnExpiration"] = o.DeleteOnExpiration
	}
	if o.AlertingProfileId.IsSet() {
		toSerialize["alertingProfileId"] = o.AlertingProfileId.Get()
	}
	if !IsNil(o.ExposeHostname) {
		toSerialize["exposeHostname"] = o.ExposeHostname
	}
	if !IsNil(o.WorkloadResources) {
		toSerialize["workloadResources"] = o.WorkloadResources
	}
	if !IsNil(o.CpuLimits) {
		toSerialize["cpuLimits"] = o.CpuLimits
	}
	if !IsNil(o.RamLimits) {
		toSerialize["ramLimits"] = o.RamLimits
	}
	if !IsNil(o.EphemeralStorageLimits) {
		toSerialize["ephemeralStorageLimits"] = o.EphemeralStorageLimits
	}
	return toSerialize, nil
}

type NullableCreateVirtualClusterCommand struct {
	value *CreateVirtualClusterCommand
	isSet bool
}

func (v NullableCreateVirtualClusterCommand) Get() *CreateVirtualClusterCommand {
	return v.value
}

func (v *NullableCreateVirtualClusterCommand) Set(val *CreateVirtualClusterCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVirtualClusterCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVirtualClusterCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVirtualClusterCommand(val *CreateVirtualClusterCommand) *NullableCreateVirtualClusterCommand {
	return &NullableCreateVirtualClusterCommand{value: val, isSet: true}
}

func (v NullableCreateVirtualClusterCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVirtualClusterCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


