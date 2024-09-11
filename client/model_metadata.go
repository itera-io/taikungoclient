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

// checks if the Metadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Metadata{}

// Metadata struct for Metadata
type Metadata struct {
	InfracostCommand *string `json:"infracostCommand,omitempty"`
	VcsBranch *string `json:"vcsBranch,omitempty"`
	VcsCommitSha *string `json:"vcsCommitSha,omitempty"`
	VcsCommitAuthorName *string `json:"vcsCommitAuthorName,omitempty"`
	VcsCommitAuthorEmail *string `json:"vcsCommitAuthorEmail,omitempty"`
	VcsCommitTimestamp *string `json:"vcsCommitTimestamp,omitempty"`
	VcsCommitMessage *string `json:"vcsCommitMessage,omitempty"`
	VcsRepositoryUrl *string `json:"vcsRepositoryUrl,omitempty"`
	UsageApiEnabled *bool `json:"usageApiEnabled,omitempty"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}
	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetInfracostCommand returns the InfracostCommand field value if set, zero value otherwise.
func (o *Metadata) GetInfracostCommand() string {
	if o == nil || IsNil(o.InfracostCommand) {
		var ret string
		return ret
	}
	return *o.InfracostCommand
}

// GetInfracostCommandOk returns a tuple with the InfracostCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetInfracostCommandOk() (*string, bool) {
	if o == nil || IsNil(o.InfracostCommand) {
		return nil, false
	}
	return o.InfracostCommand, true
}

// HasInfracostCommand returns a boolean if a field has been set.
func (o *Metadata) HasInfracostCommand() bool {
	if o != nil && !IsNil(o.InfracostCommand) {
		return true
	}

	return false
}

// SetInfracostCommand gets a reference to the given string and assigns it to the InfracostCommand field.
func (o *Metadata) SetInfracostCommand(v string) {
	o.InfracostCommand = &v
}

// GetVcsBranch returns the VcsBranch field value if set, zero value otherwise.
func (o *Metadata) GetVcsBranch() string {
	if o == nil || IsNil(o.VcsBranch) {
		var ret string
		return ret
	}
	return *o.VcsBranch
}

// GetVcsBranchOk returns a tuple with the VcsBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsBranchOk() (*string, bool) {
	if o == nil || IsNil(o.VcsBranch) {
		return nil, false
	}
	return o.VcsBranch, true
}

// HasVcsBranch returns a boolean if a field has been set.
func (o *Metadata) HasVcsBranch() bool {
	if o != nil && !IsNil(o.VcsBranch) {
		return true
	}

	return false
}

// SetVcsBranch gets a reference to the given string and assigns it to the VcsBranch field.
func (o *Metadata) SetVcsBranch(v string) {
	o.VcsBranch = &v
}

// GetVcsCommitSha returns the VcsCommitSha field value if set, zero value otherwise.
func (o *Metadata) GetVcsCommitSha() string {
	if o == nil || IsNil(o.VcsCommitSha) {
		var ret string
		return ret
	}
	return *o.VcsCommitSha
}

// GetVcsCommitShaOk returns a tuple with the VcsCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.VcsCommitSha) {
		return nil, false
	}
	return o.VcsCommitSha, true
}

// HasVcsCommitSha returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitSha() bool {
	if o != nil && !IsNil(o.VcsCommitSha) {
		return true
	}

	return false
}

// SetVcsCommitSha gets a reference to the given string and assigns it to the VcsCommitSha field.
func (o *Metadata) SetVcsCommitSha(v string) {
	o.VcsCommitSha = &v
}

// GetVcsCommitAuthorName returns the VcsCommitAuthorName field value if set, zero value otherwise.
func (o *Metadata) GetVcsCommitAuthorName() string {
	if o == nil || IsNil(o.VcsCommitAuthorName) {
		var ret string
		return ret
	}
	return *o.VcsCommitAuthorName
}

// GetVcsCommitAuthorNameOk returns a tuple with the VcsCommitAuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsCommitAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VcsCommitAuthorName) {
		return nil, false
	}
	return o.VcsCommitAuthorName, true
}

// HasVcsCommitAuthorName returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitAuthorName() bool {
	if o != nil && !IsNil(o.VcsCommitAuthorName) {
		return true
	}

	return false
}

// SetVcsCommitAuthorName gets a reference to the given string and assigns it to the VcsCommitAuthorName field.
func (o *Metadata) SetVcsCommitAuthorName(v string) {
	o.VcsCommitAuthorName = &v
}

// GetVcsCommitAuthorEmail returns the VcsCommitAuthorEmail field value if set, zero value otherwise.
func (o *Metadata) GetVcsCommitAuthorEmail() string {
	if o == nil || IsNil(o.VcsCommitAuthorEmail) {
		var ret string
		return ret
	}
	return *o.VcsCommitAuthorEmail
}

// GetVcsCommitAuthorEmailOk returns a tuple with the VcsCommitAuthorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsCommitAuthorEmailOk() (*string, bool) {
	if o == nil || IsNil(o.VcsCommitAuthorEmail) {
		return nil, false
	}
	return o.VcsCommitAuthorEmail, true
}

// HasVcsCommitAuthorEmail returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitAuthorEmail() bool {
	if o != nil && !IsNil(o.VcsCommitAuthorEmail) {
		return true
	}

	return false
}

// SetVcsCommitAuthorEmail gets a reference to the given string and assigns it to the VcsCommitAuthorEmail field.
func (o *Metadata) SetVcsCommitAuthorEmail(v string) {
	o.VcsCommitAuthorEmail = &v
}

// GetVcsCommitTimestamp returns the VcsCommitTimestamp field value if set, zero value otherwise.
func (o *Metadata) GetVcsCommitTimestamp() string {
	if o == nil || IsNil(o.VcsCommitTimestamp) {
		var ret string
		return ret
	}
	return *o.VcsCommitTimestamp
}

// GetVcsCommitTimestampOk returns a tuple with the VcsCommitTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsCommitTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.VcsCommitTimestamp) {
		return nil, false
	}
	return o.VcsCommitTimestamp, true
}

// HasVcsCommitTimestamp returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitTimestamp() bool {
	if o != nil && !IsNil(o.VcsCommitTimestamp) {
		return true
	}

	return false
}

// SetVcsCommitTimestamp gets a reference to the given string and assigns it to the VcsCommitTimestamp field.
func (o *Metadata) SetVcsCommitTimestamp(v string) {
	o.VcsCommitTimestamp = &v
}

// GetVcsCommitMessage returns the VcsCommitMessage field value if set, zero value otherwise.
func (o *Metadata) GetVcsCommitMessage() string {
	if o == nil || IsNil(o.VcsCommitMessage) {
		var ret string
		return ret
	}
	return *o.VcsCommitMessage
}

// GetVcsCommitMessageOk returns a tuple with the VcsCommitMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsCommitMessageOk() (*string, bool) {
	if o == nil || IsNil(o.VcsCommitMessage) {
		return nil, false
	}
	return o.VcsCommitMessage, true
}

// HasVcsCommitMessage returns a boolean if a field has been set.
func (o *Metadata) HasVcsCommitMessage() bool {
	if o != nil && !IsNil(o.VcsCommitMessage) {
		return true
	}

	return false
}

// SetVcsCommitMessage gets a reference to the given string and assigns it to the VcsCommitMessage field.
func (o *Metadata) SetVcsCommitMessage(v string) {
	o.VcsCommitMessage = &v
}

// GetVcsRepositoryUrl returns the VcsRepositoryUrl field value if set, zero value otherwise.
func (o *Metadata) GetVcsRepositoryUrl() string {
	if o == nil || IsNil(o.VcsRepositoryUrl) {
		var ret string
		return ret
	}
	return *o.VcsRepositoryUrl
}

// GetVcsRepositoryUrlOk returns a tuple with the VcsRepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetVcsRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VcsRepositoryUrl) {
		return nil, false
	}
	return o.VcsRepositoryUrl, true
}

// HasVcsRepositoryUrl returns a boolean if a field has been set.
func (o *Metadata) HasVcsRepositoryUrl() bool {
	if o != nil && !IsNil(o.VcsRepositoryUrl) {
		return true
	}

	return false
}

// SetVcsRepositoryUrl gets a reference to the given string and assigns it to the VcsRepositoryUrl field.
func (o *Metadata) SetVcsRepositoryUrl(v string) {
	o.VcsRepositoryUrl = &v
}

// GetUsageApiEnabled returns the UsageApiEnabled field value if set, zero value otherwise.
func (o *Metadata) GetUsageApiEnabled() bool {
	if o == nil || IsNil(o.UsageApiEnabled) {
		var ret bool
		return ret
	}
	return *o.UsageApiEnabled
}

// GetUsageApiEnabledOk returns a tuple with the UsageApiEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadata) GetUsageApiEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UsageApiEnabled) {
		return nil, false
	}
	return o.UsageApiEnabled, true
}

// HasUsageApiEnabled returns a boolean if a field has been set.
func (o *Metadata) HasUsageApiEnabled() bool {
	if o != nil && !IsNil(o.UsageApiEnabled) {
		return true
	}

	return false
}

// SetUsageApiEnabled gets a reference to the given bool and assigns it to the UsageApiEnabled field.
func (o *Metadata) SetUsageApiEnabled(v bool) {
	o.UsageApiEnabled = &v
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Metadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InfracostCommand) {
		toSerialize["infracostCommand"] = o.InfracostCommand
	}
	if !IsNil(o.VcsBranch) {
		toSerialize["vcsBranch"] = o.VcsBranch
	}
	if !IsNil(o.VcsCommitSha) {
		toSerialize["vcsCommitSha"] = o.VcsCommitSha
	}
	if !IsNil(o.VcsCommitAuthorName) {
		toSerialize["vcsCommitAuthorName"] = o.VcsCommitAuthorName
	}
	if !IsNil(o.VcsCommitAuthorEmail) {
		toSerialize["vcsCommitAuthorEmail"] = o.VcsCommitAuthorEmail
	}
	if !IsNil(o.VcsCommitTimestamp) {
		toSerialize["vcsCommitTimestamp"] = o.VcsCommitTimestamp
	}
	if !IsNil(o.VcsCommitMessage) {
		toSerialize["vcsCommitMessage"] = o.VcsCommitMessage
	}
	if !IsNil(o.VcsRepositoryUrl) {
		toSerialize["vcsRepositoryUrl"] = o.VcsRepositoryUrl
	}
	if !IsNil(o.UsageApiEnabled) {
		toSerialize["usageApiEnabled"] = o.UsageApiEnabled
	}
	return toSerialize, nil
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


