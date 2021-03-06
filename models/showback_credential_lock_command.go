// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ShowbackCredentialLockCommand showback credential lock command
//
// swagger:model ShowbackCredentialLockCommand
type ShowbackCredentialLockCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`
}

// Validate validates this showback credential lock command
func (m *ShowbackCredentialLockCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback credential lock command based on context it is used
func (m *ShowbackCredentialLockCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShowbackCredentialLockCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShowbackCredentialLockCommand) UnmarshalBinary(b []byte) error {
	var res ShowbackCredentialLockCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
