// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TryDemoCommand try demo command
//
// swagger:model TryDemoCommand
type TryDemoCommand struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this try demo command
func (m *TryDemoCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this try demo command based on context it is used
func (m *TryDemoCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TryDemoCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TryDemoCommand) UnmarshalBinary(b []byte) error {
	var res TryDemoCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
