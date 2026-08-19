# AzureAksClusterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroupName** | **NullableString** |  | 
**ClusterName** | **NullableString** |  | 
**Location** | **NullableString** |  | 

## Methods

### NewAzureAksClusterDto

`func NewAzureAksClusterDto(resourceGroupName NullableString, clusterName NullableString, location NullableString, ) *AzureAksClusterDto`

NewAzureAksClusterDto instantiates a new AzureAksClusterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAksClusterDtoWithDefaults

`func NewAzureAksClusterDtoWithDefaults() *AzureAksClusterDto`

NewAzureAksClusterDtoWithDefaults instantiates a new AzureAksClusterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroupName

`func (o *AzureAksClusterDto) GetResourceGroupName() string`

GetResourceGroupName returns the ResourceGroupName field if non-nil, zero value otherwise.

### GetResourceGroupNameOk

`func (o *AzureAksClusterDto) GetResourceGroupNameOk() (*string, bool)`

GetResourceGroupNameOk returns a tuple with the ResourceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupName

`func (o *AzureAksClusterDto) SetResourceGroupName(v string)`

SetResourceGroupName sets ResourceGroupName field to given value.


### SetResourceGroupNameNil

`func (o *AzureAksClusterDto) SetResourceGroupNameNil(b bool)`

 SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil

### UnsetResourceGroupName
`func (o *AzureAksClusterDto) UnsetResourceGroupName()`

UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
### GetClusterName

`func (o *AzureAksClusterDto) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *AzureAksClusterDto) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *AzureAksClusterDto) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### SetClusterNameNil

`func (o *AzureAksClusterDto) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *AzureAksClusterDto) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil
### GetLocation

`func (o *AzureAksClusterDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureAksClusterDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureAksClusterDto) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *AzureAksClusterDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AzureAksClusterDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


