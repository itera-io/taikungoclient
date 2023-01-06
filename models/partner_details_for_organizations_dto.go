// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PartnerDetailsForOrganizationsDto partner details for organizations dto
//
// swagger:model PartnerDetailsForOrganizationsDto
type PartnerDetailsForOrganizationsDto struct {

	// domain
	Domain string `json:"domain,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this partner details for organizations dto
func (m *PartnerDetailsForOrganizationsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner details for organizations dto based on context it is used
func (m *PartnerDetailsForOrganizationsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartnerDetailsForOrganizationsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartnerDetailsForOrganizationsDto) UnmarshalBinary(b []byte) error {
	var res PartnerDetailsForOrganizationsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}