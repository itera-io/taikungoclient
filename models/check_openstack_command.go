// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CheckOpenstackCommand check openstack command
//
// swagger:model CheckOpenstackCommand
type CheckOpenstackCommand struct {

	// is admin
	IsAdmin bool `json:"isAdmin"`

	// open stack domain
	OpenStackDomain string `json:"openStackDomain,omitempty"`

	// open stack password
	OpenStackPassword string `json:"openStackPassword,omitempty"`

	// open stack Url
	OpenStackURL string `json:"openStackUrl,omitempty"`

	// open stack user
	OpenStackUser string `json:"openStackUser,omitempty"`
}

// Validate validates this check openstack command
func (m *CheckOpenstackCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this check openstack command based on context it is used
func (m *CheckOpenstackCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckOpenstackCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckOpenstackCommand) UnmarshalBinary(b []byte) error {
	var res CheckOpenstackCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
