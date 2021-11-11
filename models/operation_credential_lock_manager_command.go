// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OperationCredentialLockManagerCommand operation credential lock manager command
//
// swagger:model OperationCredentialLockManagerCommand
type OperationCredentialLockManagerCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`
}

// Validate validates this operation credential lock manager command
func (m *OperationCredentialLockManagerCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this operation credential lock manager command based on context it is used
func (m *OperationCredentialLockManagerCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperationCredentialLockManagerCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperationCredentialLockManagerCommand) UnmarshalBinary(b []byte) error {
	var res OperationCredentialLockManagerCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
