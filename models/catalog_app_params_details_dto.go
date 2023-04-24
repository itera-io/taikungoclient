// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogAppParamsDetailsDto catalog app params details dto
//
// swagger:model CatalogAppParamsDetailsDto
type CatalogAppParamsDetailsDto struct {

	// catalog app name
	CatalogAppName string `json:"catalogAppName,omitempty"`

	// has Json schema
	HasJSONSchema bool `json:"hasJsonSchema"`

	// id
	ID int32 `json:"id,omitempty"`

	// is editable after installation
	IsEditableAfterInstallation bool `json:"isEditableAfterInstallation"`

	// is editable when installing
	IsEditableWhenInstalling bool `json:"isEditableWhenInstalling"`

	// is mandatory
	IsMandatory bool `json:"isMandatory"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this catalog app params details dto
func (m *CatalogAppParamsDetailsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog app params details dto based on context it is used
func (m *CatalogAppParamsDetailsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogAppParamsDetailsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogAppParamsDetailsDto) UnmarshalBinary(b []byte) error {
	var res CatalogAppParamsDetailsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
