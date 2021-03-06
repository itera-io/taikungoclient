// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonStringBasedDropdownDto common string based dropdown dto
//
// swagger:model CommonStringBasedDropdownDto
type CommonStringBasedDropdownDto struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this common string based dropdown dto
func (m *CommonStringBasedDropdownDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common string based dropdown dto based on context it is used
func (m *CommonStringBasedDropdownDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonStringBasedDropdownDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonStringBasedDropdownDto) UnmarshalBinary(b []byte) error {
	var res CommonStringBasedDropdownDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
