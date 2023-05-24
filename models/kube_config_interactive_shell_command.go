// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubeConfigInteractiveShellCommand kube config interactive shell command
//
// swagger:model KubeConfigInteractiveShellCommand
type KubeConfigInteractiveShellCommand struct {

	// kube config Id
	// Required: true
	// Minimum: > 0
	KubeConfigID *int32 `json:"kubeConfigId"`

	// project Id
	// Required: true
	// Minimum: > 0
	ProjectID *int32 `json:"projectId"`

	// token
	// Required: true
	// Min Length: 1
	Token *string `json:"token"`
}

// Validate validates this kube config interactive shell command
func (m *KubeConfigInteractiveShellCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKubeConfigID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeConfigInteractiveShellCommand) validateKubeConfigID(formats strfmt.Registry) error {

	if err := validate.Required("kubeConfigId", "body", m.KubeConfigID); err != nil {
		return err
	}

	if err := validate.MinimumInt("kubeConfigId", "body", int64(*m.KubeConfigID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *KubeConfigInteractiveShellCommand) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("projectId", "body", int64(*m.ProjectID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *KubeConfigInteractiveShellCommand) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if err := validate.MinLength("token", "body", *m.Token, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kube config interactive shell command based on context it is used
func (m *KubeConfigInteractiveShellCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeConfigInteractiveShellCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeConfigInteractiveShellCommand) UnmarshalBinary(b []byte) error {
	var res KubeConfigInteractiveShellCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}