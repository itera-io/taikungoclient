# AwsEksNodeGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**Role** | **NullableString** |  | 
**AmiType** | **NullableString** |  | 
**AmiReleaseVersion** | **NullableString** |  | 
**Arn** | **NullableString** |  | 
**CapacityType** | **NullableString** |  | 
**HealthIssues** | **[]string** |  | 
**Status** | **NullableString** |  | 
**Version** | **NullableString** |  | 
**Created** | **NullableString** |  | 
**DiskSize** | **int32** |  | 
**DesiredSize** | **int32** |  | 
**MaxSize** | **int32** |  | 
**MinSize** | **int32** |  | 
**InstanceTypes** | **[]string** |  | 

## Methods

### NewAwsEksNodeGroupDto

`func NewAwsEksNodeGroupDto(name NullableString, role NullableString, amiType NullableString, amiReleaseVersion NullableString, arn NullableString, capacityType NullableString, healthIssues []string, status NullableString, version NullableString, created NullableString, diskSize int32, desiredSize int32, maxSize int32, minSize int32, instanceTypes []string, ) *AwsEksNodeGroupDto`

NewAwsEksNodeGroupDto instantiates a new AwsEksNodeGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEksNodeGroupDtoWithDefaults

`func NewAwsEksNodeGroupDtoWithDefaults() *AwsEksNodeGroupDto`

NewAwsEksNodeGroupDtoWithDefaults instantiates a new AwsEksNodeGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AwsEksNodeGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsEksNodeGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsEksNodeGroupDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AwsEksNodeGroupDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AwsEksNodeGroupDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRole

`func (o *AwsEksNodeGroupDto) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AwsEksNodeGroupDto) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AwsEksNodeGroupDto) SetRole(v string)`

SetRole sets Role field to given value.


### SetRoleNil

`func (o *AwsEksNodeGroupDto) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AwsEksNodeGroupDto) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetAmiType

`func (o *AwsEksNodeGroupDto) GetAmiType() string`

GetAmiType returns the AmiType field if non-nil, zero value otherwise.

### GetAmiTypeOk

`func (o *AwsEksNodeGroupDto) GetAmiTypeOk() (*string, bool)`

GetAmiTypeOk returns a tuple with the AmiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiType

`func (o *AwsEksNodeGroupDto) SetAmiType(v string)`

SetAmiType sets AmiType field to given value.


### SetAmiTypeNil

`func (o *AwsEksNodeGroupDto) SetAmiTypeNil(b bool)`

 SetAmiTypeNil sets the value for AmiType to be an explicit nil

### UnsetAmiType
`func (o *AwsEksNodeGroupDto) UnsetAmiType()`

UnsetAmiType ensures that no value is present for AmiType, not even an explicit nil
### GetAmiReleaseVersion

`func (o *AwsEksNodeGroupDto) GetAmiReleaseVersion() string`

GetAmiReleaseVersion returns the AmiReleaseVersion field if non-nil, zero value otherwise.

### GetAmiReleaseVersionOk

`func (o *AwsEksNodeGroupDto) GetAmiReleaseVersionOk() (*string, bool)`

GetAmiReleaseVersionOk returns a tuple with the AmiReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiReleaseVersion

`func (o *AwsEksNodeGroupDto) SetAmiReleaseVersion(v string)`

SetAmiReleaseVersion sets AmiReleaseVersion field to given value.


### SetAmiReleaseVersionNil

`func (o *AwsEksNodeGroupDto) SetAmiReleaseVersionNil(b bool)`

 SetAmiReleaseVersionNil sets the value for AmiReleaseVersion to be an explicit nil

### UnsetAmiReleaseVersion
`func (o *AwsEksNodeGroupDto) UnsetAmiReleaseVersion()`

UnsetAmiReleaseVersion ensures that no value is present for AmiReleaseVersion, not even an explicit nil
### GetArn

`func (o *AwsEksNodeGroupDto) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### GetArnOk

`func (o *AwsEksNodeGroupDto) GetArnOk() (*string, bool)`

GetArnOk returns a tuple with the Arn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArn

`func (o *AwsEksNodeGroupDto) SetArn(v string)`

SetArn sets Arn field to given value.


### SetArnNil

`func (o *AwsEksNodeGroupDto) SetArnNil(b bool)`

 SetArnNil sets the value for Arn to be an explicit nil

### UnsetArn
`func (o *AwsEksNodeGroupDto) UnsetArn()`

UnsetArn ensures that no value is present for Arn, not even an explicit nil
### GetCapacityType

`func (o *AwsEksNodeGroupDto) GetCapacityType() string`

GetCapacityType returns the CapacityType field if non-nil, zero value otherwise.

### GetCapacityTypeOk

`func (o *AwsEksNodeGroupDto) GetCapacityTypeOk() (*string, bool)`

GetCapacityTypeOk returns a tuple with the CapacityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityType

`func (o *AwsEksNodeGroupDto) SetCapacityType(v string)`

SetCapacityType sets CapacityType field to given value.


### SetCapacityTypeNil

`func (o *AwsEksNodeGroupDto) SetCapacityTypeNil(b bool)`

 SetCapacityTypeNil sets the value for CapacityType to be an explicit nil

### UnsetCapacityType
`func (o *AwsEksNodeGroupDto) UnsetCapacityType()`

UnsetCapacityType ensures that no value is present for CapacityType, not even an explicit nil
### GetHealthIssues

`func (o *AwsEksNodeGroupDto) GetHealthIssues() []string`

GetHealthIssues returns the HealthIssues field if non-nil, zero value otherwise.

### GetHealthIssuesOk

`func (o *AwsEksNodeGroupDto) GetHealthIssuesOk() (*[]string, bool)`

GetHealthIssuesOk returns a tuple with the HealthIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthIssues

`func (o *AwsEksNodeGroupDto) SetHealthIssues(v []string)`

SetHealthIssues sets HealthIssues field to given value.


### SetHealthIssuesNil

`func (o *AwsEksNodeGroupDto) SetHealthIssuesNil(b bool)`

 SetHealthIssuesNil sets the value for HealthIssues to be an explicit nil

### UnsetHealthIssues
`func (o *AwsEksNodeGroupDto) UnsetHealthIssues()`

UnsetHealthIssues ensures that no value is present for HealthIssues, not even an explicit nil
### GetStatus

`func (o *AwsEksNodeGroupDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsEksNodeGroupDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsEksNodeGroupDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *AwsEksNodeGroupDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AwsEksNodeGroupDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVersion

`func (o *AwsEksNodeGroupDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AwsEksNodeGroupDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AwsEksNodeGroupDto) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *AwsEksNodeGroupDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *AwsEksNodeGroupDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetCreated

`func (o *AwsEksNodeGroupDto) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AwsEksNodeGroupDto) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AwsEksNodeGroupDto) SetCreated(v string)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *AwsEksNodeGroupDto) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *AwsEksNodeGroupDto) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetDiskSize

`func (o *AwsEksNodeGroupDto) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *AwsEksNodeGroupDto) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *AwsEksNodeGroupDto) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.


### GetDesiredSize

`func (o *AwsEksNodeGroupDto) GetDesiredSize() int32`

GetDesiredSize returns the DesiredSize field if non-nil, zero value otherwise.

### GetDesiredSizeOk

`func (o *AwsEksNodeGroupDto) GetDesiredSizeOk() (*int32, bool)`

GetDesiredSizeOk returns a tuple with the DesiredSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredSize

`func (o *AwsEksNodeGroupDto) SetDesiredSize(v int32)`

SetDesiredSize sets DesiredSize field to given value.


### GetMaxSize

`func (o *AwsEksNodeGroupDto) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *AwsEksNodeGroupDto) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *AwsEksNodeGroupDto) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.


### GetMinSize

`func (o *AwsEksNodeGroupDto) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *AwsEksNodeGroupDto) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *AwsEksNodeGroupDto) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.


### GetInstanceTypes

`func (o *AwsEksNodeGroupDto) GetInstanceTypes() []string`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *AwsEksNodeGroupDto) GetInstanceTypesOk() (*[]string, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *AwsEksNodeGroupDto) SetInstanceTypes(v []string)`

SetInstanceTypes sets InstanceTypes field to given value.


### SetInstanceTypesNil

`func (o *AwsEksNodeGroupDto) SetInstanceTypesNil(b bool)`

 SetInstanceTypesNil sets the value for InstanceTypes to be an explicit nil

### UnsetInstanceTypes
`func (o *AwsEksNodeGroupDto) UnsetInstanceTypes()`

UnsetInstanceTypes ensures that no value is present for InstanceTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


