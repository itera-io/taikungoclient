# CreateAiCredentialCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**ApiKey** | **NullableString** |  | 
**OrganizationId** | **NullableInt32** |  | 

## Methods

### NewCreateAiCredentialCommand

`func NewCreateAiCredentialCommand(name NullableString, apiKey NullableString, organizationId NullableInt32, ) *CreateAiCredentialCommand`

NewCreateAiCredentialCommand instantiates a new CreateAiCredentialCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAiCredentialCommandWithDefaults

`func NewCreateAiCredentialCommandWithDefaults() *CreateAiCredentialCommand`

NewCreateAiCredentialCommandWithDefaults instantiates a new CreateAiCredentialCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAiCredentialCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAiCredentialCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAiCredentialCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateAiCredentialCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAiCredentialCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetApiKey

`func (o *CreateAiCredentialCommand) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateAiCredentialCommand) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateAiCredentialCommand) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *CreateAiCredentialCommand) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *CreateAiCredentialCommand) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetOrganizationId

`func (o *CreateAiCredentialCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAiCredentialCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAiCredentialCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *CreateAiCredentialCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAiCredentialCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


