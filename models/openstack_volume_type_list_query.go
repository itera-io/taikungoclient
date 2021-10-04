// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackVolumeTypeListQuery openstack volume type list query
//
// swagger:model OpenstackVolumeTypeListQuery
type OpenstackVolumeTypeListQuery struct {

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
}

// Validate validates this openstack volume type list query
func (m *OpenstackVolumeTypeListQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack volume type list query based on context it is used
func (m *OpenstackVolumeTypeListQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackVolumeTypeListQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackVolumeTypeListQuery) UnmarshalBinary(b []byte) error {
	var res OpenstackVolumeTypeListQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
