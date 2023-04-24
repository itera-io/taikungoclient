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

// UpdateTanzuCommand update tanzu command
//
// swagger:model UpdateTanzuCommand
type UpdateTanzuCommand struct {

	// id
	// Required: true
	// Minimum: > 0
	ID *int32 `json:"id"`

	// name
	// Required: true
	// Max Length: 30
	// Min Length: 3
	Name *string `json:"name"`

	// password
	// Required: true
	// Max Length: 100
	// Min Length: 3
	Password *string `json:"password"`
}

// Validate validates this update tanzu command
func (m *UpdateTanzuCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTanzuCommand) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTanzuCommand) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 30); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTanzuCommand) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", *m.Password, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", *m.Password, 100); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update tanzu command based on context it is used
func (m *UpdateTanzuCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTanzuCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTanzuCommand) UnmarshalBinary(b []byte) error {
	var res UpdateTanzuCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
