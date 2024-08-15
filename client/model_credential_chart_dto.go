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

// checks if the CredentialChartDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialChartDto{}

// CredentialChartDto struct for CredentialChartDto
type CredentialChartDto struct {
	Aws *int32 `json:"aws,omitempty"`
	Azure *int32 `json:"azure,omitempty"`
	Openstack *int32 `json:"openstack,omitempty"`
	Google *int32 `json:"google,omitempty"`
	Tanzu *int32 `json:"tanzu,omitempty"`
	Proxmox *int32 `json:"proxmox,omitempty"`
	Openshift *int32 `json:"openshift,omitempty"`
	Vsphere *int32 `json:"vsphere,omitempty"`
	Zadara *int32 `json:"zadara,omitempty"`
	Zededa *int32 `json:"zededa,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewCredentialChartDto instantiates a new CredentialChartDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialChartDto() *CredentialChartDto {
	this := CredentialChartDto{}
	return &this
}

// NewCredentialChartDtoWithDefaults instantiates a new CredentialChartDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialChartDtoWithDefaults() *CredentialChartDto {
	this := CredentialChartDto{}
	return &this
}

// GetAws returns the Aws field value if set, zero value otherwise.
func (o *CredentialChartDto) GetAws() int32 {
	if o == nil || IsNil(o.Aws) {
		var ret int32
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetAwsOk() (*int32, bool) {
	if o == nil || IsNil(o.Aws) {
		return nil, false
	}
	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *CredentialChartDto) HasAws() bool {
	if o != nil && !IsNil(o.Aws) {
		return true
	}

	return false
}

// SetAws gets a reference to the given int32 and assigns it to the Aws field.
func (o *CredentialChartDto) SetAws(v int32) {
	o.Aws = &v
}

// GetAzure returns the Azure field value if set, zero value otherwise.
func (o *CredentialChartDto) GetAzure() int32 {
	if o == nil || IsNil(o.Azure) {
		var ret int32
		return ret
	}
	return *o.Azure
}

// GetAzureOk returns a tuple with the Azure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetAzureOk() (*int32, bool) {
	if o == nil || IsNil(o.Azure) {
		return nil, false
	}
	return o.Azure, true
}

// HasAzure returns a boolean if a field has been set.
func (o *CredentialChartDto) HasAzure() bool {
	if o != nil && !IsNil(o.Azure) {
		return true
	}

	return false
}

// SetAzure gets a reference to the given int32 and assigns it to the Azure field.
func (o *CredentialChartDto) SetAzure(v int32) {
	o.Azure = &v
}

// GetOpenstack returns the Openstack field value if set, zero value otherwise.
func (o *CredentialChartDto) GetOpenstack() int32 {
	if o == nil || IsNil(o.Openstack) {
		var ret int32
		return ret
	}
	return *o.Openstack
}

// GetOpenstackOk returns a tuple with the Openstack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetOpenstackOk() (*int32, bool) {
	if o == nil || IsNil(o.Openstack) {
		return nil, false
	}
	return o.Openstack, true
}

// HasOpenstack returns a boolean if a field has been set.
func (o *CredentialChartDto) HasOpenstack() bool {
	if o != nil && !IsNil(o.Openstack) {
		return true
	}

	return false
}

// SetOpenstack gets a reference to the given int32 and assigns it to the Openstack field.
func (o *CredentialChartDto) SetOpenstack(v int32) {
	o.Openstack = &v
}

// GetGoogle returns the Google field value if set, zero value otherwise.
func (o *CredentialChartDto) GetGoogle() int32 {
	if o == nil || IsNil(o.Google) {
		var ret int32
		return ret
	}
	return *o.Google
}

// GetGoogleOk returns a tuple with the Google field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetGoogleOk() (*int32, bool) {
	if o == nil || IsNil(o.Google) {
		return nil, false
	}
	return o.Google, true
}

// HasGoogle returns a boolean if a field has been set.
func (o *CredentialChartDto) HasGoogle() bool {
	if o != nil && !IsNil(o.Google) {
		return true
	}

	return false
}

// SetGoogle gets a reference to the given int32 and assigns it to the Google field.
func (o *CredentialChartDto) SetGoogle(v int32) {
	o.Google = &v
}

// GetTanzu returns the Tanzu field value if set, zero value otherwise.
func (o *CredentialChartDto) GetTanzu() int32 {
	if o == nil || IsNil(o.Tanzu) {
		var ret int32
		return ret
	}
	return *o.Tanzu
}

// GetTanzuOk returns a tuple with the Tanzu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetTanzuOk() (*int32, bool) {
	if o == nil || IsNil(o.Tanzu) {
		return nil, false
	}
	return o.Tanzu, true
}

// HasTanzu returns a boolean if a field has been set.
func (o *CredentialChartDto) HasTanzu() bool {
	if o != nil && !IsNil(o.Tanzu) {
		return true
	}

	return false
}

// SetTanzu gets a reference to the given int32 and assigns it to the Tanzu field.
func (o *CredentialChartDto) SetTanzu(v int32) {
	o.Tanzu = &v
}

// GetProxmox returns the Proxmox field value if set, zero value otherwise.
func (o *CredentialChartDto) GetProxmox() int32 {
	if o == nil || IsNil(o.Proxmox) {
		var ret int32
		return ret
	}
	return *o.Proxmox
}

// GetProxmoxOk returns a tuple with the Proxmox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetProxmoxOk() (*int32, bool) {
	if o == nil || IsNil(o.Proxmox) {
		return nil, false
	}
	return o.Proxmox, true
}

// HasProxmox returns a boolean if a field has been set.
func (o *CredentialChartDto) HasProxmox() bool {
	if o != nil && !IsNil(o.Proxmox) {
		return true
	}

	return false
}

// SetProxmox gets a reference to the given int32 and assigns it to the Proxmox field.
func (o *CredentialChartDto) SetProxmox(v int32) {
	o.Proxmox = &v
}

// GetOpenshift returns the Openshift field value if set, zero value otherwise.
func (o *CredentialChartDto) GetOpenshift() int32 {
	if o == nil || IsNil(o.Openshift) {
		var ret int32
		return ret
	}
	return *o.Openshift
}

// GetOpenshiftOk returns a tuple with the Openshift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetOpenshiftOk() (*int32, bool) {
	if o == nil || IsNil(o.Openshift) {
		return nil, false
	}
	return o.Openshift, true
}

// HasOpenshift returns a boolean if a field has been set.
func (o *CredentialChartDto) HasOpenshift() bool {
	if o != nil && !IsNil(o.Openshift) {
		return true
	}

	return false
}

// SetOpenshift gets a reference to the given int32 and assigns it to the Openshift field.
func (o *CredentialChartDto) SetOpenshift(v int32) {
	o.Openshift = &v
}

// GetVsphere returns the Vsphere field value if set, zero value otherwise.
func (o *CredentialChartDto) GetVsphere() int32 {
	if o == nil || IsNil(o.Vsphere) {
		var ret int32
		return ret
	}
	return *o.Vsphere
}

// GetVsphereOk returns a tuple with the Vsphere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetVsphereOk() (*int32, bool) {
	if o == nil || IsNil(o.Vsphere) {
		return nil, false
	}
	return o.Vsphere, true
}

// HasVsphere returns a boolean if a field has been set.
func (o *CredentialChartDto) HasVsphere() bool {
	if o != nil && !IsNil(o.Vsphere) {
		return true
	}

	return false
}

// SetVsphere gets a reference to the given int32 and assigns it to the Vsphere field.
func (o *CredentialChartDto) SetVsphere(v int32) {
	o.Vsphere = &v
}

// GetZadara returns the Zadara field value if set, zero value otherwise.
func (o *CredentialChartDto) GetZadara() int32 {
	if o == nil || IsNil(o.Zadara) {
		var ret int32
		return ret
	}
	return *o.Zadara
}

// GetZadaraOk returns a tuple with the Zadara field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetZadaraOk() (*int32, bool) {
	if o == nil || IsNil(o.Zadara) {
		return nil, false
	}
	return o.Zadara, true
}

// HasZadara returns a boolean if a field has been set.
func (o *CredentialChartDto) HasZadara() bool {
	if o != nil && !IsNil(o.Zadara) {
		return true
	}

	return false
}

// SetZadara gets a reference to the given int32 and assigns it to the Zadara field.
func (o *CredentialChartDto) SetZadara(v int32) {
	o.Zadara = &v
}

// GetZededa returns the Zededa field value if set, zero value otherwise.
func (o *CredentialChartDto) GetZededa() int32 {
	if o == nil || IsNil(o.Zededa) {
		var ret int32
		return ret
	}
	return *o.Zededa
}

// GetZededaOk returns a tuple with the Zededa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetZededaOk() (*int32, bool) {
	if o == nil || IsNil(o.Zededa) {
		return nil, false
	}
	return o.Zededa, true
}

// HasZededa returns a boolean if a field has been set.
func (o *CredentialChartDto) HasZededa() bool {
	if o != nil && !IsNil(o.Zededa) {
		return true
	}

	return false
}

// SetZededa gets a reference to the given int32 and assigns it to the Zededa field.
func (o *CredentialChartDto) SetZededa(v int32) {
	o.Zededa = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CredentialChartDto) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialChartDto) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CredentialChartDto) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *CredentialChartDto) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o CredentialChartDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialChartDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aws) {
		toSerialize["aws"] = o.Aws
	}
	if !IsNil(o.Azure) {
		toSerialize["azure"] = o.Azure
	}
	if !IsNil(o.Openstack) {
		toSerialize["openstack"] = o.Openstack
	}
	if !IsNil(o.Google) {
		toSerialize["google"] = o.Google
	}
	if !IsNil(o.Tanzu) {
		toSerialize["tanzu"] = o.Tanzu
	}
	if !IsNil(o.Proxmox) {
		toSerialize["proxmox"] = o.Proxmox
	}
	if !IsNil(o.Openshift) {
		toSerialize["openshift"] = o.Openshift
	}
	if !IsNil(o.Vsphere) {
		toSerialize["vsphere"] = o.Vsphere
	}
	if !IsNil(o.Zadara) {
		toSerialize["zadara"] = o.Zadara
	}
	if !IsNil(o.Zededa) {
		toSerialize["zededa"] = o.Zededa
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableCredentialChartDto struct {
	value *CredentialChartDto
	isSet bool
}

func (v NullableCredentialChartDto) Get() *CredentialChartDto {
	return v.value
}

func (v *NullableCredentialChartDto) Set(val *CredentialChartDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialChartDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialChartDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialChartDto(val *CredentialChartDto) *NullableCredentialChartDto {
	return &NullableCredentialChartDto{value: val, isSet: true}
}

func (v NullableCredentialChartDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialChartDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


