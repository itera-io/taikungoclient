# GetCatalogAppValueCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | **NullableString** |  | 
**Version** | **NullableString** |  | 

## Methods

### NewGetCatalogAppValueCommand

`func NewGetCatalogAppValueCommand(packageId NullableString, version NullableString, ) *GetCatalogAppValueCommand`

NewGetCatalogAppValueCommand instantiates a new GetCatalogAppValueCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogAppValueCommandWithDefaults

`func NewGetCatalogAppValueCommandWithDefaults() *GetCatalogAppValueCommand`

NewGetCatalogAppValueCommandWithDefaults instantiates a new GetCatalogAppValueCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *GetCatalogAppValueCommand) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *GetCatalogAppValueCommand) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *GetCatalogAppValueCommand) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### SetPackageIdNil

`func (o *GetCatalogAppValueCommand) SetPackageIdNil(b bool)`

 SetPackageIdNil sets the value for PackageId to be an explicit nil

### UnsetPackageId
`func (o *GetCatalogAppValueCommand) UnsetPackageId()`

UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
### GetVersion

`func (o *GetCatalogAppValueCommand) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetCatalogAppValueCommand) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetCatalogAppValueCommand) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *GetCatalogAppValueCommand) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *GetCatalogAppValueCommand) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


