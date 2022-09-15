// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MonitoringOperationsCommand monitoring operations command
//
// swagger:model MonitoringOperationsCommand
type MonitoringOperationsCommand struct {

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this monitoring operations command
func (m *MonitoringOperationsCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this monitoring operations command based on context it is used
func (m *MonitoringOperationsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MonitoringOperationsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MonitoringOperationsCommand) UnmarshalBinary(b []byte) error {
	var res MonitoringOperationsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
