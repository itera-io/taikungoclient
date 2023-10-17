/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the KeycloakCreateCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeycloakCreateCommand{}

// KeycloakCreateCommand struct for KeycloakCreateCommand
type KeycloakCreateCommand struct {
	Name         NullableString `json:"name,omitempty"`
	Url          NullableString `json:"url,omitempty"`
	RealmsName   NullableString `json:"realmsName,omitempty"`
	ClientId     NullableString `json:"clientId,omitempty"`
	ClientSecret NullableString `json:"clientSecret,omitempty"`
}

// NewKeycloakCreateCommand instantiates a new KeycloakCreateCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeycloakCreateCommand() *KeycloakCreateCommand {
	this := KeycloakCreateCommand{}
	return &this
}

// NewKeycloakCreateCommandWithDefaults instantiates a new KeycloakCreateCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeycloakCreateCommandWithDefaults() *KeycloakCreateCommand {
	this := KeycloakCreateCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakCreateCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakCreateCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *KeycloakCreateCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *KeycloakCreateCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *KeycloakCreateCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *KeycloakCreateCommand) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakCreateCommand) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakCreateCommand) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *KeycloakCreateCommand) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *KeycloakCreateCommand) SetUrl(v string) {
	o.Url.Set(&v)
}

// SetUrlNil sets the value for Url to be an explicit nil
func (o *KeycloakCreateCommand) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *KeycloakCreateCommand) UnsetUrl() {
	o.Url.Unset()
}

// GetRealmsName returns the RealmsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakCreateCommand) GetRealmsName() string {
	if o == nil || IsNil(o.RealmsName.Get()) {
		var ret string
		return ret
	}
	return *o.RealmsName.Get()
}

// GetRealmsNameOk returns a tuple with the RealmsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakCreateCommand) GetRealmsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RealmsName.Get(), o.RealmsName.IsSet()
}

// HasRealmsName returns a boolean if a field has been set.
func (o *KeycloakCreateCommand) HasRealmsName() bool {
	if o != nil && o.RealmsName.IsSet() {
		return true
	}

	return false
}

// SetRealmsName gets a reference to the given NullableString and assigns it to the RealmsName field.
func (o *KeycloakCreateCommand) SetRealmsName(v string) {
	o.RealmsName.Set(&v)
}

// SetRealmsNameNil sets the value for RealmsName to be an explicit nil
func (o *KeycloakCreateCommand) SetRealmsNameNil() {
	o.RealmsName.Set(nil)
}

// UnsetRealmsName ensures that no value is present for RealmsName, not even an explicit nil
func (o *KeycloakCreateCommand) UnsetRealmsName() {
	o.RealmsName.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakCreateCommand) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakCreateCommand) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *KeycloakCreateCommand) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *KeycloakCreateCommand) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *KeycloakCreateCommand) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *KeycloakCreateCommand) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KeycloakCreateCommand) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakCreateCommand) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *KeycloakCreateCommand) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *KeycloakCreateCommand) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}

// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *KeycloakCreateCommand) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *KeycloakCreateCommand) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

func (o KeycloakCreateCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeycloakCreateCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.RealmsName.IsSet() {
		toSerialize["realmsName"] = o.RealmsName.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["clientSecret"] = o.ClientSecret.Get()
	}
	return toSerialize, nil
}

type NullableKeycloakCreateCommand struct {
	value *KeycloakCreateCommand
	isSet bool
}

func (v NullableKeycloakCreateCommand) Get() *KeycloakCreateCommand {
	return v.value
}

func (v *NullableKeycloakCreateCommand) Set(val *KeycloakCreateCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableKeycloakCreateCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableKeycloakCreateCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeycloakCreateCommand(val *KeycloakCreateCommand) *NullableKeycloakCreateCommand {
	return &NullableKeycloakCreateCommand{value: val, isSet: true}
}

func (v NullableKeycloakCreateCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeycloakCreateCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
