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

// checks if the GroupedBillingListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedBillingListQuery{}

// GroupedBillingListQuery struct for GroupedBillingListQuery
type GroupedBillingListQuery struct {
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	PeriodDuration *BillingPeriod `json:"periodDuration,omitempty"`
	IsDeleted NullableBool `json:"isDeleted,omitempty"`
}

// NewGroupedBillingListQuery instantiates a new GroupedBillingListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedBillingListQuery() *GroupedBillingListQuery {
	this := GroupedBillingListQuery{}
	return &this
}

// NewGroupedBillingListQueryWithDefaults instantiates a new GroupedBillingListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedBillingListQueryWithDefaults() *GroupedBillingListQuery {
	this := GroupedBillingListQuery{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedBillingListQuery) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedBillingListQuery) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GroupedBillingListQuery) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *GroupedBillingListQuery) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *GroupedBillingListQuery) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *GroupedBillingListQuery) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetPeriodDuration returns the PeriodDuration field value if set, zero value otherwise.
func (o *GroupedBillingListQuery) GetPeriodDuration() BillingPeriod {
	if o == nil || IsNil(o.PeriodDuration) {
		var ret BillingPeriod
		return ret
	}
	return *o.PeriodDuration
}

// GetPeriodDurationOk returns a tuple with the PeriodDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedBillingListQuery) GetPeriodDurationOk() (*BillingPeriod, bool) {
	if o == nil || IsNil(o.PeriodDuration) {
		return nil, false
	}
	return o.PeriodDuration, true
}

// HasPeriodDuration returns a boolean if a field has been set.
func (o *GroupedBillingListQuery) HasPeriodDuration() bool {
	if o != nil && !IsNil(o.PeriodDuration) {
		return true
	}

	return false
}

// SetPeriodDuration gets a reference to the given BillingPeriod and assigns it to the PeriodDuration field.
func (o *GroupedBillingListQuery) SetPeriodDuration(v BillingPeriod) {
	o.PeriodDuration = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedBillingListQuery) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDeleted.Get()
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedBillingListQuery) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDeleted.Get(), o.IsDeleted.IsSet()
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *GroupedBillingListQuery) HasIsDeleted() bool {
	if o != nil && o.IsDeleted.IsSet() {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given NullableBool and assigns it to the IsDeleted field.
func (o *GroupedBillingListQuery) SetIsDeleted(v bool) {
	o.IsDeleted.Set(&v)
}
// SetIsDeletedNil sets the value for IsDeleted to be an explicit nil
func (o *GroupedBillingListQuery) SetIsDeletedNil() {
	o.IsDeleted.Set(nil)
}

// UnsetIsDeleted ensures that no value is present for IsDeleted, not even an explicit nil
func (o *GroupedBillingListQuery) UnsetIsDeleted() {
	o.IsDeleted.Unset()
}

func (o GroupedBillingListQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedBillingListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if !IsNil(o.PeriodDuration) {
		toSerialize["periodDuration"] = o.PeriodDuration
	}
	if o.IsDeleted.IsSet() {
		toSerialize["isDeleted"] = o.IsDeleted.Get()
	}
	return toSerialize, nil
}

type NullableGroupedBillingListQuery struct {
	value *GroupedBillingListQuery
	isSet bool
}

func (v NullableGroupedBillingListQuery) Get() *GroupedBillingListQuery {
	return v.value
}

func (v *NullableGroupedBillingListQuery) Set(val *GroupedBillingListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedBillingListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedBillingListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedBillingListQuery(val *GroupedBillingListQuery) *NullableGroupedBillingListQuery {
	return &NullableGroupedBillingListQuery{value: val, isSet: true}
}

func (v NullableGroupedBillingListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedBillingListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


