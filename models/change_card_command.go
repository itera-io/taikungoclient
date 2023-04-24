// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeCardCommand change card command
//
// swagger:model ChangeCardCommand
type ChangeCardCommand struct {

	// payment method Id
	// Required: true
	// Min Length: 1
	PaymentMethodID *string `json:"paymentMethodId"`
}

// Validate validates this change card command
func (m *ChangeCardCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethodID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeCardCommand) validatePaymentMethodID(formats strfmt.Registry) error {

	if err := validate.Required("paymentMethodId", "body", m.PaymentMethodID); err != nil {
		return err
	}

	if err := validate.MinLength("paymentMethodId", "body", *m.PaymentMethodID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this change card command based on context it is used
func (m *ChangeCardCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeCardCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeCardCommand) UnmarshalBinary(b []byte) error {
	var res ChangeCardCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
