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

// checks if the GroupedPrometheusBillingListQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedPrometheusBillingListQuery{}

// GroupedPrometheusBillingListQuery struct for GroupedPrometheusBillingListQuery
type GroupedPrometheusBillingListQuery struct {
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	PeriodDuration *BillingPeriod `json:"periodDuration,omitempty"`
}

// NewGroupedPrometheusBillingListQuery instantiates a new GroupedPrometheusBillingListQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedPrometheusBillingListQuery() *GroupedPrometheusBillingListQuery {
	this := GroupedPrometheusBillingListQuery{}
	return &this
}

// NewGroupedPrometheusBillingListQueryWithDefaults instantiates a new GroupedPrometheusBillingListQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedPrometheusBillingListQueryWithDefaults() *GroupedPrometheusBillingListQuery {
	this := GroupedPrometheusBillingListQuery{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedPrometheusBillingListQuery) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedPrometheusBillingListQuery) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GroupedPrometheusBillingListQuery) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *GroupedPrometheusBillingListQuery) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *GroupedPrometheusBillingListQuery) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *GroupedPrometheusBillingListQuery) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetPeriodDuration returns the PeriodDuration field value if set, zero value otherwise.
func (o *GroupedPrometheusBillingListQuery) GetPeriodDuration() BillingPeriod {
	if o == nil || IsNil(o.PeriodDuration) {
		var ret BillingPeriod
		return ret
	}
	return *o.PeriodDuration
}

// GetPeriodDurationOk returns a tuple with the PeriodDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupedPrometheusBillingListQuery) GetPeriodDurationOk() (*BillingPeriod, bool) {
	if o == nil || IsNil(o.PeriodDuration) {
		return nil, false
	}
	return o.PeriodDuration, true
}

// HasPeriodDuration returns a boolean if a field has been set.
func (o *GroupedPrometheusBillingListQuery) HasPeriodDuration() bool {
	if o != nil && !IsNil(o.PeriodDuration) {
		return true
	}

	return false
}

// SetPeriodDuration gets a reference to the given BillingPeriod and assigns it to the PeriodDuration field.
func (o *GroupedPrometheusBillingListQuery) SetPeriodDuration(v BillingPeriod) {
	o.PeriodDuration = &v
}

func (o GroupedPrometheusBillingListQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedPrometheusBillingListQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if !IsNil(o.PeriodDuration) {
		toSerialize["periodDuration"] = o.PeriodDuration
	}
	return toSerialize, nil
}

type NullableGroupedPrometheusBillingListQuery struct {
	value *GroupedPrometheusBillingListQuery
	isSet bool
}

func (v NullableGroupedPrometheusBillingListQuery) Get() *GroupedPrometheusBillingListQuery {
	return v.value
}

func (v *NullableGroupedPrometheusBillingListQuery) Set(val *GroupedPrometheusBillingListQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedPrometheusBillingListQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedPrometheusBillingListQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedPrometheusBillingListQuery(val *GroupedPrometheusBillingListQuery) *NullableGroupedPrometheusBillingListQuery {
	return &NullableGroupedPrometheusBillingListQuery{value: val, isSet: true}
}

func (v NullableGroupedPrometheusBillingListQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedPrometheusBillingListQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


