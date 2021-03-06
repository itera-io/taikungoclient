// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpaProfileSyncCommand opa profile sync command
//
// swagger:model OpaProfileSyncCommand
type OpaProfileSyncCommand struct {

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this opa profile sync command
func (m *OpaProfileSyncCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this opa profile sync command based on context it is used
func (m *OpaProfileSyncCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpaProfileSyncCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpaProfileSyncCommand) UnmarshalBinary(b []byte) error {
	var res OpaProfileSyncCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
