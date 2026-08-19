# WebhookHeaderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWebhookHeaderDto

`func NewWebhookHeaderDto() *WebhookHeaderDto`

NewWebhookHeaderDto instantiates a new WebhookHeaderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookHeaderDtoWithDefaults

`func NewWebhookHeaderDtoWithDefaults() *WebhookHeaderDto`

NewWebhookHeaderDtoWithDefaults instantiates a new WebhookHeaderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookHeaderDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookHeaderDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookHeaderDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookHeaderDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *WebhookHeaderDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookHeaderDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookHeaderDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WebhookHeaderDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *WebhookHeaderDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *WebhookHeaderDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *WebhookHeaderDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookHeaderDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookHeaderDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WebhookHeaderDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *WebhookHeaderDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WebhookHeaderDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


