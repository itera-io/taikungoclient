# CreateAwsSubnetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsSubnetId** | Pointer to **NullableString** |  | [optional] 
**SubnetType** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateAwsSubnetDto

`func NewCreateAwsSubnetDto() *CreateAwsSubnetDto`

NewCreateAwsSubnetDto instantiates a new CreateAwsSubnetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAwsSubnetDtoWithDefaults

`func NewCreateAwsSubnetDtoWithDefaults() *CreateAwsSubnetDto`

NewCreateAwsSubnetDtoWithDefaults instantiates a new CreateAwsSubnetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsSubnetId

`func (o *CreateAwsSubnetDto) GetAwsSubnetId() string`

GetAwsSubnetId returns the AwsSubnetId field if non-nil, zero value otherwise.

### GetAwsSubnetIdOk

`func (o *CreateAwsSubnetDto) GetAwsSubnetIdOk() (*string, bool)`

GetAwsSubnetIdOk returns a tuple with the AwsSubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSubnetId

`func (o *CreateAwsSubnetDto) SetAwsSubnetId(v string)`

SetAwsSubnetId sets AwsSubnetId field to given value.

### HasAwsSubnetId

`func (o *CreateAwsSubnetDto) HasAwsSubnetId() bool`

HasAwsSubnetId returns a boolean if a field has been set.

### SetAwsSubnetIdNil

`func (o *CreateAwsSubnetDto) SetAwsSubnetIdNil(b bool)`

 SetAwsSubnetIdNil sets the value for AwsSubnetId to be an explicit nil

### UnsetAwsSubnetId
`func (o *CreateAwsSubnetDto) UnsetAwsSubnetId()`

UnsetAwsSubnetId ensures that no value is present for AwsSubnetId, not even an explicit nil
### GetSubnetType

`func (o *CreateAwsSubnetDto) GetSubnetType() map[string]interface{}`

GetSubnetType returns the SubnetType field if non-nil, zero value otherwise.

### GetSubnetTypeOk

`func (o *CreateAwsSubnetDto) GetSubnetTypeOk() (*map[string]interface{}, bool)`

GetSubnetTypeOk returns a tuple with the SubnetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetType

`func (o *CreateAwsSubnetDto) SetSubnetType(v map[string]interface{})`

SetSubnetType sets SubnetType field to given value.

### HasSubnetType

`func (o *CreateAwsSubnetDto) HasSubnetType() bool`

HasSubnetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


