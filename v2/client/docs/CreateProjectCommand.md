# CreateProjectCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**KubernetesVersion** | Pointer to **NullableString** |  | [optional] 
**CloudCredentialId** | Pointer to **int32** |  | [optional] 
**S3CredentialId** | Pointer to **NullableInt32** |  | [optional] 
**AccessProfileId** | Pointer to **NullableInt32** |  | [optional] 
**OpaProfileId** | Pointer to **NullableInt32** |  | [optional] 
**KubernetesProfileId** | Pointer to **NullableInt32** |  | [optional] 
**IsKubernetes** | Pointer to **bool** |  | [optional] 
**IsAutoUpgrade** | Pointer to **bool** |  | [optional] 
**IsBackupEnabled** | Pointer to **bool** |  | [optional] 
**IsMonitoringEnabled** | Pointer to **bool** |  | [optional] 
**AiEnabled** | Pointer to **bool** |  | [optional] 
**AiCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**Flavors** | Pointer to **[]string** |  | [optional] 
**AlertingProfileId** | Pointer to **NullableInt32** |  | [optional] 
**DnsCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**IsDnsEnabled** | Pointer to **bool** |  | [optional] 
**IsCertTrustEnabled** | Pointer to **bool** |  | [optional] 
**CertificateProfileId** | Pointer to **NullableInt32** |  | [optional] 
**TaikunLBFlavor** | Pointer to **NullableString** |  | [optional] 
**RouterIdStartRange** | Pointer to **NullableInt32** |  | [optional] 
**RouterIdEndRange** | Pointer to **NullableInt32** |  | [optional] 
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**DeleteOnExpiration** | Pointer to **bool** |  | [optional] 
**AllowFullSpotKubernetes** | Pointer to **bool** |  | [optional] 
**AllowSpotWorkers** | Pointer to **bool** |  | [optional] 
**AllowSpotVMs** | Pointer to **bool** |  | [optional] 
**MaxSpotPrice** | Pointer to **NullableFloat64** |  | [optional] 
**AutoscalingEnabled** | Pointer to **bool** |  | [optional] 
**MinSize** | Pointer to **int32** |  | [optional] 
**MaxSize** | Pointer to **int32** |  | [optional] 
**DiskSize** | Pointer to **float64** |  | [optional] 
**AutoscalingFlavor** | Pointer to **NullableString** |  | [optional] 
**AutoscalingSpotEnabled** | Pointer to **bool** |  | [optional] 
**IsCskEnabled** | Pointer to **bool** |  | [optional] 
**Cidr** | Pointer to **NullableString** |  | [optional] 
**NetMask** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateProjectCommand

`func NewCreateProjectCommand() *CreateProjectCommand`

NewCreateProjectCommand instantiates a new CreateProjectCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectCommandWithDefaults

`func NewCreateProjectCommandWithDefaults() *CreateProjectCommand`

NewCreateProjectCommandWithDefaults instantiates a new CreateProjectCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateProjectCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateProjectCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateProjectCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateProjectCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKubernetesVersion

`func (o *CreateProjectCommand) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *CreateProjectCommand) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *CreateProjectCommand) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *CreateProjectCommand) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### SetKubernetesVersionNil

`func (o *CreateProjectCommand) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *CreateProjectCommand) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetCloudCredentialId

`func (o *CreateProjectCommand) GetCloudCredentialId() int32`

GetCloudCredentialId returns the CloudCredentialId field if non-nil, zero value otherwise.

### GetCloudCredentialIdOk

`func (o *CreateProjectCommand) GetCloudCredentialIdOk() (*int32, bool)`

GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudCredentialId

`func (o *CreateProjectCommand) SetCloudCredentialId(v int32)`

SetCloudCredentialId sets CloudCredentialId field to given value.

### HasCloudCredentialId

`func (o *CreateProjectCommand) HasCloudCredentialId() bool`

HasCloudCredentialId returns a boolean if a field has been set.

### GetS3CredentialId

`func (o *CreateProjectCommand) GetS3CredentialId() int32`

GetS3CredentialId returns the S3CredentialId field if non-nil, zero value otherwise.

### GetS3CredentialIdOk

`func (o *CreateProjectCommand) GetS3CredentialIdOk() (*int32, bool)`

GetS3CredentialIdOk returns a tuple with the S3CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CredentialId

`func (o *CreateProjectCommand) SetS3CredentialId(v int32)`

SetS3CredentialId sets S3CredentialId field to given value.

### HasS3CredentialId

`func (o *CreateProjectCommand) HasS3CredentialId() bool`

HasS3CredentialId returns a boolean if a field has been set.

### SetS3CredentialIdNil

`func (o *CreateProjectCommand) SetS3CredentialIdNil(b bool)`

 SetS3CredentialIdNil sets the value for S3CredentialId to be an explicit nil

### UnsetS3CredentialId
`func (o *CreateProjectCommand) UnsetS3CredentialId()`

UnsetS3CredentialId ensures that no value is present for S3CredentialId, not even an explicit nil
### GetAccessProfileId

`func (o *CreateProjectCommand) GetAccessProfileId() int32`

GetAccessProfileId returns the AccessProfileId field if non-nil, zero value otherwise.

### GetAccessProfileIdOk

`func (o *CreateProjectCommand) GetAccessProfileIdOk() (*int32, bool)`

GetAccessProfileIdOk returns a tuple with the AccessProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessProfileId

`func (o *CreateProjectCommand) SetAccessProfileId(v int32)`

SetAccessProfileId sets AccessProfileId field to given value.

### HasAccessProfileId

`func (o *CreateProjectCommand) HasAccessProfileId() bool`

HasAccessProfileId returns a boolean if a field has been set.

### SetAccessProfileIdNil

`func (o *CreateProjectCommand) SetAccessProfileIdNil(b bool)`

 SetAccessProfileIdNil sets the value for AccessProfileId to be an explicit nil

### UnsetAccessProfileId
`func (o *CreateProjectCommand) UnsetAccessProfileId()`

UnsetAccessProfileId ensures that no value is present for AccessProfileId, not even an explicit nil
### GetOpaProfileId

`func (o *CreateProjectCommand) GetOpaProfileId() int32`

GetOpaProfileId returns the OpaProfileId field if non-nil, zero value otherwise.

### GetOpaProfileIdOk

`func (o *CreateProjectCommand) GetOpaProfileIdOk() (*int32, bool)`

GetOpaProfileIdOk returns a tuple with the OpaProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaProfileId

`func (o *CreateProjectCommand) SetOpaProfileId(v int32)`

SetOpaProfileId sets OpaProfileId field to given value.

### HasOpaProfileId

`func (o *CreateProjectCommand) HasOpaProfileId() bool`

HasOpaProfileId returns a boolean if a field has been set.

### SetOpaProfileIdNil

`func (o *CreateProjectCommand) SetOpaProfileIdNil(b bool)`

 SetOpaProfileIdNil sets the value for OpaProfileId to be an explicit nil

### UnsetOpaProfileId
`func (o *CreateProjectCommand) UnsetOpaProfileId()`

UnsetOpaProfileId ensures that no value is present for OpaProfileId, not even an explicit nil
### GetKubernetesProfileId

`func (o *CreateProjectCommand) GetKubernetesProfileId() int32`

GetKubernetesProfileId returns the KubernetesProfileId field if non-nil, zero value otherwise.

### GetKubernetesProfileIdOk

`func (o *CreateProjectCommand) GetKubernetesProfileIdOk() (*int32, bool)`

GetKubernetesProfileIdOk returns a tuple with the KubernetesProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesProfileId

`func (o *CreateProjectCommand) SetKubernetesProfileId(v int32)`

SetKubernetesProfileId sets KubernetesProfileId field to given value.

### HasKubernetesProfileId

`func (o *CreateProjectCommand) HasKubernetesProfileId() bool`

HasKubernetesProfileId returns a boolean if a field has been set.

### SetKubernetesProfileIdNil

`func (o *CreateProjectCommand) SetKubernetesProfileIdNil(b bool)`

 SetKubernetesProfileIdNil sets the value for KubernetesProfileId to be an explicit nil

### UnsetKubernetesProfileId
`func (o *CreateProjectCommand) UnsetKubernetesProfileId()`

UnsetKubernetesProfileId ensures that no value is present for KubernetesProfileId, not even an explicit nil
### GetIsKubernetes

`func (o *CreateProjectCommand) GetIsKubernetes() bool`

GetIsKubernetes returns the IsKubernetes field if non-nil, zero value otherwise.

### GetIsKubernetesOk

`func (o *CreateProjectCommand) GetIsKubernetesOk() (*bool, bool)`

GetIsKubernetesOk returns a tuple with the IsKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKubernetes

`func (o *CreateProjectCommand) SetIsKubernetes(v bool)`

SetIsKubernetes sets IsKubernetes field to given value.

### HasIsKubernetes

`func (o *CreateProjectCommand) HasIsKubernetes() bool`

HasIsKubernetes returns a boolean if a field has been set.

### GetIsAutoUpgrade

`func (o *CreateProjectCommand) GetIsAutoUpgrade() bool`

GetIsAutoUpgrade returns the IsAutoUpgrade field if non-nil, zero value otherwise.

### GetIsAutoUpgradeOk

`func (o *CreateProjectCommand) GetIsAutoUpgradeOk() (*bool, bool)`

GetIsAutoUpgradeOk returns a tuple with the IsAutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgrade

`func (o *CreateProjectCommand) SetIsAutoUpgrade(v bool)`

SetIsAutoUpgrade sets IsAutoUpgrade field to given value.

### HasIsAutoUpgrade

`func (o *CreateProjectCommand) HasIsAutoUpgrade() bool`

HasIsAutoUpgrade returns a boolean if a field has been set.

### GetIsBackupEnabled

`func (o *CreateProjectCommand) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *CreateProjectCommand) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *CreateProjectCommand) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.

### HasIsBackupEnabled

`func (o *CreateProjectCommand) HasIsBackupEnabled() bool`

HasIsBackupEnabled returns a boolean if a field has been set.

### GetIsMonitoringEnabled

`func (o *CreateProjectCommand) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *CreateProjectCommand) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *CreateProjectCommand) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.

### HasIsMonitoringEnabled

`func (o *CreateProjectCommand) HasIsMonitoringEnabled() bool`

HasIsMonitoringEnabled returns a boolean if a field has been set.

### GetAiEnabled

`func (o *CreateProjectCommand) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *CreateProjectCommand) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *CreateProjectCommand) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.

### HasAiEnabled

`func (o *CreateProjectCommand) HasAiEnabled() bool`

HasAiEnabled returns a boolean if a field has been set.

### GetAiCredentialId

`func (o *CreateProjectCommand) GetAiCredentialId() int32`

GetAiCredentialId returns the AiCredentialId field if non-nil, zero value otherwise.

### GetAiCredentialIdOk

`func (o *CreateProjectCommand) GetAiCredentialIdOk() (*int32, bool)`

GetAiCredentialIdOk returns a tuple with the AiCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiCredentialId

`func (o *CreateProjectCommand) SetAiCredentialId(v int32)`

SetAiCredentialId sets AiCredentialId field to given value.

### HasAiCredentialId

`func (o *CreateProjectCommand) HasAiCredentialId() bool`

HasAiCredentialId returns a boolean if a field has been set.

### SetAiCredentialIdNil

`func (o *CreateProjectCommand) SetAiCredentialIdNil(b bool)`

 SetAiCredentialIdNil sets the value for AiCredentialId to be an explicit nil

### UnsetAiCredentialId
`func (o *CreateProjectCommand) UnsetAiCredentialId()`

UnsetAiCredentialId ensures that no value is present for AiCredentialId, not even an explicit nil
### GetFlavors

`func (o *CreateProjectCommand) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *CreateProjectCommand) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *CreateProjectCommand) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *CreateProjectCommand) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### SetFlavorsNil

`func (o *CreateProjectCommand) SetFlavorsNil(b bool)`

 SetFlavorsNil sets the value for Flavors to be an explicit nil

### UnsetFlavors
`func (o *CreateProjectCommand) UnsetFlavors()`

UnsetFlavors ensures that no value is present for Flavors, not even an explicit nil
### GetAlertingProfileId

`func (o *CreateProjectCommand) GetAlertingProfileId() int32`

GetAlertingProfileId returns the AlertingProfileId field if non-nil, zero value otherwise.

### GetAlertingProfileIdOk

`func (o *CreateProjectCommand) GetAlertingProfileIdOk() (*int32, bool)`

GetAlertingProfileIdOk returns a tuple with the AlertingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfileId

`func (o *CreateProjectCommand) SetAlertingProfileId(v int32)`

SetAlertingProfileId sets AlertingProfileId field to given value.

### HasAlertingProfileId

`func (o *CreateProjectCommand) HasAlertingProfileId() bool`

HasAlertingProfileId returns a boolean if a field has been set.

### SetAlertingProfileIdNil

`func (o *CreateProjectCommand) SetAlertingProfileIdNil(b bool)`

 SetAlertingProfileIdNil sets the value for AlertingProfileId to be an explicit nil

### UnsetAlertingProfileId
`func (o *CreateProjectCommand) UnsetAlertingProfileId()`

UnsetAlertingProfileId ensures that no value is present for AlertingProfileId, not even an explicit nil
### GetDnsCredentialId

`func (o *CreateProjectCommand) GetDnsCredentialId() int32`

GetDnsCredentialId returns the DnsCredentialId field if non-nil, zero value otherwise.

### GetDnsCredentialIdOk

`func (o *CreateProjectCommand) GetDnsCredentialIdOk() (*int32, bool)`

GetDnsCredentialIdOk returns a tuple with the DnsCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCredentialId

`func (o *CreateProjectCommand) SetDnsCredentialId(v int32)`

SetDnsCredentialId sets DnsCredentialId field to given value.

### HasDnsCredentialId

`func (o *CreateProjectCommand) HasDnsCredentialId() bool`

HasDnsCredentialId returns a boolean if a field has been set.

### SetDnsCredentialIdNil

`func (o *CreateProjectCommand) SetDnsCredentialIdNil(b bool)`

 SetDnsCredentialIdNil sets the value for DnsCredentialId to be an explicit nil

### UnsetDnsCredentialId
`func (o *CreateProjectCommand) UnsetDnsCredentialId()`

UnsetDnsCredentialId ensures that no value is present for DnsCredentialId, not even an explicit nil
### GetIsDnsEnabled

`func (o *CreateProjectCommand) GetIsDnsEnabled() bool`

GetIsDnsEnabled returns the IsDnsEnabled field if non-nil, zero value otherwise.

### GetIsDnsEnabledOk

`func (o *CreateProjectCommand) GetIsDnsEnabledOk() (*bool, bool)`

GetIsDnsEnabledOk returns a tuple with the IsDnsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDnsEnabled

`func (o *CreateProjectCommand) SetIsDnsEnabled(v bool)`

SetIsDnsEnabled sets IsDnsEnabled field to given value.

### HasIsDnsEnabled

`func (o *CreateProjectCommand) HasIsDnsEnabled() bool`

HasIsDnsEnabled returns a boolean if a field has been set.

### GetIsCertTrustEnabled

`func (o *CreateProjectCommand) GetIsCertTrustEnabled() bool`

GetIsCertTrustEnabled returns the IsCertTrustEnabled field if non-nil, zero value otherwise.

### GetIsCertTrustEnabledOk

`func (o *CreateProjectCommand) GetIsCertTrustEnabledOk() (*bool, bool)`

GetIsCertTrustEnabledOk returns a tuple with the IsCertTrustEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCertTrustEnabled

`func (o *CreateProjectCommand) SetIsCertTrustEnabled(v bool)`

SetIsCertTrustEnabled sets IsCertTrustEnabled field to given value.

### HasIsCertTrustEnabled

`func (o *CreateProjectCommand) HasIsCertTrustEnabled() bool`

HasIsCertTrustEnabled returns a boolean if a field has been set.

### GetCertificateProfileId

`func (o *CreateProjectCommand) GetCertificateProfileId() int32`

GetCertificateProfileId returns the CertificateProfileId field if non-nil, zero value otherwise.

### GetCertificateProfileIdOk

`func (o *CreateProjectCommand) GetCertificateProfileIdOk() (*int32, bool)`

GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateProfileId

`func (o *CreateProjectCommand) SetCertificateProfileId(v int32)`

SetCertificateProfileId sets CertificateProfileId field to given value.

### HasCertificateProfileId

`func (o *CreateProjectCommand) HasCertificateProfileId() bool`

HasCertificateProfileId returns a boolean if a field has been set.

### SetCertificateProfileIdNil

`func (o *CreateProjectCommand) SetCertificateProfileIdNil(b bool)`

 SetCertificateProfileIdNil sets the value for CertificateProfileId to be an explicit nil

### UnsetCertificateProfileId
`func (o *CreateProjectCommand) UnsetCertificateProfileId()`

UnsetCertificateProfileId ensures that no value is present for CertificateProfileId, not even an explicit nil
### GetTaikunLBFlavor

`func (o *CreateProjectCommand) GetTaikunLBFlavor() string`

GetTaikunLBFlavor returns the TaikunLBFlavor field if non-nil, zero value otherwise.

### GetTaikunLBFlavorOk

`func (o *CreateProjectCommand) GetTaikunLBFlavorOk() (*string, bool)`

GetTaikunLBFlavorOk returns a tuple with the TaikunLBFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBFlavor

`func (o *CreateProjectCommand) SetTaikunLBFlavor(v string)`

SetTaikunLBFlavor sets TaikunLBFlavor field to given value.

### HasTaikunLBFlavor

`func (o *CreateProjectCommand) HasTaikunLBFlavor() bool`

HasTaikunLBFlavor returns a boolean if a field has been set.

### SetTaikunLBFlavorNil

`func (o *CreateProjectCommand) SetTaikunLBFlavorNil(b bool)`

 SetTaikunLBFlavorNil sets the value for TaikunLBFlavor to be an explicit nil

### UnsetTaikunLBFlavor
`func (o *CreateProjectCommand) UnsetTaikunLBFlavor()`

UnsetTaikunLBFlavor ensures that no value is present for TaikunLBFlavor, not even an explicit nil
### GetRouterIdStartRange

`func (o *CreateProjectCommand) GetRouterIdStartRange() int32`

GetRouterIdStartRange returns the RouterIdStartRange field if non-nil, zero value otherwise.

### GetRouterIdStartRangeOk

`func (o *CreateProjectCommand) GetRouterIdStartRangeOk() (*int32, bool)`

GetRouterIdStartRangeOk returns a tuple with the RouterIdStartRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIdStartRange

`func (o *CreateProjectCommand) SetRouterIdStartRange(v int32)`

SetRouterIdStartRange sets RouterIdStartRange field to given value.

### HasRouterIdStartRange

`func (o *CreateProjectCommand) HasRouterIdStartRange() bool`

HasRouterIdStartRange returns a boolean if a field has been set.

### SetRouterIdStartRangeNil

`func (o *CreateProjectCommand) SetRouterIdStartRangeNil(b bool)`

 SetRouterIdStartRangeNil sets the value for RouterIdStartRange to be an explicit nil

### UnsetRouterIdStartRange
`func (o *CreateProjectCommand) UnsetRouterIdStartRange()`

UnsetRouterIdStartRange ensures that no value is present for RouterIdStartRange, not even an explicit nil
### GetRouterIdEndRange

`func (o *CreateProjectCommand) GetRouterIdEndRange() int32`

GetRouterIdEndRange returns the RouterIdEndRange field if non-nil, zero value otherwise.

### GetRouterIdEndRangeOk

`func (o *CreateProjectCommand) GetRouterIdEndRangeOk() (*int32, bool)`

GetRouterIdEndRangeOk returns a tuple with the RouterIdEndRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIdEndRange

`func (o *CreateProjectCommand) SetRouterIdEndRange(v int32)`

SetRouterIdEndRange sets RouterIdEndRange field to given value.

### HasRouterIdEndRange

`func (o *CreateProjectCommand) HasRouterIdEndRange() bool`

HasRouterIdEndRange returns a boolean if a field has been set.

### SetRouterIdEndRangeNil

`func (o *CreateProjectCommand) SetRouterIdEndRangeNil(b bool)`

 SetRouterIdEndRangeNil sets the value for RouterIdEndRange to be an explicit nil

### UnsetRouterIdEndRange
`func (o *CreateProjectCommand) UnsetRouterIdEndRange()`

UnsetRouterIdEndRange ensures that no value is present for RouterIdEndRange, not even an explicit nil
### GetExpiredAt

`func (o *CreateProjectCommand) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *CreateProjectCommand) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *CreateProjectCommand) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *CreateProjectCommand) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *CreateProjectCommand) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *CreateProjectCommand) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetDeleteOnExpiration

`func (o *CreateProjectCommand) GetDeleteOnExpiration() bool`

GetDeleteOnExpiration returns the DeleteOnExpiration field if non-nil, zero value otherwise.

### GetDeleteOnExpirationOk

`func (o *CreateProjectCommand) GetDeleteOnExpirationOk() (*bool, bool)`

GetDeleteOnExpirationOk returns a tuple with the DeleteOnExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnExpiration

`func (o *CreateProjectCommand) SetDeleteOnExpiration(v bool)`

SetDeleteOnExpiration sets DeleteOnExpiration field to given value.

### HasDeleteOnExpiration

`func (o *CreateProjectCommand) HasDeleteOnExpiration() bool`

HasDeleteOnExpiration returns a boolean if a field has been set.

### GetAllowFullSpotKubernetes

`func (o *CreateProjectCommand) GetAllowFullSpotKubernetes() bool`

GetAllowFullSpotKubernetes returns the AllowFullSpotKubernetes field if non-nil, zero value otherwise.

### GetAllowFullSpotKubernetesOk

`func (o *CreateProjectCommand) GetAllowFullSpotKubernetesOk() (*bool, bool)`

GetAllowFullSpotKubernetesOk returns a tuple with the AllowFullSpotKubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFullSpotKubernetes

`func (o *CreateProjectCommand) SetAllowFullSpotKubernetes(v bool)`

SetAllowFullSpotKubernetes sets AllowFullSpotKubernetes field to given value.

### HasAllowFullSpotKubernetes

`func (o *CreateProjectCommand) HasAllowFullSpotKubernetes() bool`

HasAllowFullSpotKubernetes returns a boolean if a field has been set.

### GetAllowSpotWorkers

`func (o *CreateProjectCommand) GetAllowSpotWorkers() bool`

GetAllowSpotWorkers returns the AllowSpotWorkers field if non-nil, zero value otherwise.

### GetAllowSpotWorkersOk

`func (o *CreateProjectCommand) GetAllowSpotWorkersOk() (*bool, bool)`

GetAllowSpotWorkersOk returns a tuple with the AllowSpotWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotWorkers

`func (o *CreateProjectCommand) SetAllowSpotWorkers(v bool)`

SetAllowSpotWorkers sets AllowSpotWorkers field to given value.

### HasAllowSpotWorkers

`func (o *CreateProjectCommand) HasAllowSpotWorkers() bool`

HasAllowSpotWorkers returns a boolean if a field has been set.

### GetAllowSpotVMs

`func (o *CreateProjectCommand) GetAllowSpotVMs() bool`

GetAllowSpotVMs returns the AllowSpotVMs field if non-nil, zero value otherwise.

### GetAllowSpotVMsOk

`func (o *CreateProjectCommand) GetAllowSpotVMsOk() (*bool, bool)`

GetAllowSpotVMsOk returns a tuple with the AllowSpotVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSpotVMs

`func (o *CreateProjectCommand) SetAllowSpotVMs(v bool)`

SetAllowSpotVMs sets AllowSpotVMs field to given value.

### HasAllowSpotVMs

`func (o *CreateProjectCommand) HasAllowSpotVMs() bool`

HasAllowSpotVMs returns a boolean if a field has been set.

### GetMaxSpotPrice

`func (o *CreateProjectCommand) GetMaxSpotPrice() float64`

GetMaxSpotPrice returns the MaxSpotPrice field if non-nil, zero value otherwise.

### GetMaxSpotPriceOk

`func (o *CreateProjectCommand) GetMaxSpotPriceOk() (*float64, bool)`

GetMaxSpotPriceOk returns a tuple with the MaxSpotPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpotPrice

`func (o *CreateProjectCommand) SetMaxSpotPrice(v float64)`

SetMaxSpotPrice sets MaxSpotPrice field to given value.

### HasMaxSpotPrice

`func (o *CreateProjectCommand) HasMaxSpotPrice() bool`

HasMaxSpotPrice returns a boolean if a field has been set.

### SetMaxSpotPriceNil

`func (o *CreateProjectCommand) SetMaxSpotPriceNil(b bool)`

 SetMaxSpotPriceNil sets the value for MaxSpotPrice to be an explicit nil

### UnsetMaxSpotPrice
`func (o *CreateProjectCommand) UnsetMaxSpotPrice()`

UnsetMaxSpotPrice ensures that no value is present for MaxSpotPrice, not even an explicit nil
### GetAutoscalingEnabled

`func (o *CreateProjectCommand) GetAutoscalingEnabled() bool`

GetAutoscalingEnabled returns the AutoscalingEnabled field if non-nil, zero value otherwise.

### GetAutoscalingEnabledOk

`func (o *CreateProjectCommand) GetAutoscalingEnabledOk() (*bool, bool)`

GetAutoscalingEnabledOk returns a tuple with the AutoscalingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingEnabled

`func (o *CreateProjectCommand) SetAutoscalingEnabled(v bool)`

SetAutoscalingEnabled sets AutoscalingEnabled field to given value.

### HasAutoscalingEnabled

`func (o *CreateProjectCommand) HasAutoscalingEnabled() bool`

HasAutoscalingEnabled returns a boolean if a field has been set.

### GetMinSize

`func (o *CreateProjectCommand) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *CreateProjectCommand) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *CreateProjectCommand) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *CreateProjectCommand) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### GetMaxSize

`func (o *CreateProjectCommand) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *CreateProjectCommand) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *CreateProjectCommand) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *CreateProjectCommand) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetDiskSize

`func (o *CreateProjectCommand) GetDiskSize() float64`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *CreateProjectCommand) GetDiskSizeOk() (*float64, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *CreateProjectCommand) SetDiskSize(v float64)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *CreateProjectCommand) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetAutoscalingFlavor

`func (o *CreateProjectCommand) GetAutoscalingFlavor() string`

GetAutoscalingFlavor returns the AutoscalingFlavor field if non-nil, zero value otherwise.

### GetAutoscalingFlavorOk

`func (o *CreateProjectCommand) GetAutoscalingFlavorOk() (*string, bool)`

GetAutoscalingFlavorOk returns a tuple with the AutoscalingFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingFlavor

`func (o *CreateProjectCommand) SetAutoscalingFlavor(v string)`

SetAutoscalingFlavor sets AutoscalingFlavor field to given value.

### HasAutoscalingFlavor

`func (o *CreateProjectCommand) HasAutoscalingFlavor() bool`

HasAutoscalingFlavor returns a boolean if a field has been set.

### SetAutoscalingFlavorNil

`func (o *CreateProjectCommand) SetAutoscalingFlavorNil(b bool)`

 SetAutoscalingFlavorNil sets the value for AutoscalingFlavor to be an explicit nil

### UnsetAutoscalingFlavor
`func (o *CreateProjectCommand) UnsetAutoscalingFlavor()`

UnsetAutoscalingFlavor ensures that no value is present for AutoscalingFlavor, not even an explicit nil
### GetAutoscalingSpotEnabled

`func (o *CreateProjectCommand) GetAutoscalingSpotEnabled() bool`

GetAutoscalingSpotEnabled returns the AutoscalingSpotEnabled field if non-nil, zero value otherwise.

### GetAutoscalingSpotEnabledOk

`func (o *CreateProjectCommand) GetAutoscalingSpotEnabledOk() (*bool, bool)`

GetAutoscalingSpotEnabledOk returns a tuple with the AutoscalingSpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingSpotEnabled

`func (o *CreateProjectCommand) SetAutoscalingSpotEnabled(v bool)`

SetAutoscalingSpotEnabled sets AutoscalingSpotEnabled field to given value.

### HasAutoscalingSpotEnabled

`func (o *CreateProjectCommand) HasAutoscalingSpotEnabled() bool`

HasAutoscalingSpotEnabled returns a boolean if a field has been set.

### GetIsCskEnabled

`func (o *CreateProjectCommand) GetIsCskEnabled() bool`

GetIsCskEnabled returns the IsCskEnabled field if non-nil, zero value otherwise.

### GetIsCskEnabledOk

`func (o *CreateProjectCommand) GetIsCskEnabledOk() (*bool, bool)`

GetIsCskEnabledOk returns a tuple with the IsCskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCskEnabled

`func (o *CreateProjectCommand) SetIsCskEnabled(v bool)`

SetIsCskEnabled sets IsCskEnabled field to given value.

### HasIsCskEnabled

`func (o *CreateProjectCommand) HasIsCskEnabled() bool`

HasIsCskEnabled returns a boolean if a field has been set.

### GetCidr

`func (o *CreateProjectCommand) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateProjectCommand) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateProjectCommand) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateProjectCommand) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### SetCidrNil

`func (o *CreateProjectCommand) SetCidrNil(b bool)`

 SetCidrNil sets the value for Cidr to be an explicit nil

### UnsetCidr
`func (o *CreateProjectCommand) UnsetCidr()`

UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
### GetNetMask

`func (o *CreateProjectCommand) GetNetMask() int32`

GetNetMask returns the NetMask field if non-nil, zero value otherwise.

### GetNetMaskOk

`func (o *CreateProjectCommand) GetNetMaskOk() (*int32, bool)`

GetNetMaskOk returns a tuple with the NetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetMask

`func (o *CreateProjectCommand) SetNetMask(v int32)`

SetNetMask sets NetMask field to given value.

### HasNetMask

`func (o *CreateProjectCommand) HasNetMask() bool`

HasNetMask returns a boolean if a field has been set.

### SetNetMaskNil

`func (o *CreateProjectCommand) SetNetMaskNil(b bool)`

 SetNetMaskNil sets the value for NetMask to be an explicit nil

### UnsetNetMask
`func (o *CreateProjectCommand) UnsetNetMask()`

UnsetNetMask ensures that no value is present for NetMask, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


