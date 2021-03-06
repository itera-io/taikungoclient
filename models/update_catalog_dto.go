// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateCatalogDto update catalog dto
//
// swagger:model UpdateCatalogDto
type UpdateCatalogDto struct {

	// is bound
	IsBound bool `json:"isBound"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this update catalog dto
func (m *UpdateCatalogDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update catalog dto based on context it is used
func (m *UpdateCatalogDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCatalogDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCatalogDto) UnmarshalBinary(b []byte) error {
	var res UpdateCatalogDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
