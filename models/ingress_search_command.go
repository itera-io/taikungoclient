// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IngressSearchCommand ingress search command
//
// swagger:model IngressSearchCommand
type IngressSearchCommand struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// search term
	SearchTerm string `json:"searchTerm,omitempty"`
}

// Validate validates this ingress search command
func (m *IngressSearchCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ingress search command based on context it is used
func (m *IngressSearchCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IngressSearchCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngressSearchCommand) UnmarshalBinary(b []byte) error {
	var res IngressSearchCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
