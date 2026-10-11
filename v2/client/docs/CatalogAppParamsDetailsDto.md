# CatalogAppParamsDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CatalogAppName** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsEditableWhenInstalling** | Pointer to **bool** |  | [optional] 
**IsEditableAfterInstallation** | Pointer to **bool** |  | [optional] 
**IsMandatory** | Pointer to **bool** |  | [optional] 
**HasJsonSchema** | Pointer to **bool** |  | [optional] 
**IsTaikunLink** | Pointer to **bool** |  | [optional] 

## Methods

### NewCatalogAppParamsDetailsDto

`func NewCatalogAppParamsDetailsDto() *CatalogAppParamsDetailsDto`

NewCatalogAppParamsDetailsDto instantiates a new CatalogAppParamsDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogAppParamsDetailsDtoWithDefaults

`func NewCatalogAppParamsDetailsDtoWithDefaults() *CatalogAppParamsDetailsDto`

NewCatalogAppParamsDetailsDtoWithDefaults instantiates a new CatalogAppParamsDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogAppParamsDetailsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogAppParamsDetailsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogAppParamsDetailsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogAppParamsDetailsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCatalogAppName

`func (o *CatalogAppParamsDetailsDto) GetCatalogAppName() string`

GetCatalogAppName returns the CatalogAppName field if non-nil, zero value otherwise.

### GetCatalogAppNameOk

`func (o *CatalogAppParamsDetailsDto) GetCatalogAppNameOk() (*string, bool)`

GetCatalogAppNameOk returns a tuple with the CatalogAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogAppName

`func (o *CatalogAppParamsDetailsDto) SetCatalogAppName(v string)`

SetCatalogAppName sets CatalogAppName field to given value.

### HasCatalogAppName

`func (o *CatalogAppParamsDetailsDto) HasCatalogAppName() bool`

HasCatalogAppName returns a boolean if a field has been set.

### SetCatalogAppNameNil

`func (o *CatalogAppParamsDetailsDto) SetCatalogAppNameNil(b bool)`

 SetCatalogAppNameNil sets the value for CatalogAppName to be an explicit nil

### UnsetCatalogAppName
`func (o *CatalogAppParamsDetailsDto) UnsetCatalogAppName()`

UnsetCatalogAppName ensures that no value is present for CatalogAppName, not even an explicit nil
### GetKey

`func (o *CatalogAppParamsDetailsDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CatalogAppParamsDetailsDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CatalogAppParamsDetailsDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CatalogAppParamsDetailsDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *CatalogAppParamsDetailsDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *CatalogAppParamsDetailsDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *CatalogAppParamsDetailsDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CatalogAppParamsDetailsDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CatalogAppParamsDetailsDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CatalogAppParamsDetailsDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CatalogAppParamsDetailsDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CatalogAppParamsDetailsDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsEditableWhenInstalling

`func (o *CatalogAppParamsDetailsDto) GetIsEditableWhenInstalling() bool`

GetIsEditableWhenInstalling returns the IsEditableWhenInstalling field if non-nil, zero value otherwise.

### GetIsEditableWhenInstallingOk

`func (o *CatalogAppParamsDetailsDto) GetIsEditableWhenInstallingOk() (*bool, bool)`

GetIsEditableWhenInstallingOk returns a tuple with the IsEditableWhenInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditableWhenInstalling

`func (o *CatalogAppParamsDetailsDto) SetIsEditableWhenInstalling(v bool)`

SetIsEditableWhenInstalling sets IsEditableWhenInstalling field to given value.

### HasIsEditableWhenInstalling

`func (o *CatalogAppParamsDetailsDto) HasIsEditableWhenInstalling() bool`

HasIsEditableWhenInstalling returns a boolean if a field has been set.

### GetIsEditableAfterInstallation

`func (o *CatalogAppParamsDetailsDto) GetIsEditableAfterInstallation() bool`

GetIsEditableAfterInstallation returns the IsEditableAfterInstallation field if non-nil, zero value otherwise.

### GetIsEditableAfterInstallationOk

`func (o *CatalogAppParamsDetailsDto) GetIsEditableAfterInstallationOk() (*bool, bool)`

GetIsEditableAfterInstallationOk returns a tuple with the IsEditableAfterInstallation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEditableAfterInstallation

`func (o *CatalogAppParamsDetailsDto) SetIsEditableAfterInstallation(v bool)`

SetIsEditableAfterInstallation sets IsEditableAfterInstallation field to given value.

### HasIsEditableAfterInstallation

`func (o *CatalogAppParamsDetailsDto) HasIsEditableAfterInstallation() bool`

HasIsEditableAfterInstallation returns a boolean if a field has been set.

### GetIsMandatory

`func (o *CatalogAppParamsDetailsDto) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *CatalogAppParamsDetailsDto) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *CatalogAppParamsDetailsDto) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *CatalogAppParamsDetailsDto) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.

### GetHasJsonSchema

`func (o *CatalogAppParamsDetailsDto) GetHasJsonSchema() bool`

GetHasJsonSchema returns the HasJsonSchema field if non-nil, zero value otherwise.

### GetHasJsonSchemaOk

`func (o *CatalogAppParamsDetailsDto) GetHasJsonSchemaOk() (*bool, bool)`

GetHasJsonSchemaOk returns a tuple with the HasJsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJsonSchema

`func (o *CatalogAppParamsDetailsDto) SetHasJsonSchema(v bool)`

SetHasJsonSchema sets HasJsonSchema field to given value.

### HasHasJsonSchema

`func (o *CatalogAppParamsDetailsDto) HasHasJsonSchema() bool`

HasHasJsonSchema returns a boolean if a field has been set.

### GetIsTaikunLink

`func (o *CatalogAppParamsDetailsDto) GetIsTaikunLink() bool`

GetIsTaikunLink returns the IsTaikunLink field if non-nil, zero value otherwise.

### GetIsTaikunLinkOk

`func (o *CatalogAppParamsDetailsDto) GetIsTaikunLinkOk() (*bool, bool)`

GetIsTaikunLinkOk returns a tuple with the IsTaikunLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTaikunLink

`func (o *CatalogAppParamsDetailsDto) SetIsTaikunLink(v bool)`

SetIsTaikunLink sets IsTaikunLink field to given value.

### HasIsTaikunLink

`func (o *CatalogAppParamsDetailsDto) HasIsTaikunLink() bool`

HasIsTaikunLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


