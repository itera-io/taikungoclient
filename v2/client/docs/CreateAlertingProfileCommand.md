# CreateAlertingProfileCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SlackConfigurationId** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**Emails** | Pointer to [**[]AlertingEmailDto**](AlertingEmailDto.md) |  | [optional] 
**Webhooks** | Pointer to [**[]AlertingWebhookDto**](AlertingWebhookDto.md) |  | [optional] 
**AlertingIntegrations** | Pointer to [**[]AlertingIntegrationDto**](AlertingIntegrationDto.md) |  | [optional] 
**Reminder** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateAlertingProfileCommand

`func NewCreateAlertingProfileCommand() *CreateAlertingProfileCommand`

NewCreateAlertingProfileCommand instantiates a new CreateAlertingProfileCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertingProfileCommandWithDefaults

`func NewCreateAlertingProfileCommandWithDefaults() *CreateAlertingProfileCommand`

NewCreateAlertingProfileCommandWithDefaults instantiates a new CreateAlertingProfileCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAlertingProfileCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAlertingProfileCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAlertingProfileCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAlertingProfileCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateAlertingProfileCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAlertingProfileCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlackConfigurationId

`func (o *CreateAlertingProfileCommand) GetSlackConfigurationId() int32`

GetSlackConfigurationId returns the SlackConfigurationId field if non-nil, zero value otherwise.

### GetSlackConfigurationIdOk

`func (o *CreateAlertingProfileCommand) GetSlackConfigurationIdOk() (*int32, bool)`

GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationId

`func (o *CreateAlertingProfileCommand) SetSlackConfigurationId(v int32)`

SetSlackConfigurationId sets SlackConfigurationId field to given value.

### HasSlackConfigurationId

`func (o *CreateAlertingProfileCommand) HasSlackConfigurationId() bool`

HasSlackConfigurationId returns a boolean if a field has been set.

### SetSlackConfigurationIdNil

`func (o *CreateAlertingProfileCommand) SetSlackConfigurationIdNil(b bool)`

 SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil

### UnsetSlackConfigurationId
`func (o *CreateAlertingProfileCommand) UnsetSlackConfigurationId()`

UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
### GetOrganizationId

`func (o *CreateAlertingProfileCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAlertingProfileCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAlertingProfileCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAlertingProfileCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateAlertingProfileCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAlertingProfileCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetEmails

`func (o *CreateAlertingProfileCommand) GetEmails() []AlertingEmailDto`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateAlertingProfileCommand) GetEmailsOk() (*[]AlertingEmailDto, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateAlertingProfileCommand) SetEmails(v []AlertingEmailDto)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *CreateAlertingProfileCommand) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### SetEmailsNil

`func (o *CreateAlertingProfileCommand) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *CreateAlertingProfileCommand) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetWebhooks

`func (o *CreateAlertingProfileCommand) GetWebhooks() []AlertingWebhookDto`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *CreateAlertingProfileCommand) GetWebhooksOk() (*[]AlertingWebhookDto, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *CreateAlertingProfileCommand) SetWebhooks(v []AlertingWebhookDto)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *CreateAlertingProfileCommand) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooksNil

`func (o *CreateAlertingProfileCommand) SetWebhooksNil(b bool)`

 SetWebhooksNil sets the value for Webhooks to be an explicit nil

### UnsetWebhooks
`func (o *CreateAlertingProfileCommand) UnsetWebhooks()`

UnsetWebhooks ensures that no value is present for Webhooks, not even an explicit nil
### GetAlertingIntegrations

`func (o *CreateAlertingProfileCommand) GetAlertingIntegrations() []AlertingIntegrationDto`

GetAlertingIntegrations returns the AlertingIntegrations field if non-nil, zero value otherwise.

### GetAlertingIntegrationsOk

`func (o *CreateAlertingProfileCommand) GetAlertingIntegrationsOk() (*[]AlertingIntegrationDto, bool)`

GetAlertingIntegrationsOk returns a tuple with the AlertingIntegrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrations

`func (o *CreateAlertingProfileCommand) SetAlertingIntegrations(v []AlertingIntegrationDto)`

SetAlertingIntegrations sets AlertingIntegrations field to given value.

### HasAlertingIntegrations

`func (o *CreateAlertingProfileCommand) HasAlertingIntegrations() bool`

HasAlertingIntegrations returns a boolean if a field has been set.

### SetAlertingIntegrationsNil

`func (o *CreateAlertingProfileCommand) SetAlertingIntegrationsNil(b bool)`

 SetAlertingIntegrationsNil sets the value for AlertingIntegrations to be an explicit nil

### UnsetAlertingIntegrations
`func (o *CreateAlertingProfileCommand) UnsetAlertingIntegrations()`

UnsetAlertingIntegrations ensures that no value is present for AlertingIntegrations, not even an explicit nil
### GetReminder

`func (o *CreateAlertingProfileCommand) GetReminder() map[string]interface{}`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *CreateAlertingProfileCommand) GetReminderOk() (*map[string]interface{}, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *CreateAlertingProfileCommand) SetReminder(v map[string]interface{})`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *CreateAlertingProfileCommand) HasReminder() bool`

HasReminder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


