// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StopStandaloneVMCommand stop standalone Vm command
//
// swagger:model StopStandaloneVmCommand
type StopStandaloneVMCommand struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this stop standalone Vm command
func (m *StopStandaloneVMCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stop standalone Vm command based on context it is used
func (m *StopStandaloneVMCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StopStandaloneVMCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StopStandaloneVMCommand) UnmarshalBinary(b []byte) error {
	var res StopStandaloneVMCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
