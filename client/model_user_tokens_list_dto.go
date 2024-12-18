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

// checks if the UserTokensListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTokensListDto{}

// UserTokensListDto struct for UserTokensListDto
type UserTokensListDto struct {
	Id NullableString `json:"id"`
	Name NullableString `json:"name"`
	CreatedAt NullableString `json:"createdAt"`
	IsReadonly bool `json:"isReadonly"`
	ExpireDate NullableString `json:"expireDate"`
	AccessKey NullableString `json:"accessKey"`
}

type _UserTokensListDto UserTokensListDto

// NewUserTokensListDto instantiates a new UserTokensListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTokensListDto(id NullableString, name NullableString, createdAt NullableString, isReadonly bool, expireDate NullableString, accessKey NullableString) *UserTokensListDto {
	this := UserTokensListDto{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.IsReadonly = isReadonly
	this.ExpireDate = expireDate
	this.AccessKey = accessKey
	return &this
}

// NewUserTokensListDtoWithDefaults instantiates a new UserTokensListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTokensListDtoWithDefaults() *UserTokensListDto {
	this := UserTokensListDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserTokensListDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokensListDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *UserTokensListDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserTokensListDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokensListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *UserTokensListDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserTokensListDto) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokensListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *UserTokensListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetIsReadonly returns the IsReadonly field value
func (o *UserTokensListDto) GetIsReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReadonly
}

// GetIsReadonlyOk returns a tuple with the IsReadonly field value
// and a boolean to check if the value has been set.
func (o *UserTokensListDto) GetIsReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReadonly, true
}

// SetIsReadonly sets field value
func (o *UserTokensListDto) SetIsReadonly(v bool) {
	o.IsReadonly = v
}

// GetExpireDate returns the ExpireDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserTokensListDto) GetExpireDate() string {
	if o == nil || o.ExpireDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpireDate.Get()
}

// GetExpireDateOk returns a tuple with the ExpireDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokensListDto) GetExpireDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpireDate.Get(), o.ExpireDate.IsSet()
}

// SetExpireDate sets field value
func (o *UserTokensListDto) SetExpireDate(v string) {
	o.ExpireDate.Set(&v)
}

// GetAccessKey returns the AccessKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserTokensListDto) GetAccessKey() string {
	if o == nil || o.AccessKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessKey.Get()
}

// GetAccessKeyOk returns a tuple with the AccessKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserTokensListDto) GetAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessKey.Get(), o.AccessKey.IsSet()
}

// SetAccessKey sets field value
func (o *UserTokensListDto) SetAccessKey(v string) {
	o.AccessKey.Set(&v)
}

func (o UserTokensListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTokensListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["isReadonly"] = o.IsReadonly
	toSerialize["expireDate"] = o.ExpireDate.Get()
	toSerialize["accessKey"] = o.AccessKey.Get()
	return toSerialize, nil
}

func (o *UserTokensListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"createdAt",
		"isReadonly",
		"expireDate",
		"accessKey",
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

	varUserTokensListDto := _UserTokensListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserTokensListDto)

	if err != nil {
		return err
	}

	*o = UserTokensListDto(varUserTokensListDto)

	return err
}

type NullableUserTokensListDto struct {
	value *UserTokensListDto
	isSet bool
}

func (v NullableUserTokensListDto) Get() *UserTokensListDto {
	return v.value
}

func (v *NullableUserTokensListDto) Set(val *UserTokensListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTokensListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTokensListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTokensListDto(val *UserTokensListDto) *NullableUserTokensListDto {
	return &NullableUserTokensListDto{value: val, isSet: true}
}

func (v NullableUserTokensListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTokensListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


