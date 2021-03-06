// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigMapDto config map dto
//
// swagger:model ConfigMapDto
type ConfigMapDto struct {

	// age
	Age string `json:"age,omitempty"`

	// metadata name
	MetadataName string `json:"metadataName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this config map dto
func (m *ConfigMapDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config map dto based on context it is used
func (m *ConfigMapDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigMapDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigMapDto) UnmarshalBinary(b []byte) error {
	var res ConfigMapDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
