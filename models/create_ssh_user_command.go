// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateSSHUserCommand create Ssh user command
//
// swagger:model CreateSshUserCommand
type CreateSSHUserCommand struct {

	// access profile Id
	AccessProfileID int32 `json:"accessProfileId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ssh public key
	SSHPublicKey string `json:"sshPublicKey,omitempty"`
}

// Validate validates this create Ssh user command
func (m *CreateSSHUserCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create Ssh user command based on context it is used
func (m *CreateSSHUserCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSSHUserCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSSHUserCommand) UnmarshalBinary(b []byte) error {
	var res CreateSSHUserCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
