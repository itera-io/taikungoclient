// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListForPartnersDto list for partners dto
//
// swagger:model ListForPartnersDto
type ListForPartnersDto struct {

	// cloud credential limit
	CloudCredentialLimit int32 `json:"cloudCredentialLimit,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`

	// exceeded cloud credential
	ExceededCloudCredential bool `json:"exceededCloudCredential,omitempty"`

	// exceeded project
	ExceededProject bool `json:"exceededProject,omitempty"`

	// exceeded servers
	ExceededServers bool `json:"exceededServers,omitempty"`

	// exceeded user
	ExceededUser bool `json:"exceededUser,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is active
	IsActive bool `json:"isActive,omitempty"`

	// is deprecated
	IsDeprecated bool `json:"isDeprecated,omitempty"`

	// is yearly
	IsYearly bool `json:"isYearly,omitempty"`

	// monthly price
	MonthlyPrice float64 `json:"monthlyPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// partner
	Partner *PartnerDetailsForSubscription `json:"partner,omitempty"`

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

// Validate validates this list for partners dto
func (m *ListForPartnersDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListForPartnersDto) validatePartner(formats strfmt.Registry) error {
	if swag.IsZero(m.Partner) { // not required
		return nil
	}

	if m.Partner != nil {
		if err := m.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this list for partners dto based on the context it is used
func (m *ListForPartnersDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListForPartnersDto) contextValidatePartner(ctx context.Context, formats strfmt.Registry) error {

	if m.Partner != nil {
		if err := m.Partner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListForPartnersDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListForPartnersDto) UnmarshalBinary(b []byte) error {
	var res ListForPartnersDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
