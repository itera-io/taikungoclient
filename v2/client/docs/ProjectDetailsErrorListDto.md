# ProjectDetailsErrorListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProjectDetailsErrorType**](ProjectDetailsErrorType.md) |  | 
**Message** | **[]string** |  | 
**Kind** | [**ProjectType**](ProjectType.md) |  | 

## Methods

### NewProjectDetailsErrorListDto

`func NewProjectDetailsErrorListDto(type_ ProjectDetailsErrorType, message []string, kind ProjectType, ) *ProjectDetailsErrorListDto`

NewProjectDetailsErrorListDto instantiates a new ProjectDetailsErrorListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectDetailsErrorListDtoWithDefaults

`func NewProjectDetailsErrorListDtoWithDefaults() *ProjectDetailsErrorListDto`

NewProjectDetailsErrorListDtoWithDefaults instantiates a new ProjectDetailsErrorListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProjectDetailsErrorListDto) GetType() ProjectDetailsErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectDetailsErrorListDto) GetTypeOk() (*ProjectDetailsErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectDetailsErrorListDto) SetType(v ProjectDetailsErrorType)`

SetType sets Type field to given value.


### GetMessage

`func (o *ProjectDetailsErrorListDto) GetMessage() []string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProjectDetailsErrorListDto) GetMessageOk() (*[]string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProjectDetailsErrorListDto) SetMessage(v []string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ProjectDetailsErrorListDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ProjectDetailsErrorListDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetKind

`func (o *ProjectDetailsErrorListDto) GetKind() ProjectType`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProjectDetailsErrorListDto) GetKindOk() (*ProjectType, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProjectDetailsErrorListDto) SetKind(v ProjectType)`

SetKind sets Kind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


