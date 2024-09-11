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

// checks if the ServerForCreateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerForCreateDto{}

// ServerForCreateDto struct for ServerForCreateDto
type ServerForCreateDto struct {
	Name *string `json:"name,omitempty"`
	Role *CloudRole `json:"role,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	DiskSize *float64 `json:"diskSize,omitempty"`
	Flavor *string `json:"flavor,omitempty"`
	Count *int32 `json:"count,omitempty"`
	SpotPrice NullableFloat64 `json:"spotPrice,omitempty"`
	SpotInstance *bool `json:"spotInstance,omitempty"`
	WasmEnabled *bool `json:"wasmEnabled,omitempty"`
	AutoscalingGroup *string `json:"autoscalingGroup,omitempty"`
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	ProxmoxExtraDiskSize *int32 `json:"proxmoxExtraDiskSize,omitempty"`
	ProxmoxRole *ProxmoxRole `json:"proxmoxRole,omitempty"`
	Hypervisor *string `json:"hypervisor,omitempty"`
	KubernetesNodeLabels []KubernetesNodeLabelsDto `json:"kubernetesNodeLabels,omitempty"`
	ReplicaCount NullableInt32 `json:"replicaCount,omitempty"`
}

// NewServerForCreateDto instantiates a new ServerForCreateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerForCreateDto() *ServerForCreateDto {
	this := ServerForCreateDto{}
	return &this
}

// NewServerForCreateDtoWithDefaults instantiates a new ServerForCreateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerForCreateDtoWithDefaults() *ServerForCreateDto {
	this := ServerForCreateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerForCreateDto) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetRole() CloudRole {
	if o == nil || IsNil(o.Role) {
		var ret CloudRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetRoleOk() (*CloudRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given CloudRole and assigns it to the Role field.
func (o *ServerForCreateDto) SetRole(v CloudRole) {
	o.Role = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *ServerForCreateDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetDiskSize() float64 {
	if o == nil || IsNil(o.DiskSize) {
		var ret float64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetDiskSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.DiskSize) {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasDiskSize() bool {
	if o != nil && !IsNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given float64 and assigns it to the DiskSize field.
func (o *ServerForCreateDto) SetDiskSize(v float64) {
	o.DiskSize = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetFlavor() string {
	if o == nil || IsNil(o.Flavor) {
		var ret string
		return ret
	}
	return *o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetFlavorOk() (*string, bool) {
	if o == nil || IsNil(o.Flavor) {
		return nil, false
	}
	return o.Flavor, true
}

// HasFlavor returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasFlavor() bool {
	if o != nil && !IsNil(o.Flavor) {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given string and assigns it to the Flavor field.
func (o *ServerForCreateDto) SetFlavor(v string) {
	o.Flavor = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ServerForCreateDto) SetCount(v int32) {
	o.Count = &v
}

// GetSpotPrice returns the SpotPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerForCreateDto) GetSpotPrice() float64 {
	if o == nil || IsNil(o.SpotPrice.Get()) {
		var ret float64
		return ret
	}
	return *o.SpotPrice.Get()
}

// GetSpotPriceOk returns a tuple with the SpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerForCreateDto) GetSpotPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpotPrice.Get(), o.SpotPrice.IsSet()
}

// HasSpotPrice returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasSpotPrice() bool {
	if o != nil && o.SpotPrice.IsSet() {
		return true
	}

	return false
}

// SetSpotPrice gets a reference to the given NullableFloat64 and assigns it to the SpotPrice field.
func (o *ServerForCreateDto) SetSpotPrice(v float64) {
	o.SpotPrice.Set(&v)
}
// SetSpotPriceNil sets the value for SpotPrice to be an explicit nil
func (o *ServerForCreateDto) SetSpotPriceNil() {
	o.SpotPrice.Set(nil)
}

// UnsetSpotPrice ensures that no value is present for SpotPrice, not even an explicit nil
func (o *ServerForCreateDto) UnsetSpotPrice() {
	o.SpotPrice.Unset()
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetSpotInstance() bool {
	if o == nil || IsNil(o.SpotInstance) {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.SpotInstance) {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasSpotInstance() bool {
	if o != nil && !IsNil(o.SpotInstance) {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *ServerForCreateDto) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetWasmEnabled returns the WasmEnabled field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetWasmEnabled() bool {
	if o == nil || IsNil(o.WasmEnabled) {
		var ret bool
		return ret
	}
	return *o.WasmEnabled
}

// GetWasmEnabledOk returns a tuple with the WasmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetWasmEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WasmEnabled) {
		return nil, false
	}
	return o.WasmEnabled, true
}

// HasWasmEnabled returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasWasmEnabled() bool {
	if o != nil && !IsNil(o.WasmEnabled) {
		return true
	}

	return false
}

// SetWasmEnabled gets a reference to the given bool and assigns it to the WasmEnabled field.
func (o *ServerForCreateDto) SetWasmEnabled(v bool) {
	o.WasmEnabled = &v
}

// GetAutoscalingGroup returns the AutoscalingGroup field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetAutoscalingGroup() string {
	if o == nil || IsNil(o.AutoscalingGroup) {
		var ret string
		return ret
	}
	return *o.AutoscalingGroup
}

// GetAutoscalingGroupOk returns a tuple with the AutoscalingGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetAutoscalingGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AutoscalingGroup) {
		return nil, false
	}
	return o.AutoscalingGroup, true
}

// HasAutoscalingGroup returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasAutoscalingGroup() bool {
	if o != nil && !IsNil(o.AutoscalingGroup) {
		return true
	}

	return false
}

// SetAutoscalingGroup gets a reference to the given string and assigns it to the AutoscalingGroup field.
func (o *ServerForCreateDto) SetAutoscalingGroup(v string) {
	o.AutoscalingGroup = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetAvailabilityZone() string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret string
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *ServerForCreateDto) SetAvailabilityZone(v string) {
	o.AvailabilityZone = &v
}

// GetProxmoxExtraDiskSize returns the ProxmoxExtraDiskSize field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetProxmoxExtraDiskSize() int32 {
	if o == nil || IsNil(o.ProxmoxExtraDiskSize) {
		var ret int32
		return ret
	}
	return *o.ProxmoxExtraDiskSize
}

// GetProxmoxExtraDiskSizeOk returns a tuple with the ProxmoxExtraDiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetProxmoxExtraDiskSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ProxmoxExtraDiskSize) {
		return nil, false
	}
	return o.ProxmoxExtraDiskSize, true
}

// HasProxmoxExtraDiskSize returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasProxmoxExtraDiskSize() bool {
	if o != nil && !IsNil(o.ProxmoxExtraDiskSize) {
		return true
	}

	return false
}

// SetProxmoxExtraDiskSize gets a reference to the given int32 and assigns it to the ProxmoxExtraDiskSize field.
func (o *ServerForCreateDto) SetProxmoxExtraDiskSize(v int32) {
	o.ProxmoxExtraDiskSize = &v
}

// GetProxmoxRole returns the ProxmoxRole field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetProxmoxRole() ProxmoxRole {
	if o == nil || IsNil(o.ProxmoxRole) {
		var ret ProxmoxRole
		return ret
	}
	return *o.ProxmoxRole
}

// GetProxmoxRoleOk returns a tuple with the ProxmoxRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetProxmoxRoleOk() (*ProxmoxRole, bool) {
	if o == nil || IsNil(o.ProxmoxRole) {
		return nil, false
	}
	return o.ProxmoxRole, true
}

// HasProxmoxRole returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasProxmoxRole() bool {
	if o != nil && !IsNil(o.ProxmoxRole) {
		return true
	}

	return false
}

// SetProxmoxRole gets a reference to the given ProxmoxRole and assigns it to the ProxmoxRole field.
func (o *ServerForCreateDto) SetProxmoxRole(v ProxmoxRole) {
	o.ProxmoxRole = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetHypervisor() string {
	if o == nil || IsNil(o.Hypervisor) {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetHypervisorOk() (*string, bool) {
	if o == nil || IsNil(o.Hypervisor) {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasHypervisor() bool {
	if o != nil && !IsNil(o.Hypervisor) {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *ServerForCreateDto) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetKubernetesNodeLabels returns the KubernetesNodeLabels field value if set, zero value otherwise.
func (o *ServerForCreateDto) GetKubernetesNodeLabels() []KubernetesNodeLabelsDto {
	if o == nil || IsNil(o.KubernetesNodeLabels) {
		var ret []KubernetesNodeLabelsDto
		return ret
	}
	return o.KubernetesNodeLabels
}

// GetKubernetesNodeLabelsOk returns a tuple with the KubernetesNodeLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerForCreateDto) GetKubernetesNodeLabelsOk() ([]KubernetesNodeLabelsDto, bool) {
	if o == nil || IsNil(o.KubernetesNodeLabels) {
		return nil, false
	}
	return o.KubernetesNodeLabels, true
}

// HasKubernetesNodeLabels returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasKubernetesNodeLabels() bool {
	if o != nil && !IsNil(o.KubernetesNodeLabels) {
		return true
	}

	return false
}

// SetKubernetesNodeLabels gets a reference to the given []KubernetesNodeLabelsDto and assigns it to the KubernetesNodeLabels field.
func (o *ServerForCreateDto) SetKubernetesNodeLabels(v []KubernetesNodeLabelsDto) {
	o.KubernetesNodeLabels = v
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServerForCreateDto) GetReplicaCount() int32 {
	if o == nil || IsNil(o.ReplicaCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReplicaCount.Get()
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServerForCreateDto) GetReplicaCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReplicaCount.Get(), o.ReplicaCount.IsSet()
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *ServerForCreateDto) HasReplicaCount() bool {
	if o != nil && o.ReplicaCount.IsSet() {
		return true
	}

	return false
}

// SetReplicaCount gets a reference to the given NullableInt32 and assigns it to the ReplicaCount field.
func (o *ServerForCreateDto) SetReplicaCount(v int32) {
	o.ReplicaCount.Set(&v)
}
// SetReplicaCountNil sets the value for ReplicaCount to be an explicit nil
func (o *ServerForCreateDto) SetReplicaCountNil() {
	o.ReplicaCount.Set(nil)
}

// UnsetReplicaCount ensures that no value is present for ReplicaCount, not even an explicit nil
func (o *ServerForCreateDto) UnsetReplicaCount() {
	o.ReplicaCount.Unset()
}

func (o ServerForCreateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerForCreateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.DiskSize) {
		toSerialize["diskSize"] = o.DiskSize
	}
	if !IsNil(o.Flavor) {
		toSerialize["flavor"] = o.Flavor
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.SpotPrice.IsSet() {
		toSerialize["spotPrice"] = o.SpotPrice.Get()
	}
	if !IsNil(o.SpotInstance) {
		toSerialize["spotInstance"] = o.SpotInstance
	}
	if !IsNil(o.WasmEnabled) {
		toSerialize["wasmEnabled"] = o.WasmEnabled
	}
	if !IsNil(o.AutoscalingGroup) {
		toSerialize["autoscalingGroup"] = o.AutoscalingGroup
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.ProxmoxExtraDiskSize) {
		toSerialize["proxmoxExtraDiskSize"] = o.ProxmoxExtraDiskSize
	}
	if !IsNil(o.ProxmoxRole) {
		toSerialize["proxmoxRole"] = o.ProxmoxRole
	}
	if !IsNil(o.Hypervisor) {
		toSerialize["hypervisor"] = o.Hypervisor
	}
	if !IsNil(o.KubernetesNodeLabels) {
		toSerialize["kubernetesNodeLabels"] = o.KubernetesNodeLabels
	}
	if o.ReplicaCount.IsSet() {
		toSerialize["replicaCount"] = o.ReplicaCount.Get()
	}
	return toSerialize, nil
}

type NullableServerForCreateDto struct {
	value *ServerForCreateDto
	isSet bool
}

func (v NullableServerForCreateDto) Get() *ServerForCreateDto {
	return v.value
}

func (v *NullableServerForCreateDto) Set(val *ServerForCreateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServerForCreateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServerForCreateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerForCreateDto(val *ServerForCreateDto) *NullableServerForCreateDto {
	return &NullableServerForCreateDto{value: val, isSet: true}
}

func (v NullableServerForCreateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerForCreateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


