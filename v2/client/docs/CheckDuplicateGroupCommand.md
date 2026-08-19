# CheckDuplicateGroupCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **NullableInt32** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewCheckDuplicateGroupCommand

`func NewCheckDuplicateGroupCommand(domainId NullableInt32, name NullableString, ) *CheckDuplicateGroupCommand`

NewCheckDuplicateGroupCommand instantiates a new CheckDuplicateGroupCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDuplicateGroupCommandWithDefaults

`func NewCheckDuplicateGroupCommandWithDefaults() *CheckDuplicateGroupCommand`

NewCheckDuplicateGroupCommandWithDefaults instantiates a new CheckDuplicateGroupCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *CheckDuplicateGroupCommand) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CheckDuplicateGroupCommand) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CheckDuplicateGroupCommand) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### SetDomainIdNil

`func (o *CheckDuplicateGroupCommand) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *CheckDuplicateGroupCommand) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetName

`func (o *CheckDuplicateGroupCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckDuplicateGroupCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckDuplicateGroupCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CheckDuplicateGroupCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CheckDuplicateGroupCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


