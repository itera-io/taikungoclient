# LokiLogsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to **NullableString** |  | [optional] 
**Filters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 
**Direction** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLokiLogsQuery

`func NewLokiLogsQuery() *LokiLogsQuery`

NewLokiLogsQuery instantiates a new LokiLogsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLokiLogsQueryWithDefaults

`func NewLokiLogsQueryWithDefaults() *LokiLogsQuery`

NewLokiLogsQueryWithDefaults instantiates a new LokiLogsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *LokiLogsQuery) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LokiLogsQuery) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LokiLogsQuery) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LokiLogsQuery) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParameters

`func (o *LokiLogsQuery) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LokiLogsQuery) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LokiLogsQuery) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *LokiLogsQuery) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *LokiLogsQuery) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *LokiLogsQuery) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetFilters

`func (o *LokiLogsQuery) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LokiLogsQuery) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LokiLogsQuery) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *LokiLogsQuery) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *LokiLogsQuery) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *LokiLogsQuery) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetStartDate

`func (o *LokiLogsQuery) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LokiLogsQuery) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LokiLogsQuery) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LokiLogsQuery) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *LokiLogsQuery) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *LokiLogsQuery) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *LokiLogsQuery) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *LokiLogsQuery) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *LokiLogsQuery) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *LokiLogsQuery) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *LokiLogsQuery) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *LokiLogsQuery) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetLimit

`func (o *LokiLogsQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LokiLogsQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LokiLogsQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LokiLogsQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *LokiLogsQuery) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *LokiLogsQuery) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetDirection

`func (o *LokiLogsQuery) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LokiLogsQuery) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LokiLogsQuery) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LokiLogsQuery) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *LokiLogsQuery) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *LokiLogsQuery) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


