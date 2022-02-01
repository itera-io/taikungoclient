// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateStandaloneVMDiskSizeCommand update standalone Vm disk size command
//
// swagger:model UpdateStandaloneVmDiskSizeCommand
type UpdateStandaloneVMDiskSizeCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this update standalone Vm disk size command
func (m *UpdateStandaloneVMDiskSizeCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update standalone Vm disk size command based on context it is used
func (m *UpdateStandaloneVMDiskSizeCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateStandaloneVMDiskSizeCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateStandaloneVMDiskSizeCommand) UnmarshalBinary(b []byte) error {
	var res UpdateStandaloneVMDiskSizeCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
