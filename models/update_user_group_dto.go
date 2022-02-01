// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateUserGroupDto update user group dto
//
// swagger:model UpdateUserGroupDto
type UpdateUserGroupDto struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update user group dto
func (m *UpdateUserGroupDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update user group dto based on context it is used
func (m *UpdateUserGroupDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUserGroupDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUserGroupDto) UnmarshalBinary(b []byte) error {
	var res UpdateUserGroupDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
