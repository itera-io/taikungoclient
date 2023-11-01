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

// checks if the CatalogAppDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogAppDetailsDto{}

// CatalogAppDetailsDto struct for CatalogAppDetailsDto
type CatalogAppDetailsDto struct {
	Id                      *int32                    `json:"id,omitempty"`
	Name                    NullableString            `json:"name,omitempty"`
	AppRepoName             NullableString            `json:"appRepoName,omitempty"`
	AppRepoOrganizationName NullableString            `json:"appRepoOrganizationName,omitempty"`
	AppRepoId               *int32                    `json:"appRepoId,omitempty"`
	CatalogName             NullableString            `json:"catalogName,omitempty"`
	CatalogId               *int32                    `json:"catalogId,omitempty"`
	PackageId               NullableString            `json:"packageId,omitempty"`
	Version                 NullableString            `json:"version,omitempty"`
	LogoId                  NullableString            `json:"logoId,omitempty"`
	ProjectApps             []ProjectAppDto           `json:"projectApps,omitempty"`
	Description             NullableString            `json:"description,omitempty"`
	Readme                  NullableString            `json:"readme,omitempty"`
	SecurityReport          *SecurityReportSummaryDto `json:"securityReport,omitempty"`
	AppVersion              NullableString            `json:"appVersion,omitempty"`
	Stars                   *int32                    `json:"stars,omitempty"`
	VerifiedPublisher       *bool                     `json:"verifiedPublisher,omitempty"`
	Official                *bool                     `json:"official,omitempty"`
	HasJsonSchema           *bool                     `json:"hasJsonSchema,omitempty"`
}

// NewCatalogAppDetailsDto instantiates a new CatalogAppDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogAppDetailsDto() *CatalogAppDetailsDto {
	this := CatalogAppDetailsDto{}
	return &this
}

// NewCatalogAppDetailsDtoWithDefaults instantiates a new CatalogAppDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogAppDetailsDtoWithDefaults() *CatalogAppDetailsDto {
	this := CatalogAppDetailsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CatalogAppDetailsDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CatalogAppDetailsDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CatalogAppDetailsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetName() {
	o.Name.Unset()
}

// GetAppRepoName returns the AppRepoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetAppRepoName() string {
	if o == nil || IsNil(o.AppRepoName.Get()) {
		var ret string
		return ret
	}
	return *o.AppRepoName.Get()
}

// GetAppRepoNameOk returns a tuple with the AppRepoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetAppRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppRepoName.Get(), o.AppRepoName.IsSet()
}

// HasAppRepoName returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasAppRepoName() bool {
	if o != nil && o.AppRepoName.IsSet() {
		return true
	}

	return false
}

// SetAppRepoName gets a reference to the given NullableString and assigns it to the AppRepoName field.
func (o *CatalogAppDetailsDto) SetAppRepoName(v string) {
	o.AppRepoName.Set(&v)
}

// SetAppRepoNameNil sets the value for AppRepoName to be an explicit nil
func (o *CatalogAppDetailsDto) SetAppRepoNameNil() {
	o.AppRepoName.Set(nil)
}

// UnsetAppRepoName ensures that no value is present for AppRepoName, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetAppRepoName() {
	o.AppRepoName.Unset()
}

// GetAppRepoOrganizationName returns the AppRepoOrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetAppRepoOrganizationName() string {
	if o == nil || IsNil(o.AppRepoOrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.AppRepoOrganizationName.Get()
}

// GetAppRepoOrganizationNameOk returns a tuple with the AppRepoOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetAppRepoOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppRepoOrganizationName.Get(), o.AppRepoOrganizationName.IsSet()
}

// HasAppRepoOrganizationName returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasAppRepoOrganizationName() bool {
	if o != nil && o.AppRepoOrganizationName.IsSet() {
		return true
	}

	return false
}

// SetAppRepoOrganizationName gets a reference to the given NullableString and assigns it to the AppRepoOrganizationName field.
func (o *CatalogAppDetailsDto) SetAppRepoOrganizationName(v string) {
	o.AppRepoOrganizationName.Set(&v)
}

// SetAppRepoOrganizationNameNil sets the value for AppRepoOrganizationName to be an explicit nil
func (o *CatalogAppDetailsDto) SetAppRepoOrganizationNameNil() {
	o.AppRepoOrganizationName.Set(nil)
}

// UnsetAppRepoOrganizationName ensures that no value is present for AppRepoOrganizationName, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetAppRepoOrganizationName() {
	o.AppRepoOrganizationName.Unset()
}

// GetAppRepoId returns the AppRepoId field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetAppRepoId() int32 {
	if o == nil || IsNil(o.AppRepoId) {
		var ret int32
		return ret
	}
	return *o.AppRepoId
}

// GetAppRepoIdOk returns a tuple with the AppRepoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetAppRepoIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AppRepoId) {
		return nil, false
	}
	return o.AppRepoId, true
}

// HasAppRepoId returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasAppRepoId() bool {
	if o != nil && !IsNil(o.AppRepoId) {
		return true
	}

	return false
}

// SetAppRepoId gets a reference to the given int32 and assigns it to the AppRepoId field.
func (o *CatalogAppDetailsDto) SetAppRepoId(v int32) {
	o.AppRepoId = &v
}

// GetCatalogName returns the CatalogName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetCatalogName() string {
	if o == nil || IsNil(o.CatalogName.Get()) {
		var ret string
		return ret
	}
	return *o.CatalogName.Get()
}

// GetCatalogNameOk returns a tuple with the CatalogName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetCatalogNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CatalogName.Get(), o.CatalogName.IsSet()
}

// HasCatalogName returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasCatalogName() bool {
	if o != nil && o.CatalogName.IsSet() {
		return true
	}

	return false
}

// SetCatalogName gets a reference to the given NullableString and assigns it to the CatalogName field.
func (o *CatalogAppDetailsDto) SetCatalogName(v string) {
	o.CatalogName.Set(&v)
}

// SetCatalogNameNil sets the value for CatalogName to be an explicit nil
func (o *CatalogAppDetailsDto) SetCatalogNameNil() {
	o.CatalogName.Set(nil)
}

// UnsetCatalogName ensures that no value is present for CatalogName, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetCatalogName() {
	o.CatalogName.Unset()
}

// GetCatalogId returns the CatalogId field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetCatalogId() int32 {
	if o == nil || IsNil(o.CatalogId) {
		var ret int32
		return ret
	}
	return *o.CatalogId
}

// GetCatalogIdOk returns a tuple with the CatalogId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetCatalogIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CatalogId) {
		return nil, false
	}
	return o.CatalogId, true
}

// HasCatalogId returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasCatalogId() bool {
	if o != nil && !IsNil(o.CatalogId) {
		return true
	}

	return false
}

// SetCatalogId gets a reference to the given int32 and assigns it to the CatalogId field.
func (o *CatalogAppDetailsDto) SetCatalogId(v int32) {
	o.CatalogId = &v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetPackageId() string {
	if o == nil || IsNil(o.PackageId.Get()) {
		var ret string
		return ret
	}
	return *o.PackageId.Get()
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageId.Get(), o.PackageId.IsSet()
}

// HasPackageId returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasPackageId() bool {
	if o != nil && o.PackageId.IsSet() {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given NullableString and assigns it to the PackageId field.
func (o *CatalogAppDetailsDto) SetPackageId(v string) {
	o.PackageId.Set(&v)
}

// SetPackageIdNil sets the value for PackageId to be an explicit nil
func (o *CatalogAppDetailsDto) SetPackageIdNil() {
	o.PackageId.Set(nil)
}

// UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetPackageId() {
	o.PackageId.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *CatalogAppDetailsDto) SetVersion(v string) {
	o.Version.Set(&v)
}

// SetVersionNil sets the value for Version to be an explicit nil
func (o *CatalogAppDetailsDto) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetVersion() {
	o.Version.Unset()
}

// GetLogoId returns the LogoId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetLogoId() string {
	if o == nil || IsNil(o.LogoId.Get()) {
		var ret string
		return ret
	}
	return *o.LogoId.Get()
}

// GetLogoIdOk returns a tuple with the LogoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetLogoIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoId.Get(), o.LogoId.IsSet()
}

// HasLogoId returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasLogoId() bool {
	if o != nil && o.LogoId.IsSet() {
		return true
	}

	return false
}

// SetLogoId gets a reference to the given NullableString and assigns it to the LogoId field.
func (o *CatalogAppDetailsDto) SetLogoId(v string) {
	o.LogoId.Set(&v)
}

// SetLogoIdNil sets the value for LogoId to be an explicit nil
func (o *CatalogAppDetailsDto) SetLogoIdNil() {
	o.LogoId.Set(nil)
}

// UnsetLogoId ensures that no value is present for LogoId, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetLogoId() {
	o.LogoId.Unset()
}

// GetProjectApps returns the ProjectApps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetProjectApps() []ProjectAppDto {
	if o == nil {
		var ret []ProjectAppDto
		return ret
	}
	return o.ProjectApps
}

// GetProjectAppsOk returns a tuple with the ProjectApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetProjectAppsOk() ([]ProjectAppDto, bool) {
	if o == nil || IsNil(o.ProjectApps) {
		return nil, false
	}
	return o.ProjectApps, true
}

// HasProjectApps returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasProjectApps() bool {
	if o != nil && IsNil(o.ProjectApps) {
		return true
	}

	return false
}

// SetProjectApps gets a reference to the given []ProjectAppDto and assigns it to the ProjectApps field.
func (o *CatalogAppDetailsDto) SetProjectApps(v []ProjectAppDto) {
	o.ProjectApps = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CatalogAppDetailsDto) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CatalogAppDetailsDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetDescription() {
	o.Description.Unset()
}

// GetReadme returns the Readme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetReadme() string {
	if o == nil || IsNil(o.Readme.Get()) {
		var ret string
		return ret
	}
	return *o.Readme.Get()
}

// GetReadmeOk returns a tuple with the Readme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetReadmeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Readme.Get(), o.Readme.IsSet()
}

// HasReadme returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasReadme() bool {
	if o != nil && o.Readme.IsSet() {
		return true
	}

	return false
}

// SetReadme gets a reference to the given NullableString and assigns it to the Readme field.
func (o *CatalogAppDetailsDto) SetReadme(v string) {
	o.Readme.Set(&v)
}

// SetReadmeNil sets the value for Readme to be an explicit nil
func (o *CatalogAppDetailsDto) SetReadmeNil() {
	o.Readme.Set(nil)
}

// UnsetReadme ensures that no value is present for Readme, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetReadme() {
	o.Readme.Unset()
}

// GetSecurityReport returns the SecurityReport field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetSecurityReport() SecurityReportSummaryDto {
	if o == nil || IsNil(o.SecurityReport) {
		var ret SecurityReportSummaryDto
		return ret
	}
	return *o.SecurityReport
}

// GetSecurityReportOk returns a tuple with the SecurityReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetSecurityReportOk() (*SecurityReportSummaryDto, bool) {
	if o == nil || IsNil(o.SecurityReport) {
		return nil, false
	}
	return o.SecurityReport, true
}

// HasSecurityReport returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasSecurityReport() bool {
	if o != nil && !IsNil(o.SecurityReport) {
		return true
	}

	return false
}

// SetSecurityReport gets a reference to the given SecurityReportSummaryDto and assigns it to the SecurityReport field.
func (o *CatalogAppDetailsDto) SetSecurityReport(v SecurityReportSummaryDto) {
	o.SecurityReport = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CatalogAppDetailsDto) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CatalogAppDetailsDto) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// HasAppVersion returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasAppVersion() bool {
	if o != nil && o.AppVersion.IsSet() {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given NullableString and assigns it to the AppVersion field.
func (o *CatalogAppDetailsDto) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}

// SetAppVersionNil sets the value for AppVersion to be an explicit nil
func (o *CatalogAppDetailsDto) SetAppVersionNil() {
	o.AppVersion.Set(nil)
}

// UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
func (o *CatalogAppDetailsDto) UnsetAppVersion() {
	o.AppVersion.Unset()
}

// GetStars returns the Stars field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetStars() int32 {
	if o == nil || IsNil(o.Stars) {
		var ret int32
		return ret
	}
	return *o.Stars
}

// GetStarsOk returns a tuple with the Stars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetStarsOk() (*int32, bool) {
	if o == nil || IsNil(o.Stars) {
		return nil, false
	}
	return o.Stars, true
}

// HasStars returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasStars() bool {
	if o != nil && !IsNil(o.Stars) {
		return true
	}

	return false
}

// SetStars gets a reference to the given int32 and assigns it to the Stars field.
func (o *CatalogAppDetailsDto) SetStars(v int32) {
	o.Stars = &v
}

// GetVerifiedPublisher returns the VerifiedPublisher field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetVerifiedPublisher() bool {
	if o == nil || IsNil(o.VerifiedPublisher) {
		var ret bool
		return ret
	}
	return *o.VerifiedPublisher
}

// GetVerifiedPublisherOk returns a tuple with the VerifiedPublisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetVerifiedPublisherOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifiedPublisher) {
		return nil, false
	}
	return o.VerifiedPublisher, true
}

// HasVerifiedPublisher returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasVerifiedPublisher() bool {
	if o != nil && !IsNil(o.VerifiedPublisher) {
		return true
	}

	return false
}

// SetVerifiedPublisher gets a reference to the given bool and assigns it to the VerifiedPublisher field.
func (o *CatalogAppDetailsDto) SetVerifiedPublisher(v bool) {
	o.VerifiedPublisher = &v
}

// GetOfficial returns the Official field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetOfficial() bool {
	if o == nil || IsNil(o.Official) {
		var ret bool
		return ret
	}
	return *o.Official
}

// GetOfficialOk returns a tuple with the Official field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetOfficialOk() (*bool, bool) {
	if o == nil || IsNil(o.Official) {
		return nil, false
	}
	return o.Official, true
}

// HasOfficial returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasOfficial() bool {
	if o != nil && !IsNil(o.Official) {
		return true
	}

	return false
}

// SetOfficial gets a reference to the given bool and assigns it to the Official field.
func (o *CatalogAppDetailsDto) SetOfficial(v bool) {
	o.Official = &v
}

// GetHasJsonSchema returns the HasJsonSchema field value if set, zero value otherwise.
func (o *CatalogAppDetailsDto) GetHasJsonSchema() bool {
	if o == nil || IsNil(o.HasJsonSchema) {
		var ret bool
		return ret
	}
	return *o.HasJsonSchema
}

// GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogAppDetailsDto) GetHasJsonSchemaOk() (*bool, bool) {
	if o == nil || IsNil(o.HasJsonSchema) {
		return nil, false
	}
	return o.HasJsonSchema, true
}

// HasHasJsonSchema returns a boolean if a field has been set.
func (o *CatalogAppDetailsDto) HasHasJsonSchema() bool {
	if o != nil && !IsNil(o.HasJsonSchema) {
		return true
	}

	return false
}

// SetHasJsonSchema gets a reference to the given bool and assigns it to the HasJsonSchema field.
func (o *CatalogAppDetailsDto) SetHasJsonSchema(v bool) {
	o.HasJsonSchema = &v
}

func (o CatalogAppDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogAppDetailsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.AppRepoName.IsSet() {
		toSerialize["appRepoName"] = o.AppRepoName.Get()
	}
	if o.AppRepoOrganizationName.IsSet() {
		toSerialize["appRepoOrganizationName"] = o.AppRepoOrganizationName.Get()
	}
	if !IsNil(o.AppRepoId) {
		toSerialize["appRepoId"] = o.AppRepoId
	}
	if o.CatalogName.IsSet() {
		toSerialize["catalogName"] = o.CatalogName.Get()
	}
	if !IsNil(o.CatalogId) {
		toSerialize["catalogId"] = o.CatalogId
	}
	if o.PackageId.IsSet() {
		toSerialize["packageId"] = o.PackageId.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.LogoId.IsSet() {
		toSerialize["logoId"] = o.LogoId.Get()
	}
	if o.ProjectApps != nil {
		toSerialize["projectApps"] = o.ProjectApps
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Readme.IsSet() {
		toSerialize["readme"] = o.Readme.Get()
	}
	if !IsNil(o.SecurityReport) {
		toSerialize["securityReport"] = o.SecurityReport
	}
	if o.AppVersion.IsSet() {
		toSerialize["appVersion"] = o.AppVersion.Get()
	}
	if !IsNil(o.Stars) {
		toSerialize["stars"] = o.Stars
	}
	if !IsNil(o.VerifiedPublisher) {
		toSerialize["verifiedPublisher"] = o.VerifiedPublisher
	}
	if !IsNil(o.Official) {
		toSerialize["official"] = o.Official
	}
	if !IsNil(o.HasJsonSchema) {
		toSerialize["hasJsonSchema"] = o.HasJsonSchema
	}
	return toSerialize, nil
}

type NullableCatalogAppDetailsDto struct {
	value *CatalogAppDetailsDto
	isSet bool
}

func (v NullableCatalogAppDetailsDto) Get() *CatalogAppDetailsDto {
	return v.value
}

func (v *NullableCatalogAppDetailsDto) Set(val *CatalogAppDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogAppDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogAppDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogAppDetailsDto(val *CatalogAppDetailsDto) *NullableCatalogAppDetailsDto {
	return &NullableCatalogAppDetailsDto{value: val, isSet: true}
}

func (v NullableCatalogAppDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogAppDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}