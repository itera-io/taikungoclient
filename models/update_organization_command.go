// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateOrganizationCommand update organization command
//
// swagger:model UpdateOrganizationCommand
type UpdateOrganizationCommand struct {

	// address
	Address string `json:"address,omitempty"`

	// billing email
	BillingEmail string `json:"billingEmail,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// discount rate
	DiscountRate float64 `json:"discountRate"`

	// email
	Email string `json:"email,omitempty"`

	// full name
	FullName string `json:"fullName,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is eligible update subscription
	IsEligibleUpdateSubscription bool `json:"isEligibleUpdateSubscription"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// name
	Name string `json:"name,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`
}

// Validate validates this update organization command
func (m *UpdateOrganizationCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization command based on context it is used
func (m *UpdateOrganizationCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateOrganizationCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateOrganizationCommand) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
