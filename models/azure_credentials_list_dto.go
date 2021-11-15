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

// AzureCredentialsListDto azure credentials list dto
//
// swagger:model AzureCredentialsListDto
type AzureCredentialsListDto struct {

	// availability zone
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project count
	ProjectCount int32 `json:"projectCount,omitempty"`

	// projects
	Projects []*CommonDropdownDto `json:"projects"`

	// tenant Id
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this azure credentials list dto
func (m *AzureCredentialsListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCredentialsListDto) validateProjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Projects) { // not required
		return nil
	}

	for i := 0; i < len(m.Projects); i++ {
		if swag.IsZero(m.Projects[i]) { // not required
			continue
		}

		if m.Projects[i] != nil {
			if err := m.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this azure credentials list dto based on the context it is used
func (m *AzureCredentialsListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureCredentialsListDto) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Projects); i++ {

		if m.Projects[i] != nil {
			if err := m.Projects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureCredentialsListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureCredentialsListDto) UnmarshalBinary(b []byte) error {
	var res AzureCredentialsListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
