# AlertingWebhookDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Headers** | Pointer to [**[]WebhookHeaderDto**](WebhookHeaderDto.md) |  | [optional] 

## Methods

### NewAlertingWebhookDto

`func NewAlertingWebhookDto() *AlertingWebhookDto`

NewAlertingWebhookDto instantiates a new AlertingWebhookDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingWebhookDtoWithDefaults

`func NewAlertingWebhookDtoWithDefaults() *AlertingWebhookDto`

NewAlertingWebhookDtoWithDefaults instantiates a new AlertingWebhookDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertingWebhookDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingWebhookDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingWebhookDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertingWebhookDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *AlertingWebhookDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AlertingWebhookDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AlertingWebhookDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AlertingWebhookDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AlertingWebhookDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AlertingWebhookDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetHeaders

`func (o *AlertingWebhookDto) GetHeaders() []WebhookHeaderDto`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AlertingWebhookDto) GetHeadersOk() (*[]WebhookHeaderDto, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AlertingWebhookDto) SetHeaders(v []WebhookHeaderDto)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AlertingWebhookDto) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *AlertingWebhookDto) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *AlertingWebhookDto) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


