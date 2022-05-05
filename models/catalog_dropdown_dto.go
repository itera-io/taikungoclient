// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogDropdownDto catalog dropdown dto
//
// swagger:model CatalogDropdownDto
type CatalogDropdownDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// package ids
	PackageIds []string `json:"packageIds"`
}

// Validate validates this catalog dropdown dto
func (m *CatalogDropdownDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog dropdown dto based on context it is used
func (m *CatalogDropdownDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogDropdownDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogDropdownDto) UnmarshalBinary(b []byte) error {
	var res CatalogDropdownDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
