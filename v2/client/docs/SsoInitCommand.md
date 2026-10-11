# SsoInitCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **NullableString** |  | [optional] 
**UiRedirectHost** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSsoInitCommand

`func NewSsoInitCommand() *SsoInitCommand`

NewSsoInitCommand instantiates a new SsoInitCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoInitCommandWithDefaults

`func NewSsoInitCommandWithDefaults() *SsoInitCommand`

NewSsoInitCommandWithDefaults instantiates a new SsoInitCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *SsoInitCommand) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SsoInitCommand) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SsoInitCommand) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SsoInitCommand) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *SsoInitCommand) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *SsoInitCommand) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetUiRedirectHost

`func (o *SsoInitCommand) GetUiRedirectHost() string`

GetUiRedirectHost returns the UiRedirectHost field if non-nil, zero value otherwise.

### GetUiRedirectHostOk

`func (o *SsoInitCommand) GetUiRedirectHostOk() (*string, bool)`

GetUiRedirectHostOk returns a tuple with the UiRedirectHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiRedirectHost

`func (o *SsoInitCommand) SetUiRedirectHost(v string)`

SetUiRedirectHost sets UiRedirectHost field to given value.

### HasUiRedirectHost

`func (o *SsoInitCommand) HasUiRedirectHost() bool`

HasUiRedirectHost returns a boolean if a field has been set.

### SetUiRedirectHostNil

`func (o *SsoInitCommand) SetUiRedirectHostNil(b bool)`

 SetUiRedirectHostNil sets the value for UiRedirectHost to be an explicit nil

### UnsetUiRedirectHost
`func (o *SsoInitCommand) UnsetUiRedirectHost()`

UnsetUiRedirectHost ensures that no value is present for UiRedirectHost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


