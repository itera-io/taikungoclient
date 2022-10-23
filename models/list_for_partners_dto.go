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

	// description
	Description string `json:"description,omitempty"`

	// exceeded cloud credential
	ExceededCloudCredential bool `json:"exceededCloudCredential"`

	// exceeded project
	ExceededProject bool `json:"exceededProject"`

	// exceeded servers
	ExceededServers bool `json:"exceededServers"`

	// exceeded user
	ExceededUser bool `json:"exceededUser"`

	// id
	ID int32 `json:"id,omitempty"`

	// is active
	IsActive bool `json:"isActive"`

	// is demo
	IsDemo bool `json:"isDemo"`

	// is deprecated
	IsDeprecated bool `json:"isDeprecated"`

	// is yearly
	IsYearly bool `json:"isYearly"`

	// monthly price
	MonthlyPrice float64 `json:"monthlyPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// partner
	Partner *ListForPartnersDtoPartner `json:"partner,omitempty"`

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

// ListForPartnersDtoPartner list for partners dto partner
//
// swagger:model ListForPartnersDtoPartner
type ListForPartnersDtoPartner struct {

	// id
	ID int32 `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this list for partners dto partner
func (m *ListForPartnersDtoPartner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list for partners dto partner based on context it is used
func (m *ListForPartnersDtoPartner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListForPartnersDtoPartner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListForPartnersDtoPartner) UnmarshalBinary(b []byte) error {
	var res ListForPartnersDtoPartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
