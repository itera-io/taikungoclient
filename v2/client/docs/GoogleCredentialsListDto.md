# GoogleCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**PartnerLogo** | **NullableString** |  | 
**PartnerName** | **NullableString** |  | 
**FolderId** | **NullableString** |  | 
**ProjectId** | **NullableString** |  | 
**BillingAccountId** | **NullableString** |  | 
**Zones** | **[]string** |  | 
**AvailabilityZonesCount** | **int32** |  | 
**Region** | **string** |  | 
**IsLocked** | **bool** |  | 
**IsDefault** | **bool** |  | 
**BillingAccountName** | **NullableString** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**CreatedAt** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**ContinentName** | **NullableString** |  | 
**VpcMode** | Pointer to **map[string]interface{}** |  | [optional] 
**IpMode** | Pointer to [**IpMode**](IpMode.md) |  | [optional] 
**ImportedGcpProjectId** | Pointer to **NullableString** |  | [optional] 
**ImportedVpcName** | Pointer to **NullableString** |  | [optional] 
**SharedFileSystemEnabled** | Pointer to **NullableBool** |  | [optional] 
**ImportedFilestoreId** | Pointer to **NullableString** |  | [optional] 
**ImportedFilestoreName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGoogleCredentialsListDto

`func NewGoogleCredentialsListDto(id int32, name string, projects []CommonDropdownDto, organizationId int32, organizationName string, partnerLogo NullableString, partnerName NullableString, folderId NullableString, projectId NullableString, billingAccountId NullableString, zones []string, availabilityZonesCount int32, region string, isLocked bool, isDefault bool, billingAccountName NullableString, createdBy AuditUserDto, createdAt NullableString, lastModified NullableString, lastModifiedBy AuditUserDto, continentName NullableString, ) *GoogleCredentialsListDto`

NewGoogleCredentialsListDto instantiates a new GoogleCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCredentialsListDtoWithDefaults

`func NewGoogleCredentialsListDtoWithDefaults() *GoogleCredentialsListDto`

NewGoogleCredentialsListDtoWithDefaults instantiates a new GoogleCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GoogleCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GoogleCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GoogleCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GoogleCredentialsListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleCredentialsListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleCredentialsListDto) SetName(v string)`

SetName sets Name field to given value.


### GetProjects

`func (o *GoogleCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GoogleCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GoogleCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetOrganizationId

`func (o *GoogleCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GoogleCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GoogleCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *GoogleCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GoogleCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GoogleCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetPartnerLogo

`func (o *GoogleCredentialsListDto) GetPartnerLogo() string`

GetPartnerLogo returns the PartnerLogo field if non-nil, zero value otherwise.

### GetPartnerLogoOk

`func (o *GoogleCredentialsListDto) GetPartnerLogoOk() (*string, bool)`

GetPartnerLogoOk returns a tuple with the PartnerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLogo

`func (o *GoogleCredentialsListDto) SetPartnerLogo(v string)`

SetPartnerLogo sets PartnerLogo field to given value.


### SetPartnerLogoNil

`func (o *GoogleCredentialsListDto) SetPartnerLogoNil(b bool)`

 SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil

### UnsetPartnerLogo
`func (o *GoogleCredentialsListDto) UnsetPartnerLogo()`

UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
### GetPartnerName

`func (o *GoogleCredentialsListDto) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *GoogleCredentialsListDto) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *GoogleCredentialsListDto) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### SetPartnerNameNil

`func (o *GoogleCredentialsListDto) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *GoogleCredentialsListDto) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetFolderId

`func (o *GoogleCredentialsListDto) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GoogleCredentialsListDto) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GoogleCredentialsListDto) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### SetFolderIdNil

`func (o *GoogleCredentialsListDto) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *GoogleCredentialsListDto) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetProjectId

`func (o *GoogleCredentialsListDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleCredentialsListDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleCredentialsListDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *GoogleCredentialsListDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *GoogleCredentialsListDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetBillingAccountId

`func (o *GoogleCredentialsListDto) GetBillingAccountId() string`

GetBillingAccountId returns the BillingAccountId field if non-nil, zero value otherwise.

### GetBillingAccountIdOk

`func (o *GoogleCredentialsListDto) GetBillingAccountIdOk() (*string, bool)`

GetBillingAccountIdOk returns a tuple with the BillingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountId

`func (o *GoogleCredentialsListDto) SetBillingAccountId(v string)`

SetBillingAccountId sets BillingAccountId field to given value.


### SetBillingAccountIdNil

`func (o *GoogleCredentialsListDto) SetBillingAccountIdNil(b bool)`

 SetBillingAccountIdNil sets the value for BillingAccountId to be an explicit nil

### UnsetBillingAccountId
`func (o *GoogleCredentialsListDto) UnsetBillingAccountId()`

UnsetBillingAccountId ensures that no value is present for BillingAccountId, not even an explicit nil
### GetZones

`func (o *GoogleCredentialsListDto) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *GoogleCredentialsListDto) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *GoogleCredentialsListDto) SetZones(v []string)`

SetZones sets Zones field to given value.


### GetAvailabilityZonesCount

`func (o *GoogleCredentialsListDto) GetAvailabilityZonesCount() int32`

GetAvailabilityZonesCount returns the AvailabilityZonesCount field if non-nil, zero value otherwise.

### GetAvailabilityZonesCountOk

`func (o *GoogleCredentialsListDto) GetAvailabilityZonesCountOk() (*int32, bool)`

GetAvailabilityZonesCountOk returns a tuple with the AvailabilityZonesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZonesCount

`func (o *GoogleCredentialsListDto) SetAvailabilityZonesCount(v int32)`

SetAvailabilityZonesCount sets AvailabilityZonesCount field to given value.


### GetRegion

`func (o *GoogleCredentialsListDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GoogleCredentialsListDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GoogleCredentialsListDto) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetIsLocked

`func (o *GoogleCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GoogleCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GoogleCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsDefault

`func (o *GoogleCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GoogleCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GoogleCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetBillingAccountName

`func (o *GoogleCredentialsListDto) GetBillingAccountName() string`

GetBillingAccountName returns the BillingAccountName field if non-nil, zero value otherwise.

### GetBillingAccountNameOk

`func (o *GoogleCredentialsListDto) GetBillingAccountNameOk() (*string, bool)`

GetBillingAccountNameOk returns a tuple with the BillingAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccountName

`func (o *GoogleCredentialsListDto) SetBillingAccountName(v string)`

SetBillingAccountName sets BillingAccountName field to given value.


### SetBillingAccountNameNil

`func (o *GoogleCredentialsListDto) SetBillingAccountNameNil(b bool)`

 SetBillingAccountNameNil sets the value for BillingAccountName to be an explicit nil

### UnsetBillingAccountName
`func (o *GoogleCredentialsListDto) UnsetBillingAccountName()`

UnsetBillingAccountName ensures that no value is present for BillingAccountName, not even an explicit nil
### GetCreatedBy

`func (o *GoogleCredentialsListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GoogleCredentialsListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GoogleCredentialsListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *GoogleCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GoogleCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GoogleCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GoogleCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GoogleCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastModified

`func (o *GoogleCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *GoogleCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *GoogleCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *GoogleCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *GoogleCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *GoogleCredentialsListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *GoogleCredentialsListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *GoogleCredentialsListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetContinentName

`func (o *GoogleCredentialsListDto) GetContinentName() string`

GetContinentName returns the ContinentName field if non-nil, zero value otherwise.

### GetContinentNameOk

`func (o *GoogleCredentialsListDto) GetContinentNameOk() (*string, bool)`

GetContinentNameOk returns a tuple with the ContinentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinentName

`func (o *GoogleCredentialsListDto) SetContinentName(v string)`

SetContinentName sets ContinentName field to given value.


### SetContinentNameNil

`func (o *GoogleCredentialsListDto) SetContinentNameNil(b bool)`

 SetContinentNameNil sets the value for ContinentName to be an explicit nil

### UnsetContinentName
`func (o *GoogleCredentialsListDto) UnsetContinentName()`

UnsetContinentName ensures that no value is present for ContinentName, not even an explicit nil
### GetVpcMode

`func (o *GoogleCredentialsListDto) GetVpcMode() map[string]interface{}`

GetVpcMode returns the VpcMode field if non-nil, zero value otherwise.

### GetVpcModeOk

`func (o *GoogleCredentialsListDto) GetVpcModeOk() (*map[string]interface{}, bool)`

GetVpcModeOk returns a tuple with the VpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMode

`func (o *GoogleCredentialsListDto) SetVpcMode(v map[string]interface{})`

SetVpcMode sets VpcMode field to given value.

### HasVpcMode

`func (o *GoogleCredentialsListDto) HasVpcMode() bool`

HasVpcMode returns a boolean if a field has been set.

### GetIpMode

`func (o *GoogleCredentialsListDto) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *GoogleCredentialsListDto) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *GoogleCredentialsListDto) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *GoogleCredentialsListDto) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetImportedGcpProjectId

`func (o *GoogleCredentialsListDto) GetImportedGcpProjectId() string`

GetImportedGcpProjectId returns the ImportedGcpProjectId field if non-nil, zero value otherwise.

### GetImportedGcpProjectIdOk

`func (o *GoogleCredentialsListDto) GetImportedGcpProjectIdOk() (*string, bool)`

GetImportedGcpProjectIdOk returns a tuple with the ImportedGcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedGcpProjectId

`func (o *GoogleCredentialsListDto) SetImportedGcpProjectId(v string)`

SetImportedGcpProjectId sets ImportedGcpProjectId field to given value.

### HasImportedGcpProjectId

`func (o *GoogleCredentialsListDto) HasImportedGcpProjectId() bool`

HasImportedGcpProjectId returns a boolean if a field has been set.

### SetImportedGcpProjectIdNil

`func (o *GoogleCredentialsListDto) SetImportedGcpProjectIdNil(b bool)`

 SetImportedGcpProjectIdNil sets the value for ImportedGcpProjectId to be an explicit nil

### UnsetImportedGcpProjectId
`func (o *GoogleCredentialsListDto) UnsetImportedGcpProjectId()`

UnsetImportedGcpProjectId ensures that no value is present for ImportedGcpProjectId, not even an explicit nil
### GetImportedVpcName

`func (o *GoogleCredentialsListDto) GetImportedVpcName() string`

GetImportedVpcName returns the ImportedVpcName field if non-nil, zero value otherwise.

### GetImportedVpcNameOk

`func (o *GoogleCredentialsListDto) GetImportedVpcNameOk() (*string, bool)`

GetImportedVpcNameOk returns a tuple with the ImportedVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedVpcName

`func (o *GoogleCredentialsListDto) SetImportedVpcName(v string)`

SetImportedVpcName sets ImportedVpcName field to given value.

### HasImportedVpcName

`func (o *GoogleCredentialsListDto) HasImportedVpcName() bool`

HasImportedVpcName returns a boolean if a field has been set.

### SetImportedVpcNameNil

`func (o *GoogleCredentialsListDto) SetImportedVpcNameNil(b bool)`

 SetImportedVpcNameNil sets the value for ImportedVpcName to be an explicit nil

### UnsetImportedVpcName
`func (o *GoogleCredentialsListDto) UnsetImportedVpcName()`

UnsetImportedVpcName ensures that no value is present for ImportedVpcName, not even an explicit nil
### GetSharedFileSystemEnabled

`func (o *GoogleCredentialsListDto) GetSharedFileSystemEnabled() bool`

GetSharedFileSystemEnabled returns the SharedFileSystemEnabled field if non-nil, zero value otherwise.

### GetSharedFileSystemEnabledOk

`func (o *GoogleCredentialsListDto) GetSharedFileSystemEnabledOk() (*bool, bool)`

GetSharedFileSystemEnabledOk returns a tuple with the SharedFileSystemEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFileSystemEnabled

`func (o *GoogleCredentialsListDto) SetSharedFileSystemEnabled(v bool)`

SetSharedFileSystemEnabled sets SharedFileSystemEnabled field to given value.

### HasSharedFileSystemEnabled

`func (o *GoogleCredentialsListDto) HasSharedFileSystemEnabled() bool`

HasSharedFileSystemEnabled returns a boolean if a field has been set.

### SetSharedFileSystemEnabledNil

`func (o *GoogleCredentialsListDto) SetSharedFileSystemEnabledNil(b bool)`

 SetSharedFileSystemEnabledNil sets the value for SharedFileSystemEnabled to be an explicit nil

### UnsetSharedFileSystemEnabled
`func (o *GoogleCredentialsListDto) UnsetSharedFileSystemEnabled()`

UnsetSharedFileSystemEnabled ensures that no value is present for SharedFileSystemEnabled, not even an explicit nil
### GetImportedFilestoreId

`func (o *GoogleCredentialsListDto) GetImportedFilestoreId() string`

GetImportedFilestoreId returns the ImportedFilestoreId field if non-nil, zero value otherwise.

### GetImportedFilestoreIdOk

`func (o *GoogleCredentialsListDto) GetImportedFilestoreIdOk() (*string, bool)`

GetImportedFilestoreIdOk returns a tuple with the ImportedFilestoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFilestoreId

`func (o *GoogleCredentialsListDto) SetImportedFilestoreId(v string)`

SetImportedFilestoreId sets ImportedFilestoreId field to given value.

### HasImportedFilestoreId

`func (o *GoogleCredentialsListDto) HasImportedFilestoreId() bool`

HasImportedFilestoreId returns a boolean if a field has been set.

### SetImportedFilestoreIdNil

`func (o *GoogleCredentialsListDto) SetImportedFilestoreIdNil(b bool)`

 SetImportedFilestoreIdNil sets the value for ImportedFilestoreId to be an explicit nil

### UnsetImportedFilestoreId
`func (o *GoogleCredentialsListDto) UnsetImportedFilestoreId()`

UnsetImportedFilestoreId ensures that no value is present for ImportedFilestoreId, not even an explicit nil
### GetImportedFilestoreName

`func (o *GoogleCredentialsListDto) GetImportedFilestoreName() string`

GetImportedFilestoreName returns the ImportedFilestoreName field if non-nil, zero value otherwise.

### GetImportedFilestoreNameOk

`func (o *GoogleCredentialsListDto) GetImportedFilestoreNameOk() (*string, bool)`

GetImportedFilestoreNameOk returns a tuple with the ImportedFilestoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFilestoreName

`func (o *GoogleCredentialsListDto) SetImportedFilestoreName(v string)`

SetImportedFilestoreName sets ImportedFilestoreName field to given value.

### HasImportedFilestoreName

`func (o *GoogleCredentialsListDto) HasImportedFilestoreName() bool`

HasImportedFilestoreName returns a boolean if a field has been set.

### SetImportedFilestoreNameNil

`func (o *GoogleCredentialsListDto) SetImportedFilestoreNameNil(b bool)`

 SetImportedFilestoreNameNil sets the value for ImportedFilestoreName to be an explicit nil

### UnsetImportedFilestoreName
`func (o *GoogleCredentialsListDto) UnsetImportedFilestoreName()`

UnsetImportedFilestoreName ensures that no value is present for ImportedFilestoreName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


