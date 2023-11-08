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

// checks if the DashboardChart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardChart{}

// DashboardChart struct for DashboardChart
type DashboardChart struct {
	Organization *OrganizationEntityForDashboard `json:"organization,omitempty"`
	Projects *ProjectChartDto `json:"projects,omitempty"`
	CloudCredentials *CredentialChartDto `json:"cloudCredentials,omitempty"`
	Servers *ServerChartDto `json:"servers,omitempty"`
	StandAloneVms *ServerChartDto `json:"standAloneVms,omitempty"`
}

// NewDashboardChart instantiates a new DashboardChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardChart() *DashboardChart {
	this := DashboardChart{}
	return &this
}

// NewDashboardChartWithDefaults instantiates a new DashboardChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardChartWithDefaults() *DashboardChart {
	this := DashboardChart{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DashboardChart) GetOrganization() OrganizationEntityForDashboard {
	if o == nil || IsNil(o.Organization) {
		var ret OrganizationEntityForDashboard
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetOrganizationOk() (*OrganizationEntityForDashboard, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DashboardChart) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationEntityForDashboard and assigns it to the Organization field.
func (o *DashboardChart) SetOrganization(v OrganizationEntityForDashboard) {
	o.Organization = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *DashboardChart) GetProjects() ProjectChartDto {
	if o == nil || IsNil(o.Projects) {
		var ret ProjectChartDto
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetProjectsOk() (*ProjectChartDto, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *DashboardChart) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given ProjectChartDto and assigns it to the Projects field.
func (o *DashboardChart) SetProjects(v ProjectChartDto) {
	o.Projects = &v
}

// GetCloudCredentials returns the CloudCredentials field value if set, zero value otherwise.
func (o *DashboardChart) GetCloudCredentials() CredentialChartDto {
	if o == nil || IsNil(o.CloudCredentials) {
		var ret CredentialChartDto
		return ret
	}
	return *o.CloudCredentials
}

// GetCloudCredentialsOk returns a tuple with the CloudCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetCloudCredentialsOk() (*CredentialChartDto, bool) {
	if o == nil || IsNil(o.CloudCredentials) {
		return nil, false
	}
	return o.CloudCredentials, true
}

// HasCloudCredentials returns a boolean if a field has been set.
func (o *DashboardChart) HasCloudCredentials() bool {
	if o != nil && !IsNil(o.CloudCredentials) {
		return true
	}

	return false
}

// SetCloudCredentials gets a reference to the given CredentialChartDto and assigns it to the CloudCredentials field.
func (o *DashboardChart) SetCloudCredentials(v CredentialChartDto) {
	o.CloudCredentials = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *DashboardChart) GetServers() ServerChartDto {
	if o == nil || IsNil(o.Servers) {
		var ret ServerChartDto
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetServersOk() (*ServerChartDto, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *DashboardChart) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given ServerChartDto and assigns it to the Servers field.
func (o *DashboardChart) SetServers(v ServerChartDto) {
	o.Servers = &v
}

// GetStandAloneVms returns the StandAloneVms field value if set, zero value otherwise.
func (o *DashboardChart) GetStandAloneVms() ServerChartDto {
	if o == nil || IsNil(o.StandAloneVms) {
		var ret ServerChartDto
		return ret
	}
	return *o.StandAloneVms
}

// GetStandAloneVmsOk returns a tuple with the StandAloneVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardChart) GetStandAloneVmsOk() (*ServerChartDto, bool) {
	if o == nil || IsNil(o.StandAloneVms) {
		return nil, false
	}
	return o.StandAloneVms, true
}

// HasStandAloneVms returns a boolean if a field has been set.
func (o *DashboardChart) HasStandAloneVms() bool {
	if o != nil && !IsNil(o.StandAloneVms) {
		return true
	}

	return false
}

// SetStandAloneVms gets a reference to the given ServerChartDto and assigns it to the StandAloneVms field.
func (o *DashboardChart) SetStandAloneVms(v ServerChartDto) {
	o.StandAloneVms = &v
}

func (o DashboardChart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardChart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.CloudCredentials) {
		toSerialize["cloudCredentials"] = o.CloudCredentials
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.StandAloneVms) {
		toSerialize["standAloneVms"] = o.StandAloneVms
	}
	return toSerialize, nil
}

type NullableDashboardChart struct {
	value *DashboardChart
	isSet bool
}

func (v NullableDashboardChart) Get() *DashboardChart {
	return v.value
}

func (v *NullableDashboardChart) Set(val *DashboardChart) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardChart) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardChart(val *DashboardChart) *NullableDashboardChart {
	return &NullableDashboardChart{value: val, isSet: true}
}

func (v NullableDashboardChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


