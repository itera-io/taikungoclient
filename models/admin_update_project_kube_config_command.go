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

// AdminUpdateProjectKubeConfigCommand admin update project kube config command
//
// swagger:model AdminUpdateProjectKubeConfigCommand
type AdminUpdateProjectKubeConfigCommand struct {

	// project Id
	// Required: true
	// Minimum: > 0
	ProjectID *int32 `json:"projectId"`
}

// Validate validates this admin update project kube config command
func (m *AdminUpdateProjectKubeConfigCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminUpdateProjectKubeConfigCommand) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("projectId", "body", int64(*m.ProjectID), 0, true); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this admin update project kube config command based on context it is used
func (m *AdminUpdateProjectKubeConfigCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminUpdateProjectKubeConfigCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminUpdateProjectKubeConfigCommand) UnmarshalBinary(b []byte) error {
	var res AdminUpdateProjectKubeConfigCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
