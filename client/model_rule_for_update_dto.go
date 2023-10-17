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

// checks if the RuleForUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleForUpdateDto{}

// RuleForUpdateDto struct for RuleForUpdateDto
type RuleForUpdateDto struct {
	Name                  NullableString             `json:"name,omitempty"`
	MetricName            NullableString             `json:"metricName,omitempty"`
	Type                  *PrometheusType            `json:"type,omitempty"`
	Price                 NullableFloat64            `json:"price,omitempty"`
	LabelsToAdd           []PrometheusLabelListDto   `json:"labelsToAdd,omitempty"`
	LabelsToDelete        []PrometheusLabelDeleteDto `json:"labelsToDelete,omitempty"`
	LabelsToUpdate        []PrometheusLabelUpdateDto `json:"labelsToUpdate,omitempty"`
	OrganizationId        NullableInt32              `json:"organizationId,omitempty"`
	RuleDiscountRate      NullableInt32              `json:"ruleDiscountRate,omitempty"`
	OperationCredentialId NullableInt32              `json:"operationCredentialId,omitempty"`
}

// NewRuleForUpdateDto instantiates a new RuleForUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleForUpdateDto() *RuleForUpdateDto {
	this := RuleForUpdateDto{}
	return &this
}

// NewRuleForUpdateDtoWithDefaults instantiates a new RuleForUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleForUpdateDtoWithDefaults() *RuleForUpdateDto {
	this := RuleForUpdateDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RuleForUpdateDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *RuleForUpdateDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RuleForUpdateDto) UnsetName() {
	o.Name.Unset()
}

// GetMetricName returns the MetricName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetMetricName() string {
	if o == nil || IsNil(o.MetricName.Get()) {
		var ret string
		return ret
	}
	return *o.MetricName.Get()
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricName.Get(), o.MetricName.IsSet()
}

// HasMetricName returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasMetricName() bool {
	if o != nil && o.MetricName.IsSet() {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given NullableString and assigns it to the MetricName field.
func (o *RuleForUpdateDto) SetMetricName(v string) {
	o.MetricName.Set(&v)
}

// SetMetricNameNil sets the value for MetricName to be an explicit nil
func (o *RuleForUpdateDto) SetMetricNameNil() {
	o.MetricName.Set(nil)
}

// UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
func (o *RuleForUpdateDto) UnsetMetricName() {
	o.MetricName.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RuleForUpdateDto) GetType() PrometheusType {
	if o == nil || IsNil(o.Type) {
		var ret PrometheusType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleForUpdateDto) GetTypeOk() (*PrometheusType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PrometheusType and assigns it to the Type field.
func (o *RuleForUpdateDto) SetType(v PrometheusType) {
	o.Type = &v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *RuleForUpdateDto) SetPrice(v float64) {
	o.Price.Set(&v)
}

// SetPriceNil sets the value for Price to be an explicit nil
func (o *RuleForUpdateDto) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *RuleForUpdateDto) UnsetPrice() {
	o.Price.Unset()
}

// GetLabelsToAdd returns the LabelsToAdd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetLabelsToAdd() []PrometheusLabelListDto {
	if o == nil {
		var ret []PrometheusLabelListDto
		return ret
	}
	return o.LabelsToAdd
}

// GetLabelsToAddOk returns a tuple with the LabelsToAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetLabelsToAddOk() ([]PrometheusLabelListDto, bool) {
	if o == nil || IsNil(o.LabelsToAdd) {
		return nil, false
	}
	return o.LabelsToAdd, true
}

// HasLabelsToAdd returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasLabelsToAdd() bool {
	if o != nil && IsNil(o.LabelsToAdd) {
		return true
	}

	return false
}

// SetLabelsToAdd gets a reference to the given []PrometheusLabelListDto and assigns it to the LabelsToAdd field.
func (o *RuleForUpdateDto) SetLabelsToAdd(v []PrometheusLabelListDto) {
	o.LabelsToAdd = v
}

// GetLabelsToDelete returns the LabelsToDelete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetLabelsToDelete() []PrometheusLabelDeleteDto {
	if o == nil {
		var ret []PrometheusLabelDeleteDto
		return ret
	}
	return o.LabelsToDelete
}

// GetLabelsToDeleteOk returns a tuple with the LabelsToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetLabelsToDeleteOk() ([]PrometheusLabelDeleteDto, bool) {
	if o == nil || IsNil(o.LabelsToDelete) {
		return nil, false
	}
	return o.LabelsToDelete, true
}

// HasLabelsToDelete returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasLabelsToDelete() bool {
	if o != nil && IsNil(o.LabelsToDelete) {
		return true
	}

	return false
}

// SetLabelsToDelete gets a reference to the given []PrometheusLabelDeleteDto and assigns it to the LabelsToDelete field.
func (o *RuleForUpdateDto) SetLabelsToDelete(v []PrometheusLabelDeleteDto) {
	o.LabelsToDelete = v
}

// GetLabelsToUpdate returns the LabelsToUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetLabelsToUpdate() []PrometheusLabelUpdateDto {
	if o == nil {
		var ret []PrometheusLabelUpdateDto
		return ret
	}
	return o.LabelsToUpdate
}

// GetLabelsToUpdateOk returns a tuple with the LabelsToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetLabelsToUpdateOk() ([]PrometheusLabelUpdateDto, bool) {
	if o == nil || IsNil(o.LabelsToUpdate) {
		return nil, false
	}
	return o.LabelsToUpdate, true
}

// HasLabelsToUpdate returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasLabelsToUpdate() bool {
	if o != nil && IsNil(o.LabelsToUpdate) {
		return true
	}

	return false
}

// SetLabelsToUpdate gets a reference to the given []PrometheusLabelUpdateDto and assigns it to the LabelsToUpdate field.
func (o *RuleForUpdateDto) SetLabelsToUpdate(v []PrometheusLabelUpdateDto) {
	o.LabelsToUpdate = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *RuleForUpdateDto) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *RuleForUpdateDto) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *RuleForUpdateDto) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetRuleDiscountRate returns the RuleDiscountRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetRuleDiscountRate() int32 {
	if o == nil || IsNil(o.RuleDiscountRate.Get()) {
		var ret int32
		return ret
	}
	return *o.RuleDiscountRate.Get()
}

// GetRuleDiscountRateOk returns a tuple with the RuleDiscountRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetRuleDiscountRateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RuleDiscountRate.Get(), o.RuleDiscountRate.IsSet()
}

// HasRuleDiscountRate returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasRuleDiscountRate() bool {
	if o != nil && o.RuleDiscountRate.IsSet() {
		return true
	}

	return false
}

// SetRuleDiscountRate gets a reference to the given NullableInt32 and assigns it to the RuleDiscountRate field.
func (o *RuleForUpdateDto) SetRuleDiscountRate(v int32) {
	o.RuleDiscountRate.Set(&v)
}

// SetRuleDiscountRateNil sets the value for RuleDiscountRate to be an explicit nil
func (o *RuleForUpdateDto) SetRuleDiscountRateNil() {
	o.RuleDiscountRate.Set(nil)
}

// UnsetRuleDiscountRate ensures that no value is present for RuleDiscountRate, not even an explicit nil
func (o *RuleForUpdateDto) UnsetRuleDiscountRate() {
	o.RuleDiscountRate.Unset()
}

// GetOperationCredentialId returns the OperationCredentialId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RuleForUpdateDto) GetOperationCredentialId() int32 {
	if o == nil || IsNil(o.OperationCredentialId.Get()) {
		var ret int32
		return ret
	}
	return *o.OperationCredentialId.Get()
}

// GetOperationCredentialIdOk returns a tuple with the OperationCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RuleForUpdateDto) GetOperationCredentialIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationCredentialId.Get(), o.OperationCredentialId.IsSet()
}

// HasOperationCredentialId returns a boolean if a field has been set.
func (o *RuleForUpdateDto) HasOperationCredentialId() bool {
	if o != nil && o.OperationCredentialId.IsSet() {
		return true
	}

	return false
}

// SetOperationCredentialId gets a reference to the given NullableInt32 and assigns it to the OperationCredentialId field.
func (o *RuleForUpdateDto) SetOperationCredentialId(v int32) {
	o.OperationCredentialId.Set(&v)
}

// SetOperationCredentialIdNil sets the value for OperationCredentialId to be an explicit nil
func (o *RuleForUpdateDto) SetOperationCredentialIdNil() {
	o.OperationCredentialId.Set(nil)
}

// UnsetOperationCredentialId ensures that no value is present for OperationCredentialId, not even an explicit nil
func (o *RuleForUpdateDto) UnsetOperationCredentialId() {
	o.OperationCredentialId.Unset()
}

func (o RuleForUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleForUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.MetricName.IsSet() {
		toSerialize["metricName"] = o.MetricName.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.LabelsToAdd != nil {
		toSerialize["labelsToAdd"] = o.LabelsToAdd
	}
	if o.LabelsToDelete != nil {
		toSerialize["labelsToDelete"] = o.LabelsToDelete
	}
	if o.LabelsToUpdate != nil {
		toSerialize["labelsToUpdate"] = o.LabelsToUpdate
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.RuleDiscountRate.IsSet() {
		toSerialize["ruleDiscountRate"] = o.RuleDiscountRate.Get()
	}
	if o.OperationCredentialId.IsSet() {
		toSerialize["operationCredentialId"] = o.OperationCredentialId.Get()
	}
	return toSerialize, nil
}

type NullableRuleForUpdateDto struct {
	value *RuleForUpdateDto
	isSet bool
}

func (v NullableRuleForUpdateDto) Get() *RuleForUpdateDto {
	return v.value
}

func (v *NullableRuleForUpdateDto) Set(val *RuleForUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleForUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleForUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleForUpdateDto(val *RuleForUpdateDto) *NullableRuleForUpdateDto {
	return &NullableRuleForUpdateDto{value: val, isSet: true}
}

func (v NullableRuleForUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleForUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
