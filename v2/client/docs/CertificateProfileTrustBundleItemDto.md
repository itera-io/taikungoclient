# CertificateProfileTrustBundleItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**PemBundle** | **string** |  | 

## Methods

### NewCertificateProfileTrustBundleItemDto

`func NewCertificateProfileTrustBundleItemDto(id int32, name string, pemBundle string, ) *CertificateProfileTrustBundleItemDto`

NewCertificateProfileTrustBundleItemDto instantiates a new CertificateProfileTrustBundleItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateProfileTrustBundleItemDtoWithDefaults

`func NewCertificateProfileTrustBundleItemDtoWithDefaults() *CertificateProfileTrustBundleItemDto`

NewCertificateProfileTrustBundleItemDtoWithDefaults instantiates a new CertificateProfileTrustBundleItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateProfileTrustBundleItemDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateProfileTrustBundleItemDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateProfileTrustBundleItemDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CertificateProfileTrustBundleItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateProfileTrustBundleItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateProfileTrustBundleItemDto) SetName(v string)`

SetName sets Name field to given value.


### GetPemBundle

`func (o *CertificateProfileTrustBundleItemDto) GetPemBundle() string`

GetPemBundle returns the PemBundle field if non-nil, zero value otherwise.

### GetPemBundleOk

`func (o *CertificateProfileTrustBundleItemDto) GetPemBundleOk() (*string, bool)`

GetPemBundleOk returns a tuple with the PemBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPemBundle

`func (o *CertificateProfileTrustBundleItemDto) SetPemBundle(v string)`

SetPemBundle sets PemBundle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


