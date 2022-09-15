// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectAppParamDto project app param dto
//
// swagger:model ProjectAppParamDto
type ProjectAppParamDto struct {

	// is changeable
	IsChangeable bool `json:"isChangeable"`

	// is readonly
	IsReadonly bool `json:"isReadonly"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this project app param dto
func (m *ProjectAppParamDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project app param dto based on context it is used
func (m *ProjectAppParamDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectAppParamDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectAppParamDto) UnmarshalBinary(b []byte) error {
	var res ProjectAppParamDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
