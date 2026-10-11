# EditAlertingIntegrationCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Token** | Pointer to **NullableString** |  | [optional] 
**AlertingIntegrationType** | Pointer to [**AlertingIntegrationType**](AlertingIntegrationType.md) |  | [optional] 
**AlertingProfileId** | Pointer to **int32** |  | [optional] 

## Methods

### NewEditAlertingIntegrationCommand

`func NewEditAlertingIntegrationCommand() *EditAlertingIntegrationCommand`

NewEditAlertingIntegrationCommand instantiates a new EditAlertingIntegrationCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditAlertingIntegrationCommandWithDefaults

`func NewEditAlertingIntegrationCommandWithDefaults() *EditAlertingIntegrationCommand`

NewEditAlertingIntegrationCommandWithDefaults instantiates a new EditAlertingIntegrationCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditAlertingIntegrationCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditAlertingIntegrationCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditAlertingIntegrationCommand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EditAlertingIntegrationCommand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *EditAlertingIntegrationCommand) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EditAlertingIntegrationCommand) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EditAlertingIntegrationCommand) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EditAlertingIntegrationCommand) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EditAlertingIntegrationCommand) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EditAlertingIntegrationCommand) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetToken

`func (o *EditAlertingIntegrationCommand) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EditAlertingIntegrationCommand) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EditAlertingIntegrationCommand) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EditAlertingIntegrationCommand) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *EditAlertingIntegrationCommand) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *EditAlertingIntegrationCommand) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetAlertingIntegrationType

`func (o *EditAlertingIntegrationCommand) GetAlertingIntegrationType() AlertingIntegrationType`

GetAlertingIntegrationType returns the AlertingIntegrationType field if non-nil, zero value otherwise.

### GetAlertingIntegrationTypeOk

`func (o *EditAlertingIntegrationCommand) GetAlertingIntegrationTypeOk() (*AlertingIntegrationType, bool)`

GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrationType

`func (o *EditAlertingIntegrationCommand) SetAlertingIntegrationType(v AlertingIntegrationType)`

SetAlertingIntegrationType sets AlertingIntegrationType field to given value.

### HasAlertingIntegrationType

`func (o *EditAlertingIntegrationCommand) HasAlertingIntegrationType() bool`

HasAlertingIntegrationType returns a boolean if a field has been set.

### GetAlertingProfileId

`func (o *EditAlertingIntegrationCommand) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *EditAlertingIntegrationCommand) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *EditAlertingIntegrationCommand) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.

### HasAlertingProfileId

`func (o *EditAlertingIntegrationCommand) HasAlertingProfileId() bool`

HasAlertingProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


