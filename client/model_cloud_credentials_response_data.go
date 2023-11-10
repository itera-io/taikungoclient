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

// checks if the CloudCredentialsResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudCredentialsResponseData{}

// CloudCredentialsResponseData struct for CloudCredentialsResponseData
type CloudCredentialsResponseData struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	OrganizationName NullableString `json:"organizationName,omitempty"`
	CloudType NullableString `json:"cloudType,omitempty"`
}

// NewCloudCredentialsResponseData instantiates a new CloudCredentialsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCredentialsResponseData() *CloudCredentialsResponseData {
	this := CloudCredentialsResponseData{}
	return &this
}

// NewCloudCredentialsResponseDataWithDefaults instantiates a new CloudCredentialsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCredentialsResponseDataWithDefaults() *CloudCredentialsResponseData {
	this := CloudCredentialsResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudCredentialsResponseData) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCredentialsResponseData) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudCredentialsResponseData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CloudCredentialsResponseData) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsResponseData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CloudCredentialsResponseData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CloudCredentialsResponseData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CloudCredentialsResponseData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CloudCredentialsResponseData) UnsetName() {
	o.Name.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsResponseData) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsResponseData) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CloudCredentialsResponseData) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CloudCredentialsResponseData) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CloudCredentialsResponseData) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CloudCredentialsResponseData) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsResponseData) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName.Get()) {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsResponseData) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *CloudCredentialsResponseData) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *CloudCredentialsResponseData) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *CloudCredentialsResponseData) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *CloudCredentialsResponseData) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetCloudType returns the CloudType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudCredentialsResponseData) GetCloudType() string {
	if o == nil || IsNil(o.CloudType.Get()) {
		var ret string
		return ret
	}
	return *o.CloudType.Get()
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudCredentialsResponseData) GetCloudTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudType.Get(), o.CloudType.IsSet()
}

// HasCloudType returns a boolean if a field has been set.
func (o *CloudCredentialsResponseData) HasCloudType() bool {
	if o != nil && o.CloudType.IsSet() {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given NullableString and assigns it to the CloudType field.
func (o *CloudCredentialsResponseData) SetCloudType(v string) {
	o.CloudType.Set(&v)
}
// SetCloudTypeNil sets the value for CloudType to be an explicit nil
func (o *CloudCredentialsResponseData) SetCloudTypeNil() {
	o.CloudType.Set(nil)
}

// UnsetCloudType ensures that no value is present for CloudType, not even an explicit nil
func (o *CloudCredentialsResponseData) UnsetCloudType() {
	o.CloudType.Unset()
}

func (o CloudCredentialsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudCredentialsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.CloudType.IsSet() {
		toSerialize["cloudType"] = o.CloudType.Get()
	}
	return toSerialize, nil
}

type NullableCloudCredentialsResponseData struct {
	value *CloudCredentialsResponseData
	isSet bool
}

func (v NullableCloudCredentialsResponseData) Get() *CloudCredentialsResponseData {
	return v.value
}

func (v *NullableCloudCredentialsResponseData) Set(val *CloudCredentialsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCredentialsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCredentialsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCredentialsResponseData(val *CloudCredentialsResponseData) *NullableCloudCredentialsResponseData {
	return &NullableCloudCredentialsResponseData{value: val, isSet: true}
}

func (v NullableCloudCredentialsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCredentialsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


