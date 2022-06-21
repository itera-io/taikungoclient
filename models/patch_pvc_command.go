// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchPvcCommand patch pvc command
//
// swagger:model PatchPvcCommand
type PatchPvcCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// yaml
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this patch pvc command
func (m *PatchPvcCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch pvc command based on context it is used
func (m *PatchPvcCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchPvcCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchPvcCommand) UnmarshalBinary(b []byte) error {
	var res PatchPvcCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
