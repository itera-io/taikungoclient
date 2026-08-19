# ProjectDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **string** |  | 
**CloudType** | [**ECloudCredentialType**](ECloudCredentialType.md) |  | 
**Status** | [**ProjectStatus**](ProjectStatus.md) |  | 
**Health** | [**ProjectHealth**](ProjectHealth.md) |  | 
**AlertCount** | **int32** |  | 
**CurrentKubernetesVersion** | **string** |  | 
**ModifiedAt** | **string** |  | 
**ModifiedBy** | **string** |  | 
**CreatedAt** | **string** |  | 
**ExpiredAt** | **string** |  | 
**HasKubeConfigFile** | **bool** |  | 
**Users** | [**[]GroupedOrganizationUserDto**](GroupedOrganizationUserDto.md) |  | 

## Methods

### NewProjectDetailsDto

`func NewProjectDetailsDto(id int32, name string, organizationId int32, organizationName string, cloudType ECloudCredentialType, status ProjectStatus, health ProjectHealth, alertCount int32, currentKubernetesVersion string, modifiedAt string, modifiedBy string, createdAt string, expiredAt string, hasKubeConfigFile bool, users []GroupedOrganizationUserDto, ) *ProjectDetailsDto`

NewProjectDetailsDto instantiates a new ProjectDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsDtoWithDefaults

`func NewProjectDetailsDtoWithDefaults() *ProjectDetailsDto`

NewProjectDetailsDtoWithDefaults instantiates a new ProjectDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectDetailsDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectDetailsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectDetailsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectDetailsDto) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ProjectDetailsDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectDetailsDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectDetailsDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *ProjectDetailsDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectDetailsDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectDetailsDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### GetCloudType

`func (o *ProjectDetailsDto) GetCloudType() ECloudCredentialType`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *ProjectDetailsDto) GetCloudTypeOk() (*ECloudCredentialType, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *ProjectDetailsDto) SetCloudType(v ECloudCredentialType)`

SetCloudType sets CloudType field to given value.


### GetStatus

`func (o *ProjectDetailsDto) GetStatus() ProjectStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectDetailsDto) GetStatusOk() (*ProjectStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectDetailsDto) SetStatus(v ProjectStatus)`

SetStatus sets Status field to given value.


### GetHealth

`func (o *ProjectDetailsDto) GetHealth() ProjectHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ProjectDetailsDto) GetHealthOk() (*ProjectHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ProjectDetailsDto) SetHealth(v ProjectHealth)`

SetHealth sets Health field to given value.


### GetAlertCount

`func (o *ProjectDetailsDto) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *ProjectDetailsDto) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *ProjectDetailsDto) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.


### GetCurrentKubernetesVersion

`func (o *ProjectDetailsDto) GetCurrentKubernetesVersion() string`

GetCurrentKubernetesVersion returns the CurrentKubernetesVersion field if non-nil, zero value otherwise.

### GetCurrentKubernetesVersionOk

`func (o *ProjectDetailsDto) GetCurrentKubernetesVersionOk() (*string, bool)`

GetCurrentKubernetesVersionOk returns a tuple with the CurrentKubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKubernetesVersion

`func (o *ProjectDetailsDto) SetCurrentKubernetesVersion(v string)`

SetCurrentKubernetesVersion sets CurrentKubernetesVersion field to given value.


### GetModifiedAt

`func (o *ProjectDetailsDto) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ProjectDetailsDto) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ProjectDetailsDto) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ProjectDetailsDto) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ProjectDetailsDto) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ProjectDetailsDto) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetCreatedAt

`func (o *ProjectDetailsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectDetailsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectDetailsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiredAt

`func (o *ProjectDetailsDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectDetailsDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectDetailsDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### GetHasKubeConfigFile

`func (o *ProjectDetailsDto) GetHasKubeConfigFile() bool`

GetHasKubeConfigFile returns the HasKubeConfigFile field if non-nil, zero value otherwise.

### GetHasKubeConfigFileOk

`func (o *ProjectDetailsDto) GetHasKubeConfigFileOk() (*bool, bool)`

GetHasKubeConfigFileOk returns a tuple with the HasKubeConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKubeConfigFile

`func (o *ProjectDetailsDto) SetHasKubeConfigFile(v bool)`

SetHasKubeConfigFile sets HasKubeConfigFile field to given value.


### GetUsers

`func (o *ProjectDetailsDto) GetUsers() []GroupedOrganizationUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProjectDetailsDto) GetUsersOk() (*[]GroupedOrganizationUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProjectDetailsDto) SetUsers(v []GroupedOrganizationUserDto)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


