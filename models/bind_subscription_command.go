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

// BindSubscriptionCommand bind subscription command
//
// swagger:model BindSubscriptionCommand
type BindSubscriptionCommand struct {

	// id
	// Required: true
	// Minimum: > 0
	ID *int32 `json:"id"`

	// yearly
	Yearly bool `json:"yearly"`
}

// Validate validates this bind subscription command
func (m *BindSubscriptionCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BindSubscriptionCommand) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 0, true); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bind subscription command based on context it is used
func (m *BindSubscriptionCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BindSubscriptionCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BindSubscriptionCommand) UnmarshalBinary(b []byte) error {
	var res BindSubscriptionCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
