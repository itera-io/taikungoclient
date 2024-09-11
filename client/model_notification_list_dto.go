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

// checks if the NotificationListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationListDto{}

// NotificationListDto struct for NotificationListDto
type NotificationListDto struct {
	CreatedAt NullableString `json:"createdAt,omitempty"`
	ActionMessage NullableString `json:"actionMessage,omitempty"`
	ActionStatus *ActionStatus `json:"actionStatus,omitempty"`
	Username NullableString `json:"username,omitempty"`
	Category *ActionType `json:"category,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	ProjectId *int32 `json:"projectId,omitempty"`
	IsDeleted *bool `json:"isDeleted,omitempty"`
}

// NewNotificationListDto instantiates a new NotificationListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationListDto() *NotificationListDto {
	this := NotificationListDto{}
	return &this
}

// NewNotificationListDtoWithDefaults instantiates a new NotificationListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationListDtoWithDefaults() *NotificationListDto {
	this := NotificationListDto{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NotificationListDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *NotificationListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *NotificationListDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *NotificationListDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetActionMessage returns the ActionMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationListDto) GetActionMessage() string {
	if o == nil || IsNil(o.ActionMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ActionMessage.Get()
}

// GetActionMessageOk returns a tuple with the ActionMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetActionMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionMessage.Get(), o.ActionMessage.IsSet()
}

// HasActionMessage returns a boolean if a field has been set.
func (o *NotificationListDto) HasActionMessage() bool {
	if o != nil && o.ActionMessage.IsSet() {
		return true
	}

	return false
}

// SetActionMessage gets a reference to the given NullableString and assigns it to the ActionMessage field.
func (o *NotificationListDto) SetActionMessage(v string) {
	o.ActionMessage.Set(&v)
}
// SetActionMessageNil sets the value for ActionMessage to be an explicit nil
func (o *NotificationListDto) SetActionMessageNil() {
	o.ActionMessage.Set(nil)
}

// UnsetActionMessage ensures that no value is present for ActionMessage, not even an explicit nil
func (o *NotificationListDto) UnsetActionMessage() {
	o.ActionMessage.Unset()
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *NotificationListDto) GetActionStatus() ActionStatus {
	if o == nil || IsNil(o.ActionStatus) {
		var ret ActionStatus
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetActionStatusOk() (*ActionStatus, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *NotificationListDto) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given ActionStatus and assigns it to the ActionStatus field.
func (o *NotificationListDto) SetActionStatus(v ActionStatus) {
	o.ActionStatus = &v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationListDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *NotificationListDto) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *NotificationListDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *NotificationListDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *NotificationListDto) UnsetUsername() {
	o.Username.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *NotificationListDto) GetCategory() ActionType {
	if o == nil || IsNil(o.Category) {
		var ret ActionType
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetCategoryOk() (*ActionType, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *NotificationListDto) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given ActionType and assigns it to the Category field.
func (o *NotificationListDto) SetCategory(v ActionType) {
	o.Category = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationListDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *NotificationListDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *NotificationListDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *NotificationListDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *NotificationListDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *NotificationListDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *NotificationListDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *NotificationListDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *NotificationListDto) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *NotificationListDto) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *NotificationListDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o NotificationListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.ActionMessage.IsSet() {
		toSerialize["actionMessage"] = o.ActionMessage.Get()
	}
	if !IsNil(o.ActionStatus) {
		toSerialize["actionStatus"] = o.ActionStatus
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableNotificationListDto struct {
	value *NotificationListDto
	isSet bool
}

func (v NullableNotificationListDto) Get() *NotificationListDto {
	return v.value
}

func (v *NullableNotificationListDto) Set(val *NotificationListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationListDto(val *NotificationListDto) *NullableNotificationListDto {
	return &NullableNotificationListDto{value: val, isSet: true}
}

func (v NullableNotificationListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


