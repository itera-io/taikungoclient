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

// checks if the AzureQuotaListRecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureQuotaListRecordDto{}

// AzureQuotaListRecordDto struct for AzureQuotaListRecordDto
type AzureQuotaListRecordDto struct {
	TotalCores *int64 `json:"totalCores,omitempty"`
	CurrentUsage *int32 `json:"currentUsage,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewAzureQuotaListRecordDto instantiates a new AzureQuotaListRecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureQuotaListRecordDto() *AzureQuotaListRecordDto {
	this := AzureQuotaListRecordDto{}
	return &this
}

// NewAzureQuotaListRecordDtoWithDefaults instantiates a new AzureQuotaListRecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureQuotaListRecordDtoWithDefaults() *AzureQuotaListRecordDto {
	this := AzureQuotaListRecordDto{}
	return &this
}

// GetTotalCores returns the TotalCores field value if set, zero value otherwise.
func (o *AzureQuotaListRecordDto) GetTotalCores() int64 {
	if o == nil || IsNil(o.TotalCores) {
		var ret int64
		return ret
	}
	return *o.TotalCores
}

// GetTotalCoresOk returns a tuple with the TotalCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureQuotaListRecordDto) GetTotalCoresOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCores) {
		return nil, false
	}
	return o.TotalCores, true
}

// HasTotalCores returns a boolean if a field has been set.
func (o *AzureQuotaListRecordDto) HasTotalCores() bool {
	if o != nil && !IsNil(o.TotalCores) {
		return true
	}

	return false
}

// SetTotalCores gets a reference to the given int64 and assigns it to the TotalCores field.
func (o *AzureQuotaListRecordDto) SetTotalCores(v int64) {
	o.TotalCores = &v
}

// GetCurrentUsage returns the CurrentUsage field value if set, zero value otherwise.
func (o *AzureQuotaListRecordDto) GetCurrentUsage() int32 {
	if o == nil || IsNil(o.CurrentUsage) {
		var ret int32
		return ret
	}
	return *o.CurrentUsage
}

// GetCurrentUsageOk returns a tuple with the CurrentUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureQuotaListRecordDto) GetCurrentUsageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentUsage) {
		return nil, false
	}
	return o.CurrentUsage, true
}

// HasCurrentUsage returns a boolean if a field has been set.
func (o *AzureQuotaListRecordDto) HasCurrentUsage() bool {
	if o != nil && !IsNil(o.CurrentUsage) {
		return true
	}

	return false
}

// SetCurrentUsage gets a reference to the given int32 and assigns it to the CurrentUsage field.
func (o *AzureQuotaListRecordDto) SetCurrentUsage(v int32) {
	o.CurrentUsage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AzureQuotaListRecordDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureQuotaListRecordDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AzureQuotaListRecordDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AzureQuotaListRecordDto) SetName(v string) {
	o.Name = &v
}

func (o AzureQuotaListRecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureQuotaListRecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCores) {
		toSerialize["totalCores"] = o.TotalCores
	}
	if !IsNil(o.CurrentUsage) {
		toSerialize["currentUsage"] = o.CurrentUsage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAzureQuotaListRecordDto struct {
	value *AzureQuotaListRecordDto
	isSet bool
}

func (v NullableAzureQuotaListRecordDto) Get() *AzureQuotaListRecordDto {
	return v.value
}

func (v *NullableAzureQuotaListRecordDto) Set(val *AzureQuotaListRecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureQuotaListRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureQuotaListRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureQuotaListRecordDto(val *AzureQuotaListRecordDto) *NullableAzureQuotaListRecordDto {
	return &NullableAzureQuotaListRecordDto{value: val, isSet: true}
}

func (v NullableAzureQuotaListRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureQuotaListRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


