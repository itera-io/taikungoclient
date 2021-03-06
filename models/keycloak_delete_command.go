// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KeycloakDeleteCommand keycloak delete command
//
// swagger:model KeycloakDeleteCommand
type KeycloakDeleteCommand struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this keycloak delete command
func (m *KeycloakDeleteCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this keycloak delete command based on context it is used
func (m *KeycloakDeleteCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeycloakDeleteCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeycloakDeleteCommand) UnmarshalBinary(b []byte) error {
	var res KeycloakDeleteCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
