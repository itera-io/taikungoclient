# InternalExporterDomainDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**UsersCount** | **int32** |  | 

## Methods

### NewInternalExporterDomainDto

`func NewInternalExporterDomainDto(id int32, name NullableString, usersCount int32, ) *InternalExporterDomainDto`

NewInternalExporterDomainDto instantiates a new InternalExporterDomainDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalExporterDomainDtoWithDefaults

`func NewInternalExporterDomainDtoWithDefaults() *InternalExporterDomainDto`

NewInternalExporterDomainDtoWithDefaults instantiates a new InternalExporterDomainDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalExporterDomainDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalExporterDomainDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalExporterDomainDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *InternalExporterDomainDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalExporterDomainDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalExporterDomainDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *InternalExporterDomainDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InternalExporterDomainDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUsersCount

`func (o *InternalExporterDomainDto) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *InternalExporterDomainDto) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *InternalExporterDomainDto) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


