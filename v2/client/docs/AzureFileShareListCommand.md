# AzureFileShareListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureClientId** | **NullableString** |  | 
**AzureClientSecret** | **NullableString** |  | 
**AzureTenantId** | **NullableString** |  | 
**AzureSubscriptionId** | **NullableString** |  | 
**ResourceGroupName** | **NullableString** |  | 
**CloudId** | **NullableInt32** |  | 

## Methods

### NewAzureFileShareListCommand

`func NewAzureFileShareListCommand(azureClientId NullableString, azureClientSecret NullableString, azureTenantId NullableString, azureSubscriptionId NullableString, resourceGroupName NullableString, cloudId NullableInt32, ) *AzureFileShareListCommand`

NewAzureFileShareListCommand instantiates a new AzureFileShareListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureFileShareListCommandWithDefaults

`func NewAzureFileShareListCommandWithDefaults() *AzureFileShareListCommand`

NewAzureFileShareListCommandWithDefaults instantiates a new AzureFileShareListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureClientId

`func (o *AzureFileShareListCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *AzureFileShareListCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *AzureFileShareListCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### SetAzureClientIdNil

`func (o *AzureFileShareListCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *AzureFileShareListCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *AzureFileShareListCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *AzureFileShareListCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *AzureFileShareListCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### SetAzureClientSecretNil

`func (o *AzureFileShareListCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *AzureFileShareListCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetAzureTenantId

`func (o *AzureFileShareListCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *AzureFileShareListCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *AzureFileShareListCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### SetAzureTenantIdNil

`func (o *AzureFileShareListCommand) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *AzureFileShareListCommand) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureSubscriptionId

`func (o *AzureFileShareListCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AzureFileShareListCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AzureFileShareListCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### SetAzureSubscriptionIdNil

`func (o *AzureFileShareListCommand) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *AzureFileShareListCommand) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
### GetResourceGroupName

`func (o *AzureFileShareListCommand) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureFileShareListCommand) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureFileShareListCommand) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### SetResourceGroupNameNil

`func (o *AzureFileShareListCommand) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureFileShareListCommand) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetCloudId

`func (o *AzureFileShareListCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AzureFileShareListCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AzureFileShareListCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *AzureFileShareListCommand) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *AzureFileShareListCommand) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


