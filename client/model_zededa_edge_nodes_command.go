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

// checks if the ZededaEdgeNodesCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZededaEdgeNodesCommand{}

// ZededaEdgeNodesCommand struct for ZededaEdgeNodesCommand
type ZededaEdgeNodesCommand struct {
	ApiUrl NullableString `json:"apiUrl,omitempty"`
	ApiToken NullableString `json:"apiToken,omitempty"`
	Project NullableString `json:"project,omitempty"`
}

// NewZededaEdgeNodesCommand instantiates a new ZededaEdgeNodesCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZededaEdgeNodesCommand() *ZededaEdgeNodesCommand {
	this := ZededaEdgeNodesCommand{}
	return &this
}

// NewZededaEdgeNodesCommandWithDefaults instantiates a new ZededaEdgeNodesCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZededaEdgeNodesCommandWithDefaults() *ZededaEdgeNodesCommand {
	this := ZededaEdgeNodesCommand{}
	return &this
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZededaEdgeNodesCommand) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ApiUrl.Get()
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZededaEdgeNodesCommand) GetApiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUrl.Get(), o.ApiUrl.IsSet()
}

// HasApiUrl returns a boolean if a field has been set.
func (o *ZededaEdgeNodesCommand) HasApiUrl() bool {
	if o != nil && o.ApiUrl.IsSet() {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given NullableString and assigns it to the ApiUrl field.
func (o *ZededaEdgeNodesCommand) SetApiUrl(v string) {
	o.ApiUrl.Set(&v)
}
// SetApiUrlNil sets the value for ApiUrl to be an explicit nil
func (o *ZededaEdgeNodesCommand) SetApiUrlNil() {
	o.ApiUrl.Set(nil)
}

// UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
func (o *ZededaEdgeNodesCommand) UnsetApiUrl() {
	o.ApiUrl.Unset()
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZededaEdgeNodesCommand) GetApiToken() string {
	if o == nil || IsNil(o.ApiToken.Get()) {
		var ret string
		return ret
	}
	return *o.ApiToken.Get()
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZededaEdgeNodesCommand) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiToken.Get(), o.ApiToken.IsSet()
}

// HasApiToken returns a boolean if a field has been set.
func (o *ZededaEdgeNodesCommand) HasApiToken() bool {
	if o != nil && o.ApiToken.IsSet() {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given NullableString and assigns it to the ApiToken field.
func (o *ZededaEdgeNodesCommand) SetApiToken(v string) {
	o.ApiToken.Set(&v)
}
// SetApiTokenNil sets the value for ApiToken to be an explicit nil
func (o *ZededaEdgeNodesCommand) SetApiTokenNil() {
	o.ApiToken.Set(nil)
}

// UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil
func (o *ZededaEdgeNodesCommand) UnsetApiToken() {
	o.ApiToken.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZededaEdgeNodesCommand) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZededaEdgeNodesCommand) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *ZededaEdgeNodesCommand) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *ZededaEdgeNodesCommand) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *ZededaEdgeNodesCommand) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *ZededaEdgeNodesCommand) UnsetProject() {
	o.Project.Unset()
}

func (o ZededaEdgeNodesCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZededaEdgeNodesCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiUrl.IsSet() {
		toSerialize["apiUrl"] = o.ApiUrl.Get()
	}
	if o.ApiToken.IsSet() {
		toSerialize["apiToken"] = o.ApiToken.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	return toSerialize, nil
}

type NullableZededaEdgeNodesCommand struct {
	value *ZededaEdgeNodesCommand
	isSet bool
}

func (v NullableZededaEdgeNodesCommand) Get() *ZededaEdgeNodesCommand {
	return v.value
}

func (v *NullableZededaEdgeNodesCommand) Set(val *ZededaEdgeNodesCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableZededaEdgeNodesCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableZededaEdgeNodesCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZededaEdgeNodesCommand(val *ZededaEdgeNodesCommand) *NullableZededaEdgeNodesCommand {
	return &NullableZededaEdgeNodesCommand{value: val, isSet: true}
}

func (v NullableZededaEdgeNodesCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZededaEdgeNodesCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


