# ProjectDescribeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**ProjectName** | **NullableString** |  | 
**IsMonitoringEnabled** | **bool** |  | 
**IsBackupEnabled** | **bool** |  | 
**IsAutoUpgradeEnabled** | **bool** |  | 
**ServersCount** | **int32** |  | 
**VmsCount** | **int32** |  | 
**AlertsCount** | **int32** |  | 
**LastModifiedBy** | **NullableString** |  | 
**CreatedBy** | **NullableString** |  | 
**LastModified** | **NullableString** |  | 
**Flavors** | **[]string** |  | 
**Images** | **[]string** |  | 
**AccessIp** | **NullableString** |  | 
**FailureReason** | **NullableString** |  | 

## Methods

### NewProjectDescribeDto

`func NewProjectDescribeDto(projectId int32, projectName NullableString, isMonitoringEnabled bool, isBackupEnabled bool, isAutoUpgradeEnabled bool, serversCount int32, vmsCount int32, alertsCount int32, lastModifiedBy NullableString, createdBy NullableString, lastModified NullableString, flavors []string, images []string, accessIp NullableString, failureReason NullableString, ) *ProjectDescribeDto`

NewProjectDescribeDto instantiates a new ProjectDescribeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDescribeDtoWithDefaults

`func NewProjectDescribeDtoWithDefaults() *ProjectDescribeDto`

NewProjectDescribeDtoWithDefaults instantiates a new ProjectDescribeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ProjectDescribeDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectDescribeDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectDescribeDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *ProjectDescribeDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ProjectDescribeDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ProjectDescribeDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### SetProjectNameNil

`func (o *ProjectDescribeDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ProjectDescribeDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetIsMonitoringEnabled

`func (o *ProjectDescribeDto) GetIsMonitoringEnabled() bool`

GetIsMonitoringEnabled returns the IsMonitoringEnabled field if non-nil, zero value otherwise.

### GetIsMonitoringEnabledOk

`func (o *ProjectDescribeDto) GetIsMonitoringEnabledOk() (*bool, bool)`

GetIsMonitoringEnabledOk returns a tuple with the IsMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonitoringEnabled

`func (o *ProjectDescribeDto) SetIsMonitoringEnabled(v bool)`

SetIsMonitoringEnabled sets IsMonitoringEnabled field to given value.


### GetIsBackupEnabled

`func (o *ProjectDescribeDto) GetIsBackupEnabled() bool`

GetIsBackupEnabled returns the IsBackupEnabled field if non-nil, zero value otherwise.

### GetIsBackupEnabledOk

`func (o *ProjectDescribeDto) GetIsBackupEnabledOk() (*bool, bool)`

GetIsBackupEnabledOk returns a tuple with the IsBackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBackupEnabled

`func (o *ProjectDescribeDto) SetIsBackupEnabled(v bool)`

SetIsBackupEnabled sets IsBackupEnabled field to given value.


### GetIsAutoUpgradeEnabled

`func (o *ProjectDescribeDto) GetIsAutoUpgradeEnabled() bool`

GetIsAutoUpgradeEnabled returns the IsAutoUpgradeEnabled field if non-nil, zero value otherwise.

### GetIsAutoUpgradeEnabledOk

`func (o *ProjectDescribeDto) GetIsAutoUpgradeEnabledOk() (*bool, bool)`

GetIsAutoUpgradeEnabledOk returns a tuple with the IsAutoUpgradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpgradeEnabled

`func (o *ProjectDescribeDto) SetIsAutoUpgradeEnabled(v bool)`

SetIsAutoUpgradeEnabled sets IsAutoUpgradeEnabled field to given value.


### GetServersCount

`func (o *ProjectDescribeDto) GetServersCount() int32`

GetServersCount returns the ServersCount field if non-nil, zero value otherwise.

### GetServersCountOk

`func (o *ProjectDescribeDto) GetServersCountOk() (*int32, bool)`

GetServersCountOk returns a tuple with the ServersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersCount

`func (o *ProjectDescribeDto) SetServersCount(v int32)`

SetServersCount sets ServersCount field to given value.


### GetVmsCount

`func (o *ProjectDescribeDto) GetVmsCount() int32`

GetVmsCount returns the VmsCount field if non-nil, zero value otherwise.

### GetVmsCountOk

`func (o *ProjectDescribeDto) GetVmsCountOk() (*int32, bool)`

GetVmsCountOk returns a tuple with the VmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmsCount

`func (o *ProjectDescribeDto) SetVmsCount(v int32)`

SetVmsCount sets VmsCount field to given value.


### GetAlertsCount

`func (o *ProjectDescribeDto) GetAlertsCount() int32`

GetAlertsCount returns the AlertsCount field if non-nil, zero value otherwise.

### GetAlertsCountOk

`func (o *ProjectDescribeDto) GetAlertsCountOk() (*int32, bool)`

GetAlertsCountOk returns a tuple with the AlertsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsCount

`func (o *ProjectDescribeDto) SetAlertsCount(v int32)`

SetAlertsCount sets AlertsCount field to given value.


### GetLastModifiedBy

`func (o *ProjectDescribeDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ProjectDescribeDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ProjectDescribeDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### SetLastModifiedByNil

`func (o *ProjectDescribeDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *ProjectDescribeDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
### GetCreatedBy

`func (o *ProjectDescribeDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ProjectDescribeDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ProjectDescribeDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *ProjectDescribeDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ProjectDescribeDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *ProjectDescribeDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ProjectDescribeDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ProjectDescribeDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *ProjectDescribeDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ProjectDescribeDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetFlavors

`func (o *ProjectDescribeDto) GetFlavors() []string`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *ProjectDescribeDto) GetFlavorsOk() (*[]string, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *ProjectDescribeDto) SetFlavors(v []string)`

SetFlavors sets Flavors field to given value.


### SetFlavorsNil

`func (o *ProjectDescribeDto) SetFlavorsNil(b bool)`

 SetFlavorsNil sets the value for Flavors to be an explicit nil

### UnsetFlavors
`func (o *ProjectDescribeDto) UnsetFlavors()`

UnsetFlavors ensures that no value is present for Flavors, not even an explicit nil
### GetImages

`func (o *ProjectDescribeDto) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ProjectDescribeDto) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ProjectDescribeDto) SetImages(v []string)`

SetImages sets Images field to given value.


### SetImagesNil

`func (o *ProjectDescribeDto) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *ProjectDescribeDto) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetAccessIp

`func (o *ProjectDescribeDto) GetAccessIp() string`

GetAccessIp returns the AccessIp field if non-nil, zero value otherwise.

### GetAccessIpOk

`func (o *ProjectDescribeDto) GetAccessIpOk() (*string, bool)`

GetAccessIpOk returns a tuple with the AccessIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIp

`func (o *ProjectDescribeDto) SetAccessIp(v string)`

SetAccessIp sets AccessIp field to given value.


### SetAccessIpNil

`func (o *ProjectDescribeDto) SetAccessIpNil(b bool)`

 SetAccessIpNil sets the value for AccessIp to be an explicit nil

### UnsetAccessIp
`func (o *ProjectDescribeDto) UnsetAccessIp()`

UnsetAccessIp ensures that no value is present for AccessIp, not even an explicit nil
### GetFailureReason

`func (o *ProjectDescribeDto) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProjectDescribeDto) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProjectDescribeDto) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *ProjectDescribeDto) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *ProjectDescribeDto) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


