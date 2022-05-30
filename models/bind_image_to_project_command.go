// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BindImageToProjectCommand bind image to project command
//
// swagger:model BindImageToProjectCommand
type BindImageToProjectCommand struct {

	// images
	Images []string `json:"images"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this bind image to project command
func (m *BindImageToProjectCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bind image to project command based on context it is used
func (m *BindImageToProjectCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BindImageToProjectCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BindImageToProjectCommand) UnmarshalBinary(b []byte) error {
	var res BindImageToProjectCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}