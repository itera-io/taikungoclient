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

// checks if the CreateZededaCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateZededaCommand{}

// CreateZededaCommand struct for CreateZededaCommand
type CreateZededaCommand struct {
	Name NullableString `json:"name,omitempty"`
	ApiToken NullableString `json:"apiToken,omitempty"`
	ApiUrl NullableString `json:"apiUrl,omitempty"`
	Project NullableString `json:"project,omitempty"`
	Continent NullableString `json:"continent,omitempty"`
	OrganizationId NullableInt32 `json:"organizationId,omitempty"`
	EdgeNodes []string `json:"edgeNodes,omitempty"`
	PublicNetwork *CreateZededaNetworkDto `json:"publicNetwork,omitempty"`
	PrivateNetwork *CreateZededaNetworkDto `json:"privateNetwork,omitempty"`
}

// NewCreateZededaCommand instantiates a new CreateZededaCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateZededaCommand() *CreateZededaCommand {
	this := CreateZededaCommand{}
	return &this
}

// NewCreateZededaCommandWithDefaults instantiates a new CreateZededaCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateZededaCommandWithDefaults() *CreateZededaCommand {
	this := CreateZededaCommand{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateZededaCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateZededaCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateZededaCommand) UnsetName() {
	o.Name.Unset()
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetApiToken() string {
	if o == nil || IsNil(o.ApiToken.Get()) {
		var ret string
		return ret
	}
	return *o.ApiToken.Get()
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiToken.Get(), o.ApiToken.IsSet()
}

// HasApiToken returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasApiToken() bool {
	if o != nil && o.ApiToken.IsSet() {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given NullableString and assigns it to the ApiToken field.
func (o *CreateZededaCommand) SetApiToken(v string) {
	o.ApiToken.Set(&v)
}
// SetApiTokenNil sets the value for ApiToken to be an explicit nil
func (o *CreateZededaCommand) SetApiTokenNil() {
	o.ApiToken.Set(nil)
}

// UnsetApiToken ensures that no value is present for ApiToken, not even an explicit nil
func (o *CreateZededaCommand) UnsetApiToken() {
	o.ApiToken.Unset()
}

// GetApiUrl returns the ApiUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetApiUrl() string {
	if o == nil || IsNil(o.ApiUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ApiUrl.Get()
}

// GetApiUrlOk returns a tuple with the ApiUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetApiUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiUrl.Get(), o.ApiUrl.IsSet()
}

// HasApiUrl returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasApiUrl() bool {
	if o != nil && o.ApiUrl.IsSet() {
		return true
	}

	return false
}

// SetApiUrl gets a reference to the given NullableString and assigns it to the ApiUrl field.
func (o *CreateZededaCommand) SetApiUrl(v string) {
	o.ApiUrl.Set(&v)
}
// SetApiUrlNil sets the value for ApiUrl to be an explicit nil
func (o *CreateZededaCommand) SetApiUrlNil() {
	o.ApiUrl.Set(nil)
}

// UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
func (o *CreateZededaCommand) UnsetApiUrl() {
	o.ApiUrl.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetProject() string {
	if o == nil || IsNil(o.Project.Get()) {
		var ret string
		return ret
	}
	return *o.Project.Get()
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Project.Get(), o.Project.IsSet()
}

// HasProject returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasProject() bool {
	if o != nil && o.Project.IsSet() {
		return true
	}

	return false
}

// SetProject gets a reference to the given NullableString and assigns it to the Project field.
func (o *CreateZededaCommand) SetProject(v string) {
	o.Project.Set(&v)
}
// SetProjectNil sets the value for Project to be an explicit nil
func (o *CreateZededaCommand) SetProjectNil() {
	o.Project.Set(nil)
}

// UnsetProject ensures that no value is present for Project, not even an explicit nil
func (o *CreateZededaCommand) UnsetProject() {
	o.Project.Unset()
}

// GetContinent returns the Continent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetContinent() string {
	if o == nil || IsNil(o.Continent.Get()) {
		var ret string
		return ret
	}
	return *o.Continent.Get()
}

// GetContinentOk returns a tuple with the Continent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetContinentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Continent.Get(), o.Continent.IsSet()
}

// HasContinent returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasContinent() bool {
	if o != nil && o.Continent.IsSet() {
		return true
	}

	return false
}

// SetContinent gets a reference to the given NullableString and assigns it to the Continent field.
func (o *CreateZededaCommand) SetContinent(v string) {
	o.Continent.Set(&v)
}
// SetContinentNil sets the value for Continent to be an explicit nil
func (o *CreateZededaCommand) SetContinentNil() {
	o.Continent.Set(nil)
}

// UnsetContinent ensures that no value is present for Continent, not even an explicit nil
func (o *CreateZededaCommand) UnsetContinent() {
	o.Continent.Unset()
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId.Get()) {
		var ret int32
		return ret
	}
	return *o.OrganizationId.Get()
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrganizationId.Get(), o.OrganizationId.IsSet()
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasOrganizationId() bool {
	if o != nil && o.OrganizationId.IsSet() {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given NullableInt32 and assigns it to the OrganizationId field.
func (o *CreateZededaCommand) SetOrganizationId(v int32) {
	o.OrganizationId.Set(&v)
}
// SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil
func (o *CreateZededaCommand) SetOrganizationIdNil() {
	o.OrganizationId.Set(nil)
}

// UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
func (o *CreateZededaCommand) UnsetOrganizationId() {
	o.OrganizationId.Unset()
}

// GetEdgeNodes returns the EdgeNodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateZededaCommand) GetEdgeNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EdgeNodes
}

// GetEdgeNodesOk returns a tuple with the EdgeNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateZededaCommand) GetEdgeNodesOk() ([]string, bool) {
	if o == nil || IsNil(o.EdgeNodes) {
		return nil, false
	}
	return o.EdgeNodes, true
}

// HasEdgeNodes returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasEdgeNodes() bool {
	if o != nil && !IsNil(o.EdgeNodes) {
		return true
	}

	return false
}

// SetEdgeNodes gets a reference to the given []string and assigns it to the EdgeNodes field.
func (o *CreateZededaCommand) SetEdgeNodes(v []string) {
	o.EdgeNodes = v
}

// GetPublicNetwork returns the PublicNetwork field value if set, zero value otherwise.
func (o *CreateZededaCommand) GetPublicNetwork() CreateZededaNetworkDto {
	if o == nil || IsNil(o.PublicNetwork) {
		var ret CreateZededaNetworkDto
		return ret
	}
	return *o.PublicNetwork
}

// GetPublicNetworkOk returns a tuple with the PublicNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZededaCommand) GetPublicNetworkOk() (*CreateZededaNetworkDto, bool) {
	if o == nil || IsNil(o.PublicNetwork) {
		return nil, false
	}
	return o.PublicNetwork, true
}

// HasPublicNetwork returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasPublicNetwork() bool {
	if o != nil && !IsNil(o.PublicNetwork) {
		return true
	}

	return false
}

// SetPublicNetwork gets a reference to the given CreateZededaNetworkDto and assigns it to the PublicNetwork field.
func (o *CreateZededaCommand) SetPublicNetwork(v CreateZededaNetworkDto) {
	o.PublicNetwork = &v
}

// GetPrivateNetwork returns the PrivateNetwork field value if set, zero value otherwise.
func (o *CreateZededaCommand) GetPrivateNetwork() CreateZededaNetworkDto {
	if o == nil || IsNil(o.PrivateNetwork) {
		var ret CreateZededaNetworkDto
		return ret
	}
	return *o.PrivateNetwork
}

// GetPrivateNetworkOk returns a tuple with the PrivateNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateZededaCommand) GetPrivateNetworkOk() (*CreateZededaNetworkDto, bool) {
	if o == nil || IsNil(o.PrivateNetwork) {
		return nil, false
	}
	return o.PrivateNetwork, true
}

// HasPrivateNetwork returns a boolean if a field has been set.
func (o *CreateZededaCommand) HasPrivateNetwork() bool {
	if o != nil && !IsNil(o.PrivateNetwork) {
		return true
	}

	return false
}

// SetPrivateNetwork gets a reference to the given CreateZededaNetworkDto and assigns it to the PrivateNetwork field.
func (o *CreateZededaCommand) SetPrivateNetwork(v CreateZededaNetworkDto) {
	o.PrivateNetwork = &v
}

func (o CreateZededaCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateZededaCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ApiToken.IsSet() {
		toSerialize["apiToken"] = o.ApiToken.Get()
	}
	if o.ApiUrl.IsSet() {
		toSerialize["apiUrl"] = o.ApiUrl.Get()
	}
	if o.Project.IsSet() {
		toSerialize["project"] = o.Project.Get()
	}
	if o.Continent.IsSet() {
		toSerialize["continent"] = o.Continent.Get()
	}
	if o.OrganizationId.IsSet() {
		toSerialize["organizationId"] = o.OrganizationId.Get()
	}
	if o.EdgeNodes != nil {
		toSerialize["edgeNodes"] = o.EdgeNodes
	}
	if !IsNil(o.PublicNetwork) {
		toSerialize["publicNetwork"] = o.PublicNetwork
	}
	if !IsNil(o.PrivateNetwork) {
		toSerialize["privateNetwork"] = o.PrivateNetwork
	}
	return toSerialize, nil
}

type NullableCreateZededaCommand struct {
	value *CreateZededaCommand
	isSet bool
}

func (v NullableCreateZededaCommand) Get() *CreateZededaCommand {
	return v.value
}

func (v *NullableCreateZededaCommand) Set(val *CreateZededaCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateZededaCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateZededaCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateZededaCommand(val *CreateZededaCommand) *NullableCreateZededaCommand {
	return &NullableCreateZededaCommand{value: val, isSet: true}
}

func (v NullableCreateZededaCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateZededaCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


