// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateProjectFromTemplateCommand create project from template command
//
// swagger:model CreateProjectFromTemplateCommand
type CreateProjectFromTemplateCommand struct {

	// can commit
	CanCommit bool `json:"canCommit"`

	// id
	ID int32 `json:"id,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this create project from template command
func (m *CreateProjectFromTemplateCommand) Validate(formats strfmt.Registry) error {
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