// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchCrdCommand patch crd command
//
// swagger:model PatchCrdCommand
type PatchCrdCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// yaml
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this patch crd command
func (m *PatchCrdCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch crd command based on context it is used
func (m *PatchCrdCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchCrdCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchCrdCommand) UnmarshalBinary(b []byte) error {
	var res PatchCrdCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}