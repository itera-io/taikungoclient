/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the ListForPartnersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListForPartnersDto{}

// ListForPartnersDto struct for ListForPartnersDto
type ListForPartnersDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ProjectLimit *int32 `json:"projectLimit,omitempty"`
	ServerLimit *int32 `json:"serverLimit,omitempty"`
	UserLimit *int32 `json:"userLimit,omitempty"`
	CloudCredentialLimit *int32 `json:"cloudCredentialLimit,omitempty"`
	MonthlyPrice *float64 `json:"monthlyPrice,omitempty"`
	YearlyPrice *float64 `json:"yearlyPrice,omitempty"`
	TcuPrice *float64 `json:"tcuPrice,omitempty"`
	IsDeprecated *bool `json:"isDeprecated,omitempty"`
	Currency *string `json:"currency,omitempty"`
	IsEnterprise *bool `json:"isEnterprise,omitempty"`
	Partner *PartnerDetailsForSubscription `json:"partner,omitempty"`
	ExceededUser *bool `json:"exceededUser,omitempty"`
	ExceededProject *bool `json:"exceededProject,omitempty"`
	ExceededCloudCredential *bool `json:"exceededCloudCredential,omitempty"`
	ExceededServers *bool `json:"exceededServers,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	IsYearly *bool `json:"isYearly,omitempty"`
	TrialDays *int32 `json:"trialDays,omitempty"`
	Description *string `json:"description,omitempty"`
	IsDemo *bool `json:"isDemo,omitempty"`
}

// NewListForPartnersDto instantiates a new ListForPartnersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListForPartnersDto() *ListForPartnersDto {
	this := ListForPartnersDto{}
	return &this
}

// NewListForPartnersDtoWithDefaults instantiates a new ListForPartnersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListForPartnersDtoWithDefaults() *ListForPartnersDto {
	this := ListForPartnersDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListForPartnersDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListForPartnersDto) SetName(v string) {
	o.Name = &v
}

// GetProjectLimit returns the ProjectLimit field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetProjectLimit() int32 {
	if o == nil || IsNil(o.ProjectLimit) {
		var ret int32
		return ret
	}
	return *o.ProjectLimit
}

// GetProjectLimitOk returns a tuple with the ProjectLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetProjectLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectLimit) {
		return nil, false
	}
	return o.ProjectLimit, true
}

// HasProjectLimit returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasProjectLimit() bool {
	if o != nil && !IsNil(o.ProjectLimit) {
		return true
	}

	return false
}

// SetProjectLimit gets a reference to the given int32 and assigns it to the ProjectLimit field.
func (o *ListForPartnersDto) SetProjectLimit(v int32) {
	o.ProjectLimit = &v
}

// GetServerLimit returns the ServerLimit field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetServerLimit() int32 {
	if o == nil || IsNil(o.ServerLimit) {
		var ret int32
		return ret
	}
	return *o.ServerLimit
}

// GetServerLimitOk returns a tuple with the ServerLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetServerLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerLimit) {
		return nil, false
	}
	return o.ServerLimit, true
}

// HasServerLimit returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasServerLimit() bool {
	if o != nil && !IsNil(o.ServerLimit) {
		return true
	}

	return false
}

// SetServerLimit gets a reference to the given int32 and assigns it to the ServerLimit field.
func (o *ListForPartnersDto) SetServerLimit(v int32) {
	o.ServerLimit = &v
}

// GetUserLimit returns the UserLimit field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetUserLimit() int32 {
	if o == nil || IsNil(o.UserLimit) {
		var ret int32
		return ret
	}
	return *o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetUserLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UserLimit) {
		return nil, false
	}
	return o.UserLimit, true
}

// HasUserLimit returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasUserLimit() bool {
	if o != nil && !IsNil(o.UserLimit) {
		return true
	}

	return false
}

// SetUserLimit gets a reference to the given int32 and assigns it to the UserLimit field.
func (o *ListForPartnersDto) SetUserLimit(v int32) {
	o.UserLimit = &v
}

// GetCloudCredentialLimit returns the CloudCredentialLimit field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetCloudCredentialLimit() int32 {
	if o == nil || IsNil(o.CloudCredentialLimit) {
		var ret int32
		return ret
	}
	return *o.CloudCredentialLimit
}

// GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetCloudCredentialLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudCredentialLimit) {
		return nil, false
	}
	return o.CloudCredentialLimit, true
}

// HasCloudCredentialLimit returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasCloudCredentialLimit() bool {
	if o != nil && !IsNil(o.CloudCredentialLimit) {
		return true
	}

	return false
}

// SetCloudCredentialLimit gets a reference to the given int32 and assigns it to the CloudCredentialLimit field.
func (o *ListForPartnersDto) SetCloudCredentialLimit(v int32) {
	o.CloudCredentialLimit = &v
}

// GetMonthlyPrice returns the MonthlyPrice field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetMonthlyPrice() float64 {
	if o == nil || IsNil(o.MonthlyPrice) {
		var ret float64
		return ret
	}
	return *o.MonthlyPrice
}

// GetMonthlyPriceOk returns a tuple with the MonthlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetMonthlyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.MonthlyPrice) {
		return nil, false
	}
	return o.MonthlyPrice, true
}

// HasMonthlyPrice returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasMonthlyPrice() bool {
	if o != nil && !IsNil(o.MonthlyPrice) {
		return true
	}

	return false
}

// SetMonthlyPrice gets a reference to the given float64 and assigns it to the MonthlyPrice field.
func (o *ListForPartnersDto) SetMonthlyPrice(v float64) {
	o.MonthlyPrice = &v
}

// GetYearlyPrice returns the YearlyPrice field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetYearlyPrice() float64 {
	if o == nil || IsNil(o.YearlyPrice) {
		var ret float64
		return ret
	}
	return *o.YearlyPrice
}

// GetYearlyPriceOk returns a tuple with the YearlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetYearlyPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.YearlyPrice) {
		return nil, false
	}
	return o.YearlyPrice, true
}

// HasYearlyPrice returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasYearlyPrice() bool {
	if o != nil && !IsNil(o.YearlyPrice) {
		return true
	}

	return false
}

// SetYearlyPrice gets a reference to the given float64 and assigns it to the YearlyPrice field.
func (o *ListForPartnersDto) SetYearlyPrice(v float64) {
	o.YearlyPrice = &v
}

// GetTcuPrice returns the TcuPrice field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetTcuPrice() float64 {
	if o == nil || IsNil(o.TcuPrice) {
		var ret float64
		return ret
	}
	return *o.TcuPrice
}

// GetTcuPriceOk returns a tuple with the TcuPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetTcuPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.TcuPrice) {
		return nil, false
	}
	return o.TcuPrice, true
}

// HasTcuPrice returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasTcuPrice() bool {
	if o != nil && !IsNil(o.TcuPrice) {
		return true
	}

	return false
}

// SetTcuPrice gets a reference to the given float64 and assigns it to the TcuPrice field.
func (o *ListForPartnersDto) SetTcuPrice(v float64) {
	o.TcuPrice = &v
}

// GetIsDeprecated returns the IsDeprecated field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetIsDeprecated() bool {
	if o == nil || IsNil(o.IsDeprecated) {
		var ret bool
		return ret
	}
	return *o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeprecated) {
		return nil, false
	}
	return o.IsDeprecated, true
}

// HasIsDeprecated returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasIsDeprecated() bool {
	if o != nil && !IsNil(o.IsDeprecated) {
		return true
	}

	return false
}

// SetIsDeprecated gets a reference to the given bool and assigns it to the IsDeprecated field.
func (o *ListForPartnersDto) SetIsDeprecated(v bool) {
	o.IsDeprecated = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ListForPartnersDto) SetCurrency(v string) {
	o.Currency = &v
}

// GetIsEnterprise returns the IsEnterprise field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetIsEnterprise() bool {
	if o == nil || IsNil(o.IsEnterprise) {
		var ret bool
		return ret
	}
	return *o.IsEnterprise
}

// GetIsEnterpriseOk returns a tuple with the IsEnterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsEnterpriseOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnterprise) {
		return nil, false
	}
	return o.IsEnterprise, true
}

// HasIsEnterprise returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasIsEnterprise() bool {
	if o != nil && !IsNil(o.IsEnterprise) {
		return true
	}

	return false
}

// SetIsEnterprise gets a reference to the given bool and assigns it to the IsEnterprise field.
func (o *ListForPartnersDto) SetIsEnterprise(v bool) {
	o.IsEnterprise = &v
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetPartner() PartnerDetailsForSubscription {
	if o == nil || IsNil(o.Partner) {
		var ret PartnerDetailsForSubscription
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetPartnerOk() (*PartnerDetailsForSubscription, bool) {
	if o == nil || IsNil(o.Partner) {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasPartner() bool {
	if o != nil && !IsNil(o.Partner) {
		return true
	}

	return false
}

// SetPartner gets a reference to the given PartnerDetailsForSubscription and assigns it to the Partner field.
func (o *ListForPartnersDto) SetPartner(v PartnerDetailsForSubscription) {
	o.Partner = &v
}

// GetExceededUser returns the ExceededUser field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetExceededUser() bool {
	if o == nil || IsNil(o.ExceededUser) {
		var ret bool
		return ret
	}
	return *o.ExceededUser
}

// GetExceededUserOk returns a tuple with the ExceededUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededUserOk() (*bool, bool) {
	if o == nil || IsNil(o.ExceededUser) {
		return nil, false
	}
	return o.ExceededUser, true
}

// HasExceededUser returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasExceededUser() bool {
	if o != nil && !IsNil(o.ExceededUser) {
		return true
	}

	return false
}

// SetExceededUser gets a reference to the given bool and assigns it to the ExceededUser field.
func (o *ListForPartnersDto) SetExceededUser(v bool) {
	o.ExceededUser = &v
}

// GetExceededProject returns the ExceededProject field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetExceededProject() bool {
	if o == nil || IsNil(o.ExceededProject) {
		var ret bool
		return ret
	}
	return *o.ExceededProject
}

// GetExceededProjectOk returns a tuple with the ExceededProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededProjectOk() (*bool, bool) {
	if o == nil || IsNil(o.ExceededProject) {
		return nil, false
	}
	return o.ExceededProject, true
}

// HasExceededProject returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasExceededProject() bool {
	if o != nil && !IsNil(o.ExceededProject) {
		return true
	}

	return false
}

// SetExceededProject gets a reference to the given bool and assigns it to the ExceededProject field.
func (o *ListForPartnersDto) SetExceededProject(v bool) {
	o.ExceededProject = &v
}

// GetExceededCloudCredential returns the ExceededCloudCredential field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetExceededCloudCredential() bool {
	if o == nil || IsNil(o.ExceededCloudCredential) {
		var ret bool
		return ret
	}
	return *o.ExceededCloudCredential
}

// GetExceededCloudCredentialOk returns a tuple with the ExceededCloudCredential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededCloudCredentialOk() (*bool, bool) {
	if o == nil || IsNil(o.ExceededCloudCredential) {
		return nil, false
	}
	return o.ExceededCloudCredential, true
}

// HasExceededCloudCredential returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasExceededCloudCredential() bool {
	if o != nil && !IsNil(o.ExceededCloudCredential) {
		return true
	}

	return false
}

// SetExceededCloudCredential gets a reference to the given bool and assigns it to the ExceededCloudCredential field.
func (o *ListForPartnersDto) SetExceededCloudCredential(v bool) {
	o.ExceededCloudCredential = &v
}

// GetExceededServers returns the ExceededServers field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetExceededServers() bool {
	if o == nil || IsNil(o.ExceededServers) {
		var ret bool
		return ret
	}
	return *o.ExceededServers
}

// GetExceededServersOk returns a tuple with the ExceededServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededServersOk() (*bool, bool) {
	if o == nil || IsNil(o.ExceededServers) {
		return nil, false
	}
	return o.ExceededServers, true
}

// HasExceededServers returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasExceededServers() bool {
	if o != nil && !IsNil(o.ExceededServers) {
		return true
	}

	return false
}

// SetExceededServers gets a reference to the given bool and assigns it to the ExceededServers field.
func (o *ListForPartnersDto) SetExceededServers(v bool) {
	o.ExceededServers = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ListForPartnersDto) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsYearly returns the IsYearly field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetIsYearly() bool {
	if o == nil || IsNil(o.IsYearly) {
		var ret bool
		return ret
	}
	return *o.IsYearly
}

// GetIsYearlyOk returns a tuple with the IsYearly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsYearlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsYearly) {
		return nil, false
	}
	return o.IsYearly, true
}

// HasIsYearly returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasIsYearly() bool {
	if o != nil && !IsNil(o.IsYearly) {
		return true
	}

	return false
}

// SetIsYearly gets a reference to the given bool and assigns it to the IsYearly field.
func (o *ListForPartnersDto) SetIsYearly(v bool) {
	o.IsYearly = &v
}

// GetTrialDays returns the TrialDays field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetTrialDays() int32 {
	if o == nil || IsNil(o.TrialDays) {
		var ret int32
		return ret
	}
	return *o.TrialDays
}

// GetTrialDaysOk returns a tuple with the TrialDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetTrialDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialDays) {
		return nil, false
	}
	return o.TrialDays, true
}

// HasTrialDays returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasTrialDays() bool {
	if o != nil && !IsNil(o.TrialDays) {
		return true
	}

	return false
}

// SetTrialDays gets a reference to the given int32 and assigns it to the TrialDays field.
func (o *ListForPartnersDto) SetTrialDays(v int32) {
	o.TrialDays = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListForPartnersDto) SetDescription(v string) {
	o.Description = &v
}

// GetIsDemo returns the IsDemo field value if set, zero value otherwise.
func (o *ListForPartnersDto) GetIsDemo() bool {
	if o == nil || IsNil(o.IsDemo) {
		var ret bool
		return ret
	}
	return *o.IsDemo
}

// GetIsDemoOk returns a tuple with the IsDemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsDemoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDemo) {
		return nil, false
	}
	return o.IsDemo, true
}

// HasIsDemo returns a boolean if a field has been set.
func (o *ListForPartnersDto) HasIsDemo() bool {
	if o != nil && !IsNil(o.IsDemo) {
		return true
	}

	return false
}

// SetIsDemo gets a reference to the given bool and assigns it to the IsDemo field.
func (o *ListForPartnersDto) SetIsDemo(v bool) {
	o.IsDemo = &v
}

func (o ListForPartnersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListForPartnersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProjectLimit) {
		toSerialize["projectLimit"] = o.ProjectLimit
	}
	if !IsNil(o.ServerLimit) {
		toSerialize["serverLimit"] = o.ServerLimit
	}
	if !IsNil(o.UserLimit) {
		toSerialize["userLimit"] = o.UserLimit
	}
	if !IsNil(o.CloudCredentialLimit) {
		toSerialize["cloudCredentialLimit"] = o.CloudCredentialLimit
	}
	if !IsNil(o.MonthlyPrice) {
		toSerialize["monthlyPrice"] = o.MonthlyPrice
	}
	if !IsNil(o.YearlyPrice) {
		toSerialize["yearlyPrice"] = o.YearlyPrice
	}
	if !IsNil(o.TcuPrice) {
		toSerialize["tcuPrice"] = o.TcuPrice
	}
	if !IsNil(o.IsDeprecated) {
		toSerialize["isDeprecated"] = o.IsDeprecated
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.IsEnterprise) {
		toSerialize["isEnterprise"] = o.IsEnterprise
	}
	if !IsNil(o.Partner) {
		toSerialize["partner"] = o.Partner
	}
	if !IsNil(o.ExceededUser) {
		toSerialize["exceededUser"] = o.ExceededUser
	}
	if !IsNil(o.ExceededProject) {
		toSerialize["exceededProject"] = o.ExceededProject
	}
	if !IsNil(o.ExceededCloudCredential) {
		toSerialize["exceededCloudCredential"] = o.ExceededCloudCredential
	}
	if !IsNil(o.ExceededServers) {
		toSerialize["exceededServers"] = o.ExceededServers
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.IsYearly) {
		toSerialize["isYearly"] = o.IsYearly
	}
	if !IsNil(o.TrialDays) {
		toSerialize["trialDays"] = o.TrialDays
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsDemo) {
		toSerialize["isDemo"] = o.IsDemo
	}
	return toSerialize, nil
}

type NullableListForPartnersDto struct {
	value *ListForPartnersDto
	isSet bool
}

func (v NullableListForPartnersDto) Get() *ListForPartnersDto {
	return v.value
}

func (v *NullableListForPartnersDto) Set(val *ListForPartnersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableListForPartnersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableListForPartnersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListForPartnersDto(val *ListForPartnersDto) *NullableListForPartnersDto {
	return &NullableListForPartnersDto{value: val, isSet: true}
}

func (v NullableListForPartnersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListForPartnersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


