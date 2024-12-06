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

// checks if the PodDisruptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PodDisruptionDto{}

// PodDisruptionDto struct for PodDisruptionDto
type PodDisruptionDto struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	MinAvailable interface{} `json:"minAvailable"`
	MaxAvailable interface{} `json:"maxAvailable"`
	AllowedDisruptions interface{} `json:"allowedDisruptions"`
	CreatedAt NullableString `json:"createdAt"`
}

type _PodDisruptionDto PodDisruptionDto

// NewPodDisruptionDto instantiates a new PodDisruptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodDisruptionDto(name string, namespace string, minAvailable interface{}, maxAvailable interface{}, allowedDisruptions interface{}, createdAt NullableString) *PodDisruptionDto {
	this := PodDisruptionDto{}
	this.Name = name
	this.Namespace = namespace
	this.MinAvailable = minAvailable
	this.MaxAvailable = maxAvailable
	this.AllowedDisruptions = allowedDisruptions
	this.CreatedAt = createdAt
	return &this
}

// NewPodDisruptionDtoWithDefaults instantiates a new PodDisruptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodDisruptionDtoWithDefaults() *PodDisruptionDto {
	this := PodDisruptionDto{}
	return &this
}

// GetName returns the Name field value
func (o *PodDisruptionDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PodDisruptionDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PodDisruptionDto) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *PodDisruptionDto) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PodDisruptionDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PodDisruptionDto) SetNamespace(v string) {
	o.Namespace = v
}

// GetMinAvailable returns the MinAvailable field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PodDisruptionDto) GetMinAvailable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MinAvailable
}

// GetMinAvailableOk returns a tuple with the MinAvailable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetMinAvailableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MinAvailable) {
		return nil, false
	}
	return &o.MinAvailable, true
}

// SetMinAvailable sets field value
func (o *PodDisruptionDto) SetMinAvailable(v interface{}) {
	o.MinAvailable = v
}

// GetMaxAvailable returns the MaxAvailable field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PodDisruptionDto) GetMaxAvailable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MaxAvailable
}

// GetMaxAvailableOk returns a tuple with the MaxAvailable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetMaxAvailableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MaxAvailable) {
		return nil, false
	}
	return &o.MaxAvailable, true
}

// SetMaxAvailable sets field value
func (o *PodDisruptionDto) SetMaxAvailable(v interface{}) {
	o.MaxAvailable = v
}

// GetAllowedDisruptions returns the AllowedDisruptions field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PodDisruptionDto) GetAllowedDisruptions() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AllowedDisruptions
}

// GetAllowedDisruptionsOk returns a tuple with the AllowedDisruptions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetAllowedDisruptionsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowedDisruptions) {
		return nil, false
	}
	return &o.AllowedDisruptions, true
}

// SetAllowedDisruptions sets field value
func (o *PodDisruptionDto) SetAllowedDisruptions(v interface{}) {
	o.AllowedDisruptions = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PodDisruptionDto) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodDisruptionDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *PodDisruptionDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

func (o PodDisruptionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodDisruptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	if o.MinAvailable != nil {
		toSerialize["minAvailable"] = o.MinAvailable
	}
	if o.MaxAvailable != nil {
		toSerialize["maxAvailable"] = o.MaxAvailable
	}
	if o.AllowedDisruptions != nil {
		toSerialize["allowedDisruptions"] = o.AllowedDisruptions
	}
	toSerialize["createdAt"] = o.CreatedAt.Get()
	return toSerialize, nil
}

func (o *PodDisruptionDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"namespace",
		"minAvailable",
		"maxAvailable",
		"allowedDisruptions",
		"createdAt",
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

	varPodDisruptionDto := _PodDisruptionDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPodDisruptionDto)

	if err != nil {
		return err
	}

	*o = PodDisruptionDto(varPodDisruptionDto)

	return err
}

type NullablePodDisruptionDto struct {
	value *PodDisruptionDto
	isSet bool
}

func (v NullablePodDisruptionDto) Get() *PodDisruptionDto {
	return v.value
}

func (v *NullablePodDisruptionDto) Set(val *PodDisruptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePodDisruptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePodDisruptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodDisruptionDto(val *PodDisruptionDto) *NullablePodDisruptionDto {
	return &NullablePodDisruptionDto{value: val, isSet: true}
}

func (v NullablePodDisruptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodDisruptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


