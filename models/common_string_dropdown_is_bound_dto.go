// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CommonStringDropdownIsBoundDto common string dropdown is bound dto
//
// swagger:model CommonStringDropdownIsBoundDto
type CommonStringDropdownIsBoundDto struct {

	// id
	ID string `json:"id,omitempty"`

	// is bound
	IsBound bool `json:"isBound"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this common string dropdown is bound dto
func (m *CommonStringDropdownIsBoundDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this common string dropdown is bound dto based on context it is used
func (m *CommonStringDropdownIsBoundDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonStringDropdownIsBoundDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonStringDropdownIsBoundDto) UnmarshalBinary(b []byte) error {
	var res CommonStringDropdownIsBoundDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
