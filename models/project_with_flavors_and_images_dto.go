// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectWithFlavorsAndImagesDto project with flavors and images dto
//
// swagger:model ProjectWithFlavorsAndImagesDto
type ProjectWithFlavorsAndImagesDto struct {

	// flavors
	Flavors []string `json:"flavors"`

	// id
	ID int32 `json:"id,omitempty"`

	// images
	Images []string `json:"images"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this project with flavors and images dto
func (m *ProjectWithFlavorsAndImagesDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project with flavors and images dto based on context it is used
func (m *ProjectWithFlavorsAndImagesDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectWithFlavorsAndImagesDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectWithFlavorsAndImagesDto) UnmarshalBinary(b []byte) error {
	var res ProjectWithFlavorsAndImagesDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}