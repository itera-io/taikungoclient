# TransitEncryptionStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransitEnabled** | **bool** |  | 
**TenantSecretsEncryptedAtCallSites** | **bool** |  | 
**WrappedDekRowCount** | **int32** |  | 
**Columns** | [**[]EncryptedColumnStatusDto**](EncryptedColumnStatusDto.md) |  | 
**PartnerPoolDekAnomalies** | [**[]PartnerPoolDekAnomalyDto**](PartnerPoolDekAnomalyDto.md) |  | 
**LastTransitFailureUtc** | **NullableTime** |  | 

## Methods

### NewTransitEncryptionStatusDto

`func NewTransitEncryptionStatusDto(transitEnabled bool, tenantSecretsEncryptedAtCallSites bool, wrappedDekRowCount int32, columns []EncryptedColumnStatusDto, partnerPoolDekAnomalies []PartnerPoolDekAnomalyDto, lastTransitFailureUtc NullableTime, ) *TransitEncryptionStatusDto`

NewTransitEncryptionStatusDto instantiates a new TransitEncryptionStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitEncryptionStatusDtoWithDefaults

`func NewTransitEncryptionStatusDtoWithDefaults() *TransitEncryptionStatusDto`

NewTransitEncryptionStatusDtoWithDefaults instantiates a new TransitEncryptionStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransitEnabled

`func (o *TransitEncryptionStatusDto) GetTransitEnabled() bool`

GetTransitEnabled returns the TransitEnabled field if non-nil, zero value otherwise.

### GetTransitEnabledOk

`func (o *TransitEncryptionStatusDto) GetTransitEnabledOk() (*bool, bool)`

GetTransitEnabledOk returns a tuple with the TransitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitEnabled

`func (o *TransitEncryptionStatusDto) SetTransitEnabled(v bool)`

SetTransitEnabled sets TransitEnabled field to given value.


### GetTenantSecretsEncryptedAtCallSites

`func (o *TransitEncryptionStatusDto) GetTenantSecretsEncryptedAtCallSites() bool`

GetTenantSecretsEncryptedAtCallSites returns the TenantSecretsEncryptedAtCallSites field if non-nil, zero value otherwise.

### GetTenantSecretsEncryptedAtCallSitesOk

`func (o *TransitEncryptionStatusDto) GetTenantSecretsEncryptedAtCallSitesOk() (*bool, bool)`

GetTenantSecretsEncryptedAtCallSitesOk returns a tuple with the TenantSecretsEncryptedAtCallSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantSecretsEncryptedAtCallSites

`func (o *TransitEncryptionStatusDto) SetTenantSecretsEncryptedAtCallSites(v bool)`

SetTenantSecretsEncryptedAtCallSites sets TenantSecretsEncryptedAtCallSites field to given value.


### GetWrappedDekRowCount

`func (o *TransitEncryptionStatusDto) GetWrappedDekRowCount() int32`

GetWrappedDekRowCount returns the WrappedDekRowCount field if non-nil, zero value otherwise.

### GetWrappedDekRowCountOk

`func (o *TransitEncryptionStatusDto) GetWrappedDekRowCountOk() (*int32, bool)`

GetWrappedDekRowCountOk returns a tuple with the WrappedDekRowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrappedDekRowCount

`func (o *TransitEncryptionStatusDto) SetWrappedDekRowCount(v int32)`

SetWrappedDekRowCount sets WrappedDekRowCount field to given value.


### GetColumns

`func (o *TransitEncryptionStatusDto) GetColumns() []EncryptedColumnStatusDto`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TransitEncryptionStatusDto) GetColumnsOk() (*[]EncryptedColumnStatusDto, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TransitEncryptionStatusDto) SetColumns(v []EncryptedColumnStatusDto)`

SetColumns sets Columns field to given value.


### SetColumnsNil

`func (o *TransitEncryptionStatusDto) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *TransitEncryptionStatusDto) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetPartnerPoolDekAnomalies

`func (o *TransitEncryptionStatusDto) GetPartnerPoolDekAnomalies() []PartnerPoolDekAnomalyDto`

GetPartnerPoolDekAnomalies returns the PartnerPoolDekAnomalies field if non-nil, zero value otherwise.

### GetPartnerPoolDekAnomaliesOk

`func (o *TransitEncryptionStatusDto) GetPartnerPoolDekAnomaliesOk() (*[]PartnerPoolDekAnomalyDto, bool)`

GetPartnerPoolDekAnomaliesOk returns a tuple with the PartnerPoolDekAnomalies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPoolDekAnomalies

`func (o *TransitEncryptionStatusDto) SetPartnerPoolDekAnomalies(v []PartnerPoolDekAnomalyDto)`

SetPartnerPoolDekAnomalies sets PartnerPoolDekAnomalies field to given value.


### SetPartnerPoolDekAnomaliesNil

`func (o *TransitEncryptionStatusDto) SetPartnerPoolDekAnomaliesNil(b bool)`

 SetPartnerPoolDekAnomaliesNil sets the value for PartnerPoolDekAnomalies to be an explicit nil

### UnsetPartnerPoolDekAnomalies
`func (o *TransitEncryptionStatusDto) UnsetPartnerPoolDekAnomalies()`

UnsetPartnerPoolDekAnomalies ensures that no value is present for PartnerPoolDekAnomalies, not even an explicit nil
### GetLastTransitFailureUtc

`func (o *TransitEncryptionStatusDto) GetLastTransitFailureUtc() time.Time`

GetLastTransitFailureUtc returns the LastTransitFailureUtc field if non-nil, zero value otherwise.

### GetLastTransitFailureUtcOk

`func (o *TransitEncryptionStatusDto) GetLastTransitFailureUtcOk() (*time.Time, bool)`

GetLastTransitFailureUtcOk returns a tuple with the LastTransitFailureUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitFailureUtc

`func (o *TransitEncryptionStatusDto) SetLastTransitFailureUtc(v time.Time)`

SetLastTransitFailureUtc sets LastTransitFailureUtc field to given value.


### SetLastTransitFailureUtcNil

`func (o *TransitEncryptionStatusDto) SetLastTransitFailureUtcNil(b bool)`

 SetLastTransitFailureUtcNil sets the value for LastTransitFailureUtc to be an explicit nil

### UnsetLastTransitFailureUtc
`func (o *TransitEncryptionStatusDto) UnsetLastTransitFailureUtc()`

UnsetLastTransitFailureUtc ensures that no value is present for LastTransitFailureUtc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


