// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenStackZoneListQuery open stack zone list query
//
// swagger:model OpenStackZoneListQuery
type OpenStackZoneListQuery struct {

	// application cred enabled
	ApplicationCredEnabled bool `json:"applicationCredEnabled"`

	// is admin
	IsAdmin bool `json:"isAdmin"`

	// open stack domain
	OpenStackDomain string `json:"openStackDomain,omitempty"`

	// open stack password
	OpenStackPassword string `json:"openStackPassword,omitempty"`

	// open stack region
	OpenStackRegion string `json:"openStackRegion,omitempty"`

	// open stack Url
	OpenStackURL string `json:"openStackUrl,omitempty"`

	// open stack user
	OpenStackUser string `json:"openStackUser,omitempty"`

	// openstack project
	OpenstackProject string `json:"openstackProject,omitempty"`
}

// Validate validates this open stack zone list query
func (m *OpenStackZoneListQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this open stack zone list query based on context it is used
func (m *OpenStackZoneListQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenStackZoneListQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenStackZoneListQuery) UnmarshalBinary(b []byte) error {
	var res OpenStackZoneListQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
