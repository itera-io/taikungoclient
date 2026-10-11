# EditCatalogCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**Description** | **NullableString** |  | 

## Methods

### NewEditCatalogCommand

`func NewEditCatalogCommand(id int32, name NullableString, description NullableString, ) *EditCatalogCommand`

NewEditCatalogCommand instantiates a new EditCatalogCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditCatalogCommandWithDefaults

`func NewEditCatalogCommandWithDefaults() *EditCatalogCommand`

NewEditCatalogCommandWithDefaults instantiates a new EditCatalogCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditCatalogCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditCatalogCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditCatalogCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *EditCatalogCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditCatalogCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditCatalogCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *EditCatalogCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditCatalogCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EditCatalogCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditCatalogCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditCatalogCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EditCatalogCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditCatalogCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


