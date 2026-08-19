# CreateSlackConfigurationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **NullableInt32** |  | 
**Name** | **NullableString** |  | 
**Url** | **NullableString** |  | 
**Channel** | **NullableString** |  | 
**SlackType** | [**SlackType**](SlackType.md) |  | 

## Methods

### NewCreateSlackConfigurationCommand

`func NewCreateSlackConfigurationCommand(organizationId NullableInt32, name NullableString, url NullableString, channel NullableString, slackType SlackType, ) *CreateSlackConfigurationCommand`

NewCreateSlackConfigurationCommand instantiates a new CreateSlackConfigurationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSlackConfigurationCommandWithDefaults

`func NewCreateSlackConfigurationCommandWithDefaults() *CreateSlackConfigurationCommand`

NewCreateSlackConfigurationCommandWithDefaults instantiates a new CreateSlackConfigurationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *CreateSlackConfigurationCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateSlackConfigurationCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateSlackConfigurationCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *CreateSlackConfigurationCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateSlackConfigurationCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetName

`func (o *CreateSlackConfigurationCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSlackConfigurationCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSlackConfigurationCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateSlackConfigurationCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateSlackConfigurationCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *CreateSlackConfigurationCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateSlackConfigurationCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateSlackConfigurationCommand) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *CreateSlackConfigurationCommand) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateSlackConfigurationCommand) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetChannel

`func (o *CreateSlackConfigurationCommand) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateSlackConfigurationCommand) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateSlackConfigurationCommand) SetChannel(v string)`

SetChannel sets Channel field to given value.


### SetChannelNil

`func (o *CreateSlackConfigurationCommand) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *CreateSlackConfigurationCommand) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetSlackType

`func (o *CreateSlackConfigurationCommand) GetSlackType() SlackType`

GetSlackType returns the SlackType field if non-nil, zero value otherwise.

### GetSlackTypeOk

`func (o *CreateSlackConfigurationCommand) GetSlackTypeOk() (*SlackType, bool)`

GetSlackTypeOk returns a tuple with the SlackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackType

`func (o *CreateSlackConfigurationCommand) SetSlackType(v SlackType)`

SetSlackType sets SlackType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


