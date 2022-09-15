// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CidrCommand cidr command
//
// swagger:model CidrCommand
type CidrCommand struct {

	// cidr
	Cidr string `json:"cidr,omitempty"`
}

// Validate validates this cidr command
func (m *CidrCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cidr command based on context it is used
func (m *CidrCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CidrCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CidrCommand) UnmarshalBinary(b []byte) error {
	var res CidrCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
