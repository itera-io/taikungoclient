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
	"bytes"
	"fmt"
)

// checks if the NotificationListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationListDto{}

// NotificationListDto struct for NotificationListDto
type NotificationListDto struct {
	CreatedAt NullableString `json:"createdAt"`
	ActionMessage string `json:"actionMessage"`
	ActionStatus ActionStatus `json:"actionStatus"`
	Username NullableString `json:"username"`
	Category ActionType `json:"category"`
	ProjectName NullableString `json:"projectName"`
	ProjectId NullableInt32 `json:"projectId"`
	IsDeleted bool `json:"isDeleted"`
}

type _NotificationListDto NotificationListDto

// NewNotificationListDto instantiates a new NotificationListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationListDto(createdAt NullableString, actionMessage string, actionStatus ActionStatus, username NullableString, category ActionType, projectName NullableString, projectId NullableInt32, isDeleted bool) *NotificationListDto {
	this := NotificationListDto{}
	this.CreatedAt = createdAt
	this.ActionMessage = actionMessage
	this.ActionStatus = actionStatus
	this.Username = username
	this.Category = category
	this.ProjectName = projectName
	this.ProjectId = projectId
	this.IsDeleted = isDeleted
	return &this
}

// NewNotificationListDtoWithDefaults instantiates a new NotificationListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationListDtoWithDefaults() *NotificationListDto {
	this := NotificationListDto{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationListDto) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *NotificationListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetActionMessage returns the ActionMessage field value
func (o *NotificationListDto) GetActionMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionMessage
}

// GetActionMessageOk returns a tuple with the ActionMessage field value
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetActionMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionMessage, true
}

// SetActionMessage sets field value
func (o *NotificationListDto) SetActionMessage(v string) {
	o.ActionMessage = v
}

// GetActionStatus returns the ActionStatus field value
func (o *NotificationListDto) GetActionStatus() ActionStatus {
	if o == nil {
		var ret ActionStatus
		return ret
	}

	return o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetActionStatusOk() (*ActionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionStatus, true
}

// SetActionStatus sets field value
func (o *NotificationListDto) SetActionStatus(v ActionStatus) {
	o.ActionStatus = v
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationListDto) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}

	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// SetUsername sets field value
func (o *NotificationListDto) SetUsername(v string) {
	o.Username.Set(&v)
}

// GetCategory returns the Category field value
func (o *NotificationListDto) GetCategory() ActionType {
	if o == nil {
		var ret ActionType
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetCategoryOk() (*ActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *NotificationListDto) SetCategory(v ActionType) {
	o.Category = v
}

// GetProjectName returns the ProjectName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationListDto) GetProjectName() string {
	if o == nil || o.ProjectName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// SetProjectName sets field value
func (o *NotificationListDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// GetProjectId returns the ProjectId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NotificationListDto) GetProjectId() int32 {
	if o == nil || o.ProjectId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationListDto) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// SetProjectId sets field value
func (o *NotificationListDto) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}

// GetIsDeleted returns the IsDeleted field value
func (o *NotificationListDto) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *NotificationListDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *NotificationListDto) SetIsDeleted(v bool) {
	o.IsDeleted = v
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
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["actionMessage"] = o.ActionMessage
	toSerialize["actionStatus"] = o.ActionStatus
	toSerialize["username"] = o.Username.Get()
	toSerialize["category"] = o.Category
	toSerialize["projectName"] = o.ProjectName.Get()
	toSerialize["projectId"] = o.ProjectId.Get()
	toSerialize["isDeleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *NotificationListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"actionMessage",
		"actionStatus",
		"username",
		"category",
		"projectName",
		"projectId",
		"isDeleted",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNotificationListDto := _NotificationListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationListDto)

	if err != nil {
		return err
	}

	*o = NotificationListDto(varNotificationListDto)

	return err
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


