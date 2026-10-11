# CreateAwsCloudCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** |  | 
**AwsSecretAccessKey** | Pointer to **NullableString** |  | [optional] 
**AwsAccessKeyId** | Pointer to **NullableString** |  | [optional] 
**RoleArn** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**AzCount** | Pointer to **NullableInt32** |  | [optional] 
**AwsRegion** | **NullableString** |  | 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**ImportedVpcId** | Pointer to **NullableString** |  | [optional] 
**PrivateOnly** | Pointer to **bool** |  | [optional] 
**IpMode** | Pointer to [**IpMode**](IpMode.md) |  | [optional] 
**VpcMode** | Pointer to **map[string]interface{}** |  | [optional] 
**Subnets** | Pointer to [**[]CreateAwsSubnetDto**](CreateAwsSubnetDto.md) |  | [optional] 
**SharedFileSystemEnabled** | Pointer to **bool** |  | [optional] 
**EfsFileSystemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateAwsCloudCommand

`func NewCreateAwsCloudCommand(name NullableString, awsRegion NullableString, ) *CreateAwsCloudCommand`

NewCreateAwsCloudCommand instantiates a new CreateAwsCloudCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAwsCloudCommandWithDefaults

`func NewCreateAwsCloudCommandWithDefaults() *CreateAwsCloudCommand`

NewCreateAwsCloudCommandWithDefaults instantiates a new CreateAwsCloudCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAwsCloudCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAwsCloudCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAwsCloudCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateAwsCloudCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAwsCloudCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *CreateAwsCloudCommand) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *CreateAwsCloudCommand) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *CreateAwsCloudCommand) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *CreateAwsCloudCommand) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### SetAwsSecretAccessKeyNil

`func (o *CreateAwsCloudCommand) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *CreateAwsCloudCommand) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetAwsAccessKeyId

`func (o *CreateAwsCloudCommand) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *CreateAwsCloudCommand) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *CreateAwsCloudCommand) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *CreateAwsCloudCommand) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### SetAwsAccessKeyIdNil

`func (o *CreateAwsCloudCommand) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *CreateAwsCloudCommand) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetRoleArn

`func (o *CreateAwsCloudCommand) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *CreateAwsCloudCommand) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *CreateAwsCloudCommand) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *CreateAwsCloudCommand) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### SetRoleArnNil

`func (o *CreateAwsCloudCommand) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *CreateAwsCloudCommand) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetExternalId

`func (o *CreateAwsCloudCommand) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateAwsCloudCommand) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateAwsCloudCommand) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateAwsCloudCommand) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CreateAwsCloudCommand) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CreateAwsCloudCommand) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAzCount

`func (o *CreateAwsCloudCommand) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *CreateAwsCloudCommand) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *CreateAwsCloudCommand) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *CreateAwsCloudCommand) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### SetAzCountNil

`func (o *CreateAwsCloudCommand) SetAzCountNil(b bool)`

 SetAzCountNil sets the value for AzCount to be an explicit nil

### UnsetAzCount
`func (o *CreateAwsCloudCommand) UnsetAzCount()`

UnsetAzCount ensures that no value is present for AzCount, not even an explicit nil
### GetAwsRegion

`func (o *CreateAwsCloudCommand) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *CreateAwsCloudCommand) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *CreateAwsCloudCommand) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### SetAwsRegionNil

`func (o *CreateAwsCloudCommand) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *CreateAwsCloudCommand) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetOrganizationId

`func (o *CreateAwsCloudCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAwsCloudCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAwsCloudCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAwsCloudCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateAwsCloudCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAwsCloudCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetImportedVpcId

`func (o *CreateAwsCloudCommand) GetImportedVpcId() string`

GetImportedVpcId returns the ImportedVpcId field if non-nil, zero value otherwise.

### GetImportedVpcIdOk

`func (o *CreateAwsCloudCommand) GetImportedVpcIdOk() (*string, bool)`

GetImportedVpcIdOk returns a tuple with the ImportedVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedVpcId

`func (o *CreateAwsCloudCommand) SetImportedVpcId(v string)`

SetImportedVpcId sets ImportedVpcId field to given value.

### HasImportedVpcId

`func (o *CreateAwsCloudCommand) HasImportedVpcId() bool`

HasImportedVpcId returns a boolean if a field has been set.

### SetImportedVpcIdNil

`func (o *CreateAwsCloudCommand) SetImportedVpcIdNil(b bool)`

 SetImportedVpcIdNil sets the value for ImportedVpcId to be an explicit nil

### UnsetImportedVpcId
`func (o *CreateAwsCloudCommand) UnsetImportedVpcId()`

UnsetImportedVpcId ensures that no value is present for ImportedVpcId, not even an explicit nil
### GetPrivateOnly

`func (o *CreateAwsCloudCommand) GetPrivateOnly() bool`

GetPrivateOnly returns the PrivateOnly field if non-nil, zero value otherwise.

### GetPrivateOnlyOk

`func (o *CreateAwsCloudCommand) GetPrivateOnlyOk() (*bool, bool)`

GetPrivateOnlyOk returns a tuple with the PrivateOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOnly

`func (o *CreateAwsCloudCommand) SetPrivateOnly(v bool)`

SetPrivateOnly sets PrivateOnly field to given value.

### HasPrivateOnly

`func (o *CreateAwsCloudCommand) HasPrivateOnly() bool`

HasPrivateOnly returns a boolean if a field has been set.

### GetIpMode

`func (o *CreateAwsCloudCommand) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *CreateAwsCloudCommand) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *CreateAwsCloudCommand) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *CreateAwsCloudCommand) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetVpcMode

`func (o *CreateAwsCloudCommand) GetVpcMode() map[string]interface{}`

GetVpcMode returns the VpcMode field if non-nil, zero value otherwise.

### GetVpcModeOk

`func (o *CreateAwsCloudCommand) GetVpcModeOk() (*map[string]interface{}, bool)`

GetVpcModeOk returns a tuple with the VpcMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcMode

`func (o *CreateAwsCloudCommand) SetVpcMode(v map[string]interface{})`

SetVpcMode sets VpcMode field to given value.

### HasVpcMode

`func (o *CreateAwsCloudCommand) HasVpcMode() bool`

HasVpcMode returns a boolean if a field has been set.

### GetSubnets

`func (o *CreateAwsCloudCommand) GetSubnets() []CreateAwsSubnetDto`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateAwsCloudCommand) GetSubnetsOk() (*[]CreateAwsSubnetDto, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateAwsCloudCommand) SetSubnets(v []CreateAwsSubnetDto)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateAwsCloudCommand) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *CreateAwsCloudCommand) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *CreateAwsCloudCommand) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil
### GetSharedFileSystemEnabled

`func (o *CreateAwsCloudCommand) GetSharedFileSystemEnabled() bool`

GetSharedFileSystemEnabled returns the SharedFileSystemEnabled field if non-nil, zero value otherwise.

### GetSharedFileSystemEnabledOk

`func (o *CreateAwsCloudCommand) GetSharedFileSystemEnabledOk() (*bool, bool)`

GetSharedFileSystemEnabledOk returns a tuple with the SharedFileSystemEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFileSystemEnabled

`func (o *CreateAwsCloudCommand) SetSharedFileSystemEnabled(v bool)`

SetSharedFileSystemEnabled sets SharedFileSystemEnabled field to given value.

### HasSharedFileSystemEnabled

`func (o *CreateAwsCloudCommand) HasSharedFileSystemEnabled() bool`

HasSharedFileSystemEnabled returns a boolean if a field has been set.

### GetEfsFileSystemId

`func (o *CreateAwsCloudCommand) GetEfsFileSystemId() string`

GetEfsFileSystemId returns the EfsFileSystemId field if non-nil, zero value otherwise.

### GetEfsFileSystemIdOk

`func (o *CreateAwsCloudCommand) GetEfsFileSystemIdOk() (*string, bool)`

GetEfsFileSystemIdOk returns a tuple with the EfsFileSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfsFileSystemId

`func (o *CreateAwsCloudCommand) SetEfsFileSystemId(v string)`

SetEfsFileSystemId sets EfsFileSystemId field to given value.

### HasEfsFileSystemId

`func (o *CreateAwsCloudCommand) HasEfsFileSystemId() bool`

HasEfsFileSystemId returns a boolean if a field has been set.

### SetEfsFileSystemIdNil

`func (o *CreateAwsCloudCommand) SetEfsFileSystemIdNil(b bool)`

 SetEfsFileSystemIdNil sets the value for EfsFileSystemId to be an explicit nil

### UnsetEfsFileSystemId
`func (o *CreateAwsCloudCommand) UnsetEfsFileSystemId()`

UnsetEfsFileSystemId ensures that no value is present for EfsFileSystemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


