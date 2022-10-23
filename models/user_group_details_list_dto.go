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
)

// UserGroupDetailsListDto user group details list dto
//
// swagger:model UserGroupDetailsListDto
type UserGroupDetailsListDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// project groups
	ProjectGroups []*UserGroupDetailsListDtoProjectGroupsItems0 `json:"projectGroups"`

	// users
	Users []*UserGroupDetailsListDtoUsersItems0 `json:"users"`
}

// Validate validates this user group details list dto
func (m *UserGroupDetailsListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGroupDetailsListDto) validateProjectGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectGroups); i++ {
		if swag.IsZero(m.ProjectGroups[i]) { // not required
			continue
		}

		if m.ProjectGroups[i] != nil {
			if err := m.ProjectGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserGroupDetailsListDto) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this user group details list dto based on the context it is used
func (m *UserGroupDetailsListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjectGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserGroupDetailsListDto) contextValidateProjectGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectGroups); i++ {

		if m.ProjectGroups[i] != nil {
			if err := m.ProjectGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserGroupDetailsListDto) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserGroupDetailsListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroupDetailsListDto) UnmarshalBinary(b []byte) error {
	var res UserGroupDetailsListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserGroupDetailsListDtoProjectGroupsItems0 user group details list dto project groups items0
//
// swagger:model UserGroupDetailsListDtoProjectGroupsItems0
type UserGroupDetailsListDtoProjectGroupsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user group details list dto project groups items0
func (m *UserGroupDetailsListDtoProjectGroupsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user group details list dto project groups items0 based on context it is used
func (m *UserGroupDetailsListDtoProjectGroupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGroupDetailsListDtoProjectGroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroupDetailsListDtoProjectGroupsItems0) UnmarshalBinary(b []byte) error {
	var res UserGroupDetailsListDtoProjectGroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserGroupDetailsListDtoUsersItems0 user group details list dto users items0
//
// swagger:model UserGroupDetailsListDtoUsersItems0
type UserGroupDetailsListDtoUsersItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user group details list dto users items0
func (m *UserGroupDetailsListDtoUsersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user group details list dto users items0 based on context it is used
func (m *UserGroupDetailsListDtoUsersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGroupDetailsListDtoUsersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroupDetailsListDtoUsersItems0) UnmarshalBinary(b []byte) error {
	var res UserGroupDetailsListDtoUsersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
