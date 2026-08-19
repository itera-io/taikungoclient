# AzureZonesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureSubscriptionId** | **NullableString** |  | 
**AzureClientId** | **NullableString** |  | 
**AzureClientSecret** | **NullableString** |  | 
**AzureTenantId** | **NullableString** |  | 
**AzureLocation** | **NullableString** |  | 
**CloudId** | **NullableInt32** |  | 

## Methods

### NewAzureZonesCommand

`func NewAzureZonesCommand(azureSubscriptionId NullableString, azureClientId NullableString, azureClientSecret NullableString, azureTenantId NullableString, azureLocation NullableString, cloudId NullableInt32, ) *AzureZonesCommand`

NewAzureZonesCommand instantiates a new AzureZonesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureZonesCommandWithDefaults

`func NewAzureZonesCommandWithDefaults() *AzureZonesCommand`

NewAzureZonesCommandWithDefaults instantiates a new AzureZonesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureSubscriptionId

`func (o *AzureZonesCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AzureZonesCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AzureZonesCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### SetAzureSubscriptionIdNil

`func (o *AzureZonesCommand) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *AzureZonesCommand) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
### GetAzureClientId

`func (o *AzureZonesCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *AzureZonesCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *AzureZonesCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.


### SetAzureClientIdNil

`func (o *AzureZonesCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *AzureZonesCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *AzureZonesCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *AzureZonesCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *AzureZonesCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.


### SetAzureClientSecretNil

`func (o *AzureZonesCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *AzureZonesCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetAzureTenantId

`func (o *AzureZonesCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *AzureZonesCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *AzureZonesCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.


### SetAzureTenantIdNil

`func (o *AzureZonesCommand) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *AzureZonesCommand) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureLocation

`func (o *AzureZonesCommand) GetAzureLocation() string`

GetAzureLocation returns the AzureLocation field if non-nil, zero value otherwise.

### GetAzureLocationOk

`func (o *AzureZonesCommand) GetAzureLocationOk() (*string, bool)`

GetAzureLocationOk returns a tuple with the AzureLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureLocation

`func (o *AzureZonesCommand) SetAzureLocation(v string)`

SetAzureLocation sets AzureLocation field to given value.


### SetAzureLocationNil

`func (o *AzureZonesCommand) SetAzureLocationNil(b bool)`

 SetAzureLocationNil sets the value for AzureLocation to be an explicit nil

### UnsetAzureLocation
`func (o *AzureZonesCommand) UnsetAzureLocation()`

UnsetAzureLocation ensures that no value is present for AzureLocation, not even an explicit nil
### GetCloudId

`func (o *AzureZonesCommand) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *AzureZonesCommand) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *AzureZonesCommand) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### SetCloudIdNil

`func (o *AzureZonesCommand) SetCloudIdNil(b bool)`

 SetCloudIdNil sets the value for CloudId to be an explicit nil

### UnsetCloudId
`func (o *AzureZonesCommand) UnsetCloudId()`

UnsetCloudId ensures that no value is present for CloudId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


