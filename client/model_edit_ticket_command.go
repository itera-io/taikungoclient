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

// checks if the EditTicketCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditTicketCommand{}

// EditTicketCommand struct for EditTicketCommand
type EditTicketCommand struct {
	TicketId NullableString `json:"ticketId,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
}

// NewEditTicketCommand instantiates a new EditTicketCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditTicketCommand() *EditTicketCommand {
	this := EditTicketCommand{}
	return &this
}

// NewEditTicketCommandWithDefaults instantiates a new EditTicketCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditTicketCommandWithDefaults() *EditTicketCommand {
	this := EditTicketCommand{}
	return &this
}

// GetTicketId returns the TicketId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditTicketCommand) GetTicketId() string {
	if o == nil || IsNil(o.TicketId.Get()) {
		var ret string
		return ret
	}
	return *o.TicketId.Get()
}

// GetTicketIdOk returns a tuple with the TicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditTicketCommand) GetTicketIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketId.Get(), o.TicketId.IsSet()
}

// HasTicketId returns a boolean if a field has been set.
func (o *EditTicketCommand) HasTicketId() bool {
	if o != nil && o.TicketId.IsSet() {
		return true
	}

	return false
}

// SetTicketId gets a reference to the given NullableString and assigns it to the TicketId field.
func (o *EditTicketCommand) SetTicketId(v string) {
	o.TicketId.Set(&v)
}
// SetTicketIdNil sets the value for TicketId to be an explicit nil
func (o *EditTicketCommand) SetTicketIdNil() {
	o.TicketId.Set(nil)
}

// UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
func (o *EditTicketCommand) UnsetTicketId() {
	o.TicketId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditTicketCommand) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditTicketCommand) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EditTicketCommand) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EditTicketCommand) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EditTicketCommand) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EditTicketCommand) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditTicketCommand) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditTicketCommand) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EditTicketCommand) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *EditTicketCommand) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *EditTicketCommand) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *EditTicketCommand) UnsetDescription() {
	o.Description.Unset()
}

func (o EditTicketCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditTicketCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TicketId.IsSet() {
		toSerialize["ticketId"] = o.TicketId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullableEditTicketCommand struct {
	value *EditTicketCommand
	isSet bool
}

func (v NullableEditTicketCommand) Get() *EditTicketCommand {
	return v.value
}

func (v *NullableEditTicketCommand) Set(val *EditTicketCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableEditTicketCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableEditTicketCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditTicketCommand(val *EditTicketCommand) *NullableEditTicketCommand {
	return &NullableEditTicketCommand{value: val, isSet: true}
}

func (v NullableEditTicketCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditTicketCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


