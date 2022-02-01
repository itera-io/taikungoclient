// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StandAloneVMDiskForDetailsDto stand alone Vm disk for details dto
//
// swagger:model StandAloneVmDiskForDetailsDto
type StandAloneVMDiskForDetailsDto struct {

	// current size
	CurrentSize int64 `json:"currentSize,omitempty"`

	// device name
	DeviceName string `json:"deviceName,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// lun Id
	LunID string `json:"lunId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// target size
	TargetSize int64 `json:"targetSize,omitempty"`

	// volume type
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this stand alone Vm disk for details dto
func (m *StandAloneVMDiskForDetailsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone Vm disk for details dto based on context it is used
func (m *StandAloneVMDiskForDetailsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StandAloneVMDiskForDetailsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandAloneVMDiskForDetailsDto) UnmarshalBinary(b []byte) error {
	var res StandAloneVMDiskForDetailsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
