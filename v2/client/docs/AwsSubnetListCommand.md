# AwsSubnetListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKeyId** | **NullableString** |  | 
**AwsSecretAccessKey** | **NullableString** |  | 
**RoleArn** | **NullableString** |  | 
**ExternalId** | **NullableString** |  | 
**AwsRegion** | **NullableString** |  | 
**VpcId** | **NullableString** |  | 

## Methods

### NewAwsSubnetListCommand

`func NewAwsSubnetListCommand(awsAccessKeyId NullableString, awsSecretAccessKey NullableString, roleArn NullableString, externalId NullableString, awsRegion NullableString, vpcId NullableString, ) *AwsSubnetListCommand`

NewAwsSubnetListCommand instantiates a new AwsSubnetListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSubnetListCommandWithDefaults

`func NewAwsSubnetListCommandWithDefaults() *AwsSubnetListCommand`

NewAwsSubnetListCommandWithDefaults instantiates a new AwsSubnetListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKeyId

`func (o *AwsSubnetListCommand) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *AwsSubnetListCommand) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *AwsSubnetListCommand) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### SetAwsAccessKeyIdNil

`func (o *AwsSubnetListCommand) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *AwsSubnetListCommand) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *AwsSubnetListCommand) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AwsSubnetListCommand) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AwsSubnetListCommand) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### SetAwsSecretAccessKeyNil

`func (o *AwsSubnetListCommand) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *AwsSubnetListCommand) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetRoleArn

`func (o *AwsSubnetListCommand) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsSubnetListCommand) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsSubnetListCommand) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### SetRoleArnNil

`func (o *AwsSubnetListCommand) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AwsSubnetListCommand) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetExternalId

`func (o *AwsSubnetListCommand) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsSubnetListCommand) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsSubnetListCommand) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *AwsSubnetListCommand) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsSubnetListCommand) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAwsRegion

`func (o *AwsSubnetListCommand) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsSubnetListCommand) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsSubnetListCommand) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### SetAwsRegionNil

`func (o *AwsSubnetListCommand) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsSubnetListCommand) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetVpcId

`func (o *AwsSubnetListCommand) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AwsSubnetListCommand) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AwsSubnetListCommand) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *AwsSubnetListCommand) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *AwsSubnetListCommand) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


