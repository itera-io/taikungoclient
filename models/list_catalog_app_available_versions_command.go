// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListCatalogAppAvailableVersionsCommand list catalog app available versions command
//
// swagger:model ListCatalogAppAvailableVersionsCommand
type ListCatalogAppAvailableVersionsCommand struct {

	// package name
	PackageName string `json:"packageName,omitempty"`

	// repo name
	RepoName string `json:"repoName,omitempty"`
}

// Validate validates this list catalog app available versions command
func (m *ListCatalogAppAvailableVersionsCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list catalog app available versions command based on context it is used
func (m *ListCatalogAppAvailableVersionsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListCatalogAppAvailableVersionsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCatalogAppAvailableVersionsCommand) UnmarshalBinary(b []byte) error {
	var res ListCatalogAppAvailableVersionsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
