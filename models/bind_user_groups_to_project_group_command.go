// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BindUserGroupsToProjectGroupCommand bind user groups to project group command
//
// swagger:model BindUserGroupsToProjectGroupCommand
type BindUserGroupsToProjectGroupCommand struct {

	// project group Id
	// Required: true
	// Minimum: > 0
	ProjectGroupID *int32 `json:"projectGroupId"`

	// project group name
	ProjectGroupName string `json:"projectGroupName,omitempty"`

	// user groups
	UserGroups []*UpdateProjectUserGroupDto `json:"userGroups"`
}

// Validate validates this bind user groups to project group command
func (m *BindUserGroupsToProjectGroupCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BindUserGroupsToProjectGroupCommand) validateProjectGroupID(formats strfmt.Registry) error {

	if err := validate.Required("projectGroupId", "body", m.ProjectGroupID); err != nil {
		return err
	}

	if err := validate.MinimumInt("projectGroupId", "body", int64(*m.ProjectGroupID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *BindUserGroupsToProjectGroupCommand) validateUserGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.UserGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.UserGroups); i++ {
		if swag.IsZero(m.UserGroups[i]) { // not required
			continue
		}

		if m.UserGroups[i] != nil {
			if err := m.UserGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bind user groups to project group command based on the context it is used
func (m *BindUserGroupsToProjectGroupCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BindUserGroupsToProjectGroupCommand) contextValidateUserGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserGroups); i++ {

		if m.UserGroups[i] != nil {
			if err := m.UserGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BindUserGroupsToProjectGroupCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BindUserGroupsToProjectGroupCommand) UnmarshalBinary(b []byte) error {
	var res BindUserGroupsToProjectGroupCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
