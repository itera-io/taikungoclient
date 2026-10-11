# UpdateAwsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**AwsSecretAccessKey** | **NullableString** |  | 
**AwsAccessKeyId** | **NullableString** |  | 
**RoleArn** | **NullableString** |  | 
**ExternalId** | **NullableString** |  | 

## Methods

### NewUpdateAwsCommand

`func NewUpdateAwsCommand(id int32, name NullableString, awsSecretAccessKey NullableString, awsAccessKeyId NullableString, roleArn NullableString, externalId NullableString, ) *UpdateAwsCommand`

NewUpdateAwsCommand instantiates a new UpdateAwsCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAwsCommandWithDefaults

`func NewUpdateAwsCommandWithDefaults() *UpdateAwsCommand`

NewUpdateAwsCommandWithDefaults instantiates a new UpdateAwsCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAwsCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAwsCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAwsCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAwsCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAwsCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAwsCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateAwsCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAwsCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *UpdateAwsCommand) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *UpdateAwsCommand) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *UpdateAwsCommand) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### SetAwsSecretAccessKeyNil

`func (o *UpdateAwsCommand) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *UpdateAwsCommand) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetAwsAccessKeyId

`func (o *UpdateAwsCommand) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *UpdateAwsCommand) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *UpdateAwsCommand) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### SetAwsAccessKeyIdNil

`func (o *UpdateAwsCommand) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *UpdateAwsCommand) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetRoleArn

`func (o *UpdateAwsCommand) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *UpdateAwsCommand) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *UpdateAwsCommand) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### SetRoleArnNil

`func (o *UpdateAwsCommand) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *UpdateAwsCommand) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetExternalId

`func (o *UpdateAwsCommand) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateAwsCommand) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateAwsCommand) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *UpdateAwsCommand) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateAwsCommand) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


