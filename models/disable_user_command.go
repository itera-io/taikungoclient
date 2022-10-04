// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DisableUserCommand disable user command
//
// swagger:model DisableUserCommand
type DisableUserCommand struct {

	// disable
	Disable bool `json:"disable"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this disable user command
func (m *DisableUserCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disable user command based on context it is used
func (m *DisableUserCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DisableUserCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisableUserCommand) UnmarshalBinary(b []byte) error {
	var res DisableUserCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}