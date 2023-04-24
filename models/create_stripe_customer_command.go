// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
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
	Country string `json:"country,omitempty"`

	// legal name
	LegalName string `json:"legalName,omitempty"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`
}

// Validate validates this create stripe customer command
func (m *CreateStripeCustomerCommand) Validate(formats strfmt.Registry) error {
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
