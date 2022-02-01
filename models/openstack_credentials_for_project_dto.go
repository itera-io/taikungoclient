// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackCredentialsForProjectDto openstack credentials for project dto
//
// swagger:model OpenstackCredentialsForProjectDto
type OpenstackCredentialsForProjectDto struct {

	// open stack availability zone
	OpenStackAvailabilityZone string `json:"openStackAvailabilityZone,omitempty"`

	// open stack domain
	OpenStackDomain string `json:"openStackDomain,omitempty"`

	// open stack import network
	OpenStackImportNetwork bool `json:"openStackImportNetwork"`

	// open stack internal subnet Id
	OpenStackInternalSubnetID string `json:"openStackInternalSubnetId,omitempty"`

	// open stack password
	OpenStackPassword string `json:"openStackPassword,omitempty"`

	// open stack project
	OpenStackProject string `json:"openStackProject,omitempty"`

	// open stack public network
	OpenStackPublicNetwork string `json:"openStackPublicNetwork,omitempty"`

	// open stack region
	OpenStackRegion string `json:"openStackRegion,omitempty"`

	// open stack tenant Id
	OpenStackTenantID string `json:"openStackTenantId,omitempty"`

	// open stack Url
	OpenStackURL string `json:"openStackUrl,omitempty"`

	// open stack user
	OpenStackUser string `json:"openStackUser,omitempty"`

	// open stack volume type
	OpenStackVolumeType string `json:"openStackVolumeType,omitempty"`
}

// Validate validates this openstack credentials for project dto
func (m *OpenstackCredentialsForProjectDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack credentials for project dto based on context it is used
func (m *OpenstackCredentialsForProjectDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackCredentialsForProjectDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackCredentialsForProjectDto) UnmarshalBinary(b []byte) error {
	var res OpenstackCredentialsForProjectDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
