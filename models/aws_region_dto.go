// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsRegionDto aws region dto
//
// swagger:model AwsRegionDto
type AwsRegionDto struct {

	// name
	Name string `json:"name,omitempty"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this aws region dto
func (m *AwsRegionDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws region dto based on context it is used
func (m *AwsRegionDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsRegionDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsRegionDto) UnmarshalBinary(b []byte) error {
	var res AwsRegionDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
