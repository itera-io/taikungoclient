# CertificateProfileLockCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Mode** | **NullableString** |  | 

## Methods

### NewCertificateProfileLockCommand

`func NewCertificateProfileLockCommand(id int32, mode NullableString, ) *CertificateProfileLockCommand`

NewCertificateProfileLockCommand instantiates a new CertificateProfileLockCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileLockCommandWithDefaults

`func NewCertificateProfileLockCommandWithDefaults() *CertificateProfileLockCommand`

NewCertificateProfileLockCommandWithDefaults instantiates a new CertificateProfileLockCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileLockCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileLockCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileLockCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetMode

`func (o *CertificateProfileLockCommand) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CertificateProfileLockCommand) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CertificateProfileLockCommand) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *CertificateProfileLockCommand) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *CertificateProfileLockCommand) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


