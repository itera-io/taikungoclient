# StandAloneProfilesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**PublicKey** | **NullableString** |  | 
**IsLocked** | **bool** |  | 
**StandaloneVms** | [**[]StandAloneVmSmallDetailDto**](StandAloneVmSmallDetailDto.md) |  | 
**OrganizationId** | **int32** |  | 
**OrganizationName** | **NullableString** |  | 
**PartnerLogo** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 

## Methods

### NewStandAloneProfilesListDto

`func NewStandAloneProfilesListDto(id int32, name NullableString, publicKey NullableString, isLocked bool, standaloneVms []StandAloneVmSmallDetailDto, organizationId int32, organizationName NullableString, partnerLogo NullableString, createdAt NullableString, ) *StandAloneProfilesListDto`

NewStandAloneProfilesListDto instantiates a new StandAloneProfilesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandAloneProfilesListDtoWithDefaults

`func NewStandAloneProfilesListDtoWithDefaults() *StandAloneProfilesListDto`

NewStandAloneProfilesListDtoWithDefaults instantiates a new StandAloneProfilesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandAloneProfilesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandAloneProfilesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandAloneProfilesListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *StandAloneProfilesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandAloneProfilesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandAloneProfilesListDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *StandAloneProfilesListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StandAloneProfilesListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublicKey

`func (o *StandAloneProfilesListDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *StandAloneProfilesListDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *StandAloneProfilesListDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### SetPublicKeyNil

`func (o *StandAloneProfilesListDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *StandAloneProfilesListDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetIsLocked

`func (o *StandAloneProfilesListDto) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *StandAloneProfilesListDto) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *StandAloneProfilesListDto) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetStandaloneVms

`func (o *StandAloneProfilesListDto) GetStandaloneVms() []StandAloneVmSmallDetailDto`

GetStandaloneVms returns the StandaloneVms field if non-nil, zero value otherwise.

### GetStandaloneVmsOk

`func (o *StandAloneProfilesListDto) GetStandaloneVmsOk() (*[]StandAloneVmSmallDetailDto, bool)`

GetStandaloneVmsOk returns a tuple with the StandaloneVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVms

`func (o *StandAloneProfilesListDto) SetStandaloneVms(v []StandAloneVmSmallDetailDto)`

SetStandaloneVms sets StandaloneVms field to given value.


### SetStandaloneVmsNil

`func (o *StandAloneProfilesListDto) SetStandaloneVmsNil(b bool)`

 SetStandaloneVmsNil sets the value for StandaloneVms to be an explicit nil

### UnsetStandaloneVms
`func (o *StandAloneProfilesListDto) UnsetStandaloneVms()`

UnsetStandaloneVms ensures that no value is present for StandaloneVms, not even an explicit nil
### GetOrganizationId

`func (o *StandAloneProfilesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *StandAloneProfilesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *StandAloneProfilesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOrganizationName

`func (o *StandAloneProfilesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *StandAloneProfilesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *StandAloneProfilesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *StandAloneProfilesListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *StandAloneProfilesListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetPartnerLogo

`func (o *StandAloneProfilesListDto) GetPartnerLogo() string`

GetPartnerLogo returns the PartnerLogo field if non-nil, zero value otherwise.

### GetPartnerLogoOk

`func (o *StandAloneProfilesListDto) GetPartnerLogoOk() (*string, bool)`

GetPartnerLogoOk returns a tuple with the PartnerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerLogo

`func (o *StandAloneProfilesListDto) SetPartnerLogo(v string)`

SetPartnerLogo sets PartnerLogo field to given value.


### SetPartnerLogoNil

`func (o *StandAloneProfilesListDto) SetPartnerLogoNil(b bool)`

 SetPartnerLogoNil sets the value for PartnerLogo to be an explicit nil

### UnsetPartnerLogo
`func (o *StandAloneProfilesListDto) UnsetPartnerLogo()`

UnsetPartnerLogo ensures that no value is present for PartnerLogo, not even an explicit nil
### GetCreatedAt

`func (o *StandAloneProfilesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StandAloneProfilesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StandAloneProfilesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *StandAloneProfilesListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StandAloneProfilesListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


