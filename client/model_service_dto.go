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

// checks if the ServiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDto{}

// ServiceDto struct for ServiceDto
type ServiceDto struct {
	MetadataName string `json:"metadataName"`
	Namespace string `json:"namespace"`
	Age NullableString `json:"age"`
	Type NullableString `json:"type"`
	Ip interface{} `json:"ip"`
}

type _ServiceDto ServiceDto

// NewServiceDto instantiates a new ServiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDto(metadataName string, namespace string, age NullableString, type_ NullableString, ip interface{}) *ServiceDto {
	this := ServiceDto{}
	this.MetadataName = metadataName
	this.Namespace = namespace
	this.Age = age
	this.Type = type_
	this.Ip = ip
	return &this
}

// NewServiceDtoWithDefaults instantiates a new ServiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDtoWithDefaults() *ServiceDto {
	this := ServiceDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value
func (o *ServiceDto) GetMetadataName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataName, true
}

// SetMetadataName sets field value
func (o *ServiceDto) SetMetadataName(v string) {
	o.MetadataName = v
}

// GetNamespace returns the Namespace field value
func (o *ServiceDto) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ServiceDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ServiceDto) SetNamespace(v string) {
	o.Namespace = v
}

// GetAge returns the Age field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceDto) GetAge() string {
	if o == nil || o.Age.Get() == nil {
		var ret string
		return ret
	}

	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// SetAge sets field value
func (o *ServiceDto) SetAge(v string) {
	o.Age.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ServiceDto) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *ServiceDto) SetType(v string) {
	o.Type.Set(&v)
}

// GetIp returns the Ip field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ServiceDto) GetIp() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceDto) GetIpOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *ServiceDto) SetIp(v interface{}) {
	o.Ip = v
}

func (o ServiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadataName"] = o.MetadataName
	toSerialize["namespace"] = o.Namespace
	toSerialize["age"] = o.Age.Get()
	toSerialize["type"] = o.Type.Get()
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

func (o *ServiceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataName",
		"namespace",
		"age",
		"type",
		"ip",
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

	varServiceDto := _ServiceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceDto)

	if err != nil {
		return err
	}

	*o = ServiceDto(varServiceDto)

	return err
}

type NullableServiceDto struct {
	value *ServiceDto
	isSet bool
}

func (v NullableServiceDto) Get() *ServiceDto {
	return v.value
}

func (v *NullableServiceDto) Set(val *ServiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDto(val *ServiceDto) *NullableServiceDto {
	return &NullableServiceDto{value: val, isSet: true}
}

func (v NullableServiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


