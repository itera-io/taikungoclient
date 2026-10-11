# AlertingIntegrationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Token** | **NullableString** |  | 
**AlertingIntegrationType** | [**AlertingIntegrationType**](AlertingIntegrationType.md) |  | 

## Methods

### NewAlertingIntegrationDto

`func NewAlertingIntegrationDto(url string, token NullableString, alertingIntegrationType AlertingIntegrationType, ) *AlertingIntegrationDto`

NewAlertingIntegrationDto instantiates a new AlertingIntegrationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingIntegrationDtoWithDefaults

`func NewAlertingIntegrationDtoWithDefaults() *AlertingIntegrationDto`

NewAlertingIntegrationDtoWithDefaults instantiates a new AlertingIntegrationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AlertingIntegrationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AlertingIntegrationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AlertingIntegrationDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetToken

`func (o *AlertingIntegrationDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AlertingIntegrationDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AlertingIntegrationDto) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *AlertingIntegrationDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *AlertingIntegrationDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetAlertingIntegrationType

`func (o *AlertingIntegrationDto) GetAlertingIntegrationType() AlertingIntegrationType`

GetAlertingIntegrationType returns the AlertingIntegrationType field if non-nil, zero value otherwise.

### GetAlertingIntegrationTypeOk

`func (o *AlertingIntegrationDto) GetAlertingIntegrationTypeOk() (*AlertingIntegrationType, bool)`

GetAlertingIntegrationTypeOk returns a tuple with the AlertingIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingIntegrationType

`func (o *AlertingIntegrationDto) SetAlertingIntegrationType(v AlertingIntegrationType)`

SetAlertingIntegrationType sets AlertingIntegrationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


