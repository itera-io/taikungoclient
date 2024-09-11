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

// checks if the StandAloneProfilesListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandAloneProfilesListDto{}

// StandAloneProfilesListDto struct for StandAloneProfilesListDto
type StandAloneProfilesListDto struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
	IsLocked *bool `json:"isLocked,omitempty"`
	StandaloneVms []StandAloneVmSmallDetailDto `json:"standaloneVms,omitempty"`
	OrganizationId *int32 `json:"organizationId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	PartnerLogo *string `json:"partnerLogo,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewStandAloneProfilesListDto instantiates a new StandAloneProfilesListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandAloneProfilesListDto() *StandAloneProfilesListDto {
	this := StandAloneProfilesListDto{}
	return &this
}

// NewStandAloneProfilesListDtoWithDefaults instantiates a new StandAloneProfilesListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandAloneProfilesListDtoWithDefaults() *StandAloneProfilesListDto {
	this := StandAloneProfilesListDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandAloneProfilesListDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StandAloneProfilesListDto) SetName(v string) {
	o.Name = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *StandAloneProfilesListDto) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetIsLocked() bool {
	if o == nil || IsNil(o.IsLocked) {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocked) {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasIsLocked() bool {
	if o != nil && !IsNil(o.IsLocked) {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *StandAloneProfilesListDto) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetStandaloneVms returns the StandaloneVms field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetStandaloneVms() []StandAloneVmSmallDetailDto {
	if o == nil || IsNil(o.StandaloneVms) {
		var ret []StandAloneVmSmallDetailDto
		return ret
	}
	return o.StandaloneVms
}

// GetStandaloneVmsOk returns a tuple with the StandaloneVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetStandaloneVmsOk() ([]StandAloneVmSmallDetailDto, bool) {
	if o == nil || IsNil(o.StandaloneVms) {
		return nil, false
	}
	return o.StandaloneVms, true
}

// HasStandaloneVms returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasStandaloneVms() bool {
	if o != nil && !IsNil(o.StandaloneVms) {
		return true
	}

	return false
}

// SetStandaloneVms gets a reference to the given []StandAloneVmSmallDetailDto and assigns it to the StandaloneVms field.
func (o *StandAloneProfilesListDto) SetStandaloneVms(v []StandAloneVmSmallDetailDto) {
	o.StandaloneVms = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *StandAloneProfilesListDto) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *StandAloneProfilesListDto) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetPartnerLogo returns the PartnerLogo field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetPartnerLogo() string {
	if o == nil || IsNil(o.PartnerLogo) {
		var ret string
		return ret
	}
	return *o.PartnerLogo
}

// GetPartnerLogoOk returns a tuple with the PartnerLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetPartnerLogoOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerLogo) {
		return nil, false
	}
	return o.PartnerLogo, true
}

// HasPartnerLogo returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasPartnerLogo() bool {
	if o != nil && !IsNil(o.PartnerLogo) {
		return true
	}

	return false
}

// SetPartnerLogo gets a reference to the given string and assigns it to the PartnerLogo field.
func (o *StandAloneProfilesListDto) SetPartnerLogo(v string) {
	o.PartnerLogo = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StandAloneProfilesListDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandAloneProfilesListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StandAloneProfilesListDto) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *StandAloneProfilesListDto) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o StandAloneProfilesListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandAloneProfilesListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !IsNil(o.IsLocked) {
		toSerialize["isLocked"] = o.IsLocked
	}
	if !IsNil(o.StandaloneVms) {
		toSerialize["standaloneVms"] = o.StandaloneVms
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.PartnerLogo) {
		toSerialize["partnerLogo"] = o.PartnerLogo
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableStandAloneProfilesListDto struct {
	value *StandAloneProfilesListDto
	isSet bool
}

func (v NullableStandAloneProfilesListDto) Get() *StandAloneProfilesListDto {
	return v.value
}

func (v *NullableStandAloneProfilesListDto) Set(val *StandAloneProfilesListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandAloneProfilesListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandAloneProfilesListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandAloneProfilesListDto(val *StandAloneProfilesListDto) *NullableStandAloneProfilesListDto {
	return &NullableStandAloneProfilesListDto{value: val, isSet: true}
}

func (v NullableStandAloneProfilesListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandAloneProfilesListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


