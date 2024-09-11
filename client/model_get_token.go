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
	"time"
)

// checks if the GetToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetToken{}

// GetToken struct for GetToken
type GetToken struct {
	Token *string `json:"token,omitempty"`
	RefreshToken *string `json:"refreshToken,omitempty"`
	RefreshTokenExpireTime *time.Time `json:"refreshTokenExpireTime,omitempty"`
	TempToken *string `json:"tempToken,omitempty"`
}

// NewGetToken instantiates a new GetToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetToken() *GetToken {
	this := GetToken{}
	return &this
}

// NewGetTokenWithDefaults instantiates a new GetToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokenWithDefaults() *GetToken {
	this := GetToken{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetToken) SetToken(v string) {
	o.Token = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *GetToken) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *GetToken) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *GetToken) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetRefreshTokenExpireTime returns the RefreshTokenExpireTime field value if set, zero value otherwise.
func (o *GetToken) GetRefreshTokenExpireTime() time.Time {
	if o == nil || IsNil(o.RefreshTokenExpireTime) {
		var ret time.Time
		return ret
	}
	return *o.RefreshTokenExpireTime
}

// GetRefreshTokenExpireTimeOk returns a tuple with the RefreshTokenExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken) GetRefreshTokenExpireTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RefreshTokenExpireTime) {
		return nil, false
	}
	return o.RefreshTokenExpireTime, true
}

// HasRefreshTokenExpireTime returns a boolean if a field has been set.
func (o *GetToken) HasRefreshTokenExpireTime() bool {
	if o != nil && !IsNil(o.RefreshTokenExpireTime) {
		return true
	}

	return false
}

// SetRefreshTokenExpireTime gets a reference to the given time.Time and assigns it to the RefreshTokenExpireTime field.
func (o *GetToken) SetRefreshTokenExpireTime(v time.Time) {
	o.RefreshTokenExpireTime = &v
}

// GetTempToken returns the TempToken field value if set, zero value otherwise.
func (o *GetToken) GetTempToken() string {
	if o == nil || IsNil(o.TempToken) {
		var ret string
		return ret
	}
	return *o.TempToken
}

// GetTempTokenOk returns a tuple with the TempToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken) GetTempTokenOk() (*string, bool) {
	if o == nil || IsNil(o.TempToken) {
		return nil, false
	}
	return o.TempToken, true
}

// HasTempToken returns a boolean if a field has been set.
func (o *GetToken) HasTempToken() bool {
	if o != nil && !IsNil(o.TempToken) {
		return true
	}

	return false
}

// SetTempToken gets a reference to the given string and assigns it to the TempToken field.
func (o *GetToken) SetTempToken(v string) {
	o.TempToken = &v
}

func (o GetToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !IsNil(o.RefreshTokenExpireTime) {
		toSerialize["refreshTokenExpireTime"] = o.RefreshTokenExpireTime
	}
	if !IsNil(o.TempToken) {
		toSerialize["tempToken"] = o.TempToken
	}
	return toSerialize, nil
}

type NullableGetToken struct {
	value *GetToken
	isSet bool
}

func (v NullableGetToken) Get() *GetToken {
	return v.value
}

func (v *NullableGetToken) Set(val *GetToken) {
	v.value = val
	v.isSet = true
}

func (v NullableGetToken) IsSet() bool {
	return v.isSet
}

func (v *NullableGetToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetToken(val *GetToken) *NullableGetToken {
	return &NullableGetToken{value: val, isSet: true}
}

func (v NullableGetToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


