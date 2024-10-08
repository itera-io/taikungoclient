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

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	Name NullableString `json:"name,omitempty"`
	ResourceType NullableString `json:"resourceType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	HourlyCost NullableString `json:"hourlyCost,omitempty"`
	MonthlyCost NullableString `json:"monthlyCost,omitempty"`
	MonthlyUsageCost NullableString `json:"monthlyUsageCost,omitempty"`
	CostComponents []CostComponent `json:"costComponents,omitempty"`
	Subresources []Subresource `json:"subresources,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource() *Resource {
	this := Resource{}
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Resource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Resource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Resource) UnsetName() {
	o.Name.Unset()
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// HasResourceType returns a boolean if a field has been set.
func (o *Resource) HasResourceType() bool {
	if o != nil && o.ResourceType.IsSet() {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given NullableString and assigns it to the ResourceType field.
func (o *Resource) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}
// SetResourceTypeNil sets the value for ResourceType to be an explicit nil
func (o *Resource) SetResourceTypeNil() {
	o.ResourceType.Set(nil)
}

// UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
func (o *Resource) UnsetResourceType() {
	o.ResourceType.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Resource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Resource) SetTags(v map[string]string) {
	o.Tags = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Resource) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Resource) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetHourlyCost returns the HourlyCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetHourlyCost() string {
	if o == nil || IsNil(o.HourlyCost.Get()) {
		var ret string
		return ret
	}
	return *o.HourlyCost.Get()
}

// GetHourlyCostOk returns a tuple with the HourlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetHourlyCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HourlyCost.Get(), o.HourlyCost.IsSet()
}

// HasHourlyCost returns a boolean if a field has been set.
func (o *Resource) HasHourlyCost() bool {
	if o != nil && o.HourlyCost.IsSet() {
		return true
	}

	return false
}

// SetHourlyCost gets a reference to the given NullableString and assigns it to the HourlyCost field.
func (o *Resource) SetHourlyCost(v string) {
	o.HourlyCost.Set(&v)
}
// SetHourlyCostNil sets the value for HourlyCost to be an explicit nil
func (o *Resource) SetHourlyCostNil() {
	o.HourlyCost.Set(nil)
}

// UnsetHourlyCost ensures that no value is present for HourlyCost, not even an explicit nil
func (o *Resource) UnsetHourlyCost() {
	o.HourlyCost.Unset()
}

// GetMonthlyCost returns the MonthlyCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetMonthlyCost() string {
	if o == nil || IsNil(o.MonthlyCost.Get()) {
		var ret string
		return ret
	}
	return *o.MonthlyCost.Get()
}

// GetMonthlyCostOk returns a tuple with the MonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetMonthlyCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyCost.Get(), o.MonthlyCost.IsSet()
}

// HasMonthlyCost returns a boolean if a field has been set.
func (o *Resource) HasMonthlyCost() bool {
	if o != nil && o.MonthlyCost.IsSet() {
		return true
	}

	return false
}

// SetMonthlyCost gets a reference to the given NullableString and assigns it to the MonthlyCost field.
func (o *Resource) SetMonthlyCost(v string) {
	o.MonthlyCost.Set(&v)
}
// SetMonthlyCostNil sets the value for MonthlyCost to be an explicit nil
func (o *Resource) SetMonthlyCostNil() {
	o.MonthlyCost.Set(nil)
}

// UnsetMonthlyCost ensures that no value is present for MonthlyCost, not even an explicit nil
func (o *Resource) UnsetMonthlyCost() {
	o.MonthlyCost.Unset()
}

// GetMonthlyUsageCost returns the MonthlyUsageCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetMonthlyUsageCost() string {
	if o == nil || IsNil(o.MonthlyUsageCost.Get()) {
		var ret string
		return ret
	}
	return *o.MonthlyUsageCost.Get()
}

// GetMonthlyUsageCostOk returns a tuple with the MonthlyUsageCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetMonthlyUsageCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MonthlyUsageCost.Get(), o.MonthlyUsageCost.IsSet()
}

// HasMonthlyUsageCost returns a boolean if a field has been set.
func (o *Resource) HasMonthlyUsageCost() bool {
	if o != nil && o.MonthlyUsageCost.IsSet() {
		return true
	}

	return false
}

// SetMonthlyUsageCost gets a reference to the given NullableString and assigns it to the MonthlyUsageCost field.
func (o *Resource) SetMonthlyUsageCost(v string) {
	o.MonthlyUsageCost.Set(&v)
}
// SetMonthlyUsageCostNil sets the value for MonthlyUsageCost to be an explicit nil
func (o *Resource) SetMonthlyUsageCostNil() {
	o.MonthlyUsageCost.Set(nil)
}

// UnsetMonthlyUsageCost ensures that no value is present for MonthlyUsageCost, not even an explicit nil
func (o *Resource) UnsetMonthlyUsageCost() {
	o.MonthlyUsageCost.Unset()
}

// GetCostComponents returns the CostComponents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetCostComponents() []CostComponent {
	if o == nil {
		var ret []CostComponent
		return ret
	}
	return o.CostComponents
}

// GetCostComponentsOk returns a tuple with the CostComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetCostComponentsOk() ([]CostComponent, bool) {
	if o == nil || IsNil(o.CostComponents) {
		return nil, false
	}
	return o.CostComponents, true
}

// HasCostComponents returns a boolean if a field has been set.
func (o *Resource) HasCostComponents() bool {
	if o != nil && !IsNil(o.CostComponents) {
		return true
	}

	return false
}

// SetCostComponents gets a reference to the given []CostComponent and assigns it to the CostComponents field.
func (o *Resource) SetCostComponents(v []CostComponent) {
	o.CostComponents = v
}

// GetSubresources returns the Subresources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Resource) GetSubresources() []Subresource {
	if o == nil {
		var ret []Subresource
		return ret
	}
	return o.Subresources
}

// GetSubresourcesOk returns a tuple with the Subresources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Resource) GetSubresourcesOk() ([]Subresource, bool) {
	if o == nil || IsNil(o.Subresources) {
		return nil, false
	}
	return o.Subresources, true
}

// HasSubresources returns a boolean if a field has been set.
func (o *Resource) HasSubresources() bool {
	if o != nil && !IsNil(o.Subresources) {
		return true
	}

	return false
}

// SetSubresources gets a reference to the given []Subresource and assigns it to the Subresources field.
func (o *Resource) SetSubresources(v []Subresource) {
	o.Subresources = v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ResourceType.IsSet() {
		toSerialize["resourceType"] = o.ResourceType.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.HourlyCost.IsSet() {
		toSerialize["hourlyCost"] = o.HourlyCost.Get()
	}
	if o.MonthlyCost.IsSet() {
		toSerialize["monthlyCost"] = o.MonthlyCost.Get()
	}
	if o.MonthlyUsageCost.IsSet() {
		toSerialize["monthlyUsageCost"] = o.MonthlyUsageCost.Get()
	}
	if o.CostComponents != nil {
		toSerialize["costComponents"] = o.CostComponents
	}
	if o.Subresources != nil {
		toSerialize["subresources"] = o.Subresources
	}
	return toSerialize, nil
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


