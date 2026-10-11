# CatalogDropdownDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**PackageIds** | **[]string** |  | 
**IsDefault** | **bool** |  | 

## Methods

### NewCatalogDropdownDto

`func NewCatalogDropdownDto(id int32, name NullableString, packageIds []string, isDefault bool, ) *CatalogDropdownDto`

NewCatalogDropdownDto instantiates a new CatalogDropdownDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogDropdownDtoWithDefaults

`func NewCatalogDropdownDtoWithDefaults() *CatalogDropdownDto`

NewCatalogDropdownDtoWithDefaults instantiates a new CatalogDropdownDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogDropdownDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogDropdownDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogDropdownDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CatalogDropdownDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogDropdownDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogDropdownDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CatalogDropdownDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CatalogDropdownDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPackageIds

`func (o *CatalogDropdownDto) GetPackageIds() []string`

GetPackageIds returns the PackageIds field if non-nil, zero value otherwise.

### GetPackageIdsOk

`func (o *CatalogDropdownDto) GetPackageIdsOk() (*[]string, bool)`

GetPackageIdsOk returns a tuple with the PackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageIds

`func (o *CatalogDropdownDto) SetPackageIds(v []string)`

SetPackageIds sets PackageIds field to given value.


### SetPackageIdsNil

`func (o *CatalogDropdownDto) SetPackageIdsNil(b bool)`

 SetPackageIdsNil sets the value for PackageIds to be an explicit nil

### UnsetPackageIds
`func (o *CatalogDropdownDto) UnsetPackageIds()`

UnsetPackageIds ensures that no value is present for PackageIds, not even an explicit nil
### GetIsDefault

`func (o *CatalogDropdownDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CatalogDropdownDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CatalogDropdownDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


