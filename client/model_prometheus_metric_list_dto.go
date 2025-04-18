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

// checks if the PrometheusMetricListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrometheusMetricListDto{}

// PrometheusMetricListDto struct for PrometheusMetricListDto
type PrometheusMetricListDto struct {
	Status NullableString `json:"status,omitempty"`
	Data *MetricData `json:"data,omitempty"`
}

// NewPrometheusMetricListDto instantiates a new PrometheusMetricListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusMetricListDto() *PrometheusMetricListDto {
	this := PrometheusMetricListDto{}
	return &this
}

// NewPrometheusMetricListDtoWithDefaults instantiates a new PrometheusMetricListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusMetricListDtoWithDefaults() *PrometheusMetricListDto {
	this := PrometheusMetricListDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrometheusMetricListDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrometheusMetricListDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PrometheusMetricListDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PrometheusMetricListDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PrometheusMetricListDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PrometheusMetricListDto) UnsetStatus() {
	o.Status.Unset()
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PrometheusMetricListDto) GetData() MetricData {
	if o == nil || IsNil(o.Data) {
		var ret MetricData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusMetricListDto) GetDataOk() (*MetricData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PrometheusMetricListDto) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MetricData and assigns it to the Data field.
func (o *PrometheusMetricListDto) SetData(v MetricData) {
	o.Data = &v
}

func (o PrometheusMetricListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrometheusMetricListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePrometheusMetricListDto struct {
	value *PrometheusMetricListDto
	isSet bool
}

func (v NullablePrometheusMetricListDto) Get() *PrometheusMetricListDto {
	return v.value
}

func (v *NullablePrometheusMetricListDto) Set(val *PrometheusMetricListDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheusMetricListDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheusMetricListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheusMetricListDto(val *PrometheusMetricListDto) *NullablePrometheusMetricListDto {
	return &NullablePrometheusMetricListDto{value: val, isSet: true}
}

func (v NullablePrometheusMetricListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheusMetricListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


