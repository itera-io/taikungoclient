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

// checks if the UpdateUserCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserCommand{}

// UpdateUserCommand struct for UpdateUserCommand
type UpdateUserCommand struct {
	Id NullableString `json:"id,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Role *UserRole `json:"role,omitempty"`
	ForceToResetPassword *bool `json:"forceToResetPassword,omitempty"`
	Disable *bool `json:"disable,omitempty"`
	IsApprovedByPartner *bool `json:"isApprovedByPartner,omitempty"`
}

// NewUpdateUserCommand instantiates a new UpdateUserCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserCommand() *UpdateUserCommand {
	this := UpdateUserCommand{}
	return &this
}

// NewUpdateUserCommandWithDefaults instantiates a new UpdateUserCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserCommandWithDefaults() *UpdateUserCommand {
	this := UpdateUserCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserCommand) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserCommand) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *UpdateUserCommand) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UpdateUserCommand) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UpdateUserCommand) UnsetId() {
	o.Id.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserCommand) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserCommand) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UpdateUserCommand) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UpdateUserCommand) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UpdateUserCommand) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserCommand) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserCommand) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *UpdateUserCommand) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *UpdateUserCommand) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *UpdateUserCommand) UnsetUsername() {
	o.Username.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserCommand) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserCommand) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateUserCommand) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateUserCommand) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateUserCommand) UnsetEmail() {
	o.Email.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UpdateUserCommand) GetRole() UserRole {
	if o == nil || IsNil(o.Role) {
		var ret UserRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserCommand) GetRoleOk() (*UserRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRole and assigns it to the Role field.
func (o *UpdateUserCommand) SetRole(v UserRole) {
	o.Role = &v
}

// GetForceToResetPassword returns the ForceToResetPassword field value if set, zero value otherwise.
func (o *UpdateUserCommand) GetForceToResetPassword() bool {
	if o == nil || IsNil(o.ForceToResetPassword) {
		var ret bool
		return ret
	}
	return *o.ForceToResetPassword
}

// GetForceToResetPasswordOk returns a tuple with the ForceToResetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserCommand) GetForceToResetPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceToResetPassword) {
		return nil, false
	}
	return o.ForceToResetPassword, true
}

// HasForceToResetPassword returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasForceToResetPassword() bool {
	if o != nil && !IsNil(o.ForceToResetPassword) {
		return true
	}

	return false
}

// SetForceToResetPassword gets a reference to the given bool and assigns it to the ForceToResetPassword field.
func (o *UpdateUserCommand) SetForceToResetPassword(v bool) {
	o.ForceToResetPassword = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *UpdateUserCommand) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserCommand) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *UpdateUserCommand) SetDisable(v bool) {
	o.Disable = &v
}

// GetIsApprovedByPartner returns the IsApprovedByPartner field value if set, zero value otherwise.
func (o *UpdateUserCommand) GetIsApprovedByPartner() bool {
	if o == nil || IsNil(o.IsApprovedByPartner) {
		var ret bool
		return ret
	}
	return *o.IsApprovedByPartner
}

// GetIsApprovedByPartnerOk returns a tuple with the IsApprovedByPartner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserCommand) GetIsApprovedByPartnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApprovedByPartner) {
		return nil, false
	}
	return o.IsApprovedByPartner, true
}

// HasIsApprovedByPartner returns a boolean if a field has been set.
func (o *UpdateUserCommand) HasIsApprovedByPartner() bool {
	if o != nil && !IsNil(o.IsApprovedByPartner) {
		return true
	}

	return false
}

// SetIsApprovedByPartner gets a reference to the given bool and assigns it to the IsApprovedByPartner field.
func (o *UpdateUserCommand) SetIsApprovedByPartner(v bool) {
	o.IsApprovedByPartner = &v
}

func (o UpdateUserCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.ForceToResetPassword) {
		toSerialize["forceToResetPassword"] = o.ForceToResetPassword
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.IsApprovedByPartner) {
		toSerialize["isApprovedByPartner"] = o.IsApprovedByPartner
	}
	return toSerialize, nil
}

type NullableUpdateUserCommand struct {
	value *UpdateUserCommand
	isSet bool
}

func (v NullableUpdateUserCommand) Get() *UpdateUserCommand {
	return v.value
}

func (v *NullableUpdateUserCommand) Set(val *UpdateUserCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserCommand(val *UpdateUserCommand) *NullableUpdateUserCommand {
	return &NullableUpdateUserCommand{value: val, isSet: true}
}

func (v NullableUpdateUserCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


