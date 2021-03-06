// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudCredentialsResponseData cloud credentials response data
//
// swagger:model CloudCredentialsResponseData
type CloudCredentialsResponseData struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`
}

// Validate validates this cloud credentials response data
func (m *CloudCredentialsResponseData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud credentials response data based on context it is used
func (m *CloudCredentialsResponseData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudCredentialsResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudCredentialsResponseData) UnmarshalBinary(b []byte) error {
	var res CloudCredentialsResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
