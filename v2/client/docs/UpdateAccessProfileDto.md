# UpdateAccessProfileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**HttpProxy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateAccessProfileDto

`func NewUpdateAccessProfileDto() *UpdateAccessProfileDto`

NewUpdateAccessProfileDto instantiates a new UpdateAccessProfileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessProfileDtoWithDefaults

`func NewUpdateAccessProfileDtoWithDefaults() *UpdateAccessProfileDto`

NewUpdateAccessProfileDtoWithDefaults instantiates a new UpdateAccessProfileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAccessProfileDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccessProfileDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccessProfileDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccessProfileDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateAccessProfileDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAccessProfileDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHttpProxy

`func (o *UpdateAccessProfileDto) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *UpdateAccessProfileDto) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *UpdateAccessProfileDto) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *UpdateAccessProfileDto) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *UpdateAccessProfileDto) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *UpdateAccessProfileDto) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


