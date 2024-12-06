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

// checks if the PodListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PodListDto{}

// PodListDto struct for PodListDto
type PodListDto struct {
	MetadataName string `json:"metadataName"`
	RestartCount int32 `json:"restartCount"`
	Namespace string `json:"namespace"`
	Node NullableString `json:"node"`
	Age NullableString `json:"age"`
	Status NullableString `json:"status,omitempty"`
	Container []string `json:"container"`
	State *KubernetesStateDto `json:"state,omitempty"`
	Type interface{} `json:"type,omitempty"`
}

type _PodListDto PodListDto

// NewPodListDto instantiates a new PodListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodListDto(metadataName string, restartCount int32, namespace string, node NullableString, age NullableString, container []string) *PodListDto {
	this := PodListDto{}
	this.MetadataName = metadataName
	this.RestartCount = restartCount
	this.Namespace = namespace
	this.Node = node
	this.Age = age
	this.Container = container
	return &this
}

// NewPodListDtoWithDefaults instantiates a new PodListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodListDtoWithDefaults() *PodListDto {
	this := PodListDto{}
	return &this
}

// GetMetadataName returns the MetadataName field value
func (o *PodListDto) GetMetadataName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataName
}

// GetMetadataNameOk returns a tuple with the MetadataName field value
// and a boolean to check if the value has been set.
func (o *PodListDto) GetMetadataNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataName, true
}

// SetMetadataName sets field value
func (o *PodListDto) SetMetadataName(v string) {
	o.MetadataName = v
}

// GetRestartCount returns the RestartCount field value
func (o *PodListDto) GetRestartCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value
// and a boolean to check if the value has been set.
func (o *PodListDto) GetRestartCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartCount, true
}

// SetRestartCount sets field value
func (o *PodListDto) SetRestartCount(v int32) {
	o.RestartCount = v
}

// GetNamespace returns the Namespace field value
func (o *PodListDto) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *PodListDto) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *PodListDto) SetNamespace(v string) {
	o.Namespace = v
}

// GetNode returns the Node field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PodListDto) GetNode() string {
	if o == nil || o.Node.Get() == nil {
		var ret string
		return ret
	}

	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodListDto) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// SetNode sets field value
func (o *PodListDto) SetNode(v string) {
	o.Node.Set(&v)
}

// GetAge returns the Age field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PodListDto) GetAge() string {
	if o == nil || o.Age.Get() == nil {
		var ret string
		return ret
	}

	return *o.Age.Get()
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodListDto) GetAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Age.Get(), o.Age.IsSet()
}

// SetAge sets field value
func (o *PodListDto) SetAge(v string) {
	o.Age.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodListDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodListDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PodListDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PodListDto) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PodListDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PodListDto) UnsetStatus() {
	o.Status.Unset()
}

// GetContainer returns the Container field value
func (o *PodListDto) GetContainer() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *PodListDto) GetContainerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container, true
}

// SetContainer sets field value
func (o *PodListDto) SetContainer(v []string) {
	o.Container = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PodListDto) GetState() KubernetesStateDto {
	if o == nil || IsNil(o.State) {
		var ret KubernetesStateDto
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodListDto) GetStateOk() (*KubernetesStateDto, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PodListDto) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given KubernetesStateDto and assigns it to the State field.
func (o *PodListDto) SetState(v KubernetesStateDto) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PodListDto) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PodListDto) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PodListDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PodListDto) SetType(v interface{}) {
	o.Type = v
}

func (o PodListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadataName"] = o.MetadataName
	toSerialize["restartCount"] = o.RestartCount
	toSerialize["namespace"] = o.Namespace
	toSerialize["node"] = o.Node.Get()
	toSerialize["age"] = o.Age.Get()
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["container"] = o.Container
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *PodListDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadataName",
		"restartCount",
		"namespace",
		"node",
		"age",
		"container",
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

	varPodListDto := _PodListDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPodListDto)

	if err != nil {
		return err
	}

	*o = PodListDto(varPodListDto)

	return err
}

type NullablePodListDto struct {
	value *PodListDto
	isSet bool
}

func (v NullablePodListDto) Get() *PodListDto {
	return v.value
}

func (v *NullablePodListDto) Set(val *PodListDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePodListDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePodListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodListDto(val *PodListDto) *NullablePodListDto {
	return &NullablePodListDto{value: val, isSet: true}
}

func (v NullablePodListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


