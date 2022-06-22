// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateAccessProfileDto update access profile dto
//
// swagger:model UpdateAccessProfileDto
type UpdateAccessProfileDto struct {

	// http proxy
	HTTPProxy string `json:"httpProxy,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update access profile dto
func (m *UpdateAccessProfileDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update access profile dto based on context it is used
func (m *UpdateAccessProfileDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateAccessProfileDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateAccessProfileDto) UnmarshalBinary(b []byte) error {
	var res UpdateAccessProfileDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}