# CreateAzureCloudCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**AzureSubscriptionId** | Pointer to **NullableString** |  | [optional] 
**AzureClientId** | Pointer to **NullableString** |  | [optional] 
**AzureClientSecret** | Pointer to **NullableString** |  | [optional] 
**AzureTenantId** | Pointer to **NullableString** |  | [optional] 
**AzureLocation** | Pointer to **NullableString** |  | [optional] 
**AzCount** | Pointer to **NullableInt32** |  | [optional] 
**OrganizationId** | Pointer to **NullableInt32** |  | [optional] 
**IpMode** | Pointer to [**IpMode**](IpMode.md) |  | [optional] 
**VnetMode** | Pointer to [**VnetMode**](VnetMode.md) |  | [optional] 
**ExistingVirtualNetworkResourceGroupName** | Pointer to **NullableString** |  | [optional] 
**ExistingVirtualNetworkName** | Pointer to **NullableString** |  | [optional] 
**SharedFileSystemEnabled** | Pointer to **bool** |  | [optional] 
**ImportedFileShareId** | Pointer to **NullableString** |  | [optional] 
**ImportedFileShareResourceGroupName** | Pointer to **NullableString** |  | [optional] 
**Subnets** | Pointer to [**[]CreateAzureSubnetDto**](CreateAzureSubnetDto.md) |  | [optional] 

## Methods

### NewCreateAzureCloudCommand

`func NewCreateAzureCloudCommand() *CreateAzureCloudCommand`

NewCreateAzureCloudCommand instantiates a new CreateAzureCloudCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzureCloudCommandWithDefaults

`func NewCreateAzureCloudCommandWithDefaults() *CreateAzureCloudCommand`

NewCreateAzureCloudCommandWithDefaults instantiates a new CreateAzureCloudCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAzureCloudCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAzureCloudCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAzureCloudCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAzureCloudCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateAzureCloudCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateAzureCloudCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAzureSubscriptionId

`func (o *CreateAzureCloudCommand) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *CreateAzureCloudCommand) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *CreateAzureCloudCommand) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *CreateAzureCloudCommand) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### SetAzureSubscriptionIdNil

`func (o *CreateAzureCloudCommand) SetAzureSubscriptionIdNil(b bool)`

 SetAzureSubscriptionIdNil sets the value for AzureSubscriptionId to be an explicit nil

### UnsetAzureSubscriptionId
`func (o *CreateAzureCloudCommand) UnsetAzureSubscriptionId()`

UnsetAzureSubscriptionId ensures that no value is present for AzureSubscriptionId, not even an explicit nil
### GetAzureClientId

`func (o *CreateAzureCloudCommand) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *CreateAzureCloudCommand) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *CreateAzureCloudCommand) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *CreateAzureCloudCommand) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### SetAzureClientIdNil

`func (o *CreateAzureCloudCommand) SetAzureClientIdNil(b bool)`

 SetAzureClientIdNil sets the value for AzureClientId to be an explicit nil

### UnsetAzureClientId
`func (o *CreateAzureCloudCommand) UnsetAzureClientId()`

UnsetAzureClientId ensures that no value is present for AzureClientId, not even an explicit nil
### GetAzureClientSecret

`func (o *CreateAzureCloudCommand) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### GetAzureClientSecretOk

`func (o *CreateAzureCloudCommand) GetAzureClientSecretOk() (*string, bool)`

GetAzureClientSecretOk returns a tuple with the AzureClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientSecret

`func (o *CreateAzureCloudCommand) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### HasAzureClientSecret

`func (o *CreateAzureCloudCommand) HasAzureClientSecret() bool`

HasAzureClientSecret returns a boolean if a field has been set.

### SetAzureClientSecretNil

`func (o *CreateAzureCloudCommand) SetAzureClientSecretNil(b bool)`

 SetAzureClientSecretNil sets the value for AzureClientSecret to be an explicit nil

### UnsetAzureClientSecret
`func (o *CreateAzureCloudCommand) UnsetAzureClientSecret()`

UnsetAzureClientSecret ensures that no value is present for AzureClientSecret, not even an explicit nil
### GetAzureTenantId

`func (o *CreateAzureCloudCommand) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *CreateAzureCloudCommand) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *CreateAzureCloudCommand) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *CreateAzureCloudCommand) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### SetAzureTenantIdNil

`func (o *CreateAzureCloudCommand) SetAzureTenantIdNil(b bool)`

 SetAzureTenantIdNil sets the value for AzureTenantId to be an explicit nil

### UnsetAzureTenantId
`func (o *CreateAzureCloudCommand) UnsetAzureTenantId()`

UnsetAzureTenantId ensures that no value is present for AzureTenantId, not even an explicit nil
### GetAzureLocation

`func (o *CreateAzureCloudCommand) GetAzureLocation() string`

GetAzureLocation returns the AzureLocation field if non-nil, zero value otherwise.

### GetAzureLocationOk

`func (o *CreateAzureCloudCommand) GetAzureLocationOk() (*string, bool)`

GetAzureLocationOk returns a tuple with the AzureLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureLocation

`func (o *CreateAzureCloudCommand) SetAzureLocation(v string)`

SetAzureLocation sets AzureLocation field to given value.

### HasAzureLocation

`func (o *CreateAzureCloudCommand) HasAzureLocation() bool`

HasAzureLocation returns a boolean if a field has been set.

### SetAzureLocationNil

`func (o *CreateAzureCloudCommand) SetAzureLocationNil(b bool)`

 SetAzureLocationNil sets the value for AzureLocation to be an explicit nil

### UnsetAzureLocation
`func (o *CreateAzureCloudCommand) UnsetAzureLocation()`

UnsetAzureLocation ensures that no value is present for AzureLocation, not even an explicit nil
### GetAzCount

`func (o *CreateAzureCloudCommand) GetAzCount() int32`

GetAzCount returns the AzCount field if non-nil, zero value otherwise.

### GetAzCountOk

`func (o *CreateAzureCloudCommand) GetAzCountOk() (*int32, bool)`

GetAzCountOk returns a tuple with the AzCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzCount

`func (o *CreateAzureCloudCommand) SetAzCount(v int32)`

SetAzCount sets AzCount field to given value.

### HasAzCount

`func (o *CreateAzureCloudCommand) HasAzCount() bool`

HasAzCount returns a boolean if a field has been set.

### SetAzCountNil

`func (o *CreateAzureCloudCommand) SetAzCountNil(b bool)`

 SetAzCountNil sets the value for AzCount to be an explicit nil

### UnsetAzCount
`func (o *CreateAzureCloudCommand) UnsetAzCount()`

UnsetAzCount ensures that no value is present for AzCount, not even an explicit nil
### GetOrganizationId

`func (o *CreateAzureCloudCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateAzureCloudCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateAzureCloudCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateAzureCloudCommand) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *CreateAzureCloudCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *CreateAzureCloudCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetIpMode

`func (o *CreateAzureCloudCommand) GetIpMode() IpMode`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *CreateAzureCloudCommand) GetIpModeOk() (*IpMode, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *CreateAzureCloudCommand) SetIpMode(v IpMode)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *CreateAzureCloudCommand) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetVnetMode

`func (o *CreateAzureCloudCommand) GetVnetMode() VnetMode`

GetVnetMode returns the VnetMode field if non-nil, zero value otherwise.

### GetVnetModeOk

`func (o *CreateAzureCloudCommand) GetVnetModeOk() (*VnetMode, bool)`

GetVnetModeOk returns a tuple with the VnetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetMode

`func (o *CreateAzureCloudCommand) SetVnetMode(v VnetMode)`

SetVnetMode sets VnetMode field to given value.

### HasVnetMode

`func (o *CreateAzureCloudCommand) HasVnetMode() bool`

HasVnetMode returns a boolean if a field has been set.

### GetExistingVirtualNetworkResourceGroupName

`func (o *CreateAzureCloudCommand) GetExistingVirtualNetworkResourceGroupName() string`

GetExistingVirtualNetworkResourceGroupName returns the ExistingVirtualNetworkResourceGroupName field if non-nil, zero value otherwise.

### GetExistingVirtualNetworkResourceGroupNameOk

`func (o *CreateAzureCloudCommand) GetExistingVirtualNetworkResourceGroupNameOk() (*string, bool)`

GetExistingVirtualNetworkResourceGroupNameOk returns a tuple with the ExistingVirtualNetworkResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingVirtualNetworkResourceGroupName

`func (o *CreateAzureCloudCommand) SetExistingVirtualNetworkResourceGroupName(v string)`

SetExistingVirtualNetworkResourceGroupName sets ExistingVirtualNetworkResourceGroupName field to given value.

### HasExistingVirtualNetworkResourceGroupName

`func (o *CreateAzureCloudCommand) HasExistingVirtualNetworkResourceGroupName() bool`

HasExistingVirtualNetworkResourceGroupName returns a boolean if a field has been set.

### SetExistingVirtualNetworkResourceGroupNameNil

`func (o *CreateAzureCloudCommand) SetExistingVirtualNetworkResourceGroupNameNil(b bool)`

 SetExistingVirtualNetworkResourceGroupNameNil sets the value for ExistingVirtualNetworkResourceGroupName to be an explicit nil

### UnsetExistingVirtualNetworkResourceGroupName
`func (o *CreateAzureCloudCommand) UnsetExistingVirtualNetworkResourceGroupName()`

UnsetExistingVirtualNetworkResourceGroupName ensures that no value is present for ExistingVirtualNetworkResourceGroupName, not even an explicit nil
### GetExistingVirtualNetworkName

`func (o *CreateAzureCloudCommand) GetExistingVirtualNetworkName() string`

GetExistingVirtualNetworkName returns the ExistingVirtualNetworkName field if non-nil, zero value otherwise.

### GetExistingVirtualNetworkNameOk

`func (o *CreateAzureCloudCommand) GetExistingVirtualNetworkNameOk() (*string, bool)`

GetExistingVirtualNetworkNameOk returns a tuple with the ExistingVirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingVirtualNetworkName

`func (o *CreateAzureCloudCommand) SetExistingVirtualNetworkName(v string)`

SetExistingVirtualNetworkName sets ExistingVirtualNetworkName field to given value.

### HasExistingVirtualNetworkName

`func (o *CreateAzureCloudCommand) HasExistingVirtualNetworkName() bool`

HasExistingVirtualNetworkName returns a boolean if a field has been set.

### SetExistingVirtualNetworkNameNil

`func (o *CreateAzureCloudCommand) SetExistingVirtualNetworkNameNil(b bool)`

 SetExistingVirtualNetworkNameNil sets the value for ExistingVirtualNetworkName to be an explicit nil

### UnsetExistingVirtualNetworkName
`func (o *CreateAzureCloudCommand) UnsetExistingVirtualNetworkName()`

UnsetExistingVirtualNetworkName ensures that no value is present for ExistingVirtualNetworkName, not even an explicit nil
### GetSharedFileSystemEnabled

`func (o *CreateAzureCloudCommand) GetSharedFileSystemEnabled() bool`

GetSharedFileSystemEnabled returns the SharedFileSystemEnabled field if non-nil, zero value otherwise.

### GetSharedFileSystemEnabledOk

`func (o *CreateAzureCloudCommand) GetSharedFileSystemEnabledOk() (*bool, bool)`

GetSharedFileSystemEnabledOk returns a tuple with the SharedFileSystemEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedFileSystemEnabled

`func (o *CreateAzureCloudCommand) SetSharedFileSystemEnabled(v bool)`

SetSharedFileSystemEnabled sets SharedFileSystemEnabled field to given value.

### HasSharedFileSystemEnabled

`func (o *CreateAzureCloudCommand) HasSharedFileSystemEnabled() bool`

HasSharedFileSystemEnabled returns a boolean if a field has been set.

### GetImportedFileShareId

`func (o *CreateAzureCloudCommand) GetImportedFileShareId() string`

GetImportedFileShareId returns the ImportedFileShareId field if non-nil, zero value otherwise.

### GetImportedFileShareIdOk

`func (o *CreateAzureCloudCommand) GetImportedFileShareIdOk() (*string, bool)`

GetImportedFileShareIdOk returns a tuple with the ImportedFileShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFileShareId

`func (o *CreateAzureCloudCommand) SetImportedFileShareId(v string)`

SetImportedFileShareId sets ImportedFileShareId field to given value.

### HasImportedFileShareId

`func (o *CreateAzureCloudCommand) HasImportedFileShareId() bool`

HasImportedFileShareId returns a boolean if a field has been set.

### SetImportedFileShareIdNil

`func (o *CreateAzureCloudCommand) SetImportedFileShareIdNil(b bool)`

 SetImportedFileShareIdNil sets the value for ImportedFileShareId to be an explicit nil

### UnsetImportedFileShareId
`func (o *CreateAzureCloudCommand) UnsetImportedFileShareId()`

UnsetImportedFileShareId ensures that no value is present for ImportedFileShareId, not even an explicit nil
### GetImportedFileShareResourceGroupName

`func (o *CreateAzureCloudCommand) GetImportedFileShareResourceGroupName() string`

GetImportedFileShareResourceGroupName returns the ImportedFileShareResourceGroupName field if non-nil, zero value otherwise.

### GetImportedFileShareResourceGroupNameOk

`func (o *CreateAzureCloudCommand) GetImportedFileShareResourceGroupNameOk() (*string, bool)`

GetImportedFileShareResourceGroupNameOk returns a tuple with the ImportedFileShareResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedFileShareResourceGroupName

`func (o *CreateAzureCloudCommand) SetImportedFileShareResourceGroupName(v string)`

SetImportedFileShareResourceGroupName sets ImportedFileShareResourceGroupName field to given value.

### HasImportedFileShareResourceGroupName

`func (o *CreateAzureCloudCommand) HasImportedFileShareResourceGroupName() bool`

HasImportedFileShareResourceGroupName returns a boolean if a field has been set.

### SetImportedFileShareResourceGroupNameNil

`func (o *CreateAzureCloudCommand) SetImportedFileShareResourceGroupNameNil(b bool)`

 SetImportedFileShareResourceGroupNameNil sets the value for ImportedFileShareResourceGroupName to be an explicit nil

### UnsetImportedFileShareResourceGroupName
`func (o *CreateAzureCloudCommand) UnsetImportedFileShareResourceGroupName()`

UnsetImportedFileShareResourceGroupName ensures that no value is present for ImportedFileShareResourceGroupName, not even an explicit nil
### GetSubnets

`func (o *CreateAzureCloudCommand) GetSubnets() []CreateAzureSubnetDto`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateAzureCloudCommand) GetSubnetsOk() (*[]CreateAzureSubnetDto, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateAzureCloudCommand) SetSubnets(v []CreateAzureSubnetDto)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *CreateAzureCloudCommand) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *CreateAzureCloudCommand) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *CreateAzureCloudCommand) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


