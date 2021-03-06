// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesPodLogsCommand kubernetes pod logs command
//
// swagger:model KubernetesPodLogsCommand
type KubernetesPodLogsCommand struct {

	// container
	Container string `json:"container,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this kubernetes pod logs command
func (m *KubernetesPodLogsCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes pod logs command based on context it is used
func (m *KubernetesPodLogsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesPodLogsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesPodLogsCommand) UnmarshalBinary(b []byte) error {
	var res KubernetesPodLogsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
