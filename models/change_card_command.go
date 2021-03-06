// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChangeCardCommand change card command
//
// swagger:model ChangeCardCommand
type ChangeCardCommand struct {

	// payment method Id
	PaymentMethodID string `json:"paymentMethodId,omitempty"`
}

// Validate validates this change card command
func (m *ChangeCardCommand) Validate(formats strfmt.Registry) error {
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
