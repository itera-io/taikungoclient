// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteScheduleCommand delete schedule command
//
// swagger:model DeleteScheduleCommand
type DeleteScheduleCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this delete schedule command
func (m *DeleteScheduleCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete schedule command based on context it is used
func (m *DeleteScheduleCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteScheduleCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteScheduleCommand) UnmarshalBinary(b []byte) error {
	var res DeleteScheduleCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
