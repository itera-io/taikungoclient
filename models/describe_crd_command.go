// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeCrdCommand describe crd command
//
// swagger:model DescribeCrdCommand
type DescribeCrdCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this describe crd command
func (m *DescribeCrdCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe crd command based on context it is used
func (m *DescribeCrdCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeCrdCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeCrdCommand) UnmarshalBinary(b []byte) error {
	var res DescribeCrdCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
