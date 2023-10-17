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

// checks if the BillingSummaryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingSummaryDto{}

// BillingSummaryDto struct for BillingSummaryDto
type BillingSummaryDto struct {
	ProjectId   *int32         `json:"projectId,omitempty"`
	ProjectName NullableString `json:"projectName,omitempty"`
	StartDate   *time.Time     `json:"startDate,omitempty"`
	EndDate     NullableTime   `json:"endDate,omitempty"`
	Tcu         *float64       `json:"tcu,omitempty"`
	IsDeleted   *bool          `json:"isDeleted,omitempty"`
}

// NewBillingSummaryDto instantiates a new BillingSummaryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingSummaryDto() *BillingSummaryDto {
	this := BillingSummaryDto{}
	return &this
}

// NewBillingSummaryDtoWithDefaults instantiates a new BillingSummaryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingSummaryDtoWithDefaults() *BillingSummaryDto {
	this := BillingSummaryDto{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BillingSummaryDto) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int32
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSummaryDto) GetProjectIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BillingSummaryDto) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *BillingSummaryDto) SetProjectId(v int32) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingSummaryDto) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingSummaryDto) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *BillingSummaryDto) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *BillingSummaryDto) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *BillingSummaryDto) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *BillingSummaryDto) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingSummaryDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSummaryDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingSummaryDto) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingSummaryDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingSummaryDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingSummaryDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingSummaryDto) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableTime and assigns it to the EndDate field.
func (o *BillingSummaryDto) SetEndDate(v time.Time) {
	o.EndDate.Set(&v)
}

// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *BillingSummaryDto) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *BillingSummaryDto) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetTcu returns the Tcu field value if set, zero value otherwise.
func (o *BillingSummaryDto) GetTcu() float64 {
	if o == nil || IsNil(o.Tcu) {
		var ret float64
		return ret
	}
	return *o.Tcu
}

// GetTcuOk returns a tuple with the Tcu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSummaryDto) GetTcuOk() (*float64, bool) {
	if o == nil || IsNil(o.Tcu) {
		return nil, false
	}
	return o.Tcu, true
}

// HasTcu returns a boolean if a field has been set.
func (o *BillingSummaryDto) HasTcu() bool {
	if o != nil && !IsNil(o.Tcu) {
		return true
	}

	return false
}

// SetTcu gets a reference to the given float64 and assigns it to the Tcu field.
func (o *BillingSummaryDto) SetTcu(v float64) {
	o.Tcu = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *BillingSummaryDto) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSummaryDto) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *BillingSummaryDto) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *BillingSummaryDto) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o BillingSummaryDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingSummaryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate.IsSet() {
		toSerialize["endDate"] = o.EndDate.Get()
	}
	if !IsNil(o.Tcu) {
		toSerialize["tcu"] = o.Tcu
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableBillingSummaryDto struct {
	value *BillingSummaryDto
	isSet bool
}

func (v NullableBillingSummaryDto) Get() *BillingSummaryDto {
	return v.value
}

func (v *NullableBillingSummaryDto) Set(val *BillingSummaryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingSummaryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingSummaryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingSummaryDto(val *BillingSummaryDto) *NullableBillingSummaryDto {
	return &NullableBillingSummaryDto{value: val, isSet: true}
}

func (v NullableBillingSummaryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingSummaryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
