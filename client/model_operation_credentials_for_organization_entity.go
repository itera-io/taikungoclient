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

// checks if the OperationCredentialsForOrganizationEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationCredentialsForOrganizationEntity{}

// OperationCredentialsForOrganizationEntity struct for OperationCredentialsForOrganizationEntity
type OperationCredentialsForOrganizationEntity struct {
	OperationCredentialId NullableInt32 `json:"operationCredentialId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
}

// NewOperationCredentialsForOrganizationEntity instantiates a new OperationCredentialsForOrganizationEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationCredentialsForOrganizationEntity() *OperationCredentialsForOrganizationEntity {
	this := OperationCredentialsForOrganizationEntity{}
	return &this
}

// NewOperationCredentialsForOrganizationEntityWithDefaults instantiates a new OperationCredentialsForOrganizationEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationCredentialsForOrganizationEntityWithDefaults() *OperationCredentialsForOrganizationEntity {
	this := OperationCredentialsForOrganizationEntity{}
	return &this
}

// GetOperationCredentialId returns the OperationCredentialId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsForOrganizationEntity) GetOperationCredentialId() int32 {
	if o == nil || IsNil(o.OperationCredentialId.Get()) {
		var ret int32
		return ret
	}
	return *o.OperationCredentialId.Get()
}

// GetOperationCredentialIdOk returns a tuple with the OperationCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsForOrganizationEntity) GetOperationCredentialIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationCredentialId.Get(), o.OperationCredentialId.IsSet()
}

// HasOperationCredentialId returns a boolean if a field has been set.
func (o *OperationCredentialsForOrganizationEntity) HasOperationCredentialId() bool {
	if o != nil && o.OperationCredentialId.IsSet() {
		return true
	}

	return false
}

// SetOperationCredentialId gets a reference to the given NullableInt32 and assigns it to the OperationCredentialId field.
func (o *OperationCredentialsForOrganizationEntity) SetOperationCredentialId(v int32) {
	o.OperationCredentialId.Set(&v)
}
// SetOperationCredentialIdNil sets the value for OperationCredentialId to be an explicit nil
func (o *OperationCredentialsForOrganizationEntity) SetOperationCredentialIdNil() {
	o.OperationCredentialId.Set(nil)
}

// UnsetOperationCredentialId ensures that no value is present for OperationCredentialId, not even an explicit nil
func (o *OperationCredentialsForOrganizationEntity) UnsetOperationCredentialId() {
	o.OperationCredentialId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OperationCredentialsForOrganizationEntity) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OperationCredentialsForOrganizationEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OperationCredentialsForOrganizationEntity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OperationCredentialsForOrganizationEntity) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OperationCredentialsForOrganizationEntity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OperationCredentialsForOrganizationEntity) UnsetName() {
	o.Name.Unset()
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *OperationCredentialsForOrganizationEntity) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationCredentialsForOrganizationEntity) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *OperationCredentialsForOrganizationEntity) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *OperationCredentialsForOrganizationEntity) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o OperationCredentialsForOrganizationEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationCredentialsForOrganizationEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OperationCredentialId.IsSet() {
		toSerialize["operationCredentialId"] = o.OperationCredentialId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableOperationCredentialsForOrganizationEntity struct {
	value *OperationCredentialsForOrganizationEntity
	isSet bool
}

func (v NullableOperationCredentialsForOrganizationEntity) Get() *OperationCredentialsForOrganizationEntity {
	return v.value
}

func (v *NullableOperationCredentialsForOrganizationEntity) Set(val *OperationCredentialsForOrganizationEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationCredentialsForOrganizationEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationCredentialsForOrganizationEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationCredentialsForOrganizationEntity(val *OperationCredentialsForOrganizationEntity) *NullableOperationCredentialsForOrganizationEntity {
	return &NullableOperationCredentialsForOrganizationEntity{value: val, isSet: true}
}

func (v NullableOperationCredentialsForOrganizationEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationCredentialsForOrganizationEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


