// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PurgeStandAloneVMDiskCommand purge stand alone Vm disk command
//
// swagger:model PurgeStandAloneVmDiskCommand
type PurgeStandAloneVMDiskCommand struct {

	// standalone Vm Id
	StandaloneVMID int32 `json:"standaloneVmId,omitempty"`

	// vm disk ids
	VMDiskIds []int32 `json:"vmDiskIds"`
}

// Validate validates this purge stand alone Vm disk command
func (m *PurgeStandAloneVMDiskCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this purge stand alone Vm disk command based on context it is used
func (m *PurgeStandAloneVMDiskCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PurgeStandAloneVMDiskCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PurgeStandAloneVMDiskCommand) UnmarshalBinary(b []byte) error {
	var res PurgeStandAloneVMDiskCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}