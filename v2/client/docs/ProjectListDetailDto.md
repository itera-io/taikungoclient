# ProjectListDetailDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**IsKubernetes** | **bool** |  | 
**IsLocked** | **bool** |  | 
**IsVirtualCluster** | **bool** |  | 
**IsMonitoringEnabled** | **bool** |  | 
**HasKubeConfigFile** | **bool** |  | 
**IsMaintenanceModeEnabled** | **bool** |  | 
**CloudCredentialName** | Pointer to **NullableString** |  | [optional] 
**OrganizationName** | **string** |  | 
**OrganizationId** | **int32** |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**KubesprayCurrentVersion** | **string** |  | 
**KubesprayTargetVersion** | **string** |  | 
**KubernetesCurrentVersion** | **string** |  | 
**KubernetesTargetVersion** | **string** |  | 
**CreatedAt** | **string** |  | 
**AlertsCount** | **int32** |  | 
**TotalServersCount** | **int32** |  | 
**TotalStandaloneVmsCount** | **int32** |  | 
**BoundUsers** | [**[]UserDto**](UserDto.md) |  | 
**CreatedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**LastModified** | **NullableString** |  | 
**ExpiredAt** | **NullableString** |  | 
**DeleteOnExpiration** | **bool** |  | 
**CertificateExpiredAt** | **NullableString** |  | 
**LastModifiedBy** | [**AuditUserDto**](AuditUserDto.md) |  | 
**QuotaId** | **int32** |  | 
**AllowFullSpotKubernetes** | **bool** |  | 
**AllowSpotWorkers** | **bool** |  | 
**AllowSpotVMs** | **bool** |  | 
**MaxSpotPrice** | **NullableFloat64** |  | 
**ProjectAction** | **bool** |  | 
**HasExpirationWarning** | **bool** |  | 
**IsAutoscalingEnabled** | **bool** |  | 
**IsAutoscalingSpotEnabled** | **bool** |  | 
**GatewayApiModern** | **bool** |  | 
**AiEnabled** | **bool** |  | 
**AnyServer** | **bool** |  | 
**AnyVm** | **bool** |  | 
**IsBackupEnabled** | **bool** |  | 
**IsProjectMaintenanceModeEnabled** | **bool** |  | 
**ParentProjectId** | **NullableInt32** |  | 
**AlertingProfileId** | **NullableInt32** |  | 
**OpaProfileId** | **NullableInt32** |  | 
**LockButton** | Pointer to [**ButtonStatusDto**](ButtonStatusDto.md) |  | [optional] 
**UnlockButton** | Pointer to [**ButtonStatusDto**](ButtonStatusDto.md) |  | [optional] 
**DeleteButton** | Pointer to [**ButtonStatusDto**](ButtonStatusDto.md) |  | [optional] 
**KubeInfoButton** | Pointer to [**ButtonStatusDto**](ButtonStatusDto.md) |  | [optional] 
**SetExpirationDateButton** | Pointer to [**ButtonStatusDto**](ButtonStatusDto.md) |  | [optional] 
**ResetStatusButton** | Pointer to [**ButtonStatusDto**](ButtonStatusDto.md) |  | [optional] 
**ImportClusterType** | [**ImportClusterType**](ImportClusterType.md) |  | 
**EImportClusterType** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProjectListDetailDto

`func NewProjectListDetailDto(id int32, name string, isKubernetes bool, isLocked bool, isVirtualCluster bool, isMonitoringEnabled bool, hasKubeConfigFile bool, isMaintenanceModeEnabled bool, organizationName string, organizationId int32, status ProjectStatus, health ProjectHealth, cloudType ECloudCredentialType, kubesprayCurrentVersion string, kubesprayTargetVersion string, kubernetesCurrentVersion string, kubernetesTargetVersion string, createdAt string, alertsCount int32, totalServersCount int32, totalStandaloneVmsCount int32, boundUsers []UserDto, createdBy AuditUserDto, lastModified NullableString, expiredAt NullableString, deleteOnExpiration bool, certificateExpiredAt NullableString, lastModifiedBy AuditUserDto, quotaId int32, allowFullSpotKubernetes bool, allowSpotWorkers bool, allowSpotVMs bool, maxSpotPrice NullableFloat64, projectAction bool, hasExpirationWarning bool, isAutoscalingEnabled bool, isAutoscalingSpotEnabled bool, gatewayApiModern bool, aiEnabled bool, anyServer bool, anyVm bool, isBackupEnabled bool, isProjectMaintenanceModeEnabled bool, parentProjectId NullableInt32, alertingProfileId NullableInt32, opaProfileId NullableInt32, importClusterType ImportClusterType, ) *ProjectListDetailDto`

NewProjectListDetailDto instantiates a new ProjectListDetailDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListDetailDtoWithDefaults

`func NewProjectListDetailDtoWithDefaults() *ProjectListDetailDto`

NewProjectListDetailDtoWithDefaults instantiates a new ProjectListDetailDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectListDetailDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectListDetailDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectListDetailDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectListDetailDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectListDetailDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectListDetailDto) SetName(v string)`

SetName sets Name field to given value.


### GetIsKubernetes

`func (o *ProjectListDetailDto) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *ProjectListDetailDto) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *ProjectListDetailDto) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.


### GetIsLocked

`func (o *ProjectListDetailDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ProjectListDetailDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ProjectListDetailDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetIsVirtualCluster

`func (o *ProjectListDetailDto) GetIsVirtualCluster() bool`

GetIsVirtualCluster returns the IsVirtualCluster field if non-nil, zero value otherwise.

### GetIsVirtualClusterOk

`func (o *ProjectListDetailDto) GetIsVirtualClusterOk() (*bool, bool)`

GetIsVirtualClusterOk returns a tuple with the IsVirtualCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualCluster

`func (o *ProjectListDetailDto) SetIsVirtualCluster(v bool)`

SetIsVirtualCluster sets IsVirtualCluster field to given value.


### GetIsMonitoringEnabled

`func (o *ProjectListDetailDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectListDetailDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectListDetailDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.


### GetHasKubeConfigFile

`func (o *ProjectListDetailDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectListDetailDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectListDetailDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.


### GetIsMaintenanceModeEnabled

`func (o *ProjectListDetailDto) GetIsMaintenanceModeEnabled() bool`

GetIsMaintenanceModeEnabled returns the IsMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsMaintenanceModeEnabledOk

`func (o *ProjectListDetailDto) GetIsMaintenanceModeEnabledOk() (*bool, bool)`

GetIsMaintenanceModeEnabledOk returns a tuple with the IsMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaintenanceModeEnabled

`func (o *ProjectListDetailDto) SetIsMaintenanceModeEnabled(v bool)`

SetIsMaintenanceModeEnabled sets IsMaintenanceModeEnabled field to given value.


### GetCloudCredentialName

`func (o *ProjectListDetailDto) GetCloudCredentialName() string`

GetCloudCredentialName returns the CloudCredentialName field if non-nil, zero value otherwise.

### GetCloudCredentialNameOk

`func (o *ProjectListDetailDto) GetCloudCredentialNameOk() (*string, bool)`

GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialName

`func (o *ProjectListDetailDto) SetCloudCredentialName(v string)`

SetCloudCredentialName sets CloudCredentialName field to given value.

### HasCloudCredentialName

`func (o *ProjectListDetailDto) HasCloudCredentialName() bool`

HasCloudCredentialName returns a boolean if a field has been set.

### SetCloudCredentialNameNil

`func (o *ProjectListDetailDto) SetCloudCredentialNameNil(b bool)`

 SetCloudCredentialNameNil sets the value for CloudCredentialName to be an explicit nil

### UnsetCloudCredentialName
`func (o *ProjectListDetailDto) UnsetCloudCredentialName()`

UnsetCloudCredentialName ensures that no value is present for CloudCredentialName, not even an explicit nil
### GetOrganizationName

`func (o *ProjectListDetailDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectListDetailDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectListDetailDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetOrganizationId

`func (o *ProjectListDetailDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectListDetailDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectListDetailDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetStatus

`func (o *ProjectListDetailDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectListDetailDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectListDetailDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *ProjectListDetailDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectListDetailDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectListDetailDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetCloudType

`func (o *ProjectListDetailDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectListDetailDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectListDetailDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetKubesprayCurrentVersion

`func (o *ProjectListDetailDto) GetKubesprayCurrentVersion() string`

GetKubesprayCurrentVersion returns the KubesprayCurrentVersion field if non-nil, zero value otherwise.

### GetKubesprayCurrentVersionOk

`func (o *ProjectListDetailDto) GetKubesprayCurrentVersionOk() (*string, bool)`

GetKubesprayCurrentVersionOk returns a tuple with the KubesprayCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayCurrentVersion

`func (o *ProjectListDetailDto) SetKubesprayCurrentVersion(v string)`

SetKubesprayCurrentVersion sets KubesprayCurrentVersion field to given value.


### GetKubesprayTargetVersion

`func (o *ProjectListDetailDto) GetKubesprayTargetVersion() string`

GetKubesprayTargetVersion returns the KubesprayTargetVersion field if non-nil, zero value otherwise.

### GetKubesprayTargetVersionOk

`func (o *ProjectListDetailDto) GetKubesprayTargetVersionOk() (*string, bool)`

GetKubesprayTargetVersionOk returns a tuple with the KubesprayTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesprayTargetVersion

`func (o *ProjectListDetailDto) SetKubesprayTargetVersion(v string)`

SetKubesprayTargetVersion sets KubesprayTargetVersion field to given value.


### GetKubernetesCurrentVersion

`func (o *ProjectListDetailDto) GetKubernetesCurrentVersion() string`

GetKubernetesCurrentVersion returns the KubernetesCurrentVersion field if non-nil, zero value otherwise.

### GetKubernetesCurrentVersionOk

`func (o *ProjectListDetailDto) GetKubernetesCurrentVersionOk() (*string, bool)`

GetKubernetesCurrentVersionOk returns a tuple with the KubernetesCurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesCurrentVersion

`func (o *ProjectListDetailDto) SetKubernetesCurrentVersion(v string)`

SetKubernetesCurrentVersion sets KubernetesCurrentVersion field to given value.


### GetKubernetesTargetVersion

`func (o *ProjectListDetailDto) GetKubernetesTargetVersion() string`

GetKubernetesTargetVersion returns the KubernetesTargetVersion field if non-nil, zero value otherwise.

### GetKubernetesTargetVersionOk

`func (o *ProjectListDetailDto) GetKubernetesTargetVersionOk() (*string, bool)`

GetKubernetesTargetVersionOk returns a tuple with the KubernetesTargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetVersion

`func (o *ProjectListDetailDto) SetKubernetesTargetVersion(v string)`

SetKubernetesTargetVersion sets KubernetesTargetVersion field to given value.


### GetCreatedAt

`func (o *ProjectListDetailDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectListDetailDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectListDetailDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetAlertsCount

`func (o *ProjectListDetailDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *ProjectListDetailDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *ProjectListDetailDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetTotalServersCount

`func (o *ProjectListDetailDto) GetTotalServersCount() int32`

GetTotalServersCount returns the TotalServersCount field if non-nil, zero value otherwise.

### GetTotalServersCountOk

`func (o *ProjectListDetailDto) GetTotalServersCountOk() (*int32, bool)`

GetTotalServersCountOk returns a tuple with the TotalServersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServersCount

`func (o *ProjectListDetailDto) SetTotalServersCount(v int32)`

SetTotalServersCount sets TotalServersCount field to given value.


### GetTotalStandaloneVmsCount

`func (o *ProjectListDetailDto) GetTotalStandaloneVmsCount() int32`

GetTotalStandaloneVmsCount returns the TotalStandaloneVmsCount field if non-nil, zero value otherwise.

### GetTotalStandaloneVmsCountOk

`func (o *ProjectListDetailDto) GetTotalStandaloneVmsCountOk() (*int32, bool)`

GetTotalStandaloneVmsCountOk returns a tuple with the TotalStandaloneVmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStandaloneVmsCount

`func (o *ProjectListDetailDto) SetTotalStandaloneVmsCount(v int32)`

SetTotalStandaloneVmsCount sets TotalStandaloneVmsCount field to given value.


### GetBoundUsers

`func (o *ProjectListDetailDto) GetBoundUsers() []UserDto`

GetBoundUsers returns the BoundUsers field if non-nil, zero value otherwise.

### GetBoundUsersOk

`func (o *ProjectListDetailDto) GetBoundUsersOk() (*[]UserDto, bool)`

GetBoundUsersOk returns a tuple with the BoundUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundUsers

`func (o *ProjectListDetailDto) SetBoundUsers(v []UserDto)`

SetBoundUsers sets BoundUsers field to given value.


### GetCreatedBy

`func (o *ProjectListDetailDto) GetCreatedBy() AuditUserDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectListDetailDto) GetCreatedByOk() (*AuditUserDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectListDetailDto) SetCreatedBy(v AuditUserDto)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *ProjectListDetailDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ProjectListDetailDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ProjectListDetailDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *ProjectListDetailDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ProjectListDetailDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetExpiredAt

`func (o *ProjectListDetailDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectListDetailDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectListDetailDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *ProjectListDetailDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ProjectListDetailDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetDeleteOnExpiration

`func (o *ProjectListDetailDto) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *ProjectListDetailDto) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *ProjectListDetailDto) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.


### GetCertificateExpiredAt

`func (o *ProjectListDetailDto) GetCertificateExpiredAt() string`

GetCertificateExpiredAt returns the CertificateExpiredAt field if non-nil, zero value otherwise.

### GetCertificateExpiredAtOk

`func (o *ProjectListDetailDto) GetCertificateExpiredAtOk() (*string, bool)`

GetCertificateExpiredAtOk returns a tuple with the CertificateExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateExpiredAt

`func (o *ProjectListDetailDto) SetCertificateExpiredAt(v string)`

SetCertificateExpiredAt sets CertificateExpiredAt field to given value.


### SetCertificateExpiredAtNil

`func (o *ProjectListDetailDto) SetCertificateExpiredAtNil(b bool)`

 SetCertificateExpiredAtNil sets the value for CertificateExpiredAt to be an explicit nil

### UnsetCertificateExpiredAt
`func (o *ProjectListDetailDto) UnsetCertificateExpiredAt()`

UnsetCertificateExpiredAt ensures that no value is present for CertificateExpiredAt, not even an explicit nil
### GetLastModifiedBy

`func (o *ProjectListDetailDto) GetLastModifiedBy() AuditUserDto`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ProjectListDetailDto) GetLastModifiedByOk() (*AuditUserDto, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ProjectListDetailDto) SetLastModifiedBy(v AuditUserDto)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetQuotaId

`func (o *ProjectListDetailDto) GetQuotaId() int32`

GetQuotaId returns the QuotaId field if non-nil, zero value otherwise.

### GetQuotaIdOk

`func (o *ProjectListDetailDto) GetQuotaIdOk() (*int32, bool)`

GetQuotaIdOk returns a tuple with the QuotaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaId

`func (o *ProjectListDetailDto) SetQuotaId(v int32)`

SetQuotaId sets QuotaId field to given value.


### GetAllowFullSpotKubernetes

`func (o *ProjectListDetailDto) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *ProjectListDetailDto) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *ProjectListDetailDto) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.


### GetAllowSpotWorkers

`func (o *ProjectListDetailDto) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *ProjectListDetailDto) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *ProjectListDetailDto) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.


### GetAllowSpotVMs

`func (o *ProjectListDetailDto) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *ProjectListDetailDto) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *ProjectListDetailDto) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.


### GetMaxSpotPrice

`func (o *ProjectListDetailDto) GetMaxSpotPrice() float64`

GetMaxSpotPrice returns the MaxSpotPrice field if non-nil, zero value otherwise.

### GetMaxSpotPriceOk

`func (o *ProjectListDetailDto) GetMaxSpotPriceOk() (*float64, bool)`

GetMaxSpotPriceOk returns a tuple with the MaxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotPrice

`func (o *ProjectListDetailDto) SetMaxSpotPrice(v float64)`

SetMaxSpotPrice sets MaxSpotPrice field to given value.


### SetMaxSpotPriceNil

`func (o *ProjectListDetailDto) SetMaxSpotPriceNil(b bool)`

 SetMaxSpotPriceNil sets the value for MaxSpotPrice to be an explicit nil

### UnsetMaxSpotPrice
`func (o *ProjectListDetailDto) UnsetMaxSpotPrice()`

UnsetMaxSpotPrice ensures that no value is present for MaxSpotPrice, not even an explicit nil
### GetProjectAction

`func (o *ProjectListDetailDto) GetProjectAction() bool`

GetProjectAction returns the ProjectAction field if non-nil, zero value otherwise.

### GetProjectActionOk

`func (o *ProjectListDetailDto) GetProjectActionOk() (*bool, bool)`

GetProjectActionOk returns a tuple with the ProjectAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAction

`func (o *ProjectListDetailDto) SetProjectAction(v bool)`

SetProjectAction sets ProjectAction field to given value.


### GetHasExpirationWarning

`func (o *ProjectListDetailDto) GetHasExpirationWarning() bool`

GetHasExpirationWarning returns the HasExpirationWarning field if non-nil, zero value otherwise.

### GetHasExpirationWarningOk

`func (o *ProjectListDetailDto) GetHasExpirationWarningOk() (*bool, bool)`

GetHasExpirationWarningOk returns a tuple with the HasExpirationWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpirationWarning

`func (o *ProjectListDetailDto) SetHasExpirationWarning(v bool)`

SetHasExpirationWarning sets HasExpirationWarning field to given value.


### GetIsAutoscalingEnabled

`func (o *ProjectListDetailDto) GetIsAutoscalingEnabled() bool`

GetIsAutoscalingEnabled returns the IsAutoscalingEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingEnabledOk

`func (o *ProjectListDetailDto) GetIsAutoscalingEnabledOk() (*bool, bool)`

GetIsAutoscalingEnabledOk returns a tuple with the IsAutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingEnabled

`func (o *ProjectListDetailDto) SetIsAutoscalingEnabled(v bool)`

SetIsAutoscalingEnabled sets IsAutoscalingEnabled field to given value.


### GetIsAutoscalingSpotEnabled

`func (o *ProjectListDetailDto) GetIsAutoscalingSpotEnabled() bool`

GetIsAutoscalingSpotEnabled returns the IsAutoscalingSpotEnabled field if non-nil, zero value otherwise.

### GetIsAutoscalingSpotEnabledOk

`func (o *ProjectListDetailDto) GetIsAutoscalingSpotEnabledOk() (*bool, bool)`

GetIsAutoscalingSpotEnabledOk returns a tuple with the IsAutoscalingSpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalingSpotEnabled

`func (o *ProjectListDetailDto) SetIsAutoscalingSpotEnabled(v bool)`

SetIsAutoscalingSpotEnabled sets IsAutoscalingSpotEnabled field to given value.


### GetGatewayApiModern

`func (o *ProjectListDetailDto) GetGatewayApiModern() bool`

GetGatewayApiModern returns the GatewayApiModern field if non-nil, zero value otherwise.

### GetGatewayApiModernOk

`func (o *ProjectListDetailDto) GetGatewayApiModernOk() (*bool, bool)`

GetGatewayApiModernOk returns a tuple with the GatewayApiModern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayApiModern

`func (o *ProjectListDetailDto) SetGatewayApiModern(v bool)`

SetGatewayApiModern sets GatewayApiModern field to given value.


### GetAiEnabled

`func (o *ProjectListDetailDto) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *ProjectListDetailDto) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *ProjectListDetailDto) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.


### GetAnyServer

`func (o *ProjectListDetailDto) GetAnyServer() bool`

GetAnyServer returns the AnyServer field if non-nil, zero value otherwise.

### GetAnyServerOk

`func (o *ProjectListDetailDto) GetAnyServerOk() (*bool, bool)`

GetAnyServerOk returns a tuple with the AnyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyServer

`func (o *ProjectListDetailDto) SetAnyServer(v bool)`

SetAnyServer sets AnyServer field to given value.


### GetAnyVm

`func (o *ProjectListDetailDto) GetAnyVm() bool`

GetAnyVm returns the AnyVm field if non-nil, zero value otherwise.

### GetAnyVmOk

`func (o *ProjectListDetailDto) GetAnyVmOk() (*bool, bool)`

GetAnyVmOk returns a tuple with the AnyVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyVm

`func (o *ProjectListDetailDto) SetAnyVm(v bool)`

SetAnyVm sets AnyVm field to given value.


### GetIsBackupEnabled

`func (o *ProjectListDetailDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectListDetailDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectListDetailDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.


### GetIsProjectMaintenanceModeEnabled

`func (o *ProjectListDetailDto) GetIsProjectMaintenanceModeEnabled() bool`

GetIsProjectMaintenanceModeEnabled returns the IsProjectMaintenanceModeEnabled field if non-nil, zero value otherwise.

### GetIsProjectMaintenanceModeEnabledOk

`func (o *ProjectListDetailDto) GetIsProjectMaintenanceModeEnabledOk() (*bool, bool)`

GetIsProjectMaintenanceModeEnabledOk returns a tuple with the IsProjectMaintenanceModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProjectMaintenanceModeEnabled

`func (o *ProjectListDetailDto) SetIsProjectMaintenanceModeEnabled(v bool)`

SetIsProjectMaintenanceModeEnabled sets IsProjectMaintenanceModeEnabled field to given value.


### GetParentProjectId

`func (o *ProjectListDetailDto) GetParentProjectId() int32`

GetParentProjectId returns the ParentProjectId field if non-nil, zero value otherwise.

### GetParentProjectIdOk

`func (o *ProjectListDetailDto) GetParentProjectIdOk() (*int32, bool)`

GetParentProjectIdOk returns a tuple with the ParentProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProjectId

`func (o *ProjectListDetailDto) SetParentProjectId(v int32)`

SetParentProjectId sets ParentProjectId field to given value.


### SetParentProjectIdNil

`func (o *ProjectListDetailDto) SetParentProjectIdNil(b bool)`

 SetParentProjectIdNil sets the value for ParentProjectId to be an explicit nil

### UnsetParentProjectId
`func (o *ProjectListDetailDto) UnsetParentProjectId()`

UnsetParentProjectId ensures that no value is present for ParentProjectId, not even an explicit nil
### GetAlertingProfileId

`func (o *ProjectListDetailDto) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *ProjectListDetailDto) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *ProjectListDetailDto) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.


### SetAlertingProfileIdNil

`func (o *ProjectListDetailDto) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *ProjectListDetailDto) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetOpaProfileId

`func (o *ProjectListDetailDto) GetOpaProfileId() int32`

GetOpaProfileId returns the OpaProfileId field if non-nil, zero value otherwise.

### GetOpaProfileIdOk

`func (o *ProjectListDetailDto) GetOpaProfileIdOk() (*int32, bool)`

GetOpaProfileIdOk returns a tuple with the OpaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileId

`func (o *ProjectListDetailDto) SetOpaProfileId(v int32)`

SetOpaProfileId sets OpaProfileId field to given value.


### SetOpaProfileIdNil

`func (o *ProjectListDetailDto) SetOpaProfileIdNil(b bool)`

 SetOpaProfileIdNil sets the value for OpaProfileId to be an explicit nil

### UnsetOpaProfileId
`func (o *ProjectListDetailDto) UnsetOpaProfileId()`

UnsetOpaProfileId ensures that no value is present for OpaProfileId, not even an explicit nil
### GetLockButton

`func (o *ProjectListDetailDto) GetLockButton() ButtonStatusDto`

GetLockButton returns the LockButton field if non-nil, zero value otherwise.

### GetLockButtonOk

`func (o *ProjectListDetailDto) GetLockButtonOk() (*ButtonStatusDto, bool)`

GetLockButtonOk returns a tuple with the LockButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockButton

`func (o *ProjectListDetailDto) SetLockButton(v ButtonStatusDto)`

SetLockButton sets LockButton field to given value.

### HasLockButton

`func (o *ProjectListDetailDto) HasLockButton() bool`

HasLockButton returns a boolean if a field has been set.

### GetUnlockButton

`func (o *ProjectListDetailDto) GetUnlockButton() ButtonStatusDto`

GetUnlockButton returns the UnlockButton field if non-nil, zero value otherwise.

### GetUnlockButtonOk

`func (o *ProjectListDetailDto) GetUnlockButtonOk() (*ButtonStatusDto, bool)`

GetUnlockButtonOk returns a tuple with the UnlockButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockButton

`func (o *ProjectListDetailDto) SetUnlockButton(v ButtonStatusDto)`

SetUnlockButton sets UnlockButton field to given value.

### HasUnlockButton

`func (o *ProjectListDetailDto) HasUnlockButton() bool`

HasUnlockButton returns a boolean if a field has been set.

### GetDeleteButton

`func (o *ProjectListDetailDto) GetDeleteButton() ButtonStatusDto`

GetDeleteButton returns the DeleteButton field if non-nil, zero value otherwise.

### GetDeleteButtonOk

`func (o *ProjectListDetailDto) GetDeleteButtonOk() (*ButtonStatusDto, bool)`

GetDeleteButtonOk returns a tuple with the DeleteButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteButton

`func (o *ProjectListDetailDto) SetDeleteButton(v ButtonStatusDto)`

SetDeleteButton sets DeleteButton field to given value.

### HasDeleteButton

`func (o *ProjectListDetailDto) HasDeleteButton() bool`

HasDeleteButton returns a boolean if a field has been set.

### GetKubeInfoButton

`func (o *ProjectListDetailDto) GetKubeInfoButton() ButtonStatusDto`

GetKubeInfoButton returns the KubeInfoButton field if non-nil, zero value otherwise.

### GetKubeInfoButtonOk

`func (o *ProjectListDetailDto) GetKubeInfoButtonOk() (*ButtonStatusDto, bool)`

GetKubeInfoButtonOk returns a tuple with the KubeInfoButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeInfoButton

`func (o *ProjectListDetailDto) SetKubeInfoButton(v ButtonStatusDto)`

SetKubeInfoButton sets KubeInfoButton field to given value.

### HasKubeInfoButton

`func (o *ProjectListDetailDto) HasKubeInfoButton() bool`

HasKubeInfoButton returns a boolean if a field has been set.

### GetSetExpirationDateButton

`func (o *ProjectListDetailDto) GetSetExpirationDateButton() ButtonStatusDto`

GetSetExpirationDateButton returns the SetExpirationDateButton field if non-nil, zero value otherwise.

### GetSetExpirationDateButtonOk

`func (o *ProjectListDetailDto) GetSetExpirationDateButtonOk() (*ButtonStatusDto, bool)`

GetSetExpirationDateButtonOk returns a tuple with the SetExpirationDateButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetExpirationDateButton

`func (o *ProjectListDetailDto) SetSetExpirationDateButton(v ButtonStatusDto)`

SetSetExpirationDateButton sets SetExpirationDateButton field to given value.

### HasSetExpirationDateButton

`func (o *ProjectListDetailDto) HasSetExpirationDateButton() bool`

HasSetExpirationDateButton returns a boolean if a field has been set.

### GetResetStatusButton

`func (o *ProjectListDetailDto) GetResetStatusButton() ButtonStatusDto`

GetResetStatusButton returns the ResetStatusButton field if non-nil, zero value otherwise.

### GetResetStatusButtonOk

`func (o *ProjectListDetailDto) GetResetStatusButtonOk() (*ButtonStatusDto, bool)`

GetResetStatusButtonOk returns a tuple with the ResetStatusButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetStatusButton

`func (o *ProjectListDetailDto) SetResetStatusButton(v ButtonStatusDto)`

SetResetStatusButton sets ResetStatusButton field to given value.

### HasResetStatusButton

`func (o *ProjectListDetailDto) HasResetStatusButton() bool`

HasResetStatusButton returns a boolean if a field has been set.

### GetImportClusterType

`func (o *ProjectListDetailDto) GetImportClusterType() ImportClusterType`

GetImportClusterType returns the ImportClusterType field if non-nil, zero value otherwise.

### GetImportClusterTypeOk

`func (o *ProjectListDetailDto) GetImportClusterTypeOk() (*ImportClusterType, bool)`

GetImportClusterTypeOk returns a tuple with the ImportClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportClusterType

`func (o *ProjectListDetailDto) SetImportClusterType(v ImportClusterType)`

SetImportClusterType sets ImportClusterType field to given value.


### GetEImportClusterType

`func (o *ProjectListDetailDto) GetEImportClusterType() map[string]interface{}`

GetEImportClusterType returns the EImportClusterType field if non-nil, zero value otherwise.

### GetEImportClusterTypeOk

`func (o *ProjectListDetailDto) GetEImportClusterTypeOk() (*map[string]interface{}, bool)`

GetEImportClusterTypeOk returns a tuple with the EImportClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEImportClusterType

`func (o *ProjectListDetailDto) SetEImportClusterType(v map[string]interface{})`

SetEImportClusterType sets EImportClusterType field to given value.

### HasEImportClusterType

`func (o *ProjectListDetailDto) HasEImportClusterType() bool`

HasEImportClusterType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


