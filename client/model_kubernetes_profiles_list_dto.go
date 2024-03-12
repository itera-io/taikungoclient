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

// checks if the KubernetesProfilesListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesProfilesListDto{}

// KubernetesProfilesListDto struct for KubernetesProfilesListDto
type KubernetesProfilesListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	Cni NullableString `json:"cni,omitempty"`
	OctaviaEnabled *bool `json:"octaviaEnabled,omitempty"`
	ExposeNodePortOnBastion *bool `json:"exposeNodePortOnBastion,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	TaikunLBEnabled *bool `json:"taikunLBEnabled,omitempty"`
	AllowSchedulingOnMaster *bool `json:"allowSchedulingOnMaster,omitempty"`
	UniqueClusterName *bool `json:"uniqueClusterName,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	LastModified NullableString `json:"lastModified,omitempty"`
	LastModifiedBy NullableString `json:"lastModifiedBy,omitempty"`
	ProxmoxStorage NullableString `json:"proxmoxStorage,omitempty"`
	NvidiaGpuOperatorEnabled *bool `json:"nvidiaGpuOperatorEnabled,omitempty"`
	WasmEnabled *bool `json:"wasmEnabled,omitempty"`
}

// NewKubernetesProfilesListDto instantiates a new KubernetesProfilesListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesProfilesListDto() *KubernetesProfilesListDto {
	this := KubernetesProfilesListDto{}
	return &this
}

// NewKubernetesProfilesListDtoWithDefaults instantiates a new KubernetesProfilesListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesProfilesListDtoWithDefaults() *KubernetesProfilesListDto {
	this := KubernetesProfilesListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KubernetesProfilesListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KubernetesProfilesListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *KubernetesProfilesListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *KubernetesProfilesListDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *KubernetesProfilesListDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *KubernetesProfilesListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *KubernetesProfilesListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetCni returns the Cni field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetCni() string {
	if o == nil || IsNil(o.Cni.Get()) {
		var ret string
		return ret
	}
	return *o.Cni.Get()
}

// GetCniOk returns a tuple with the Cni field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetCniOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cni.Get(), o.Cni.IsSet()
}

// HasCni returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasCni() bool {
	if o != nil && o.Cni.IsSet() {
		return true
	}

	return false
}

// SetCni gets a reference to the given NullableString and assigns it to the Cni field.
func (o *KubernetesProfilesListDto) SetCni(v string) {
	o.Cni.Set(&v)
}
// SetCniNil sets the value for Cni to be an explicit nil
func (o *KubernetesProfilesListDto) SetCniNil() {
	o.Cni.Set(nil)
}

// UnsetCni ensures that no value is present for Cni, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetCni() {
	o.Cni.Unset()
}

// GetOctaviaEnabled returns the OctaviaEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetOctaviaEnabled() bool {
	if o == nil || IsNil(o.OctaviaEnabled) {
		var ret bool
		return ret
	}
	return *o.OctaviaEnabled
}

// GetOctaviaEnabledOk returns a tuple with the OctaviaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetOctaviaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OctaviaEnabled) {
		return nil, false
	}
	return o.OctaviaEnabled, true
}

// HasOctaviaEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasOctaviaEnabled() bool {
	if o != nil && !IsNil(o.OctaviaEnabled) {
		return true
	}

	return false
}

// SetOctaviaEnabled gets a reference to the given bool and assigns it to the OctaviaEnabled field.
func (o *KubernetesProfilesListDto) SetOctaviaEnabled(v bool) {
	o.OctaviaEnabled = &v
}

// GetExposeNodePortOnBastion returns the ExposeNodePortOnBastion field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetExposeNodePortOnBastion() bool {
	if o == nil || IsNil(o.ExposeNodePortOnBastion) {
		var ret bool
		return ret
	}
	return *o.ExposeNodePortOnBastion
}

// GetExposeNodePortOnBastionOk returns a tuple with the ExposeNodePortOnBastion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetExposeNodePortOnBastionOk() (*bool, bool) {
	if o == nil || IsNil(o.ExposeNodePortOnBastion) {
		return nil, false
	}
	return o.ExposeNodePortOnBastion, true
}

// HasExposeNodePortOnBastion returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasExposeNodePortOnBastion() bool {
	if o != nil && !IsNil(o.ExposeNodePortOnBastion) {
		return true
	}

	return false
}

// SetExposeNodePortOnBastion gets a reference to the given bool and assigns it to the ExposeNodePortOnBastion field.
func (o *KubernetesProfilesListDto) SetExposeNodePortOnBastion(v bool) {
	o.ExposeNodePortOnBastion = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *KubernetesProfilesListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetTaikunLBEnabled returns the TaikunLBEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetTaikunLBEnabled() bool {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		var ret bool
		return ret
	}
	return *o.TaikunLBEnabled
}

// GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetTaikunLBEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TaikunLBEnabled) {
		return nil, false
	}
	return o.TaikunLBEnabled, true
}

// HasTaikunLBEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasTaikunLBEnabled() bool {
	if o != nil && !IsNil(o.TaikunLBEnabled) {
		return true
	}

	return false
}

// SetTaikunLBEnabled gets a reference to the given bool and assigns it to the TaikunLBEnabled field.
func (o *KubernetesProfilesListDto) SetTaikunLBEnabled(v bool) {
	o.TaikunLBEnabled = &v
}

// GetAllowSchedulingOnMaster returns the AllowSchedulingOnMaster field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetAllowSchedulingOnMaster() bool {
	if o == nil || IsNil(o.AllowSchedulingOnMaster) {
		var ret bool
		return ret
	}
	return *o.AllowSchedulingOnMaster
}

// GetAllowSchedulingOnMasterOk returns a tuple with the AllowSchedulingOnMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetAllowSchedulingOnMasterOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSchedulingOnMaster) {
		return nil, false
	}
	return o.AllowSchedulingOnMaster, true
}

// HasAllowSchedulingOnMaster returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasAllowSchedulingOnMaster() bool {
	if o != nil && !IsNil(o.AllowSchedulingOnMaster) {
		return true
	}

	return false
}

// SetAllowSchedulingOnMaster gets a reference to the given bool and assigns it to the AllowSchedulingOnMaster field.
func (o *KubernetesProfilesListDto) SetAllowSchedulingOnMaster(v bool) {
	o.AllowSchedulingOnMaster = &v
}

// GetUniqueClusterName returns the UniqueClusterName field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetUniqueClusterName() bool {
	if o == nil || IsNil(o.UniqueClusterName) {
		var ret bool
		return ret
	}
	return *o.UniqueClusterName
}

// GetUniqueClusterNameOk returns a tuple with the UniqueClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetUniqueClusterNameOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueClusterName) {
		return nil, false
	}
	return o.UniqueClusterName, true
}

// HasUniqueClusterName returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasUniqueClusterName() bool {
	if o != nil && !IsNil(o.UniqueClusterName) {
		return true
	}

	return false
}

// SetUniqueClusterName gets a reference to the given bool and assigns it to the UniqueClusterName field.
func (o *KubernetesProfilesListDto) SetUniqueClusterName(v bool) {
	o.UniqueClusterName = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *KubernetesProfilesListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *KubernetesProfilesListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *KubernetesProfilesListDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *KubernetesProfilesListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *KubernetesProfilesListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *KubernetesProfilesListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}
// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *KubernetesProfilesListDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *KubernetesProfilesListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}
// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *KubernetesProfilesListDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetProxmoxStorage returns the ProxmoxStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesProfilesListDto) GetProxmoxStorage() string {
	if o == nil || IsNil(o.ProxmoxStorage.Get()) {
		var ret string
		return ret
	}
	return *o.ProxmoxStorage.Get()
}

// GetProxmoxStorageOk returns a tuple with the ProxmoxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesProfilesListDto) GetProxmoxStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxmoxStorage.Get(), o.ProxmoxStorage.IsSet()
}

// HasProxmoxStorage returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasProxmoxStorage() bool {
	if o != nil && o.ProxmoxStorage.IsSet() {
		return true
	}

	return false
}

// SetProxmoxStorage gets a reference to the given NullableString and assigns it to the ProxmoxStorage field.
func (o *KubernetesProfilesListDto) SetProxmoxStorage(v string) {
	o.ProxmoxStorage.Set(&v)
}
// SetProxmoxStorageNil sets the value for ProxmoxStorage to be an explicit nil
func (o *KubernetesProfilesListDto) SetProxmoxStorageNil() {
	o.ProxmoxStorage.Set(nil)
}

// UnsetProxmoxStorage ensures that no value is present for ProxmoxStorage, not even an explicit nil
func (o *KubernetesProfilesListDto) UnsetProxmoxStorage() {
	o.ProxmoxStorage.Unset()
}

// GetNvidiaGpuOperatorEnabled returns the NvidiaGpuOperatorEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetNvidiaGpuOperatorEnabled() bool {
	if o == nil || IsNil(o.NvidiaGpuOperatorEnabled) {
		var ret bool
		return ret
	}
	return *o.NvidiaGpuOperatorEnabled
}

// GetNvidiaGpuOperatorEnabledOk returns a tuple with the NvidiaGpuOperatorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetNvidiaGpuOperatorEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NvidiaGpuOperatorEnabled) {
		return nil, false
	}
	return o.NvidiaGpuOperatorEnabled, true
}

// HasNvidiaGpuOperatorEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasNvidiaGpuOperatorEnabled() bool {
	if o != nil && !IsNil(o.NvidiaGpuOperatorEnabled) {
		return true
	}

	return false
}

// SetNvidiaGpuOperatorEnabled gets a reference to the given bool and assigns it to the NvidiaGpuOperatorEnabled field.
func (o *KubernetesProfilesListDto) SetNvidiaGpuOperatorEnabled(v bool) {
	o.NvidiaGpuOperatorEnabled = &v
}

// GetWasmEnabled returns the WasmEnabled field value if set, zero value otherwise.
func (o *KubernetesProfilesListDto) GetWasmEnabled() bool {
	if o == nil || IsNil(o.WasmEnabled) {
		var ret bool
		return ret
	}
	return *o.WasmEnabled
}

// GetWasmEnabledOk returns a tuple with the WasmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesProfilesListDto) GetWasmEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WasmEnabled) {
		return nil, false
	}
	return o.WasmEnabled, true
}

// HasWasmEnabled returns a boolean if a field has been set.
func (o *KubernetesProfilesListDto) HasWasmEnabled() bool {
	if o != nil && !IsNil(o.WasmEnabled) {
		return true
	}

	return false
}

// SetWasmEnabled gets a reference to the given bool and assigns it to the WasmEnabled field.
func (o *KubernetesProfilesListDto) SetWasmEnabled(v bool) {
	o.WasmEnabled = &v
}

func (o KubernetesProfilesListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesProfilesListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.Cni.IsSet() {
		toSerialize["cni"] = o.Cni.Get()
	}
	if !IsNil(o.OctaviaEnabled) {
		toSerialize["octaviaEnabled"] = o.OctaviaEnabled
	}
	if !IsNil(o.ExposeNodePortOnBastion) {
		toSerialize["exposeNodePortOnBastion"] = o.ExposeNodePortOnBastion
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.TaikunLBEnabled) {
		toSerialize["taikunLBEnabled"] = o.TaikunLBEnabled
	}
	if !IsNil(o.AllowSchedulingOnMaster) {
		toSerialize["allowSchedulingOnMaster"] = o.AllowSchedulingOnMaster
	}
	if !IsNil(o.UniqueClusterName) {
		toSerialize["uniqueClusterName"] = o.UniqueClusterName
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if o.ProxmoxStorage.IsSet() {
		toSerialize["proxmoxStorage"] = o.ProxmoxStorage.Get()
	}
	if !IsNil(o.NvidiaGpuOperatorEnabled) {
		toSerialize["nvidiaGpuOperatorEnabled"] = o.NvidiaGpuOperatorEnabled
	}
	if !IsNil(o.WasmEnabled) {
		toSerialize["wasmEnabled"] = o.WasmEnabled
	}
	return toSerialize, nil
}

type NullableKubernetesProfilesListDto struct {
	value *KubernetesProfilesListDto
	isSet bool
}

func (v NullableKubernetesProfilesListDto) Get() *KubernetesProfilesListDto {
	return v.value
}

func (v *NullableKubernetesProfilesListDto) Set(val *KubernetesProfilesListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesProfilesListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesProfilesListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesProfilesListDto(val *KubernetesProfilesListDto) *NullableKubernetesProfilesListDto {
	return &NullableKubernetesProfilesListDto{value: val, isSet: true}
}

func (v NullableKubernetesProfilesListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesProfilesListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


