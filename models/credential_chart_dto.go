// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CredentialChartDto credential chart dto
//
// swagger:model CredentialChartDto
type CredentialChartDto struct {

	// aws
	Aws int32 `json:"aws,omitempty"`

	// azure
	Azure int32 `json:"azure,omitempty"`

	// google
	Google int32 `json:"google,omitempty"`

	// openstack
	Openstack int32 `json:"openstack,omitempty"`

	// tanzu
	Tanzu int32 `json:"tanzu,omitempty"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this credential chart dto
func (m *CredentialChartDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this credential chart dto based on context it is used
func (m *CredentialChartDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CredentialChartDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialChartDto) UnmarshalBinary(b []byte) error {
	var res CredentialChartDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
