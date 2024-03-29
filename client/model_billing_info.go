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

// checks if the BillingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfo{}

// BillingInfo struct for BillingInfo
type BillingInfo struct {
	Data []BillingSummaryDto `json:"data,omitempty"`
	TotalTcu *float64 `json:"totalTcu,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewBillingInfo instantiates a new BillingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfo() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// NewBillingInfoWithDefaults instantiates a new BillingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoWithDefaults() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfo) GetData() []BillingSummaryDto {
	if o == nil {
		var ret []BillingSummaryDto
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfo) GetDataOk() ([]BillingSummaryDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BillingInfo) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BillingSummaryDto and assigns it to the Data field.
func (o *BillingInfo) SetData(v []BillingSummaryDto) {
	o.Data = v
}

// GetTotalTcu returns the TotalTcu field value if set, zero value otherwise.
func (o *BillingInfo) GetTotalTcu() float64 {
	if o == nil || IsNil(o.TotalTcu) {
		var ret float64
		return ret
	}
	return *o.TotalTcu
}

// GetTotalTcuOk returns a tuple with the TotalTcu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetTotalTcuOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalTcu) {
		return nil, false
	}
	return o.TotalTcu, true
}

// HasTotalTcu returns a boolean if a field has been set.
func (o *BillingInfo) HasTotalTcu() bool {
	if o != nil && !IsNil(o.TotalTcu) {
		return true
	}

	return false
}

// SetTotalTcu gets a reference to the given float64 and assigns it to the TotalTcu field.
func (o *BillingInfo) SetTotalTcu(v float64) {
	o.TotalTcu = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *BillingInfo) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *BillingInfo) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *BillingInfo) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o BillingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.TotalTcu) {
		toSerialize["totalTcu"] = o.TotalTcu
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableBillingInfo struct {
	value *BillingInfo
	isSet bool
}

func (v NullableBillingInfo) Get() *BillingInfo {
	return v.value
}

func (v *NullableBillingInfo) Set(val *BillingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfo(val *BillingInfo) *NullableBillingInfo {
	return &NullableBillingInfo{value: val, isSet: true}
}

func (v NullableBillingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


