/*
Taikun - Showback API

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the GroupedCredentialForShowbackSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupedCredentialForShowbackSummary{}

// GroupedCredentialForShowbackSummary struct for GroupedCredentialForShowbackSummary
type GroupedCredentialForShowbackSummary struct {
	Id        NullableInt32  `json:"id,omitempty"`
	ProjectId NullableInt32  `json:"projectId,omitempty"`
	Name      NullableString `json:"name,omitempty"`
}

// NewGroupedCredentialForShowbackSummary instantiates a new GroupedCredentialForShowbackSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupedCredentialForShowbackSummary() *GroupedCredentialForShowbackSummary {
	this := GroupedCredentialForShowbackSummary{}
	return &this
}

// NewGroupedCredentialForShowbackSummaryWithDefaults instantiates a new GroupedCredentialForShowbackSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupedCredentialForShowbackSummaryWithDefaults() *GroupedCredentialForShowbackSummary {
	this := GroupedCredentialForShowbackSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedCredentialForShowbackSummary) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedCredentialForShowbackSummary) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GroupedCredentialForShowbackSummary) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GroupedCredentialForShowbackSummary) SetId(v int32) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *GroupedCredentialForShowbackSummary) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GroupedCredentialForShowbackSummary) UnsetId() {
	o.Id.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedCredentialForShowbackSummary) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedCredentialForShowbackSummary) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *GroupedCredentialForShowbackSummary) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt32 and assigns it to the ProjectId field.
func (o *GroupedCredentialForShowbackSummary) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}

// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *GroupedCredentialForShowbackSummary) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *GroupedCredentialForShowbackSummary) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupedCredentialForShowbackSummary) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupedCredentialForShowbackSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GroupedCredentialForShowbackSummary) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GroupedCredentialForShowbackSummary) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *GroupedCredentialForShowbackSummary) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GroupedCredentialForShowbackSummary) UnsetName() {
	o.Name.Unset()
}

func (o GroupedCredentialForShowbackSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupedCredentialForShowbackSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ProjectId.IsSet() {
		toSerialize["projectId"] = o.ProjectId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableGroupedCredentialForShowbackSummary struct {
	value *GroupedCredentialForShowbackSummary
	isSet bool
}

func (v NullableGroupedCredentialForShowbackSummary) Get() *GroupedCredentialForShowbackSummary {
	return v.value
}

func (v *NullableGroupedCredentialForShowbackSummary) Set(val *GroupedCredentialForShowbackSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupedCredentialForShowbackSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupedCredentialForShowbackSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupedCredentialForShowbackSummary(val *GroupedCredentialForShowbackSummary) *NullableGroupedCredentialForShowbackSummary {
	return &NullableGroupedCredentialForShowbackSummary{value: val, isSet: true}
}

func (v NullableGroupedCredentialForShowbackSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupedCredentialForShowbackSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}