# EditRobotUserCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**ExpiresAt** | **NullableTime** |  | 
**Ips** | **[]string** |  | 

## Methods

### NewEditRobotUserCommand

`func NewEditRobotUserCommand(id NullableString, name NullableString, description NullableString, expiresAt NullableTime, ips []string, ) *EditRobotUserCommand`

NewEditRobotUserCommand instantiates a new EditRobotUserCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditRobotUserCommandWithDefaults

`func NewEditRobotUserCommandWithDefaults() *EditRobotUserCommand`

NewEditRobotUserCommandWithDefaults instantiates a new EditRobotUserCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditRobotUserCommand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditRobotUserCommand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditRobotUserCommand) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *EditRobotUserCommand) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EditRobotUserCommand) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *EditRobotUserCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditRobotUserCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditRobotUserCommand) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *EditRobotUserCommand) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditRobotUserCommand) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *EditRobotUserCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditRobotUserCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditRobotUserCommand) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EditRobotUserCommand) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditRobotUserCommand) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpiresAt

`func (o *EditRobotUserCommand) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EditRobotUserCommand) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EditRobotUserCommand) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *EditRobotUserCommand) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *EditRobotUserCommand) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetIps

`func (o *EditRobotUserCommand) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *EditRobotUserCommand) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *EditRobotUserCommand) SetIps(v []string)`

SetIps sets Ips field to given value.


### SetIpsNil

`func (o *EditRobotUserCommand) SetIpsNil(b bool)`

 SetIpsNil sets the value for Ips to be an explicit nil

### UnsetIps
`func (o *EditRobotUserCommand) UnsetIps()`

UnsetIps ensures that no value is present for Ips, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


