# UpdateAlertingProfileCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SlackConfigurationId** | Pointer to **NullableInt32** |  | [optional] 
**Reminder** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateAlertingProfileCommand

`func NewUpdateAlertingProfileCommand() *UpdateAlertingProfileCommand`

NewUpdateAlertingProfileCommand instantiates a new UpdateAlertingProfileCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertingProfileCommandWithDefaults

`func NewUpdateAlertingProfileCommandWithDefaults() *UpdateAlertingProfileCommand`

NewUpdateAlertingProfileCommandWithDefaults instantiates a new UpdateAlertingProfileCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAlertingProfileCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAlertingProfileCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAlertingProfileCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateAlertingProfileCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateAlertingProfileCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAlertingProfileCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAlertingProfileCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAlertingProfileCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateAlertingProfileCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAlertingProfileCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlackConfigurationId

`func (o *UpdateAlertingProfileCommand) GetSlackConfigurationId() int32`

GetSlackConfigurationId returns the SlackConfigurationId field if non-nil, zero value otherwise.

### GetSlackConfigurationIdOk

`func (o *UpdateAlertingProfileCommand) GetSlackConfigurationIdOk() (*int32, bool)`

GetSlackConfigurationIdOk returns a tuple with the SlackConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfigurationId

`func (o *UpdateAlertingProfileCommand) SetSlackConfigurationId(v int32)`

SetSlackConfigurationId sets SlackConfigurationId field to given value.

### HasSlackConfigurationId

`func (o *UpdateAlertingProfileCommand) HasSlackConfigurationId() bool`

HasSlackConfigurationId returns a boolean if a field has been set.

### SetSlackConfigurationIdNil

`func (o *UpdateAlertingProfileCommand) SetSlackConfigurationIdNil(b bool)`

 SetSlackConfigurationIdNil sets the value for SlackConfigurationId to be an explicit nil

### UnsetSlackConfigurationId
`func (o *UpdateAlertingProfileCommand) UnsetSlackConfigurationId()`

UnsetSlackConfigurationId ensures that no value is present for SlackConfigurationId, not even an explicit nil
### GetReminder

`func (o *UpdateAlertingProfileCommand) GetReminder() map[string]interface{}`

GetReminder returns the Reminder field if non-nil, zero value otherwise.

### GetReminderOk

`func (o *UpdateAlertingProfileCommand) GetReminderOk() (*map[string]interface{}, bool)`

GetReminderOk returns a tuple with the Reminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminder

`func (o *UpdateAlertingProfileCommand) SetReminder(v map[string]interface{})`

SetReminder sets Reminder field to given value.

### HasReminder

`func (o *UpdateAlertingProfileCommand) HasReminder() bool`

HasReminder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


