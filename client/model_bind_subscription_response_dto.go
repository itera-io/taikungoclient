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

// checks if the BindSubscriptionResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BindSubscriptionResponseDto{}

// BindSubscriptionResponseDto struct for BindSubscriptionResponseDto
type BindSubscriptionResponseDto struct {
	Status *string `json:"status,omitempty"`
	PaymentIntentClientSecret *string `json:"paymentIntentClientSecret,omitempty"`
	PaymentIntentId *string `json:"paymentIntentId,omitempty"`
	InvoiceFailureCode *string `json:"invoiceFailureCode,omitempty"`
	InvoiceFailureMessage *string `json:"invoiceFailureMessage,omitempty"`
	InvoiceFailureReason *string `json:"invoiceFailureReason,omitempty"`
}

// NewBindSubscriptionResponseDto instantiates a new BindSubscriptionResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBindSubscriptionResponseDto() *BindSubscriptionResponseDto {
	this := BindSubscriptionResponseDto{}
	return &this
}

// NewBindSubscriptionResponseDtoWithDefaults instantiates a new BindSubscriptionResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBindSubscriptionResponseDtoWithDefaults() *BindSubscriptionResponseDto {
	this := BindSubscriptionResponseDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BindSubscriptionResponseDto) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionResponseDto) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BindSubscriptionResponseDto) SetStatus(v string) {
	o.Status = &v
}

// GetPaymentIntentClientSecret returns the PaymentIntentClientSecret field value if set, zero value otherwise.
func (o *BindSubscriptionResponseDto) GetPaymentIntentClientSecret() string {
	if o == nil || IsNil(o.PaymentIntentClientSecret) {
		var ret string
		return ret
	}
	return *o.PaymentIntentClientSecret
}

// GetPaymentIntentClientSecretOk returns a tuple with the PaymentIntentClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionResponseDto) GetPaymentIntentClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentIntentClientSecret) {
		return nil, false
	}
	return o.PaymentIntentClientSecret, true
}

// HasPaymentIntentClientSecret returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasPaymentIntentClientSecret() bool {
	if o != nil && !IsNil(o.PaymentIntentClientSecret) {
		return true
	}

	return false
}

// SetPaymentIntentClientSecret gets a reference to the given string and assigns it to the PaymentIntentClientSecret field.
func (o *BindSubscriptionResponseDto) SetPaymentIntentClientSecret(v string) {
	o.PaymentIntentClientSecret = &v
}

// GetPaymentIntentId returns the PaymentIntentId field value if set, zero value otherwise.
func (o *BindSubscriptionResponseDto) GetPaymentIntentId() string {
	if o == nil || IsNil(o.PaymentIntentId) {
		var ret string
		return ret
	}
	return *o.PaymentIntentId
}

// GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionResponseDto) GetPaymentIntentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentIntentId) {
		return nil, false
	}
	return o.PaymentIntentId, true
}

// HasPaymentIntentId returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasPaymentIntentId() bool {
	if o != nil && !IsNil(o.PaymentIntentId) {
		return true
	}

	return false
}

// SetPaymentIntentId gets a reference to the given string and assigns it to the PaymentIntentId field.
func (o *BindSubscriptionResponseDto) SetPaymentIntentId(v string) {
	o.PaymentIntentId = &v
}

// GetInvoiceFailureCode returns the InvoiceFailureCode field value if set, zero value otherwise.
func (o *BindSubscriptionResponseDto) GetInvoiceFailureCode() string {
	if o == nil || IsNil(o.InvoiceFailureCode) {
		var ret string
		return ret
	}
	return *o.InvoiceFailureCode
}

// GetInvoiceFailureCodeOk returns a tuple with the InvoiceFailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionResponseDto) GetInvoiceFailureCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFailureCode) {
		return nil, false
	}
	return o.InvoiceFailureCode, true
}

// HasInvoiceFailureCode returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasInvoiceFailureCode() bool {
	if o != nil && !IsNil(o.InvoiceFailureCode) {
		return true
	}

	return false
}

// SetInvoiceFailureCode gets a reference to the given string and assigns it to the InvoiceFailureCode field.
func (o *BindSubscriptionResponseDto) SetInvoiceFailureCode(v string) {
	o.InvoiceFailureCode = &v
}

// GetInvoiceFailureMessage returns the InvoiceFailureMessage field value if set, zero value otherwise.
func (o *BindSubscriptionResponseDto) GetInvoiceFailureMessage() string {
	if o == nil || IsNil(o.InvoiceFailureMessage) {
		var ret string
		return ret
	}
	return *o.InvoiceFailureMessage
}

// GetInvoiceFailureMessageOk returns a tuple with the InvoiceFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionResponseDto) GetInvoiceFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFailureMessage) {
		return nil, false
	}
	return o.InvoiceFailureMessage, true
}

// HasInvoiceFailureMessage returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasInvoiceFailureMessage() bool {
	if o != nil && !IsNil(o.InvoiceFailureMessage) {
		return true
	}

	return false
}

// SetInvoiceFailureMessage gets a reference to the given string and assigns it to the InvoiceFailureMessage field.
func (o *BindSubscriptionResponseDto) SetInvoiceFailureMessage(v string) {
	o.InvoiceFailureMessage = &v
}

// GetInvoiceFailureReason returns the InvoiceFailureReason field value if set, zero value otherwise.
func (o *BindSubscriptionResponseDto) GetInvoiceFailureReason() string {
	if o == nil || IsNil(o.InvoiceFailureReason) {
		var ret string
		return ret
	}
	return *o.InvoiceFailureReason
}

// GetInvoiceFailureReasonOk returns a tuple with the InvoiceFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BindSubscriptionResponseDto) GetInvoiceFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceFailureReason) {
		return nil, false
	}
	return o.InvoiceFailureReason, true
}

// HasInvoiceFailureReason returns a boolean if a field has been set.
func (o *BindSubscriptionResponseDto) HasInvoiceFailureReason() bool {
	if o != nil && !IsNil(o.InvoiceFailureReason) {
		return true
	}

	return false
}

// SetInvoiceFailureReason gets a reference to the given string and assigns it to the InvoiceFailureReason field.
func (o *BindSubscriptionResponseDto) SetInvoiceFailureReason(v string) {
	o.InvoiceFailureReason = &v
}

func (o BindSubscriptionResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BindSubscriptionResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.PaymentIntentClientSecret) {
		toSerialize["paymentIntentClientSecret"] = o.PaymentIntentClientSecret
	}
	if !IsNil(o.PaymentIntentId) {
		toSerialize["paymentIntentId"] = o.PaymentIntentId
	}
	if !IsNil(o.InvoiceFailureCode) {
		toSerialize["invoiceFailureCode"] = o.InvoiceFailureCode
	}
	if !IsNil(o.InvoiceFailureMessage) {
		toSerialize["invoiceFailureMessage"] = o.InvoiceFailureMessage
	}
	if !IsNil(o.InvoiceFailureReason) {
		toSerialize["invoiceFailureReason"] = o.InvoiceFailureReason
	}
	return toSerialize, nil
}

type NullableBindSubscriptionResponseDto struct {
	value *BindSubscriptionResponseDto
	isSet bool
}

func (v NullableBindSubscriptionResponseDto) Get() *BindSubscriptionResponseDto {
	return v.value
}

func (v *NullableBindSubscriptionResponseDto) Set(val *BindSubscriptionResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBindSubscriptionResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBindSubscriptionResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindSubscriptionResponseDto(val *BindSubscriptionResponseDto) *NullableBindSubscriptionResponseDto {
	return &NullableBindSubscriptionResponseDto{value: val, isSet: true}
}

func (v NullableBindSubscriptionResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindSubscriptionResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


