# InteractiveShellDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfig** | **NullableString** |  | 
**AdminKubeConfig** | **NullableString** |  | 

## Methods

### NewInteractiveShellDto

`func NewInteractiveShellDto(kubeConfig NullableString, adminKubeConfig NullableString, ) *InteractiveShellDto`

NewInteractiveShellDto instantiates a new InteractiveShellDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractiveShellDtoWithDefaults

`func NewInteractiveShellDtoWithDefaults() *InteractiveShellDto`

NewInteractiveShellDtoWithDefaults instantiates a new InteractiveShellDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfig

`func (o *InteractiveShellDto) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *InteractiveShellDto) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *InteractiveShellDto) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.


### SetKubeConfigNil

`func (o *InteractiveShellDto) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *InteractiveShellDto) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetAdminKubeConfig

`func (o *InteractiveShellDto) GetAdminKubeConfig() string`

GetAdminKubeConfig returns the AdminKubeConfig field if non-nil, zero value otherwise.

### GetAdminKubeConfigOk

`func (o *InteractiveShellDto) GetAdminKubeConfigOk() (*string, bool)`

GetAdminKubeConfigOk returns a tuple with the AdminKubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKubeConfig

`func (o *InteractiveShellDto) SetAdminKubeConfig(v string)`

SetAdminKubeConfig sets AdminKubeConfig field to given value.


### SetAdminKubeConfigNil

`func (o *InteractiveShellDto) SetAdminKubeConfigNil(b bool)`

 SetAdminKubeConfigNil sets the value for AdminKubeConfig to be an explicit nil

### UnsetAdminKubeConfig
`func (o *InteractiveShellDto) UnsetAdminKubeConfig()`

UnsetAdminKubeConfig ensures that no value is present for AdminKubeConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


