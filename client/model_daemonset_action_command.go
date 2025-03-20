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

// checks if the DaemonsetActionCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaemonsetActionCommand{}

// DaemonsetActionCommand struct for DaemonsetActionCommand
type DaemonsetActionCommand struct {
	ProjectId int32 `json:"projectId"`
	Name NullableString `json:"name"`
	Namespace NullableString `json:"namespace"`
	Action EDaemonSetAction `json:"action"`
}

type _DaemonsetActionCommand DaemonsetActionCommand

// NewDaemonsetActionCommand instantiates a new DaemonsetActionCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaemonsetActionCommand(projectId int32, name NullableString, namespace NullableString, action EDaemonSetAction) *DaemonsetActionCommand {
	this := DaemonsetActionCommand{}
	this.ProjectId = projectId
	this.Name = name
	this.Namespace = namespace
	this.Action = action
	return &this
}

// NewDaemonsetActionCommandWithDefaults instantiates a new DaemonsetActionCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaemonsetActionCommandWithDefaults() *DaemonsetActionCommand {
	this := DaemonsetActionCommand{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DaemonsetActionCommand) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DaemonsetActionCommand) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DaemonsetActionCommand) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DaemonsetActionCommand) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DaemonsetActionCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *DaemonsetActionCommand) SetName(v string) {
	o.Name.Set(&v)
}

// GetNamespace returns the Namespace field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DaemonsetActionCommand) GetNamespace() string {
	if o == nil || o.Namespace.Get() == nil {
		var ret string
		return ret
	}

	return *o.Namespace.Get()
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DaemonsetActionCommand) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace.Get(), o.Namespace.IsSet()
}

// SetNamespace sets field value
func (o *DaemonsetActionCommand) SetNamespace(v string) {
	o.Namespace.Set(&v)
}

// GetAction returns the Action field value
func (o *DaemonsetActionCommand) GetAction() EDaemonSetAction {
	if o == nil {
		var ret EDaemonSetAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DaemonsetActionCommand) GetActionOk() (*EDaemonSetAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *DaemonsetActionCommand) SetAction(v EDaemonSetAction) {
	o.Action = v
}

func (o DaemonsetActionCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaemonsetActionCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name.Get()
	toSerialize["namespace"] = o.Namespace.Get()
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *DaemonsetActionCommand) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
		"name",
		"namespace",
		"action",
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

	varDaemonsetActionCommand := _DaemonsetActionCommand{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDaemonsetActionCommand)

	if err != nil {
		return err
	}

	*o = DaemonsetActionCommand(varDaemonsetActionCommand)

	return err
}

type NullableDaemonsetActionCommand struct {
	value *DaemonsetActionCommand
	isSet bool
}

func (v NullableDaemonsetActionCommand) Get() *DaemonsetActionCommand {
	return v.value
}

func (v *NullableDaemonsetActionCommand) Set(val *DaemonsetActionCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableDaemonsetActionCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableDaemonsetActionCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaemonsetActionCommand(val *DaemonsetActionCommand) *NullableDaemonsetActionCommand {
	return &NullableDaemonsetActionCommand{value: val, isSet: true}
}

func (v NullableDaemonsetActionCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaemonsetActionCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


