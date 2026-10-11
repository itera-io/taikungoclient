# ImportedClusterDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**IsLocked** | **bool** |  | 
**AccessIp** | **NullableString** |  | 
**KubernetesVersion** | **NullableString** |  | 
**ImportClusterType** | [**ImportClusterType**](ImportClusterType.md) |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**CloudCredentialName** | Pointer to **NullableString** |  | [optional] 
**CloudCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**IsMonitoringEnabled** | **bool** |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**AlertingProfileName** | **NullableString** |  | 
**IsOpaEnabled** | **bool** |  | 
**OpaProfileId** | **NullableInt32** |  | 
**IsBackupEnabled** | **bool** |  | 
**S3CredentialId** | **NullableInt32** |  | 
**AiEnabled** | **bool** |  | 
**AiCredentialId** | **NullableInt32** |  | 
**ExpiredAt** | **string** |  | 
**AlertsCount** | **int32** |  | 

## Methods

### NewImportedClusterDetailsDto

`func NewImportedClusterDetailsDto(id int32, name string, isLocked bool, accessIp NullableString, kubernetesVersion NullableString, importClusterType ImportClusterType, organizationId int32, organizationName string, health ProjectHealth, cloudType ECloudCredentialType, status ProjectStatus, isMonitoringEnabled bool, alertingProfileId NullableInt32, alertingProfileName NullableString, isOpaEnabled bool, opaProfileId NullableInt32, isBackupEnabled bool, s3CredentialId NullableInt32, aiEnabled bool, aiCredentialId NullableInt32, expiredAt string, alertsCount int32, ) *ImportedClusterDetailsDto`

NewImportedClusterDetailsDto instantiates a new ImportedClusterDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportedClusterDetailsDtoWithDefaults

`func NewImportedClusterDetailsDtoWithDefaults() *ImportedClusterDetailsDto`

NewImportedClusterDetailsDtoWithDefaults instantiates a new ImportedClusterDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportedClusterDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportedClusterDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportedClusterDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ImportedClusterDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportedClusterDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportedClusterDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsLocked

`func (o *ImportedClusterDetailsDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ImportedClusterDetailsDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ImportedClusterDetailsDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetAccessIp

`func (o *ImportedClusterDetailsDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ImportedClusterDetailsDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ImportedClusterDetailsDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.


### SetAccessIpNil

`func (o *ImportedClusterDetailsDto) SetAccessIpNil(b bool)`

 SetAccessIpNil sets the value for AccessIp to be an explicit nil

### UnsetAccessIp
`func (o *ImportedClusterDetailsDto) UnsetAccessIp()`

UnsetAccessIp ensures that no value is present for AccessIp, not even an explicit nil
### GetKubernetesVersion

`func (o *ImportedClusterDetailsDto) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ImportedClusterDetailsDto) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ImportedClusterDetailsDto) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *ImportedClusterDetailsDto) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *ImportedClusterDetailsDto) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetImportClusterType

`func (o *ImportedClusterDetailsDto) GetImportClusterType() ImportClusterType`

GetImportClusterType returns the ImportClusterType field if non-nil, zero value otherwise.

### GetImportClusterTypeOk

`func (o *ImportedClusterDetailsDto) GetImportClusterTypeOk() (*ImportClusterType, bool)`

GetImportClusterTypeOk returns a tuple with the ImportClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportClusterType

`func (o *ImportedClusterDetailsDto) SetImportClusterType(v ImportClusterType)`

SetImportClusterType sets ImportClusterType field to given value.


### GetOrganizationId

`func (o *ImportedClusterDetailsDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ImportedClusterDetailsDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ImportedClusterDetailsDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *ImportedClusterDetailsDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ImportedClusterDetailsDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ImportedClusterDetailsDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetCloudCredentialName

`func (o *ImportedClusterDetailsDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *ImportedClusterDetailsDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *ImportedClusterDetailsDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *ImportedClusterDetailsDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### SetCloudCredentialNameNil

`func (o *ImportedClusterDetailsDto) SetCloudCredentialNameNil(b bool)`

 SetCloudCredentialNameNil sets the value for CloudCredentialName to be an explicit nil

### UnsetCloudCredentialName
`func (o *ImportedClusterDetailsDto) UnsetCloudCredentialName()`

UnsetCloudCredentialName ensures that no value is present for CloudCredentialName, not even an explicit nil
### GetCloudCredentialId

`func (o *ImportedClusterDetailsDto) GetCloudCredentialId() int32`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *ImportedClusterDetailsDto) GetCloudCredentialIdOk() (*int32, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *ImportedClusterDetailsDto) SetCloudCredentialId(v int32)`

SetCloudCredentialId sets CloudCredentialId field to given value.

### HasCloudCredentialId

`func (o *ImportedClusterDetailsDto) HasCloudCredentialId() bool`

HasCloudCredentialId returns a boolean if a field has been set.

### SetCloudCredentialIdNil

`func (o *ImportedClusterDetailsDto) SetCloudCredentialIdNil(b bool)`

 SetCloudCredentialIdNil sets the value for CloudCredentialId to be an explicit nil

### UnsetCloudCredentialId
`func (o *ImportedClusterDetailsDto) UnsetCloudCredentialId()`

UnsetCloudCredentialId ensures that no value is present for CloudCredentialId, not even an explicit nil
### GetHealth

`func (o *ImportedClusterDetailsDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ImportedClusterDetailsDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ImportedClusterDetailsDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetCloudType

`func (o *ImportedClusterDetailsDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ImportedClusterDetailsDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ImportedClusterDetailsDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetStatus

`func (o *ImportedClusterDetailsDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportedClusterDetailsDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportedClusterDetailsDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetIsMonitoringEnabled

`func (o *ImportedClusterDetailsDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ImportedClusterDetailsDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ImportedClusterDetailsDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.


### GetAlertingProfileId

`func (o *ImportedClusterDetailsDto) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *ImportedClusterDetailsDto) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *ImportedClusterDetailsDto) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *ImportedClusterDetailsDto) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *ImportedClusterDetailsDto) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetAlertingProfileName

`func (o *ImportedClusterDetailsDto) GetAlertingProfileName() string`

GetAlertingProfileName returns the AlertingProfileName field if non-nil, zero value otherwise.

### GetAlertingProfileNameOk

`func (o *ImportedClusterDetailsDto) GetAlertingProfileNameOk() (*string, bool)`

GetAlertingProfileNameOk returns a tuple with the AlertingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileName

`func (o *ImportedClusterDetailsDto) SetAlertingProfileName(v string)`

SetAlertingProfileName sets AlertingProfileName field to given value.


### SetAlertingProfileNameNil

`func (o *ImportedClusterDetailsDto) SetAlertingProfileNameNil(b bool)`

 SetAlertingProfileNameNil sets the value for AlertingProfileName to be an explicit nil

### UnsetAlertingProfileName
`func (o *ImportedClusterDetailsDto) UnsetAlertingProfileName()`

UnsetAlertingProfileName ensures that no value is present for AlertingProfileName, not even an explicit nil
### GetIsOpaEnabled

`func (o *ImportedClusterDetailsDto) GetIsOpaEnabled() bool`

GetIsOpaEnabled returns the IsOpaEnabled field if non-nil, zero value otherwise.

### GetIsOpaEnabledOk

`func (o *ImportedClusterDetailsDto) GetIsOpaEnabledOk() (*bool, bool)`

GetIsOpaEnabledOk returns a tuple with the IsOpaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpaEnabled

`func (o *ImportedClusterDetailsDto) SetIsOpaEnabled(v bool)`

SetIsOpaEnabled sets IsOpaEnabled field to given value.


### GetOpaProfileId

`func (o *ImportedClusterDetailsDto) GetOpaProfileId() int32`

GetOpaProfileId returns the OpaProfileId field if non-nil, zero value otherwise.

### GetOpaProfileIdOk

`func (o *ImportedClusterDetailsDto) GetOpaProfileIdOk() (*int32, bool)`

GetOpaProfileIdOk returns a tuple with the OpaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileId

`func (o *ImportedClusterDetailsDto) SetOpaProfileId(v int32)`

SetOpaProfileId sets OpaProfileId field to given value.


### SetOpaProfileIdNil

`func (o *ImportedClusterDetailsDto) SetOpaProfileIdNil(b bool)`

 SetOpaProfileIdNil sets the value for OpaProfileId to be an explicit nil

### UnsetOpaProfileId
`func (o *ImportedClusterDetailsDto) UnsetOpaProfileId()`

UnsetOpaProfileId ensures that no value is present for OpaProfileId, not even an explicit nil
### GetIsBackupEnabled

`func (o *ImportedClusterDetailsDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ImportedClusterDetailsDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ImportedClusterDetailsDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.


### GetS3CredentialId

`func (o *ImportedClusterDetailsDto) GetS3CredentialId() int32`

GetS3CredentialId returns the S3CredentialId field if non-nil, zero value otherwise.

### GetS3CredentialIdOk

`func (o *ImportedClusterDetailsDto) GetS3CredentialIdOk() (*int32, bool)`

GetS3CredentialIdOk returns a tuple with the S3CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CredentialId

`func (o *ImportedClusterDetailsDto) SetS3CredentialId(v int32)`

SetS3CredentialId sets S3CredentialId field to given value.


### SetS3CredentialIdNil

`func (o *ImportedClusterDetailsDto) SetS3CredentialIdNil(b bool)`

 SetS3CredentialIdNil sets the value for S3CredentialId to be an explicit nil

### UnsetS3CredentialId
`func (o *ImportedClusterDetailsDto) UnsetS3CredentialId()`

UnsetS3CredentialId ensures that no value is present for S3CredentialId, not even an explicit nil
### GetAiEnabled

`func (o *ImportedClusterDetailsDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *ImportedClusterDetailsDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *ImportedClusterDetailsDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.


### GetAiCredentialId

`func (o *ImportedClusterDetailsDto) GetAiCredentialId() int32`

GetAiCredentialId returns the AiCredentialId field if non-nil, zero value otherwise.

### GetAiCredentialIdOk

`func (o *ImportedClusterDetailsDto) GetAiCredentialIdOk() (*int32, bool)`

GetAiCredentialIdOk returns a tuple with the AiCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiCredentialId

`func (o *ImportedClusterDetailsDto) SetAiCredentialId(v int32)`

SetAiCredentialId sets AiCredentialId field to given value.


### SetAiCredentialIdNil

`func (o *ImportedClusterDetailsDto) SetAiCredentialIdNil(b bool)`

 SetAiCredentialIdNil sets the value for AiCredentialId to be an explicit nil

### UnsetAiCredentialId
`func (o *ImportedClusterDetailsDto) UnsetAiCredentialId()`

UnsetAiCredentialId ensures that no value is present for AiCredentialId, not even an explicit nil
### GetExpiredAt

`func (o *ImportedClusterDetailsDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ImportedClusterDetailsDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ImportedClusterDetailsDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### GetAlertsCount

`func (o *ImportedClusterDetailsDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *ImportedClusterDetailsDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *ImportedClusterDetailsDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


