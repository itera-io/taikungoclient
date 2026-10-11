# AwsEksClusterDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **NullableString** |  | 
**TargetUrl** | **NullableString** |  | 
**Certificate** | **NullableString** |  | 

## Methods

### NewAwsEksClusterDto

`func NewAwsEksClusterDto(token NullableString, targetUrl NullableString, certificate NullableString, ) *AwsEksClusterDto`

NewAwsEksClusterDto instantiates a new AwsEksClusterDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsEksClusterDtoWithDefaults

`func NewAwsEksClusterDtoWithDefaults() *AwsEksClusterDto`

NewAwsEksClusterDtoWithDefaults instantiates a new AwsEksClusterDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AwsEksClusterDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AwsEksClusterDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AwsEksClusterDto) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *AwsEksClusterDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *AwsEksClusterDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetTargetUrl

`func (o *AwsEksClusterDto) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *AwsEksClusterDto) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *AwsEksClusterDto) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.


### SetTargetUrlNil

`func (o *AwsEksClusterDto) SetTargetUrlNil(b bool)`

 SetTargetUrlNil sets the value for TargetUrl to be an explicit nil

### UnsetTargetUrl
`func (o *AwsEksClusterDto) UnsetTargetUrl()`

UnsetTargetUrl ensures that no value is present for TargetUrl, not even an explicit nil
### GetCertificate

`func (o *AwsEksClusterDto) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AwsEksClusterDto) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AwsEksClusterDto) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### SetCertificateNil

`func (o *AwsEksClusterDto) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *AwsEksClusterDto) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


