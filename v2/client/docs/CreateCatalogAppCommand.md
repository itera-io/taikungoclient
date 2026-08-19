# CreateCatalogAppCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoName** | Pointer to **NullableString** |  | [optional] 
**PackageName** | Pointer to **NullableString** |  | [optional] 
**CatalogId** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to [**[]CatalogAppParamsDto**](CatalogAppParamsDto.md) |  | [optional] 

## Methods

### NewCreateCatalogAppCommand

`func NewCreateCatalogAppCommand() *CreateCatalogAppCommand`

NewCreateCatalogAppCommand instantiates a new CreateCatalogAppCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCatalogAppCommandWithDefaults

`func NewCreateCatalogAppCommandWithDefaults() *CreateCatalogAppCommand`

NewCreateCatalogAppCommandWithDefaults instantiates a new CreateCatalogAppCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoName

`func (o *CreateCatalogAppCommand) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *CreateCatalogAppCommand) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *CreateCatalogAppCommand) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.

### HasRepoName

`func (o *CreateCatalogAppCommand) HasRepoName() bool`

HasRepoName returns a boolean if a field has been set.

### SetRepoNameNil

`func (o *CreateCatalogAppCommand) SetRepoNameNil(b bool)`

 SetRepoNameNil sets the value for RepoName to be an explicit nil

### UnsetRepoName
`func (o *CreateCatalogAppCommand) UnsetRepoName()`

UnsetRepoName ensures that no value is present for RepoName, not even an explicit nil
### GetPackageName

`func (o *CreateCatalogAppCommand) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *CreateCatalogAppCommand) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *CreateCatalogAppCommand) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *CreateCatalogAppCommand) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### SetPackageNameNil

`func (o *CreateCatalogAppCommand) SetPackageNameNil(b bool)`

 SetPackageNameNil sets the value for PackageName to be an explicit nil

### UnsetPackageName
`func (o *CreateCatalogAppCommand) UnsetPackageName()`

UnsetPackageName ensures that no value is present for PackageName, not even an explicit nil
### GetCatalogId

`func (o *CreateCatalogAppCommand) GetCatalogId() int32`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CreateCatalogAppCommand) GetCatalogIdOk() (*int32, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CreateCatalogAppCommand) SetCatalogId(v int32)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *CreateCatalogAppCommand) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetVersion

`func (o *CreateCatalogAppCommand) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateCatalogAppCommand) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateCatalogAppCommand) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateCatalogAppCommand) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *CreateCatalogAppCommand) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *CreateCatalogAppCommand) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetParameters

`func (o *CreateCatalogAppCommand) GetParameters() []CatalogAppParamsDto`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateCatalogAppCommand) GetParametersOk() (*[]CatalogAppParamsDto, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateCatalogAppCommand) SetParameters(v []CatalogAppParamsDto)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CreateCatalogAppCommand) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *CreateCatalogAppCommand) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *CreateCatalogAppCommand) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


