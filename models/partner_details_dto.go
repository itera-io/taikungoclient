// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PartnerDetailsDto partner details dto
//
// swagger:model PartnerDetailsDto
type PartnerDetailsDto struct {

	// address
	Address string `json:"address,omitempty"`

	// allow registration
	AllowRegistration bool `json:"allowRegistration,omitempty"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organizations
	Organizations []*CommonDropdownDto `json:"organizations"`

	// payment enabled
	PaymentEnabled bool `json:"paymentEnabled,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// required user approval
	RequiredUserApproval bool `json:"requiredUserApproval,omitempty"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`

	// white list domains
	WhiteListDomains []*WhiteListDomainDto `json:"whiteListDomains"`
}

// Validate validates this partner details dto
func (m *PartnerDetailsDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhiteListDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartnerDetailsDto) validateOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {
		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PartnerDetailsDto) validateWhiteListDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.WhiteListDomains) { // not required
		return nil
	}

	for i := 0; i < len(m.WhiteListDomains); i++ {
		if swag.IsZero(m.WhiteListDomains[i]) { // not required
			continue
		}

		if m.WhiteListDomains[i] != nil {
			if err := m.WhiteListDomains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("whiteListDomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this partner details dto based on the context it is used
func (m *PartnerDetailsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhiteListDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PartnerDetailsDto) contextValidateOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Organizations); i++ {

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PartnerDetailsDto) contextValidateWhiteListDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WhiteListDomains); i++ {

		if m.WhiteListDomains[i] != nil {
			if err := m.WhiteListDomains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("whiteListDomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PartnerDetailsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartnerDetailsDto) UnmarshalBinary(b []byte) error {
	var res PartnerDetailsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
