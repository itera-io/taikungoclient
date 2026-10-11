# SeedQaTenantResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **NullableString** |  | 
**Template** | **NullableString** |  | 
**DomainId** | **int32** |  | 
**OrganizationId** | **int32** |  | 
**UserId** | **NullableString** |  | 
**DomainName** | **NullableString** |  | 
**CleanupToken** | **NullableString** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**AccessKey** | Pointer to **NullableString** |  | [optional] 
**SecretKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSeedQaTenantResult

`func NewSeedQaTenantResult(runId NullableString, template NullableString, domainId int32, organizationId int32, userId NullableString, domainName NullableString, cleanupToken NullableString, ) *SeedQaTenantResult`

NewSeedQaTenantResult instantiates a new SeedQaTenantResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeedQaTenantResultWithDefaults

`func NewSeedQaTenantResultWithDefaults() *SeedQaTenantResult`

NewSeedQaTenantResultWithDefaults instantiates a new SeedQaTenantResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *SeedQaTenantResult) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *SeedQaTenantResult) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *SeedQaTenantResult) SetRunId(v string)`

SetRunId sets RunId field to given value.


### SetRunIdNil

`func (o *SeedQaTenantResult) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *SeedQaTenantResult) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetTemplate

`func (o *SeedQaTenantResult) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SeedQaTenantResult) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SeedQaTenantResult) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### SetTemplateNil

`func (o *SeedQaTenantResult) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *SeedQaTenantResult) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetDomainId

`func (o *SeedQaTenantResult) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *SeedQaTenantResult) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *SeedQaTenantResult) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### GetOrganizationId

`func (o *SeedQaTenantResult) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SeedQaTenantResult) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SeedQaTenantResult) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *SeedQaTenantResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SeedQaTenantResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SeedQaTenantResult) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *SeedQaTenantResult) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SeedQaTenantResult) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDomainName

`func (o *SeedQaTenantResult) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SeedQaTenantResult) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SeedQaTenantResult) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### SetDomainNameNil

`func (o *SeedQaTenantResult) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *SeedQaTenantResult) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetCleanupToken

`func (o *SeedQaTenantResult) GetCleanupToken() string`

GetCleanupToken returns the CleanupToken field if non-nil, zero value otherwise.

### GetCleanupTokenOk

`func (o *SeedQaTenantResult) GetCleanupTokenOk() (*string, bool)`

GetCleanupTokenOk returns a tuple with the CleanupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupToken

`func (o *SeedQaTenantResult) SetCleanupToken(v string)`

SetCleanupToken sets CleanupToken field to given value.


### SetCleanupTokenNil

`func (o *SeedQaTenantResult) SetCleanupTokenNil(b bool)`

 SetCleanupTokenNil sets the value for CleanupToken to be an explicit nil

### UnsetCleanupToken
`func (o *SeedQaTenantResult) UnsetCleanupToken()`

UnsetCleanupToken ensures that no value is present for CleanupToken, not even an explicit nil
### GetEmail

`func (o *SeedQaTenantResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SeedQaTenantResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SeedQaTenantResult) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SeedQaTenantResult) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SeedQaTenantResult) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SeedQaTenantResult) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *SeedQaTenantResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SeedQaTenantResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SeedQaTenantResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SeedQaTenantResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *SeedQaTenantResult) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SeedQaTenantResult) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetAccessKey

`func (o *SeedQaTenantResult) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SeedQaTenantResult) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SeedQaTenantResult) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *SeedQaTenantResult) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *SeedQaTenantResult) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *SeedQaTenantResult) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetSecretKey

`func (o *SeedQaTenantResult) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SeedQaTenantResult) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SeedQaTenantResult) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SeedQaTenantResult) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *SeedQaTenantResult) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *SeedQaTenantResult) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


