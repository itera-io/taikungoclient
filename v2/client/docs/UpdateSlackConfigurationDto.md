# UpdateSlackConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Channel** | Pointer to **NullableString** |  | [optional] 
**SlackType** | Pointer to [**SlackType**](SlackType.md) |  | [optional] 

## Methods

### NewUpdateSlackConfigurationDto

`func NewUpdateSlackConfigurationDto() *UpdateSlackConfigurationDto`

NewUpdateSlackConfigurationDto instantiates a new UpdateSlackConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSlackConfigurationDtoWithDefaults

`func NewUpdateSlackConfigurationDtoWithDefaults() *UpdateSlackConfigurationDto`

NewUpdateSlackConfigurationDtoWithDefaults instantiates a new UpdateSlackConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *UpdateSlackConfigurationDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateSlackConfigurationDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateSlackConfigurationDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UpdateSlackConfigurationDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *UpdateSlackConfigurationDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *UpdateSlackConfigurationDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetName

`func (o *UpdateSlackConfigurationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSlackConfigurationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSlackConfigurationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSlackConfigurationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSlackConfigurationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSlackConfigurationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *UpdateSlackConfigurationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateSlackConfigurationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateSlackConfigurationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateSlackConfigurationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *UpdateSlackConfigurationDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UpdateSlackConfigurationDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetChannel

`func (o *UpdateSlackConfigurationDto) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateSlackConfigurationDto) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateSlackConfigurationDto) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateSlackConfigurationDto) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *UpdateSlackConfigurationDto) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *UpdateSlackConfigurationDto) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetSlackType

`func (o *UpdateSlackConfigurationDto) GetSlackType() SlackType`

GetSlackType returns the SlackType field if non-nil, zero value otherwise.

### GetSlackTypeOk

`func (o *UpdateSlackConfigurationDto) GetSlackTypeOk() (*SlackType, bool)`

GetSlackTypeOk returns a tuple with the SlackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackType

`func (o *UpdateSlackConfigurationDto) SetSlackType(v SlackType)`

SetSlackType sets SlackType field to given value.

### HasSlackType

`func (o *UpdateSlackConfigurationDto) HasSlackType() bool`

HasSlackType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


