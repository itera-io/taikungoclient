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

// StandAloneProfileSecurityGroupDto stand alone profile security group dto
//
// swagger:model StandAloneProfileSecurityGroupDto
type StandAloneProfileSecurityGroupDto struct {

	// name
	Name string `json:"name,omitempty"`

	// port max range
	PortMaxRange int32 `json:"portMaxRange,omitempty"`

	// port min range
	PortMinRange int32 `json:"portMinRange,omitempty"`

	// protocol
	Protocol SecurityGroupProtocol `json:"protocol,omitempty"`

	// remote Ip prefix
	RemoteIPPrefix string `json:"remoteIpPrefix,omitempty"`
}

// Validate validates this stand alone profile security group dto
func (m *StandAloneProfileSecurityGroupDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandAloneProfileSecurityGroupDto) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("protocol")
		}
		return err
	}

	return nil
}

// ContextValidate validate this stand alone profile security group dto based on the context it is used
func (m *StandAloneProfileSecurityGroupDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandAloneProfileSecurityGroupDto) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("protocol")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandAloneProfileSecurityGroupDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandAloneProfileSecurityGroupDto) UnmarshalBinary(b []byte) error {
	var res StandAloneProfileSecurityGroupDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
