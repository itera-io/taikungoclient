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

// KubernetesProfilesListDto kubernetes profiles list dto
//
// swagger:model KubernetesProfilesListDto
type KubernetesProfilesListDto struct {

	// allow scheduling on master
	AllowSchedulingOnMaster bool `json:"allowSchedulingOnMaster"`

	// cni
	Cni string `json:"cni,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// expose node port on bastion
	ExposeNodePortOnBastion bool `json:"exposeNodePortOnBastion"`

	// id
	ID int32 `json:"id,omitempty"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// octavia enabled
	OctaviaEnabled bool `json:"octaviaEnabled"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// projects
	Projects []*KubernetesProfilesListDtoProjectsItems0 `json:"projects"`

	// taikun l b enabled
	TaikunLBEnabled bool `json:"taikunLBEnabled"`

	// unique cluster name
	UniqueClusterName bool `json:"uniqueClusterName"`
}

// Validate validates this kubernetes profiles list dto
func (m *KubernetesProfilesListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesProfilesListDto) validateProjects(formats strfmt.Registry) error {
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

// ContextValidate validate this kubernetes profiles list dto based on the context it is used
func (m *KubernetesProfilesListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesProfilesListDto) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

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
func (m *KubernetesProfilesListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesProfilesListDto) UnmarshalBinary(b []byte) error {
	var res KubernetesProfilesListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubernetesProfilesListDtoProjectsItems0 kubernetes profiles list dto projects items0
//
// swagger:model KubernetesProfilesListDtoProjectsItems0
type KubernetesProfilesListDtoProjectsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this kubernetes profiles list dto projects items0
func (m *KubernetesProfilesListDtoProjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes profiles list dto projects items0 based on context it is used
func (m *KubernetesProfilesListDtoProjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesProfilesListDtoProjectsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesProfilesListDtoProjectsItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesProfilesListDtoProjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
