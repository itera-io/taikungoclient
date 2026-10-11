# AmazonAvailabilityZonesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **NullableString** |  | 
**AwsAccessKeyId** | **NullableString** |  | 
**AwsSecretAccessKey** | **NullableString** |  | 
**RoleArn** | **NullableString** |  | 
**ExternalId** | **NullableString** |  | 
**CloudId** | **NullableInt32** |  | 

## Methods

### NewAmazonAvailabilityZonesCommand

`func NewAmazonAvailabilityZonesCommand(region NullableString, awsAccessKeyId NullableString, awsSecretAccessKey NullableString, roleArn NullableString, externalId NullableString, cloudId NullableInt32, ) *AmazonAvailabilityZonesCommand`

NewAmazonAvailabilityZonesCommand instantiates a new AmazonAvailabilityZonesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonAvailabilityZonesCommandWithDefaults

`func NewAmazonAvailabilityZonesCommandWithDefaults() *AmazonAvailabilityZonesCommand`

NewAmazonAvailabilityZonesCommandWithDefaults instantiates a new AmazonAvailabilityZonesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *AmazonAvailabilityZonesCommand) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AmazonAvailabilityZonesCommand) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AmazonAvailabilityZonesCommand) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AmazonAvailabilityZonesCommand) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AmazonAvailabilityZonesCommand) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetAwsAccessKeyId

`func (o *AmazonAvailabilityZonesCommand) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *AmazonAvailabilityZonesCommand) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *AmazonAvailabilityZonesCommand) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.


### SetAwsAccessKeyIdNil

`func (o *AmazonAvailabilityZonesCommand) SetAwsAccessKeyIdNil(b bool)`

 SetAwsAccessKeyIdNil sets the value for AwsAccessKeyId to be an explicit nil

### UnsetAwsAccessKeyId
`func (o *AmazonAvailabilityZonesCommand) UnsetAwsAccessKeyId()`

UnsetAwsAccessKeyId ensures that no value is present for AwsAccessKeyId, not even an explicit nil
### GetAwsSecretAccessKey

`func (o *AmazonAvailabilityZonesCommand) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *AmazonAvailabilityZonesCommand) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *AmazonAvailabilityZonesCommand) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.


### SetAwsSecretAccessKeyNil

`func (o *AmazonAvailabilityZonesCommand) SetAwsSecretAccessKeyNil(b bool)`

 SetAwsSecretAccessKeyNil sets the value for AwsSecretAccessKey to be an explicit nil

### UnsetAwsSecretAccessKey
`func (o *AmazonAvailabilityZonesCommand) UnsetAwsSecretAccessKey()`

UnsetAwsSecretAccessKey ensures that no value is present for AwsSecretAccessKey, not even an explicit nil
### GetRoleArn

`func (o *AmazonAvailabilityZonesCommand) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *AmazonAvailabilityZonesCommand) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *AmazonAvailabilityZonesCommand) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.


### SetRoleArnNil

`func (o *AmazonAvailabilityZonesCommand) SetRoleArnNil(b bool)`

 SetRoleArnNil sets the value for RoleArn to be an explicit nil

### UnsetRoleArn
`func (o *AmazonAvailabilityZonesCommand) UnsetRoleArn()`

UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
### GetExternalId

`func (o *AmazonAvailabilityZonesCommand) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AmazonAvailabilityZonesCommand) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AmazonAvailabilityZonesCommand) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *AmazonAvailabilityZonesCommand) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AmazonAvailabilityZonesCommand) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCloudId

`func (o *AmazonAvailabilityZonesCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AmazonAvailabilityZonesCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AmazonAvailabilityZonesCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *AmazonAvailabilityZonesCommand) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *AmazonAvailabilityZonesCommand) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


