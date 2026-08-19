# DeploymentEnableBackupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int32** |  | 
**S3CredentialId** | **int32** |  | 

## Methods

### NewDeploymentEnableBackupCommand

`func NewDeploymentEnableBackupCommand(projectId int32, s3CredentialId int32, ) *DeploymentEnableBackupCommand`

NewDeploymentEnableBackupCommand instantiates a new DeploymentEnableBackupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentEnableBackupCommandWithDefaults

`func NewDeploymentEnableBackupCommandWithDefaults() *DeploymentEnableBackupCommand`

NewDeploymentEnableBackupCommandWithDefaults instantiates a new DeploymentEnableBackupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeploymentEnableBackupCommand) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeploymentEnableBackupCommand) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeploymentEnableBackupCommand) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetS3CredentialId

`func (o *DeploymentEnableBackupCommand) GetS3CredentialId() int32`

GetS3CredentialId returns the S3CredentialId field if non-nil, zero value otherwise.

### GetS3CredentialIdOk

`func (o *DeploymentEnableBackupCommand) GetS3CredentialIdOk() (*int32, bool)`

GetS3CredentialIdOk returns a tuple with the S3CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3CredentialId

`func (o *DeploymentEnableBackupCommand) SetS3CredentialId(v int32)`

SetS3CredentialId sets S3CredentialId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


