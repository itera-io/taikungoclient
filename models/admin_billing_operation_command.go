// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminBillingOperationCommand admin billing operation command
//
// swagger:model AdminBillingOperationCommand
type AdminBillingOperationCommand struct {

	// cloud credential Id
	CloudCredentialID int32 `json:"cloudCredentialId,omitempty"`
}

// Validate validates this admin billing operation command
func (m *AdminBillingOperationCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin billing operation command based on context it is used
func (m *AdminBillingOperationCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminBillingOperationCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminBillingOperationCommand) UnmarshalBinary(b []byte) error {
	var res AdminBillingOperationCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}