// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CountryListDto country list dto
//
// swagger:model CountryListDto
type CountryListDto struct {

	// is eu
	IsEu bool `json:"isEu"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this country list dto
func (m *CountryListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this country list dto based on context it is used
func (m *CountryListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CountryListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountryListDto) UnmarshalBinary(b []byte) error {
	var res CountryListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}