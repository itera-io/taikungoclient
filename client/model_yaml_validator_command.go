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

// checks if the YamlValidatorCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YamlValidatorCommand{}

// YamlValidatorCommand struct for YamlValidatorCommand
type YamlValidatorCommand struct {
	Yaml NullableString `json:"yaml,omitempty"`
}

// NewYamlValidatorCommand instantiates a new YamlValidatorCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYamlValidatorCommand() *YamlValidatorCommand {
	this := YamlValidatorCommand{}
	return &this
}

// NewYamlValidatorCommandWithDefaults instantiates a new YamlValidatorCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYamlValidatorCommandWithDefaults() *YamlValidatorCommand {
	this := YamlValidatorCommand{}
	return &this
}

// GetYaml returns the Yaml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YamlValidatorCommand) GetYaml() string {
	if o == nil || IsNil(o.Yaml.Get()) {
		var ret string
		return ret
	}
	return *o.Yaml.Get()
}

// GetYamlOk returns a tuple with the Yaml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YamlValidatorCommand) GetYamlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yaml.Get(), o.Yaml.IsSet()
}

// HasYaml returns a boolean if a field has been set.
func (o *YamlValidatorCommand) HasYaml() bool {
	if o != nil && o.Yaml.IsSet() {
		return true
	}

	return false
}

// SetYaml gets a reference to the given NullableString and assigns it to the Yaml field.
func (o *YamlValidatorCommand) SetYaml(v string) {
	o.Yaml.Set(&v)
}
// SetYamlNil sets the value for Yaml to be an explicit nil
func (o *YamlValidatorCommand) SetYamlNil() {
	o.Yaml.Set(nil)
}

// UnsetYaml ensures that no value is present for Yaml, not even an explicit nil
func (o *YamlValidatorCommand) UnsetYaml() {
	o.Yaml.Unset()
}

func (o YamlValidatorCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YamlValidatorCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Yaml.IsSet() {
		toSerialize["yaml"] = o.Yaml.Get()
	}
	return toSerialize, nil
}

type NullableYamlValidatorCommand struct {
	value *YamlValidatorCommand
	isSet bool
}

func (v NullableYamlValidatorCommand) Get() *YamlValidatorCommand {
	return v.value
}

func (v *NullableYamlValidatorCommand) Set(val *YamlValidatorCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableYamlValidatorCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableYamlValidatorCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYamlValidatorCommand(val *YamlValidatorCommand) *NullableYamlValidatorCommand {
	return &NullableYamlValidatorCommand{value: val, isSet: true}
}

func (v NullableYamlValidatorCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYamlValidatorCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


