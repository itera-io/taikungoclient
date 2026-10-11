# BackupCredentialsListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**S3Name** | **string** |  | 
**S3AccessKeyId** | **string** |  | 
**S3Endpoint** | **string** |  | 
**S3Region** | **string** |  | 
**OrganizationId** | **NullableInt32** |  | 
**OrganizationName** | **NullableString** |  | 
**Projects** | [**[]CommonDropdownDto**](CommonDropdownDto.md) |  | 
**IsLocked** | **bool** |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**CreatedAt** | **NullableString** |  | 
**IsDefault** | **bool** |  | 
**IsInfra** | **bool** |  | 

## Methods

### NewBackupCredentialsListDto

`func NewBackupCredentialsListDto(id int32, s3Name string, s3AccessKeyId string, s3Endpoint string, s3Region string, organizationId NullableInt32, organizationName NullableString, projects []CommonDropdownDto, isLocked bool, createdBy AuditUserDto, lastModified NullableString, lastModifiedBy AuditUserDto, createdAt NullableString, isDefault bool, isInfra bool, ) *BackupCredentialsListDto`

NewBackupCredentialsListDto instantiates a new BackupCredentialsListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCredentialsListDtoWithDefaults

`func NewBackupCredentialsListDtoWithDefaults() *BackupCredentialsListDto`

NewBackupCredentialsListDtoWithDefaults instantiates a new BackupCredentialsListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupCredentialsListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupCredentialsListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupCredentialsListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetS3Name

`func (o *BackupCredentialsListDto) GetS3Name() string`

GetS3Name returns the S3Name field if non-nil, zero value otherwise.

### GetS3NameOk

`func (o *BackupCredentialsListDto) GetS3NameOk() (*string, bool)`

GetS3NameOk returns a tuple with the S3Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Name

`func (o *BackupCredentialsListDto) SetS3Name(v string)`

SetS3Name sets S3Name field to given value.


### GetS3AccessKeyId

`func (o *BackupCredentialsListDto) GetS3AccessKeyId() string`

GetS3AccessKeyId returns the S3AccessKeyId field if non-nil, zero value otherwise.

### GetS3AccessKeyIdOk

`func (o *BackupCredentialsListDto) GetS3AccessKeyIdOk() (*string, bool)`

GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3AccessKeyId

`func (o *BackupCredentialsListDto) SetS3AccessKeyId(v string)`

SetS3AccessKeyId sets S3AccessKeyId field to given value.


### GetS3Endpoint

`func (o *BackupCredentialsListDto) GetS3Endpoint() string`

GetS3Endpoint returns the S3Endpoint field if non-nil, zero value otherwise.

### GetS3EndpointOk

`func (o *BackupCredentialsListDto) GetS3EndpointOk() (*string, bool)`

GetS3EndpointOk returns a tuple with the S3Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Endpoint

`func (o *BackupCredentialsListDto) SetS3Endpoint(v string)`

SetS3Endpoint sets S3Endpoint field to given value.


### GetS3Region

`func (o *BackupCredentialsListDto) GetS3Region() string`

GetS3Region returns the S3Region field if non-nil, zero value otherwise.

### GetS3RegionOk

`func (o *BackupCredentialsListDto) GetS3RegionOk() (*string, bool)`

GetS3RegionOk returns a tuple with the S3Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Region

`func (o *BackupCredentialsListDto) SetS3Region(v string)`

SetS3Region sets S3Region field to given value.


### GetOrganizationId

`func (o *BackupCredentialsListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BackupCredentialsListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BackupCredentialsListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *BackupCredentialsListDto) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *BackupCredentialsListDto) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetOrganizationName

`func (o *BackupCredentialsListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *BackupCredentialsListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *BackupCredentialsListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *BackupCredentialsListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *BackupCredentialsListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetProjects

`func (o *BackupCredentialsListDto) GetProjects() []CommonDropdownDto`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BackupCredentialsListDto) GetProjectsOk() (*[]CommonDropdownDto, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BackupCredentialsListDto) SetProjects(v []CommonDropdownDto)`

SetProjects sets Projects field to given value.


### GetIsLocked

`func (o *BackupCredentialsListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *BackupCredentialsListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *BackupCredentialsListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetCreatedBy

`func (o *BackupCredentialsListDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupCredentialsListDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupCredentialsListDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *BackupCredentialsListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *BackupCredentialsListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *BackupCredentialsListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *BackupCredentialsListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *BackupCredentialsListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *BackupCredentialsListDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *BackupCredentialsListDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *BackupCredentialsListDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetCreatedAt

`func (o *BackupCredentialsListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupCredentialsListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupCredentialsListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *BackupCredentialsListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BackupCredentialsListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsDefault

`func (o *BackupCredentialsListDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BackupCredentialsListDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BackupCredentialsListDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetIsInfra

`func (o *BackupCredentialsListDto) GetIsInfra() bool`

GetIsInfra returns the IsInfra field if non-nil, zero value otherwise.

### GetIsInfraOk

`func (o *BackupCredentialsListDto) GetIsInfraOk() (*bool, bool)`

GetIsInfraOk returns a tuple with the IsInfra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfra

`func (o *BackupCredentialsListDto) SetIsInfra(v bool)`

SetIsInfra sets IsInfra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


