# BindAppRepositoryCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilteringElements** | [**[]FilteringElementDto**](FilteringElementDto.md) |  | 
**OrganizationId** | **NullableInt32** |  | 

## Methods

### NewBindAppRepositoryCommand

`func NewBindAppRepositoryCommand(filteringElements []FilteringElementDto, organizationId NullableInt32, ) *BindAppRepositoryCommand`

NewBindAppRepositoryCommand instantiates a new BindAppRepositoryCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindAppRepositoryCommandWithDefaults

`func NewBindAppRepositoryCommandWithDefaults() *BindAppRepositoryCommand`

NewBindAppRepositoryCommandWithDefaults instantiates a new BindAppRepositoryCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilteringElements

`func (o *BindAppRepositoryCommand) GetFilteringElements() []FilteringElementDto`

GetFilteringElements returns the FilteringElements field if non-nil, zero value otherwise.

### GetFilteringElementsOk

`func (o *BindAppRepositoryCommand) GetFilteringElementsOk() (*[]FilteringElementDto, bool)`

GetFilteringElementsOk returns a tuple with the FilteringElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteringElements

`func (o *BindAppRepositoryCommand) SetFilteringElements(v []FilteringElementDto)`

SetFilteringElements sets FilteringElements field to given value.


### SetFilteringElementsNil

`func (o *BindAppRepositoryCommand) SetFilteringElementsNil(b bool)`

 SetFilteringElementsNil sets the value for FilteringElements to be an explicit nil

### UnsetFilteringElements
`func (o *BindAppRepositoryCommand) UnsetFilteringElements()`

UnsetFilteringElements ensures that no value is present for FilteringElements, not even an explicit nil
### GetOrganizationId

`func (o *BindAppRepositoryCommand) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BindAppRepositoryCommand) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BindAppRepositoryCommand) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### SetOrganizationIdNil

`func (o *BindAppRepositoryCommand) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *BindAppRepositoryCommand) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


