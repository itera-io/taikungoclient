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

// checks if the OpenStackNetworkListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenStackNetworkListQuery{}

// OpenStackNetworkListQuery struct for OpenStackNetworkListQuery
type OpenStackNetworkListQuery struct {
	OpenStackUser          NullableString `json:"openStackUser,omitempty"`
	OpenStackPassword      NullableString `json:"openStackPassword,omitempty"`
	OpenStackUrl           NullableString `json:"openStackUrl,omitempty"`
	OpenStackProjectId     NullableString `json:"openStackProjectId,omitempty"`
	OpenStackDomain        NullableString `json:"openStackDomain,omitempty"`
	OpenStackRegion        NullableString `json:"openStackRegion,omitempty"`
	ApplicationCredEnabled *bool          `json:"applicationCredEnabled,omitempty"`
	IsAdmin                *bool          `json:"isAdmin,omitempty"`
}

// NewOpenStackNetworkListQuery instantiates a new OpenStackNetworkListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenStackNetworkListQuery() *OpenStackNetworkListQuery {
	this := OpenStackNetworkListQuery{}
	return &this
}

// NewOpenStackNetworkListQueryWithDefaults instantiates a new OpenStackNetworkListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenStackNetworkListQueryWithDefaults() *OpenStackNetworkListQuery {
	this := OpenStackNetworkListQuery{}
	return &this
}

// GetOpenStackUser returns the OpenStackUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackNetworkListQuery) GetOpenStackUser() string {
	if o == nil || IsNil(o.OpenStackUser.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUser.Get()
}

// GetOpenStackUserOk returns a tuple with the OpenStackUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackNetworkListQuery) GetOpenStackUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUser.Get(), o.OpenStackUser.IsSet()
}

// HasOpenStackUser returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasOpenStackUser() bool {
	if o != nil && o.OpenStackUser.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUser gets a reference to the given NullableString and assigns it to the OpenStackUser field.
func (o *OpenStackNetworkListQuery) SetOpenStackUser(v string) {
	o.OpenStackUser.Set(&v)
}

// SetOpenStackUserNil sets the value for OpenStackUser to be an explicit nil
func (o *OpenStackNetworkListQuery) SetOpenStackUserNil() {
	o.OpenStackUser.Set(nil)
}

// UnsetOpenStackUser ensures that no value is present for OpenStackUser, not even an explicit nil
func (o *OpenStackNetworkListQuery) UnsetOpenStackUser() {
	o.OpenStackUser.Unset()
}

// GetOpenStackPassword returns the OpenStackPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackNetworkListQuery) GetOpenStackPassword() string {
	if o == nil || IsNil(o.OpenStackPassword.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackPassword.Get()
}

// GetOpenStackPasswordOk returns a tuple with the OpenStackPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackNetworkListQuery) GetOpenStackPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackPassword.Get(), o.OpenStackPassword.IsSet()
}

// HasOpenStackPassword returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasOpenStackPassword() bool {
	if o != nil && o.OpenStackPassword.IsSet() {
		return true
	}

	return false
}

// SetOpenStackPassword gets a reference to the given NullableString and assigns it to the OpenStackPassword field.
func (o *OpenStackNetworkListQuery) SetOpenStackPassword(v string) {
	o.OpenStackPassword.Set(&v)
}

// SetOpenStackPasswordNil sets the value for OpenStackPassword to be an explicit nil
func (o *OpenStackNetworkListQuery) SetOpenStackPasswordNil() {
	o.OpenStackPassword.Set(nil)
}

// UnsetOpenStackPassword ensures that no value is present for OpenStackPassword, not even an explicit nil
func (o *OpenStackNetworkListQuery) UnsetOpenStackPassword() {
	o.OpenStackPassword.Unset()
}

// GetOpenStackUrl returns the OpenStackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackNetworkListQuery) GetOpenStackUrl() string {
	if o == nil || IsNil(o.OpenStackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackUrl.Get()
}

// GetOpenStackUrlOk returns a tuple with the OpenStackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackNetworkListQuery) GetOpenStackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackUrl.Get(), o.OpenStackUrl.IsSet()
}

// HasOpenStackUrl returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasOpenStackUrl() bool {
	if o != nil && o.OpenStackUrl.IsSet() {
		return true
	}

	return false
}

// SetOpenStackUrl gets a reference to the given NullableString and assigns it to the OpenStackUrl field.
func (o *OpenStackNetworkListQuery) SetOpenStackUrl(v string) {
	o.OpenStackUrl.Set(&v)
}

// SetOpenStackUrlNil sets the value for OpenStackUrl to be an explicit nil
func (o *OpenStackNetworkListQuery) SetOpenStackUrlNil() {
	o.OpenStackUrl.Set(nil)
}

// UnsetOpenStackUrl ensures that no value is present for OpenStackUrl, not even an explicit nil
func (o *OpenStackNetworkListQuery) UnsetOpenStackUrl() {
	o.OpenStackUrl.Unset()
}

// GetOpenStackProjectId returns the OpenStackProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackNetworkListQuery) GetOpenStackProjectId() string {
	if o == nil || IsNil(o.OpenStackProjectId.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackProjectId.Get()
}

// GetOpenStackProjectIdOk returns a tuple with the OpenStackProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackNetworkListQuery) GetOpenStackProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackProjectId.Get(), o.OpenStackProjectId.IsSet()
}

// HasOpenStackProjectId returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasOpenStackProjectId() bool {
	if o != nil && o.OpenStackProjectId.IsSet() {
		return true
	}

	return false
}

// SetOpenStackProjectId gets a reference to the given NullableString and assigns it to the OpenStackProjectId field.
func (o *OpenStackNetworkListQuery) SetOpenStackProjectId(v string) {
	o.OpenStackProjectId.Set(&v)
}

// SetOpenStackProjectIdNil sets the value for OpenStackProjectId to be an explicit nil
func (o *OpenStackNetworkListQuery) SetOpenStackProjectIdNil() {
	o.OpenStackProjectId.Set(nil)
}

// UnsetOpenStackProjectId ensures that no value is present for OpenStackProjectId, not even an explicit nil
func (o *OpenStackNetworkListQuery) UnsetOpenStackProjectId() {
	o.OpenStackProjectId.Unset()
}

// GetOpenStackDomain returns the OpenStackDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackNetworkListQuery) GetOpenStackDomain() string {
	if o == nil || IsNil(o.OpenStackDomain.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackDomain.Get()
}

// GetOpenStackDomainOk returns a tuple with the OpenStackDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackNetworkListQuery) GetOpenStackDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackDomain.Get(), o.OpenStackDomain.IsSet()
}

// HasOpenStackDomain returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasOpenStackDomain() bool {
	if o != nil && o.OpenStackDomain.IsSet() {
		return true
	}

	return false
}

// SetOpenStackDomain gets a reference to the given NullableString and assigns it to the OpenStackDomain field.
func (o *OpenStackNetworkListQuery) SetOpenStackDomain(v string) {
	o.OpenStackDomain.Set(&v)
}

// SetOpenStackDomainNil sets the value for OpenStackDomain to be an explicit nil
func (o *OpenStackNetworkListQuery) SetOpenStackDomainNil() {
	o.OpenStackDomain.Set(nil)
}

// UnsetOpenStackDomain ensures that no value is present for OpenStackDomain, not even an explicit nil
func (o *OpenStackNetworkListQuery) UnsetOpenStackDomain() {
	o.OpenStackDomain.Unset()
}

// GetOpenStackRegion returns the OpenStackRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenStackNetworkListQuery) GetOpenStackRegion() string {
	if o == nil || IsNil(o.OpenStackRegion.Get()) {
		var ret string
		return ret
	}
	return *o.OpenStackRegion.Get()
}

// GetOpenStackRegionOk returns a tuple with the OpenStackRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenStackNetworkListQuery) GetOpenStackRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenStackRegion.Get(), o.OpenStackRegion.IsSet()
}

// HasOpenStackRegion returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasOpenStackRegion() bool {
	if o != nil && o.OpenStackRegion.IsSet() {
		return true
	}

	return false
}

// SetOpenStackRegion gets a reference to the given NullableString and assigns it to the OpenStackRegion field.
func (o *OpenStackNetworkListQuery) SetOpenStackRegion(v string) {
	o.OpenStackRegion.Set(&v)
}

// SetOpenStackRegionNil sets the value for OpenStackRegion to be an explicit nil
func (o *OpenStackNetworkListQuery) SetOpenStackRegionNil() {
	o.OpenStackRegion.Set(nil)
}

// UnsetOpenStackRegion ensures that no value is present for OpenStackRegion, not even an explicit nil
func (o *OpenStackNetworkListQuery) UnsetOpenStackRegion() {
	o.OpenStackRegion.Unset()
}

// GetApplicationCredEnabled returns the ApplicationCredEnabled field value if set, zero value otherwise.
func (o *OpenStackNetworkListQuery) GetApplicationCredEnabled() bool {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		var ret bool
		return ret
	}
	return *o.ApplicationCredEnabled
}

// GetApplicationCredEnabledOk returns a tuple with the ApplicationCredEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackNetworkListQuery) GetApplicationCredEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplicationCredEnabled) {
		return nil, false
	}
	return o.ApplicationCredEnabled, true
}

// HasApplicationCredEnabled returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasApplicationCredEnabled() bool {
	if o != nil && !IsNil(o.ApplicationCredEnabled) {
		return true
	}

	return false
}

// SetApplicationCredEnabled gets a reference to the given bool and assigns it to the ApplicationCredEnabled field.
func (o *OpenStackNetworkListQuery) SetApplicationCredEnabled(v bool) {
	o.ApplicationCredEnabled = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *OpenStackNetworkListQuery) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenStackNetworkListQuery) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *OpenStackNetworkListQuery) HasIsAdmin() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *OpenStackNetworkListQuery) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

func (o OpenStackNetworkListQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenStackNetworkListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenStackUser.IsSet() {
		toSerialize["openStackUser"] = o.OpenStackUser.Get()
	}
	if o.OpenStackPassword.IsSet() {
		toSerialize["openStackPassword"] = o.OpenStackPassword.Get()
	}
	if o.OpenStackUrl.IsSet() {
		toSerialize["openStackUrl"] = o.OpenStackUrl.Get()
	}
	if o.OpenStackProjectId.IsSet() {
		toSerialize["openStackProjectId"] = o.OpenStackProjectId.Get()
	}
	if o.OpenStackDomain.IsSet() {
		toSerialize["openStackDomain"] = o.OpenStackDomain.Get()
	}
	if o.OpenStackRegion.IsSet() {
		toSerialize["openStackRegion"] = o.OpenStackRegion.Get()
	}
	if !IsNil(o.ApplicationCredEnabled) {
		toSerialize["applicationCredEnabled"] = o.ApplicationCredEnabled
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	return toSerialize, nil
}

type NullableOpenStackNetworkListQuery struct {
	value *OpenStackNetworkListQuery
	isSet bool
}

func (v NullableOpenStackNetworkListQuery) Get() *OpenStackNetworkListQuery {
	return v.value
}

func (v *NullableOpenStackNetworkListQuery) Set(val *OpenStackNetworkListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenStackNetworkListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenStackNetworkListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenStackNetworkListQuery(val *OpenStackNetworkListQuery) *NullableOpenStackNetworkListQuery {
	return &NullableOpenStackNetworkListQuery{value: val, isSet: true}
}

func (v NullableOpenStackNetworkListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenStackNetworkListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
