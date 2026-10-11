# ShowbackRulesListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ByLabel** | Pointer to **NullableString** |  | [optional] 
**MetricName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Kind** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**ProjectAlertLimit** | Pointer to **int32** |  | [optional] 
**GlobalAlertLimit** | Pointer to **int32** |  | [optional] 
**OrganizationName** | Pointer to **NullableString** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**BillingStartDate** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to [**[]ShowbackLabelCreateDto**](ShowbackLabelCreateDto.md) |  | [optional] 
**ShowbackCredentialName** | Pointer to **NullableString** |  | [optional] 
**ShowbackCredentialId** | Pointer to **NullableInt32** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**LastModified** | Pointer to **NullableString** |  | [optional] 
**LastModifiedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewShowbackRulesListDto

`func NewShowbackRulesListDto() *ShowbackRulesListDto`

NewShowbackRulesListDto instantiates a new ShowbackRulesListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowbackRulesListDtoWithDefaults

`func NewShowbackRulesListDtoWithDefaults() *ShowbackRulesListDto`

NewShowbackRulesListDtoWithDefaults instantiates a new ShowbackRulesListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShowbackRulesListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShowbackRulesListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShowbackRulesListDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShowbackRulesListDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShowbackRulesListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShowbackRulesListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShowbackRulesListDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShowbackRulesListDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ShowbackRulesListDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ShowbackRulesListDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetByLabel

`func (o *ShowbackRulesListDto) GetByLabel() string`

GetByLabel returns the ByLabel field if non-nil, zero value otherwise.

### GetByLabelOk

`func (o *ShowbackRulesListDto) GetByLabelOk() (*string, bool)`

GetByLabelOk returns a tuple with the ByLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByLabel

`func (o *ShowbackRulesListDto) SetByLabel(v string)`

SetByLabel sets ByLabel field to given value.

### HasByLabel

`func (o *ShowbackRulesListDto) HasByLabel() bool`

HasByLabel returns a boolean if a field has been set.

### SetByLabelNil

`func (o *ShowbackRulesListDto) SetByLabelNil(b bool)`

 SetByLabelNil sets the value for ByLabel to be an explicit nil

### UnsetByLabel
`func (o *ShowbackRulesListDto) UnsetByLabel()`

UnsetByLabel ensures that no value is present for ByLabel, not even an explicit nil
### GetMetricName

`func (o *ShowbackRulesListDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *ShowbackRulesListDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *ShowbackRulesListDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *ShowbackRulesListDto) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### SetMetricNameNil

`func (o *ShowbackRulesListDto) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *ShowbackRulesListDto) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetCreatedAt

`func (o *ShowbackRulesListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShowbackRulesListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShowbackRulesListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ShowbackRulesListDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ShowbackRulesListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ShowbackRulesListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetKind

`func (o *ShowbackRulesListDto) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ShowbackRulesListDto) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ShowbackRulesListDto) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ShowbackRulesListDto) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *ShowbackRulesListDto) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *ShowbackRulesListDto) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetType

`func (o *ShowbackRulesListDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShowbackRulesListDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShowbackRulesListDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShowbackRulesListDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ShowbackRulesListDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShowbackRulesListDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPrice

`func (o *ShowbackRulesListDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShowbackRulesListDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShowbackRulesListDto) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ShowbackRulesListDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProjectAlertLimit

`func (o *ShowbackRulesListDto) GetProjectAlertLimit() int32`

GetProjectAlertLimit returns the ProjectAlertLimit field if non-nil, zero value otherwise.

### GetProjectAlertLimitOk

`func (o *ShowbackRulesListDto) GetProjectAlertLimitOk() (*int32, bool)`

GetProjectAlertLimitOk returns a tuple with the ProjectAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAlertLimit

`func (o *ShowbackRulesListDto) SetProjectAlertLimit(v int32)`

SetProjectAlertLimit sets ProjectAlertLimit field to given value.

### HasProjectAlertLimit

`func (o *ShowbackRulesListDto) HasProjectAlertLimit() bool`

HasProjectAlertLimit returns a boolean if a field has been set.

### GetGlobalAlertLimit

`func (o *ShowbackRulesListDto) GetGlobalAlertLimit() int32`

GetGlobalAlertLimit returns the GlobalAlertLimit field if non-nil, zero value otherwise.

### GetGlobalAlertLimitOk

`func (o *ShowbackRulesListDto) GetGlobalAlertLimitOk() (*int32, bool)`

GetGlobalAlertLimitOk returns a tuple with the GlobalAlertLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAlertLimit

`func (o *ShowbackRulesListDto) SetGlobalAlertLimit(v int32)`

SetGlobalAlertLimit sets GlobalAlertLimit field to given value.

### HasGlobalAlertLimit

`func (o *ShowbackRulesListDto) HasGlobalAlertLimit() bool`

HasGlobalAlertLimit returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ShowbackRulesListDto) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ShowbackRulesListDto) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ShowbackRulesListDto) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ShowbackRulesListDto) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *ShowbackRulesListDto) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *ShowbackRulesListDto) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetOrganizationId

`func (o *ShowbackRulesListDto) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ShowbackRulesListDto) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ShowbackRulesListDto) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ShowbackRulesListDto) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *ShowbackRulesListDto) GetBillingStartDate() string`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *ShowbackRulesListDto) GetBillingStartDateOk() (*string, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *ShowbackRulesListDto) SetBillingStartDate(v string)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *ShowbackRulesListDto) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### SetBillingStartDateNil

`func (o *ShowbackRulesListDto) SetBillingStartDateNil(b bool)`

 SetBillingStartDateNil sets the value for BillingStartDate to be an explicit nil

### UnsetBillingStartDate
`func (o *ShowbackRulesListDto) UnsetBillingStartDate()`

UnsetBillingStartDate ensures that no value is present for BillingStartDate, not even an explicit nil
### GetLabels

`func (o *ShowbackRulesListDto) GetLabels() []ShowbackLabelCreateDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ShowbackRulesListDto) GetLabelsOk() (*[]ShowbackLabelCreateDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ShowbackRulesListDto) SetLabels(v []ShowbackLabelCreateDto)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ShowbackRulesListDto) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ShowbackRulesListDto) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ShowbackRulesListDto) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetShowbackCredentialName

`func (o *ShowbackRulesListDto) GetShowbackCredentialName() string`

GetShowbackCredentialName returns the ShowbackCredentialName field if non-nil, zero value otherwise.

### GetShowbackCredentialNameOk

`func (o *ShowbackRulesListDto) GetShowbackCredentialNameOk() (*string, bool)`

GetShowbackCredentialNameOk returns a tuple with the ShowbackCredentialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowbackCredentialName

`func (o *ShowbackRulesListDto) SetShowbackCredentialName(v string)`

SetShowbackCredentialName sets ShowbackCredentialName field to given value.

### HasShowbackCredentialName

`func (o *ShowbackRulesListDto) HasShowbackCredentialName() bool`

HasShowbackCredentialName returns a boolean if a field has been set.

### SetShowbackCredentialNameNil

`func (o *ShowbackRulesListDto) SetShowbackCredentialNameNil(b bool)`

 SetShowbackCredentialNameNil sets the value for ShowbackCredentialName to be an explicit nil

### UnsetShowbackCredentialName
`func (o *ShowbackRulesListDto) UnsetShowbackCredentialName()`

UnsetShowbackCredentialName ensures that no value is present for ShowbackCredentialName, not even an explicit nil
### GetShowbackCredentialId

`func (o *ShowbackRulesListDto) GetShowbackCredentialId() int32`

GetShowbackCredentialId returns the ShowbackCredentialId field if non-nil, zero value otherwise.

### GetShowbackCredentialIdOk

`func (o *ShowbackRulesListDto) GetShowbackCredentialIdOk() (*int32, bool)`

GetShowbackCredentialIdOk returns a tuple with the ShowbackCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowbackCredentialId

`func (o *ShowbackRulesListDto) SetShowbackCredentialId(v int32)`

SetShowbackCredentialId sets ShowbackCredentialId field to given value.

### HasShowbackCredentialId

`func (o *ShowbackRulesListDto) HasShowbackCredentialId() bool`

HasShowbackCredentialId returns a boolean if a field has been set.

### SetShowbackCredentialIdNil

`func (o *ShowbackRulesListDto) SetShowbackCredentialIdNil(b bool)`

 SetShowbackCredentialIdNil sets the value for ShowbackCredentialId to be an explicit nil

### UnsetShowbackCredentialId
`func (o *ShowbackRulesListDto) UnsetShowbackCredentialId()`

UnsetShowbackCredentialId ensures that no value is present for ShowbackCredentialId, not even an explicit nil
### GetCreatedBy

`func (o *ShowbackRulesListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ShowbackRulesListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ShowbackRulesListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ShowbackRulesListDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ShowbackRulesListDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ShowbackRulesListDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetLastModified

`func (o *ShowbackRulesListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ShowbackRulesListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ShowbackRulesListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ShowbackRulesListDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### SetLastModifiedNil

`func (o *ShowbackRulesListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *ShowbackRulesListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *ShowbackRulesListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ShowbackRulesListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ShowbackRulesListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *ShowbackRulesListDto) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### SetLastModifiedByNil

`func (o *ShowbackRulesListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *ShowbackRulesListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


