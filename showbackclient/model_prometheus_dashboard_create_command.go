/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the PrometheusDashboardCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusDashboardCreateCommand{}

// PrometheusDashboardCreateCommand struct for PrometheusDashboardCreateCommand
type PrometheusDashboardCreateCommand struct {
	ProjectId    *int32         `json:"projectId,omitempty"`
	Name         NullableString `json:"name,omitempty"`
	Expression   NullableString `json:"expression,omitempty"`
	Description  NullableString `json:"description,omitempty"`
	CategoryName NullableString `json:"categoryName,omitempty"`
}

// NewPrometheusDashboardCreateCommand instantiates a new PrometheusDashboardCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusDashboardCreateCommand() *PrometheusDashboardCreateCommand {
	this := PrometheusDashboardCreateCommand{}
	return &this
}

// NewPrometheusDashboardCreateCommandWithDefaults instantiates a new PrometheusDashboardCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusDashboardCreateCommandWithDefaults() *PrometheusDashboardCreateCommand {
	this := PrometheusDashboardCreateCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *PrometheusDashboardCreateCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusDashboardCreateCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *PrometheusDashboardCreateCommand) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *PrometheusDashboardCreateCommand) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardCreateCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PrometheusDashboardCreateCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PrometheusDashboardCreateCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PrometheusDashboardCreateCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PrometheusDashboardCreateCommand) UnsetName() {
	o.Name.Unset()
}

// GetExpression returns the Expression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardCreateCommand) GetExpression() string {
	if o == nil || IsNil(o.Expression.Get()) {
		var ret string
		return ret
	}
	return *o.Expression.Get()
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardCreateCommand) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expression.Get(), o.Expression.IsSet()
}

// HasExpression returns a boolean if a field has been set.
func (o *PrometheusDashboardCreateCommand) HasExpression() bool {
	if o != nil && o.Expression.IsSet() {
		return true
	}

	return false
}

// SetExpression gets a reference to the given NullableString and assigns it to the Expression field.
func (o *PrometheusDashboardCreateCommand) SetExpression(v string) {
	o.Expression.Set(&v)
}

// SetExpressionNil sets the value for Expression to be an explicit nil
func (o *PrometheusDashboardCreateCommand) SetExpressionNil() {
	o.Expression.Set(nil)
}

// UnsetExpression ensures that no value is present for Expression, not even an explicit nil
func (o *PrometheusDashboardCreateCommand) UnsetExpression() {
	o.Expression.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardCreateCommand) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardCreateCommand) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PrometheusDashboardCreateCommand) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PrometheusDashboardCreateCommand) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PrometheusDashboardCreateCommand) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PrometheusDashboardCreateCommand) UnsetDescription() {
	o.Description.Unset()
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusDashboardCreateCommand) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryName.Get()
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusDashboardCreateCommand) GetCategoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryName.Get(), o.CategoryName.IsSet()
}

// HasCategoryName returns a boolean if a field has been set.
func (o *PrometheusDashboardCreateCommand) HasCategoryName() bool {
	if o != nil && o.CategoryName.IsSet() {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given NullableString and assigns it to the CategoryName field.
func (o *PrometheusDashboardCreateCommand) SetCategoryName(v string) {
	o.CategoryName.Set(&v)
}

// SetCategoryNameNil sets the value for CategoryName to be an explicit nil
func (o *PrometheusDashboardCreateCommand) SetCategoryNameNil() {
	o.CategoryName.Set(nil)
}

// UnsetCategoryName ensures that no value is present for CategoryName, not even an explicit nil
func (o *PrometheusDashboardCreateCommand) UnsetCategoryName() {
	o.CategoryName.Unset()
}

func (o PrometheusDashboardCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusDashboardCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Expression.IsSet() {
		toSerialize["expression"] = o.Expression.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CategoryName.IsSet() {
		toSerialize["categoryName"] = o.CategoryName.Get()
	}
	return toSerialize, nil
}

type NullablePrometheusDashboardCreateCommand struct {
	value *PrometheusDashboardCreateCommand
	isSet bool
}

func (v NullablePrometheusDashboardCreateCommand) Get() *PrometheusDashboardCreateCommand {
	return v.value
}

func (v *NullablePrometheusDashboardCreateCommand) Set(val *PrometheusDashboardCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusDashboardCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusDashboardCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusDashboardCreateCommand(val *PrometheusDashboardCreateCommand) *NullablePrometheusDashboardCreateCommand {
	return &NullablePrometheusDashboardCreateCommand{value: val, isSet: true}
}

func (v NullablePrometheusDashboardCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusDashboardCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
