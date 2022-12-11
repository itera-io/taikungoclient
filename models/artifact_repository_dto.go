// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArtifactRepositoryDto artifact repository dto
//
// swagger:model ArtifactRepositoryDto
type ArtifactRepositoryDto struct {

	// disabled
	Disabled bool `json:"disabled"`

	// display name
	DisplayName string `json:"displayName"`

	// has catalog app
	HasCatalogApp bool `json:"hasCatalogApp"`

	// is bound
	IsBound bool `json:"isBound"`

	// name
	Name string `json:"name,omitempty"`

	// official
	Official bool `json:"official"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// repository Id
	RepositoryID string `json:"repositoryId,omitempty"`

	// true Url
	TrueURL string `json:"trueUrl,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// verified publisher
	VerifiedPublisher bool `json:"verifiedPublisher"`
}

// Validate validates this artifact repository dto
func (m *ArtifactRepositoryDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this artifact repository dto based on context it is used
func (m *ArtifactRepositoryDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactRepositoryDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactRepositoryDto) UnmarshalBinary(b []byte) error {
	var res ArtifactRepositoryDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
