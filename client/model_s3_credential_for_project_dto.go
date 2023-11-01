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

// checks if the S3CredentialForProjectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3CredentialForProjectDto{}

// S3CredentialForProjectDto struct for S3CredentialForProjectDto
type S3CredentialForProjectDto struct {
	S3AccessKeyId NullableString `json:"s3AccessKeyId,omitempty"`
	S3SecretKey   NullableString `json:"s3SecretKey,omitempty"`
	S3Endpoint    NullableString `json:"s3Endpoint,omitempty"`
	S3Region      NullableString `json:"s3Region,omitempty"`
}

// NewS3CredentialForProjectDto instantiates a new S3CredentialForProjectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3CredentialForProjectDto() *S3CredentialForProjectDto {
	this := S3CredentialForProjectDto{}
	return &this
}

// NewS3CredentialForProjectDtoWithDefaults instantiates a new S3CredentialForProjectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3CredentialForProjectDtoWithDefaults() *S3CredentialForProjectDto {
	this := S3CredentialForProjectDto{}
	return &this
}

// GetS3AccessKeyId returns the S3AccessKeyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3CredentialForProjectDto) GetS3AccessKeyId() string {
	if o == nil || IsNil(o.S3AccessKeyId.Get()) {
		var ret string
		return ret
	}
	return *o.S3AccessKeyId.Get()
}

// GetS3AccessKeyIdOk returns a tuple with the S3AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3CredentialForProjectDto) GetS3AccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3AccessKeyId.Get(), o.S3AccessKeyId.IsSet()
}

// HasS3AccessKeyId returns a boolean if a field has been set.
func (o *S3CredentialForProjectDto) HasS3AccessKeyId() bool {
	if o != nil && o.S3AccessKeyId.IsSet() {
		return true
	}

	return false
}

// SetS3AccessKeyId gets a reference to the given NullableString and assigns it to the S3AccessKeyId field.
func (o *S3CredentialForProjectDto) SetS3AccessKeyId(v string) {
	o.S3AccessKeyId.Set(&v)
}

// SetS3AccessKeyIdNil sets the value for S3AccessKeyId to be an explicit nil
func (o *S3CredentialForProjectDto) SetS3AccessKeyIdNil() {
	o.S3AccessKeyId.Set(nil)
}

// UnsetS3AccessKeyId ensures that no value is present for S3AccessKeyId, not even an explicit nil
func (o *S3CredentialForProjectDto) UnsetS3AccessKeyId() {
	o.S3AccessKeyId.Unset()
}

// GetS3SecretKey returns the S3SecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3CredentialForProjectDto) GetS3SecretKey() string {
	if o == nil || IsNil(o.S3SecretKey.Get()) {
		var ret string
		return ret
	}
	return *o.S3SecretKey.Get()
}

// GetS3SecretKeyOk returns a tuple with the S3SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3CredentialForProjectDto) GetS3SecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3SecretKey.Get(), o.S3SecretKey.IsSet()
}

// HasS3SecretKey returns a boolean if a field has been set.
func (o *S3CredentialForProjectDto) HasS3SecretKey() bool {
	if o != nil && o.S3SecretKey.IsSet() {
		return true
	}

	return false
}

// SetS3SecretKey gets a reference to the given NullableString and assigns it to the S3SecretKey field.
func (o *S3CredentialForProjectDto) SetS3SecretKey(v string) {
	o.S3SecretKey.Set(&v)
}

// SetS3SecretKeyNil sets the value for S3SecretKey to be an explicit nil
func (o *S3CredentialForProjectDto) SetS3SecretKeyNil() {
	o.S3SecretKey.Set(nil)
}

// UnsetS3SecretKey ensures that no value is present for S3SecretKey, not even an explicit nil
func (o *S3CredentialForProjectDto) UnsetS3SecretKey() {
	o.S3SecretKey.Unset()
}

// GetS3Endpoint returns the S3Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3CredentialForProjectDto) GetS3Endpoint() string {
	if o == nil || IsNil(o.S3Endpoint.Get()) {
		var ret string
		return ret
	}
	return *o.S3Endpoint.Get()
}

// GetS3EndpointOk returns a tuple with the S3Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3CredentialForProjectDto) GetS3EndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Endpoint.Get(), o.S3Endpoint.IsSet()
}

// HasS3Endpoint returns a boolean if a field has been set.
func (o *S3CredentialForProjectDto) HasS3Endpoint() bool {
	if o != nil && o.S3Endpoint.IsSet() {
		return true
	}

	return false
}

// SetS3Endpoint gets a reference to the given NullableString and assigns it to the S3Endpoint field.
func (o *S3CredentialForProjectDto) SetS3Endpoint(v string) {
	o.S3Endpoint.Set(&v)
}

// SetS3EndpointNil sets the value for S3Endpoint to be an explicit nil
func (o *S3CredentialForProjectDto) SetS3EndpointNil() {
	o.S3Endpoint.Set(nil)
}

// UnsetS3Endpoint ensures that no value is present for S3Endpoint, not even an explicit nil
func (o *S3CredentialForProjectDto) UnsetS3Endpoint() {
	o.S3Endpoint.Unset()
}

// GetS3Region returns the S3Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *S3CredentialForProjectDto) GetS3Region() string {
	if o == nil || IsNil(o.S3Region.Get()) {
		var ret string
		return ret
	}
	return *o.S3Region.Get()
}

// GetS3RegionOk returns a tuple with the S3Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3CredentialForProjectDto) GetS3RegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.S3Region.Get(), o.S3Region.IsSet()
}

// HasS3Region returns a boolean if a field has been set.
func (o *S3CredentialForProjectDto) HasS3Region() bool {
	if o != nil && o.S3Region.IsSet() {
		return true
	}

	return false
}

// SetS3Region gets a reference to the given NullableString and assigns it to the S3Region field.
func (o *S3CredentialForProjectDto) SetS3Region(v string) {
	o.S3Region.Set(&v)
}

// SetS3RegionNil sets the value for S3Region to be an explicit nil
func (o *S3CredentialForProjectDto) SetS3RegionNil() {
	o.S3Region.Set(nil)
}

// UnsetS3Region ensures that no value is present for S3Region, not even an explicit nil
func (o *S3CredentialForProjectDto) UnsetS3Region() {
	o.S3Region.Unset()
}

func (o S3CredentialForProjectDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3CredentialForProjectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.S3AccessKeyId.IsSet() {
		toSerialize["s3AccessKeyId"] = o.S3AccessKeyId.Get()
	}
	if o.S3SecretKey.IsSet() {
		toSerialize["s3SecretKey"] = o.S3SecretKey.Get()
	}
	if o.S3Endpoint.IsSet() {
		toSerialize["s3Endpoint"] = o.S3Endpoint.Get()
	}
	if o.S3Region.IsSet() {
		toSerialize["s3Region"] = o.S3Region.Get()
	}
	return toSerialize, nil
}

type NullableS3CredentialForProjectDto struct {
	value *S3CredentialForProjectDto
	isSet bool
}

func (v NullableS3CredentialForProjectDto) Get() *S3CredentialForProjectDto {
	return v.value
}

func (v *NullableS3CredentialForProjectDto) Set(val *S3CredentialForProjectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableS3CredentialForProjectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableS3CredentialForProjectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3CredentialForProjectDto(val *S3CredentialForProjectDto) *NullableS3CredentialForProjectDto {
	return &NullableS3CredentialForProjectDto{value: val, isSet: true}
}

func (v NullableS3CredentialForProjectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3CredentialForProjectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}