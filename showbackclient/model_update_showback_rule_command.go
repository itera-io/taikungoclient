/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the UpdateShowbackRuleCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShowbackRuleCommand{}

// UpdateShowbackRuleCommand struct for UpdateShowbackRuleCommand
type UpdateShowbackRuleCommand struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	ByLabel NullableString `json:"byLabel,omitempty"`
	MetricName NullableString `json:"metricName,omitempty"`
	Kind *EShowbackType `json:"kind,omitempty"`
	Type *EPrometheusType `json:"type,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	ProjectAlertLimit NullableInt32 `json:"projectAlertLimit,omitempty"`
	GlobalAlertLimit NullableInt32 `json:"globalAlertLimit,omitempty"`
	Labels []ShowbackLabelCreateDto `json:"labels,omitempty"`
}

// NewUpdateShowbackRuleCommand instantiates a new UpdateShowbackRuleCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShowbackRuleCommand() *UpdateShowbackRuleCommand {
	this := UpdateShowbackRuleCommand{}
	return &this
}

// NewUpdateShowbackRuleCommandWithDefaults instantiates a new UpdateShowbackRuleCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShowbackRuleCommandWithDefaults() *UpdateShowbackRuleCommand {
	this := UpdateShowbackRuleCommand{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateShowbackRuleCommand) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShowbackRuleCommand) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateShowbackRuleCommand) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateShowbackRuleCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateShowbackRuleCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateShowbackRuleCommand) UnsetName() {
	o.Name.Unset()
}

// GetByLabel returns the ByLabel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetByLabel() string {
	if o == nil || IsNil(o.ByLabel.Get()) {
		var ret string
		return ret
	}
	return *o.ByLabel.Get()
}

// GetByLabelOk returns a tuple with the ByLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetByLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ByLabel.Get(), o.ByLabel.IsSet()
}

// HasByLabel returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasByLabel() bool {
	if o != nil && o.ByLabel.IsSet() {
		return true
	}

	return false
}

// SetByLabel gets a reference to the given NullableString and assigns it to the ByLabel field.
func (o *UpdateShowbackRuleCommand) SetByLabel(v string) {
	o.ByLabel.Set(&v)
}
// SetByLabelNil sets the value for ByLabel to be an explicit nil
func (o *UpdateShowbackRuleCommand) SetByLabelNil() {
	o.ByLabel.Set(nil)
}

// UnsetByLabel ensures that no value is present for ByLabel, not even an explicit nil
func (o *UpdateShowbackRuleCommand) UnsetByLabel() {
	o.ByLabel.Unset()
}

// GetMetricName returns the MetricName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetMetricName() string {
	if o == nil || IsNil(o.MetricName.Get()) {
		var ret string
		return ret
	}
	return *o.MetricName.Get()
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricName.Get(), o.MetricName.IsSet()
}

// HasMetricName returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasMetricName() bool {
	if o != nil && o.MetricName.IsSet() {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given NullableString and assigns it to the MetricName field.
func (o *UpdateShowbackRuleCommand) SetMetricName(v string) {
	o.MetricName.Set(&v)
}
// SetMetricNameNil sets the value for MetricName to be an explicit nil
func (o *UpdateShowbackRuleCommand) SetMetricNameNil() {
	o.MetricName.Set(nil)
}

// UnsetMetricName ensures that no value is present for MetricName, not even an explicit nil
func (o *UpdateShowbackRuleCommand) UnsetMetricName() {
	o.MetricName.Unset()
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *UpdateShowbackRuleCommand) GetKind() EShowbackType {
	if o == nil || IsNil(o.Kind) {
		var ret EShowbackType
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShowbackRuleCommand) GetKindOk() (*EShowbackType, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given EShowbackType and assigns it to the Kind field.
func (o *UpdateShowbackRuleCommand) SetKind(v EShowbackType) {
	o.Kind = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateShowbackRuleCommand) GetType() EPrometheusType {
	if o == nil || IsNil(o.Type) {
		var ret EPrometheusType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateShowbackRuleCommand) GetTypeOk() (*EPrometheusType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EPrometheusType and assigns it to the Type field.
func (o *UpdateShowbackRuleCommand) SetType(v EPrometheusType) {
	o.Type = &v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *UpdateShowbackRuleCommand) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *UpdateShowbackRuleCommand) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *UpdateShowbackRuleCommand) UnsetPrice() {
	o.Price.Unset()
}

// GetProjectAlertLimit returns the ProjectAlertLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetProjectAlertLimit() int32 {
	if o == nil || IsNil(o.ProjectAlertLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectAlertLimit.Get()
}

// GetProjectAlertLimitOk returns a tuple with the ProjectAlertLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetProjectAlertLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectAlertLimit.Get(), o.ProjectAlertLimit.IsSet()
}

// HasProjectAlertLimit returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasProjectAlertLimit() bool {
	if o != nil && o.ProjectAlertLimit.IsSet() {
		return true
	}

	return false
}

// SetProjectAlertLimit gets a reference to the given NullableInt32 and assigns it to the ProjectAlertLimit field.
func (o *UpdateShowbackRuleCommand) SetProjectAlertLimit(v int32) {
	o.ProjectAlertLimit.Set(&v)
}
// SetProjectAlertLimitNil sets the value for ProjectAlertLimit to be an explicit nil
func (o *UpdateShowbackRuleCommand) SetProjectAlertLimitNil() {
	o.ProjectAlertLimit.Set(nil)
}

// UnsetProjectAlertLimit ensures that no value is present for ProjectAlertLimit, not even an explicit nil
func (o *UpdateShowbackRuleCommand) UnsetProjectAlertLimit() {
	o.ProjectAlertLimit.Unset()
}

// GetGlobalAlertLimit returns the GlobalAlertLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetGlobalAlertLimit() int32 {
	if o == nil || IsNil(o.GlobalAlertLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.GlobalAlertLimit.Get()
}

// GetGlobalAlertLimitOk returns a tuple with the GlobalAlertLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetGlobalAlertLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalAlertLimit.Get(), o.GlobalAlertLimit.IsSet()
}

// HasGlobalAlertLimit returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasGlobalAlertLimit() bool {
	if o != nil && o.GlobalAlertLimit.IsSet() {
		return true
	}

	return false
}

// SetGlobalAlertLimit gets a reference to the given NullableInt32 and assigns it to the GlobalAlertLimit field.
func (o *UpdateShowbackRuleCommand) SetGlobalAlertLimit(v int32) {
	o.GlobalAlertLimit.Set(&v)
}
// SetGlobalAlertLimitNil sets the value for GlobalAlertLimit to be an explicit nil
func (o *UpdateShowbackRuleCommand) SetGlobalAlertLimitNil() {
	o.GlobalAlertLimit.Set(nil)
}

// UnsetGlobalAlertLimit ensures that no value is present for GlobalAlertLimit, not even an explicit nil
func (o *UpdateShowbackRuleCommand) UnsetGlobalAlertLimit() {
	o.GlobalAlertLimit.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateShowbackRuleCommand) GetLabels() []ShowbackLabelCreateDto {
	if o == nil {
		var ret []ShowbackLabelCreateDto
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateShowbackRuleCommand) GetLabelsOk() ([]ShowbackLabelCreateDto, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateShowbackRuleCommand) HasLabels() bool {
	if o != nil && IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ShowbackLabelCreateDto and assigns it to the Labels field.
func (o *UpdateShowbackRuleCommand) SetLabels(v []ShowbackLabelCreateDto) {
	o.Labels = v
}

func (o UpdateShowbackRuleCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateShowbackRuleCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ByLabel.IsSet() {
		toSerialize["byLabel"] = o.ByLabel.Get()
	}
	if o.MetricName.IsSet() {
		toSerialize["metricName"] = o.MetricName.Get()
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.ProjectAlertLimit.IsSet() {
		toSerialize["projectAlertLimit"] = o.ProjectAlertLimit.Get()
	}
	if o.GlobalAlertLimit.IsSet() {
		toSerialize["globalAlertLimit"] = o.GlobalAlertLimit.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableUpdateShowbackRuleCommand struct {
	value *UpdateShowbackRuleCommand
	isSet bool
}

func (v NullableUpdateShowbackRuleCommand) Get() *UpdateShowbackRuleCommand {
	return v.value
}

func (v *NullableUpdateShowbackRuleCommand) Set(val *UpdateShowbackRuleCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShowbackRuleCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShowbackRuleCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShowbackRuleCommand(val *UpdateShowbackRuleCommand) *NullableUpdateShowbackRuleCommand {
	return &NullableUpdateShowbackRuleCommand{value: val, isSet: true}
}

func (v NullableUpdateShowbackRuleCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateShowbackRuleCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


