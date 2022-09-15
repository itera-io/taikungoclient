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

// ProjectListDetailDto project list detail dto
//
// swagger:model ProjectListDetailDto
type ProjectListDetailDto struct {

	// alerts count
	AlertsCount int32 `json:"alertsCount,omitempty"`

	// allow full spot kubernetes
	AllowFullSpotKubernetes bool `json:"allowFullSpotKubernetes"`

	// allow spot v ms
	AllowSpotVMs bool `json:"allowSpotVMs"`

	// allow spot workers
	AllowSpotWorkers bool `json:"allowSpotWorkers"`

	// bound users
	BoundUsers []*UserDto `json:"boundUsers"`

	// certification expired at
	CertificationExpiredAt string `json:"certificationExpiredAt,omitempty"`

	// cloud credential name
	CloudCredentialName string `json:"cloudCredentialName,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// delete on expiration
	DeleteOnExpiration bool `json:"deleteOnExpiration"`

	// expired at
	ExpiredAt string `json:"expiredAt,omitempty"`

	// has kube config file
	HasKubeConfigFile bool `json:"hasKubeConfigFile"`

	// health
	Health string `json:"health,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes"`

	// is locked
	IsLocked bool `json:"isLocked"`

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

	// max spot price
	MaxSpotPrice float64 `json:"maxSpotPrice,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// project action
	ProjectAction bool `json:"projectAction"`

	// quota Id
	QuotaID int32 `json:"quotaId,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// total servers count
	TotalServersCount int32 `json:"totalServersCount,omitempty"`

	// total standalone vms count
	TotalStandaloneVmsCount int32 `json:"totalStandaloneVmsCount,omitempty"`
}

// Validate validates this project list detail dto
func (m *ProjectListDetailDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListDetailDto) validateBoundUsers(formats strfmt.Registry) error {
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

// ContextValidate validate this project list detail dto based on the context it is used
func (m *ProjectListDetailDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoundUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectListDetailDto) contextValidateBoundUsers(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ProjectListDetailDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectListDetailDto) UnmarshalBinary(b []byte) error {
	var res ProjectListDetailDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
