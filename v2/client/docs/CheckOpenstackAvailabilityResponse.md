# CheckOpenstackAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsOpenstackAvailable** | **bool** |  | 
**IsManilaAvailable** | **bool** |  | 
**ErrorMessage** | **NullableString** |  | 

## Methods

### NewCheckOpenstackAvailabilityResponse

`func NewCheckOpenstackAvailabilityResponse(isOpenstackAvailable bool, isManilaAvailable bool, errorMessage NullableString, ) *CheckOpenstackAvailabilityResponse`

NewCheckOpenstackAvailabilityResponse instantiates a new CheckOpenstackAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckOpenstackAvailabilityResponseWithDefaults

`func NewCheckOpenstackAvailabilityResponseWithDefaults() *CheckOpenstackAvailabilityResponse`

NewCheckOpenstackAvailabilityResponseWithDefaults instantiates a new CheckOpenstackAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsOpenstackAvailable

`func (o *CheckOpenstackAvailabilityResponse) GetIsOpenstackAvailable() bool`

GetIsOpenstackAvailable returns the IsOpenstackAvailable field if non-nil, zero value otherwise.

### GetIsOpenstackAvailableOk

`func (o *CheckOpenstackAvailabilityResponse) GetIsOpenstackAvailableOk() (*bool, bool)`

GetIsOpenstackAvailableOk returns a tuple with the IsOpenstackAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenstackAvailable

`func (o *CheckOpenstackAvailabilityResponse) SetIsOpenstackAvailable(v bool)`

SetIsOpenstackAvailable sets IsOpenstackAvailable field to given value.


### GetIsManilaAvailable

`func (o *CheckOpenstackAvailabilityResponse) GetIsManilaAvailable() bool`

GetIsManilaAvailable returns the IsManilaAvailable field if non-nil, zero value otherwise.

### GetIsManilaAvailableOk

`func (o *CheckOpenstackAvailabilityResponse) GetIsManilaAvailableOk() (*bool, bool)`

GetIsManilaAvailableOk returns a tuple with the IsManilaAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManilaAvailable

`func (o *CheckOpenstackAvailabilityResponse) SetIsManilaAvailable(v bool)`

SetIsManilaAvailable sets IsManilaAvailable field to given value.


### GetErrorMessage

`func (o *CheckOpenstackAvailabilityResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CheckOpenstackAvailabilityResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CheckOpenstackAvailabilityResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *CheckOpenstackAvailabilityResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *CheckOpenstackAvailabilityResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


