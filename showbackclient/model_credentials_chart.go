/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikunshowback

import (
	"encoding/json"
)

// checks if the CredentialsChart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsChart{}

// CredentialsChart struct for CredentialsChart
type CredentialsChart struct {
	Amazon              []AmazonCredentialsListDto    `json:"amazon,omitempty"`
	Openstack           []OpenstackCredentialsListDto `json:"openstack,omitempty"`
	Azure               []AzureCredentialsListDto     `json:"azure,omitempty"`
	Google              []GoogleCredentialsListDto    `json:"google,omitempty"`
	Tanzu               []TanzuCredentialsListDto     `json:"tanzu,omitempty"`
	TotalCountOpenstack *int32                        `json:"totalCountOpenstack,omitempty"`
	TotalCountAws       *int32                        `json:"totalCountAws,omitempty"`
	TotalCountAzure     *int32                        `json:"totalCountAzure,omitempty"`
	TotalCountGoogle    *int32                        `json:"totalCountGoogle,omitempty"`
	TotalCountTanzu     *int32                        `json:"totalCountTanzu,omitempty"`
}

// NewCredentialsChart instantiates a new CredentialsChart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsChart() *CredentialsChart {
	this := CredentialsChart{}
	return &this
}

// NewCredentialsChartWithDefaults instantiates a new CredentialsChart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsChartWithDefaults() *CredentialsChart {
	this := CredentialsChart{}
	return &this
}

// GetAmazon returns the Amazon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChart) GetAmazon() []AmazonCredentialsListDto {
	if o == nil {
		var ret []AmazonCredentialsListDto
		return ret
	}
	return o.Amazon
}

// GetAmazonOk returns a tuple with the Amazon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChart) GetAmazonOk() ([]AmazonCredentialsListDto, bool) {
	if o == nil || IsNil(o.Amazon) {
		return nil, false
	}
	return o.Amazon, true
}

// HasAmazon returns a boolean if a field has been set.
func (o *CredentialsChart) HasAmazon() bool {
	if o != nil && IsNil(o.Amazon) {
		return true
	}

	return false
}

// SetAmazon gets a reference to the given []AmazonCredentialsListDto and assigns it to the Amazon field.
func (o *CredentialsChart) SetAmazon(v []AmazonCredentialsListDto) {
	o.Amazon = v
}

// GetOpenstack returns the Openstack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChart) GetOpenstack() []OpenstackCredentialsListDto {
	if o == nil {
		var ret []OpenstackCredentialsListDto
		return ret
	}
	return o.Openstack
}

// GetOpenstackOk returns a tuple with the Openstack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChart) GetOpenstackOk() ([]OpenstackCredentialsListDto, bool) {
	if o == nil || IsNil(o.Openstack) {
		return nil, false
	}
	return o.Openstack, true
}

// HasOpenstack returns a boolean if a field has been set.
func (o *CredentialsChart) HasOpenstack() bool {
	if o != nil && IsNil(o.Openstack) {
		return true
	}

	return false
}

// SetOpenstack gets a reference to the given []OpenstackCredentialsListDto and assigns it to the Openstack field.
func (o *CredentialsChart) SetOpenstack(v []OpenstackCredentialsListDto) {
	o.Openstack = v
}

// GetAzure returns the Azure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChart) GetAzure() []AzureCredentialsListDto {
	if o == nil {
		var ret []AzureCredentialsListDto
		return ret
	}
	return o.Azure
}

// GetAzureOk returns a tuple with the Azure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChart) GetAzureOk() ([]AzureCredentialsListDto, bool) {
	if o == nil || IsNil(o.Azure) {
		return nil, false
	}
	return o.Azure, true
}

// HasAzure returns a boolean if a field has been set.
func (o *CredentialsChart) HasAzure() bool {
	if o != nil && IsNil(o.Azure) {
		return true
	}

	return false
}

// SetAzure gets a reference to the given []AzureCredentialsListDto and assigns it to the Azure field.
func (o *CredentialsChart) SetAzure(v []AzureCredentialsListDto) {
	o.Azure = v
}

// GetGoogle returns the Google field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChart) GetGoogle() []GoogleCredentialsListDto {
	if o == nil {
		var ret []GoogleCredentialsListDto
		return ret
	}
	return o.Google
}

// GetGoogleOk returns a tuple with the Google field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChart) GetGoogleOk() ([]GoogleCredentialsListDto, bool) {
	if o == nil || IsNil(o.Google) {
		return nil, false
	}
	return o.Google, true
}

// HasGoogle returns a boolean if a field has been set.
func (o *CredentialsChart) HasGoogle() bool {
	if o != nil && IsNil(o.Google) {
		return true
	}

	return false
}

// SetGoogle gets a reference to the given []GoogleCredentialsListDto and assigns it to the Google field.
func (o *CredentialsChart) SetGoogle(v []GoogleCredentialsListDto) {
	o.Google = v
}

// GetTanzu returns the Tanzu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CredentialsChart) GetTanzu() []TanzuCredentialsListDto {
	if o == nil {
		var ret []TanzuCredentialsListDto
		return ret
	}
	return o.Tanzu
}

// GetTanzuOk returns a tuple with the Tanzu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CredentialsChart) GetTanzuOk() ([]TanzuCredentialsListDto, bool) {
	if o == nil || IsNil(o.Tanzu) {
		return nil, false
	}
	return o.Tanzu, true
}

// HasTanzu returns a boolean if a field has been set.
func (o *CredentialsChart) HasTanzu() bool {
	if o != nil && IsNil(o.Tanzu) {
		return true
	}

	return false
}

// SetTanzu gets a reference to the given []TanzuCredentialsListDto and assigns it to the Tanzu field.
func (o *CredentialsChart) SetTanzu(v []TanzuCredentialsListDto) {
	o.Tanzu = v
}

// GetTotalCountOpenstack returns the TotalCountOpenstack field value if set, zero value otherwise.
func (o *CredentialsChart) GetTotalCountOpenstack() int32 {
	if o == nil || IsNil(o.TotalCountOpenstack) {
		var ret int32
		return ret
	}
	return *o.TotalCountOpenstack
}

// GetTotalCountOpenstackOk returns a tuple with the TotalCountOpenstack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsChart) GetTotalCountOpenstackOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCountOpenstack) {
		return nil, false
	}
	return o.TotalCountOpenstack, true
}

// HasTotalCountOpenstack returns a boolean if a field has been set.
func (o *CredentialsChart) HasTotalCountOpenstack() bool {
	if o != nil && !IsNil(o.TotalCountOpenstack) {
		return true
	}

	return false
}

// SetTotalCountOpenstack gets a reference to the given int32 and assigns it to the TotalCountOpenstack field.
func (o *CredentialsChart) SetTotalCountOpenstack(v int32) {
	o.TotalCountOpenstack = &v
}

// GetTotalCountAws returns the TotalCountAws field value if set, zero value otherwise.
func (o *CredentialsChart) GetTotalCountAws() int32 {
	if o == nil || IsNil(o.TotalCountAws) {
		var ret int32
		return ret
	}
	return *o.TotalCountAws
}

// GetTotalCountAwsOk returns a tuple with the TotalCountAws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsChart) GetTotalCountAwsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCountAws) {
		return nil, false
	}
	return o.TotalCountAws, true
}

// HasTotalCountAws returns a boolean if a field has been set.
func (o *CredentialsChart) HasTotalCountAws() bool {
	if o != nil && !IsNil(o.TotalCountAws) {
		return true
	}

	return false
}

// SetTotalCountAws gets a reference to the given int32 and assigns it to the TotalCountAws field.
func (o *CredentialsChart) SetTotalCountAws(v int32) {
	o.TotalCountAws = &v
}

// GetTotalCountAzure returns the TotalCountAzure field value if set, zero value otherwise.
func (o *CredentialsChart) GetTotalCountAzure() int32 {
	if o == nil || IsNil(o.TotalCountAzure) {
		var ret int32
		return ret
	}
	return *o.TotalCountAzure
}

// GetTotalCountAzureOk returns a tuple with the TotalCountAzure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsChart) GetTotalCountAzureOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCountAzure) {
		return nil, false
	}
	return o.TotalCountAzure, true
}

// HasTotalCountAzure returns a boolean if a field has been set.
func (o *CredentialsChart) HasTotalCountAzure() bool {
	if o != nil && !IsNil(o.TotalCountAzure) {
		return true
	}

	return false
}

// SetTotalCountAzure gets a reference to the given int32 and assigns it to the TotalCountAzure field.
func (o *CredentialsChart) SetTotalCountAzure(v int32) {
	o.TotalCountAzure = &v
}

// GetTotalCountGoogle returns the TotalCountGoogle field value if set, zero value otherwise.
func (o *CredentialsChart) GetTotalCountGoogle() int32 {
	if o == nil || IsNil(o.TotalCountGoogle) {
		var ret int32
		return ret
	}
	return *o.TotalCountGoogle
}

// GetTotalCountGoogleOk returns a tuple with the TotalCountGoogle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsChart) GetTotalCountGoogleOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCountGoogle) {
		return nil, false
	}
	return o.TotalCountGoogle, true
}

// HasTotalCountGoogle returns a boolean if a field has been set.
func (o *CredentialsChart) HasTotalCountGoogle() bool {
	if o != nil && !IsNil(o.TotalCountGoogle) {
		return true
	}

	return false
}

// SetTotalCountGoogle gets a reference to the given int32 and assigns it to the TotalCountGoogle field.
func (o *CredentialsChart) SetTotalCountGoogle(v int32) {
	o.TotalCountGoogle = &v
}

// GetTotalCountTanzu returns the TotalCountTanzu field value if set, zero value otherwise.
func (o *CredentialsChart) GetTotalCountTanzu() int32 {
	if o == nil || IsNil(o.TotalCountTanzu) {
		var ret int32
		return ret
	}
	return *o.TotalCountTanzu
}

// GetTotalCountTanzuOk returns a tuple with the TotalCountTanzu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsChart) GetTotalCountTanzuOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCountTanzu) {
		return nil, false
	}
	return o.TotalCountTanzu, true
}

// HasTotalCountTanzu returns a boolean if a field has been set.
func (o *CredentialsChart) HasTotalCountTanzu() bool {
	if o != nil && !IsNil(o.TotalCountTanzu) {
		return true
	}

	return false
}

// SetTotalCountTanzu gets a reference to the given int32 and assigns it to the TotalCountTanzu field.
func (o *CredentialsChart) SetTotalCountTanzu(v int32) {
	o.TotalCountTanzu = &v
}

func (o CredentialsChart) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsChart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amazon != nil {
		toSerialize["amazon"] = o.Amazon
	}
	if o.Openstack != nil {
		toSerialize["openstack"] = o.Openstack
	}
	if o.Azure != nil {
		toSerialize["azure"] = o.Azure
	}
	if o.Google != nil {
		toSerialize["google"] = o.Google
	}
	if o.Tanzu != nil {
		toSerialize["tanzu"] = o.Tanzu
	}
	if !IsNil(o.TotalCountOpenstack) {
		toSerialize["totalCountOpenstack"] = o.TotalCountOpenstack
	}
	if !IsNil(o.TotalCountAws) {
		toSerialize["totalCountAws"] = o.TotalCountAws
	}
	if !IsNil(o.TotalCountAzure) {
		toSerialize["totalCountAzure"] = o.TotalCountAzure
	}
	if !IsNil(o.TotalCountGoogle) {
		toSerialize["totalCountGoogle"] = o.TotalCountGoogle
	}
	if !IsNil(o.TotalCountTanzu) {
		toSerialize["totalCountTanzu"] = o.TotalCountTanzu
	}
	return toSerialize, nil
}

type NullableCredentialsChart struct {
	value *CredentialsChart
	isSet bool
}

func (v NullableCredentialsChart) Get() *CredentialsChart {
	return v.value
}

func (v *NullableCredentialsChart) Set(val *CredentialsChart) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsChart) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsChart(val *CredentialsChart) *NullableCredentialsChart {
	return &NullableCredentialsChart{value: val, isSet: true}
}

func (v NullableCredentialsChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
