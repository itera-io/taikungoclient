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

// CreateProjectFromTemplateCommand create project from template command
//
// swagger:model CreateProjectFromTemplateCommand
type CreateProjectFromTemplateCommand struct {

	// can commit
	CanCommit bool `json:"canCommit"`

	// id
	// Required: true
	// Minimum: > 0
	ID *int32 `json:"id"`

	// project name
	// Required: true
	// Max Length: 30
	// Min Length: 3
	ProjectName *string `json:"projectName"`
}

// Validate validates this create project from template command
func (m *CreateProjectFromTemplateCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProjectFromTemplateCommand) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *CreateProjectFromTemplateCommand) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("projectName", "body", m.ProjectName); err != nil {
		return err
	}

	if err := validate.MinLength("projectName", "body", *m.ProjectName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("projectName", "body", *m.ProjectName, 30); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create project from template command based on context it is used
func (m *CreateProjectFromTemplateCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateProjectFromTemplateCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProjectFromTemplateCommand) UnmarshalBinary(b []byte) error {
	var res CreateProjectFromTemplateCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
