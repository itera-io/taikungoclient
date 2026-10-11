# AlertingProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**OrganizationId** | **NullableInt32** |  | 
**OrganizationName** | **NullableString** |  | 
**SlackConfigurationId** | **NullableInt32** |  | 
**SlackConfigurationName** | **NullableString** |  | 
**IsLocked** | **bool** |  | 
**Emails** | [**[]AlertingEmailDto**](AlertingEmailDto.md) |  | 
**Webhooks** | [**[]AlertingWebhookDto**](AlertingWebhookDto.md) |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**Reminder** | **map[string]interface{}** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewAlertingProfilesListDto

`func NewAlertingProfilesListDto(id int32, name string, organizationId NullableInt32, organizationName NullableString, slackConfigurationId NullableInt32, slackConfigurationName NullableString, isLocked bool, emails []AlertingEmailDto, webhooks []AlertingWebhookDto, projects []CommonDropdownDto, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, reminder map[string]interface{}, createdAt string, ) *AlertingProfilesListDto`

NewAlertingProfilesListDto instantiates a new AlertingProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingProfilesListDtoWithDefaults

`func NewAlertingProfilesListDtoWithDefaults() *AlertingProfilesListDto`

NewAlertingProfilesListDtoWithDefaults instantiates a new AlertingProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertingProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AlertingProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertingProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertingProfilesListDto) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *AlertingProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AlertingProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AlertingProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *AlertingProfilesListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *AlertingProfilesListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *AlertingProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AlertingProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AlertingProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *AlertingProfilesListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *AlertingProfilesListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetSlackConfigurationId

`func (o *AlertingProfilesListDto) GetSlackConfigurationId() int32`

GetSlackConfigurationId returns the SlackConfigurationId field if non-nil, zero value otherwise.

### GetSlackConfigurationIdOk

`func (o *AlertingProfilesListDto) GetSlackConfigurationIdOk() (*int32, bool)`

GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationId

`func (o *AlertingProfilesListDto) SetSlackConfigurationId(v int32)`

SetSlackConfigurationId sets SlackConfigurationId field to given value.


### SetSlackConfigurationIdNil

`func (o *AlertingProfilesListDto) SetSlackConfigurationIdNil(b bool)`

 SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil

### UnsetSlackConfigurationId
`func (o *AlertingProfilesListDto) UnsetSlackConfigurationId()`

UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
### GetSlackConfigurationName

`func (o *AlertingProfilesListDto) GetSlackConfigurationName() string`

GetSlackConfigurationName returns the SlackConfigurationName field if non-nil, zero value otherwise.

### GetSlackConfigurationNameOk

`func (o *AlertingProfilesListDto) GetSlackConfigurationNameOk() (*string, bool)`

GetSlackConfigurationNameOk returns a tuple with the SlackConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationName

`func (o *AlertingProfilesListDto) SetSlackConfigurationName(v string)`

SetSlackConfigurationName sets SlackConfigurationName field to given value.


### SetSlackConfigurationNameNil

`func (o *AlertingProfilesListDto) SetSlackConfigurationNameNil(b bool)`

 SetSlackConfigurationNameNil sets the value for SlackConfigurationName to be an explicit nil

### UnsetSlackConfigurationName
`func (o *AlertingProfilesListDto) UnsetSlackConfigurationName()`

UnsetSlackConfigurationName ensures that no value is present for SlackConfigurationName, not even an explicit nil
### GetIsLocked

`func (o *AlertingProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AlertingProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AlertingProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetEmails

`func (o *AlertingProfilesListDto) GetEmails() []AlertingEmailDto`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AlertingProfilesListDto) GetEmailsOk() (*[]AlertingEmailDto, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AlertingProfilesListDto) SetEmails(v []AlertingEmailDto)`

SetEmails sets Emails field to given value.


### GetWebhooks

`func (o *AlertingProfilesListDto) GetWebhooks() []AlertingWebhookDto`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *AlertingProfilesListDto) GetWebhooksOk() (*[]AlertingWebhookDto, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *AlertingProfilesListDto) SetWebhooks(v []AlertingWebhookDto)`

SetWebhooks sets Webhooks field to given value.


### GetProjects

`func (o *AlertingProfilesListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *AlertingProfilesListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *AlertingProfilesListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetCreatedBy

`func (o *AlertingProfilesListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertingProfilesListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertingProfilesListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *AlertingProfilesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AlertingProfilesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AlertingProfilesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *AlertingProfilesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *AlertingProfilesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *AlertingProfilesListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *AlertingProfilesListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *AlertingProfilesListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetReminder

`func (o *AlertingProfilesListDto) GetReminder() map[string]interface{}`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *AlertingProfilesListDto) GetReminderOk() (*map[string]interface{}, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *AlertingProfilesListDto) SetReminder(v map[string]interface{})`

SetReminder sets Reminder field to given value.


### GetCreatedAt

`func (o *AlertingProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertingProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertingProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


