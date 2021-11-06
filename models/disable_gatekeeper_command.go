// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DisableGatekeeperCommand disable gatekeeper command
//
// swagger:model DisableGatekeeperCommand
type DisableGatekeeperCommand struct {

	// opa profile Id
	OpaProfileID int32 `json:"opaProfileId,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this disable gatekeeper command
func (m *DisableGatekeeperCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disable gatekeeper command based on context it is used
func (m *DisableGatekeeperCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DisableGatekeeperCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisableGatekeeperCommand) UnmarshalBinary(b []byte) error {
	var res DisableGatekeeperCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
