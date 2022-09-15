// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteProjectCommand delete project command
//
// swagger:model DeleteProjectCommand
type DeleteProjectCommand struct {

	// is force delete
	IsForceDelete bool `json:"isForceDelete"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this delete project command
func (m *DeleteProjectCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete project command based on context it is used
func (m *DeleteProjectCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteProjectCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteProjectCommand) UnmarshalBinary(b []byte) error {
	var res DeleteProjectCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
