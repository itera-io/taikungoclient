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
	"bytes"
	"fmt"
)

// checks if the ProjectsForBillingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectsForBillingDto{}

// ProjectsForBillingDto struct for ProjectsForBillingDto
type ProjectsForBillingDto struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	CreatedAt NullableTime `json:"createdAt"`
	BillingStartDate NullableTime `json:"billingStartDate"`
	OrganizationName NullableString `json:"organizationName"`
	Price float64 `json:"price"`
	Servers []ServersForBillingDto `json:"servers"`
	StandaloneVms []StandaloneVmsForBillingDto `json:"standaloneVms"`
	BillingEnabled bool `json:"billingEnabled"`
}

type _ProjectsForBillingDto ProjectsForBillingDto

// NewProjectsForBillingDto instantiates a new ProjectsForBillingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsForBillingDto(id int32, name NullableString, createdAt NullableTime, billingStartDate NullableTime, organizationName NullableString, price float64, servers []ServersForBillingDto, standaloneVms []StandaloneVmsForBillingDto, billingEnabled bool) *ProjectsForBillingDto {
	this := ProjectsForBillingDto{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.BillingStartDate = billingStartDate
	this.OrganizationName = organizationName
	this.Price = price
	this.Servers = servers
	this.StandaloneVms = standaloneVms
	this.BillingEnabled = billingEnabled
	return &this
}

// NewProjectsForBillingDtoWithDefaults instantiates a new ProjectsForBillingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsForBillingDtoWithDefaults() *ProjectsForBillingDto {
	this := ProjectsForBillingDto{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectsForBillingDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectsForBillingDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectsForBillingDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectsForBillingDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ProjectsForBillingDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ProjectsForBillingDto) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *ProjectsForBillingDto) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetBillingStartDate returns the BillingStartDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ProjectsForBillingDto) GetBillingStartDate() time.Time {
	if o == nil || o.BillingStartDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.BillingStartDate.Get()
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetBillingStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingStartDate.Get(), o.BillingStartDate.IsSet()
}

// SetBillingStartDate sets field value
func (o *ProjectsForBillingDto) SetBillingStartDate(v time.Time) {
	o.BillingStartDate.Set(&v)
}

// GetOrganizationName returns the OrganizationName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectsForBillingDto) GetOrganizationName() string {
	if o == nil || o.OrganizationName.Get() == nil {
		var ret string
		return ret
	}

	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// SetOrganizationName sets field value
func (o *ProjectsForBillingDto) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}

// GetPrice returns the Price field value
func (o *ProjectsForBillingDto) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ProjectsForBillingDto) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ProjectsForBillingDto) SetPrice(v float64) {
	o.Price = v
}

// GetServers returns the Servers field value
// If the value is explicit nil, the zero value for []ServersForBillingDto will be returned
func (o *ProjectsForBillingDto) GetServers() []ServersForBillingDto {
	if o == nil {
		var ret []ServersForBillingDto
		return ret
	}

	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetServersOk() ([]ServersForBillingDto, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// SetServers sets field value
func (o *ProjectsForBillingDto) SetServers(v []ServersForBillingDto) {
	o.Servers = v
}

// GetStandaloneVms returns the StandaloneVms field value
// If the value is explicit nil, the zero value for []StandaloneVmsForBillingDto will be returned
func (o *ProjectsForBillingDto) GetStandaloneVms() []StandaloneVmsForBillingDto {
	if o == nil {
		var ret []StandaloneVmsForBillingDto
		return ret
	}

	return o.StandaloneVms
}

// GetStandaloneVmsOk returns a tuple with the StandaloneVms field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsForBillingDto) GetStandaloneVmsOk() ([]StandaloneVmsForBillingDto, bool) {
	if o == nil || IsNil(o.StandaloneVms) {
		return nil, false
	}
	return o.StandaloneVms, true
}

// SetStandaloneVms sets field value
func (o *ProjectsForBillingDto) SetStandaloneVms(v []StandaloneVmsForBillingDto) {
	o.StandaloneVms = v
}

// GetBillingEnabled returns the BillingEnabled field value
func (o *ProjectsForBillingDto) GetBillingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BillingEnabled
}

// GetBillingEnabledOk returns a tuple with the BillingEnabled field value
// and a boolean to check if the value has been set.
func (o *ProjectsForBillingDto) GetBillingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingEnabled, true
}

// SetBillingEnabled sets field value
func (o *ProjectsForBillingDto) SetBillingEnabled(v bool) {
	o.BillingEnabled = v
}

func (o ProjectsForBillingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectsForBillingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["billingStartDate"] = o.BillingStartDate.Get()
	toSerialize["organizationName"] = o.OrganizationName.Get()
	toSerialize["price"] = o.Price
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.StandaloneVms != nil {
		toSerialize["standaloneVms"] = o.StandaloneVms
	}
	toSerialize["billingEnabled"] = o.BillingEnabled
	return toSerialize, nil
}

func (o *ProjectsForBillingDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"createdAt",
		"billingStartDate",
		"organizationName",
		"price",
		"servers",
		"standaloneVms",
		"billingEnabled",
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

	varProjectsForBillingDto := _ProjectsForBillingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectsForBillingDto)

	if err != nil {
		return err
	}

	*o = ProjectsForBillingDto(varProjectsForBillingDto)

	return err
}

type NullableProjectsForBillingDto struct {
	value *ProjectsForBillingDto
	isSet bool
}

func (v NullableProjectsForBillingDto) Get() *ProjectsForBillingDto {
	return v.value
}

func (v *NullableProjectsForBillingDto) Set(val *ProjectsForBillingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsForBillingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsForBillingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsForBillingDto(val *ProjectsForBillingDto) *NullableProjectsForBillingDto {
	return &NullableProjectsForBillingDto{value: val, isSet: true}
}

func (v NullableProjectsForBillingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsForBillingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


