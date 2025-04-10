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

// checks if the InvoiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDto{}

// InvoiceDto struct for InvoiceDto
type InvoiceDto struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name"`
	DocumentNumber NullableString `json:"documentNumber"`
	OrganizationSubscriptionId int32 `json:"organizationSubscriptionId"`
	IsPaid bool `json:"isPaid"`
	RequiredPaymentAction bool `json:"requiredPaymentAction"`
	StripeInvoiceId NullableString `json:"stripeInvoiceId"`
	Price float64 `json:"price"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
	DueDate time.Time `json:"dueDate"`
}

type _InvoiceDto InvoiceDto

// NewInvoiceDto instantiates a new InvoiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDto(id int32, name NullableString, documentNumber NullableString, organizationSubscriptionId int32, isPaid bool, requiredPaymentAction bool, stripeInvoiceId NullableString, price float64, startDate time.Time, endDate time.Time, dueDate time.Time) *InvoiceDto {
	this := InvoiceDto{}
	this.Id = id
	this.Name = name
	this.DocumentNumber = documentNumber
	this.OrganizationSubscriptionId = organizationSubscriptionId
	this.IsPaid = isPaid
	this.RequiredPaymentAction = requiredPaymentAction
	this.StripeInvoiceId = stripeInvoiceId
	this.Price = price
	this.StartDate = startDate
	this.EndDate = endDate
	this.DueDate = dueDate
	return &this
}

// NewInvoiceDtoWithDefaults instantiates a new InvoiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDtoWithDefaults() *InvoiceDto {
	this := InvoiceDto{}
	return &this
}

// GetId returns the Id field value
func (o *InvoiceDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceDto) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *InvoiceDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetDocumentNumber returns the DocumentNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceDto) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDto) GetDocumentNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// SetDocumentNumber sets field value
func (o *InvoiceDto) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}

// GetOrganizationSubscriptionId returns the OrganizationSubscriptionId field value
func (o *InvoiceDto) GetOrganizationSubscriptionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationSubscriptionId
}

// GetOrganizationSubscriptionIdOk returns a tuple with the OrganizationSubscriptionId field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetOrganizationSubscriptionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationSubscriptionId, true
}

// SetOrganizationSubscriptionId sets field value
func (o *InvoiceDto) SetOrganizationSubscriptionId(v int32) {
	o.OrganizationSubscriptionId = v
}

// GetIsPaid returns the IsPaid field value
func (o *InvoiceDto) GetIsPaid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPaid
}

// GetIsPaidOk returns a tuple with the IsPaid field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetIsPaidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPaid, true
}

// SetIsPaid sets field value
func (o *InvoiceDto) SetIsPaid(v bool) {
	o.IsPaid = v
}

// GetRequiredPaymentAction returns the RequiredPaymentAction field value
func (o *InvoiceDto) GetRequiredPaymentAction() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequiredPaymentAction
}

// GetRequiredPaymentActionOk returns a tuple with the RequiredPaymentAction field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetRequiredPaymentActionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiredPaymentAction, true
}

// SetRequiredPaymentAction sets field value
func (o *InvoiceDto) SetRequiredPaymentAction(v bool) {
	o.RequiredPaymentAction = v
}

// GetStripeInvoiceId returns the StripeInvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoiceDto) GetStripeInvoiceId() string {
	if o == nil || o.StripeInvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.StripeInvoiceId.Get()
}

// GetStripeInvoiceIdOk returns a tuple with the StripeInvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceDto) GetStripeInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StripeInvoiceId.Get(), o.StripeInvoiceId.IsSet()
}

// SetStripeInvoiceId sets field value
func (o *InvoiceDto) SetStripeInvoiceId(v string) {
	o.StripeInvoiceId.Set(&v)
}

// GetPrice returns the Price field value
func (o *InvoiceDto) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InvoiceDto) SetPrice(v float64) {
	o.Price = v
}

// GetStartDate returns the StartDate field value
func (o *InvoiceDto) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *InvoiceDto) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *InvoiceDto) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *InvoiceDto) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetDueDate returns the DueDate field value
func (o *InvoiceDto) GetDueDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *InvoiceDto) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *InvoiceDto) SetDueDate(v time.Time) {
	o.DueDate = v
}

func (o InvoiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name.Get()
	toSerialize["documentNumber"] = o.DocumentNumber.Get()
	toSerialize["organizationSubscriptionId"] = o.OrganizationSubscriptionId
	toSerialize["isPaid"] = o.IsPaid
	toSerialize["requiredPaymentAction"] = o.RequiredPaymentAction
	toSerialize["stripeInvoiceId"] = o.StripeInvoiceId.Get()
	toSerialize["price"] = o.Price
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	toSerialize["dueDate"] = o.DueDate
	return toSerialize, nil
}

func (o *InvoiceDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"documentNumber",
		"organizationSubscriptionId",
		"isPaid",
		"requiredPaymentAction",
		"stripeInvoiceId",
		"price",
		"startDate",
		"endDate",
		"dueDate",
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

	varInvoiceDto := _InvoiceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvoiceDto)

	if err != nil {
		return err
	}

	*o = InvoiceDto(varInvoiceDto)

	return err
}

type NullableInvoiceDto struct {
	value *InvoiceDto
	isSet bool
}

func (v NullableInvoiceDto) Get() *InvoiceDto {
	return v.value
}

func (v *NullableInvoiceDto) Set(val *InvoiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDto(val *InvoiceDto) *NullableInvoiceDto {
	return &NullableInvoiceDto{value: val, isSet: true}
}

func (v NullableInvoiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


