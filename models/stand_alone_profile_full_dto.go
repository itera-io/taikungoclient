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

// StandAloneProfileFullDto stand alone profile full dto
//
// swagger:model StandAloneProfileFullDto
type StandAloneProfileFullDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// stand alone profile security groups
	StandAloneProfileSecurityGroups []*StandAloneProfileSecurityGroupFullDto `json:"standAloneProfileSecurityGroups"`
}

// Validate validates this stand alone profile full dto
func (m *StandAloneProfileFullDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandAloneProfileSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandAloneProfileFullDto) validateStandAloneProfileSecurityGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.StandAloneProfileSecurityGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.StandAloneProfileSecurityGroups); i++ {
		if swag.IsZero(m.StandAloneProfileSecurityGroups[i]) { // not required
			continue
		}

		if m.StandAloneProfileSecurityGroups[i] != nil {
			if err := m.StandAloneProfileSecurityGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standAloneProfileSecurityGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standAloneProfileSecurityGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stand alone profile full dto based on the context it is used
func (m *StandAloneProfileFullDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStandAloneProfileSecurityGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandAloneProfileFullDto) contextValidateStandAloneProfileSecurityGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StandAloneProfileSecurityGroups); i++ {

		if m.StandAloneProfileSecurityGroups[i] != nil {
			if err := m.StandAloneProfileSecurityGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standAloneProfileSecurityGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standAloneProfileSecurityGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandAloneProfileFullDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandAloneProfileFullDto) UnmarshalBinary(b []byte) error {
	var res StandAloneProfileFullDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
