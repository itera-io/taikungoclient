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

// checks if the KeycloakListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeycloakListDto{}

// KeycloakListDto struct for KeycloakListDto
type KeycloakListDto struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	ClientId string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RealmsName string `json:"realmsName"`
	OrganizationId int32 `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	PartnerLogo NullableString `json:"partnerLogo"`
	Enabled bool `json:"enabled"`
}

type _KeycloakListDto KeycloakListDto

// NewKeycloakListDto instantiates a new KeycloakListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeycloakListDto(id int32, name string, url string, clientId string, clientSecret string, realmsName string, organizationId int32, organizationName string, partnerLogo NullableString, enabled bool) *KeycloakListDto {
	this := KeycloakListDto{}
	this.Id = id
	this.Name = name
	this.Url = url
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.RealmsName = realmsName
	this.OrganizationId = organizationId
	this.OrganizationName = organizationName
	this.PartnerLogo = partnerLogo
	this.Enabled = enabled
	return &this
}

// NewKeycloakListDtoWithDefaults instantiates a new KeycloakListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeycloakListDtoWithDefaults() *KeycloakListDto {
	this := KeycloakListDto{}
	return &this
}

// GetId returns the Id field value
func (o *KeycloakListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KeycloakListDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KeycloakListDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KeycloakListDto) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *KeycloakListDto) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *KeycloakListDto) SetUrl(v string) {
	o.Url = v
}

// GetClientId returns the ClientId field value
func (o *KeycloakListDto) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *KeycloakListDto) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *KeycloakListDto) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *KeycloakListDto) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetRealmsName returns the RealmsName field value
func (o *KeycloakListDto) GetRealmsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RealmsName
}

// GetRealmsNameOk returns a tuple with the RealmsName field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetRealmsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RealmsName, true
}

// SetRealmsName sets field value
func (o *KeycloakListDto) SetRealmsName(v string) {
	o.RealmsName = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *KeycloakListDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *KeycloakListDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *KeycloakListDto) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *KeycloakListDto) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetPartnerLogo returns the PartnerLogo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KeycloakListDto) GetPartnerLogo() string {
	if o == nil || o.PartnerLogo.Get() == nil {
		var ret string
		return ret
	}

	return *o.PartnerLogo.Get()
}

// GetPartnerLogoOk returns a tuple with the PartnerLogo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KeycloakListDto) GetPartnerLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerLogo.Get(), o.PartnerLogo.IsSet()
}

// SetPartnerLogo sets field value
func (o *KeycloakListDto) SetPartnerLogo(v string) {
	o.PartnerLogo.Set(&v)
}

// GetEnabled returns the Enabled field value
func (o *KeycloakListDto) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *KeycloakListDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *KeycloakListDto) SetEnabled(v bool) {
	o.Enabled = v
}

func (o KeycloakListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeycloakListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["realmsName"] = o.RealmsName
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["organizationName"] = o.OrganizationName
	toSerialize["partnerLogo"] = o.PartnerLogo.Get()
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *KeycloakListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"url",
		"clientId",
		"clientSecret",
		"realmsName",
		"organizationId",
		"organizationName",
		"partnerLogo",
		"enabled",
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

	varKeycloakListDto := _KeycloakListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeycloakListDto)

	if err != nil {
		return err
	}

	*o = KeycloakListDto(varKeycloakListDto)

	return err
}

type NullableKeycloakListDto struct {
	value *KeycloakListDto
	isSet bool
}

func (v NullableKeycloakListDto) Get() *KeycloakListDto {
	return v.value
}

func (v *NullableKeycloakListDto) Set(val *KeycloakListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableKeycloakListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableKeycloakListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeycloakListDto(val *KeycloakListDto) *NullableKeycloakListDto {
	return &NullableKeycloakListDto{value: val, isSet: true}
}

func (v NullableKeycloakListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeycloakListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


