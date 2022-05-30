// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GoogleCredentialForProjectDto google credential for project dto
//
// swagger:model GoogleCredentialForProjectDto
type GoogleCredentialForProjectDto struct {

	// billing account Id
	BillingAccountID string `json:"billingAccountId,omitempty"`

	// config
	Config string `json:"config,omitempty"`

	// folder Id
	FolderID string `json:"folderId,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// zone
	Zone string `json:"zone,omitempty"`
}

// Validate validates this google credential for project dto
func (m *GoogleCredentialForProjectDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google credential for project dto based on context it is used
func (m *GoogleCredentialForProjectDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GoogleCredentialForProjectDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleCredentialForProjectDto) UnmarshalBinary(b []byte) error {
	var res GoogleCredentialForProjectDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}