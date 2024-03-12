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

// checks if the UserForListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserForListDto{}

// UserForListDto struct for UserForListDto
type UserForListDto struct {
	Id NullableString `json:"id,omitempty"`
	Username NullableString `json:"username,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	HasCustomerId *bool `json:"hasCustomerId,omitempty"`
	HasPaymentMethod *bool `json:"hasPaymentMethod,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	Role *UserRole `json:"role,omitempty"`
	Email NullableString `json:"email,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	CreatedAt NullableString `json:"createdAt,omitempty"`
	IsEmailConfirmed *bool `json:"isEmailConfirmed,omitempty"`
	IsEmailNotificationEnabled *bool `json:"isEmailNotificationEnabled,omitempty"`
	IsForcedToResetPassword *bool `json:"isForcedToResetPassword,omitempty"`
	IsCsm *bool `json:"isCsm,omitempty"`
	IsEligibleUpdateSubscription *bool `json:"isEligibleUpdateSubscription,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	IsApprovedByPartner *bool `json:"isApprovedByPartner,omitempty"`
	Owner *bool `json:"owner,omitempty"`
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
	HasRepo *bool `json:"hasRepo,omitempty"`
	DemoModeEnabled *bool `json:"demoModeEnabled,omitempty"`
	IsNewOrganization *bool `json:"isNewOrganization,omitempty"`
	Is2FAEnabled *bool `json:"is2FAEnabled,omitempty"`
	LastLoginAt NullableString `json:"lastLoginAt,omitempty"`
	BoundProjects []ProjectDto `json:"boundProjects,omitempty"`
	Partner *PartnerDetailsForUserDto `json:"partner,omitempty"`
}

// NewUserForListDto instantiates a new UserForListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserForListDto() *UserForListDto {
	this := UserForListDto{}
	return &this
}

// NewUserForListDtoWithDefaults instantiates a new UserForListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserForListDtoWithDefaults() *UserForListDto {
	this := UserForListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UserForListDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *UserForListDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UserForListDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UserForListDto) UnsetId() {
	o.Id.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UserForListDto) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UserForListDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UserForListDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UserForListDto) UnsetUsername() {
	o.Username.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *UserForListDto) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *UserForListDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *UserForListDto) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *UserForListDto) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetHasCustomerId returns the HasCustomerId field value if set, zero value otherwise.
func (o *UserForListDto) GetHasCustomerId() bool {
	if o == nil || IsNil(o.HasCustomerId) {
		var ret bool
		return ret
	}
	return *o.HasCustomerId
}

// GetHasCustomerIdOk returns a tuple with the HasCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetHasCustomerIdOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCustomerId) {
		return nil, false
	}
	return o.HasCustomerId, true
}

// HasHasCustomerId returns a boolean if a field has been set.
func (o *UserForListDto) HasHasCustomerId() bool {
	if o != nil && !IsNil(o.HasCustomerId) {
		return true
	}

	return false
}

// SetHasCustomerId gets a reference to the given bool and assigns it to the HasCustomerId field.
func (o *UserForListDto) SetHasCustomerId(v bool) {
	o.HasCustomerId = &v
}

// GetHasPaymentMethod returns the HasPaymentMethod field value if set, zero value otherwise.
func (o *UserForListDto) GetHasPaymentMethod() bool {
	if o == nil || IsNil(o.HasPaymentMethod) {
		var ret bool
		return ret
	}
	return *o.HasPaymentMethod
}

// GetHasPaymentMethodOk returns a tuple with the HasPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetHasPaymentMethodOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPaymentMethod) {
		return nil, false
	}
	return o.HasPaymentMethod, true
}

// HasHasPaymentMethod returns a boolean if a field has been set.
func (o *UserForListDto) HasHasPaymentMethod() bool {
	if o != nil && !IsNil(o.HasPaymentMethod) {
		return true
	}

	return false
}

// SetHasPaymentMethod gets a reference to the given bool and assigns it to the HasPaymentMethod field.
func (o *UserForListDto) SetHasPaymentMethod(v bool) {
	o.HasPaymentMethod = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *UserForListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *UserForListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *UserForListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserForListDto) GetRole() UserRole {
	if o == nil || IsNil(o.Role) {
		var ret UserRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetRoleOk() (*UserRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserForListDto) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRole and assigns it to the Role field.
func (o *UserForListDto) SetRole(v UserRole) {
	o.Role = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserForListDto) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserForListDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserForListDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserForListDto) UnsetEmail() {
	o.Email.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserForListDto) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UserForListDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UserForListDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UserForListDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserForListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *UserForListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *UserForListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *UserForListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetIsEmailConfirmed returns the IsEmailConfirmed field value if set, zero value otherwise.
func (o *UserForListDto) GetIsEmailConfirmed() bool {
	if o == nil || IsNil(o.IsEmailConfirmed) {
		var ret bool
		return ret
	}
	return *o.IsEmailConfirmed
}

// GetIsEmailConfirmedOk returns a tuple with the IsEmailConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsEmailConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEmailConfirmed) {
		return nil, false
	}
	return o.IsEmailConfirmed, true
}

// HasIsEmailConfirmed returns a boolean if a field has been set.
func (o *UserForListDto) HasIsEmailConfirmed() bool {
	if o != nil && !IsNil(o.IsEmailConfirmed) {
		return true
	}

	return false
}

// SetIsEmailConfirmed gets a reference to the given bool and assigns it to the IsEmailConfirmed field.
func (o *UserForListDto) SetIsEmailConfirmed(v bool) {
	o.IsEmailConfirmed = &v
}

// GetIsEmailNotificationEnabled returns the IsEmailNotificationEnabled field value if set, zero value otherwise.
func (o *UserForListDto) GetIsEmailNotificationEnabled() bool {
	if o == nil || IsNil(o.IsEmailNotificationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEmailNotificationEnabled
}

// GetIsEmailNotificationEnabledOk returns a tuple with the IsEmailNotificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsEmailNotificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEmailNotificationEnabled) {
		return nil, false
	}
	return o.IsEmailNotificationEnabled, true
}

// HasIsEmailNotificationEnabled returns a boolean if a field has been set.
func (o *UserForListDto) HasIsEmailNotificationEnabled() bool {
	if o != nil && !IsNil(o.IsEmailNotificationEnabled) {
		return true
	}

	return false
}

// SetIsEmailNotificationEnabled gets a reference to the given bool and assigns it to the IsEmailNotificationEnabled field.
func (o *UserForListDto) SetIsEmailNotificationEnabled(v bool) {
	o.IsEmailNotificationEnabled = &v
}

// GetIsForcedToResetPassword returns the IsForcedToResetPassword field value if set, zero value otherwise.
func (o *UserForListDto) GetIsForcedToResetPassword() bool {
	if o == nil || IsNil(o.IsForcedToResetPassword) {
		var ret bool
		return ret
	}
	return *o.IsForcedToResetPassword
}

// GetIsForcedToResetPasswordOk returns a tuple with the IsForcedToResetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsForcedToResetPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.IsForcedToResetPassword) {
		return nil, false
	}
	return o.IsForcedToResetPassword, true
}

// HasIsForcedToResetPassword returns a boolean if a field has been set.
func (o *UserForListDto) HasIsForcedToResetPassword() bool {
	if o != nil && !IsNil(o.IsForcedToResetPassword) {
		return true
	}

	return false
}

// SetIsForcedToResetPassword gets a reference to the given bool and assigns it to the IsForcedToResetPassword field.
func (o *UserForListDto) SetIsForcedToResetPassword(v bool) {
	o.IsForcedToResetPassword = &v
}

// GetIsCsm returns the IsCsm field value if set, zero value otherwise.
func (o *UserForListDto) GetIsCsm() bool {
	if o == nil || IsNil(o.IsCsm) {
		var ret bool
		return ret
	}
	return *o.IsCsm
}

// GetIsCsmOk returns a tuple with the IsCsm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsCsmOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCsm) {
		return nil, false
	}
	return o.IsCsm, true
}

// HasIsCsm returns a boolean if a field has been set.
func (o *UserForListDto) HasIsCsm() bool {
	if o != nil && !IsNil(o.IsCsm) {
		return true
	}

	return false
}

// SetIsCsm gets a reference to the given bool and assigns it to the IsCsm field.
func (o *UserForListDto) SetIsCsm(v bool) {
	o.IsCsm = &v
}

// GetIsEligibleUpdateSubscription returns the IsEligibleUpdateSubscription field value if set, zero value otherwise.
func (o *UserForListDto) GetIsEligibleUpdateSubscription() bool {
	if o == nil || IsNil(o.IsEligibleUpdateSubscription) {
		var ret bool
		return ret
	}
	return *o.IsEligibleUpdateSubscription
}

// GetIsEligibleUpdateSubscriptionOk returns a tuple with the IsEligibleUpdateSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsEligibleUpdateSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEligibleUpdateSubscription) {
		return nil, false
	}
	return o.IsEligibleUpdateSubscription, true
}

// HasIsEligibleUpdateSubscription returns a boolean if a field has been set.
func (o *UserForListDto) HasIsEligibleUpdateSubscription() bool {
	if o != nil && !IsNil(o.IsEligibleUpdateSubscription) {
		return true
	}

	return false
}

// SetIsEligibleUpdateSubscription gets a reference to the given bool and assigns it to the IsEligibleUpdateSubscription field.
func (o *UserForListDto) SetIsEligibleUpdateSubscription(v bool) {
	o.IsEligibleUpdateSubscription = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *UserForListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *UserForListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *UserForListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetIsApprovedByPartner returns the IsApprovedByPartner field value if set, zero value otherwise.
func (o *UserForListDto) GetIsApprovedByPartner() bool {
	if o == nil || IsNil(o.IsApprovedByPartner) {
		var ret bool
		return ret
	}
	return *o.IsApprovedByPartner
}

// GetIsApprovedByPartnerOk returns a tuple with the IsApprovedByPartner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsApprovedByPartnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApprovedByPartner) {
		return nil, false
	}
	return o.IsApprovedByPartner, true
}

// HasIsApprovedByPartner returns a boolean if a field has been set.
func (o *UserForListDto) HasIsApprovedByPartner() bool {
	if o != nil && !IsNil(o.IsApprovedByPartner) {
		return true
	}

	return false
}

// SetIsApprovedByPartner gets a reference to the given bool and assigns it to the IsApprovedByPartner field.
func (o *UserForListDto) SetIsApprovedByPartner(v bool) {
	o.IsApprovedByPartner = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *UserForListDto) GetOwner() bool {
	if o == nil || IsNil(o.Owner) {
		var ret bool
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *UserForListDto) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given bool and assigns it to the Owner field.
func (o *UserForListDto) SetOwner(v bool) {
	o.Owner = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *UserForListDto) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *UserForListDto) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *UserForListDto) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetHasRepo returns the HasRepo field value if set, zero value otherwise.
func (o *UserForListDto) GetHasRepo() bool {
	if o == nil || IsNil(o.HasRepo) {
		var ret bool
		return ret
	}
	return *o.HasRepo
}

// GetHasRepoOk returns a tuple with the HasRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetHasRepoOk() (*bool, bool) {
	if o == nil || IsNil(o.HasRepo) {
		return nil, false
	}
	return o.HasRepo, true
}

// HasHasRepo returns a boolean if a field has been set.
func (o *UserForListDto) HasHasRepo() bool {
	if o != nil && !IsNil(o.HasRepo) {
		return true
	}

	return false
}

// SetHasRepo gets a reference to the given bool and assigns it to the HasRepo field.
func (o *UserForListDto) SetHasRepo(v bool) {
	o.HasRepo = &v
}

// GetDemoModeEnabled returns the DemoModeEnabled field value if set, zero value otherwise.
func (o *UserForListDto) GetDemoModeEnabled() bool {
	if o == nil || IsNil(o.DemoModeEnabled) {
		var ret bool
		return ret
	}
	return *o.DemoModeEnabled
}

// GetDemoModeEnabledOk returns a tuple with the DemoModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetDemoModeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DemoModeEnabled) {
		return nil, false
	}
	return o.DemoModeEnabled, true
}

// HasDemoModeEnabled returns a boolean if a field has been set.
func (o *UserForListDto) HasDemoModeEnabled() bool {
	if o != nil && !IsNil(o.DemoModeEnabled) {
		return true
	}

	return false
}

// SetDemoModeEnabled gets a reference to the given bool and assigns it to the DemoModeEnabled field.
func (o *UserForListDto) SetDemoModeEnabled(v bool) {
	o.DemoModeEnabled = &v
}

// GetIsNewOrganization returns the IsNewOrganization field value if set, zero value otherwise.
func (o *UserForListDto) GetIsNewOrganization() bool {
	if o == nil || IsNil(o.IsNewOrganization) {
		var ret bool
		return ret
	}
	return *o.IsNewOrganization
}

// GetIsNewOrganizationOk returns a tuple with the IsNewOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIsNewOrganizationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNewOrganization) {
		return nil, false
	}
	return o.IsNewOrganization, true
}

// HasIsNewOrganization returns a boolean if a field has been set.
func (o *UserForListDto) HasIsNewOrganization() bool {
	if o != nil && !IsNil(o.IsNewOrganization) {
		return true
	}

	return false
}

// SetIsNewOrganization gets a reference to the given bool and assigns it to the IsNewOrganization field.
func (o *UserForListDto) SetIsNewOrganization(v bool) {
	o.IsNewOrganization = &v
}

// GetIs2FAEnabled returns the Is2FAEnabled field value if set, zero value otherwise.
func (o *UserForListDto) GetIs2FAEnabled() bool {
	if o == nil || IsNil(o.Is2FAEnabled) {
		var ret bool
		return ret
	}
	return *o.Is2FAEnabled
}

// GetIs2FAEnabledOk returns a tuple with the Is2FAEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetIs2FAEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Is2FAEnabled) {
		return nil, false
	}
	return o.Is2FAEnabled, true
}

// HasIs2FAEnabled returns a boolean if a field has been set.
func (o *UserForListDto) HasIs2FAEnabled() bool {
	if o != nil && !IsNil(o.Is2FAEnabled) {
		return true
	}

	return false
}

// SetIs2FAEnabled gets a reference to the given bool and assigns it to the Is2FAEnabled field.
func (o *UserForListDto) SetIs2FAEnabled(v bool) {
	o.Is2FAEnabled = &v
}

// GetLastLoginAt returns the LastLoginAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetLastLoginAt() string {
	if o == nil || IsNil(o.LastLoginAt.Get()) {
		var ret string
		return ret
	}
	return *o.LastLoginAt.Get()
}

// GetLastLoginAtOk returns a tuple with the LastLoginAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetLastLoginAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoginAt.Get(), o.LastLoginAt.IsSet()
}

// HasLastLoginAt returns a boolean if a field has been set.
func (o *UserForListDto) HasLastLoginAt() bool {
	if o != nil && o.LastLoginAt.IsSet() {
		return true
	}

	return false
}

// SetLastLoginAt gets a reference to the given NullableString and assigns it to the LastLoginAt field.
func (o *UserForListDto) SetLastLoginAt(v string) {
	o.LastLoginAt.Set(&v)
}
// SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil
func (o *UserForListDto) SetLastLoginAtNil() {
	o.LastLoginAt.Set(nil)
}

// UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
func (o *UserForListDto) UnsetLastLoginAt() {
	o.LastLoginAt.Unset()
}

// GetBoundProjects returns the BoundProjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserForListDto) GetBoundProjects() []ProjectDto {
	if o == nil {
		var ret []ProjectDto
		return ret
	}
	return o.BoundProjects
}

// GetBoundProjectsOk returns a tuple with the BoundProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserForListDto) GetBoundProjectsOk() ([]ProjectDto, bool) {
	if o == nil || IsNil(o.BoundProjects) {
		return nil, false
	}
	return o.BoundProjects, true
}

// HasBoundProjects returns a boolean if a field has been set.
func (o *UserForListDto) HasBoundProjects() bool {
	if o != nil && !IsNil(o.BoundProjects) {
		return true
	}

	return false
}

// SetBoundProjects gets a reference to the given []ProjectDto and assigns it to the BoundProjects field.
func (o *UserForListDto) SetBoundProjects(v []ProjectDto) {
	o.BoundProjects = v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *UserForListDto) GetPartner() PartnerDetailsForUserDto {
	if o == nil || IsNil(o.Partner) {
		var ret PartnerDetailsForUserDto
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserForListDto) GetPartnerOk() (*PartnerDetailsForUserDto, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *UserForListDto) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given PartnerDetailsForUserDto and assigns it to the Partner field.
func (o *UserForListDto) SetPartner(v PartnerDetailsForUserDto) {
	o.Partner = &v
}

func (o UserForListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserForListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if !IsNil(o.HasCustomerId) {
		toSerialize["hasCustomerId"] = o.HasCustomerId
	}
	if !IsNil(o.HasPaymentMethod) {
		toSerialize["hasPaymentMethod"] = o.HasPaymentMethod
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.IsEmailConfirmed) {
		toSerialize["isEmailConfirmed"] = o.IsEmailConfirmed
	}
	if !IsNil(o.IsEmailNotificationEnabled) {
		toSerialize["isEmailNotificationEnabled"] = o.IsEmailNotificationEnabled
	}
	if !IsNil(o.IsForcedToResetPassword) {
		toSerialize["isForcedToResetPassword"] = o.IsForcedToResetPassword
	}
	if !IsNil(o.IsCsm) {
		toSerialize["isCsm"] = o.IsCsm
	}
	if !IsNil(o.IsEligibleUpdateSubscription) {
		toSerialize["isEligibleUpdateSubscription"] = o.IsEligibleUpdateSubscription
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.IsApprovedByPartner) {
		toSerialize["isApprovedByPartner"] = o.IsApprovedByPartner
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.IsReadOnly) {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if !IsNil(o.HasRepo) {
		toSerialize["hasRepo"] = o.HasRepo
	}
	if !IsNil(o.DemoModeEnabled) {
		toSerialize["demoModeEnabled"] = o.DemoModeEnabled
	}
	if !IsNil(o.IsNewOrganization) {
		toSerialize["isNewOrganization"] = o.IsNewOrganization
	}
	if !IsNil(o.Is2FAEnabled) {
		toSerialize["is2FAEnabled"] = o.Is2FAEnabled
	}
	if o.LastLoginAt.IsSet() {
		toSerialize["lastLoginAt"] = o.LastLoginAt.Get()
	}
	if o.BoundProjects != nil {
		toSerialize["boundProjects"] = o.BoundProjects
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	return toSerialize, nil
}

type NullableUserForListDto struct {
	value *UserForListDto
	isSet bool
}

func (v NullableUserForListDto) Get() *UserForListDto {
	return v.value
}

func (v *NullableUserForListDto) Set(val *UserForListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserForListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserForListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserForListDto(val *UserForListDto) *NullableUserForListDto {
	return &NullableUserForListDto{value: val, isSet: true}
}

func (v NullableUserForListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserForListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


