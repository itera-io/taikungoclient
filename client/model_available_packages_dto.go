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

// checks if the AvailablePackagesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailablePackagesDto{}

// AvailablePackagesDto struct for AvailablePackagesDto
type AvailablePackagesDto struct {
	PackageId NullableString `json:"packageId,omitempty"`
	CatalogAppId NullableInt32 `json:"catalogAppId,omitempty"`
	InstalledInstanceCount NullableInt32 `json:"installedInstanceCount,omitempty"`
	Name NullableString `json:"name,omitempty"`
	NormalizedName NullableString `json:"normalizedName,omitempty"`
	LogoImageId NullableString `json:"logoImageId,omitempty"`
	Stars *int64 `json:"stars,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Version NullableString `json:"version,omitempty"`
	AppVersion NullableString `json:"appVersion,omitempty"`
	Deprecated *bool `json:"deprecated,omitempty"`
	Signed *bool `json:"signed,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	SecurityReportSummary *SecurityReportSummary `json:"securityReportSummary,omitempty"`
	Ts NullableString `json:"ts,omitempty"`
	Repository *Repository `json:"repository,omitempty"`
}

// NewAvailablePackagesDto instantiates a new AvailablePackagesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailablePackagesDto() *AvailablePackagesDto {
	this := AvailablePackagesDto{}
	return &this
}

// NewAvailablePackagesDtoWithDefaults instantiates a new AvailablePackagesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailablePackagesDtoWithDefaults() *AvailablePackagesDto {
	this := AvailablePackagesDto{}
	return &this
}

// GetPackageId returns the PackageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetPackageId() string {
	if o == nil || IsNil(o.PackageId.Get()) {
		var ret string
		return ret
	}
	return *o.PackageId.Get()
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageId.Get(), o.PackageId.IsSet()
}

// HasPackageId returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasPackageId() bool {
	if o != nil && o.PackageId.IsSet() {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given NullableString and assigns it to the PackageId field.
func (o *AvailablePackagesDto) SetPackageId(v string) {
	o.PackageId.Set(&v)
}
// SetPackageIdNil sets the value for PackageId to be an explicit nil
func (o *AvailablePackagesDto) SetPackageIdNil() {
	o.PackageId.Set(nil)
}

// UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
func (o *AvailablePackagesDto) UnsetPackageId() {
	o.PackageId.Unset()
}

// GetCatalogAppId returns the CatalogAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetCatalogAppId() int32 {
	if o == nil || IsNil(o.CatalogAppId.Get()) {
		var ret int32
		return ret
	}
	return *o.CatalogAppId.Get()
}

// GetCatalogAppIdOk returns a tuple with the CatalogAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetCatalogAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogAppId.Get(), o.CatalogAppId.IsSet()
}

// HasCatalogAppId returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasCatalogAppId() bool {
	if o != nil && o.CatalogAppId.IsSet() {
		return true
	}

	return false
}

// SetCatalogAppId gets a reference to the given NullableInt32 and assigns it to the CatalogAppId field.
func (o *AvailablePackagesDto) SetCatalogAppId(v int32) {
	o.CatalogAppId.Set(&v)
}
// SetCatalogAppIdNil sets the value for CatalogAppId to be an explicit nil
func (o *AvailablePackagesDto) SetCatalogAppIdNil() {
	o.CatalogAppId.Set(nil)
}

// UnsetCatalogAppId ensures that no value is present for CatalogAppId, not even an explicit nil
func (o *AvailablePackagesDto) UnsetCatalogAppId() {
	o.CatalogAppId.Unset()
}

// GetInstalledInstanceCount returns the InstalledInstanceCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetInstalledInstanceCount() int32 {
	if o == nil || IsNil(o.InstalledInstanceCount.Get()) {
		var ret int32
		return ret
	}
	return *o.InstalledInstanceCount.Get()
}

// GetInstalledInstanceCountOk returns a tuple with the InstalledInstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetInstalledInstanceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledInstanceCount.Get(), o.InstalledInstanceCount.IsSet()
}

// HasInstalledInstanceCount returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasInstalledInstanceCount() bool {
	if o != nil && o.InstalledInstanceCount.IsSet() {
		return true
	}

	return false
}

// SetInstalledInstanceCount gets a reference to the given NullableInt32 and assigns it to the InstalledInstanceCount field.
func (o *AvailablePackagesDto) SetInstalledInstanceCount(v int32) {
	o.InstalledInstanceCount.Set(&v)
}
// SetInstalledInstanceCountNil sets the value for InstalledInstanceCount to be an explicit nil
func (o *AvailablePackagesDto) SetInstalledInstanceCountNil() {
	o.InstalledInstanceCount.Set(nil)
}

// UnsetInstalledInstanceCount ensures that no value is present for InstalledInstanceCount, not even an explicit nil
func (o *AvailablePackagesDto) UnsetInstalledInstanceCount() {
	o.InstalledInstanceCount.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AvailablePackagesDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AvailablePackagesDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AvailablePackagesDto) UnsetName() {
	o.Name.Unset()
}

// GetNormalizedName returns the NormalizedName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetNormalizedName() string {
	if o == nil || IsNil(o.NormalizedName.Get()) {
		var ret string
		return ret
	}
	return *o.NormalizedName.Get()
}

// GetNormalizedNameOk returns a tuple with the NormalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetNormalizedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NormalizedName.Get(), o.NormalizedName.IsSet()
}

// HasNormalizedName returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasNormalizedName() bool {
	if o != nil && o.NormalizedName.IsSet() {
		return true
	}

	return false
}

// SetNormalizedName gets a reference to the given NullableString and assigns it to the NormalizedName field.
func (o *AvailablePackagesDto) SetNormalizedName(v string) {
	o.NormalizedName.Set(&v)
}
// SetNormalizedNameNil sets the value for NormalizedName to be an explicit nil
func (o *AvailablePackagesDto) SetNormalizedNameNil() {
	o.NormalizedName.Set(nil)
}

// UnsetNormalizedName ensures that no value is present for NormalizedName, not even an explicit nil
func (o *AvailablePackagesDto) UnsetNormalizedName() {
	o.NormalizedName.Unset()
}

// GetLogoImageId returns the LogoImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetLogoImageId() string {
	if o == nil || IsNil(o.LogoImageId.Get()) {
		var ret string
		return ret
	}
	return *o.LogoImageId.Get()
}

// GetLogoImageIdOk returns a tuple with the LogoImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetLogoImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoImageId.Get(), o.LogoImageId.IsSet()
}

// HasLogoImageId returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasLogoImageId() bool {
	if o != nil && o.LogoImageId.IsSet() {
		return true
	}

	return false
}

// SetLogoImageId gets a reference to the given NullableString and assigns it to the LogoImageId field.
func (o *AvailablePackagesDto) SetLogoImageId(v string) {
	o.LogoImageId.Set(&v)
}
// SetLogoImageIdNil sets the value for LogoImageId to be an explicit nil
func (o *AvailablePackagesDto) SetLogoImageIdNil() {
	o.LogoImageId.Set(nil)
}

// UnsetLogoImageId ensures that no value is present for LogoImageId, not even an explicit nil
func (o *AvailablePackagesDto) UnsetLogoImageId() {
	o.LogoImageId.Unset()
}

// GetStars returns the Stars field value if set, zero value otherwise.
func (o *AvailablePackagesDto) GetStars() int64 {
	if o == nil || IsNil(o.Stars) {
		var ret int64
		return ret
	}
	return *o.Stars
}

// GetStarsOk returns a tuple with the Stars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesDto) GetStarsOk() (*int64, bool) {
	if o == nil || IsNil(o.Stars) {
		return nil, false
	}
	return o.Stars, true
}

// HasStars returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasStars() bool {
	if o != nil && !IsNil(o.Stars) {
		return true
	}

	return false
}

// SetStars gets a reference to the given int64 and assigns it to the Stars field.
func (o *AvailablePackagesDto) SetStars(v int64) {
	o.Stars = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AvailablePackagesDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AvailablePackagesDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AvailablePackagesDto) UnsetDescription() {
	o.Description.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *AvailablePackagesDto) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *AvailablePackagesDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *AvailablePackagesDto) UnsetVersion() {
	o.Version.Unset()
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasAppVersion() bool {
	if o != nil && o.AppVersion.IsSet() {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given NullableString and assigns it to the AppVersion field.
func (o *AvailablePackagesDto) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}
// SetAppVersionNil sets the value for AppVersion to be an explicit nil
func (o *AvailablePackagesDto) SetAppVersionNil() {
	o.AppVersion.Set(nil)
}

// UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
func (o *AvailablePackagesDto) UnsetAppVersion() {
	o.AppVersion.Unset()
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *AvailablePackagesDto) GetDeprecated() bool {
	if o == nil || IsNil(o.Deprecated) {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesDto) GetDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deprecated) {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasDeprecated() bool {
	if o != nil && !IsNil(o.Deprecated) {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *AvailablePackagesDto) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetSigned returns the Signed field value if set, zero value otherwise.
func (o *AvailablePackagesDto) GetSigned() bool {
	if o == nil || IsNil(o.Signed) {
		var ret bool
		return ret
	}
	return *o.Signed
}

// GetSignedOk returns a tuple with the Signed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesDto) GetSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.Signed) {
		return nil, false
	}
	return o.Signed, true
}

// HasSigned returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasSigned() bool {
	if o != nil && !IsNil(o.Signed) {
		return true
	}

	return false
}

// SetSigned gets a reference to the given bool and assigns it to the Signed field.
func (o *AvailablePackagesDto) SetSigned(v bool) {
	o.Signed = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *AvailablePackagesDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *AvailablePackagesDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetSecurityReportSummary returns the SecurityReportSummary field value if set, zero value otherwise.
func (o *AvailablePackagesDto) GetSecurityReportSummary() SecurityReportSummary {
	if o == nil || IsNil(o.SecurityReportSummary) {
		var ret SecurityReportSummary
		return ret
	}
	return *o.SecurityReportSummary
}

// GetSecurityReportSummaryOk returns a tuple with the SecurityReportSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesDto) GetSecurityReportSummaryOk() (*SecurityReportSummary, bool) {
	if o == nil || IsNil(o.SecurityReportSummary) {
		return nil, false
	}
	return o.SecurityReportSummary, true
}

// HasSecurityReportSummary returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasSecurityReportSummary() bool {
	if o != nil && !IsNil(o.SecurityReportSummary) {
		return true
	}

	return false
}

// SetSecurityReportSummary gets a reference to the given SecurityReportSummary and assigns it to the SecurityReportSummary field.
func (o *AvailablePackagesDto) SetSecurityReportSummary(v SecurityReportSummary) {
	o.SecurityReportSummary = &v
}

// GetTs returns the Ts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AvailablePackagesDto) GetTs() string {
	if o == nil || IsNil(o.Ts.Get()) {
		var ret string
		return ret
	}
	return *o.Ts.Get()
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AvailablePackagesDto) GetTsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ts.Get(), o.Ts.IsSet()
}

// HasTs returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasTs() bool {
	if o != nil && o.Ts.IsSet() {
		return true
	}

	return false
}

// SetTs gets a reference to the given NullableString and assigns it to the Ts field.
func (o *AvailablePackagesDto) SetTs(v string) {
	o.Ts.Set(&v)
}
// SetTsNil sets the value for Ts to be an explicit nil
func (o *AvailablePackagesDto) SetTsNil() {
	o.Ts.Set(nil)
}

// UnsetTs ensures that no value is present for Ts, not even an explicit nil
func (o *AvailablePackagesDto) UnsetTs() {
	o.Ts.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *AvailablePackagesDto) GetRepository() Repository {
	if o == nil || IsNil(o.Repository) {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailablePackagesDto) GetRepositoryOk() (*Repository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *AvailablePackagesDto) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *AvailablePackagesDto) SetRepository(v Repository) {
	o.Repository = &v
}

func (o AvailablePackagesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailablePackagesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PackageId.IsSet() {
		toSerialize["packageId"] = o.PackageId.Get()
	}
	if o.CatalogAppId.IsSet() {
		toSerialize["catalogAppId"] = o.CatalogAppId.Get()
	}
	if o.InstalledInstanceCount.IsSet() {
		toSerialize["installedInstanceCount"] = o.InstalledInstanceCount.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.NormalizedName.IsSet() {
		toSerialize["normalizedName"] = o.NormalizedName.Get()
	}
	if o.LogoImageId.IsSet() {
		toSerialize["logoImageId"] = o.LogoImageId.Get()
	}
	if !IsNil(o.Stars) {
		toSerialize["stars"] = o.Stars
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.AppVersion.IsSet() {
		toSerialize["appVersion"] = o.AppVersion.Get()
	}
	if !IsNil(o.Deprecated) {
		toSerialize["deprecated"] = o.Deprecated
	}
	if !IsNil(o.Signed) {
		toSerialize["signed"] = o.Signed
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.SecurityReportSummary) {
		toSerialize["securityReportSummary"] = o.SecurityReportSummary
	}
	if o.Ts.IsSet() {
		toSerialize["ts"] = o.Ts.Get()
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableAvailablePackagesDto struct {
	value *AvailablePackagesDto
	isSet bool
}

func (v NullableAvailablePackagesDto) Get() *AvailablePackagesDto {
	return v.value
}

func (v *NullableAvailablePackagesDto) Set(val *AvailablePackagesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailablePackagesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailablePackagesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailablePackagesDto(val *AvailablePackagesDto) *NullableAvailablePackagesDto {
	return &NullableAvailablePackagesDto{value: val, isSet: true}
}

func (v NullableAvailablePackagesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailablePackagesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


