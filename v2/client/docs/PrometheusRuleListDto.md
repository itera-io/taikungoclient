# PrometheusRuleListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Password** | **string** |  | 
**UserName** | **string** |  | 
**Url** | **string** |  | 
**MetricName** | **NullableString** |  | 
**Labels** | [**[]PrometheusLabelListDto**](PrometheusLabelListDto.md) |  | 
**BoundOrganizations** | [**[]PrometheusOrganizationDiscountDto**](PrometheusOrganizationDiscountDto.md) |  | 
**Type** | **map[string]interface{}** |  | 
**Price** | **float64** |  | 
**BillingStartDate** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**Domain** | [**DomainEntity**](DomainEntity.md) |  | 
**OperationCredential** | [**OperationCredentialsForOrganizationEntity**](OperationCredentialsForOrganizationEntity.md) |  | 
**CreatedBy** | **string** |  | 
**LastModified** | **NullableString** |  | 
**LastModifiedBy** | **NullableString** |  | 

## Methods

### NewPrometheusRuleListDto

`func NewPrometheusRuleListDto(id int32, name string, password string, userName string, url string, metricName NullableString, labels []PrometheusLabelListDto, boundOrganizations []PrometheusOrganizationDiscountDto, type_ map[string]interface{}, price float64, billingStartDate NullableString, createdAt NullableString, domain DomainEntity, operationCredential OperationCredentialsForOrganizationEntity, createdBy string, lastModified NullableString, lastModifiedBy NullableString, ) *PrometheusRuleListDto`

NewPrometheusRuleListDto instantiates a new PrometheusRuleListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrometheusRuleListDtoWithDefaults

`func NewPrometheusRuleListDtoWithDefaults() *PrometheusRuleListDto`

NewPrometheusRuleListDtoWithDefaults instantiates a new PrometheusRuleListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrometheusRuleListDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrometheusRuleListDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrometheusRuleListDto) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *PrometheusRuleListDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrometheusRuleListDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrometheusRuleListDto) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *PrometheusRuleListDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PrometheusRuleListDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PrometheusRuleListDto) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUserName

`func (o *PrometheusRuleListDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PrometheusRuleListDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PrometheusRuleListDto) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetUrl

`func (o *PrometheusRuleListDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PrometheusRuleListDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PrometheusRuleListDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetricName

`func (o *PrometheusRuleListDto) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *PrometheusRuleListDto) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *PrometheusRuleListDto) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### SetMetricNameNil

`func (o *PrometheusRuleListDto) SetMetricNameNil(b bool)`

 SetMetricNameNil sets the value for MetricName to be an explicit nil

### UnsetMetricName
`func (o *PrometheusRuleListDto) UnsetMetricName()`

UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
### GetLabels

`func (o *PrometheusRuleListDto) GetLabels() []PrometheusLabelListDto`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PrometheusRuleListDto) GetLabelsOk() (*[]PrometheusLabelListDto, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PrometheusRuleListDto) SetLabels(v []PrometheusLabelListDto)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *PrometheusRuleListDto) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *PrometheusRuleListDto) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetBoundOrganizations

`func (o *PrometheusRuleListDto) GetBoundOrganizations() []PrometheusOrganizationDiscountDto`

GetBoundOrganizations returns the BoundOrganizations field if non-nil, zero value otherwise.

### GetBoundOrganizationsOk

`func (o *PrometheusRuleListDto) GetBoundOrganizationsOk() (*[]PrometheusOrganizationDiscountDto, bool)`

GetBoundOrganizationsOk returns a tuple with the BoundOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundOrganizations

`func (o *PrometheusRuleListDto) SetBoundOrganizations(v []PrometheusOrganizationDiscountDto)`

SetBoundOrganizations sets BoundOrganizations field to given value.


### SetBoundOrganizationsNil

`func (o *PrometheusRuleListDto) SetBoundOrganizationsNil(b bool)`

 SetBoundOrganizationsNil sets the value for BoundOrganizations to be an explicit nil

### UnsetBoundOrganizations
`func (o *PrometheusRuleListDto) UnsetBoundOrganizations()`

UnsetBoundOrganizations ensures that no value is present for BoundOrganizations, not even an explicit nil
### GetType

`func (o *PrometheusRuleListDto) GetType() map[string]interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrometheusRuleListDto) GetTypeOk() (*map[string]interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrometheusRuleListDto) SetType(v map[string]interface{})`

SetType sets Type field to given value.


### GetPrice

`func (o *PrometheusRuleListDto) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PrometheusRuleListDto) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PrometheusRuleListDto) SetPrice(v float64)`

SetPrice sets Price field to given value.


### GetBillingStartDate

`func (o *PrometheusRuleListDto) GetBillingStartDate() string`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *PrometheusRuleListDto) GetBillingStartDateOk() (*string, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *PrometheusRuleListDto) SetBillingStartDate(v string)`

SetBillingStartDate sets BillingStartDate field to given value.


### SetBillingStartDateNil

`func (o *PrometheusRuleListDto) SetBillingStartDateNil(b bool)`

 SetBillingStartDateNil sets the value for BillingStartDate to be an explicit nil

### UnsetBillingStartDate
`func (o *PrometheusRuleListDto) UnsetBillingStartDate()`

UnsetBillingStartDate ensures that no value is present for BillingStartDate, not even an explicit nil
### GetCreatedAt

`func (o *PrometheusRuleListDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrometheusRuleListDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrometheusRuleListDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PrometheusRuleListDto) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PrometheusRuleListDto) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDomain

`func (o *PrometheusRuleListDto) GetDomain() DomainEntity`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PrometheusRuleListDto) GetDomainOk() (*DomainEntity, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PrometheusRuleListDto) SetDomain(v DomainEntity)`

SetDomain sets Domain field to given value.


### GetOperationCredential

`func (o *PrometheusRuleListDto) GetOperationCredential() OperationCredentialsForOrganizationEntity`

GetOperationCredential returns the OperationCredential field if non-nil, zero value otherwise.

### GetOperationCredentialOk

`func (o *PrometheusRuleListDto) GetOperationCredentialOk() (*OperationCredentialsForOrganizationEntity, bool)`

GetOperationCredentialOk returns a tuple with the OperationCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCredential

`func (o *PrometheusRuleListDto) SetOperationCredential(v OperationCredentialsForOrganizationEntity)`

SetOperationCredential sets OperationCredential field to given value.


### GetCreatedBy

`func (o *PrometheusRuleListDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrometheusRuleListDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrometheusRuleListDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModified

`func (o *PrometheusRuleListDto) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *PrometheusRuleListDto) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *PrometheusRuleListDto) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.


### SetLastModifiedNil

`func (o *PrometheusRuleListDto) SetLastModifiedNil(b bool)`

 SetLastModifiedNil sets the value for LastModified to be an explicit nil

### UnsetLastModified
`func (o *PrometheusRuleListDto) UnsetLastModified()`

UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
### GetLastModifiedBy

`func (o *PrometheusRuleListDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *PrometheusRuleListDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *PrometheusRuleListDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### SetLastModifiedByNil

`func (o *PrometheusRuleListDto) SetLastModifiedByNil(b bool)`

 SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil

### UnsetLastModifiedBy
`func (o *PrometheusRuleListDto) UnsetLastModifiedBy()`

UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


