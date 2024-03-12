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

// checks if the CreateVsphereCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVsphereCommand{}

// CreateVsphereCommand struct for CreateVsphereCommand
type CreateVsphereCommand struct {
	Name NullableString `json:"name,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Password NullableString `json:"password,omitempty"`
	DatacenterId NullableString `json:"datacenterId,omitempty"`
	DatacenterName NullableString `json:"datacenterName,omitempty"`
	DatastoreName NullableString `json:"datastoreName,omitempty"`
	ResourcePoolName NullableString `json:"resourcePoolName,omitempty"`
	DrsEnabled *bool `json:"drsEnabled,omitempty"`
	VmTemplateName NullableString `json:"vmTemplateName,omitempty"`
	Continent NullableString `json:"continent,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	Hypervisors []string `json:"hypervisors,omitempty"`
	PublicNetwork *CreateVsphereNetworkDto `json:"publicNetwork,omitempty"`
	PrivateNetwork *CreateVsphereNetworkDto `json:"privateNetwork,omitempty"`
}

// NewCreateVsphereCommand instantiates a new CreateVsphereCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVsphereCommand() *CreateVsphereCommand {
	this := CreateVsphereCommand{}
	return &this
}

// NewCreateVsphereCommandWithDefaults instantiates a new CreateVsphereCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVsphereCommandWithDefaults() *CreateVsphereCommand {
	this := CreateVsphereCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateVsphereCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateVsphereCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateVsphereCommand) UnsetName() {
	o.Name.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *CreateVsphereCommand) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *CreateVsphereCommand) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *CreateVsphereCommand) UnsetUsername() {
	o.Username.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CreateVsphereCommand) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CreateVsphereCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CreateVsphereCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CreateVsphereCommand) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CreateVsphereCommand) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CreateVsphereCommand) UnsetPassword() {
	o.Password.Unset()
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetDatacenterId() string {
	if o == nil || IsNil(o.DatacenterId.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterId.Get()
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetDatacenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterId.Get(), o.DatacenterId.IsSet()
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasDatacenterId() bool {
	if o != nil && o.DatacenterId.IsSet() {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given NullableString and assigns it to the DatacenterId field.
func (o *CreateVsphereCommand) SetDatacenterId(v string) {
	o.DatacenterId.Set(&v)
}
// SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil
func (o *CreateVsphereCommand) SetDatacenterIdNil() {
	o.DatacenterId.Set(nil)
}

// UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
func (o *CreateVsphereCommand) UnsetDatacenterId() {
	o.DatacenterId.Unset()
}

// GetDatacenterName returns the DatacenterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetDatacenterName() string {
	if o == nil || IsNil(o.DatacenterName.Get()) {
		var ret string
		return ret
	}
	return *o.DatacenterName.Get()
}

// GetDatacenterNameOk returns a tuple with the DatacenterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetDatacenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatacenterName.Get(), o.DatacenterName.IsSet()
}

// HasDatacenterName returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasDatacenterName() bool {
	if o != nil && o.DatacenterName.IsSet() {
		return true
	}

	return false
}

// SetDatacenterName gets a reference to the given NullableString and assigns it to the DatacenterName field.
func (o *CreateVsphereCommand) SetDatacenterName(v string) {
	o.DatacenterName.Set(&v)
}
// SetDatacenterNameNil sets the value for DatacenterName to be an explicit nil
func (o *CreateVsphereCommand) SetDatacenterNameNil() {
	o.DatacenterName.Set(nil)
}

// UnsetDatacenterName ensures that no value is present for DatacenterName, not even an explicit nil
func (o *CreateVsphereCommand) UnsetDatacenterName() {
	o.DatacenterName.Unset()
}

// GetDatastoreName returns the DatastoreName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetDatastoreName() string {
	if o == nil || IsNil(o.DatastoreName.Get()) {
		var ret string
		return ret
	}
	return *o.DatastoreName.Get()
}

// GetDatastoreNameOk returns a tuple with the DatastoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetDatastoreNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatastoreName.Get(), o.DatastoreName.IsSet()
}

// HasDatastoreName returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasDatastoreName() bool {
	if o != nil && o.DatastoreName.IsSet() {
		return true
	}

	return false
}

// SetDatastoreName gets a reference to the given NullableString and assigns it to the DatastoreName field.
func (o *CreateVsphereCommand) SetDatastoreName(v string) {
	o.DatastoreName.Set(&v)
}
// SetDatastoreNameNil sets the value for DatastoreName to be an explicit nil
func (o *CreateVsphereCommand) SetDatastoreNameNil() {
	o.DatastoreName.Set(nil)
}

// UnsetDatastoreName ensures that no value is present for DatastoreName, not even an explicit nil
func (o *CreateVsphereCommand) UnsetDatastoreName() {
	o.DatastoreName.Unset()
}

// GetResourcePoolName returns the ResourcePoolName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetResourcePoolName() string {
	if o == nil || IsNil(o.ResourcePoolName.Get()) {
		var ret string
		return ret
	}
	return *o.ResourcePoolName.Get()
}

// GetResourcePoolNameOk returns a tuple with the ResourcePoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetResourcePoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourcePoolName.Get(), o.ResourcePoolName.IsSet()
}

// HasResourcePoolName returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasResourcePoolName() bool {
	if o != nil && o.ResourcePoolName.IsSet() {
		return true
	}

	return false
}

// SetResourcePoolName gets a reference to the given NullableString and assigns it to the ResourcePoolName field.
func (o *CreateVsphereCommand) SetResourcePoolName(v string) {
	o.ResourcePoolName.Set(&v)
}
// SetResourcePoolNameNil sets the value for ResourcePoolName to be an explicit nil
func (o *CreateVsphereCommand) SetResourcePoolNameNil() {
	o.ResourcePoolName.Set(nil)
}

// UnsetResourcePoolName ensures that no value is present for ResourcePoolName, not even an explicit nil
func (o *CreateVsphereCommand) UnsetResourcePoolName() {
	o.ResourcePoolName.Unset()
}

// GetDrsEnabled returns the DrsEnabled field value if set, zero value otherwise.
func (o *CreateVsphereCommand) GetDrsEnabled() bool {
	if o == nil || IsNil(o.DrsEnabled) {
		var ret bool
		return ret
	}
	return *o.DrsEnabled
}

// GetDrsEnabledOk returns a tuple with the DrsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVsphereCommand) GetDrsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DrsEnabled) {
		return nil, false
	}
	return o.DrsEnabled, true
}

// HasDrsEnabled returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasDrsEnabled() bool {
	if o != nil && !IsNil(o.DrsEnabled) {
		return true
	}

	return false
}

// SetDrsEnabled gets a reference to the given bool and assigns it to the DrsEnabled field.
func (o *CreateVsphereCommand) SetDrsEnabled(v bool) {
	o.DrsEnabled = &v
}

// GetVmTemplateName returns the VmTemplateName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetVmTemplateName() string {
	if o == nil || IsNil(o.VmTemplateName.Get()) {
		var ret string
		return ret
	}
	return *o.VmTemplateName.Get()
}

// GetVmTemplateNameOk returns a tuple with the VmTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetVmTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VmTemplateName.Get(), o.VmTemplateName.IsSet()
}

// HasVmTemplateName returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasVmTemplateName() bool {
	if o != nil && o.VmTemplateName.IsSet() {
		return true
	}

	return false
}

// SetVmTemplateName gets a reference to the given NullableString and assigns it to the VmTemplateName field.
func (o *CreateVsphereCommand) SetVmTemplateName(v string) {
	o.VmTemplateName.Set(&v)
}
// SetVmTemplateNameNil sets the value for VmTemplateName to be an explicit nil
func (o *CreateVsphereCommand) SetVmTemplateNameNil() {
	o.VmTemplateName.Set(nil)
}

// UnsetVmTemplateName ensures that no value is present for VmTemplateName, not even an explicit nil
func (o *CreateVsphereCommand) UnsetVmTemplateName() {
	o.VmTemplateName.Unset()
}

// GetContinent returns the Continent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetContinent() string {
	if o == nil || IsNil(o.Continent.Get()) {
		var ret string
		return ret
	}
	return *o.Continent.Get()
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetContinentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Continent.Get(), o.Continent.IsSet()
}

// HasContinent returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasContinent() bool {
	if o != nil && o.Continent.IsSet() {
		return true
	}

	return false
}

// SetContinent gets a reference to the given NullableString and assigns it to the Continent field.
func (o *CreateVsphereCommand) SetContinent(v string) {
	o.Continent.Set(&v)
}
// SetContinentNil sets the value for Continent to be an explicit nil
func (o *CreateVsphereCommand) SetContinentNil() {
	o.Continent.Set(nil)
}

// UnsetContinent ensures that no value is present for Continent, not even an explicit nil
func (o *CreateVsphereCommand) UnsetContinent() {
	o.Continent.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateVsphereCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateVsphereCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateVsphereCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetHypervisors returns the Hypervisors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVsphereCommand) GetHypervisors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hypervisors
}

// GetHypervisorsOk returns a tuple with the Hypervisors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVsphereCommand) GetHypervisorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hypervisors) {
		return nil, false
	}
	return o.Hypervisors, true
}

// HasHypervisors returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasHypervisors() bool {
	if o != nil && !IsNil(o.Hypervisors) {
		return true
	}

	return false
}

// SetHypervisors gets a reference to the given []string and assigns it to the Hypervisors field.
func (o *CreateVsphereCommand) SetHypervisors(v []string) {
	o.Hypervisors = v
}

// GetPublicNetwork returns the PublicNetwork field value if set, zero value otherwise.
func (o *CreateVsphereCommand) GetPublicNetwork() CreateVsphereNetworkDto {
	if o == nil || IsNil(o.PublicNetwork) {
		var ret CreateVsphereNetworkDto
		return ret
	}
	return *o.PublicNetwork
}

// GetPublicNetworkOk returns a tuple with the PublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVsphereCommand) GetPublicNetworkOk() (*CreateVsphereNetworkDto, bool) {
	if o == nil || IsNil(o.PublicNetwork) {
		return nil, false
	}
	return o.PublicNetwork, true
}

// HasPublicNetwork returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasPublicNetwork() bool {
	if o != nil && !IsNil(o.PublicNetwork) {
		return true
	}

	return false
}

// SetPublicNetwork gets a reference to the given CreateVsphereNetworkDto and assigns it to the PublicNetwork field.
func (o *CreateVsphereCommand) SetPublicNetwork(v CreateVsphereNetworkDto) {
	o.PublicNetwork = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *CreateVsphereCommand) GetPrivateNetwork() CreateVsphereNetworkDto {
	if o == nil || IsNil(o.PrivateNetwork) {
		var ret CreateVsphereNetworkDto
		return ret
	}
	return *o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVsphereCommand) GetPrivateNetworkOk() (*CreateVsphereNetworkDto, bool) {
	if o == nil || IsNil(o.PrivateNetwork) {
		return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *CreateVsphereCommand) HasPrivateNetwork() bool {
	if o != nil && !IsNil(o.PrivateNetwork) {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given CreateVsphereNetworkDto and assigns it to the PrivateNetwork field.
func (o *CreateVsphereCommand) SetPrivateNetwork(v CreateVsphereNetworkDto) {
	o.PrivateNetwork = &v
}

func (o CreateVsphereCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVsphereCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
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
	if o.DatastoreName.IsSet() {
		toSerialize["datastoreName"] = o.DatastoreName.Get()
	}
	if o.ResourcePoolName.IsSet() {
		toSerialize["resourcePoolName"] = o.ResourcePoolName.Get()
	}
	if !IsNil(o.DrsEnabled) {
		toSerialize["drsEnabled"] = o.DrsEnabled
	}
	if o.VmTemplateName.IsSet() {
		toSerialize["vmTemplateName"] = o.VmTemplateName.Get()
	}
	if o.Continent.IsSet() {
		toSerialize["continent"] = o.Continent.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.Hypervisors != nil {
		toSerialize["hypervisors"] = o.Hypervisors
	}
	if !IsNil(o.PublicNetwork) {
		toSerialize["publicNetwork"] = o.PublicNetwork
	}
	if !IsNil(o.PrivateNetwork) {
		toSerialize["privateNetwork"] = o.PrivateNetwork
	}
	return toSerialize, nil
}

type NullableCreateVsphereCommand struct {
	value *CreateVsphereCommand
	isSet bool
}

func (v NullableCreateVsphereCommand) Get() *CreateVsphereCommand {
	return v.value
}

func (v *NullableCreateVsphereCommand) Set(val *CreateVsphereCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVsphereCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVsphereCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVsphereCommand(val *CreateVsphereCommand) *NullableCreateVsphereCommand {
	return &NullableCreateVsphereCommand{value: val, isSet: true}
}

func (v NullableCreateVsphereCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVsphereCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


