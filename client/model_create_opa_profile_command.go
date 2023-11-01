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

// checks if the CreateOpaProfileCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOpaProfileCommand{}

// CreateOpaProfileCommand struct for CreateOpaProfileCommand
type CreateOpaProfileCommand struct {
	Name                  NullableString `json:"name,omitempty"`
	ForbidNodePort        *bool          `json:"forbidNodePort,omitempty"`
	ForbidHttpIngress     *bool          `json:"forbidHttpIngress,omitempty"`
	RequireProbe          *bool          `json:"requireProbe,omitempty"`
	UniqueIngresses       *bool          `json:"uniqueIngresses,omitempty"`
	UniqueServiceSelector *bool          `json:"uniqueServiceSelector,omitempty"`
	ForcePodResource      *bool          `json:"forcePodResource,omitempty"`
	AllowedRepo           []string       `json:"allowedRepo,omitempty"`
	ForbidSpecificTags    []string       `json:"forbidSpecificTags,omitempty"`
	IngressWhitelist      []string       `json:"ingressWhitelist,omitempty"`
	OrganizationId        NullableInt32  `json:"organizationId,omitempty"`
}

// NewCreateOpaProfileCommand instantiates a new CreateOpaProfileCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOpaProfileCommand() *CreateOpaProfileCommand {
	this := CreateOpaProfileCommand{}
	return &this
}

// NewCreateOpaProfileCommandWithDefaults instantiates a new CreateOpaProfileCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOpaProfileCommandWithDefaults() *CreateOpaProfileCommand {
	this := CreateOpaProfileCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpaProfileCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpaProfileCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateOpaProfileCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateOpaProfileCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateOpaProfileCommand) UnsetName() {
	o.Name.Unset()
}

// GetForbidNodePort returns the ForbidNodePort field value if set, zero value otherwise.
func (o *CreateOpaProfileCommand) GetForbidNodePort() bool {
	if o == nil || IsNil(o.ForbidNodePort) {
		var ret bool
		return ret
	}
	return *o.ForbidNodePort
}

// GetForbidNodePortOk returns a tuple with the ForbidNodePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpaProfileCommand) GetForbidNodePortOk() (*bool, bool) {
	if o == nil || IsNil(o.ForbidNodePort) {
		return nil, false
	}
	return o.ForbidNodePort, true
}

// HasForbidNodePort returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasForbidNodePort() bool {
	if o != nil && !IsNil(o.ForbidNodePort) {
		return true
	}

	return false
}

// SetForbidNodePort gets a reference to the given bool and assigns it to the ForbidNodePort field.
func (o *CreateOpaProfileCommand) SetForbidNodePort(v bool) {
	o.ForbidNodePort = &v
}

// GetForbidHttpIngress returns the ForbidHttpIngress field value if set, zero value otherwise.
func (o *CreateOpaProfileCommand) GetForbidHttpIngress() bool {
	if o == nil || IsNil(o.ForbidHttpIngress) {
		var ret bool
		return ret
	}
	return *o.ForbidHttpIngress
}

// GetForbidHttpIngressOk returns a tuple with the ForbidHttpIngress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpaProfileCommand) GetForbidHttpIngressOk() (*bool, bool) {
	if o == nil || IsNil(o.ForbidHttpIngress) {
		return nil, false
	}
	return o.ForbidHttpIngress, true
}

// HasForbidHttpIngress returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasForbidHttpIngress() bool {
	if o != nil && !IsNil(o.ForbidHttpIngress) {
		return true
	}

	return false
}

// SetForbidHttpIngress gets a reference to the given bool and assigns it to the ForbidHttpIngress field.
func (o *CreateOpaProfileCommand) SetForbidHttpIngress(v bool) {
	o.ForbidHttpIngress = &v
}

// GetRequireProbe returns the RequireProbe field value if set, zero value otherwise.
func (o *CreateOpaProfileCommand) GetRequireProbe() bool {
	if o == nil || IsNil(o.RequireProbe) {
		var ret bool
		return ret
	}
	return *o.RequireProbe
}

// GetRequireProbeOk returns a tuple with the RequireProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpaProfileCommand) GetRequireProbeOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireProbe) {
		return nil, false
	}
	return o.RequireProbe, true
}

// HasRequireProbe returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasRequireProbe() bool {
	if o != nil && !IsNil(o.RequireProbe) {
		return true
	}

	return false
}

// SetRequireProbe gets a reference to the given bool and assigns it to the RequireProbe field.
func (o *CreateOpaProfileCommand) SetRequireProbe(v bool) {
	o.RequireProbe = &v
}

// GetUniqueIngresses returns the UniqueIngresses field value if set, zero value otherwise.
func (o *CreateOpaProfileCommand) GetUniqueIngresses() bool {
	if o == nil || IsNil(o.UniqueIngresses) {
		var ret bool
		return ret
	}
	return *o.UniqueIngresses
}

// GetUniqueIngressesOk returns a tuple with the UniqueIngresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpaProfileCommand) GetUniqueIngressesOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueIngresses) {
		return nil, false
	}
	return o.UniqueIngresses, true
}

// HasUniqueIngresses returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasUniqueIngresses() bool {
	if o != nil && !IsNil(o.UniqueIngresses) {
		return true
	}

	return false
}

// SetUniqueIngresses gets a reference to the given bool and assigns it to the UniqueIngresses field.
func (o *CreateOpaProfileCommand) SetUniqueIngresses(v bool) {
	o.UniqueIngresses = &v
}

// GetUniqueServiceSelector returns the UniqueServiceSelector field value if set, zero value otherwise.
func (o *CreateOpaProfileCommand) GetUniqueServiceSelector() bool {
	if o == nil || IsNil(o.UniqueServiceSelector) {
		var ret bool
		return ret
	}
	return *o.UniqueServiceSelector
}

// GetUniqueServiceSelectorOk returns a tuple with the UniqueServiceSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpaProfileCommand) GetUniqueServiceSelectorOk() (*bool, bool) {
	if o == nil || IsNil(o.UniqueServiceSelector) {
		return nil, false
	}
	return o.UniqueServiceSelector, true
}

// HasUniqueServiceSelector returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasUniqueServiceSelector() bool {
	if o != nil && !IsNil(o.UniqueServiceSelector) {
		return true
	}

	return false
}

// SetUniqueServiceSelector gets a reference to the given bool and assigns it to the UniqueServiceSelector field.
func (o *CreateOpaProfileCommand) SetUniqueServiceSelector(v bool) {
	o.UniqueServiceSelector = &v
}

// GetForcePodResource returns the ForcePodResource field value if set, zero value otherwise.
func (o *CreateOpaProfileCommand) GetForcePodResource() bool {
	if o == nil || IsNil(o.ForcePodResource) {
		var ret bool
		return ret
	}
	return *o.ForcePodResource
}

// GetForcePodResourceOk returns a tuple with the ForcePodResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOpaProfileCommand) GetForcePodResourceOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePodResource) {
		return nil, false
	}
	return o.ForcePodResource, true
}

// HasForcePodResource returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasForcePodResource() bool {
	if o != nil && !IsNil(o.ForcePodResource) {
		return true
	}

	return false
}

// SetForcePodResource gets a reference to the given bool and assigns it to the ForcePodResource field.
func (o *CreateOpaProfileCommand) SetForcePodResource(v bool) {
	o.ForcePodResource = &v
}

// GetAllowedRepo returns the AllowedRepo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpaProfileCommand) GetAllowedRepo() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedRepo
}

// GetAllowedRepoOk returns a tuple with the AllowedRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpaProfileCommand) GetAllowedRepoOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedRepo) {
		return nil, false
	}
	return o.AllowedRepo, true
}

// HasAllowedRepo returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasAllowedRepo() bool {
	if o != nil && IsNil(o.AllowedRepo) {
		return true
	}

	return false
}

// SetAllowedRepo gets a reference to the given []string and assigns it to the AllowedRepo field.
func (o *CreateOpaProfileCommand) SetAllowedRepo(v []string) {
	o.AllowedRepo = v
}

// GetForbidSpecificTags returns the ForbidSpecificTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpaProfileCommand) GetForbidSpecificTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ForbidSpecificTags
}

// GetForbidSpecificTagsOk returns a tuple with the ForbidSpecificTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpaProfileCommand) GetForbidSpecificTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.ForbidSpecificTags) {
		return nil, false
	}
	return o.ForbidSpecificTags, true
}

// HasForbidSpecificTags returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasForbidSpecificTags() bool {
	if o != nil && IsNil(o.ForbidSpecificTags) {
		return true
	}

	return false
}

// SetForbidSpecificTags gets a reference to the given []string and assigns it to the ForbidSpecificTags field.
func (o *CreateOpaProfileCommand) SetForbidSpecificTags(v []string) {
	o.ForbidSpecificTags = v
}

// GetIngressWhitelist returns the IngressWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpaProfileCommand) GetIngressWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IngressWhitelist
}

// GetIngressWhitelistOk returns a tuple with the IngressWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpaProfileCommand) GetIngressWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.IngressWhitelist) {
		return nil, false
	}
	return o.IngressWhitelist, true
}

// HasIngressWhitelist returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasIngressWhitelist() bool {
	if o != nil && IsNil(o.IngressWhitelist) {
		return true
	}

	return false
}

// SetIngressWhitelist gets a reference to the given []string and assigns it to the IngressWhitelist field.
func (o *CreateOpaProfileCommand) SetIngressWhitelist(v []string) {
	o.IngressWhitelist = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOpaProfileCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOpaProfileCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateOpaProfileCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateOpaProfileCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateOpaProfileCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateOpaProfileCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

func (o CreateOpaProfileCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOpaProfileCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.ForbidNodePort) {
		toSerialize["forbidNodePort"] = o.ForbidNodePort
	}
	if !IsNil(o.ForbidHttpIngress) {
		toSerialize["forbidHttpIngress"] = o.ForbidHttpIngress
	}
	if !IsNil(o.RequireProbe) {
		toSerialize["requireProbe"] = o.RequireProbe
	}
	if !IsNil(o.UniqueIngresses) {
		toSerialize["uniqueIngresses"] = o.UniqueIngresses
	}
	if !IsNil(o.UniqueServiceSelector) {
		toSerialize["uniqueServiceSelector"] = o.UniqueServiceSelector
	}
	if !IsNil(o.ForcePodResource) {
		toSerialize["forcePodResource"] = o.ForcePodResource
	}
	if o.AllowedRepo != nil {
		toSerialize["allowedRepo"] = o.AllowedRepo
	}
	if o.ForbidSpecificTags != nil {
		toSerialize["forbidSpecificTags"] = o.ForbidSpecificTags
	}
	if o.IngressWhitelist != nil {
		toSerialize["ingressWhitelist"] = o.IngressWhitelist
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	return toSerialize, nil
}

type NullableCreateOpaProfileCommand struct {
	value *CreateOpaProfileCommand
	isSet bool
}

func (v NullableCreateOpaProfileCommand) Get() *CreateOpaProfileCommand {
	return v.value
}

func (v *NullableCreateOpaProfileCommand) Set(val *CreateOpaProfileCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOpaProfileCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOpaProfileCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOpaProfileCommand(val *CreateOpaProfileCommand) *NullableCreateOpaProfileCommand {
	return &NullableCreateOpaProfileCommand{value: val, isSet: true}
}

func (v NullableCreateOpaProfileCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOpaProfileCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}