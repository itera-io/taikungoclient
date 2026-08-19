# IngressSearchCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **NullableInt32** |  | 
**Offset** | **NullableInt32** |  | 
**SearchTerm** | **NullableString** |  | 
**IncludePublicImportedClusters** | **NullableBool** |  | 

## Methods

### NewIngressSearchCommand

`func NewIngressSearchCommand(limit NullableInt32, offset NullableInt32, searchTerm NullableString, includePublicImportedClusters NullableBool, ) *IngressSearchCommand`

NewIngressSearchCommand instantiates a new IngressSearchCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressSearchCommandWithDefaults

`func NewIngressSearchCommandWithDefaults() *IngressSearchCommand`

NewIngressSearchCommandWithDefaults instantiates a new IngressSearchCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *IngressSearchCommand) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IngressSearchCommand) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IngressSearchCommand) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### SetLimitNil

`func (o *IngressSearchCommand) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *IngressSearchCommand) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetOffset

`func (o *IngressSearchCommand) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IngressSearchCommand) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IngressSearchCommand) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### SetOffsetNil

`func (o *IngressSearchCommand) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *IngressSearchCommand) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSearchTerm

`func (o *IngressSearchCommand) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *IngressSearchCommand) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *IngressSearchCommand) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.


### SetSearchTermNil

`func (o *IngressSearchCommand) SetSearchTermNil(b bool)`

 SetSearchTermNil sets the value for SearchTerm to be an explicit nil

### UnsetSearchTerm
`func (o *IngressSearchCommand) UnsetSearchTerm()`

UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
### GetIncludePublicImportedClusters

`func (o *IngressSearchCommand) GetIncludePublicImportedClusters() bool`

GetIncludePublicImportedClusters returns the IncludePublicImportedClusters field if non-nil, zero value otherwise.

### GetIncludePublicImportedClustersOk

`func (o *IngressSearchCommand) GetIncludePublicImportedClustersOk() (*bool, bool)`

GetIncludePublicImportedClustersOk returns a tuple with the IncludePublicImportedClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePublicImportedClusters

`func (o *IngressSearchCommand) SetIncludePublicImportedClusters(v bool)`

SetIncludePublicImportedClusters sets IncludePublicImportedClusters field to given value.


### SetIncludePublicImportedClustersNil

`func (o *IngressSearchCommand) SetIncludePublicImportedClustersNil(b bool)`

 SetIncludePublicImportedClustersNil sets the value for IncludePublicImportedClusters to be an explicit nil

### UnsetIncludePublicImportedClusters
`func (o *IngressSearchCommand) UnsetIncludePublicImportedClusters()`

UnsetIncludePublicImportedClusters ensures that no value is present for IncludePublicImportedClusters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


