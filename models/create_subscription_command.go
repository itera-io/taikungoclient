// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateSubscriptionCommand create subscription command
//
// swagger:model CreateSubscriptionCommand
type CreateSubscriptionCommand struct {

	// cloud credential limit
	CloudCredentialLimit int32 `json:"cloudCredentialLimit,omitempty"`

	// monthly price
	MonthlyPrice float64 `json:"monthlyPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project limit
	ProjectLimit int32 `json:"projectLimit,omitempty"`

	// server limit
	ServerLimit int32 `json:"serverLimit,omitempty"`

	// tcu price
	TcuPrice float64 `json:"tcuPrice,omitempty"`

	// trial days
	TrialDays int32 `json:"trialDays,omitempty"`

	// user limit
	UserLimit int32 `json:"userLimit,omitempty"`

	// yearly price
	YearlyPrice float64 `json:"yearlyPrice,omitempty"`
}

// Validate validates this create subscription command
func (m *CreateSubscriptionCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create subscription command based on context it is used
func (m *CreateSubscriptionCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSubscriptionCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSubscriptionCommand) UnmarshalBinary(b []byte) error {
	var res CreateSubscriptionCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
