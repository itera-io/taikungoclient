// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogAppParamsDto catalog app params dto
//
// swagger:model CatalogAppParamsDto
type CatalogAppParamsDto struct {

	// is readonly
	IsReadonly bool `json:"isReadonly"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this catalog app params dto
func (m *CatalogAppParamsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog app params dto based on context it is used
func (m *CatalogAppParamsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogAppParamsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogAppParamsDto) UnmarshalBinary(b []byte) error {
	var res CatalogAppParamsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
