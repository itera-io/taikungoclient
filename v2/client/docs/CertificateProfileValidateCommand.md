# CertificateProfileValidateCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateProfileId** | **NullableInt32** |  | 
**Kind** | Pointer to **NullableString** |  | [optional] 
**CertPem** | Pointer to **NullableString** |  | [optional] 
**KeyPem** | Pointer to **NullableString** |  | [optional] 
**ChainPem** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateProfileValidateCommand

`func NewCertificateProfileValidateCommand(certificateProfileId NullableInt32, ) *CertificateProfileValidateCommand`

NewCertificateProfileValidateCommand instantiates a new CertificateProfileValidateCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileValidateCommandWithDefaults

`func NewCertificateProfileValidateCommandWithDefaults() *CertificateProfileValidateCommand`

NewCertificateProfileValidateCommandWithDefaults instantiates a new CertificateProfileValidateCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateProfileId

`func (o *CertificateProfileValidateCommand) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *CertificateProfileValidateCommand) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *CertificateProfileValidateCommand) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.


### SetCertificateProfileIdNil

`func (o *CertificateProfileValidateCommand) SetCertificateProfileIdNil(b bool)`

 SetCertificateProfileIdNil sets the value for CertificateProfileId to be an explicit nil

### UnsetCertificateProfileId
`func (o *CertificateProfileValidateCommand) UnsetCertificateProfileId()`

UnsetCertificateProfileId ensures that no value is present for CertificateProfileId, not even an explicit nil
### GetKind

`func (o *CertificateProfileValidateCommand) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CertificateProfileValidateCommand) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CertificateProfileValidateCommand) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CertificateProfileValidateCommand) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *CertificateProfileValidateCommand) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *CertificateProfileValidateCommand) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetCertPem

`func (o *CertificateProfileValidateCommand) GetCertPem() string`

GetCertPem returns the CertPem field if non-nil, zero value otherwise.

### GetCertPemOk

`func (o *CertificateProfileValidateCommand) GetCertPemOk() (*string, bool)`

GetCertPemOk returns a tuple with the CertPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertPem

`func (o *CertificateProfileValidateCommand) SetCertPem(v string)`

SetCertPem sets CertPem field to given value.

### HasCertPem

`func (o *CertificateProfileValidateCommand) HasCertPem() bool`

HasCertPem returns a boolean if a field has been set.

### SetCertPemNil

`func (o *CertificateProfileValidateCommand) SetCertPemNil(b bool)`

 SetCertPemNil sets the value for CertPem to be an explicit nil

### UnsetCertPem
`func (o *CertificateProfileValidateCommand) UnsetCertPem()`

UnsetCertPem ensures that no value is present for CertPem, not even an explicit nil
### GetKeyPem

`func (o *CertificateProfileValidateCommand) GetKeyPem() string`

GetKeyPem returns the KeyPem field if non-nil, zero value otherwise.

### GetKeyPemOk

`func (o *CertificateProfileValidateCommand) GetKeyPemOk() (*string, bool)`

GetKeyPemOk returns a tuple with the KeyPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPem

`func (o *CertificateProfileValidateCommand) SetKeyPem(v string)`

SetKeyPem sets KeyPem field to given value.

### HasKeyPem

`func (o *CertificateProfileValidateCommand) HasKeyPem() bool`

HasKeyPem returns a boolean if a field has been set.

### SetKeyPemNil

`func (o *CertificateProfileValidateCommand) SetKeyPemNil(b bool)`

 SetKeyPemNil sets the value for KeyPem to be an explicit nil

### UnsetKeyPem
`func (o *CertificateProfileValidateCommand) UnsetKeyPem()`

UnsetKeyPem ensures that no value is present for KeyPem, not even an explicit nil
### GetChainPem

`func (o *CertificateProfileValidateCommand) GetChainPem() string`

GetChainPem returns the ChainPem field if non-nil, zero value otherwise.

### GetChainPemOk

`func (o *CertificateProfileValidateCommand) GetChainPemOk() (*string, bool)`

GetChainPemOk returns a tuple with the ChainPem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainPem

`func (o *CertificateProfileValidateCommand) SetChainPem(v string)`

SetChainPem sets ChainPem field to given value.

### HasChainPem

`func (o *CertificateProfileValidateCommand) HasChainPem() bool`

HasChainPem returns a boolean if a field has been set.

### SetChainPemNil

`func (o *CertificateProfileValidateCommand) SetChainPemNil(b bool)`

 SetChainPemNil sets the value for ChainPem to be an explicit nil

### UnsetChainPem
`func (o *CertificateProfileValidateCommand) UnsetChainPem()`

UnsetChainPem ensures that no value is present for ChainPem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


