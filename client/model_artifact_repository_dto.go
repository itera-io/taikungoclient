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

// checks if the ArtifactRepositoryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactRepositoryDto{}

// ArtifactRepositoryDto struct for ArtifactRepositoryDto
type ArtifactRepositoryDto struct {
	RepositoryId NullableString `json:"repositoryId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Url NullableString `json:"url,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	VerifiedPublisher *bool `json:"verifiedPublisher,omitempty"`
	Official *bool `json:"official,omitempty"`
	IsBound *bool `json:"isBound,omitempty"`
	IsTaikun *bool `json:"isTaikun,omitempty"`
	HasCatalogApp *bool `json:"hasCatalogApp,omitempty"`
}

// NewArtifactRepositoryDto instantiates a new ArtifactRepositoryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactRepositoryDto() *ArtifactRepositoryDto {
	this := ArtifactRepositoryDto{}
	return &this
}

// NewArtifactRepositoryDtoWithDefaults instantiates a new ArtifactRepositoryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactRepositoryDtoWithDefaults() *ArtifactRepositoryDto {
	this := ArtifactRepositoryDto{}
	return &this
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtifactRepositoryDto) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryId.Get()
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtifactRepositoryDto) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryId.Get(), o.RepositoryId.IsSet()
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasRepositoryId() bool {
	if o != nil && o.RepositoryId.IsSet() {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given NullableString and assigns it to the RepositoryId field.
func (o *ArtifactRepositoryDto) SetRepositoryId(v string) {
	o.RepositoryId.Set(&v)
}
// SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil
func (o *ArtifactRepositoryDto) SetRepositoryIdNil() {
	o.RepositoryId.Set(nil)
}

// UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
func (o *ArtifactRepositoryDto) UnsetRepositoryId() {
	o.RepositoryId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtifactRepositoryDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtifactRepositoryDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ArtifactRepositoryDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ArtifactRepositoryDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ArtifactRepositoryDto) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtifactRepositoryDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtifactRepositoryDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ArtifactRepositoryDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ArtifactRepositoryDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ArtifactRepositoryDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtifactRepositoryDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtifactRepositoryDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ArtifactRepositoryDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ArtifactRepositoryDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ArtifactRepositoryDto) UnsetUrl() {
	o.Url.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArtifactRepositoryDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArtifactRepositoryDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *ArtifactRepositoryDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *ArtifactRepositoryDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *ArtifactRepositoryDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ArtifactRepositoryDto) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRepositoryDto) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ArtifactRepositoryDto) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetVerifiedPublisher returns the VerifiedPublisher field value if set, zero value otherwise.
func (o *ArtifactRepositoryDto) GetVerifiedPublisher() bool {
	if o == nil || IsNil(o.VerifiedPublisher) {
		var ret bool
		return ret
	}
	return *o.VerifiedPublisher
}

// GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRepositoryDto) GetVerifiedPublisherOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifiedPublisher) {
		return nil, false
	}
	return o.VerifiedPublisher, true
}

// HasVerifiedPublisher returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasVerifiedPublisher() bool {
	if o != nil && !IsNil(o.VerifiedPublisher) {
		return true
	}

	return false
}

// SetVerifiedPublisher gets a reference to the given bool and assigns it to the VerifiedPublisher field.
func (o *ArtifactRepositoryDto) SetVerifiedPublisher(v bool) {
	o.VerifiedPublisher = &v
}

// GetOfficial returns the Official field value if set, zero value otherwise.
func (o *ArtifactRepositoryDto) GetOfficial() bool {
	if o == nil || IsNil(o.Official) {
		var ret bool
		return ret
	}
	return *o.Official
}

// GetOfficialOk returns a tuple with the Official field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRepositoryDto) GetOfficialOk() (*bool, bool) {
	if o == nil || IsNil(o.Official) {
		return nil, false
	}
	return o.Official, true
}

// HasOfficial returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasOfficial() bool {
	if o != nil && !IsNil(o.Official) {
		return true
	}

	return false
}

// SetOfficial gets a reference to the given bool and assigns it to the Official field.
func (o *ArtifactRepositoryDto) SetOfficial(v bool) {
	o.Official = &v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *ArtifactRepositoryDto) GetIsBound() bool {
	if o == nil || IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRepositoryDto) GetIsBoundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasIsBound() bool {
	if o != nil && !IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *ArtifactRepositoryDto) SetIsBound(v bool) {
	o.IsBound = &v
}

// GetIsTaikun returns the IsTaikun field value if set, zero value otherwise.
func (o *ArtifactRepositoryDto) GetIsTaikun() bool {
	if o == nil || IsNil(o.IsTaikun) {
		var ret bool
		return ret
	}
	return *o.IsTaikun
}

// GetIsTaikunOk returns a tuple with the IsTaikun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRepositoryDto) GetIsTaikunOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTaikun) {
		return nil, false
	}
	return o.IsTaikun, true
}

// HasIsTaikun returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasIsTaikun() bool {
	if o != nil && !IsNil(o.IsTaikun) {
		return true
	}

	return false
}

// SetIsTaikun gets a reference to the given bool and assigns it to the IsTaikun field.
func (o *ArtifactRepositoryDto) SetIsTaikun(v bool) {
	o.IsTaikun = &v
}

// GetHasCatalogApp returns the HasCatalogApp field value if set, zero value otherwise.
func (o *ArtifactRepositoryDto) GetHasCatalogApp() bool {
	if o == nil || IsNil(o.HasCatalogApp) {
		var ret bool
		return ret
	}
	return *o.HasCatalogApp
}

// GetHasCatalogAppOk returns a tuple with the HasCatalogApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRepositoryDto) GetHasCatalogAppOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCatalogApp) {
		return nil, false
	}
	return o.HasCatalogApp, true
}

// HasHasCatalogApp returns a boolean if a field has been set.
func (o *ArtifactRepositoryDto) HasHasCatalogApp() bool {
	if o != nil && !IsNil(o.HasCatalogApp) {
		return true
	}

	return false
}

// SetHasCatalogApp gets a reference to the given bool and assigns it to the HasCatalogApp field.
func (o *ArtifactRepositoryDto) SetHasCatalogApp(v bool) {
	o.HasCatalogApp = &v
}

func (o ArtifactRepositoryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactRepositoryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RepositoryId.IsSet() {
		toSerialize["repositoryId"] = o.RepositoryId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.VerifiedPublisher) {
		toSerialize["verifiedPublisher"] = o.VerifiedPublisher
	}
	if !IsNil(o.Official) {
		toSerialize["official"] = o.Official
	}
	if !IsNil(o.IsBound) {
		toSerialize["isBound"] = o.IsBound
	}
	if !IsNil(o.IsTaikun) {
		toSerialize["isTaikun"] = o.IsTaikun
	}
	if !IsNil(o.HasCatalogApp) {
		toSerialize["hasCatalogApp"] = o.HasCatalogApp
	}
	return toSerialize, nil
}

type NullableArtifactRepositoryDto struct {
	value *ArtifactRepositoryDto
	isSet bool
}

func (v NullableArtifactRepositoryDto) Get() *ArtifactRepositoryDto {
	return v.value
}

func (v *NullableArtifactRepositoryDto) Set(val *ArtifactRepositoryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactRepositoryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactRepositoryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactRepositoryDto(val *ArtifactRepositoryDto) *NullableArtifactRepositoryDto {
	return &NullableArtifactRepositoryDto{value: val, isSet: true}
}

func (v NullableArtifactRepositoryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactRepositoryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


