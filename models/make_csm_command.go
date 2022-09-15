// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MakeCsmCommand make csm command
//
// swagger:model MakeCsmCommand
type MakeCsmCommand struct {

	// mode
	Mode string `json:"mode,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this make csm command
func (m *MakeCsmCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this make csm command based on context it is used
func (m *MakeCsmCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MakeCsmCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MakeCsmCommand) UnmarshalBinary(b []byte) error {
	var res MakeCsmCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
