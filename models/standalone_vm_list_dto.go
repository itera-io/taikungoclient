// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StandaloneVMListDto standalone Vm list dto
//
// swagger:model StandaloneVmListDto
type StandaloneVMListDto struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// cpu
	CPU int32 `json:"cpu,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// flavor Id
	FlavorID string `json:"flavorId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// image name
	ImageName string `json:"imageName,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// is windows
	IsWindows bool `json:"isWindows"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// public Ip
	PublicIP string `json:"publicIp,omitempty"`

	// public Ip enabled
	PublicIPEnabled bool `json:"publicIpEnabled"`

	// ram
	RAM int64 `json:"ram,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// stand alone profile
	StandAloneProfile *StandaloneProfileListDto `json:"standAloneProfile,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// volume size
	VolumeSize int64 `json:"volumeSize,omitempty"`

	// volume type
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this standalone Vm list dto
func (m *StandaloneVMListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandAloneProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandaloneVMListDto) validateStandAloneProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.StandAloneProfile) { // not required
		return nil
	}

	if m.StandAloneProfile != nil {
		if err := m.StandAloneProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standAloneProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standAloneProfile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this standalone Vm list dto based on the context it is used
func (m *StandaloneVMListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStandAloneProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StandaloneVMListDto) contextValidateStandAloneProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.StandAloneProfile != nil {
		if err := m.StandAloneProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standAloneProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standAloneProfile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StandaloneVMListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandaloneVMListDto) UnmarshalBinary(b []byte) error {
	var res StandaloneVMListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
