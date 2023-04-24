// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListForLandingPageDto list for landing page dto
//
// swagger:model ListForLandingPageDto
type ListForLandingPageDto struct {

	// cloud credential limit
	CloudCredentialLimit int32 `json:"cloudCredentialLimit,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is deprecated
	IsDeprecated bool `json:"isDeprecated"`

	// is enterprise
	IsEnterprise bool `json:"isEnterprise"`

	// is free
	IsFree bool `json:"isFree"`

	// monthly price
	MonthlyPrice float64 `json:"monthlyPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// partner Id
	PartnerID int32 `json:"partnerId,omitempty"`

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

// Validate validates this list for landing page dto
func (m *ListForLandingPageDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list for landing page dto based on context it is used
func (m *ListForLandingPageDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListForLandingPageDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListForLandingPageDto) UnmarshalBinary(b []byte) error {
	var res ListForLandingPageDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
