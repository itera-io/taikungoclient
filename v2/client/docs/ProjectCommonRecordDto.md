# ProjectCommonRecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**ExpiredAt** | **NullableString** |  | 

## Methods

### NewProjectCommonRecordDto

`func NewProjectCommonRecordDto(id int32, name NullableString, expiredAt NullableString, ) *ProjectCommonRecordDto`

NewProjectCommonRecordDto instantiates a new ProjectCommonRecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCommonRecordDtoWithDefaults

`func NewProjectCommonRecordDtoWithDefaults() *ProjectCommonRecordDto`

NewProjectCommonRecordDtoWithDefaults instantiates a new ProjectCommonRecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectCommonRecordDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCommonRecordDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCommonRecordDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectCommonRecordDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCommonRecordDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCommonRecordDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProjectCommonRecordDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectCommonRecordDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetExpiredAt

`func (o *ProjectCommonRecordDto) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProjectCommonRecordDto) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProjectCommonRecordDto) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *ProjectCommonRecordDto) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ProjectCommonRecordDto) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


