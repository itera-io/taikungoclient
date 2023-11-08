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

// checks if the MonitoringCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringCredentialsListDto{}

// MonitoringCredentialsListDto struct for MonitoringCredentialsListDto
type MonitoringCredentialsListDto struct {
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	PrometheusUrl NullableString `json:"prometheusUrl,omitempty"`
	LokiUrl NullableString `json:"lokiUrl,omitempty"`
	AlertManagerUrl NullableString `json:"alertManagerUrl,omitempty"`
}

// NewMonitoringCredentialsListDto instantiates a new MonitoringCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringCredentialsListDto() *MonitoringCredentialsListDto {
	this := MonitoringCredentialsListDto{}
	return &this
}

// NewMonitoringCredentialsListDtoWithDefaults instantiates a new MonitoringCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringCredentialsListDtoWithDefaults() *MonitoringCredentialsListDto {
	this := MonitoringCredentialsListDto{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringCredentialsListDto) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringCredentialsListDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *MonitoringCredentialsListDto) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *MonitoringCredentialsListDto) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *MonitoringCredentialsListDto) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *MonitoringCredentialsListDto) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringCredentialsListDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringCredentialsListDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MonitoringCredentialsListDto) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *MonitoringCredentialsListDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *MonitoringCredentialsListDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *MonitoringCredentialsListDto) UnsetPassword() {
	o.Password.Unset()
}

// GetPrometheusUrl returns the PrometheusUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringCredentialsListDto) GetPrometheusUrl() string {
	if o == nil || IsNil(o.PrometheusUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PrometheusUrl.Get()
}

// GetPrometheusUrlOk returns a tuple with the PrometheusUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringCredentialsListDto) GetPrometheusUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrometheusUrl.Get(), o.PrometheusUrl.IsSet()
}

// HasPrometheusUrl returns a boolean if a field has been set.
func (o *MonitoringCredentialsListDto) HasPrometheusUrl() bool {
	if o != nil && o.PrometheusUrl.IsSet() {
		return true
	}

	return false
}

// SetPrometheusUrl gets a reference to the given NullableString and assigns it to the PrometheusUrl field.
func (o *MonitoringCredentialsListDto) SetPrometheusUrl(v string) {
	o.PrometheusUrl.Set(&v)
}
// SetPrometheusUrlNil sets the value for PrometheusUrl to be an explicit nil
func (o *MonitoringCredentialsListDto) SetPrometheusUrlNil() {
	o.PrometheusUrl.Set(nil)
}

// UnsetPrometheusUrl ensures that no value is present for PrometheusUrl, not even an explicit nil
func (o *MonitoringCredentialsListDto) UnsetPrometheusUrl() {
	o.PrometheusUrl.Unset()
}

// GetLokiUrl returns the LokiUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringCredentialsListDto) GetLokiUrl() string {
	if o == nil || IsNil(o.LokiUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LokiUrl.Get()
}

// GetLokiUrlOk returns a tuple with the LokiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringCredentialsListDto) GetLokiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LokiUrl.Get(), o.LokiUrl.IsSet()
}

// HasLokiUrl returns a boolean if a field has been set.
func (o *MonitoringCredentialsListDto) HasLokiUrl() bool {
	if o != nil && o.LokiUrl.IsSet() {
		return true
	}

	return false
}

// SetLokiUrl gets a reference to the given NullableString and assigns it to the LokiUrl field.
func (o *MonitoringCredentialsListDto) SetLokiUrl(v string) {
	o.LokiUrl.Set(&v)
}
// SetLokiUrlNil sets the value for LokiUrl to be an explicit nil
func (o *MonitoringCredentialsListDto) SetLokiUrlNil() {
	o.LokiUrl.Set(nil)
}

// UnsetLokiUrl ensures that no value is present for LokiUrl, not even an explicit nil
func (o *MonitoringCredentialsListDto) UnsetLokiUrl() {
	o.LokiUrl.Unset()
}

// GetAlertManagerUrl returns the AlertManagerUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitoringCredentialsListDto) GetAlertManagerUrl() string {
	if o == nil || IsNil(o.AlertManagerUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AlertManagerUrl.Get()
}

// GetAlertManagerUrlOk returns a tuple with the AlertManagerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MonitoringCredentialsListDto) GetAlertManagerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertManagerUrl.Get(), o.AlertManagerUrl.IsSet()
}

// HasAlertManagerUrl returns a boolean if a field has been set.
func (o *MonitoringCredentialsListDto) HasAlertManagerUrl() bool {
	if o != nil && o.AlertManagerUrl.IsSet() {
		return true
	}

	return false
}

// SetAlertManagerUrl gets a reference to the given NullableString and assigns it to the AlertManagerUrl field.
func (o *MonitoringCredentialsListDto) SetAlertManagerUrl(v string) {
	o.AlertManagerUrl.Set(&v)
}
// SetAlertManagerUrlNil sets the value for AlertManagerUrl to be an explicit nil
func (o *MonitoringCredentialsListDto) SetAlertManagerUrlNil() {
	o.AlertManagerUrl.Set(nil)
}

// UnsetAlertManagerUrl ensures that no value is present for AlertManagerUrl, not even an explicit nil
func (o *MonitoringCredentialsListDto) UnsetAlertManagerUrl() {
	o.AlertManagerUrl.Unset()
}

func (o MonitoringCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringCredentialsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PrometheusUrl.IsSet() {
		toSerialize["prometheusUrl"] = o.PrometheusUrl.Get()
	}
	if o.LokiUrl.IsSet() {
		toSerialize["lokiUrl"] = o.LokiUrl.Get()
	}
	if o.AlertManagerUrl.IsSet() {
		toSerialize["alertManagerUrl"] = o.AlertManagerUrl.Get()
	}
	return toSerialize, nil
}

type NullableMonitoringCredentialsListDto struct {
	value *MonitoringCredentialsListDto
	isSet bool
}

func (v NullableMonitoringCredentialsListDto) Get() *MonitoringCredentialsListDto {
	return v.value
}

func (v *NullableMonitoringCredentialsListDto) Set(val *MonitoringCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringCredentialsListDto(val *MonitoringCredentialsListDto) *NullableMonitoringCredentialsListDto {
	return &NullableMonitoringCredentialsListDto{value: val, isSet: true}
}

func (v NullableMonitoringCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


