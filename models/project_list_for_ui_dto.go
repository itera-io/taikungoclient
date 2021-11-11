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

// ProjectListForUIDto project list for Ui dto
//
// swagger:model ProjectListForUiDto
type ProjectListForUIDto struct {

	// alerts count
	AlertsCount int32 `json:"alertsCount,omitempty"`

	// bound users
	BoundUsers []*UserDto `json:"boundUsers"`

	// cloud credential name
	CloudCredentialName string `json:"cloudCredentialName,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// expired at
	ExpiredAt string `json:"expiredAt,omitempty"`

	// has kube config file
	HasKubeConfigFile bool `json:"hasKubeConfigFile,omitempty"`

	// health
	Health string `json:"health,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes,omitempty"`

	// is locked
	IsLocked bool `json:"isLocked,omitempty"`

	// kubernetes current version
	KubernetesCurrentVersion string `json:"kubernetesCurrentVersion,omitempty"`

	// kubernetes target version
	KubernetesTargetVersion string `json:"kubernetesTargetVersion,omitempty"`

	// kubespray current version
	KubesprayCurrentVersion string `json:"kubesprayCurrentVersion,omitempty"`

	// kubespray target version
	KubesprayTargetVersion string `json:"kubesprayTargetVersion,omitempty"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// total servers count
	TotalServersCount int32 `json:"totalServersCount,omitempty"`
}

// Validate validates this project list for Ui dto
func (m *ProjectListForUIDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListForUIDto) validateBoundUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.BoundUsers); i++ {
		if swag.IsZero(m.BoundUsers[i]) { // not required
			continue
		}

		if m.BoundUsers[i] != nil {
			if err := m.BoundUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this project list for Ui dto based on the context it is used
func (m *ProjectListForUIDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoundUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListForUIDto) contextValidateBoundUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BoundUsers); i++ {

		if m.BoundUsers[i] != nil {
			if err := m.BoundUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectListForUIDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectListForUIDto) UnmarshalBinary(b []byte) error {
	var res ProjectListForUIDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
