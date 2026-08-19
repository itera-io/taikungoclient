# AiCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AiCredentialsListDto**](AiCredentialsListDto.md) |  | 
**TotalCount** | **int32** |  | 

## Methods

### NewAiCredentials

`func NewAiCredentials(data []AiCredentialsListDto, totalCount int32, ) *AiCredentials`

NewAiCredentials instantiates a new AiCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiCredentialsWithDefaults

`func NewAiCredentialsWithDefaults() *AiCredentials`

NewAiCredentialsWithDefaults instantiates a new AiCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AiCredentials) GetData() []AiCredentialsListDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiCredentials) GetDataOk() (*[]AiCredentialsListDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiCredentials) SetData(v []AiCredentialsListDto)`

SetData sets Data field to given value.


### SetDataNil

`func (o *AiCredentials) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AiCredentials) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTotalCount

`func (o *AiCredentials) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AiCredentials) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AiCredentials) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


