// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StandAloneProfileSecurityGroupFullDto stand alone profile security group full dto
//
// swagger:model StandAloneProfileSecurityGroupFullDto
type StandAloneProfileSecurityGroupFullDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// port max range
	PortMaxRange int32 `json:"portMaxRange,omitempty"`

	// port min range
	PortMinRange int32 `json:"portMinRange,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// remote Ip prefix
	RemoteIPPrefix string `json:"remoteIpPrefix,omitempty"`
}

// Validate validates this stand alone profile security group full dto
func (m *StandAloneProfileSecurityGroupFullDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone profile security group full dto based on context it is used
func (m *StandAloneProfileSecurityGroupFullDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StandAloneProfileSecurityGroupFullDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandAloneProfileSecurityGroupFullDto) UnmarshalBinary(b []byte) error {
	var res StandAloneProfileSecurityGroupFullDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
