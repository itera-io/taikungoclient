# AzureVnetSubnetListCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureClientId** | **NullableString** |  | 
**AzureClientSecret** | **NullableString** |  | 
**AzureTenantId** | **NullableString** |  | 
**AzureSubscriptionId** | **NullableString** |  | 
**ResourceGroupName** | **NullableString** |  | 
**VnetName** | **NullableString** |  | 
**CloudId** | **NullableInt32** |  | 
**IpMode** | [**IpMode**](IpMode.md) |  | 

## Methods

### NewAzureVnetSubnetListCommand

`func NewAzureVnetSubnetListCommand(azureClientId NullableString, azureClientSecret NullableString, azureTenantId NullableString, azureSubscriptionId NullableString, resourceGroupName NullableString, vnetName NullableString, cloudId NullableInt32, ipMode IpMode, ) *AzureVnetSubnetListCommand`

NewAzureVnetSubnetListCommand instantiates a new AzureVnetSubnetListCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureVnetSubnetListCommandWithDefaults

`func NewAzureVnetSubnetListCommandWithDefaults() *AzureVnetSubnetListCommand`

NewAzureVnetSubnetListCommandWithDefaults instantiates a new AzureVnetSubnetListCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureClientId

`func (o *AzureVnetSubnetListCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *AzureVnetSubnetListCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *AzureVnetSubnetListCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### SetAzureClientIdNil

`func (o *AzureVnetSubnetListCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *AzureVnetSubnetListCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *AzureVnetSubnetListCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *AzureVnetSubnetListCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *AzureVnetSubnetListCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### SetAzureClientSecretNil

`func (o *AzureVnetSubnetListCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *AzureVnetSubnetListCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetAzureTenantId

`func (o *AzureVnetSubnetListCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *AzureVnetSubnetListCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *AzureVnetSubnetListCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### SetAzureTenantIdNil

`func (o *AzureVnetSubnetListCommand) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *AzureVnetSubnetListCommand) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureSubscriptionId

`func (o *AzureVnetSubnetListCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AzureVnetSubnetListCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AzureVnetSubnetListCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### SetAzureSubscriptionIdNil

`func (o *AzureVnetSubnetListCommand) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *AzureVnetSubnetListCommand) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
### GetResourceGroupName

`func (o *AzureVnetSubnetListCommand) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureVnetSubnetListCommand) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureVnetSubnetListCommand) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### SetResourceGroupNameNil

`func (o *AzureVnetSubnetListCommand) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureVnetSubnetListCommand) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetVnetName

`func (o *AzureVnetSubnetListCommand) GetVnetName() string`

GetVnetName returns the VnetName field if non-nil, zero value otherwise.

### GetVnetNameOk

`func (o *AzureVnetSubnetListCommand) GetVnetNameOk() (*string, bool)`

GetVnetNameOk returns a tuple with the VnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetName

`func (o *AzureVnetSubnetListCommand) SetVnetName(v string)`

SetVnetName sets VnetName field to given value.


### SetVnetNameNil

`func (o *AzureVnetSubnetListCommand) SetVnetNameNil(b bool)`

 SetVnetNameNil sets the value for VnetName to be an explicit nil

### UnsetVnetName
`func (o *AzureVnetSubnetListCommand) UnsetVnetName()`

UnsetVnetName ensures that no value is present for VnetName, not even an explicit nil
### GetCloudId

`func (o *AzureVnetSubnetListCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AzureVnetSubnetListCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AzureVnetSubnetListCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *AzureVnetSubnetListCommand) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *AzureVnetSubnetListCommand) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil
### GetIpMode

`func (o *AzureVnetSubnetListCommand) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *AzureVnetSubnetListCommand) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *AzureVnetSubnetListCommand) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


