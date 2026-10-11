# ProjectsForBillingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 
**BillingStartDate** | **NullableTime** |  | 
**OrganizationName** | **NullableString** |  | 
**Price** | **float64** |  | 
**Servers** | [**[]ServersForBillingDto**](ServersForBillingDto.md) |  | 
**StandaloneVms** | [**[]StandaloneVmsForBillingDto**](StandaloneVmsForBillingDto.md) |  | 
**BillingEnabled** | **bool** |  | 

## Methods

### NewProjectsForBillingDto

`func NewProjectsForBillingDto(id int32, name NullableString, createdAt NullableTime, billingStartDate NullableTime, organizationName NullableString, price float64, servers []ServersForBillingDto, standaloneVms []StandaloneVmsForBillingDto, billingEnabled bool, ) *ProjectsForBillingDto`

NewProjectsForBillingDto instantiates a new ProjectsForBillingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsForBillingDtoWithDefaults

`func NewProjectsForBillingDtoWithDefaults() *ProjectsForBillingDto`

NewProjectsForBillingDtoWithDefaults instantiates a new ProjectsForBillingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectsForBillingDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectsForBillingDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectsForBillingDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ProjectsForBillingDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsForBillingDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsForBillingDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *ProjectsForBillingDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ProjectsForBillingDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *ProjectsForBillingDto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectsForBillingDto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectsForBillingDto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ProjectsForBillingDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ProjectsForBillingDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetBillingStartDate

`func (o *ProjectsForBillingDto) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *ProjectsForBillingDto) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *ProjectsForBillingDto) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.


### SetBillingStartDateNil

`func (o *ProjectsForBillingDto) SetBillingStartDateNil(b bool)`

 SetBillingStartDateNil sets the value for BillingStartDate to be an explicit nil

### UnsetBillingStartDate
`func (o *ProjectsForBillingDto) UnsetBillingStartDate()`

UnsetBillingStartDate ensures that no value is present for BillingStartDate, not even an explicit nil
### GetOrganizationName

`func (o *ProjectsForBillingDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ProjectsForBillingDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ProjectsForBillingDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.


### SetOrganizationNameNil

`func (o *ProjectsForBillingDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ProjectsForBillingDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetPrice

`func (o *ProjectsForBillingDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProjectsForBillingDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProjectsForBillingDto) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetServers

`func (o *ProjectsForBillingDto) GetServers() []ServersForBillingDto`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ProjectsForBillingDto) GetServersOk() (*[]ServersForBillingDto, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ProjectsForBillingDto) SetServers(v []ServersForBillingDto)`

SetServers sets Servers field to given value.


### SetServersNil

`func (o *ProjectsForBillingDto) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *ProjectsForBillingDto) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetStandaloneVms

`func (o *ProjectsForBillingDto) GetStandaloneVms() []StandaloneVmsForBillingDto`

GetStandaloneVms returns the StandaloneVms field if non-nil, zero value otherwise.

### GetStandaloneVmsOk

`func (o *ProjectsForBillingDto) GetStandaloneVmsOk() (*[]StandaloneVmsForBillingDto, bool)`

GetStandaloneVmsOk returns a tuple with the StandaloneVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneVms

`func (o *ProjectsForBillingDto) SetStandaloneVms(v []StandaloneVmsForBillingDto)`

SetStandaloneVms sets StandaloneVms field to given value.


### SetStandaloneVmsNil

`func (o *ProjectsForBillingDto) SetStandaloneVmsNil(b bool)`

 SetStandaloneVmsNil sets the value for StandaloneVms to be an explicit nil

### UnsetStandaloneVms
`func (o *ProjectsForBillingDto) UnsetStandaloneVms()`

UnsetStandaloneVms ensures that no value is present for StandaloneVms, not even an explicit nil
### GetBillingEnabled

`func (o *ProjectsForBillingDto) GetBillingEnabled() bool`

GetBillingEnabled returns the BillingEnabled field if non-nil, zero value otherwise.

### GetBillingEnabledOk

`func (o *ProjectsForBillingDto) GetBillingEnabledOk() (*bool, bool)`

GetBillingEnabledOk returns a tuple with the BillingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEnabled

`func (o *ProjectsForBillingDto) SetBillingEnabled(v bool)`

SetBillingEnabled sets BillingEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


