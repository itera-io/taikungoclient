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

// checks if the BoundFlavorsForProjectsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoundFlavorsForProjectsListDto{}

// BoundFlavorsForProjectsListDto struct for BoundFlavorsForProjectsListDto
type BoundFlavorsForProjectsListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Cpu *int32 `json:"cpu,omitempty"`
	Ram *int64 `json:"ram,omitempty"`
	ProjectId NullableInt32 `json:"projectId,omitempty"`
	IsAzure *bool `json:"isAzure,omitempty"`
	IsAws *bool `json:"isAws,omitempty"`
	IsOpenstack *bool `json:"isOpenstack,omitempty"`
	IsGoogle *bool `json:"isGoogle,omitempty"`
	IsProxmox *bool `json:"isProxmox,omitempty"`
	IsTanzu *bool `json:"isTanzu,omitempty"`
	IsOpenshift *bool `json:"isOpenshift,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	MaxDataDiskCount NullableInt32 `json:"maxDataDiskCount,omitempty"`
	ExistsLinuxSpotPrice *bool `json:"existsLinuxSpotPrice,omitempty"`
	ExistsWindowsSpotPrice *bool `json:"existsWindowsSpotPrice,omitempty"`
	LinuxSpotPrice NullableString `json:"linuxSpotPrice,omitempty"`
	LinuxPrice NullableString `json:"linuxPrice,omitempty"`
	WindowsSpotPrice NullableString `json:"windowsSpotPrice,omitempty"`
	WindowsPrice NullableString `json:"windowsPrice,omitempty"`
}

// NewBoundFlavorsForProjectsListDto instantiates a new BoundFlavorsForProjectsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoundFlavorsForProjectsListDto() *BoundFlavorsForProjectsListDto {
	this := BoundFlavorsForProjectsListDto{}
	return &this
}

// NewBoundFlavorsForProjectsListDtoWithDefaults instantiates a new BoundFlavorsForProjectsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoundFlavorsForProjectsListDtoWithDefaults() *BoundFlavorsForProjectsListDto {
	this := BoundFlavorsForProjectsListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BoundFlavorsForProjectsListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BoundFlavorsForProjectsListDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetName() {
	o.Name.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *BoundFlavorsForProjectsListDto) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetRam() int64 {
	if o == nil || IsNil(o.Ram) {
		var ret int64
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetRamOk() (*int64, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given int64 and assigns it to the Ram field.
func (o *BoundFlavorsForProjectsListDto) SetRam(v int64) {
	o.Ram = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt32 and assigns it to the ProjectId field.
func (o *BoundFlavorsForProjectsListDto) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetIsAzure returns the IsAzure field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsAzure() bool {
	if o == nil || IsNil(o.IsAzure) {
		var ret bool
		return ret
	}
	return *o.IsAzure
}

// GetIsAzureOk returns a tuple with the IsAzure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsAzureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAzure) {
		return nil, false
	}
	return o.IsAzure, true
}

// HasIsAzure returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsAzure() bool {
	if o != nil && !IsNil(o.IsAzure) {
		return true
	}

	return false
}

// SetIsAzure gets a reference to the given bool and assigns it to the IsAzure field.
func (o *BoundFlavorsForProjectsListDto) SetIsAzure(v bool) {
	o.IsAzure = &v
}

// GetIsAws returns the IsAws field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsAws() bool {
	if o == nil || IsNil(o.IsAws) {
		var ret bool
		return ret
	}
	return *o.IsAws
}

// GetIsAwsOk returns a tuple with the IsAws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsAwsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAws) {
		return nil, false
	}
	return o.IsAws, true
}

// HasIsAws returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsAws() bool {
	if o != nil && !IsNil(o.IsAws) {
		return true
	}

	return false
}

// SetIsAws gets a reference to the given bool and assigns it to the IsAws field.
func (o *BoundFlavorsForProjectsListDto) SetIsAws(v bool) {
	o.IsAws = &v
}

// GetIsOpenstack returns the IsOpenstack field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsOpenstack() bool {
	if o == nil || IsNil(o.IsOpenstack) {
		var ret bool
		return ret
	}
	return *o.IsOpenstack
}

// GetIsOpenstackOk returns a tuple with the IsOpenstack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsOpenstackOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpenstack) {
		return nil, false
	}
	return o.IsOpenstack, true
}

// HasIsOpenstack returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsOpenstack() bool {
	if o != nil && !IsNil(o.IsOpenstack) {
		return true
	}

	return false
}

// SetIsOpenstack gets a reference to the given bool and assigns it to the IsOpenstack field.
func (o *BoundFlavorsForProjectsListDto) SetIsOpenstack(v bool) {
	o.IsOpenstack = &v
}

// GetIsGoogle returns the IsGoogle field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsGoogle() bool {
	if o == nil || IsNil(o.IsGoogle) {
		var ret bool
		return ret
	}
	return *o.IsGoogle
}

// GetIsGoogleOk returns a tuple with the IsGoogle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsGoogleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGoogle) {
		return nil, false
	}
	return o.IsGoogle, true
}

// HasIsGoogle returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsGoogle() bool {
	if o != nil && !IsNil(o.IsGoogle) {
		return true
	}

	return false
}

// SetIsGoogle gets a reference to the given bool and assigns it to the IsGoogle field.
func (o *BoundFlavorsForProjectsListDto) SetIsGoogle(v bool) {
	o.IsGoogle = &v
}

// GetIsProxmox returns the IsProxmox field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsProxmox() bool {
	if o == nil || IsNil(o.IsProxmox) {
		var ret bool
		return ret
	}
	return *o.IsProxmox
}

// GetIsProxmoxOk returns a tuple with the IsProxmox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsProxmoxOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProxmox) {
		return nil, false
	}
	return o.IsProxmox, true
}

// HasIsProxmox returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsProxmox() bool {
	if o != nil && !IsNil(o.IsProxmox) {
		return true
	}

	return false
}

// SetIsProxmox gets a reference to the given bool and assigns it to the IsProxmox field.
func (o *BoundFlavorsForProjectsListDto) SetIsProxmox(v bool) {
	o.IsProxmox = &v
}

// GetIsTanzu returns the IsTanzu field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsTanzu() bool {
	if o == nil || IsNil(o.IsTanzu) {
		var ret bool
		return ret
	}
	return *o.IsTanzu
}

// GetIsTanzuOk returns a tuple with the IsTanzu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsTanzuOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTanzu) {
		return nil, false
	}
	return o.IsTanzu, true
}

// HasIsTanzu returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsTanzu() bool {
	if o != nil && !IsNil(o.IsTanzu) {
		return true
	}

	return false
}

// SetIsTanzu gets a reference to the given bool and assigns it to the IsTanzu field.
func (o *BoundFlavorsForProjectsListDto) SetIsTanzu(v bool) {
	o.IsTanzu = &v
}

// GetIsOpenshift returns the IsOpenshift field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetIsOpenshift() bool {
	if o == nil || IsNil(o.IsOpenshift) {
		var ret bool
		return ret
	}
	return *o.IsOpenshift
}

// GetIsOpenshiftOk returns a tuple with the IsOpenshift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetIsOpenshiftOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpenshift) {
		return nil, false
	}
	return o.IsOpenshift, true
}

// HasIsOpenshift returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasIsOpenshift() bool {
	if o != nil && !IsNil(o.IsOpenshift) {
		return true
	}

	return false
}

// SetIsOpenshift gets a reference to the given bool and assigns it to the IsOpenshift field.
func (o *BoundFlavorsForProjectsListDto) SetIsOpenshift(v bool) {
	o.IsOpenshift = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *BoundFlavorsForProjectsListDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetMaxDataDiskCount returns the MaxDataDiskCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCount() int32 {
	if o == nil || IsNil(o.MaxDataDiskCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxDataDiskCount.Get()
}

// GetMaxDataDiskCountOk returns a tuple with the MaxDataDiskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetMaxDataDiskCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDataDiskCount.Get(), o.MaxDataDiskCount.IsSet()
}

// HasMaxDataDiskCount returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasMaxDataDiskCount() bool {
	if o != nil && o.MaxDataDiskCount.IsSet() {
		return true
	}

	return false
}

// SetMaxDataDiskCount gets a reference to the given NullableInt32 and assigns it to the MaxDataDiskCount field.
func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCount(v int32) {
	o.MaxDataDiskCount.Set(&v)
}
// SetMaxDataDiskCountNil sets the value for MaxDataDiskCount to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetMaxDataDiskCountNil() {
	o.MaxDataDiskCount.Set(nil)
}

// UnsetMaxDataDiskCount ensures that no value is present for MaxDataDiskCount, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetMaxDataDiskCount() {
	o.MaxDataDiskCount.Unset()
}

// GetExistsLinuxSpotPrice returns the ExistsLinuxSpotPrice field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetExistsLinuxSpotPrice() bool {
	if o == nil || IsNil(o.ExistsLinuxSpotPrice) {
		var ret bool
		return ret
	}
	return *o.ExistsLinuxSpotPrice
}

// GetExistsLinuxSpotPriceOk returns a tuple with the ExistsLinuxSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetExistsLinuxSpotPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.ExistsLinuxSpotPrice) {
		return nil, false
	}
	return o.ExistsLinuxSpotPrice, true
}

// HasExistsLinuxSpotPrice returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasExistsLinuxSpotPrice() bool {
	if o != nil && !IsNil(o.ExistsLinuxSpotPrice) {
		return true
	}

	return false
}

// SetExistsLinuxSpotPrice gets a reference to the given bool and assigns it to the ExistsLinuxSpotPrice field.
func (o *BoundFlavorsForProjectsListDto) SetExistsLinuxSpotPrice(v bool) {
	o.ExistsLinuxSpotPrice = &v
}

// GetExistsWindowsSpotPrice returns the ExistsWindowsSpotPrice field value if set, zero value otherwise.
func (o *BoundFlavorsForProjectsListDto) GetExistsWindowsSpotPrice() bool {
	if o == nil || IsNil(o.ExistsWindowsSpotPrice) {
		var ret bool
		return ret
	}
	return *o.ExistsWindowsSpotPrice
}

// GetExistsWindowsSpotPriceOk returns a tuple with the ExistsWindowsSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoundFlavorsForProjectsListDto) GetExistsWindowsSpotPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.ExistsWindowsSpotPrice) {
		return nil, false
	}
	return o.ExistsWindowsSpotPrice, true
}

// HasExistsWindowsSpotPrice returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasExistsWindowsSpotPrice() bool {
	if o != nil && !IsNil(o.ExistsWindowsSpotPrice) {
		return true
	}

	return false
}

// SetExistsWindowsSpotPrice gets a reference to the given bool and assigns it to the ExistsWindowsSpotPrice field.
func (o *BoundFlavorsForProjectsListDto) SetExistsWindowsSpotPrice(v bool) {
	o.ExistsWindowsSpotPrice = &v
}

// GetLinuxSpotPrice returns the LinuxSpotPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPrice() string {
	if o == nil || IsNil(o.LinuxSpotPrice.Get()) {
		var ret string
		return ret
	}
	return *o.LinuxSpotPrice.Get()
}

// GetLinuxSpotPriceOk returns a tuple with the LinuxSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetLinuxSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinuxSpotPrice.Get(), o.LinuxSpotPrice.IsSet()
}

// HasLinuxSpotPrice returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasLinuxSpotPrice() bool {
	if o != nil && o.LinuxSpotPrice.IsSet() {
		return true
	}

	return false
}

// SetLinuxSpotPrice gets a reference to the given NullableString and assigns it to the LinuxSpotPrice field.
func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPrice(v string) {
	o.LinuxSpotPrice.Set(&v)
}
// SetLinuxSpotPriceNil sets the value for LinuxSpotPrice to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetLinuxSpotPriceNil() {
	o.LinuxSpotPrice.Set(nil)
}

// UnsetLinuxSpotPrice ensures that no value is present for LinuxSpotPrice, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetLinuxSpotPrice() {
	o.LinuxSpotPrice.Unset()
}

// GetLinuxPrice returns the LinuxPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetLinuxPrice() string {
	if o == nil || IsNil(o.LinuxPrice.Get()) {
		var ret string
		return ret
	}
	return *o.LinuxPrice.Get()
}

// GetLinuxPriceOk returns a tuple with the LinuxPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetLinuxPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinuxPrice.Get(), o.LinuxPrice.IsSet()
}

// HasLinuxPrice returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasLinuxPrice() bool {
	if o != nil && o.LinuxPrice.IsSet() {
		return true
	}

	return false
}

// SetLinuxPrice gets a reference to the given NullableString and assigns it to the LinuxPrice field.
func (o *BoundFlavorsForProjectsListDto) SetLinuxPrice(v string) {
	o.LinuxPrice.Set(&v)
}
// SetLinuxPriceNil sets the value for LinuxPrice to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetLinuxPriceNil() {
	o.LinuxPrice.Set(nil)
}

// UnsetLinuxPrice ensures that no value is present for LinuxPrice, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetLinuxPrice() {
	o.LinuxPrice.Unset()
}

// GetWindowsSpotPrice returns the WindowsSpotPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPrice() string {
	if o == nil || IsNil(o.WindowsSpotPrice.Get()) {
		var ret string
		return ret
	}
	return *o.WindowsSpotPrice.Get()
}

// GetWindowsSpotPriceOk returns a tuple with the WindowsSpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetWindowsSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsSpotPrice.Get(), o.WindowsSpotPrice.IsSet()
}

// HasWindowsSpotPrice returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasWindowsSpotPrice() bool {
	if o != nil && o.WindowsSpotPrice.IsSet() {
		return true
	}

	return false
}

// SetWindowsSpotPrice gets a reference to the given NullableString and assigns it to the WindowsSpotPrice field.
func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPrice(v string) {
	o.WindowsSpotPrice.Set(&v)
}
// SetWindowsSpotPriceNil sets the value for WindowsSpotPrice to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetWindowsSpotPriceNil() {
	o.WindowsSpotPrice.Set(nil)
}

// UnsetWindowsSpotPrice ensures that no value is present for WindowsSpotPrice, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetWindowsSpotPrice() {
	o.WindowsSpotPrice.Unset()
}

// GetWindowsPrice returns the WindowsPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoundFlavorsForProjectsListDto) GetWindowsPrice() string {
	if o == nil || IsNil(o.WindowsPrice.Get()) {
		var ret string
		return ret
	}
	return *o.WindowsPrice.Get()
}

// GetWindowsPriceOk returns a tuple with the WindowsPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoundFlavorsForProjectsListDto) GetWindowsPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindowsPrice.Get(), o.WindowsPrice.IsSet()
}

// HasWindowsPrice returns a boolean if a field has been set.
func (o *BoundFlavorsForProjectsListDto) HasWindowsPrice() bool {
	if o != nil && o.WindowsPrice.IsSet() {
		return true
	}

	return false
}

// SetWindowsPrice gets a reference to the given NullableString and assigns it to the WindowsPrice field.
func (o *BoundFlavorsForProjectsListDto) SetWindowsPrice(v string) {
	o.WindowsPrice.Set(&v)
}
// SetWindowsPriceNil sets the value for WindowsPrice to be an explicit nil
func (o *BoundFlavorsForProjectsListDto) SetWindowsPriceNil() {
	o.WindowsPrice.Set(nil)
}

// UnsetWindowsPrice ensures that no value is present for WindowsPrice, not even an explicit nil
func (o *BoundFlavorsForProjectsListDto) UnsetWindowsPrice() {
	o.WindowsPrice.Unset()
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
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if !IsNil(o.IsAzure) {
		toSerialize["isAzure"] = o.IsAzure
	}
	if !IsNil(o.IsAws) {
		toSerialize["isAws"] = o.IsAws
	}
	if !IsNil(o.IsOpenstack) {
		toSerialize["isOpenstack"] = o.IsOpenstack
	}
	if !IsNil(o.IsGoogle) {
		toSerialize["isGoogle"] = o.IsGoogle
	}
	if !IsNil(o.IsProxmox) {
		toSerialize["isProxmox"] = o.IsProxmox
	}
	if !IsNil(o.IsTanzu) {
		toSerialize["isTanzu"] = o.IsTanzu
	}
	if !IsNil(o.IsOpenshift) {
		toSerialize["isOpenshift"] = o.IsOpenshift
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if o.MaxDataDiskCount.IsSet() {
		toSerialize["maxDataDiskCount"] = o.MaxDataDiskCount.Get()
	}
	if !IsNil(o.ExistsLinuxSpotPrice) {
		toSerialize["existsLinuxSpotPrice"] = o.ExistsLinuxSpotPrice
	}
	if !IsNil(o.ExistsWindowsSpotPrice) {
		toSerialize["existsWindowsSpotPrice"] = o.ExistsWindowsSpotPrice
	}
	if o.LinuxSpotPrice.IsSet() {
		toSerialize["linuxSpotPrice"] = o.LinuxSpotPrice.Get()
	}
	if o.LinuxPrice.IsSet() {
		toSerialize["linuxPrice"] = o.LinuxPrice.Get()
	}
	if o.WindowsSpotPrice.IsSet() {
		toSerialize["windowsSpotPrice"] = o.WindowsSpotPrice.Get()
	}
	if o.WindowsPrice.IsSet() {
		toSerialize["windowsPrice"] = o.WindowsPrice.Get()
	}
	return toSerialize, nil
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


