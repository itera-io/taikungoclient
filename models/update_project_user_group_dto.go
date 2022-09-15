// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateProjectUserGroupDto update project user group dto
//
// swagger:model UpdateProjectUserGroupDto
type UpdateProjectUserGroupDto struct {

	// is bound
	IsBound bool `json:"isBound"`

	// user group Id
	UserGroupID int32 `json:"userGroupId,omitempty"`
}

// Validate validates this update project user group dto
func (m *UpdateProjectUserGroupDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update project user group dto based on context it is used
func (m *UpdateProjectUserGroupDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProjectUserGroupDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProjectUserGroupDto) UnmarshalBinary(b []byte) error {
	var res UpdateProjectUserGroupDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
