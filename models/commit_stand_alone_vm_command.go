// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommitStandAloneVMCommand commit stand alone Vm command
//
// swagger:model CommitStandAloneVmCommand
type CommitStandAloneVMCommand struct {

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this commit stand alone Vm command
func (m *CommitStandAloneVMCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this commit stand alone Vm command based on context it is used
func (m *CommitStandAloneVMCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommitStandAloneVMCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommitStandAloneVMCommand) UnmarshalBinary(b []byte) error {
	var res CommitStandAloneVMCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
