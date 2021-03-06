// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RebootStandAloneVMCommand reboot stand alone Vm command
//
// swagger:model RebootStandAloneVmCommand
type RebootStandAloneVMCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this reboot stand alone Vm command
func (m *RebootStandAloneVMCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reboot stand alone Vm command based on context it is used
func (m *RebootStandAloneVMCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RebootStandAloneVMCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RebootStandAloneVMCommand) UnmarshalBinary(b []byte) error {
	var res RebootStandAloneVMCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
