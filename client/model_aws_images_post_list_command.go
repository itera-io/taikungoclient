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

// checks if the AwsImagesPostListCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsImagesPostListCommand{}

// AwsImagesPostListCommand struct for AwsImagesPostListCommand
type AwsImagesPostListCommand struct {
	CloudId *int32 `json:"cloudId,omitempty"`
	Limit NullableInt32 `json:"limit,omitempty"`
	Offset NullableInt32 `json:"offset,omitempty"`
	SortBy NullableString `json:"sortBy,omitempty"`
	SortDirection NullableString `json:"sortDirection,omitempty"`
	Search NullableString `json:"search,omitempty"`
	Latest *bool `json:"latest,omitempty"`
	Owners []string `json:"owners,omitempty"`
	ProjectId NullableInt32 `json:"projectId,omitempty"`
}

// NewAwsImagesPostListCommand instantiates a new AwsImagesPostListCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsImagesPostListCommand() *AwsImagesPostListCommand {
	this := AwsImagesPostListCommand{}
	return &this
}

// NewAwsImagesPostListCommandWithDefaults instantiates a new AwsImagesPostListCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsImagesPostListCommandWithDefaults() *AwsImagesPostListCommand {
	this := AwsImagesPostListCommand{}
	return &this
}

// GetCloudId returns the CloudId field value if set, zero value otherwise.
func (o *AwsImagesPostListCommand) GetCloudId() int32 {
	if o == nil || IsNil(o.CloudId) {
		var ret int32
		return ret
	}
	return *o.CloudId
}

// GetCloudIdOk returns a tuple with the CloudId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsImagesPostListCommand) GetCloudIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CloudId) {
		return nil, false
	}
	return o.CloudId, true
}

// HasCloudId returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasCloudId() bool {
	if o != nil && !IsNil(o.CloudId) {
		return true
	}

	return false
}

// SetCloudId gets a reference to the given int32 and assigns it to the CloudId field.
func (o *AwsImagesPostListCommand) SetCloudId(v int32) {
	o.CloudId = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *AwsImagesPostListCommand) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *AwsImagesPostListCommand) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *AwsImagesPostListCommand) UnsetLimit() {
	o.Limit.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *AwsImagesPostListCommand) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *AwsImagesPostListCommand) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *AwsImagesPostListCommand) UnsetOffset() {
	o.Offset.Unset()
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetSortBy() string {
	if o == nil || IsNil(o.SortBy.Get()) {
		var ret string
		return ret
	}
	return *o.SortBy.Get()
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetSortByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortBy.Get(), o.SortBy.IsSet()
}

// HasSortBy returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasSortBy() bool {
	if o != nil && o.SortBy.IsSet() {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given NullableString and assigns it to the SortBy field.
func (o *AwsImagesPostListCommand) SetSortBy(v string) {
	o.SortBy.Set(&v)
}
// SetSortByNil sets the value for SortBy to be an explicit nil
func (o *AwsImagesPostListCommand) SetSortByNil() {
	o.SortBy.Set(nil)
}

// UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
func (o *AwsImagesPostListCommand) UnsetSortBy() {
	o.SortBy.Unset()
}

// GetSortDirection returns the SortDirection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetSortDirection() string {
	if o == nil || IsNil(o.SortDirection.Get()) {
		var ret string
		return ret
	}
	return *o.SortDirection.Get()
}

// GetSortDirectionOk returns a tuple with the SortDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetSortDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortDirection.Get(), o.SortDirection.IsSet()
}

// HasSortDirection returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasSortDirection() bool {
	if o != nil && o.SortDirection.IsSet() {
		return true
	}

	return false
}

// SetSortDirection gets a reference to the given NullableString and assigns it to the SortDirection field.
func (o *AwsImagesPostListCommand) SetSortDirection(v string) {
	o.SortDirection.Set(&v)
}
// SetSortDirectionNil sets the value for SortDirection to be an explicit nil
func (o *AwsImagesPostListCommand) SetSortDirectionNil() {
	o.SortDirection.Set(nil)
}

// UnsetSortDirection ensures that no value is present for SortDirection, not even an explicit nil
func (o *AwsImagesPostListCommand) UnsetSortDirection() {
	o.SortDirection.Unset()
}

// GetSearch returns the Search field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetSearch() string {
	if o == nil || IsNil(o.Search.Get()) {
		var ret string
		return ret
	}
	return *o.Search.Get()
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetSearchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Search.Get(), o.Search.IsSet()
}

// HasSearch returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasSearch() bool {
	if o != nil && o.Search.IsSet() {
		return true
	}

	return false
}

// SetSearch gets a reference to the given NullableString and assigns it to the Search field.
func (o *AwsImagesPostListCommand) SetSearch(v string) {
	o.Search.Set(&v)
}
// SetSearchNil sets the value for Search to be an explicit nil
func (o *AwsImagesPostListCommand) SetSearchNil() {
	o.Search.Set(nil)
}

// UnsetSearch ensures that no value is present for Search, not even an explicit nil
func (o *AwsImagesPostListCommand) UnsetSearch() {
	o.Search.Unset()
}

// GetLatest returns the Latest field value if set, zero value otherwise.
func (o *AwsImagesPostListCommand) GetLatest() bool {
	if o == nil || IsNil(o.Latest) {
		var ret bool
		return ret
	}
	return *o.Latest
}

// GetLatestOk returns a tuple with the Latest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsImagesPostListCommand) GetLatestOk() (*bool, bool) {
	if o == nil || IsNil(o.Latest) {
		return nil, false
	}
	return o.Latest, true
}

// HasLatest returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasLatest() bool {
	if o != nil && !IsNil(o.Latest) {
		return true
	}

	return false
}

// SetLatest gets a reference to the given bool and assigns it to the Latest field.
func (o *AwsImagesPostListCommand) SetLatest(v bool) {
	o.Latest = &v
}

// GetOwners returns the Owners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetOwners() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetOwnersOk() ([]string, bool) {
	if o == nil || IsNil(o.Owners) {
		return nil, false
	}
	return o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasOwners() bool {
	if o != nil && IsNil(o.Owners) {
		return true
	}

	return false
}

// SetOwners gets a reference to the given []string and assigns it to the Owners field.
func (o *AwsImagesPostListCommand) SetOwners(v []string) {
	o.Owners = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AwsImagesPostListCommand) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AwsImagesPostListCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *AwsImagesPostListCommand) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt32 and assigns it to the ProjectId field.
func (o *AwsImagesPostListCommand) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *AwsImagesPostListCommand) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *AwsImagesPostListCommand) UnsetProjectId() {
	o.ProjectId.Unset()
}

func (o AwsImagesPostListCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsImagesPostListCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudId) {
		toSerialize["cloudId"] = o.CloudId
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.SortBy.IsSet() {
		toSerialize["sortBy"] = o.SortBy.Get()
	}
	if o.SortDirection.IsSet() {
		toSerialize["sortDirection"] = o.SortDirection.Get()
	}
	if o.Search.IsSet() {
		toSerialize["search"] = o.Search.Get()
	}
	if !IsNil(o.Latest) {
		toSerialize["latest"] = o.Latest
	}
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	return toSerialize, nil
}

type NullableAwsImagesPostListCommand struct {
	value *AwsImagesPostListCommand
	isSet bool
}

func (v NullableAwsImagesPostListCommand) Get() *AwsImagesPostListCommand {
	return v.value
}

func (v *NullableAwsImagesPostListCommand) Set(val *AwsImagesPostListCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsImagesPostListCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsImagesPostListCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsImagesPostListCommand(val *AwsImagesPostListCommand) *NullableAwsImagesPostListCommand {
	return &NullableAwsImagesPostListCommand{value: val, isSet: true}
}

func (v NullableAwsImagesPostListCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsImagesPostListCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


