// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationEntityForDashboard organization entity for dashboard
//
// swagger:model OrganizationEntityForDashboard
type OrganizationEntityForDashboard struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// users
	Users int32 `json:"users,omitempty"`
}

// Validate validates this organization entity for dashboard
func (m *OrganizationEntityForDashboard) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization entity for dashboard based on context it is used
func (m *OrganizationEntityForDashboard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationEntityForDashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationEntityForDashboard) UnmarshalBinary(b []byte) error {
	var res OrganizationEntityForDashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
