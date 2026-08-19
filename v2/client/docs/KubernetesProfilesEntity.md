# KubernetesProfilesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**TaikunLBEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubernetesProfilesEntity

`func NewKubernetesProfilesEntity() *KubernetesProfilesEntity`

NewKubernetesProfilesEntity instantiates a new KubernetesProfilesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesProfilesEntityWithDefaults

`func NewKubernetesProfilesEntityWithDefaults() *KubernetesProfilesEntity`

NewKubernetesProfilesEntityWithDefaults instantiates a new KubernetesProfilesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesProfilesEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesProfilesEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesProfilesEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KubernetesProfilesEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KubernetesProfilesEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesProfilesEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesProfilesEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesProfilesEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KubernetesProfilesEntity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubernetesProfilesEntity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaikunLBEnabled

`func (o *KubernetesProfilesEntity) GetTaikunLBEnabled() bool`

GetTaikunLBEnabled returns the TaikunLBEnabled field if non-nil, zero value otherwise.

### GetTaikunLBEnabledOk

`func (o *KubernetesProfilesEntity) GetTaikunLBEnabledOk() (*bool, bool)`

GetTaikunLBEnabledOk returns a tuple with the TaikunLBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaikunLBEnabled

`func (o *KubernetesProfilesEntity) SetTaikunLBEnabled(v bool)`

SetTaikunLBEnabled sets TaikunLBEnabled field to given value.

### HasTaikunLBEnabled

`func (o *KubernetesProfilesEntity) HasTaikunLBEnabled() bool`

HasTaikunLBEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


