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

// checks if the IngressSearchCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngressSearchCommand{}

// IngressSearchCommand struct for IngressSearchCommand
type IngressSearchCommand struct {
	Limit NullableInt32 `json:"limit,omitempty"`
	Offset NullableInt32 `json:"offset,omitempty"`
	SearchTerm NullableString `json:"searchTerm,omitempty"`
	IncludePublicImportedClusters NullableBool `json:"includePublicImportedClusters,omitempty"`
}

// NewIngressSearchCommand instantiates a new IngressSearchCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngressSearchCommand() *IngressSearchCommand {
	this := IngressSearchCommand{}
	return &this
}

// NewIngressSearchCommandWithDefaults instantiates a new IngressSearchCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngressSearchCommandWithDefaults() *IngressSearchCommand {
	this := IngressSearchCommand{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngressSearchCommand) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngressSearchCommand) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *IngressSearchCommand) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *IngressSearchCommand) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *IngressSearchCommand) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *IngressSearchCommand) UnsetLimit() {
	o.Limit.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngressSearchCommand) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngressSearchCommand) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *IngressSearchCommand) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *IngressSearchCommand) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *IngressSearchCommand) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *IngressSearchCommand) UnsetOffset() {
	o.Offset.Unset()
}

// GetSearchTerm returns the SearchTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngressSearchCommand) GetSearchTerm() string {
	if o == nil || IsNil(o.SearchTerm.Get()) {
		var ret string
		return ret
	}
	return *o.SearchTerm.Get()
}

// GetSearchTermOk returns a tuple with the SearchTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngressSearchCommand) GetSearchTermOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchTerm.Get(), o.SearchTerm.IsSet()
}

// HasSearchTerm returns a boolean if a field has been set.
func (o *IngressSearchCommand) HasSearchTerm() bool {
	if o != nil && o.SearchTerm.IsSet() {
		return true
	}

	return false
}

// SetSearchTerm gets a reference to the given NullableString and assigns it to the SearchTerm field.
func (o *IngressSearchCommand) SetSearchTerm(v string) {
	o.SearchTerm.Set(&v)
}
// SetSearchTermNil sets the value for SearchTerm to be an explicit nil
func (o *IngressSearchCommand) SetSearchTermNil() {
	o.SearchTerm.Set(nil)
}

// UnsetSearchTerm ensures that no value is present for SearchTerm, not even an explicit nil
func (o *IngressSearchCommand) UnsetSearchTerm() {
	o.SearchTerm.Unset()
}

// GetIncludePublicImportedClusters returns the IncludePublicImportedClusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngressSearchCommand) GetIncludePublicImportedClusters() bool {
	if o == nil || IsNil(o.IncludePublicImportedClusters.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludePublicImportedClusters.Get()
}

// GetIncludePublicImportedClustersOk returns a tuple with the IncludePublicImportedClusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngressSearchCommand) GetIncludePublicImportedClustersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludePublicImportedClusters.Get(), o.IncludePublicImportedClusters.IsSet()
}

// HasIncludePublicImportedClusters returns a boolean if a field has been set.
func (o *IngressSearchCommand) HasIncludePublicImportedClusters() bool {
	if o != nil && o.IncludePublicImportedClusters.IsSet() {
		return true
	}

	return false
}

// SetIncludePublicImportedClusters gets a reference to the given NullableBool and assigns it to the IncludePublicImportedClusters field.
func (o *IngressSearchCommand) SetIncludePublicImportedClusters(v bool) {
	o.IncludePublicImportedClusters.Set(&v)
}
// SetIncludePublicImportedClustersNil sets the value for IncludePublicImportedClusters to be an explicit nil
func (o *IngressSearchCommand) SetIncludePublicImportedClustersNil() {
	o.IncludePublicImportedClusters.Set(nil)
}

// UnsetIncludePublicImportedClusters ensures that no value is present for IncludePublicImportedClusters, not even an explicit nil
func (o *IngressSearchCommand) UnsetIncludePublicImportedClusters() {
	o.IncludePublicImportedClusters.Unset()
}

func (o IngressSearchCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngressSearchCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.SearchTerm.IsSet() {
		toSerialize["searchTerm"] = o.SearchTerm.Get()
	}
	if o.IncludePublicImportedClusters.IsSet() {
		toSerialize["includePublicImportedClusters"] = o.IncludePublicImportedClusters.Get()
	}
	return toSerialize, nil
}

type NullableIngressSearchCommand struct {
	value *IngressSearchCommand
	isSet bool
}

func (v NullableIngressSearchCommand) Get() *IngressSearchCommand {
	return v.value
}

func (v *NullableIngressSearchCommand) Set(val *IngressSearchCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableIngressSearchCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableIngressSearchCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngressSearchCommand(val *IngressSearchCommand) *NullableIngressSearchCommand {
	return &NullableIngressSearchCommand{value: val, isSet: true}
}

func (v NullableIngressSearchCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngressSearchCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


