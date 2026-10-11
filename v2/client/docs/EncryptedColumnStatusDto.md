# EncryptedColumnStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityName** | **NullableString** |  | 
**V1RowCount** | **int64** |  | 
**OldestV1CreatedUtc** | **NullableTime** |  | 
**SurveyError** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEncryptedColumnStatusDto

`func NewEncryptedColumnStatusDto(entityName NullableString, v1RowCount int64, oldestV1CreatedUtc NullableTime, ) *EncryptedColumnStatusDto`

NewEncryptedColumnStatusDto instantiates a new EncryptedColumnStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptedColumnStatusDtoWithDefaults

`func NewEncryptedColumnStatusDtoWithDefaults() *EncryptedColumnStatusDto`

NewEncryptedColumnStatusDtoWithDefaults instantiates a new EncryptedColumnStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityName

`func (o *EncryptedColumnStatusDto) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *EncryptedColumnStatusDto) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *EncryptedColumnStatusDto) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.


### SetEntityNameNil

`func (o *EncryptedColumnStatusDto) SetEntityNameNil(b bool)`

 SetEntityNameNil sets the value for EntityName to be an explicit nil

### UnsetEntityName
`func (o *EncryptedColumnStatusDto) UnsetEntityName()`

UnsetEntityName ensures that no value is present for EntityName, not even an explicit nil
### GetV1RowCount

`func (o *EncryptedColumnStatusDto) GetV1RowCount() int64`

GetV1RowCount returns the V1RowCount field if non-nil, zero value otherwise.

### GetV1RowCountOk

`func (o *EncryptedColumnStatusDto) GetV1RowCountOk() (*int64, bool)`

GetV1RowCountOk returns a tuple with the V1RowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1RowCount

`func (o *EncryptedColumnStatusDto) SetV1RowCount(v int64)`

SetV1RowCount sets V1RowCount field to given value.


### GetOldestV1CreatedUtc

`func (o *EncryptedColumnStatusDto) GetOldestV1CreatedUtc() time.Time`

GetOldestV1CreatedUtc returns the OldestV1CreatedUtc field if non-nil, zero value otherwise.

### GetOldestV1CreatedUtcOk

`func (o *EncryptedColumnStatusDto) GetOldestV1CreatedUtcOk() (*time.Time, bool)`

GetOldestV1CreatedUtcOk returns a tuple with the OldestV1CreatedUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestV1CreatedUtc

`func (o *EncryptedColumnStatusDto) SetOldestV1CreatedUtc(v time.Time)`

SetOldestV1CreatedUtc sets OldestV1CreatedUtc field to given value.


### SetOldestV1CreatedUtcNil

`func (o *EncryptedColumnStatusDto) SetOldestV1CreatedUtcNil(b bool)`

 SetOldestV1CreatedUtcNil sets the value for OldestV1CreatedUtc to be an explicit nil

### UnsetOldestV1CreatedUtc
`func (o *EncryptedColumnStatusDto) UnsetOldestV1CreatedUtc()`

UnsetOldestV1CreatedUtc ensures that no value is present for OldestV1CreatedUtc, not even an explicit nil
### GetSurveyError

`func (o *EncryptedColumnStatusDto) GetSurveyError() string`

GetSurveyError returns the SurveyError field if non-nil, zero value otherwise.

### GetSurveyErrorOk

`func (o *EncryptedColumnStatusDto) GetSurveyErrorOk() (*string, bool)`

GetSurveyErrorOk returns a tuple with the SurveyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyError

`func (o *EncryptedColumnStatusDto) SetSurveyError(v string)`

SetSurveyError sets SurveyError field to given value.

### HasSurveyError

`func (o *EncryptedColumnStatusDto) HasSurveyError() bool`

HasSurveyError returns a boolean if a field has been set.

### SetSurveyErrorNil

`func (o *EncryptedColumnStatusDto) SetSurveyErrorNil(b bool)`

 SetSurveyErrorNil sets the value for SurveyError to be an explicit nil

### UnsetSurveyError
`func (o *EncryptedColumnStatusDto) UnsetSurveyError()`

UnsetSurveyError ensures that no value is present for SurveyError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


