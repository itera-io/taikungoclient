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

// checks if the Repository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Repository{}

// Repository struct for Repository
type Repository struct {
	Url NullableString `json:"url,omitempty"`
	Kind *int64 `json:"kind,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Official *bool `json:"official,omitempty"`
	RepositoryId NullableString `json:"repositoryId,omitempty"`
	ScannerDisabled *bool `json:"scannerDisabled,omitempty"`
	IsImported *bool `json:"isImported,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	VerifiedPublisher *bool `json:"verifiedPublisher,omitempty"`
	OrganizationDisplayName NullableString `json:"organizationDisplayName,omitempty"`
}

// NewRepository instantiates a new Repository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepository() *Repository {
	this := Repository{}
	return &this
}

// NewRepositoryWithDefaults instantiates a new Repository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoryWithDefaults() *Repository {
	this := Repository{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Repository) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Repository) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Repository) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Repository) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Repository) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Repository) UnsetUrl() {
	o.Url.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Repository) GetKind() int64 {
	if o == nil || IsNil(o.Kind) {
		var ret int64
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetKindOk() (*int64, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Repository) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given int64 and assigns it to the Kind field.
func (o *Repository) SetKind(v int64) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Repository) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Repository) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Repository) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Repository) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Repository) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Repository) UnsetName() {
	o.Name.Unset()
}

// GetOfficial returns the Official field value if set, zero value otherwise.
func (o *Repository) GetOfficial() bool {
	if o == nil || IsNil(o.Official) {
		var ret bool
		return ret
	}
	return *o.Official
}

// GetOfficialOk returns a tuple with the Official field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetOfficialOk() (*bool, bool) {
	if o == nil || IsNil(o.Official) {
		return nil, false
	}
	return o.Official, true
}

// HasOfficial returns a boolean if a field has been set.
func (o *Repository) HasOfficial() bool {
	if o != nil && !IsNil(o.Official) {
		return true
	}

	return false
}

// SetOfficial gets a reference to the given bool and assigns it to the Official field.
func (o *Repository) SetOfficial(v bool) {
	o.Official = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Repository) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId.Get()) {
		var ret string
		return ret
	}
	return *o.RepositoryId.Get()
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Repository) GetRepositoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryId.Get(), o.RepositoryId.IsSet()
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *Repository) HasRepositoryId() bool {
	if o != nil && o.RepositoryId.IsSet() {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given NullableString and assigns it to the RepositoryId field.
func (o *Repository) SetRepositoryId(v string) {
	o.RepositoryId.Set(&v)
}
// SetRepositoryIdNil sets the value for RepositoryId to be an explicit nil
func (o *Repository) SetRepositoryIdNil() {
	o.RepositoryId.Set(nil)
}

// UnsetRepositoryId ensures that no value is present for RepositoryId, not even an explicit nil
func (o *Repository) UnsetRepositoryId() {
	o.RepositoryId.Unset()
}

// GetScannerDisabled returns the ScannerDisabled field value if set, zero value otherwise.
func (o *Repository) GetScannerDisabled() bool {
	if o == nil || IsNil(o.ScannerDisabled) {
		var ret bool
		return ret
	}
	return *o.ScannerDisabled
}

// GetScannerDisabledOk returns a tuple with the ScannerDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetScannerDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScannerDisabled) {
		return nil, false
	}
	return o.ScannerDisabled, true
}

// HasScannerDisabled returns a boolean if a field has been set.
func (o *Repository) HasScannerDisabled() bool {
	if o != nil && !IsNil(o.ScannerDisabled) {
		return true
	}

	return false
}

// SetScannerDisabled gets a reference to the given bool and assigns it to the ScannerDisabled field.
func (o *Repository) SetScannerDisabled(v bool) {
	o.ScannerDisabled = &v
}

// GetIsImported returns the IsImported field value if set, zero value otherwise.
func (o *Repository) GetIsImported() bool {
	if o == nil || IsNil(o.IsImported) {
		var ret bool
		return ret
	}
	return *o.IsImported
}

// GetIsImportedOk returns a tuple with the IsImported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetIsImportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsImported) {
		return nil, false
	}
	return o.IsImported, true
}

// HasIsImported returns a boolean if a field has been set.
func (o *Repository) HasIsImported() bool {
	if o != nil && !IsNil(o.IsImported) {
		return true
	}

	return false
}

// SetIsImported gets a reference to the given bool and assigns it to the IsImported field.
func (o *Repository) SetIsImported(v bool) {
	o.IsImported = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Repository) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Repository) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *Repository) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *Repository) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *Repository) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *Repository) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetVerifiedPublisher returns the VerifiedPublisher field value if set, zero value otherwise.
func (o *Repository) GetVerifiedPublisher() bool {
	if o == nil || IsNil(o.VerifiedPublisher) {
		var ret bool
		return ret
	}
	return *o.VerifiedPublisher
}

// GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repository) GetVerifiedPublisherOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifiedPublisher) {
		return nil, false
	}
	return o.VerifiedPublisher, true
}

// HasVerifiedPublisher returns a boolean if a field has been set.
func (o *Repository) HasVerifiedPublisher() bool {
	if o != nil && !IsNil(o.VerifiedPublisher) {
		return true
	}

	return false
}

// SetVerifiedPublisher gets a reference to the given bool and assigns it to the VerifiedPublisher field.
func (o *Repository) SetVerifiedPublisher(v bool) {
	o.VerifiedPublisher = &v
}

// GetOrganizationDisplayName returns the OrganizationDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Repository) GetOrganizationDisplayName() string {
	if o == nil || IsNil(o.OrganizationDisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationDisplayName.Get()
}

// GetOrganizationDisplayNameOk returns a tuple with the OrganizationDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Repository) GetOrganizationDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationDisplayName.Get(), o.OrganizationDisplayName.IsSet()
}

// HasOrganizationDisplayName returns a boolean if a field has been set.
func (o *Repository) HasOrganizationDisplayName() bool {
	if o != nil && o.OrganizationDisplayName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationDisplayName gets a reference to the given NullableString and assigns it to the OrganizationDisplayName field.
func (o *Repository) SetOrganizationDisplayName(v string) {
	o.OrganizationDisplayName.Set(&v)
}
// SetOrganizationDisplayNameNil sets the value for OrganizationDisplayName to be an explicit nil
func (o *Repository) SetOrganizationDisplayNameNil() {
	o.OrganizationDisplayName.Set(nil)
}

// UnsetOrganizationDisplayName ensures that no value is present for OrganizationDisplayName, not even an explicit nil
func (o *Repository) UnsetOrganizationDisplayName() {
	o.OrganizationDisplayName.Unset()
}

func (o Repository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Repository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Official) {
		toSerialize["official"] = o.Official
	}
	if o.RepositoryId.IsSet() {
		toSerialize["repositoryId"] = o.RepositoryId.Get()
	}
	if !IsNil(o.ScannerDisabled) {
		toSerialize["scannerDisabled"] = o.ScannerDisabled
	}
	if !IsNil(o.IsImported) {
		toSerialize["isImported"] = o.IsImported
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if !IsNil(o.VerifiedPublisher) {
		toSerialize["verifiedPublisher"] = o.VerifiedPublisher
	}
	if o.OrganizationDisplayName.IsSet() {
		toSerialize["organizationDisplayName"] = o.OrganizationDisplayName.Get()
	}
	return toSerialize, nil
}

type NullableRepository struct {
	value *Repository
	isSet bool
}

func (v NullableRepository) Get() *Repository {
	return v.value
}

func (v *NullableRepository) Set(val *Repository) {
	v.value = val
	v.isSet = true
}

func (v NullableRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepository(val *Repository) *NullableRepository {
	return &NullableRepository{value: val, isSet: true}
}

func (v NullableRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


