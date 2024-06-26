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

// checks if the VsphereCredentialsForProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VsphereCredentialsForProjectDto{}

// VsphereCredentialsForProjectDto struct for VsphereCredentialsForProjectDto
type VsphereCredentialsForProjectDto struct {
	Username NullableString `json:"username,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Password NullableString `json:"password,omitempty"`
	DatacenterId NullableString `json:"datacenterId,omitempty"`
	DatacenterName NullableString `json:"datacenterName,omitempty"`
	Datastore NullableString `json:"datastore,omitempty"`
	ResourcePool NullableString `json:"resourcePool,omitempty"`
	DrsEnabled *bool `json:"drsEnabled,omitempty"`
	VmTemplateName NullableString `json:"vmTemplateName,omitempty"`
	VsphereNetworks []VsphereNetworkListDto `json:"vsphereNetworks,omitempty"`
	IgnoreSsl *bool `json:"ignoreSsl,omitempty"`
}

// NewVsphereCredentialsForProjectDto instantiates a new VsphereCredentialsForProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVsphereCredentialsForProjectDto() *VsphereCredentialsForProjectDto {
	this := VsphereCredentialsForProjectDto{}
	return &this
}

// NewVsphereCredentialsForProjectDtoWithDefaults instantiates a new VsphereCredentialsForProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVsphereCredentialsForProjectDtoWithDefaults() *VsphereCredentialsForProjectDto {
	this := VsphereCredentialsForProjectDto{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *VsphereCredentialsForProjectDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetUsername() {
	o.Username.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *VsphereCredentialsForProjectDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetUrl() {
	o.Url.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *VsphereCredentialsForProjectDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetPassword() {
	o.Password.Unset()
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetDatacenterId() string {
	if o == nil || IsNil(o.DatacenterId.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterId.Get()
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterId.Get(), o.DatacenterId.IsSet()
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasDatacenterId() bool {
	if o != nil && o.DatacenterId.IsSet() {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given NullableString and assigns it to the DatacenterId field.
func (o *VsphereCredentialsForProjectDto) SetDatacenterId(v string) {
	o.DatacenterId.Set(&v)
}
// SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetDatacenterIdNil() {
	o.DatacenterId.Set(nil)
}

// UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetDatacenterId() {
	o.DatacenterId.Unset()
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetDatacenterName() string {
	if o == nil || IsNil(o.DatacenterName.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterName.Get()
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterName.Get(), o.DatacenterName.IsSet()
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasDatacenterName() bool {
	if o != nil && o.DatacenterName.IsSet() {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given NullableString and assigns it to the DatacenterName field.
func (o *VsphereCredentialsForProjectDto) SetDatacenterName(v string) {
	o.DatacenterName.Set(&v)
}
// SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetDatacenterNameNil() {
	o.DatacenterName.Set(nil)
}

// UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetDatacenterName() {
	o.DatacenterName.Unset()
}

// GetDatastore returns the Datastore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetDatastore() string {
	if o == nil || IsNil(o.Datastore.Get()) {
		var ret string
		return ret
	}
	return *o.Datastore.Get()
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetDatastoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datastore.Get(), o.Datastore.IsSet()
}

// HasDatastore returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasDatastore() bool {
	if o != nil && o.Datastore.IsSet() {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given NullableString and assigns it to the Datastore field.
func (o *VsphereCredentialsForProjectDto) SetDatastore(v string) {
	o.Datastore.Set(&v)
}
// SetDatastoreNil sets the value for Datastore to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetDatastoreNil() {
	o.Datastore.Set(nil)
}

// UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetDatastore() {
	o.Datastore.Unset()
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetResourcePool() string {
	if o == nil || IsNil(o.ResourcePool.Get()) {
		var ret string
		return ret
	}
	return *o.ResourcePool.Get()
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetResourcePoolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourcePool.Get(), o.ResourcePool.IsSet()
}

// HasResourcePool returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasResourcePool() bool {
	if o != nil && o.ResourcePool.IsSet() {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given NullableString and assigns it to the ResourcePool field.
func (o *VsphereCredentialsForProjectDto) SetResourcePool(v string) {
	o.ResourcePool.Set(&v)
}
// SetResourcePoolNil sets the value for ResourcePool to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetResourcePoolNil() {
	o.ResourcePool.Set(nil)
}

// UnsetResourcePool ensures that no value is present for ResourcePool, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetResourcePool() {
	o.ResourcePool.Unset()
}

// GetDrsEnabled returns the DrsEnabled field value if set, zero value otherwise.
func (o *VsphereCredentialsForProjectDto) GetDrsEnabled() bool {
	if o == nil || IsNil(o.DrsEnabled) {
		var ret bool
		return ret
	}
	return *o.DrsEnabled
}

// GetDrsEnabledOk returns a tuple with the DrsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereCredentialsForProjectDto) GetDrsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DrsEnabled) {
		return nil, false
	}
	return o.DrsEnabled, true
}

// HasDrsEnabled returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasDrsEnabled() bool {
	if o != nil && !IsNil(o.DrsEnabled) {
		return true
	}

	return false
}

// SetDrsEnabled gets a reference to the given bool and assigns it to the DrsEnabled field.
func (o *VsphereCredentialsForProjectDto) SetDrsEnabled(v bool) {
	o.DrsEnabled = &v
}

// GetVmTemplateName returns the VmTemplateName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetVmTemplateName() string {
	if o == nil || IsNil(o.VmTemplateName.Get()) {
		var ret string
		return ret
	}
	return *o.VmTemplateName.Get()
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetVmTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmTemplateName.Get(), o.VmTemplateName.IsSet()
}

// HasVmTemplateName returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasVmTemplateName() bool {
	if o != nil && o.VmTemplateName.IsSet() {
		return true
	}

	return false
}

// SetVmTemplateName gets a reference to the given NullableString and assigns it to the VmTemplateName field.
func (o *VsphereCredentialsForProjectDto) SetVmTemplateName(v string) {
	o.VmTemplateName.Set(&v)
}
// SetVmTemplateNameNil sets the value for VmTemplateName to be an explicit nil
func (o *VsphereCredentialsForProjectDto) SetVmTemplateNameNil() {
	o.VmTemplateName.Set(nil)
}

// UnsetVmTemplateName ensures that no value is present for VmTemplateName, not even an explicit nil
func (o *VsphereCredentialsForProjectDto) UnsetVmTemplateName() {
	o.VmTemplateName.Unset()
}

// GetVsphereNetworks returns the VsphereNetworks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VsphereCredentialsForProjectDto) GetVsphereNetworks() []VsphereNetworkListDto {
	if o == nil {
		var ret []VsphereNetworkListDto
		return ret
	}
	return o.VsphereNetworks
}

// GetVsphereNetworksOk returns a tuple with the VsphereNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VsphereCredentialsForProjectDto) GetVsphereNetworksOk() ([]VsphereNetworkListDto, bool) {
	if o == nil || IsNil(o.VsphereNetworks) {
		return nil, false
	}
	return o.VsphereNetworks, true
}

// HasVsphereNetworks returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasVsphereNetworks() bool {
	if o != nil && !IsNil(o.VsphereNetworks) {
		return true
	}

	return false
}

// SetVsphereNetworks gets a reference to the given []VsphereNetworkListDto and assigns it to the VsphereNetworks field.
func (o *VsphereCredentialsForProjectDto) SetVsphereNetworks(v []VsphereNetworkListDto) {
	o.VsphereNetworks = v
}

// GetIgnoreSsl returns the IgnoreSsl field value if set, zero value otherwise.
func (o *VsphereCredentialsForProjectDto) GetIgnoreSsl() bool {
	if o == nil || IsNil(o.IgnoreSsl) {
		var ret bool
		return ret
	}
	return *o.IgnoreSsl
}

// GetIgnoreSslOk returns a tuple with the IgnoreSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VsphereCredentialsForProjectDto) GetIgnoreSslOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreSsl) {
		return nil, false
	}
	return o.IgnoreSsl, true
}

// HasIgnoreSsl returns a boolean if a field has been set.
func (o *VsphereCredentialsForProjectDto) HasIgnoreSsl() bool {
	if o != nil && !IsNil(o.IgnoreSsl) {
		return true
	}

	return false
}

// SetIgnoreSsl gets a reference to the given bool and assigns it to the IgnoreSsl field.
func (o *VsphereCredentialsForProjectDto) SetIgnoreSsl(v bool) {
	o.IgnoreSsl = &v
}

func (o VsphereCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VsphereCredentialsForProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.DatacenterId.IsSet() {
		toSerialize["datacenterId"] = o.DatacenterId.Get()
	}
	if o.DatacenterName.IsSet() {
		toSerialize["datacenterName"] = o.DatacenterName.Get()
	}
	if o.Datastore.IsSet() {
		toSerialize["datastore"] = o.Datastore.Get()
	}
	if o.ResourcePool.IsSet() {
		toSerialize["resourcePool"] = o.ResourcePool.Get()
	}
	if !IsNil(o.DrsEnabled) {
		toSerialize["drsEnabled"] = o.DrsEnabled
	}
	if o.VmTemplateName.IsSet() {
		toSerialize["vmTemplateName"] = o.VmTemplateName.Get()
	}
	if o.VsphereNetworks != nil {
		toSerialize["vsphereNetworks"] = o.VsphereNetworks
	}
	if !IsNil(o.IgnoreSsl) {
		toSerialize["ignoreSsl"] = o.IgnoreSsl
	}
	return toSerialize, nil
}

type NullableVsphereCredentialsForProjectDto struct {
	value *VsphereCredentialsForProjectDto
	isSet bool
}

func (v NullableVsphereCredentialsForProjectDto) Get() *VsphereCredentialsForProjectDto {
	return v.value
}

func (v *NullableVsphereCredentialsForProjectDto) Set(val *VsphereCredentialsForProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableVsphereCredentialsForProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableVsphereCredentialsForProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVsphereCredentialsForProjectDto(val *VsphereCredentialsForProjectDto) *NullableVsphereCredentialsForProjectDto {
	return &NullableVsphereCredentialsForProjectDto{value: val, isSet: true}
}

func (v NullableVsphereCredentialsForProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVsphereCredentialsForProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


