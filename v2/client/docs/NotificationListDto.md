# NotificationListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **NullableString** |  | 
**ActionMessage** | **string** |  | 
**ActionStatus** | [**ActionStatus**](ActionStatus.md) |  | 
**Username** | [**AuditUserDto**](AuditUserDto.md) |  | 
**Category** | [**ActionType**](ActionType.md) |  | 
**ProjectName** | **NullableString** |  | 
**ProjectId** | **NullableInt32** |  | 
**IsDeleted** | **bool** |  | 

## Methods

### NewNotificationListDto

`func NewNotificationListDto(createdAt NullableString, actionMessage string, actionStatus ActionStatus, username AuditUserDto, category ActionType, projectName NullableString, projectId NullableInt32, isDeleted bool, ) *NotificationListDto`

NewNotificationListDto instantiates a new NotificationListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListDtoWithDefaults

`func NewNotificationListDtoWithDefaults() *NotificationListDto`

NewNotificationListDtoWithDefaults instantiates a new NotificationListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NotificationListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *NotificationListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *NotificationListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetActionMessage

`func (o *NotificationListDto) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *NotificationListDto) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *NotificationListDto) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.


### GetActionStatus

`func (o *NotificationListDto) GetActionStatus() ActionStatus`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *NotificationListDto) GetActionStatusOk() (*ActionStatus, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *NotificationListDto) SetActionStatus(v ActionStatus)`

SetActionStatus sets ActionStatus field to given value.


### GetUsername

`func (o *NotificationListDto) GetUsername() AuditUserDto`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationListDto) GetUsernameOk() (*AuditUserDto, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationListDto) SetUsername(v AuditUserDto)`

SetUsername sets Username field to given value.


### GetCategory

`func (o *NotificationListDto) GetCategory() ActionType`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NotificationListDto) GetCategoryOk() (*ActionType, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NotificationListDto) SetCategory(v ActionType)`

SetCategory sets Category field to given value.


### GetProjectName

`func (o *NotificationListDto) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *NotificationListDto) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *NotificationListDto) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### SetProjectNameNil

`func (o *NotificationListDto) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *NotificationListDto) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetProjectId

`func (o *NotificationListDto) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *NotificationListDto) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *NotificationListDto) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *NotificationListDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *NotificationListDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIsDeleted

`func (o *NotificationListDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *NotificationListDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *NotificationListDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


