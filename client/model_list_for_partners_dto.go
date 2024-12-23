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
	"bytes"
	"fmt"
)

// checks if the ListForPartnersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListForPartnersDto{}

// ListForPartnersDto struct for ListForPartnersDto
type ListForPartnersDto struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	ProjectLimit int32 `json:"projectLimit"`
	ServerLimit int32 `json:"serverLimit"`
	UserLimit int32 `json:"userLimit"`
	CloudCredentialLimit int32 `json:"cloudCredentialLimit"`
	MonthlyPrice float64 `json:"monthlyPrice"`
	YearlyPrice float64 `json:"yearlyPrice"`
	TcuPrice float64 `json:"tcuPrice"`
	IsDeprecated bool `json:"isDeprecated"`
	Currency NullableString `json:"currency"`
	IsEnterprise bool `json:"isEnterprise"`
	Partner PartnerDetailsForSubscription `json:"partner"`
	ExceededUser bool `json:"exceededUser"`
	ExceededProject bool `json:"exceededProject"`
	ExceededCloudCredential bool `json:"exceededCloudCredential"`
	ExceededServers bool `json:"exceededServers"`
	IsActive bool `json:"isActive"`
	IsYearly bool `json:"isYearly"`
	TrialDays int32 `json:"trialDays"`
	Description NullableString `json:"description"`
	IsDemo bool `json:"isDemo"`
}

type _ListForPartnersDto ListForPartnersDto

// NewListForPartnersDto instantiates a new ListForPartnersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListForPartnersDto(id int32, name NullableString, projectLimit int32, serverLimit int32, userLimit int32, cloudCredentialLimit int32, monthlyPrice float64, yearlyPrice float64, tcuPrice float64, isDeprecated bool, currency NullableString, isEnterprise bool, partner PartnerDetailsForSubscription, exceededUser bool, exceededProject bool, exceededCloudCredential bool, exceededServers bool, isActive bool, isYearly bool, trialDays int32, description NullableString, isDemo bool) *ListForPartnersDto {
	this := ListForPartnersDto{}
	this.Id = id
	this.Name = name
	this.ProjectLimit = projectLimit
	this.ServerLimit = serverLimit
	this.UserLimit = userLimit
	this.CloudCredentialLimit = cloudCredentialLimit
	this.MonthlyPrice = monthlyPrice
	this.YearlyPrice = yearlyPrice
	this.TcuPrice = tcuPrice
	this.IsDeprecated = isDeprecated
	this.Currency = currency
	this.IsEnterprise = isEnterprise
	this.Partner = partner
	this.ExceededUser = exceededUser
	this.ExceededProject = exceededProject
	this.ExceededCloudCredential = exceededCloudCredential
	this.ExceededServers = exceededServers
	this.IsActive = isActive
	this.IsYearly = isYearly
	this.TrialDays = trialDays
	this.Description = description
	this.IsDemo = isDemo
	return &this
}

// NewListForPartnersDtoWithDefaults instantiates a new ListForPartnersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListForPartnersDtoWithDefaults() *ListForPartnersDto {
	this := ListForPartnersDto{}
	return &this
}

// GetId returns the Id field value
func (o *ListForPartnersDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListForPartnersDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListForPartnersDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListForPartnersDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ListForPartnersDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetProjectLimit returns the ProjectLimit field value
func (o *ListForPartnersDto) GetProjectLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectLimit
}

// GetProjectLimitOk returns a tuple with the ProjectLimit field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetProjectLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectLimit, true
}

// SetProjectLimit sets field value
func (o *ListForPartnersDto) SetProjectLimit(v int32) {
	o.ProjectLimit = v
}

// GetServerLimit returns the ServerLimit field value
func (o *ListForPartnersDto) GetServerLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerLimit
}

// GetServerLimitOk returns a tuple with the ServerLimit field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetServerLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerLimit, true
}

// SetServerLimit sets field value
func (o *ListForPartnersDto) SetServerLimit(v int32) {
	o.ServerLimit = v
}

// GetUserLimit returns the UserLimit field value
func (o *ListForPartnersDto) GetUserLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserLimit
}

// GetUserLimitOk returns a tuple with the UserLimit field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetUserLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserLimit, true
}

// SetUserLimit sets field value
func (o *ListForPartnersDto) SetUserLimit(v int32) {
	o.UserLimit = v
}

// GetCloudCredentialLimit returns the CloudCredentialLimit field value
func (o *ListForPartnersDto) GetCloudCredentialLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudCredentialLimit
}

// GetCloudCredentialLimitOk returns a tuple with the CloudCredentialLimit field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetCloudCredentialLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudCredentialLimit, true
}

// SetCloudCredentialLimit sets field value
func (o *ListForPartnersDto) SetCloudCredentialLimit(v int32) {
	o.CloudCredentialLimit = v
}

// GetMonthlyPrice returns the MonthlyPrice field value
func (o *ListForPartnersDto) GetMonthlyPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MonthlyPrice
}

// GetMonthlyPriceOk returns a tuple with the MonthlyPrice field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetMonthlyPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyPrice, true
}

// SetMonthlyPrice sets field value
func (o *ListForPartnersDto) SetMonthlyPrice(v float64) {
	o.MonthlyPrice = v
}

// GetYearlyPrice returns the YearlyPrice field value
func (o *ListForPartnersDto) GetYearlyPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.YearlyPrice
}

// GetYearlyPriceOk returns a tuple with the YearlyPrice field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetYearlyPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.YearlyPrice, true
}

// SetYearlyPrice sets field value
func (o *ListForPartnersDto) SetYearlyPrice(v float64) {
	o.YearlyPrice = v
}

// GetTcuPrice returns the TcuPrice field value
func (o *ListForPartnersDto) GetTcuPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TcuPrice
}

// GetTcuPriceOk returns a tuple with the TcuPrice field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetTcuPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TcuPrice, true
}

// SetTcuPrice sets field value
func (o *ListForPartnersDto) SetTcuPrice(v float64) {
	o.TcuPrice = v
}

// GetIsDeprecated returns the IsDeprecated field value
func (o *ListForPartnersDto) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value
func (o *ListForPartnersDto) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListForPartnersDto) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListForPartnersDto) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *ListForPartnersDto) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetIsEnterprise returns the IsEnterprise field value
func (o *ListForPartnersDto) GetIsEnterprise() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnterprise
}

// GetIsEnterpriseOk returns a tuple with the IsEnterprise field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsEnterpriseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnterprise, true
}

// SetIsEnterprise sets field value
func (o *ListForPartnersDto) SetIsEnterprise(v bool) {
	o.IsEnterprise = v
}

// GetPartner returns the Partner field value
func (o *ListForPartnersDto) GetPartner() PartnerDetailsForSubscription {
	if o == nil {
		var ret PartnerDetailsForSubscription
		return ret
	}

	return o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetPartnerOk() (*PartnerDetailsForSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partner, true
}

// SetPartner sets field value
func (o *ListForPartnersDto) SetPartner(v PartnerDetailsForSubscription) {
	o.Partner = v
}

// GetExceededUser returns the ExceededUser field value
func (o *ListForPartnersDto) GetExceededUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExceededUser
}

// GetExceededUserOk returns a tuple with the ExceededUser field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceededUser, true
}

// SetExceededUser sets field value
func (o *ListForPartnersDto) SetExceededUser(v bool) {
	o.ExceededUser = v
}

// GetExceededProject returns the ExceededProject field value
func (o *ListForPartnersDto) GetExceededProject() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExceededProject
}

// GetExceededProjectOk returns a tuple with the ExceededProject field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededProjectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceededProject, true
}

// SetExceededProject sets field value
func (o *ListForPartnersDto) SetExceededProject(v bool) {
	o.ExceededProject = v
}

// GetExceededCloudCredential returns the ExceededCloudCredential field value
func (o *ListForPartnersDto) GetExceededCloudCredential() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExceededCloudCredential
}

// GetExceededCloudCredentialOk returns a tuple with the ExceededCloudCredential field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededCloudCredentialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceededCloudCredential, true
}

// SetExceededCloudCredential sets field value
func (o *ListForPartnersDto) SetExceededCloudCredential(v bool) {
	o.ExceededCloudCredential = v
}

// GetExceededServers returns the ExceededServers field value
func (o *ListForPartnersDto) GetExceededServers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExceededServers
}

// GetExceededServersOk returns a tuple with the ExceededServers field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetExceededServersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceededServers, true
}

// SetExceededServers sets field value
func (o *ListForPartnersDto) SetExceededServers(v bool) {
	o.ExceededServers = v
}

// GetIsActive returns the IsActive field value
func (o *ListForPartnersDto) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *ListForPartnersDto) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsYearly returns the IsYearly field value
func (o *ListForPartnersDto) GetIsYearly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsYearly
}

// GetIsYearlyOk returns a tuple with the IsYearly field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsYearlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsYearly, true
}

// SetIsYearly sets field value
func (o *ListForPartnersDto) SetIsYearly(v bool) {
	o.IsYearly = v
}

// GetTrialDays returns the TrialDays field value
func (o *ListForPartnersDto) GetTrialDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrialDays
}

// GetTrialDaysOk returns a tuple with the TrialDays field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetTrialDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrialDays, true
}

// SetTrialDays sets field value
func (o *ListForPartnersDto) SetTrialDays(v int32) {
	o.TrialDays = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ListForPartnersDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListForPartnersDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ListForPartnersDto) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetIsDemo returns the IsDemo field value
func (o *ListForPartnersDto) GetIsDemo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDemo
}

// GetIsDemoOk returns a tuple with the IsDemo field value
// and a boolean to check if the value has been set.
func (o *ListForPartnersDto) GetIsDemoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDemo, true
}

// SetIsDemo sets field value
func (o *ListForPartnersDto) SetIsDemo(v bool) {
	o.IsDemo = v
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
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["projectLimit"] = o.ProjectLimit
	toSerialize["serverLimit"] = o.ServerLimit
	toSerialize["userLimit"] = o.UserLimit
	toSerialize["cloudCredentialLimit"] = o.CloudCredentialLimit
	toSerialize["monthlyPrice"] = o.MonthlyPrice
	toSerialize["yearlyPrice"] = o.YearlyPrice
	toSerialize["tcuPrice"] = o.TcuPrice
	toSerialize["isDeprecated"] = o.IsDeprecated
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["isEnterprise"] = o.IsEnterprise
	toSerialize["partner"] = o.Partner
	toSerialize["exceededUser"] = o.ExceededUser
	toSerialize["exceededProject"] = o.ExceededProject
	toSerialize["exceededCloudCredential"] = o.ExceededCloudCredential
	toSerialize["exceededServers"] = o.ExceededServers
	toSerialize["isActive"] = o.IsActive
	toSerialize["isYearly"] = o.IsYearly
	toSerialize["trialDays"] = o.TrialDays
	toSerialize["description"] = o.Description.Get()
	toSerialize["isDemo"] = o.IsDemo
	return toSerialize, nil
}

func (o *ListForPartnersDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"projectLimit",
		"serverLimit",
		"userLimit",
		"cloudCredentialLimit",
		"monthlyPrice",
		"yearlyPrice",
		"tcuPrice",
		"isDeprecated",
		"currency",
		"isEnterprise",
		"partner",
		"exceededUser",
		"exceededProject",
		"exceededCloudCredential",
		"exceededServers",
		"isActive",
		"isYearly",
		"trialDays",
		"description",
		"isDemo",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListForPartnersDto := _ListForPartnersDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListForPartnersDto)

	if err != nil {
		return err
	}

	*o = ListForPartnersDto(varListForPartnersDto)

	return err
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


