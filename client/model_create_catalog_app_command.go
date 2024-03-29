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

// checks if the CreateCatalogAppCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCatalogAppCommand{}

// CreateCatalogAppCommand struct for CreateCatalogAppCommand
type CreateCatalogAppCommand struct {
	RepoName NullableString `json:"repoName,omitempty"`
	PackageName NullableString `json:"packageName,omitempty"`
	CatalogId *int32 `json:"catalogId,omitempty"`
	Version NullableString `json:"version,omitempty"`
	Parameters []CatalogAppParamsDto `json:"parameters,omitempty"`
}

// NewCreateCatalogAppCommand instantiates a new CreateCatalogAppCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCatalogAppCommand() *CreateCatalogAppCommand {
	this := CreateCatalogAppCommand{}
	return &this
}

// NewCreateCatalogAppCommandWithDefaults instantiates a new CreateCatalogAppCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCatalogAppCommandWithDefaults() *CreateCatalogAppCommand {
	this := CreateCatalogAppCommand{}
	return &this
}

// GetRepoName returns the RepoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCatalogAppCommand) GetRepoName() string {
	if o == nil || IsNil(o.RepoName.Get()) {
		var ret string
		return ret
	}
	return *o.RepoName.Get()
}

// GetRepoNameOk returns a tuple with the RepoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCatalogAppCommand) GetRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepoName.Get(), o.RepoName.IsSet()
}

// HasRepoName returns a boolean if a field has been set.
func (o *CreateCatalogAppCommand) HasRepoName() bool {
	if o != nil && o.RepoName.IsSet() {
		return true
	}

	return false
}

// SetRepoName gets a reference to the given NullableString and assigns it to the RepoName field.
func (o *CreateCatalogAppCommand) SetRepoName(v string) {
	o.RepoName.Set(&v)
}
// SetRepoNameNil sets the value for RepoName to be an explicit nil
func (o *CreateCatalogAppCommand) SetRepoNameNil() {
	o.RepoName.Set(nil)
}

// UnsetRepoName ensures that no value is present for RepoName, not even an explicit nil
func (o *CreateCatalogAppCommand) UnsetRepoName() {
	o.RepoName.Unset()
}

// GetPackageName returns the PackageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCatalogAppCommand) GetPackageName() string {
	if o == nil || IsNil(o.PackageName.Get()) {
		var ret string
		return ret
	}
	return *o.PackageName.Get()
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCatalogAppCommand) GetPackageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageName.Get(), o.PackageName.IsSet()
}

// HasPackageName returns a boolean if a field has been set.
func (o *CreateCatalogAppCommand) HasPackageName() bool {
	if o != nil && o.PackageName.IsSet() {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given NullableString and assigns it to the PackageName field.
func (o *CreateCatalogAppCommand) SetPackageName(v string) {
	o.PackageName.Set(&v)
}
// SetPackageNameNil sets the value for PackageName to be an explicit nil
func (o *CreateCatalogAppCommand) SetPackageNameNil() {
	o.PackageName.Set(nil)
}

// UnsetPackageName ensures that no value is present for PackageName, not even an explicit nil
func (o *CreateCatalogAppCommand) UnsetPackageName() {
	o.PackageName.Unset()
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *CreateCatalogAppCommand) GetCatalogId() int32 {
	if o == nil || IsNil(o.CatalogId) {
		var ret int32
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCatalogAppCommand) GetCatalogIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *CreateCatalogAppCommand) HasCatalogId() bool {
	if o != nil && !IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given int32 and assigns it to the CatalogId field.
func (o *CreateCatalogAppCommand) SetCatalogId(v int32) {
	o.CatalogId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCatalogAppCommand) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCatalogAppCommand) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateCatalogAppCommand) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *CreateCatalogAppCommand) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *CreateCatalogAppCommand) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *CreateCatalogAppCommand) UnsetVersion() {
	o.Version.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateCatalogAppCommand) GetParameters() []CatalogAppParamsDto {
	if o == nil {
		var ret []CatalogAppParamsDto
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCatalogAppCommand) GetParametersOk() ([]CatalogAppParamsDto, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateCatalogAppCommand) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []CatalogAppParamsDto and assigns it to the Parameters field.
func (o *CreateCatalogAppCommand) SetParameters(v []CatalogAppParamsDto) {
	o.Parameters = v
}

func (o CreateCatalogAppCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCatalogAppCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RepoName.IsSet() {
		toSerialize["repoName"] = o.RepoName.Get()
	}
	if o.PackageName.IsSet() {
		toSerialize["packageName"] = o.PackageName.Get()
	}
	if !IsNil(o.CatalogId) {
		toSerialize["catalogId"] = o.CatalogId
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableCreateCatalogAppCommand struct {
	value *CreateCatalogAppCommand
	isSet bool
}

func (v NullableCreateCatalogAppCommand) Get() *CreateCatalogAppCommand {
	return v.value
}

func (v *NullableCreateCatalogAppCommand) Set(val *CreateCatalogAppCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCatalogAppCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCatalogAppCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCatalogAppCommand(val *CreateCatalogAppCommand) *NullableCreateCatalogAppCommand {
	return &NullableCreateCatalogAppCommand{value: val, isSet: true}
}

func (v NullableCreateCatalogAppCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCatalogAppCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


