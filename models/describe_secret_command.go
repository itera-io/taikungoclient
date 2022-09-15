// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeSecretCommand describe secret command
//
// swagger:model DescribeSecretCommand
type DescribeSecretCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this describe secret command
func (m *DescribeSecretCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe secret command based on context it is used
func (m *DescribeSecretCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeSecretCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeSecretCommand) UnmarshalBinary(b []byte) error {
	var res DescribeSecretCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
