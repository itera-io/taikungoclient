# UpdateGenericKubernetesCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewUpdateGenericKubernetesCommand

`func NewUpdateGenericKubernetesCommand(id int32, name NullableString, ) *UpdateGenericKubernetesCommand`

NewUpdateGenericKubernetesCommand instantiates a new UpdateGenericKubernetesCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGenericKubernetesCommandWithDefaults

`func NewUpdateGenericKubernetesCommandWithDefaults() *UpdateGenericKubernetesCommand`

NewUpdateGenericKubernetesCommandWithDefaults instantiates a new UpdateGenericKubernetesCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateGenericKubernetesCommand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateGenericKubernetesCommand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateGenericKubernetesCommand) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateGenericKubernetesCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGenericKubernetesCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGenericKubernetesCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *UpdateGenericKubernetesCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateGenericKubernetesCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


