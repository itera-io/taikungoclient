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

// checks if the KubernetesCursorPodList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCursorPodList{}

// KubernetesCursorPodList struct for KubernetesCursorPodList
type KubernetesCursorPodList struct {
	Data []PodListDto `json:"data"`
	NextPageLink NullableString `json:"nextPageLink"`
}

type _KubernetesCursorPodList KubernetesCursorPodList

// NewKubernetesCursorPodList instantiates a new KubernetesCursorPodList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCursorPodList(data []PodListDto, nextPageLink NullableString) *KubernetesCursorPodList {
	this := KubernetesCursorPodList{}
	this.Data = data
	this.NextPageLink = nextPageLink
	return &this
}

// NewKubernetesCursorPodListWithDefaults instantiates a new KubernetesCursorPodList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesCursorPodListWithDefaults() *KubernetesCursorPodList {
	this := KubernetesCursorPodList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []PodListDto will be returned
func (o *KubernetesCursorPodList) GetData() []PodListDto {
	if o == nil {
		var ret []PodListDto
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCursorPodList) GetDataOk() ([]PodListDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *KubernetesCursorPodList) SetData(v []PodListDto) {
	o.Data = v
}

// GetNextPageLink returns the NextPageLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KubernetesCursorPodList) GetNextPageLink() string {
	if o == nil || o.NextPageLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.NextPageLink.Get()
}

// GetNextPageLinkOk returns a tuple with the NextPageLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesCursorPodList) GetNextPageLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageLink.Get(), o.NextPageLink.IsSet()
}

// SetNextPageLink sets field value
func (o *KubernetesCursorPodList) SetNextPageLink(v string) {
	o.NextPageLink.Set(&v)
}

func (o KubernetesCursorPodList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCursorPodList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["nextPageLink"] = o.NextPageLink.Get()
	return toSerialize, nil
}

func (o *KubernetesCursorPodList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"nextPageLink",
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

	varKubernetesCursorPodList := _KubernetesCursorPodList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesCursorPodList)

	if err != nil {
		return err
	}

	*o = KubernetesCursorPodList(varKubernetesCursorPodList)

	return err
}

type NullableKubernetesCursorPodList struct {
	value *KubernetesCursorPodList
	isSet bool
}

func (v NullableKubernetesCursorPodList) Get() *KubernetesCursorPodList {
	return v.value
}

func (v *NullableKubernetesCursorPodList) Set(val *KubernetesCursorPodList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCursorPodList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCursorPodList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCursorPodList(val *KubernetesCursorPodList) *NullableKubernetesCursorPodList {
	return &NullableKubernetesCursorPodList{value: val, isSet: true}
}

func (v NullableKubernetesCursorPodList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCursorPodList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


