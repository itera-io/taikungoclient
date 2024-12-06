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

// checks if the OpenshiftListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenshiftListDto{}

// OpenshiftListDto struct for OpenshiftListDto
type OpenshiftListDto struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	BaseDomain string `json:"baseDomain"`
	StorageClass string `json:"storageClass"`
	OrganizationId int32 `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	Projects []CommonDropdownDto `json:"projects"`
	ProjectCount int32 `json:"projectCount"`
	IsLocked bool `json:"isLocked"`
	CreatedBy NullableString `json:"createdBy"`
	LastModified NullableString `json:"lastModified"`
	LastModifiedBy NullableString `json:"lastModifiedBy"`
	IsDefault bool `json:"isDefault"`
	CreatedAt NullableString `json:"createdAt"`
	ContinentName string `json:"continentName"`
}

type _OpenshiftListDto OpenshiftListDto

// NewOpenshiftListDto instantiates a new OpenshiftListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenshiftListDto(id int32, name string, baseDomain string, storageClass string, organizationId int32, organizationName string, projects []CommonDropdownDto, projectCount int32, isLocked bool, createdBy NullableString, lastModified NullableString, lastModifiedBy NullableString, isDefault bool, createdAt NullableString, continentName string) *OpenshiftListDto {
	this := OpenshiftListDto{}
	this.Id = id
	this.Name = name
	this.BaseDomain = baseDomain
	this.StorageClass = storageClass
	this.OrganizationId = organizationId
	this.OrganizationName = organizationName
	this.Projects = projects
	this.ProjectCount = projectCount
	this.IsLocked = isLocked
	this.CreatedBy = createdBy
	this.LastModified = lastModified
	this.LastModifiedBy = lastModifiedBy
	this.IsDefault = isDefault
	this.CreatedAt = createdAt
	this.ContinentName = continentName
	return &this
}

// NewOpenshiftListDtoWithDefaults instantiates a new OpenshiftListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenshiftListDtoWithDefaults() *OpenshiftListDto {
	this := OpenshiftListDto{}
	return &this
}

// GetId returns the Id field value
func (o *OpenshiftListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenshiftListDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OpenshiftListDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenshiftListDto) SetName(v string) {
	o.Name = v
}

// GetBaseDomain returns the BaseDomain field value
func (o *OpenshiftListDto) GetBaseDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDomain
}

// GetBaseDomainOk returns a tuple with the BaseDomain field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetBaseDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDomain, true
}

// SetBaseDomain sets field value
func (o *OpenshiftListDto) SetBaseDomain(v string) {
	o.BaseDomain = v
}

// GetStorageClass returns the StorageClass field value
func (o *OpenshiftListDto) GetStorageClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value
func (o *OpenshiftListDto) SetStorageClass(v string) {
	o.StorageClass = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *OpenshiftListDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *OpenshiftListDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *OpenshiftListDto) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *OpenshiftListDto) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetProjects returns the Projects field value
func (o *OpenshiftListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *OpenshiftListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetProjectCount returns the ProjectCount field value
func (o *OpenshiftListDto) GetProjectCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetProjectCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectCount, true
}

// SetProjectCount sets field value
func (o *OpenshiftListDto) SetProjectCount(v int32) {
	o.ProjectCount = v
}

// GetIsLocked returns the IsLocked field value
func (o *OpenshiftListDto) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *OpenshiftListDto) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenshiftListDto) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// SetCreatedBy sets field value
func (o *OpenshiftListDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenshiftListDto) GetLastModified() string {
	if o == nil || o.LastModified.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// SetLastModified sets field value
func (o *OpenshiftListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}

// GetLastModifiedBy returns the LastModifiedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenshiftListDto) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// SetLastModifiedBy sets field value
func (o *OpenshiftListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}

// GetIsDefault returns the IsDefault field value
func (o *OpenshiftListDto) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *OpenshiftListDto) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenshiftListDto) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenshiftListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *OpenshiftListDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetContinentName returns the ContinentName field value
func (o *OpenshiftListDto) GetContinentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContinentName
}

// GetContinentNameOk returns a tuple with the ContinentName field value
// and a boolean to check if the value has been set.
func (o *OpenshiftListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinentName, true
}

// SetContinentName sets field value
func (o *OpenshiftListDto) SetContinentName(v string) {
	o.ContinentName = v
}

func (o OpenshiftListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenshiftListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["baseDomain"] = o.BaseDomain
	toSerialize["storageClass"] = o.StorageClass
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["organizationName"] = o.OrganizationName
	toSerialize["projects"] = o.Projects
	toSerialize["projectCount"] = o.ProjectCount
	toSerialize["isLocked"] = o.IsLocked
	toSerialize["createdBy"] = o.CreatedBy.Get()
	toSerialize["lastModified"] = o.LastModified.Get()
	toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	toSerialize["isDefault"] = o.IsDefault
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["continentName"] = o.ContinentName
	return toSerialize, nil
}

func (o *OpenshiftListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"baseDomain",
		"storageClass",
		"organizationId",
		"organizationName",
		"projects",
		"projectCount",
		"isLocked",
		"createdBy",
		"lastModified",
		"lastModifiedBy",
		"isDefault",
		"createdAt",
		"continentName",
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

	varOpenshiftListDto := _OpenshiftListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpenshiftListDto)

	if err != nil {
		return err
	}

	*o = OpenshiftListDto(varOpenshiftListDto)

	return err
}

type NullableOpenshiftListDto struct {
	value *OpenshiftListDto
	isSet bool
}

func (v NullableOpenshiftListDto) Get() *OpenshiftListDto {
	return v.value
}

func (v *NullableOpenshiftListDto) Set(val *OpenshiftListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenshiftListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenshiftListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenshiftListDto(val *OpenshiftListDto) *NullableOpenshiftListDto {
	return &NullableOpenshiftListDto{value: val, isSet: true}
}

func (v NullableOpenshiftListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenshiftListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


