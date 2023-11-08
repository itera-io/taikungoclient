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

// checks if the PartnerDetailsForUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerDetailsForUserDto{}

// PartnerDetailsForUserDto struct for PartnerDetailsForUserDto
type PartnerDetailsForUserDto struct {
	Name NullableString `json:"name,omitempty"`
	Logo NullableString `json:"logo,omitempty"`
	Link NullableString `json:"link,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewPartnerDetailsForUserDto instantiates a new PartnerDetailsForUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerDetailsForUserDto() *PartnerDetailsForUserDto {
	this := PartnerDetailsForUserDto{}
	return &this
}

// NewPartnerDetailsForUserDtoWithDefaults instantiates a new PartnerDetailsForUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerDetailsForUserDtoWithDefaults() *PartnerDetailsForUserDto {
	this := PartnerDetailsForUserDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerDetailsForUserDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerDetailsForUserDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PartnerDetailsForUserDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PartnerDetailsForUserDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PartnerDetailsForUserDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PartnerDetailsForUserDto) UnsetName() {
	o.Name.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerDetailsForUserDto) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerDetailsForUserDto) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *PartnerDetailsForUserDto) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *PartnerDetailsForUserDto) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *PartnerDetailsForUserDto) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *PartnerDetailsForUserDto) UnsetLogo() {
	o.Logo.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartnerDetailsForUserDto) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartnerDetailsForUserDto) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *PartnerDetailsForUserDto) HasLink() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *PartnerDetailsForUserDto) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *PartnerDetailsForUserDto) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *PartnerDetailsForUserDto) UnsetLink() {
	o.Link.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PartnerDetailsForUserDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerDetailsForUserDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartnerDetailsForUserDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PartnerDetailsForUserDto) SetId(v int32) {
	o.Id = &v
}

func (o PartnerDetailsForUserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerDetailsForUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePartnerDetailsForUserDto struct {
	value *PartnerDetailsForUserDto
	isSet bool
}

func (v NullablePartnerDetailsForUserDto) Get() *PartnerDetailsForUserDto {
	return v.value
}

func (v *NullablePartnerDetailsForUserDto) Set(val *PartnerDetailsForUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerDetailsForUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerDetailsForUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerDetailsForUserDto(val *PartnerDetailsForUserDto) *NullablePartnerDetailsForUserDto {
	return &NullablePartnerDetailsForUserDto{value: val, isSet: true}
}

func (v NullablePartnerDetailsForUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerDetailsForUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


