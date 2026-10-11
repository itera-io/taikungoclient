# CertificateProfileListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Kind** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | **NullableInt32** |  | 
**OrganizationName** | **NullableString** |  | 
**NotBefore** | Pointer to **NullableTime** |  | [optional] 
**NotAfter** | Pointer to **NullableTime** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Issuer** | Pointer to **NullableString** |  | [optional] 
**Sans** | Pointer to **NullableString** |  | [optional] 
**DaysUntilExpiry** | Pointer to **NullableInt32** |  | [optional] 
**ExpiryStatus** | Pointer to **NullableString** |  | [optional] 
**IsLocked** | **bool** |  | 
**IsDefault** | **bool** |  | 
**IsInfra** | **bool** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**InUseCount** | **int32** |  | 
**AssociatedProjects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**IssuerMode** | Pointer to [**CertificateProfileIssuerMode**](CertificateProfileIssuerMode.md) |  | [optional] 
**AcmeEmail** | Pointer to **NullableString** |  | [optional] 
**AcmeServer** | Pointer to **NullableString** |  | [optional] 
**VenafiType** | Pointer to [**CertificateProfileVenafiType**](CertificateProfileVenafiType.md) |  | [optional] 
**VenafiBaseUrl** | Pointer to **NullableString** |  | [optional] 
**VenafiZone** | Pointer to **NullableString** |  | [optional] 
**VenafiZoneUuid** | Pointer to **NullableString** |  | [optional] 
**VenafiAccessToken** | Pointer to **NullableString** | Redacted on list responses (&#x60;****&#x60; when stored). Omit or send a new value on update to rotate. | [optional] 
**VenafiApiKey** | Pointer to **NullableString** | Redacted on list responses (&#x60;****&#x60; when stored). Omit or send a new value on update to rotate. | [optional] 
**CustomCaCert** | Pointer to **NullableString** |  | [optional] 
**CustomCaPrivateKey** | Pointer to **NullableString** | Redacted on list responses (&#x60;****&#x60; when stored). Omit or send a new value on update to rotate. | [optional] 
**TrustBundleCount** | Pointer to **int32** |  | [optional] 
**TrustBundle** | [**[]CertificateProfileTrustBundleSummaryDto**](CertificateProfileTrustBundleSummaryDto.md) |  | 

## Methods

### NewCertificateProfileListDto

`func NewCertificateProfileListDto(id int32, name string, description NullableString, organizationId NullableInt32, organizationName NullableString, isLocked bool, isDefault bool, isInfra bool, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, inUseCount int32, associatedProjects []CommonDropdownDto, trustBundle []CertificateProfileTrustBundleSummaryDto, ) *CertificateProfileListDto`

NewCertificateProfileListDto instantiates a new CertificateProfileListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileListDtoWithDefaults

`func NewCertificateProfileListDtoWithDefaults() *CertificateProfileListDto`

NewCertificateProfileListDtoWithDefaults instantiates a new CertificateProfileListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CertificateProfileListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileListDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CertificateProfileListDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateProfileListDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateProfileListDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *CertificateProfileListDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateProfileListDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *CertificateProfileListDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CertificateProfileListDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CertificateProfileListDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CertificateProfileListDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *CertificateProfileListDto) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *CertificateProfileListDto) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetOrganizationId

`func (o *CertificateProfileListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CertificateProfileListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CertificateProfileListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *CertificateProfileListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CertificateProfileListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *CertificateProfileListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *CertificateProfileListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *CertificateProfileListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *CertificateProfileListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *CertificateProfileListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetNotBefore

`func (o *CertificateProfileListDto) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CertificateProfileListDto) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CertificateProfileListDto) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CertificateProfileListDto) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### SetNotBeforeNil

`func (o *CertificateProfileListDto) SetNotBeforeNil(b bool)`

 SetNotBeforeNil sets the value for NotBefore to be an explicit nil

### UnsetNotBefore
`func (o *CertificateProfileListDto) UnsetNotBefore()`

UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
### GetNotAfter

`func (o *CertificateProfileListDto) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CertificateProfileListDto) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CertificateProfileListDto) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CertificateProfileListDto) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *CertificateProfileListDto) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *CertificateProfileListDto) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
### GetSubject

`func (o *CertificateProfileListDto) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CertificateProfileListDto) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CertificateProfileListDto) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CertificateProfileListDto) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CertificateProfileListDto) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CertificateProfileListDto) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetIssuer

`func (o *CertificateProfileListDto) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CertificateProfileListDto) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CertificateProfileListDto) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *CertificateProfileListDto) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *CertificateProfileListDto) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *CertificateProfileListDto) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetSans

`func (o *CertificateProfileListDto) GetSans() string`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *CertificateProfileListDto) GetSansOk() (*string, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *CertificateProfileListDto) SetSans(v string)`

SetSans sets Sans field to given value.

### HasSans

`func (o *CertificateProfileListDto) HasSans() bool`

HasSans returns a boolean if a field has been set.

### SetSansNil

`func (o *CertificateProfileListDto) SetSansNil(b bool)`

 SetSansNil sets the value for Sans to be an explicit nil

### UnsetSans
`func (o *CertificateProfileListDto) UnsetSans()`

UnsetSans ensures that no value is present for Sans, not even an explicit nil
### GetDaysUntilExpiry

`func (o *CertificateProfileListDto) GetDaysUntilExpiry() int32`

GetDaysUntilExpiry returns the DaysUntilExpiry field if non-nil, zero value otherwise.

### GetDaysUntilExpiryOk

`func (o *CertificateProfileListDto) GetDaysUntilExpiryOk() (*int32, bool)`

GetDaysUntilExpiryOk returns a tuple with the DaysUntilExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilExpiry

`func (o *CertificateProfileListDto) SetDaysUntilExpiry(v int32)`

SetDaysUntilExpiry sets DaysUntilExpiry field to given value.

### HasDaysUntilExpiry

`func (o *CertificateProfileListDto) HasDaysUntilExpiry() bool`

HasDaysUntilExpiry returns a boolean if a field has been set.

### SetDaysUntilExpiryNil

`func (o *CertificateProfileListDto) SetDaysUntilExpiryNil(b bool)`

 SetDaysUntilExpiryNil sets the value for DaysUntilExpiry to be an explicit nil

### UnsetDaysUntilExpiry
`func (o *CertificateProfileListDto) UnsetDaysUntilExpiry()`

UnsetDaysUntilExpiry ensures that no value is present for DaysUntilExpiry, not even an explicit nil
### GetExpiryStatus

`func (o *CertificateProfileListDto) GetExpiryStatus() string`

GetExpiryStatus returns the ExpiryStatus field if non-nil, zero value otherwise.

### GetExpiryStatusOk

`func (o *CertificateProfileListDto) GetExpiryStatusOk() (*string, bool)`

GetExpiryStatusOk returns a tuple with the ExpiryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryStatus

`func (o *CertificateProfileListDto) SetExpiryStatus(v string)`

SetExpiryStatus sets ExpiryStatus field to given value.

### HasExpiryStatus

`func (o *CertificateProfileListDto) HasExpiryStatus() bool`

HasExpiryStatus returns a boolean if a field has been set.

### SetExpiryStatusNil

`func (o *CertificateProfileListDto) SetExpiryStatusNil(b bool)`

 SetExpiryStatusNil sets the value for ExpiryStatus to be an explicit nil

### UnsetExpiryStatus
`func (o *CertificateProfileListDto) UnsetExpiryStatus()`

UnsetExpiryStatus ensures that no value is present for ExpiryStatus, not even an explicit nil
### GetIsLocked

`func (o *CertificateProfileListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *CertificateProfileListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *CertificateProfileListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsDefault

`func (o *CertificateProfileListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CertificateProfileListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CertificateProfileListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsInfra

`func (o *CertificateProfileListDto) GetIsInfra() bool`

GetIsInfra returns the IsInfra field if non-nil, zero value otherwise.

### GetIsInfraOk

`func (o *CertificateProfileListDto) GetIsInfraOk() (*bool, bool)`

GetIsInfraOk returns a tuple with the IsInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfra

`func (o *CertificateProfileListDto) SetIsInfra(v bool)`

SetIsInfra sets IsInfra field to given value.


### GetCreatedBy

`func (o *CertificateProfileListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CertificateProfileListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CertificateProfileListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *CertificateProfileListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CertificateProfileListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CertificateProfileListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *CertificateProfileListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *CertificateProfileListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *CertificateProfileListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *CertificateProfileListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *CertificateProfileListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetInUseCount

`func (o *CertificateProfileListDto) GetInUseCount() int32`

GetInUseCount returns the InUseCount field if non-nil, zero value otherwise.

### GetInUseCountOk

`func (o *CertificateProfileListDto) GetInUseCountOk() (*int32, bool)`

GetInUseCountOk returns a tuple with the InUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUseCount

`func (o *CertificateProfileListDto) SetInUseCount(v int32)`

SetInUseCount sets InUseCount field to given value.


### GetAssociatedProjects

`func (o *CertificateProfileListDto) GetAssociatedProjects() []CommonDropdownDto`

GetAssociatedProjects returns the AssociatedProjects field if non-nil, zero value otherwise.

### GetAssociatedProjectsOk

`func (o *CertificateProfileListDto) GetAssociatedProjectsOk() (*[]CommonDropdownDto, bool)`

GetAssociatedProjectsOk returns a tuple with the AssociatedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedProjects

`func (o *CertificateProfileListDto) SetAssociatedProjects(v []CommonDropdownDto)`

SetAssociatedProjects sets AssociatedProjects field to given value.


### GetIssuerMode

`func (o *CertificateProfileListDto) GetIssuerMode() CertificateProfileIssuerMode`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *CertificateProfileListDto) GetIssuerModeOk() (*CertificateProfileIssuerMode, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *CertificateProfileListDto) SetIssuerMode(v CertificateProfileIssuerMode)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *CertificateProfileListDto) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetAcmeEmail

`func (o *CertificateProfileListDto) GetAcmeEmail() string`

GetAcmeEmail returns the AcmeEmail field if non-nil, zero value otherwise.

### GetAcmeEmailOk

`func (o *CertificateProfileListDto) GetAcmeEmailOk() (*string, bool)`

GetAcmeEmailOk returns a tuple with the AcmeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeEmail

`func (o *CertificateProfileListDto) SetAcmeEmail(v string)`

SetAcmeEmail sets AcmeEmail field to given value.

### HasAcmeEmail

`func (o *CertificateProfileListDto) HasAcmeEmail() bool`

HasAcmeEmail returns a boolean if a field has been set.

### SetAcmeEmailNil

`func (o *CertificateProfileListDto) SetAcmeEmailNil(b bool)`

 SetAcmeEmailNil sets the value for AcmeEmail to be an explicit nil

### UnsetAcmeEmail
`func (o *CertificateProfileListDto) UnsetAcmeEmail()`

UnsetAcmeEmail ensures that no value is present for AcmeEmail, not even an explicit nil
### GetAcmeServer

`func (o *CertificateProfileListDto) GetAcmeServer() string`

GetAcmeServer returns the AcmeServer field if non-nil, zero value otherwise.

### GetAcmeServerOk

`func (o *CertificateProfileListDto) GetAcmeServerOk() (*string, bool)`

GetAcmeServerOk returns a tuple with the AcmeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeServer

`func (o *CertificateProfileListDto) SetAcmeServer(v string)`

SetAcmeServer sets AcmeServer field to given value.

### HasAcmeServer

`func (o *CertificateProfileListDto) HasAcmeServer() bool`

HasAcmeServer returns a boolean if a field has been set.

### SetAcmeServerNil

`func (o *CertificateProfileListDto) SetAcmeServerNil(b bool)`

 SetAcmeServerNil sets the value for AcmeServer to be an explicit nil

### UnsetAcmeServer
`func (o *CertificateProfileListDto) UnsetAcmeServer()`

UnsetAcmeServer ensures that no value is present for AcmeServer, not even an explicit nil
### GetVenafiType

`func (o *CertificateProfileListDto) GetVenafiType() CertificateProfileVenafiType`

GetVenafiType returns the VenafiType field if non-nil, zero value otherwise.

### GetVenafiTypeOk

`func (o *CertificateProfileListDto) GetVenafiTypeOk() (*CertificateProfileVenafiType, bool)`

GetVenafiTypeOk returns a tuple with the VenafiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiType

`func (o *CertificateProfileListDto) SetVenafiType(v CertificateProfileVenafiType)`

SetVenafiType sets VenafiType field to given value.

### HasVenafiType

`func (o *CertificateProfileListDto) HasVenafiType() bool`

HasVenafiType returns a boolean if a field has been set.

### GetVenafiBaseUrl

`func (o *CertificateProfileListDto) GetVenafiBaseUrl() string`

GetVenafiBaseUrl returns the VenafiBaseUrl field if non-nil, zero value otherwise.

### GetVenafiBaseUrlOk

`func (o *CertificateProfileListDto) GetVenafiBaseUrlOk() (*string, bool)`

GetVenafiBaseUrlOk returns a tuple with the VenafiBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiBaseUrl

`func (o *CertificateProfileListDto) SetVenafiBaseUrl(v string)`

SetVenafiBaseUrl sets VenafiBaseUrl field to given value.

### HasVenafiBaseUrl

`func (o *CertificateProfileListDto) HasVenafiBaseUrl() bool`

HasVenafiBaseUrl returns a boolean if a field has been set.

### SetVenafiBaseUrlNil

`func (o *CertificateProfileListDto) SetVenafiBaseUrlNil(b bool)`

 SetVenafiBaseUrlNil sets the value for VenafiBaseUrl to be an explicit nil

### UnsetVenafiBaseUrl
`func (o *CertificateProfileListDto) UnsetVenafiBaseUrl()`

UnsetVenafiBaseUrl ensures that no value is present for VenafiBaseUrl, not even an explicit nil
### GetVenafiZone

`func (o *CertificateProfileListDto) GetVenafiZone() string`

GetVenafiZone returns the VenafiZone field if non-nil, zero value otherwise.

### GetVenafiZoneOk

`func (o *CertificateProfileListDto) GetVenafiZoneOk() (*string, bool)`

GetVenafiZoneOk returns a tuple with the VenafiZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZone

`func (o *CertificateProfileListDto) SetVenafiZone(v string)`

SetVenafiZone sets VenafiZone field to given value.

### HasVenafiZone

`func (o *CertificateProfileListDto) HasVenafiZone() bool`

HasVenafiZone returns a boolean if a field has been set.

### SetVenafiZoneNil

`func (o *CertificateProfileListDto) SetVenafiZoneNil(b bool)`

 SetVenafiZoneNil sets the value for VenafiZone to be an explicit nil

### UnsetVenafiZone
`func (o *CertificateProfileListDto) UnsetVenafiZone()`

UnsetVenafiZone ensures that no value is present for VenafiZone, not even an explicit nil
### GetVenafiZoneUuid

`func (o *CertificateProfileListDto) GetVenafiZoneUuid() string`

GetVenafiZoneUuid returns the VenafiZoneUuid field if non-nil, zero value otherwise.

### GetVenafiZoneUuidOk

`func (o *CertificateProfileListDto) GetVenafiZoneUuidOk() (*string, bool)`

GetVenafiZoneUuidOk returns a tuple with the VenafiZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiZoneUuid

`func (o *CertificateProfileListDto) SetVenafiZoneUuid(v string)`

SetVenafiZoneUuid sets VenafiZoneUuid field to given value.

### HasVenafiZoneUuid

`func (o *CertificateProfileListDto) HasVenafiZoneUuid() bool`

HasVenafiZoneUuid returns a boolean if a field has been set.

### SetVenafiZoneUuidNil

`func (o *CertificateProfileListDto) SetVenafiZoneUuidNil(b bool)`

 SetVenafiZoneUuidNil sets the value for VenafiZoneUuid to be an explicit nil

### UnsetVenafiZoneUuid
`func (o *CertificateProfileListDto) UnsetVenafiZoneUuid()`

UnsetVenafiZoneUuid ensures that no value is present for VenafiZoneUuid, not even an explicit nil
### GetVenafiAccessToken

`func (o *CertificateProfileListDto) GetVenafiAccessToken() string`

GetVenafiAccessToken returns the VenafiAccessToken field if non-nil, zero value otherwise.

### GetVenafiAccessTokenOk

`func (o *CertificateProfileListDto) GetVenafiAccessTokenOk() (*string, bool)`

GetVenafiAccessTokenOk returns a tuple with the VenafiAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiAccessToken

`func (o *CertificateProfileListDto) SetVenafiAccessToken(v string)`

SetVenafiAccessToken sets VenafiAccessToken field to given value.

### HasVenafiAccessToken

`func (o *CertificateProfileListDto) HasVenafiAccessToken() bool`

HasVenafiAccessToken returns a boolean if a field has been set.

### SetVenafiAccessTokenNil

`func (o *CertificateProfileListDto) SetVenafiAccessTokenNil(b bool)`

 SetVenafiAccessTokenNil sets the value for VenafiAccessToken to be an explicit nil

### UnsetVenafiAccessToken
`func (o *CertificateProfileListDto) UnsetVenafiAccessToken()`

UnsetVenafiAccessToken ensures that no value is present for VenafiAccessToken, not even an explicit nil
### GetVenafiApiKey

`func (o *CertificateProfileListDto) GetVenafiApiKey() string`

GetVenafiApiKey returns the VenafiApiKey field if non-nil, zero value otherwise.

### GetVenafiApiKeyOk

`func (o *CertificateProfileListDto) GetVenafiApiKeyOk() (*string, bool)`

GetVenafiApiKeyOk returns a tuple with the VenafiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVenafiApiKey

`func (o *CertificateProfileListDto) SetVenafiApiKey(v string)`

SetVenafiApiKey sets VenafiApiKey field to given value.

### HasVenafiApiKey

`func (o *CertificateProfileListDto) HasVenafiApiKey() bool`

HasVenafiApiKey returns a boolean if a field has been set.

### SetVenafiApiKeyNil

`func (o *CertificateProfileListDto) SetVenafiApiKeyNil(b bool)`

 SetVenafiApiKeyNil sets the value for VenafiApiKey to be an explicit nil

### UnsetVenafiApiKey
`func (o *CertificateProfileListDto) UnsetVenafiApiKey()`

UnsetVenafiApiKey ensures that no value is present for VenafiApiKey, not even an explicit nil
### GetCustomCaCert

`func (o *CertificateProfileListDto) GetCustomCaCert() string`

GetCustomCaCert returns the CustomCaCert field if non-nil, zero value otherwise.

### GetCustomCaCertOk

`func (o *CertificateProfileListDto) GetCustomCaCertOk() (*string, bool)`

GetCustomCaCertOk returns a tuple with the CustomCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCaCert

`func (o *CertificateProfileListDto) SetCustomCaCert(v string)`

SetCustomCaCert sets CustomCaCert field to given value.

### HasCustomCaCert

`func (o *CertificateProfileListDto) HasCustomCaCert() bool`

HasCustomCaCert returns a boolean if a field has been set.

### SetCustomCaCertNil

`func (o *CertificateProfileListDto) SetCustomCaCertNil(b bool)`

 SetCustomCaCertNil sets the value for CustomCaCert to be an explicit nil

### UnsetCustomCaCert
`func (o *CertificateProfileListDto) UnsetCustomCaCert()`

UnsetCustomCaCert ensures that no value is present for CustomCaCert, not even an explicit nil
### GetCustomCaPrivateKey

`func (o *CertificateProfileListDto) GetCustomCaPrivateKey() string`

GetCustomCaPrivateKey returns the CustomCaPrivateKey field if non-nil, zero value otherwise.

### GetCustomCaPrivateKeyOk

`func (o *CertificateProfileListDto) GetCustomCaPrivateKeyOk() (*string, bool)`

GetCustomCaPrivateKeyOk returns a tuple with the CustomCaPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCaPrivateKey

`func (o *CertificateProfileListDto) SetCustomCaPrivateKey(v string)`

SetCustomCaPrivateKey sets CustomCaPrivateKey field to given value.

### HasCustomCaPrivateKey

`func (o *CertificateProfileListDto) HasCustomCaPrivateKey() bool`

HasCustomCaPrivateKey returns a boolean if a field has been set.

### SetCustomCaPrivateKeyNil

`func (o *CertificateProfileListDto) SetCustomCaPrivateKeyNil(b bool)`

 SetCustomCaPrivateKeyNil sets the value for CustomCaPrivateKey to be an explicit nil

### UnsetCustomCaPrivateKey
`func (o *CertificateProfileListDto) UnsetCustomCaPrivateKey()`

UnsetCustomCaPrivateKey ensures that no value is present for CustomCaPrivateKey, not even an explicit nil
### GetTrustBundleCount

`func (o *CertificateProfileListDto) GetTrustBundleCount() int32`

GetTrustBundleCount returns the TrustBundleCount field if non-nil, zero value otherwise.

### GetTrustBundleCountOk

`func (o *CertificateProfileListDto) GetTrustBundleCountOk() (*int32, bool)`

GetTrustBundleCountOk returns a tuple with the TrustBundleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustBundleCount

`func (o *CertificateProfileListDto) SetTrustBundleCount(v int32)`

SetTrustBundleCount sets TrustBundleCount field to given value.

### HasTrustBundleCount

`func (o *CertificateProfileListDto) HasTrustBundleCount() bool`

HasTrustBundleCount returns a boolean if a field has been set.

### GetTrustBundle

`func (o *CertificateProfileListDto) GetTrustBundle() []CertificateProfileTrustBundleSummaryDto`

GetTrustBundle returns the TrustBundle field if non-nil, zero value otherwise.

### GetTrustBundleOk

`func (o *CertificateProfileListDto) GetTrustBundleOk() (*[]CertificateProfileTrustBundleSummaryDto, bool)`

GetTrustBundleOk returns a tuple with the TrustBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustBundle

`func (o *CertificateProfileListDto) SetTrustBundle(v []CertificateProfileTrustBundleSummaryDto)`

SetTrustBundle sets TrustBundle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


