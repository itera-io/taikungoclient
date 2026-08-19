# UnbindAppRepositoryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** |  | 
**OrganizationId** | **int32** |  | 

## Methods

### NewUnbindAppRepositoryCommand

`func NewUnbindAppRepositoryCommand(ids []string, organizationId int32, ) *UnbindAppRepositoryCommand`

NewUnbindAppRepositoryCommand instantiates a new UnbindAppRepositoryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnbindAppRepositoryCommandWithDefaults

`func NewUnbindAppRepositoryCommandWithDefaults() *UnbindAppRepositoryCommand`

NewUnbindAppRepositoryCommandWithDefaults instantiates a new UnbindAppRepositoryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *UnbindAppRepositoryCommand) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UnbindAppRepositoryCommand) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UnbindAppRepositoryCommand) SetIds(v []string)`

SetIds sets Ids field to given value.


### SetIdsNil

`func (o *UnbindAppRepositoryCommand) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UnbindAppRepositoryCommand) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetOrganizationId

`func (o *UnbindAppRepositoryCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UnbindAppRepositoryCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UnbindAppRepositoryCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


