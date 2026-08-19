# TrustedRegistryListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Registry** | **NullableString** |  | 

## Methods

### NewTrustedRegistryListDto

`func NewTrustedRegistryListDto(id int32, registry NullableString, ) *TrustedRegistryListDto`

NewTrustedRegistryListDto instantiates a new TrustedRegistryListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustedRegistryListDtoWithDefaults

`func NewTrustedRegistryListDtoWithDefaults() *TrustedRegistryListDto`

NewTrustedRegistryListDtoWithDefaults instantiates a new TrustedRegistryListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustedRegistryListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustedRegistryListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustedRegistryListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetRegistry

`func (o *TrustedRegistryListDto) GetRegistry() string`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *TrustedRegistryListDto) GetRegistryOk() (*string, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *TrustedRegistryListDto) SetRegistry(v string)`

SetRegistry sets Registry field to given value.


### SetRegistryNil

`func (o *TrustedRegistryListDto) SetRegistryNil(b bool)`

 SetRegistryNil sets the value for Registry to be an explicit nil

### UnsetRegistry
`func (o *TrustedRegistryListDto) UnsetRegistry()`

UnsetRegistry ensures that no value is present for Registry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


