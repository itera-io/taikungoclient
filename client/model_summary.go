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

// checks if the Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Summary{}

// Summary struct for Summary
type Summary struct {
	TotalDetectedResources *int32 `json:"totalDetectedResources,omitempty"`
	TotalSupportedResources *int32 `json:"totalSupportedResources,omitempty"`
	TotalUnsupportedResources *int32 `json:"totalUnsupportedResources,omitempty"`
	TotalUsageBasedResources *int32 `json:"totalUsageBasedResources,omitempty"`
	TotalNoPriceResources *int32 `json:"totalNoPriceResources,omitempty"`
	UnsupportedResourceCounts map[string]int32 `json:"unsupportedResourceCounts,omitempty"`
	NoPriceResourceCounts map[string]int32 `json:"noPriceResourceCounts,omitempty"`
}

// NewSummary instantiates a new Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummary() *Summary {
	this := Summary{}
	return &this
}

// NewSummaryWithDefaults instantiates a new Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryWithDefaults() *Summary {
	this := Summary{}
	return &this
}

// GetTotalDetectedResources returns the TotalDetectedResources field value if set, zero value otherwise.
func (o *Summary) GetTotalDetectedResources() int32 {
	if o == nil || IsNil(o.TotalDetectedResources) {
		var ret int32
		return ret
	}
	return *o.TotalDetectedResources
}

// GetTotalDetectedResourcesOk returns a tuple with the TotalDetectedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetTotalDetectedResourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalDetectedResources) {
		return nil, false
	}
	return o.TotalDetectedResources, true
}

// HasTotalDetectedResources returns a boolean if a field has been set.
func (o *Summary) HasTotalDetectedResources() bool {
	if o != nil && !IsNil(o.TotalDetectedResources) {
		return true
	}

	return false
}

// SetTotalDetectedResources gets a reference to the given int32 and assigns it to the TotalDetectedResources field.
func (o *Summary) SetTotalDetectedResources(v int32) {
	o.TotalDetectedResources = &v
}

// GetTotalSupportedResources returns the TotalSupportedResources field value if set, zero value otherwise.
func (o *Summary) GetTotalSupportedResources() int32 {
	if o == nil || IsNil(o.TotalSupportedResources) {
		var ret int32
		return ret
	}
	return *o.TotalSupportedResources
}

// GetTotalSupportedResourcesOk returns a tuple with the TotalSupportedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetTotalSupportedResourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSupportedResources) {
		return nil, false
	}
	return o.TotalSupportedResources, true
}

// HasTotalSupportedResources returns a boolean if a field has been set.
func (o *Summary) HasTotalSupportedResources() bool {
	if o != nil && !IsNil(o.TotalSupportedResources) {
		return true
	}

	return false
}

// SetTotalSupportedResources gets a reference to the given int32 and assigns it to the TotalSupportedResources field.
func (o *Summary) SetTotalSupportedResources(v int32) {
	o.TotalSupportedResources = &v
}

// GetTotalUnsupportedResources returns the TotalUnsupportedResources field value if set, zero value otherwise.
func (o *Summary) GetTotalUnsupportedResources() int32 {
	if o == nil || IsNil(o.TotalUnsupportedResources) {
		var ret int32
		return ret
	}
	return *o.TotalUnsupportedResources
}

// GetTotalUnsupportedResourcesOk returns a tuple with the TotalUnsupportedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetTotalUnsupportedResourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalUnsupportedResources) {
		return nil, false
	}
	return o.TotalUnsupportedResources, true
}

// HasTotalUnsupportedResources returns a boolean if a field has been set.
func (o *Summary) HasTotalUnsupportedResources() bool {
	if o != nil && !IsNil(o.TotalUnsupportedResources) {
		return true
	}

	return false
}

// SetTotalUnsupportedResources gets a reference to the given int32 and assigns it to the TotalUnsupportedResources field.
func (o *Summary) SetTotalUnsupportedResources(v int32) {
	o.TotalUnsupportedResources = &v
}

// GetTotalUsageBasedResources returns the TotalUsageBasedResources field value if set, zero value otherwise.
func (o *Summary) GetTotalUsageBasedResources() int32 {
	if o == nil || IsNil(o.TotalUsageBasedResources) {
		var ret int32
		return ret
	}
	return *o.TotalUsageBasedResources
}

// GetTotalUsageBasedResourcesOk returns a tuple with the TotalUsageBasedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetTotalUsageBasedResourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalUsageBasedResources) {
		return nil, false
	}
	return o.TotalUsageBasedResources, true
}

// HasTotalUsageBasedResources returns a boolean if a field has been set.
func (o *Summary) HasTotalUsageBasedResources() bool {
	if o != nil && !IsNil(o.TotalUsageBasedResources) {
		return true
	}

	return false
}

// SetTotalUsageBasedResources gets a reference to the given int32 and assigns it to the TotalUsageBasedResources field.
func (o *Summary) SetTotalUsageBasedResources(v int32) {
	o.TotalUsageBasedResources = &v
}

// GetTotalNoPriceResources returns the TotalNoPriceResources field value if set, zero value otherwise.
func (o *Summary) GetTotalNoPriceResources() int32 {
	if o == nil || IsNil(o.TotalNoPriceResources) {
		var ret int32
		return ret
	}
	return *o.TotalNoPriceResources
}

// GetTotalNoPriceResourcesOk returns a tuple with the TotalNoPriceResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetTotalNoPriceResourcesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalNoPriceResources) {
		return nil, false
	}
	return o.TotalNoPriceResources, true
}

// HasTotalNoPriceResources returns a boolean if a field has been set.
func (o *Summary) HasTotalNoPriceResources() bool {
	if o != nil && !IsNil(o.TotalNoPriceResources) {
		return true
	}

	return false
}

// SetTotalNoPriceResources gets a reference to the given int32 and assigns it to the TotalNoPriceResources field.
func (o *Summary) SetTotalNoPriceResources(v int32) {
	o.TotalNoPriceResources = &v
}

// GetUnsupportedResourceCounts returns the UnsupportedResourceCounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Summary) GetUnsupportedResourceCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.UnsupportedResourceCounts
}

// GetUnsupportedResourceCountsOk returns a tuple with the UnsupportedResourceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Summary) GetUnsupportedResourceCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.UnsupportedResourceCounts) {
		return nil, false
	}
	return &o.UnsupportedResourceCounts, true
}

// HasUnsupportedResourceCounts returns a boolean if a field has been set.
func (o *Summary) HasUnsupportedResourceCounts() bool {
	if o != nil && !IsNil(o.UnsupportedResourceCounts) {
		return true
	}

	return false
}

// SetUnsupportedResourceCounts gets a reference to the given map[string]int32 and assigns it to the UnsupportedResourceCounts field.
func (o *Summary) SetUnsupportedResourceCounts(v map[string]int32) {
	o.UnsupportedResourceCounts = v
}

// GetNoPriceResourceCounts returns the NoPriceResourceCounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Summary) GetNoPriceResourceCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.NoPriceResourceCounts
}

// GetNoPriceResourceCountsOk returns a tuple with the NoPriceResourceCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Summary) GetNoPriceResourceCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.NoPriceResourceCounts) {
		return nil, false
	}
	return &o.NoPriceResourceCounts, true
}

// HasNoPriceResourceCounts returns a boolean if a field has been set.
func (o *Summary) HasNoPriceResourceCounts() bool {
	if o != nil && !IsNil(o.NoPriceResourceCounts) {
		return true
	}

	return false
}

// SetNoPriceResourceCounts gets a reference to the given map[string]int32 and assigns it to the NoPriceResourceCounts field.
func (o *Summary) SetNoPriceResourceCounts(v map[string]int32) {
	o.NoPriceResourceCounts = v
}

func (o Summary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Summary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalDetectedResources) {
		toSerialize["totalDetectedResources"] = o.TotalDetectedResources
	}
	if !IsNil(o.TotalSupportedResources) {
		toSerialize["totalSupportedResources"] = o.TotalSupportedResources
	}
	if !IsNil(o.TotalUnsupportedResources) {
		toSerialize["totalUnsupportedResources"] = o.TotalUnsupportedResources
	}
	if !IsNil(o.TotalUsageBasedResources) {
		toSerialize["totalUsageBasedResources"] = o.TotalUsageBasedResources
	}
	if !IsNil(o.TotalNoPriceResources) {
		toSerialize["totalNoPriceResources"] = o.TotalNoPriceResources
	}
	if o.UnsupportedResourceCounts != nil {
		toSerialize["unsupportedResourceCounts"] = o.UnsupportedResourceCounts
	}
	if o.NoPriceResourceCounts != nil {
		toSerialize["noPriceResourceCounts"] = o.NoPriceResourceCounts
	}
	return toSerialize, nil
}

type NullableSummary struct {
	value *Summary
	isSet bool
}

func (v NullableSummary) Get() *Summary {
	return v.value
}

func (v *NullableSummary) Set(val *Summary) {
	v.value = val
	v.isSet = true
}

func (v NullableSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummary(val *Summary) *NullableSummary {
	return &NullableSummary{value: val, isSet: true}
}

func (v NullableSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


