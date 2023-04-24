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

// CreateStripeCustomerCommand create stripe customer command
//
// swagger:model CreateStripeCustomerCommand
type CreateStripeCustomerCommand struct {

	// address
	Address string `json:"address,omitempty"`

	// billing email
	BillingEmail string `json:"billingEmail,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	// Required: true
	// Max Length: 90
	// Min Length: 3
	Country *string `json:"country"`

	// legal name
	LegalName string `json:"legalName,omitempty"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`
}

// Validate validates this create stripe customer command
func (m *CreateStripeCustomerCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateStripeCustomerCommand) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("country", "body", *m.Country, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", *m.Country, 90); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create stripe customer command based on context it is used
func (m *CreateStripeCustomerCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateStripeCustomerCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStripeCustomerCommand) UnmarshalBinary(b []byte) error {
	var res CreateStripeCustomerCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
