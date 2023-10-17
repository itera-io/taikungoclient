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

// checks if the CreateAccessProfileCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccessProfileCommand{}

// CreateAccessProfileCommand struct for CreateAccessProfileCommand
type CreateAccessProfileCommand struct {
	Name           NullableString         `json:"name,omitempty"`
	HttpProxy      NullableString         `json:"httpProxy,omitempty"`
	OrganizationId NullableInt32          `json:"organizationId,omitempty"`
	SshUsers       []SshUserCreateDto     `json:"sshUsers,omitempty"`
	DnsServers     []DnsServerCreateDto   `json:"dnsServers,omitempty"`
	NtpServers     []NtpServerCreateDto   `json:"ntpServers,omitempty"`
	AllowedHosts   []AllowedHostCreateDto `json:"allowedHosts,omitempty"`
}

// NewCreateAccessProfileCommand instantiates a new CreateAccessProfileCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessProfileCommand() *CreateAccessProfileCommand {
	this := CreateAccessProfileCommand{}
	return &this
}

// NewCreateAccessProfileCommandWithDefaults instantiates a new CreateAccessProfileCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessProfileCommandWithDefaults() *CreateAccessProfileCommand {
	this := CreateAccessProfileCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateAccessProfileCommand) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateAccessProfileCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateAccessProfileCommand) UnsetName() {
	o.Name.Unset()
}

// GetHttpProxy returns the HttpProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetHttpProxy() string {
	if o == nil || IsNil(o.HttpProxy.Get()) {
		var ret string
		return ret
	}
	return *o.HttpProxy.Get()
}

// GetHttpProxyOk returns a tuple with the HttpProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetHttpProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpProxy.Get(), o.HttpProxy.IsSet()
}

// HasHttpProxy returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasHttpProxy() bool {
	if o != nil && o.HttpProxy.IsSet() {
		return true
	}

	return false
}

// SetHttpProxy gets a reference to the given NullableString and assigns it to the HttpProxy field.
func (o *CreateAccessProfileCommand) SetHttpProxy(v string) {
	o.HttpProxy.Set(&v)
}

// SetHttpProxyNil sets the value for HttpProxy to be an explicit nil
func (o *CreateAccessProfileCommand) SetHttpProxyNil() {
	o.HttpProxy.Set(nil)
}

// UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
func (o *CreateAccessProfileCommand) UnsetHttpProxy() {
	o.HttpProxy.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateAccessProfileCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}

// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateAccessProfileCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateAccessProfileCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetSshUsers returns the SshUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetSshUsers() []SshUserCreateDto {
	if o == nil {
		var ret []SshUserCreateDto
		return ret
	}
	return o.SshUsers
}

// GetSshUsersOk returns a tuple with the SshUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetSshUsersOk() ([]SshUserCreateDto, bool) {
	if o == nil || IsNil(o.SshUsers) {
		return nil, false
	}
	return o.SshUsers, true
}

// HasSshUsers returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasSshUsers() bool {
	if o != nil && IsNil(o.SshUsers) {
		return true
	}

	return false
}

// SetSshUsers gets a reference to the given []SshUserCreateDto and assigns it to the SshUsers field.
func (o *CreateAccessProfileCommand) SetSshUsers(v []SshUserCreateDto) {
	o.SshUsers = v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetDnsServers() []DnsServerCreateDto {
	if o == nil {
		var ret []DnsServerCreateDto
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetDnsServersOk() ([]DnsServerCreateDto, bool) {
	if o == nil || IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasDnsServers() bool {
	if o != nil && IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []DnsServerCreateDto and assigns it to the DnsServers field.
func (o *CreateAccessProfileCommand) SetDnsServers(v []DnsServerCreateDto) {
	o.DnsServers = v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetNtpServers() []NtpServerCreateDto {
	if o == nil {
		var ret []NtpServerCreateDto
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetNtpServersOk() ([]NtpServerCreateDto, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasNtpServers() bool {
	if o != nil && IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []NtpServerCreateDto and assigns it to the NtpServers field.
func (o *CreateAccessProfileCommand) SetNtpServers(v []NtpServerCreateDto) {
	o.NtpServers = v
}

// GetAllowedHosts returns the AllowedHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccessProfileCommand) GetAllowedHosts() []AllowedHostCreateDto {
	if o == nil {
		var ret []AllowedHostCreateDto
		return ret
	}
	return o.AllowedHosts
}

// GetAllowedHostsOk returns a tuple with the AllowedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccessProfileCommand) GetAllowedHostsOk() ([]AllowedHostCreateDto, bool) {
	if o == nil || IsNil(o.AllowedHosts) {
		return nil, false
	}
	return o.AllowedHosts, true
}

// HasAllowedHosts returns a boolean if a field has been set.
func (o *CreateAccessProfileCommand) HasAllowedHosts() bool {
	if o != nil && IsNil(o.AllowedHosts) {
		return true
	}

	return false
}

// SetAllowedHosts gets a reference to the given []AllowedHostCreateDto and assigns it to the AllowedHosts field.
func (o *CreateAccessProfileCommand) SetAllowedHosts(v []AllowedHostCreateDto) {
	o.AllowedHosts = v
}

func (o CreateAccessProfileCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccessProfileCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.HttpProxy.IsSet() {
		toSerialize["httpProxy"] = o.HttpProxy.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.SshUsers != nil {
		toSerialize["sshUsers"] = o.SshUsers
	}
	if o.DnsServers != nil {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if o.NtpServers != nil {
		toSerialize["ntpServers"] = o.NtpServers
	}
	if o.AllowedHosts != nil {
		toSerialize["allowedHosts"] = o.AllowedHosts
	}
	return toSerialize, nil
}

type NullableCreateAccessProfileCommand struct {
	value *CreateAccessProfileCommand
	isSet bool
}

func (v NullableCreateAccessProfileCommand) Get() *CreateAccessProfileCommand {
	return v.value
}

func (v *NullableCreateAccessProfileCommand) Set(val *CreateAccessProfileCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessProfileCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessProfileCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessProfileCommand(val *CreateAccessProfileCommand) *NullableCreateAccessProfileCommand {
	return &NullableCreateAccessProfileCommand{value: val, isSet: true}
}

func (v NullableCreateAccessProfileCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessProfileCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
