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

// TryForFreeCommand try for free command
//
// swagger:model TryForFreeCommand
type TryForFreeCommand struct {

	// email
	// Required: true
	// Max Length: 200
	// Min Length: 3
	// Format: email
	Email *strfmt.Email `json:"email"`

	// organization name
	// Required: true
	// Max Length: 30
	// Min Length: 3
	OrganizationName *string `json:"organizationName"`

	// username
	// Required: true
	// Max Length: 30
	// Min Length: 3
	Username *string `json:"username"`
}

// Validate validates this try for free command
func (m *TryForFreeCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TryForFreeCommand) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.MinLength("email", "body", m.Email.String(), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "body", m.Email.String(), 200); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TryForFreeCommand) validateOrganizationName(formats strfmt.Registry) error {

	if err := validate.Required("organizationName", "body", m.OrganizationName); err != nil {
		return err
	}

	if err := validate.MinLength("organizationName", "body", *m.OrganizationName, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("organizationName", "body", *m.OrganizationName, 30); err != nil {
		return err
	}

	return nil
}

func (m *TryForFreeCommand) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.MinLength("username", "body", *m.Username, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("username", "body", *m.Username, 30); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this try for free command based on context it is used
func (m *TryForFreeCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TryForFreeCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TryForFreeCommand) UnmarshalBinary(b []byte) error {
	var res TryForFreeCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
