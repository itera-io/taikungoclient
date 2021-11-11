// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfirmEmailCommand confirm email command
//
// swagger:model ConfirmEmailCommand
type ConfirmEmailCommand struct {

	// mode
	Mode string `json:"mode,omitempty"`

	// new email
	NewEmail string `json:"newEmail,omitempty"`
}

// Validate validates this confirm email command
func (m *ConfirmEmailCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this confirm email command based on context it is used
func (m *ConfirmEmailCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfirmEmailCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfirmEmailCommand) UnmarshalBinary(b []byte) error {
	var res ConfirmEmailCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
