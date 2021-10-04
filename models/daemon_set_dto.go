// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DaemonSetDto daemon set dto
//
// swagger:model DaemonSetDto
type DaemonSetDto struct {

	// age
	Age string `json:"age,omitempty"`

	// metadata name
	MetadataName string `json:"metadataName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this daemon set dto
func (m *DaemonSetDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this daemon set dto based on context it is used
func (m *DaemonSetDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DaemonSetDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DaemonSetDto) UnmarshalBinary(b []byte) error {
	var res DaemonSetDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
