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
	"bytes"
	"fmt"
)

// checks if the OrganizationEntityForDashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationEntityForDashboard{}

// OrganizationEntityForDashboard struct for OrganizationEntityForDashboard
type OrganizationEntityForDashboard struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	Users int32 `json:"users"`
}

type _OrganizationEntityForDashboard OrganizationEntityForDashboard

// NewOrganizationEntityForDashboard instantiates a new OrganizationEntityForDashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEntityForDashboard(id int32, name NullableString, users int32) *OrganizationEntityForDashboard {
	this := OrganizationEntityForDashboard{}
	this.Id = id
	this.Name = name
	this.Users = users
	return &this
}

// NewOrganizationEntityForDashboardWithDefaults instantiates a new OrganizationEntityForDashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEntityForDashboardWithDefaults() *OrganizationEntityForDashboard {
	this := OrganizationEntityForDashboard{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationEntityForDashboard) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationEntityForDashboard) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationEntityForDashboard) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OrganizationEntityForDashboard) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEntityForDashboard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *OrganizationEntityForDashboard) SetName(v string) {
	o.Name.Set(&v)
}

// GetUsers returns the Users field value
func (o *OrganizationEntityForDashboard) GetUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *OrganizationEntityForDashboard) GetUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Users, true
}

// SetUsers sets field value
func (o *OrganizationEntityForDashboard) SetUsers(v int32) {
	o.Users = v
}

func (o OrganizationEntityForDashboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationEntityForDashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *OrganizationEntityForDashboard) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"users",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrganizationEntityForDashboard := _OrganizationEntityForDashboard{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationEntityForDashboard)

	if err != nil {
		return err
	}

	*o = OrganizationEntityForDashboard(varOrganizationEntityForDashboard)

	return err
}

type NullableOrganizationEntityForDashboard struct {
	value *OrganizationEntityForDashboard
	isSet bool
}

func (v NullableOrganizationEntityForDashboard) Get() *OrganizationEntityForDashboard {
	return v.value
}

func (v *NullableOrganizationEntityForDashboard) Set(val *OrganizationEntityForDashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEntityForDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEntityForDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEntityForDashboard(val *OrganizationEntityForDashboard) *NullableOrganizationEntityForDashboard {
	return &NullableOrganizationEntityForDashboard{value: val, isSet: true}
}

func (v NullableOrganizationEntityForDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEntityForDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


