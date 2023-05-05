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

// SecretSearchCommand secret search command
//
// swagger:model SecretSearchCommand
type SecretSearchCommand struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// search term
	// Required: true
	// Max Length: 100
	// Min Length: 3
	SearchTerm *string `json:"searchTerm"`
}

// Validate validates this secret search command
func (m *SecretSearchCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSearchTerm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SecretSearchCommand) validateSearchTerm(formats strfmt.Registry) error {

	if err := validate.Required("searchTerm", "body", m.SearchTerm); err != nil {
		return err
	}

	if err := validate.MinLength("searchTerm", "body", *m.SearchTerm, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("searchTerm", "body", *m.SearchTerm, 100); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this secret search command based on context it is used
func (m *SecretSearchCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SecretSearchCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecretSearchCommand) UnmarshalBinary(b []byte) error {
	var res SecretSearchCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
