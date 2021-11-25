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

// CreateStandAloneVMCommand create stand alone Vm command
//
// swagger:model CreateStandAloneVmCommand
type CreateStandAloneVMCommand struct {

	// count
	Count int32 `json:"count,omitempty"`

	// flavor name
	FlavorName string `json:"flavorName,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// public Ip enabled
	PublicIPEnabled bool `json:"publicIpEnabled"`

	// stand alone profile Id
	StandAloneProfileID int32 `json:"standAloneProfileId,omitempty"`

	// stand alone Vm disks
	StandAloneVMDisks []*StandAloneVMDiskDto `json:"standAloneVmDisks"`

	// volume size
	VolumeSize int64 `json:"volumeSize,omitempty"`

	// volume type
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this create stand alone Vm command
func (m *CreateStandAloneVMCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandAloneVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateStandAloneVMCommand) validateStandAloneVMDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.StandAloneVMDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.StandAloneVMDisks); i++ {
		if swag.IsZero(m.StandAloneVMDisks[i]) { // not required
			continue
		}

		if m.StandAloneVMDisks[i] != nil {
			if err := m.StandAloneVMDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standAloneVmDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standAloneVmDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create stand alone Vm command based on the context it is used
func (m *CreateStandAloneVMCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStandAloneVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateStandAloneVMCommand) contextValidateStandAloneVMDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StandAloneVMDisks); i++ {

		if m.StandAloneVMDisks[i] != nil {
			if err := m.StandAloneVMDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standAloneVmDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standAloneVmDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateStandAloneVMCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateStandAloneVMCommand) UnmarshalBinary(b []byte) error {
	var res CreateStandAloneVMCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
