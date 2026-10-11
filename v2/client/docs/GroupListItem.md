# GroupListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**ClaimValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGroupListItem

`func NewGroupListItem(id int32, name string, ) *GroupListItem`

NewGroupListItem instantiates a new GroupListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupListItemWithDefaults

`func NewGroupListItemWithDefaults() *GroupListItem`

NewGroupListItemWithDefaults instantiates a new GroupListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GroupListItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupListItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupListItem) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GroupListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupListItem) SetName(v string)`

SetName sets Name field to given value.


### GetClaimValue

`func (o *GroupListItem) GetClaimValue() string`

GetClaimValue returns the ClaimValue field if non-nil, zero value otherwise.

### GetClaimValueOk

`func (o *GroupListItem) GetClaimValueOk() (*string, bool)`

GetClaimValueOk returns a tuple with the ClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimValue

`func (o *GroupListItem) SetClaimValue(v string)`

SetClaimValue sets ClaimValue field to given value.

### HasClaimValue

`func (o *GroupListItem) HasClaimValue() bool`

HasClaimValue returns a boolean if a field has been set.

### SetClaimValueNil

`func (o *GroupListItem) SetClaimValueNil(b bool)`

 SetClaimValueNil sets the value for ClaimValue to be an explicit nil

### UnsetClaimValue
`func (o *GroupListItem) UnsetClaimValue()`

UnsetClaimValue ensures that no value is present for ClaimValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


