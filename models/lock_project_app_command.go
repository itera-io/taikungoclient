// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LockProjectAppCommand lock project app command
//
// swagger:model LockProjectAppCommand
type LockProjectAppCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`
}

// Validate validates this lock project app command
func (m *LockProjectAppCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lock project app command based on context it is used
func (m *LockProjectAppCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LockProjectAppCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockProjectAppCommand) UnmarshalBinary(b []byte) error {
	var res LockProjectAppCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
