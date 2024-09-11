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

// checks if the AlertingProfilesListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertingProfilesListDto{}

// AlertingProfilesListDto struct for AlertingProfilesListDto
type AlertingProfilesListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	SlackConfigurationId NullableInt32 `json:"slackConfigurationId,omitempty"`
	SlackConfigurationName *string `json:"slackConfigurationName,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	Emails []AlertingEmailDto `json:"emails,omitempty"`
	Webhooks []AlertingWebhookDto `json:"webhooks,omitempty"`
	Projects []CommonDropdownDto `json:"projects,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	LastModified *string `json:"lastModified,omitempty"`
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	Reminder *AlertingReminder `json:"reminder,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewAlertingProfilesListDto instantiates a new AlertingProfilesListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingProfilesListDto() *AlertingProfilesListDto {
	this := AlertingProfilesListDto{}
	return &this
}

// NewAlertingProfilesListDtoWithDefaults instantiates a new AlertingProfilesListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingProfilesListDtoWithDefaults() *AlertingProfilesListDto {
	this := AlertingProfilesListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AlertingProfilesListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertingProfilesListDto) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingProfilesListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingProfilesListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *AlertingProfilesListDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *AlertingProfilesListDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *AlertingProfilesListDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *AlertingProfilesListDto) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetSlackConfigurationId returns the SlackConfigurationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlertingProfilesListDto) GetSlackConfigurationId() int32 {
	if o == nil || IsNil(o.SlackConfigurationId.Get()) {
		var ret int32
		return ret
	}
	return *o.SlackConfigurationId.Get()
}

// GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlertingProfilesListDto) GetSlackConfigurationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackConfigurationId.Get(), o.SlackConfigurationId.IsSet()
}

// HasSlackConfigurationId returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasSlackConfigurationId() bool {
	if o != nil && o.SlackConfigurationId.IsSet() {
		return true
	}

	return false
}

// SetSlackConfigurationId gets a reference to the given NullableInt32 and assigns it to the SlackConfigurationId field.
func (o *AlertingProfilesListDto) SetSlackConfigurationId(v int32) {
	o.SlackConfigurationId.Set(&v)
}
// SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil
func (o *AlertingProfilesListDto) SetSlackConfigurationIdNil() {
	o.SlackConfigurationId.Set(nil)
}

// UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
func (o *AlertingProfilesListDto) UnsetSlackConfigurationId() {
	o.SlackConfigurationId.Unset()
}

// GetSlackConfigurationName returns the SlackConfigurationName field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetSlackConfigurationName() string {
	if o == nil || IsNil(o.SlackConfigurationName) {
		var ret string
		return ret
	}
	return *o.SlackConfigurationName
}

// GetSlackConfigurationNameOk returns a tuple with the SlackConfigurationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetSlackConfigurationNameOk() (*string, bool) {
	if o == nil || IsNil(o.SlackConfigurationName) {
		return nil, false
	}
	return o.SlackConfigurationName, true
}

// HasSlackConfigurationName returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasSlackConfigurationName() bool {
	if o != nil && !IsNil(o.SlackConfigurationName) {
		return true
	}

	return false
}

// SetSlackConfigurationName gets a reference to the given string and assigns it to the SlackConfigurationName field.
func (o *AlertingProfilesListDto) SetSlackConfigurationName(v string) {
	o.SlackConfigurationName = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *AlertingProfilesListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetEmails() []AlertingEmailDto {
	if o == nil || IsNil(o.Emails) {
		var ret []AlertingEmailDto
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetEmailsOk() ([]AlertingEmailDto, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []AlertingEmailDto and assigns it to the Emails field.
func (o *AlertingProfilesListDto) SetEmails(v []AlertingEmailDto) {
	o.Emails = v
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetWebhooks() []AlertingWebhookDto {
	if o == nil || IsNil(o.Webhooks) {
		var ret []AlertingWebhookDto
		return ret
	}
	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetWebhooksOk() ([]AlertingWebhookDto, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []AlertingWebhookDto and assigns it to the Webhooks field.
func (o *AlertingProfilesListDto) SetWebhooks(v []AlertingWebhookDto) {
	o.Webhooks = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetProjects() []CommonDropdownDto {
	if o == nil || IsNil(o.Projects) {
		var ret []CommonDropdownDto
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []CommonDropdownDto and assigns it to the Projects field.
func (o *AlertingProfilesListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AlertingProfilesListDto) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *AlertingProfilesListDto) SetLastModified(v string) {
	o.LastModified = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *AlertingProfilesListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetReminder returns the Reminder field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetReminder() AlertingReminder {
	if o == nil || IsNil(o.Reminder) {
		var ret AlertingReminder
		return ret
	}
	return *o.Reminder
}

// GetReminderOk returns a tuple with the Reminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetReminderOk() (*AlertingReminder, bool) {
	if o == nil || IsNil(o.Reminder) {
		return nil, false
	}
	return o.Reminder, true
}

// HasReminder returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasReminder() bool {
	if o != nil && !IsNil(o.Reminder) {
		return true
	}

	return false
}

// SetReminder gets a reference to the given AlertingReminder and assigns it to the Reminder field.
func (o *AlertingProfilesListDto) SetReminder(v AlertingReminder) {
	o.Reminder = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AlertingProfilesListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfilesListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AlertingProfilesListDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AlertingProfilesListDto) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o AlertingProfilesListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertingProfilesListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if o.SlackConfigurationId.IsSet() {
		toSerialize["slackConfigurationId"] = o.SlackConfigurationId.Get()
	}
	if !IsNil(o.SlackConfigurationName) {
		toSerialize["slackConfigurationName"] = o.SlackConfigurationName
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.Reminder) {
		toSerialize["reminder"] = o.Reminder
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAlertingProfilesListDto struct {
	value *AlertingProfilesListDto
	isSet bool
}

func (v NullableAlertingProfilesListDto) Get() *AlertingProfilesListDto {
	return v.value
}

func (v *NullableAlertingProfilesListDto) Set(val *AlertingProfilesListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingProfilesListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingProfilesListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingProfilesListDto(val *AlertingProfilesListDto) *NullableAlertingProfilesListDto {
	return &NullableAlertingProfilesListDto{value: val, isSet: true}
}

func (v NullableAlertingProfilesListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingProfilesListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


