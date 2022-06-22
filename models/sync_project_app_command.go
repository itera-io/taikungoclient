// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SyncProjectAppCommand sync project app command
//
// swagger:model SyncProjectAppCommand
type SyncProjectAppCommand struct {

	// project app Id
	ProjectAppID int32 `json:"projectAppId,omitempty"`
}

// Validate validates this sync project app command
func (m *SyncProjectAppCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sync project app command based on context it is used
func (m *SyncProjectAppCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SyncProjectAppCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncProjectAppCommand) UnmarshalBinary(b []byte) error {
	var res SyncProjectAppCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}