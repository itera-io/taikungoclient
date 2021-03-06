// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteSSHUserCommand delete Ssh user command
//
// swagger:model DeleteSshUserCommand
type DeleteSSHUserCommand struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this delete Ssh user command
func (m *DeleteSSHUserCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete Ssh user command based on context it is used
func (m *DeleteSSHUserCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteSSHUserCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteSSHUserCommand) UnmarshalBinary(b []byte) error {
	var res DeleteSSHUserCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
