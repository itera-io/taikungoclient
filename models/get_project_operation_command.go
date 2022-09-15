// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetProjectOperationCommand get project operation command
//
// swagger:model GetProjectOperationCommand
type GetProjectOperationCommand struct {

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this get project operation command
func (m *GetProjectOperationCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get project operation command based on context it is used
func (m *GetProjectOperationCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetProjectOperationCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProjectOperationCommand) UnmarshalBinary(b []byte) error {
	var res GetProjectOperationCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
