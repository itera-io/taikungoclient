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

// WhiteListDomainCreateCommand white list domain create command
//
// swagger:model WhiteListDomainCreateCommand
type WhiteListDomainCreateCommand struct {

	// partner Id
	PartnerID int32 `json:"partnerId,omitempty"`

	// white list domains
	WhiteListDomains []*WhiteListDomainCreateDto `json:"whiteListDomains"`
}

// Validate validates this white list domain create command
func (m *WhiteListDomainCreateCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWhiteListDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhiteListDomainCreateCommand) validateWhiteListDomains(formats strfmt.Registry) error {
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

// ContextValidate validate this white list domain create command based on the context it is used
func (m *WhiteListDomainCreateCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWhiteListDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WhiteListDomainCreateCommand) contextValidateWhiteListDomains(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WhiteListDomainCreateCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhiteListDomainCreateCommand) UnmarshalBinary(b []byte) error {
	var res WhiteListDomainCreateCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
