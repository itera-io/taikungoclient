# AlertingIntegrationsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Url** | **string** |  | 
**Token** | **NullableString** |  | 
**AlertingIntegrationType** | [**AlertingIntegrationType**](AlertingIntegrationType.md) |  | 
**AlertingProfileName** | **string** |  | 

## Methods

### NewAlertingIntegrationsListDto

`func NewAlertingIntegrationsListDto(id int32, url string, token NullableString, alertingIntegrationType AlertingIntegrationType, alertingProfileName string, ) *AlertingIntegrationsListDto`

NewAlertingIntegrationsListDto instantiates a new AlertingIntegrationsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingIntegrationsListDtoWithDefaults

`func NewAlertingIntegrationsListDtoWithDefaults() *AlertingIntegrationsListDto`

NewAlertingIntegrationsListDtoWithDefaults instantiates a new AlertingIntegrationsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertingIntegrationsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingIntegrationsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingIntegrationsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *AlertingIntegrationsListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AlertingIntegrationsListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AlertingIntegrationsListDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *AlertingIntegrationsListDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AlertingIntegrationsListDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AlertingIntegrationsListDto) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *AlertingIntegrationsListDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *AlertingIntegrationsListDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetAlertingIntegrationType

`func (o *AlertingIntegrationsListDto) GetAlertingIntegrationType() AlertingIntegrationType`

GetAlertingIntegrationType returns the AlertingIntegrationType field if non-nil, zero value otherwise.

### GetAlertingIntegrationTypeOk

`func (o *AlertingIntegrationsListDto) GetAlertingIntegrationTypeOk() (*AlertingIntegrationType, bool)`

GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrationType

`func (o *AlertingIntegrationsListDto) SetAlertingIntegrationType(v AlertingIntegrationType)`

SetAlertingIntegrationType sets AlertingIntegrationType field to given value.


### GetAlertingProfileName

`func (o *AlertingIntegrationsListDto) GetAlertingProfileName() string`

GetAlertingProfileName returns the AlertingProfileName field if non-nil, zero value otherwise.

### GetAlertingProfileNameOk

`func (o *AlertingIntegrationsListDto) GetAlertingProfileNameOk() (*string, bool)`

GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileName

`func (o *AlertingIntegrationsListDto) SetAlertingProfileName(v string)`

SetAlertingProfileName sets AlertingProfileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


