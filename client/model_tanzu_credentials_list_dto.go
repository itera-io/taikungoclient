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

// checks if the TanzuCredentialsListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TanzuCredentialsListDto{}

// TanzuCredentialsListDto struct for TanzuCredentialsListDto
type TanzuCredentialsListDto struct {
	Id int32 `json:"id"`
	ProjectCount int32 `json:"projectCount"`
	IsLocked bool `json:"isLocked"`
	Name string `json:"name"`
	Username string `json:"username"`
	Url string `json:"url"`
	VolumeType string `json:"volumeType"`
	Namespace string `json:"namespace"`
	Port NullableInt32 `json:"port"`
	Projects []CommonDropdownDto `json:"projects"`
	CreatedBy string `json:"createdBy"`
	CreatedAt string `json:"createdAt"`
	LastModified NullableString `json:"lastModified"`
	LastModifiedBy NullableString `json:"lastModifiedBy"`
	IsDefault bool `json:"isDefault"`
	OrganizationId int32 `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	ContinentName NullableString `json:"continentName"`
}

type _TanzuCredentialsListDto TanzuCredentialsListDto

// NewTanzuCredentialsListDto instantiates a new TanzuCredentialsListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTanzuCredentialsListDto(id int32, projectCount int32, isLocked bool, name string, username string, url string, volumeType string, namespace string, port NullableInt32, projects []CommonDropdownDto, createdBy string, createdAt string, lastModified NullableString, lastModifiedBy NullableString, isDefault bool, organizationId int32, organizationName string, continentName NullableString) *TanzuCredentialsListDto {
	this := TanzuCredentialsListDto{}
	this.Id = id
	this.ProjectCount = projectCount
	this.IsLocked = isLocked
	this.Name = name
	this.Username = username
	this.Url = url
	this.VolumeType = volumeType
	this.Namespace = namespace
	this.Port = port
	this.Projects = projects
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.LastModified = lastModified
	this.LastModifiedBy = lastModifiedBy
	this.IsDefault = isDefault
	this.OrganizationId = organizationId
	this.OrganizationName = organizationName
	this.ContinentName = continentName
	return &this
}

// NewTanzuCredentialsListDtoWithDefaults instantiates a new TanzuCredentialsListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTanzuCredentialsListDtoWithDefaults() *TanzuCredentialsListDto {
	this := TanzuCredentialsListDto{}
	return &this
}

// GetId returns the Id field value
func (o *TanzuCredentialsListDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TanzuCredentialsListDto) SetId(v int32) {
	o.Id = v
}

// GetProjectCount returns the ProjectCount field value
func (o *TanzuCredentialsListDto) GetProjectCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectCount
}

// GetProjectCountOk returns a tuple with the ProjectCount field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetProjectCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectCount, true
}

// SetProjectCount sets field value
func (o *TanzuCredentialsListDto) SetProjectCount(v int32) {
	o.ProjectCount = v
}

// GetIsLocked returns the IsLocked field value
func (o *TanzuCredentialsListDto) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *TanzuCredentialsListDto) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetName returns the Name field value
func (o *TanzuCredentialsListDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TanzuCredentialsListDto) SetName(v string) {
	o.Name = v
}

// GetUsername returns the Username field value
func (o *TanzuCredentialsListDto) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *TanzuCredentialsListDto) SetUsername(v string) {
	o.Username = v
}

// GetUrl returns the Url field value
func (o *TanzuCredentialsListDto) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TanzuCredentialsListDto) SetUrl(v string) {
	o.Url = v
}

// GetVolumeType returns the VolumeType field value
func (o *TanzuCredentialsListDto) GetVolumeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeType, true
}

// SetVolumeType sets field value
func (o *TanzuCredentialsListDto) SetVolumeType(v string) {
	o.VolumeType = v
}

// GetNamespace returns the Namespace field value
func (o *TanzuCredentialsListDto) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *TanzuCredentialsListDto) SetNamespace(v string) {
	o.Namespace = v
}

// GetPort returns the Port field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TanzuCredentialsListDto) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsListDto) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// SetPort sets field value
func (o *TanzuCredentialsListDto) SetPort(v int32) {
	o.Port.Set(&v)
}

// GetProjects returns the Projects field value
func (o *TanzuCredentialsListDto) GetProjects() []CommonDropdownDto {
	if o == nil {
		var ret []CommonDropdownDto
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetProjectsOk() ([]CommonDropdownDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Projects, true
}

// SetProjects sets field value
func (o *TanzuCredentialsListDto) SetProjects(v []CommonDropdownDto) {
	o.Projects = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *TanzuCredentialsListDto) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *TanzuCredentialsListDto) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TanzuCredentialsListDto) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TanzuCredentialsListDto) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetLastModified returns the LastModified field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TanzuCredentialsListDto) GetLastModified() string {
	if o == nil || o.LastModified.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsListDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// SetLastModified sets field value
func (o *TanzuCredentialsListDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}

// GetLastModifiedBy returns the LastModifiedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TanzuCredentialsListDto) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsListDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// SetLastModifiedBy sets field value
func (o *TanzuCredentialsListDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}

// GetIsDefault returns the IsDefault field value
func (o *TanzuCredentialsListDto) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *TanzuCredentialsListDto) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *TanzuCredentialsListDto) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *TanzuCredentialsListDto) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *TanzuCredentialsListDto) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *TanzuCredentialsListDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *TanzuCredentialsListDto) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetContinentName returns the ContinentName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TanzuCredentialsListDto) GetContinentName() string {
	if o == nil || o.ContinentName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContinentName.Get()
}

// GetContinentNameOk returns a tuple with the ContinentName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TanzuCredentialsListDto) GetContinentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContinentName.Get(), o.ContinentName.IsSet()
}

// SetContinentName sets field value
func (o *TanzuCredentialsListDto) SetContinentName(v string) {
	o.ContinentName.Set(&v)
}

func (o TanzuCredentialsListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TanzuCredentialsListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["projectCount"] = o.ProjectCount
	toSerialize["isLocked"] = o.IsLocked
	toSerialize["name"] = o.Name
	toSerialize["username"] = o.Username
	toSerialize["url"] = o.Url
	toSerialize["volumeType"] = o.VolumeType
	toSerialize["namespace"] = o.Namespace
	toSerialize["port"] = o.Port.Get()
	toSerialize["projects"] = o.Projects
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["lastModified"] = o.LastModified.Get()
	toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	toSerialize["isDefault"] = o.IsDefault
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["organizationName"] = o.OrganizationName
	toSerialize["continentName"] = o.ContinentName.Get()
	return toSerialize, nil
}

func (o *TanzuCredentialsListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"projectCount",
		"isLocked",
		"name",
		"username",
		"url",
		"volumeType",
		"namespace",
		"port",
		"projects",
		"createdBy",
		"createdAt",
		"lastModified",
		"lastModifiedBy",
		"isDefault",
		"organizationId",
		"organizationName",
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

	varTanzuCredentialsListDto := _TanzuCredentialsListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTanzuCredentialsListDto)

	if err != nil {
		return err
	}

	*o = TanzuCredentialsListDto(varTanzuCredentialsListDto)

	return err
}

type NullableTanzuCredentialsListDto struct {
	value *TanzuCredentialsListDto
	isSet bool
}

func (v NullableTanzuCredentialsListDto) Get() *TanzuCredentialsListDto {
	return v.value
}

func (v *NullableTanzuCredentialsListDto) Set(val *TanzuCredentialsListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTanzuCredentialsListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTanzuCredentialsListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTanzuCredentialsListDto(val *TanzuCredentialsListDto) *NullableTanzuCredentialsListDto {
	return &NullableTanzuCredentialsListDto{value: val, isSet: true}
}

func (v NullableTanzuCredentialsListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTanzuCredentialsListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


