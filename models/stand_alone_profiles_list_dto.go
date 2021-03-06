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

// StandAloneProfilesListDto stand alone profiles list dto
//
// swagger:model StandAloneProfilesListDto
type StandAloneProfilesListDto struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// partner logo
	PartnerLogo string `json:"partnerLogo,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// standalone vms
	StandaloneVms []*StandAloneVMSmallDetailDto `json:"standaloneVms"`
}

// Validate validates this stand alone profiles list dto
func (m *StandAloneProfilesListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandaloneVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandAloneProfilesListDto) validateStandaloneVms(formats strfmt.Registry) error {
	if swag.IsZero(m.StandaloneVms) { // not required
		return nil
	}

	for i := 0; i < len(m.StandaloneVms); i++ {
		if swag.IsZero(m.StandaloneVms[i]) { // not required
			continue
		}

		if m.StandaloneVms[i] != nil {
			if err := m.StandaloneVms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stand alone profiles list dto based on the context it is used
func (m *StandAloneProfilesListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStandaloneVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandAloneProfilesListDto) contextValidateStandaloneVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StandaloneVms); i++ {

		if m.StandaloneVms[i] != nil {
			if err := m.StandaloneVms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandAloneProfilesListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandAloneProfilesListDto) UnmarshalBinary(b []byte) error {
	var res StandAloneProfilesListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
