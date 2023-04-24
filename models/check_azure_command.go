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

// CheckAzureCommand check azure command
//
// swagger:model CheckAzureCommand
type CheckAzureCommand struct {

	// azure client Id
	// Required: true
	// Min Length: 1
	AzureClientID *string `json:"azureClientId"`

	// azure client secret
	// Required: true
	// Min Length: 1
	AzureClientSecret *string `json:"azureClientSecret"`

	// azure tenant Id
	// Required: true
	// Min Length: 1
	AzureTenantID *string `json:"azureTenantId"`
}

// Validate validates this check azure command
func (m *CheckAzureCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckAzureCommand) validateAzureClientID(formats strfmt.Registry) error {

	if err := validate.Required("azureClientId", "body", m.AzureClientID); err != nil {
		return err
	}

	if err := validate.MinLength("azureClientId", "body", *m.AzureClientID, 1); err != nil {
		return err
	}

	return nil
}

func (m *CheckAzureCommand) validateAzureClientSecret(formats strfmt.Registry) error {

	if err := validate.Required("azureClientSecret", "body", m.AzureClientSecret); err != nil {
		return err
	}

	if err := validate.MinLength("azureClientSecret", "body", *m.AzureClientSecret, 1); err != nil {
		return err
	}

	return nil
}

func (m *CheckAzureCommand) validateAzureTenantID(formats strfmt.Registry) error {

	if err := validate.Required("azureTenantId", "body", m.AzureTenantID); err != nil {
		return err
	}

	if err := validate.MinLength("azureTenantId", "body", *m.AzureTenantID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this check azure command based on context it is used
func (m *CheckAzureCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckAzureCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckAzureCommand) UnmarshalBinary(b []byte) error {
	var res CheckAzureCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
