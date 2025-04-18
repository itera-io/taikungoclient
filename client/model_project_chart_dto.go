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

// checks if the ProjectChartDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectChartDto{}

// ProjectChartDto struct for ProjectChartDto
type ProjectChartDto struct {
	Succeeded []ProjectCommonRecordDto `json:"succeeded,omitempty"`
	Updating []ProjectCommonRecordDto `json:"updating,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
	Failed []ProjectCommonRecordDto `json:"failed,omitempty"`
	Purging []ProjectCommonRecordDto `json:"purging,omitempty"`
	Deleting []ProjectCommonRecordDto `json:"deleting,omitempty"`
	Importing []ProjectCommonRecordDto `json:"importing,omitempty"`
	FailedToImport []ProjectCommonRecordDto `json:"failedToImport,omitempty"`
}

// NewProjectChartDto instantiates a new ProjectChartDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectChartDto() *ProjectChartDto {
	this := ProjectChartDto{}
	return &this
}

// NewProjectChartDtoWithDefaults instantiates a new ProjectChartDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectChartDtoWithDefaults() *ProjectChartDto {
	this := ProjectChartDto{}
	return &this
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetSucceeded() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetSucceededOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.Succeeded) {
		return nil, false
	}
	return o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *ProjectChartDto) HasSucceeded() bool {
	if o != nil && !IsNil(o.Succeeded) {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given []ProjectCommonRecordDto and assigns it to the Succeeded field.
func (o *ProjectChartDto) SetSucceeded(v []ProjectCommonRecordDto) {
	o.Succeeded = v
}

// GetUpdating returns the Updating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetUpdating() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.Updating
}

// GetUpdatingOk returns a tuple with the Updating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetUpdatingOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.Updating) {
		return nil, false
	}
	return o.Updating, true
}

// HasUpdating returns a boolean if a field has been set.
func (o *ProjectChartDto) HasUpdating() bool {
	if o != nil && !IsNil(o.Updating) {
		return true
	}

	return false
}

// SetUpdating gets a reference to the given []ProjectCommonRecordDto and assigns it to the Updating field.
func (o *ProjectChartDto) SetUpdating(v []ProjectCommonRecordDto) {
	o.Updating = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProjectChartDto) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectChartDto) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProjectChartDto) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ProjectChartDto) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetFailed() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetFailedOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *ProjectChartDto) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given []ProjectCommonRecordDto and assigns it to the Failed field.
func (o *ProjectChartDto) SetFailed(v []ProjectCommonRecordDto) {
	o.Failed = v
}

// GetPurging returns the Purging field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetPurging() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.Purging
}

// GetPurgingOk returns a tuple with the Purging field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetPurgingOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.Purging) {
		return nil, false
	}
	return o.Purging, true
}

// HasPurging returns a boolean if a field has been set.
func (o *ProjectChartDto) HasPurging() bool {
	if o != nil && !IsNil(o.Purging) {
		return true
	}

	return false
}

// SetPurging gets a reference to the given []ProjectCommonRecordDto and assigns it to the Purging field.
func (o *ProjectChartDto) SetPurging(v []ProjectCommonRecordDto) {
	o.Purging = v
}

// GetDeleting returns the Deleting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetDeleting() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.Deleting
}

// GetDeletingOk returns a tuple with the Deleting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetDeletingOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.Deleting) {
		return nil, false
	}
	return o.Deleting, true
}

// HasDeleting returns a boolean if a field has been set.
func (o *ProjectChartDto) HasDeleting() bool {
	if o != nil && !IsNil(o.Deleting) {
		return true
	}

	return false
}

// SetDeleting gets a reference to the given []ProjectCommonRecordDto and assigns it to the Deleting field.
func (o *ProjectChartDto) SetDeleting(v []ProjectCommonRecordDto) {
	o.Deleting = v
}

// GetImporting returns the Importing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetImporting() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.Importing
}

// GetImportingOk returns a tuple with the Importing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetImportingOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.Importing) {
		return nil, false
	}
	return o.Importing, true
}

// HasImporting returns a boolean if a field has been set.
func (o *ProjectChartDto) HasImporting() bool {
	if o != nil && !IsNil(o.Importing) {
		return true
	}

	return false
}

// SetImporting gets a reference to the given []ProjectCommonRecordDto and assigns it to the Importing field.
func (o *ProjectChartDto) SetImporting(v []ProjectCommonRecordDto) {
	o.Importing = v
}

// GetFailedToImport returns the FailedToImport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectChartDto) GetFailedToImport() []ProjectCommonRecordDto {
	if o == nil {
		var ret []ProjectCommonRecordDto
		return ret
	}
	return o.FailedToImport
}

// GetFailedToImportOk returns a tuple with the FailedToImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectChartDto) GetFailedToImportOk() ([]ProjectCommonRecordDto, bool) {
	if o == nil || IsNil(o.FailedToImport) {
		return nil, false
	}
	return o.FailedToImport, true
}

// HasFailedToImport returns a boolean if a field has been set.
func (o *ProjectChartDto) HasFailedToImport() bool {
	if o != nil && !IsNil(o.FailedToImport) {
		return true
	}

	return false
}

// SetFailedToImport gets a reference to the given []ProjectCommonRecordDto and assigns it to the FailedToImport field.
func (o *ProjectChartDto) SetFailedToImport(v []ProjectCommonRecordDto) {
	o.FailedToImport = v
}

func (o ProjectChartDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectChartDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Succeeded != nil {
		toSerialize["succeeded"] = o.Succeeded
	}
	if o.Updating != nil {
		toSerialize["updating"] = o.Updating
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	if o.Purging != nil {
		toSerialize["purging"] = o.Purging
	}
	if o.Deleting != nil {
		toSerialize["deleting"] = o.Deleting
	}
	if o.Importing != nil {
		toSerialize["importing"] = o.Importing
	}
	if o.FailedToImport != nil {
		toSerialize["failedToImport"] = o.FailedToImport
	}
	return toSerialize, nil
}

type NullableProjectChartDto struct {
	value *ProjectChartDto
	isSet bool
}

func (v NullableProjectChartDto) Get() *ProjectChartDto {
	return v.value
}

func (v *NullableProjectChartDto) Set(val *ProjectChartDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectChartDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectChartDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectChartDto(val *ProjectChartDto) *NullableProjectChartDto {
	return &NullableProjectChartDto{value: val, isSet: true}
}

func (v NullableProjectChartDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectChartDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


