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

// PatchStsCommand patch sts command
//
// swagger:model PatchStsCommand
type PatchStsCommand struct {

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// namespace
	// Required: true
	// Min Length: 1
	Namespace *string `json:"namespace"`

	// project Id
	// Required: true
	// Minimum: > 0
	ProjectID *int32 `json:"projectId"`

	// yaml
	// Required: true
	// Min Length: 1
	Yaml *string `json:"yaml"`
}

// Validate validates this patch sts command
func (m *PatchStsCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYaml(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchStsCommand) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	return nil
}

func (m *PatchStsCommand) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	if err := validate.MinLength("namespace", "body", *m.Namespace, 1); err != nil {
		return err
	}

	return nil
}

func (m *PatchStsCommand) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("projectId", "body", int64(*m.ProjectID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *PatchStsCommand) validateYaml(formats strfmt.Registry) error {

	if err := validate.Required("yaml", "body", m.Yaml); err != nil {
		return err
	}

	if err := validate.MinLength("yaml", "body", *m.Yaml, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch sts command based on context it is used
func (m *PatchStsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchStsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchStsCommand) UnmarshalBinary(b []byte) error {
	var res PatchStsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
