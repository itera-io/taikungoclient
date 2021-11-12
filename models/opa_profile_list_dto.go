// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpaProfileListDto opa profile list dto
//
// swagger:model OpaProfileListDto
type OpaProfileListDto struct {

	// allowed repo
	AllowedRepo []string `json:"allowedRepo"`

	// forbid Http ingress
	ForbidHTTPIngress bool `json:"forbidHttpIngress"`

	// forbid node port
	ForbidNodePort bool `json:"forbidNodePort"`

	// forbid specific tags
	ForbidSpecificTags []string `json:"forbidSpecificTags"`

	// id
	ID int32 `json:"id,omitempty"`

	// ingress whitelist
	IngressWhitelist []string `json:"ingressWhitelist"`

	// is default
	IsDefault bool `json:"isDefault"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// require probe
	RequireProbe bool `json:"requireProbe"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// unique ingresses
	UniqueIngresses bool `json:"uniqueIngresses"`

	// unique service selector
	UniqueServiceSelector bool `json:"uniqueServiceSelector"`
}

// Validate validates this opa profile list dto
func (m *OpaProfileListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this opa profile list dto based on context it is used
func (m *OpaProfileListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpaProfileListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpaProfileListDto) UnmarshalBinary(b []byte) error {
	var res OpaProfileListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
