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

// checks if the ImportedClusterDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportedClusterDetailsDto{}

// ImportedClusterDetailsDto struct for ImportedClusterDetailsDto
type ImportedClusterDetailsDto struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	IsLocked bool `json:"isLocked"`
	AccessIp NullableString `json:"accessIp"`
	KubernetesVersion NullableString `json:"kubernetesVersion"`
	ImportClusterType ImportClusterType `json:"importClusterType"`
	OrganizationId int32 `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	CloudCredentialName NullableString `json:"cloudCredentialName,omitempty"`
	CloudCredentialId NullableInt32 `json:"cloudCredentialId,omitempty"`
	Health ProjectHealth `json:"health"`
	CloudType *CloudType `json:"cloudType,omitempty"`
	Status ProjectStatus `json:"status"`
}

type _ImportedClusterDetailsDto ImportedClusterDetailsDto

// NewImportedClusterDetailsDto instantiates a new ImportedClusterDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportedClusterDetailsDto(id int32, name string, isLocked bool, accessIp NullableString, kubernetesVersion NullableString, importClusterType ImportClusterType, organizationId int32, organizationName string, health ProjectHealth, status ProjectStatus) *ImportedClusterDetailsDto {
	this := ImportedClusterDetailsDto{}
	this.Id = id
	this.Name = name
	this.IsLocked = isLocked
	this.AccessIp = accessIp
	this.KubernetesVersion = kubernetesVersion
	this.ImportClusterType = importClusterType
	this.OrganizationId = organizationId
	this.OrganizationName = organizationName
	this.Health = health
	this.Status = status
	return &this
}

// NewImportedClusterDetailsDtoWithDefaults instantiates a new ImportedClusterDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportedClusterDetailsDtoWithDefaults() *ImportedClusterDetailsDto {
	this := ImportedClusterDetailsDto{}
	return &this
}

// GetId returns the Id field value
func (o *ImportedClusterDetailsDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImportedClusterDetailsDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ImportedClusterDetailsDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImportedClusterDetailsDto) SetName(v string) {
	o.Name = v
}

// GetIsLocked returns the IsLocked field value
func (o *ImportedClusterDetailsDto) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *ImportedClusterDetailsDto) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetAccessIp returns the AccessIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ImportedClusterDetailsDto) GetAccessIp() string {
	if o == nil || o.AccessIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccessIp.Get()
}

// GetAccessIpOk returns a tuple with the AccessIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportedClusterDetailsDto) GetAccessIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessIp.Get(), o.AccessIp.IsSet()
}

// SetAccessIp sets field value
func (o *ImportedClusterDetailsDto) SetAccessIp(v string) {
	o.AccessIp.Set(&v)
}

// GetKubernetesVersion returns the KubernetesVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ImportedClusterDetailsDto) GetKubernetesVersion() string {
	if o == nil || o.KubernetesVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.KubernetesVersion.Get()
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportedClusterDetailsDto) GetKubernetesVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KubernetesVersion.Get(), o.KubernetesVersion.IsSet()
}

// SetKubernetesVersion sets field value
func (o *ImportedClusterDetailsDto) SetKubernetesVersion(v string) {
	o.KubernetesVersion.Set(&v)
}

// GetImportClusterType returns the ImportClusterType field value
func (o *ImportedClusterDetailsDto) GetImportClusterType() ImportClusterType {
	if o == nil {
		var ret ImportClusterType
		return ret
	}

	return o.ImportClusterType
}

// GetImportClusterTypeOk returns a tuple with the ImportClusterType field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetImportClusterTypeOk() (*ImportClusterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportClusterType, true
}

// SetImportClusterType sets field value
func (o *ImportedClusterDetailsDto) SetImportClusterType(v ImportClusterType) {
	o.ImportClusterType = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *ImportedClusterDetailsDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *ImportedClusterDetailsDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *ImportedClusterDetailsDto) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *ImportedClusterDetailsDto) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetCloudCredentialName returns the CloudCredentialName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportedClusterDetailsDto) GetCloudCredentialName() string {
	if o == nil || IsNil(o.CloudCredentialName.Get()) {
		var ret string
		return ret
	}
	return *o.CloudCredentialName.Get()
}

// GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportedClusterDetailsDto) GetCloudCredentialNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudCredentialName.Get(), o.CloudCredentialName.IsSet()
}

// HasCloudCredentialName returns a boolean if a field has been set.
func (o *ImportedClusterDetailsDto) HasCloudCredentialName() bool {
	if o != nil && o.CloudCredentialName.IsSet() {
		return true
	}

	return false
}

// SetCloudCredentialName gets a reference to the given NullableString and assigns it to the CloudCredentialName field.
func (o *ImportedClusterDetailsDto) SetCloudCredentialName(v string) {
	o.CloudCredentialName.Set(&v)
}
// SetCloudCredentialNameNil sets the value for CloudCredentialName to be an explicit nil
func (o *ImportedClusterDetailsDto) SetCloudCredentialNameNil() {
	o.CloudCredentialName.Set(nil)
}

// UnsetCloudCredentialName ensures that no value is present for CloudCredentialName, not even an explicit nil
func (o *ImportedClusterDetailsDto) UnsetCloudCredentialName() {
	o.CloudCredentialName.Unset()
}

// GetCloudCredentialId returns the CloudCredentialId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportedClusterDetailsDto) GetCloudCredentialId() int32 {
	if o == nil || IsNil(o.CloudCredentialId.Get()) {
		var ret int32
		return ret
	}
	return *o.CloudCredentialId.Get()
}

// GetCloudCredentialIdOk returns a tuple with the CloudCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportedClusterDetailsDto) GetCloudCredentialIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudCredentialId.Get(), o.CloudCredentialId.IsSet()
}

// HasCloudCredentialId returns a boolean if a field has been set.
func (o *ImportedClusterDetailsDto) HasCloudCredentialId() bool {
	if o != nil && o.CloudCredentialId.IsSet() {
		return true
	}

	return false
}

// SetCloudCredentialId gets a reference to the given NullableInt32 and assigns it to the CloudCredentialId field.
func (o *ImportedClusterDetailsDto) SetCloudCredentialId(v int32) {
	o.CloudCredentialId.Set(&v)
}
// SetCloudCredentialIdNil sets the value for CloudCredentialId to be an explicit nil
func (o *ImportedClusterDetailsDto) SetCloudCredentialIdNil() {
	o.CloudCredentialId.Set(nil)
}

// UnsetCloudCredentialId ensures that no value is present for CloudCredentialId, not even an explicit nil
func (o *ImportedClusterDetailsDto) UnsetCloudCredentialId() {
	o.CloudCredentialId.Unset()
}

// GetHealth returns the Health field value
func (o *ImportedClusterDetailsDto) GetHealth() ProjectHealth {
	if o == nil {
		var ret ProjectHealth
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetHealthOk() (*ProjectHealth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *ImportedClusterDetailsDto) SetHealth(v ProjectHealth) {
	o.Health = v
}

// GetCloudType returns the CloudType field value if set, zero value otherwise.
func (o *ImportedClusterDetailsDto) GetCloudType() CloudType {
	if o == nil || IsNil(o.CloudType) {
		var ret CloudType
		return ret
	}
	return *o.CloudType
}

// GetCloudTypeOk returns a tuple with the CloudType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetCloudTypeOk() (*CloudType, bool) {
	if o == nil || IsNil(o.CloudType) {
		return nil, false
	}
	return o.CloudType, true
}

// HasCloudType returns a boolean if a field has been set.
func (o *ImportedClusterDetailsDto) HasCloudType() bool {
	if o != nil && !IsNil(o.CloudType) {
		return true
	}

	return false
}

// SetCloudType gets a reference to the given CloudType and assigns it to the CloudType field.
func (o *ImportedClusterDetailsDto) SetCloudType(v CloudType) {
	o.CloudType = &v
}

// GetStatus returns the Status field value
func (o *ImportedClusterDetailsDto) GetStatus() ProjectStatus {
	if o == nil {
		var ret ProjectStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImportedClusterDetailsDto) GetStatusOk() (*ProjectStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImportedClusterDetailsDto) SetStatus(v ProjectStatus) {
	o.Status = v
}

func (o ImportedClusterDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportedClusterDetailsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["isLocked"] = o.IsLocked
	toSerialize["accessIp"] = o.AccessIp.Get()
	toSerialize["kubernetesVersion"] = o.KubernetesVersion.Get()
	toSerialize["importClusterType"] = o.ImportClusterType
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["organizationName"] = o.OrganizationName
	if o.CloudCredentialName.IsSet() {
		toSerialize["cloudCredentialName"] = o.CloudCredentialName.Get()
	}
	if o.CloudCredentialId.IsSet() {
		toSerialize["cloudCredentialId"] = o.CloudCredentialId.Get()
	}
	toSerialize["health"] = o.Health
	if !IsNil(o.CloudType) {
		toSerialize["cloudType"] = o.CloudType
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ImportedClusterDetailsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"isLocked",
		"accessIp",
		"kubernetesVersion",
		"importClusterType",
		"organizationId",
		"organizationName",
		"health",
		"status",
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

	varImportedClusterDetailsDto := _ImportedClusterDetailsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportedClusterDetailsDto)

	if err != nil {
		return err
	}

	*o = ImportedClusterDetailsDto(varImportedClusterDetailsDto)

	return err
}

type NullableImportedClusterDetailsDto struct {
	value *ImportedClusterDetailsDto
	isSet bool
}

func (v NullableImportedClusterDetailsDto) Get() *ImportedClusterDetailsDto {
	return v.value
}

func (v *NullableImportedClusterDetailsDto) Set(val *ImportedClusterDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableImportedClusterDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableImportedClusterDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportedClusterDetailsDto(val *ImportedClusterDetailsDto) *NullableImportedClusterDetailsDto {
	return &NullableImportedClusterDetailsDto{value: val, isSet: true}
}

func (v NullableImportedClusterDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportedClusterDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


