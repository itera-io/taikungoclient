// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdatePaymentIDCommand update payment Id command
//
// swagger:model UpdatePaymentIdCommand
type UpdatePaymentIDCommand struct {

	// payment intent Id
	PaymentIntentID string `json:"paymentIntentId,omitempty"`

	// payment method Id
	PaymentMethodID string `json:"paymentMethodId,omitempty"`
}

// Validate validates this update payment Id command
func (m *UpdatePaymentIDCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update payment Id command based on context it is used
func (m *UpdatePaymentIDCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePaymentIDCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePaymentIDCommand) UnmarshalBinary(b []byte) error {
	var res UpdatePaymentIDCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
