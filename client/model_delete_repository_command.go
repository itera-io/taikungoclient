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

// checks if the DeleteRepositoryCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRepositoryCommand{}

// DeleteRepositoryCommand struct for DeleteRepositoryCommand
type DeleteRepositoryCommand struct {
	RepoName NullableString `json:"repoName,omitempty"`
}

// NewDeleteRepositoryCommand instantiates a new DeleteRepositoryCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRepositoryCommand() *DeleteRepositoryCommand {
	this := DeleteRepositoryCommand{}
	return &this
}

// NewDeleteRepositoryCommandWithDefaults instantiates a new DeleteRepositoryCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRepositoryCommandWithDefaults() *DeleteRepositoryCommand {
	this := DeleteRepositoryCommand{}
	return &this
}

// GetRepoName returns the RepoName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteRepositoryCommand) GetRepoName() string {
	if o == nil || IsNil(o.RepoName.Get()) {
		var ret string
		return ret
	}
	return *o.RepoName.Get()
}

// GetRepoNameOk returns a tuple with the RepoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteRepositoryCommand) GetRepoNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepoName.Get(), o.RepoName.IsSet()
}

// HasRepoName returns a boolean if a field has been set.
func (o *DeleteRepositoryCommand) HasRepoName() bool {
	if o != nil && o.RepoName.IsSet() {
		return true
	}

	return false
}

// SetRepoName gets a reference to the given NullableString and assigns it to the RepoName field.
func (o *DeleteRepositoryCommand) SetRepoName(v string) {
	o.RepoName.Set(&v)
}
// SetRepoNameNil sets the value for RepoName to be an explicit nil
func (o *DeleteRepositoryCommand) SetRepoNameNil() {
	o.RepoName.Set(nil)
}

// UnsetRepoName ensures that no value is present for RepoName, not even an explicit nil
func (o *DeleteRepositoryCommand) UnsetRepoName() {
	o.RepoName.Unset()
}

func (o DeleteRepositoryCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRepositoryCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RepoName.IsSet() {
		toSerialize["repoName"] = o.RepoName.Get()
	}
	return toSerialize, nil
}

type NullableDeleteRepositoryCommand struct {
	value *DeleteRepositoryCommand
	isSet bool
}

func (v NullableDeleteRepositoryCommand) Get() *DeleteRepositoryCommand {
	return v.value
}

func (v *NullableDeleteRepositoryCommand) Set(val *DeleteRepositoryCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRepositoryCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRepositoryCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRepositoryCommand(val *DeleteRepositoryCommand) *NullableDeleteRepositoryCommand {
	return &NullableDeleteRepositoryCommand{value: val, isSet: true}
}

func (v NullableDeleteRepositoryCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRepositoryCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


