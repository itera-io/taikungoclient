# AuditUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **NullableString** |  | 
**DisplayName** | **NullableString** |  | 
**DomainId** | **NullableInt32** |  | 
**IsRobot** | **bool** |  | 

## Methods

### NewAuditUserDto

`func NewAuditUserDto(userId NullableString, displayName NullableString, domainId NullableInt32, isRobot bool, ) *AuditUserDto`

NewAuditUserDto instantiates a new AuditUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditUserDtoWithDefaults

`func NewAuditUserDtoWithDefaults() *AuditUserDto`

NewAuditUserDtoWithDefaults instantiates a new AuditUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AuditUserDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditUserDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditUserDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *AuditUserDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AuditUserDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDisplayName

`func (o *AuditUserDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuditUserDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuditUserDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *AuditUserDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AuditUserDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDomainId

`func (o *AuditUserDto) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *AuditUserDto) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *AuditUserDto) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.


### SetDomainIdNil

`func (o *AuditUserDto) SetDomainIdNil(b bool)`

 SetDomainIdNil sets the value for DomainId to be an explicit nil

### UnsetDomainId
`func (o *AuditUserDto) UnsetDomainId()`

UnsetDomainId ensures that no value is present for DomainId, not even an explicit nil
### GetIsRobot

`func (o *AuditUserDto) GetIsRobot() bool`

GetIsRobot returns the IsRobot field if non-nil, zero value otherwise.

### GetIsRobotOk

`func (o *AuditUserDto) GetIsRobotOk() (*bool, bool)`

GetIsRobotOk returns a tuple with the IsRobot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRobot

`func (o *AuditUserDto) SetIsRobot(v bool)`

SetIsRobot sets IsRobot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


