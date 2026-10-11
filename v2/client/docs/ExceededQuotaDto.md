# ExceededQuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudId** | **int32** |  | 
**Name** | **NullableString** |  | 

## Methods

### NewExceededQuotaDto

`func NewExceededQuotaDto(cloudId int32, name NullableString, ) *ExceededQuotaDto`

NewExceededQuotaDto instantiates a new ExceededQuotaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceededQuotaDtoWithDefaults

`func NewExceededQuotaDtoWithDefaults() *ExceededQuotaDto`

NewExceededQuotaDtoWithDefaults instantiates a new ExceededQuotaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudId

`func (o *ExceededQuotaDto) GetCloudId() int32`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *ExceededQuotaDto) GetCloudIdOk() (*int32, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *ExceededQuotaDto) SetCloudId(v int32)`

SetCloudId sets CloudId field to given value.


### GetName

`func (o *ExceededQuotaDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExceededQuotaDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExceededQuotaDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ExceededQuotaDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExceededQuotaDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


