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
	"time"
)

// checks if the LokiResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LokiResponseDto{}

// LokiResponseDto struct for LokiResponseDto
type LokiResponseDto struct {
	ProjectId   *int32         `json:"projectId,omitempty"`
	Parameters  []Parameter    `json:"parameters,omitempty"`
	Filters     []Filter       `json:"filters,omitempty"`
	Start       NullableTime   `json:"start,omitempty"`
	End         NullableTime   `json:"end,omitempty"`
	Limit       NullableInt32  `json:"limit,omitempty"`
	Direction   NullableString `json:"direction,omitempty"`
	CanDownload *bool          `json:"canDownload,omitempty"`
}

// NewLokiResponseDto instantiates a new LokiResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLokiResponseDto() *LokiResponseDto {
	this := LokiResponseDto{}
	return &this
}

// NewLokiResponseDtoWithDefaults instantiates a new LokiResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLokiResponseDtoWithDefaults() *LokiResponseDto {
	this := LokiResponseDto{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *LokiResponseDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LokiResponseDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *LokiResponseDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *LokiResponseDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LokiResponseDto) GetParameters() []Parameter {
	if o == nil {
		var ret []Parameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LokiResponseDto) GetParametersOk() ([]Parameter, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *LokiResponseDto) HasParameters() bool {
	if o != nil && IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []Parameter and assigns it to the Parameters field.
func (o *LokiResponseDto) SetParameters(v []Parameter) {
	o.Parameters = v
}

// GetFilters returns the Filters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LokiResponseDto) GetFilters() []Filter {
	if o == nil {
		var ret []Filter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LokiResponseDto) GetFiltersOk() ([]Filter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *LokiResponseDto) HasFilters() bool {
	if o != nil && IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []Filter and assigns it to the Filters field.
func (o *LokiResponseDto) SetFilters(v []Filter) {
	o.Filters = v
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LokiResponseDto) GetStart() time.Time {
	if o == nil || IsNil(o.Start.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LokiResponseDto) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *LokiResponseDto) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableTime and assigns it to the Start field.
func (o *LokiResponseDto) SetStart(v time.Time) {
	o.Start.Set(&v)
}

// SetStartNil sets the value for Start to be an explicit nil
func (o *LokiResponseDto) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *LokiResponseDto) UnsetStart() {
	o.Start.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LokiResponseDto) GetEnd() time.Time {
	if o == nil || IsNil(o.End.Get()) {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LokiResponseDto) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *LokiResponseDto) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableTime and assigns it to the End field.
func (o *LokiResponseDto) SetEnd(v time.Time) {
	o.End.Set(&v)
}

// SetEndNil sets the value for End to be an explicit nil
func (o *LokiResponseDto) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *LokiResponseDto) UnsetEnd() {
	o.End.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LokiResponseDto) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LokiResponseDto) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *LokiResponseDto) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *LokiResponseDto) SetLimit(v int32) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil
func (o *LokiResponseDto) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *LokiResponseDto) UnsetLimit() {
	o.Limit.Unset()
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LokiResponseDto) GetDirection() string {
	if o == nil || IsNil(o.Direction.Get()) {
		var ret string
		return ret
	}
	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LokiResponseDto) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// HasDirection returns a boolean if a field has been set.
func (o *LokiResponseDto) HasDirection() bool {
	if o != nil && o.Direction.IsSet() {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NullableString and assigns it to the Direction field.
func (o *LokiResponseDto) SetDirection(v string) {
	o.Direction.Set(&v)
}

// SetDirectionNil sets the value for Direction to be an explicit nil
func (o *LokiResponseDto) SetDirectionNil() {
	o.Direction.Set(nil)
}

// UnsetDirection ensures that no value is present for Direction, not even an explicit nil
func (o *LokiResponseDto) UnsetDirection() {
	o.Direction.Unset()
}

// GetCanDownload returns the CanDownload field value if set, zero value otherwise.
func (o *LokiResponseDto) GetCanDownload() bool {
	if o == nil || IsNil(o.CanDownload) {
		var ret bool
		return ret
	}
	return *o.CanDownload
}

// GetCanDownloadOk returns a tuple with the CanDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LokiResponseDto) GetCanDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.CanDownload) {
		return nil, false
	}
	return o.CanDownload, true
}

// HasCanDownload returns a boolean if a field has been set.
func (o *LokiResponseDto) HasCanDownload() bool {
	if o != nil && !IsNil(o.CanDownload) {
		return true
	}

	return false
}

// SetCanDownload gets a reference to the given bool and assigns it to the CanDownload field.
func (o *LokiResponseDto) SetCanDownload(v bool) {
	o.CanDownload = &v
}

func (o LokiResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LokiResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Direction.IsSet() {
		toSerialize["direction"] = o.Direction.Get()
	}
	if !IsNil(o.CanDownload) {
		toSerialize["canDownload"] = o.CanDownload
	}
	return toSerialize, nil
}

type NullableLokiResponseDto struct {
	value *LokiResponseDto
	isSet bool
}

func (v NullableLokiResponseDto) Get() *LokiResponseDto {
	return v.value
}

func (v *NullableLokiResponseDto) Set(val *LokiResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLokiResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLokiResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLokiResponseDto(val *LokiResponseDto) *NullableLokiResponseDto {
	return &NullableLokiResponseDto{value: val, isSet: true}
}

func (v NullableLokiResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLokiResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
