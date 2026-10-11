# KubernetesEventCreateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**InvolvedObject** | Pointer to **map[string]interface{}** |  | [optional] 
**FirstTimeStamp** | Pointer to **NullableTime** |  | [optional] 
**LastTimeStamp** | Pointer to **NullableTime** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewKubernetesEventCreateDto

`func NewKubernetesEventCreateDto() *KubernetesEventCreateDto`

NewKubernetesEventCreateDto instantiates a new KubernetesEventCreateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEventCreateDtoWithDefaults

`func NewKubernetesEventCreateDtoWithDefaults() *KubernetesEventCreateDto`

NewKubernetesEventCreateDtoWithDefaults instantiates a new KubernetesEventCreateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KubernetesEventCreateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubernetesEventCreateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubernetesEventCreateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubernetesEventCreateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *KubernetesEventCreateDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *KubernetesEventCreateDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetReason

`func (o *KubernetesEventCreateDto) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubernetesEventCreateDto) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubernetesEventCreateDto) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubernetesEventCreateDto) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *KubernetesEventCreateDto) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *KubernetesEventCreateDto) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMessage

`func (o *KubernetesEventCreateDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubernetesEventCreateDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubernetesEventCreateDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubernetesEventCreateDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *KubernetesEventCreateDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *KubernetesEventCreateDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetMetadata

`func (o *KubernetesEventCreateDto) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubernetesEventCreateDto) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubernetesEventCreateDto) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubernetesEventCreateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSource

`func (o *KubernetesEventCreateDto) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KubernetesEventCreateDto) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KubernetesEventCreateDto) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *KubernetesEventCreateDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetInvolvedObject

`func (o *KubernetesEventCreateDto) GetInvolvedObject() map[string]interface{}`

GetInvolvedObject returns the InvolvedObject field if non-nil, zero value otherwise.

### GetInvolvedObjectOk

`func (o *KubernetesEventCreateDto) GetInvolvedObjectOk() (*map[string]interface{}, bool)`

GetInvolvedObjectOk returns a tuple with the InvolvedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvolvedObject

`func (o *KubernetesEventCreateDto) SetInvolvedObject(v map[string]interface{})`

SetInvolvedObject sets InvolvedObject field to given value.

### HasInvolvedObject

`func (o *KubernetesEventCreateDto) HasInvolvedObject() bool`

HasInvolvedObject returns a boolean if a field has been set.

### GetFirstTimeStamp

`func (o *KubernetesEventCreateDto) GetFirstTimeStamp() time.Time`

GetFirstTimeStamp returns the FirstTimeStamp field if non-nil, zero value otherwise.

### GetFirstTimeStampOk

`func (o *KubernetesEventCreateDto) GetFirstTimeStampOk() (*time.Time, bool)`

GetFirstTimeStampOk returns a tuple with the FirstTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimeStamp

`func (o *KubernetesEventCreateDto) SetFirstTimeStamp(v time.Time)`

SetFirstTimeStamp sets FirstTimeStamp field to given value.

### HasFirstTimeStamp

`func (o *KubernetesEventCreateDto) HasFirstTimeStamp() bool`

HasFirstTimeStamp returns a boolean if a field has been set.

### SetFirstTimeStampNil

`func (o *KubernetesEventCreateDto) SetFirstTimeStampNil(b bool)`

 SetFirstTimeStampNil sets the value for FirstTimeStamp to be an explicit nil

### UnsetFirstTimeStamp
`func (o *KubernetesEventCreateDto) UnsetFirstTimeStamp()`

UnsetFirstTimeStamp ensures that no value is present for FirstTimeStamp, not even an explicit nil
### GetLastTimeStamp

`func (o *KubernetesEventCreateDto) GetLastTimeStamp() time.Time`

GetLastTimeStamp returns the LastTimeStamp field if non-nil, zero value otherwise.

### GetLastTimeStampOk

`func (o *KubernetesEventCreateDto) GetLastTimeStampOk() (*time.Time, bool)`

GetLastTimeStampOk returns a tuple with the LastTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimeStamp

`func (o *KubernetesEventCreateDto) SetLastTimeStamp(v time.Time)`

SetLastTimeStamp sets LastTimeStamp field to given value.

### HasLastTimeStamp

`func (o *KubernetesEventCreateDto) HasLastTimeStamp() bool`

HasLastTimeStamp returns a boolean if a field has been set.

### SetLastTimeStampNil

`func (o *KubernetesEventCreateDto) SetLastTimeStampNil(b bool)`

 SetLastTimeStampNil sets the value for LastTimeStamp to be an explicit nil

### UnsetLastTimeStamp
`func (o *KubernetesEventCreateDto) UnsetLastTimeStamp()`

UnsetLastTimeStamp ensures that no value is present for LastTimeStamp, not even an explicit nil
### GetCount

`func (o *KubernetesEventCreateDto) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KubernetesEventCreateDto) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KubernetesEventCreateDto) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KubernetesEventCreateDto) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


