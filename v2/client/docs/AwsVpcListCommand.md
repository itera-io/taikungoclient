# AwsVpcListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKeyId** | **NullableString** |  | 
**AwsSecretAccessKey** | **NullableString** |  | 
**RoleArn** | **NullableString** |  | 
**ExternalId** | **NullableString** |  | 
**AwsRegion** | **NullableString** |  | 
**IpMode** | [**IpMode**](IpMode.md) |  | 

## Methods

### NewAwsVpcListCommand

`func NewAwsVpcListCommand(awsAccessKeyId NullableString, awsSecretAccessKey NullableString, roleArn NullableString, externalId NullableString, awsRegion NullableString, ipMode IpMode, ) *AwsVpcListCommand`

NewAwsVpcListCommand instantiates a new AwsVpcListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsVpcListCommandWithDefaults

`func NewAwsVpcListCommandWithDefaults() *AwsVpcListCommand`

NewAwsVpcListCommandWithDefaults instantiates a new AwsVpcListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKeyId

`func (o *AwsVpcListCommand) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *AwsVpcListCommand) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *AwsVpcListCommand) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### SetAwsAccessKeyIdNil

`func (o *AwsVpcListCommand) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *AwsVpcListCommand) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *AwsVpcListCommand) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AwsVpcListCommand) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AwsVpcListCommand) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### SetAwsSecretAccessKeyNil

`func (o *AwsVpcListCommand) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *AwsVpcListCommand) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetRoleArn

`func (o *AwsVpcListCommand) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AwsVpcListCommand) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AwsVpcListCommand) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### SetRoleArnNil

`func (o *AwsVpcListCommand) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AwsVpcListCommand) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetExternalId

`func (o *AwsVpcListCommand) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AwsVpcListCommand) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AwsVpcListCommand) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *AwsVpcListCommand) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AwsVpcListCommand) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetAwsRegion

`func (o *AwsVpcListCommand) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *AwsVpcListCommand) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *AwsVpcListCommand) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.


### SetAwsRegionNil

`func (o *AwsVpcListCommand) SetAwsRegionNil(b bool)`

 SetAwsRegionNil sets the value for AwsRegion to be an explicit nil

### UnsetAwsRegion
`func (o *AwsVpcListCommand) UnsetAwsRegion()`

UnsetAwsRegion ensures that no value is present for AwsRegion, not even an explicit nil
### GetIpMode

`func (o *AwsVpcListCommand) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AwsVpcListCommand) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AwsVpcListCommand) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


