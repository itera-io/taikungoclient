// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminProjectsResponseData admin projects response data
//
// swagger:model AdminProjectsResponseData
type AdminProjectsResponseData struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// kubernetes current version
	KubernetesCurrentVersion string `json:"kubernetesCurrentVersion,omitempty"`

	// kubespray current version
	KubesprayCurrentVersion string `json:"kubesprayCurrentVersion,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// servers count
	ServersCount int32 `json:"serversCount,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tcu
	Tcu int32 `json:"tcu,omitempty"`
}

// Validate validates this admin projects response data
func (m *AdminProjectsResponseData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin projects response data based on context it is used
func (m *AdminProjectsResponseData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminProjectsResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminProjectsResponseData) UnmarshalBinary(b []byte) error {
	var res AdminProjectsResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
