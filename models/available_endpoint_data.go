// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AvailableEndpointData available endpoint data
//
// swagger:model AvailableEndpointData
type AvailableEndpointData struct {

	// description
	Description string `json:"description,omitempty"`

	// http method
	HTTPMethod string `json:"httpMethod,omitempty"`

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this available endpoint data
func (m *AvailableEndpointData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this available endpoint data based on context it is used
func (m *AvailableEndpointData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailableEndpointData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableEndpointData) UnmarshalBinary(b []byte) error {
	var res AvailableEndpointData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
