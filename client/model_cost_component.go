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

// checks if the CostComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CostComponent{}

// CostComponent struct for CostComponent
type CostComponent struct {
	Name NullableString `json:"name,omitempty"`
	Unit NullableString `json:"unit,omitempty"`
	HourlyQuantity NullableString `json:"hourlyQuantity,omitempty"`
	MonthlyQuantity NullableString `json:"monthlyQuantity,omitempty"`
	Price NullableString `json:"price,omitempty"`
	HourlyCost NullableString `json:"hourlyCost,omitempty"`
	MonthlyCost NullableString `json:"monthlyCost,omitempty"`
	PriceNotFound *bool `json:"priceNotFound,omitempty"`
	UsageBased NullableBool `json:"usageBased,omitempty"`
}

// NewCostComponent instantiates a new CostComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostComponent() *CostComponent {
	this := CostComponent{}
	return &this
}

// NewCostComponentWithDefaults instantiates a new CostComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostComponentWithDefaults() *CostComponent {
	this := CostComponent{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CostComponent) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CostComponent) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CostComponent) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CostComponent) UnsetName() {
	o.Name.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetUnit() string {
	if o == nil || IsNil(o.Unit.Get()) {
		var ret string
		return ret
	}
	return *o.Unit.Get()
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unit.Get(), o.Unit.IsSet()
}

// HasUnit returns a boolean if a field has been set.
func (o *CostComponent) HasUnit() bool {
	if o != nil && o.Unit.IsSet() {
		return true
	}

	return false
}

// SetUnit gets a reference to the given NullableString and assigns it to the Unit field.
func (o *CostComponent) SetUnit(v string) {
	o.Unit.Set(&v)
}
// SetUnitNil sets the value for Unit to be an explicit nil
func (o *CostComponent) SetUnitNil() {
	o.Unit.Set(nil)
}

// UnsetUnit ensures that no value is present for Unit, not even an explicit nil
func (o *CostComponent) UnsetUnit() {
	o.Unit.Unset()
}

// GetHourlyQuantity returns the HourlyQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetHourlyQuantity() string {
	if o == nil || IsNil(o.HourlyQuantity.Get()) {
		var ret string
		return ret
	}
	return *o.HourlyQuantity.Get()
}

// GetHourlyQuantityOk returns a tuple with the HourlyQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetHourlyQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HourlyQuantity.Get(), o.HourlyQuantity.IsSet()
}

// HasHourlyQuantity returns a boolean if a field has been set.
func (o *CostComponent) HasHourlyQuantity() bool {
	if o != nil && o.HourlyQuantity.IsSet() {
		return true
	}

	return false
}

// SetHourlyQuantity gets a reference to the given NullableString and assigns it to the HourlyQuantity field.
func (o *CostComponent) SetHourlyQuantity(v string) {
	o.HourlyQuantity.Set(&v)
}
// SetHourlyQuantityNil sets the value for HourlyQuantity to be an explicit nil
func (o *CostComponent) SetHourlyQuantityNil() {
	o.HourlyQuantity.Set(nil)
}

// UnsetHourlyQuantity ensures that no value is present for HourlyQuantity, not even an explicit nil
func (o *CostComponent) UnsetHourlyQuantity() {
	o.HourlyQuantity.Unset()
}

// GetMonthlyQuantity returns the MonthlyQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetMonthlyQuantity() string {
	if o == nil || IsNil(o.MonthlyQuantity.Get()) {
		var ret string
		return ret
	}
	return *o.MonthlyQuantity.Get()
}

// GetMonthlyQuantityOk returns a tuple with the MonthlyQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetMonthlyQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyQuantity.Get(), o.MonthlyQuantity.IsSet()
}

// HasMonthlyQuantity returns a boolean if a field has been set.
func (o *CostComponent) HasMonthlyQuantity() bool {
	if o != nil && o.MonthlyQuantity.IsSet() {
		return true
	}

	return false
}

// SetMonthlyQuantity gets a reference to the given NullableString and assigns it to the MonthlyQuantity field.
func (o *CostComponent) SetMonthlyQuantity(v string) {
	o.MonthlyQuantity.Set(&v)
}
// SetMonthlyQuantityNil sets the value for MonthlyQuantity to be an explicit nil
func (o *CostComponent) SetMonthlyQuantityNil() {
	o.MonthlyQuantity.Set(nil)
}

// UnsetMonthlyQuantity ensures that no value is present for MonthlyQuantity, not even an explicit nil
func (o *CostComponent) UnsetMonthlyQuantity() {
	o.MonthlyQuantity.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetPrice() string {
	if o == nil || IsNil(o.Price.Get()) {
		var ret string
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *CostComponent) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableString and assigns it to the Price field.
func (o *CostComponent) SetPrice(v string) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *CostComponent) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *CostComponent) UnsetPrice() {
	o.Price.Unset()
}

// GetHourlyCost returns the HourlyCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetHourlyCost() string {
	if o == nil || IsNil(o.HourlyCost.Get()) {
		var ret string
		return ret
	}
	return *o.HourlyCost.Get()
}

// GetHourlyCostOk returns a tuple with the HourlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetHourlyCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HourlyCost.Get(), o.HourlyCost.IsSet()
}

// HasHourlyCost returns a boolean if a field has been set.
func (o *CostComponent) HasHourlyCost() bool {
	if o != nil && o.HourlyCost.IsSet() {
		return true
	}

	return false
}

// SetHourlyCost gets a reference to the given NullableString and assigns it to the HourlyCost field.
func (o *CostComponent) SetHourlyCost(v string) {
	o.HourlyCost.Set(&v)
}
// SetHourlyCostNil sets the value for HourlyCost to be an explicit nil
func (o *CostComponent) SetHourlyCostNil() {
	o.HourlyCost.Set(nil)
}

// UnsetHourlyCost ensures that no value is present for HourlyCost, not even an explicit nil
func (o *CostComponent) UnsetHourlyCost() {
	o.HourlyCost.Unset()
}

// GetMonthlyCost returns the MonthlyCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetMonthlyCost() string {
	if o == nil || IsNil(o.MonthlyCost.Get()) {
		var ret string
		return ret
	}
	return *o.MonthlyCost.Get()
}

// GetMonthlyCostOk returns a tuple with the MonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetMonthlyCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyCost.Get(), o.MonthlyCost.IsSet()
}

// HasMonthlyCost returns a boolean if a field has been set.
func (o *CostComponent) HasMonthlyCost() bool {
	if o != nil && o.MonthlyCost.IsSet() {
		return true
	}

	return false
}

// SetMonthlyCost gets a reference to the given NullableString and assigns it to the MonthlyCost field.
func (o *CostComponent) SetMonthlyCost(v string) {
	o.MonthlyCost.Set(&v)
}
// SetMonthlyCostNil sets the value for MonthlyCost to be an explicit nil
func (o *CostComponent) SetMonthlyCostNil() {
	o.MonthlyCost.Set(nil)
}

// UnsetMonthlyCost ensures that no value is present for MonthlyCost, not even an explicit nil
func (o *CostComponent) UnsetMonthlyCost() {
	o.MonthlyCost.Unset()
}

// GetPriceNotFound returns the PriceNotFound field value if set, zero value otherwise.
func (o *CostComponent) GetPriceNotFound() bool {
	if o == nil || IsNil(o.PriceNotFound) {
		var ret bool
		return ret
	}
	return *o.PriceNotFound
}

// GetPriceNotFoundOk returns a tuple with the PriceNotFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostComponent) GetPriceNotFoundOk() (*bool, bool) {
	if o == nil || IsNil(o.PriceNotFound) {
		return nil, false
	}
	return o.PriceNotFound, true
}

// HasPriceNotFound returns a boolean if a field has been set.
func (o *CostComponent) HasPriceNotFound() bool {
	if o != nil && !IsNil(o.PriceNotFound) {
		return true
	}

	return false
}

// SetPriceNotFound gets a reference to the given bool and assigns it to the PriceNotFound field.
func (o *CostComponent) SetPriceNotFound(v bool) {
	o.PriceNotFound = &v
}

// GetUsageBased returns the UsageBased field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CostComponent) GetUsageBased() bool {
	if o == nil || IsNil(o.UsageBased.Get()) {
		var ret bool
		return ret
	}
	return *o.UsageBased.Get()
}

// GetUsageBasedOk returns a tuple with the UsageBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CostComponent) GetUsageBasedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageBased.Get(), o.UsageBased.IsSet()
}

// HasUsageBased returns a boolean if a field has been set.
func (o *CostComponent) HasUsageBased() bool {
	if o != nil && o.UsageBased.IsSet() {
		return true
	}

	return false
}

// SetUsageBased gets a reference to the given NullableBool and assigns it to the UsageBased field.
func (o *CostComponent) SetUsageBased(v bool) {
	o.UsageBased.Set(&v)
}
// SetUsageBasedNil sets the value for UsageBased to be an explicit nil
func (o *CostComponent) SetUsageBasedNil() {
	o.UsageBased.Set(nil)
}

// UnsetUsageBased ensures that no value is present for UsageBased, not even an explicit nil
func (o *CostComponent) UnsetUsageBased() {
	o.UsageBased.Unset()
}

func (o CostComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Unit.IsSet() {
		toSerialize["unit"] = o.Unit.Get()
	}
	if o.HourlyQuantity.IsSet() {
		toSerialize["hourlyQuantity"] = o.HourlyQuantity.Get()
	}
	if o.MonthlyQuantity.IsSet() {
		toSerialize["monthlyQuantity"] = o.MonthlyQuantity.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.HourlyCost.IsSet() {
		toSerialize["hourlyCost"] = o.HourlyCost.Get()
	}
	if o.MonthlyCost.IsSet() {
		toSerialize["monthlyCost"] = o.MonthlyCost.Get()
	}
	if !IsNil(o.PriceNotFound) {
		toSerialize["priceNotFound"] = o.PriceNotFound
	}
	if o.UsageBased.IsSet() {
		toSerialize["usageBased"] = o.UsageBased.Get()
	}
	return toSerialize, nil
}

type NullableCostComponent struct {
	value *CostComponent
	isSet bool
}

func (v NullableCostComponent) Get() *CostComponent {
	return v.value
}

func (v *NullableCostComponent) Set(val *CostComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableCostComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableCostComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostComponent(val *CostComponent) *NullableCostComponent {
	return &NullableCostComponent{value: val, isSet: true}
}

func (v NullableCostComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


